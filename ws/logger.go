package ws

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/youjianglong/exchanges/io2"
)

var (
	EOL        = []byte("\n")
	inFlag     = []byte("I: ")
	outFlag    = []byte("O: ")
	msgTypeMap = map[int][]byte{
		int(MessageText):   []byte("[T] "),
		int(MessageBinary): []byte("[B] "),
		int(MessagePing):   []byte("[PI] "),
		int(MessagePong):   []byte("[PO] "),
	}
)

func GetLogger(w io.Writer) (in func(int, []byte), out func(int, []byte)) {
	mu := sync.Mutex{}

	in = func(msgType int, msg []byte) {
		mu.Lock()
		defer mu.Unlock()
		w.Write([]byte(time.Now().Format("01/02 15:04:05 ")))
		_, _ = w.Write(inFlag)
		if len(msg) == 0 {
			v := msgTypeMap[msgType]
			_, _ = w.Write(v[:len(v)-1])
		} else {
			_, _ = w.Write(msgTypeMap[msgType])
			_, _ = w.Write(msg)
		}
		_, _ = w.Write(EOL)
	}
	out = func(msgType int, msg []byte) {
		mu.Lock()
		defer mu.Unlock()
		w.Write([]byte(time.Now().Format("01/02 15:04:05 ")))
		_, _ = w.Write(outFlag)
		if len(msg) == 0 {
			v := msgTypeMap[msgType]
			_, _ = w.Write(v[:len(v)-1])
		} else {
			_, _ = w.Write(msgTypeMap[msgType])
			_, _ = w.Write(msg)
		}
		_, _ = w.Write(EOL)
	}
	return in, out
}

func GetFileLogger(dir string, name string, split time.Duration, expire time.Duration) (in func(int, []byte), out func(int, []byte)) {
	_ = os.MkdirAll(dir, 0644)
	path := fmt.Sprintf("%s/%s.log", dir, name)
	fo := io2.NewFileOutput(context.Background(), path, split, expire)
	return GetLogger(fo)
}
