package io2

import (
	"compress/gzip"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	Day = time.Hour * 24
)

type OutputWriteAdapter func(time.Time) (WriteCloser, error)

// OutputFileAdapter 默认文件适配器
func OutputFileAdapter(prefix string, split time.Duration) OutputWriteAdapter {
	return func(now time.Time) (w WriteCloser, err error) {
		var file string
		if split == 0 {
			file = prefix
		} else {
			var date string
			if split%Day == 0 {
				date = now.Format("20060102")
			} else if split%time.Hour == 0 {
				date = now.Format("2006010215")
			} else if split%time.Minute == 0 {
				date = now.Format("200601021504")
			} else {
				date = now.Format("20060102150405")
			}
			if strings.Contains(prefix, "[date]") {
				file = strings.Replace(prefix, "[date]", date, 1)
			} else {
				file = prefix + "." + date
			}
		}
		if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
			return nil, err
		}
		return os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	}
}

type OutputClearAdapter func(time.Time) error

// OutputFileClearAdapter 默认文件清理适配器
func OutputFileClearAdapter(prefix string, expire time.Duration) OutputClearAdapter {
	if expire < 1 {
		return nil
	}
	return func(now time.Time) error {
		matches, err := filepath.Glob(prefix + "*")
		if err != nil {
			return err
		}
		for _, f := range matches {
			stat, _ := os.Stat(f)
			if stat == nil || stat.IsDir() {
				continue
			}
			if stat.ModTime().Add(expire).Before(now) {
				_ = os.Remove(f)
			}
		}
		return nil
	}
}

type OutputSplitAdapter func(context.Context, Splitter)

// OutputTickSplitAdapter 定时切割适配器
func OutputTickSplitAdapter(split time.Duration) OutputSplitAdapter {
	return func(ctx context.Context, s Splitter) {
		if split == 0 {
			return
		}
		var mod time.Duration
		now := time.Now()
		if split%Day == 0 {
			next := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(Day)
			mod = next.Sub(now)
		} else if split%(time.Hour) == 0 {
			next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.Local).Add(time.Hour)
			mod = next.Sub(now)
		} else if split%(time.Minute) == 0 {
			next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.Local).Add(time.Minute)
			mod = next.Sub(now)
		}
		if mod > 0 {
			time.Sleep(mod)
		}
		now = time.Now()
		s.DoSplit()
		s.DoClear(now)
		tk := time.NewTicker(split)
		for {
			select {
			case <-ctx.Done():
				tk.Stop()
				return
			case now = <-tk.C:
				s.DoSplit()
				s.DoClear(now)
			}
		}
	}
}

type Splitter interface {
	DoSplit()
	DoClear(now time.Time)
}

type OutputAdapter interface {
	WriteAdapter() OutputWriteAdapter
	ClearAdapter() OutputClearAdapter
	SplitAdapter() OutputSplitAdapter
}

// Output 输出到文件
type Output struct {
	ctx          context.Context
	writeAdapter OutputWriteAdapter
	clearAdapter OutputClearAdapter
	splitCancel  func()

	lock   sync.RWMutex
	writer WriteCloser
}

// NewFileOutput 构造输入文件结构
func NewFileOutput(ctx context.Context, path string, split, expire time.Duration) *Output {
	o := &Output{}
	o.ctx = ctx
	o.writeAdapter = OutputFileAdapter(path, split)
	o.clearAdapter = OutputFileClearAdapter(path, expire)
	if split > 0 {
		ctx, cancel := context.WithCancel(ctx)
		go OutputTickSplitAdapter(split)(ctx, o)
		o.splitCancel = cancel
	}
	return o
}

// NewCompressFileOutput 构造带压缩功能的文件输出结构
func NewCompressFileOutput(ctx context.Context, opts CompressFileAdapterOptions) *Output {
	adapter := NewCompressFileAdapter(opts)
	o := &Output{}
	o.ctx = ctx
	o.writeAdapter = adapter.WriteAdapter()
	o.clearAdapter = adapter.ClearAdapter()
	if opts.Split > 0 {
		ctx, cancel := context.WithCancel(ctx)
		go adapter.SplitAdapter()(ctx, o)
		o.splitCancel = cancel
	}
	return o
}

func (t *Output) SetAdapter(a OutputAdapter) {
	t.lock.Lock()
	t.writeAdapter = a.WriteAdapter()
	t.clearAdapter = a.ClearAdapter()
	oldCancel := t.splitCancel
	// 关闭旧的writer，下次GetWriter时会使用新的adapter创建
	if t.writer != nil {
		if err := t.writer.Close(); err != nil {
			println("[io2] SetAdapter, close writer: " + err.Error())
		}
		t.writer = nil
	}
	t.lock.Unlock()

	if oldCancel != nil {
		oldCancel()
	}

	ctx, cancel := context.WithCancel(t.ctx)
	go a.SplitAdapter()(ctx, t)

	t.lock.Lock()
	t.splitCancel = cancel
	t.lock.Unlock()
}

// SetWriteAdapter 设置写入适配器
func (t *Output) SetWriteAdapter(a OutputWriteAdapter) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.writeAdapter = a
	// 关闭旧的writer，下次GetWriter时会使用新的adapter创建
	if t.writer != nil {
		if err := t.writer.Close(); err != nil {
			println("[io2] SetWriteAdapter, close writer: " + err.Error())
		}
		t.writer = nil
	}
}

// SetClearAdapter 设置清理适配器
func (t *Output) SetClearAdapter(a OutputClearAdapter) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.clearAdapter = a
}

// SetSplitAdapter 设置切割适配器
func (t *Output) SetSplitAdapter(a OutputSplitAdapter) {
	t.lock.Lock()
	oldCancel := t.splitCancel
	t.lock.Unlock()

	if oldCancel != nil {
		oldCancel()
	}

	ctx, cancel := context.WithCancel(t.ctx)
	go a(ctx, t)

	t.lock.Lock()
	t.splitCancel = cancel
	t.lock.Unlock()
}

// DoSplit 执行切割
func (t *Output) DoSplit() {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.writer != nil {
		if err := t.writer.Close(); err != nil {
			println("[io2] DoSplit, writer close: " + err.Error())
		}
	}
	t.writer = nil
}

// DoClear 执行清理
func (t *Output) DoClear(now time.Time) {
	t.lock.RLock()
	adapter := t.clearAdapter
	t.lock.RUnlock()

	if adapter == nil {
		return
	}
	err := adapter(now)
	if err != nil {
		println("[io2] clear failure: " + err.Error())
	}
}

// GetWriter 读取文件句柄
func (t *Output) GetWriter() (Writer, error) {
	t.lock.RLock()
	if t.writer != nil {
		t.lock.RUnlock()
		return t.writer, nil
	}
	if t.writeAdapter == nil {
		t.lock.RUnlock()
		return os.Stderr, nil
	}
	t.lock.RUnlock()

	t.lock.Lock()
	defer t.lock.Unlock()
	// Double-check: 防止在获取写锁期间其他goroutine已经创建了writer
	if t.writer != nil {
		return t.writer, nil
	}
	if t.writeAdapter == nil {
		return os.Stderr, nil
	}
	w, err := t.writeAdapter(time.Now())
	if err == nil {
		t.writer = w
	}
	return w, err
}

func (t *Output) Write(b []byte) (int, error) {
	w, err := t.GetWriter()
	if err != nil {
		return 0, err
	}
	return w.Write(b)
}

// Close 关闭输出
func (t *Output) Close() error {
	// 先取消split goroutine，避免在关闭过程中继续执行切割操作
	t.lock.Lock()
	oldCancel := t.splitCancel
	t.splitCancel = nil
	t.lock.Unlock()

	if oldCancel != nil {
		oldCancel()
	}

	t.lock.Lock()
	defer t.lock.Unlock()
	t.writeAdapter = nil
	if t.writer != nil {
		err := t.writer.Close()
		t.writer = nil
		return err
	}
	return nil
}

// CompressFileAdapterOptions 压缩文件适配器选项
type CompressFileAdapterOptions struct {
	Prefix         string        // 文件前缀
	Split          time.Duration // 切割间隔
	Expire         time.Duration // 原始文件过期时间（0表示不删除原始文件）
	CompressAge    time.Duration // 文件达到此年龄后压缩（0表示不压缩）
	CompressExpire time.Duration // 压缩文件过期时间（0表示不删除压缩文件）
	KeepOriginal   bool          // 压缩后是否保留原始文件
	CompressLevel  int           // 压缩级别（-1=默认, 0=不压缩, 1=最快, 9=最优）
}

// CompressFileAdapter 带压缩功能的文件适配器
type CompressFileAdapter struct {
	opts CompressFileAdapterOptions
}

// NewCompressFileAdapter 创建压缩文件适配器
func NewCompressFileAdapter(opts CompressFileAdapterOptions) *CompressFileAdapter {
	if opts.CompressLevel < -1 || opts.CompressLevel > 9 {
		opts.CompressLevel = gzip.DefaultCompression
	}
	return &CompressFileAdapter{opts: opts}
}

// WriteAdapter 实现 OutputAdapter 接口
func (a *CompressFileAdapter) WriteAdapter() OutputWriteAdapter {
	return OutputFileAdapter(a.opts.Prefix, a.opts.Split)
}

// SplitAdapter 实现 OutputAdapter 接口
func (a *CompressFileAdapter) SplitAdapter() OutputSplitAdapter {
	return OutputTickSplitAdapter(a.opts.Split)
}

// ClearAdapter 实现 OutputAdapter 接口
func (a *CompressFileAdapter) ClearAdapter() OutputClearAdapter {
	if a.opts.CompressAge == 0 && a.opts.Expire == 0 && a.opts.CompressExpire == 0 {
		return nil
	}

	return func(now time.Time) error {
		matches, err := filepath.Glob(a.opts.Prefix + "*")
		if err != nil {
			return err
		}

		for _, f := range matches {
			stat, _ := os.Stat(f)
			if stat == nil || stat.IsDir() {
				continue
			}

			isCompressed := strings.HasSuffix(f, ".gz")
			age := now.Sub(stat.ModTime())

			// 处理压缩文件的过期删除
			if isCompressed && a.opts.CompressExpire > 0 {
				if age >= a.opts.CompressExpire {
					_ = os.Remove(f)
					continue
				}
			}

			// 处理原始文件
			if !isCompressed {
				// 压缩旧文件
				if a.opts.CompressAge > 0 && age >= a.opts.CompressAge {
					if err := a.compressFile(f); err != nil {
						println("[io2] compress file failed: " + f + ", error: " + err.Error())
					} else {
						// 压缩成功后，根据配置决定是否删除原文件
						if !a.opts.KeepOriginal {
							_ = os.Remove(f)
						}
					}
				}

				// 删除过期的原始文件
				if a.opts.Expire > 0 && age >= a.opts.Expire {
					_ = os.Remove(f)
				}
			}
		}

		return nil
	}
}

// compressFile 压缩单个文件
func (a *CompressFileAdapter) compressFile(filename string) error {
	// 检查压缩文件是否已存在
	gzFilename := filename + ".gz"
	if _, err := os.Stat(gzFilename); err == nil {
		return nil // 已经存在压缩文件，跳过
	}

	// 打开原始文件
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建压缩文件
	dst, err := os.Create(gzFilename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 创建gzip写入器
	gzWriter, err := gzip.NewWriterLevel(dst, a.opts.CompressLevel)
	if err != nil {
		return err
	}
	defer gzWriter.Close()

	// 复制数据并压缩
	if _, err := io.Copy(gzWriter, src); err != nil {
		return err
	}

	// 确保所有数据写入
	if err := gzWriter.Close(); err != nil {
		return err
	}

	// 保持原始文件的修改时间
	srcStat, _ := os.Stat(filename)
	if srcStat != nil {
		_ = os.Chtimes(gzFilename, srcStat.ModTime(), srcStat.ModTime())
	}

	return nil
}
