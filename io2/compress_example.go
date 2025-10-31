package io2

import (
	"context"
	"time"
)

// 使用示例（仅供参考）

// Example1: 基本使用 - 按天切割，1天后压缩，7天后删除压缩文件
func exampleBasicCompress() {
	ctx := context.Background()

	output := NewCompressFileOutput(ctx, CompressFileAdapterOptions{
		Prefix:         "./logs/app.log",
		Split:          Day,     // 按天切割
		CompressAge:    Day,     // 1天后压缩
		CompressExpire: 7 * Day, // 压缩文件保留7天
		KeepOriginal:   false,   // 压缩后删除原文件
		CompressLevel:  -1,      // 使用默认压缩级别
	})
	defer output.Close()

	// 写入日志
	output.Write([]byte("test log\n"))
}

// Example2: 保留原文件 - 压缩后保留原文件，原文件3天后删除
func exampleKeepOriginal() {
	ctx := context.Background()

	output := NewCompressFileOutput(ctx, CompressFileAdapterOptions{
		Prefix:         "./logs/data.log",
		Split:          time.Hour, // 按小时切割
		Expire:         3 * Day,   // 原文件3天后删除
		CompressAge:    time.Hour, // 1小时后压缩
		CompressExpire: 30 * Day,  // 压缩文件保留30天
		KeepOriginal:   true,      // 压缩后保留原文件
		CompressLevel:  9,         // 最优压缩
	})
	defer output.Close()

	output.Write([]byte("test data\n"))
}

// Example3: 仅压缩不删除 - 只压缩文件，不删除任何文件
func exampleCompressOnly() {
	ctx := context.Background()

	output := NewCompressFileOutput(ctx, CompressFileAdapterOptions{
		Prefix:         "./logs/archive.log",
		Split:          Day,     // 按天切割
		CompressAge:    2 * Day, // 2天后压缩
		CompressExpire: 0,       // 不删除压缩文件
		Expire:         0,       // 不删除原文件
		KeepOriginal:   false,   // 压缩后删除原文件
		CompressLevel:  6,       // 平衡压缩
	})
	defer output.Close()

	output.Write([]byte("test archive\n"))
}

// Example4: 使用自定义适配器
func exampleCustomAdapter() {
	ctx := context.Background()

	adapter := NewCompressFileAdapter(CompressFileAdapterOptions{
		Prefix:         "./logs/custom.log",
		Split:          Day,
		CompressAge:    Day,
		CompressExpire: 7 * Day,
		KeepOriginal:   false,
		CompressLevel:  -1,
	})

	output := &Output{ctx: ctx}
	output.SetAdapter(adapter)
	defer output.Close()

	output.Write([]byte("test custom\n"))
}

// Example5: 高频切割 - 按分钟切割，适用于高频日志
func exampleHighFrequency() {
	ctx := context.Background()

	output := NewCompressFileOutput(ctx, CompressFileAdapterOptions{
		Prefix:         "./logs/high-freq.log",
		Split:          time.Minute,      // 按分钟切割
		CompressAge:    10 * time.Minute, // 10分钟后压缩
		CompressExpire: time.Hour,        // 压缩文件1小时后删除
		KeepOriginal:   false,
		CompressLevel:  1, // 快速压缩
	})
	defer output.Close()

	output.Write([]byte("high frequency log\n"))
}
