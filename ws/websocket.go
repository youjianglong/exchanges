package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/youjianglong/exchanges/errorx"
	"github.com/youjianglong/exchanges/types"

	"github.com/coder/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

type MessageType = websocket.MessageType

type Conn = websocket.Conn

var (
	MessageText   = websocket.MessageText
	MessageBinary = websocket.MessageBinary
	MessagePing   = websocket.MessageBinary + 1
	MessagePong   = websocket.MessageBinary + 2
)

// ErrServiceStopped 服务停止
var ErrServiceStopped = errors.New("service stopped")

type Websocket struct {
	HandshakeTimeout time.Duration
	AutoReconnect    bool

	logIn  func(int, []byte) // 输入日志
	logOut func(int, []byte) // 输出日志

	endpoint   string
	httpClient *http.Client

	wsHandler  WsHandler
	errHandler ErrHandler
	keepAlive  func(*Websocket)
	waitC      *atomic.Value
	conn       *Conn

	prevConnect    func(*Websocket) error
	startConnect   func(*Websocket, *Conn) error
	onPingReceived func(context.Context, []byte) bool
	onPongReceived func(context.Context, []byte)

	stopC    chan types.Zero
	doneC    chan types.Zero
	doneFlag *atomic.Bool

	failCount time.Duration
}

func NewWebsocket(endpoint string, handler WsHandler, errHandler ErrHandler, keepAlive func(*Websocket)) *Websocket {
	waitC := &atomic.Value{}
	waitC.Store(make(chan struct{}))
	w := &Websocket{
		HandshakeTimeout: 45 * time.Second,
		AutoReconnect:    true,
		endpoint:         endpoint,
		httpClient:       &http.Client{Timeout: 10 * time.Second},
		wsHandler:        handler,
		errHandler:       errHandler,
		keepAlive:        keepAlive,
		waitC:            waitC,
		stopC:            make(chan types.Zero),
		doneC:            make(chan types.Zero),
		doneFlag:         &atomic.Bool{},
	}
	return w
}

func (w *Websocket) SetEndpoint(endpoint string) *Websocket {
	w.endpoint = endpoint
	return w
}

func (w *Websocket) SetHttpClient(httpClient *http.Client) *Websocket {
	w.httpClient = httpClient
	return w
}

func (w *Websocket) SetWsHandler(handler WsHandler) *Websocket {
	w.wsHandler = handler
	return w
}

func (w *Websocket) SetErrHandler(handler ErrHandler) *Websocket {
	w.errHandler = handler
	return w
}

func (w *Websocket) SetPrevConnect(handler func(*Websocket) error) *Websocket {
	w.prevConnect = handler
	return w
}

func (w *Websocket) SetStartConnect(handler func(*Websocket, *websocket.Conn) error) *Websocket {
	w.startConnect = handler
	return w
}

func (w *Websocket) SetOnPingReceived(handler func(context.Context, []byte) bool) *Websocket {
	w.onPingReceived = handler
	return w
}

func (w *Websocket) SetOnPongReceived(handler func(context.Context, []byte)) *Websocket {
	w.onPongReceived = handler
	return w
}

func (w *Websocket) SetLogger(logIn func(int, []byte), logOut func(int, []byte)) *Websocket {
	w.logIn = logIn
	w.logOut = logOut
	return w
}

func (w *Websocket) closeDoneChan() {
	if w.doneFlag.CompareAndSwap(false, true) {
		ch := w.doneC
		w.doneC = types.ClosedChan
		if ch != types.ClosedChan {
			close(ch)
		}
	}
}

func (w *Websocket) resetDoneChan() {
	if w.doneFlag.CompareAndSwap(true, false) {
		w.doneC = make(chan types.Zero)
	}
}

func (w *Websocket) serve() (err error) {
	if w.prevConnect != nil {
		err = w.prevConnect(w)
		if err != nil {
			if errors.Is(err, ErrServiceStopped) {
				return
			}
			select {
			case <-w.stopC:
				return
			default:
			}
			if w.errHandler != nil {
				w.errHandler(err)
			} else {
				slog.Error(fmt.Sprintf("prev connect failed: %v", err))
			}
			w.failCount++
			time.Sleep(time.Second * 2 * w.failCount)
			return w.serve()
		}
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), w.HandshakeTimeout)
	opts := websocket.DialOptions{
		HTTPClient:      w.httpClient,
		CompressionMode: websocket.CompressionNoContextTakeover,
		OnPingReceived:  w.handlePing,
		OnPongReceived:  w.handlePong,
	}
	c, _, err := websocket.Dial(timeoutCtx, w.endpoint, &opts)
	cancel()

	if err != nil {
		if w.AutoReconnect {
			if w.errHandler != nil {
				w.errHandler(err)
			} else {
				slog.Error(fmt.Sprintf("dial: %v", err))
			}
			w.failCount++
			time.Sleep(time.Second * w.failCount)
			return w.serve()
		}
		return err
	}

	w.conn = c
	w.failCount = 0

	// 关闭waitC
	ch, _ := w.waitC.Swap(types.ClosedChan).(chan types.Zero)
	if ch != nil && ch != types.ClosedChan {
		close(ch)
	}

	w.resetDoneChan()

	go func() {
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		stopped := false

		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer func() {
			re := recover()
			if re != nil {
				if w.errHandler != nil {
					err, ok := re.(error)
					if !ok {
						err = fmt.Errorf("panic: %v", re)
					}
					w.errHandler(err)
				} else {
					stacks := errorx.FormatFrames(errorx.GetFrames(1), 2)
					slog.Error(fmt.Sprintf("panic: %v\n%s", re, stacks))
				}
			}
			w.closeDoneChan()
			if w.AutoReconnect && !stopped { // 自动重连
				time.Sleep(time.Second * 2)
				w.waitC.Store(make(chan types.Zero))
				err := w.serve()
				if err != nil {
					if w.errHandler != nil {
						w.errHandler(err)
					} else {
						slog.Error(fmt.Sprintf("reconnect: %v", err))
					}
				}
			}
		}()

		if w.startConnect != nil {
			err = w.startConnect(w, c)
			if err != nil {
				if w.errHandler != nil {
					w.errHandler(err)
				} else {
					slog.Error(fmt.Sprintf("start connect: %v", err))
				}
				return
			}
		}

		if w.keepAlive != nil {
			// This function overwrites the default ping frame handler
			// sent by the websocket API server
			w.keepAlive(w)
		}

		go func() {
			select {
			case <-w.stopC:
				stopped = true
			case <-w.doneC:
			}
			_ = c.Close(websocket.StatusNormalClosure, "")
		}()
		for {
			msgType, message, err := c.Read(context.Background())
			if err != nil {
				if !stopped {
					if w.errHandler != nil {
						w.errHandler(err)
					} else {
						slog.Error(fmt.Sprintf("read message: %v", err))
					}
				}
				return
			}
			if w.logIn != nil {
				w.logIn(int(msgType), message)
			}
			if msgType == websocket.MessageText || msgType == websocket.MessageBinary {
				w.wsHandler(message)
			}
		}
	}()
	return
}

func (w *Websocket) handlePing(ctx context.Context, message []byte) bool {
	if w.onPingReceived != nil {
		return w.onPingReceived(ctx, message)
	}
	return true
}

func (w *Websocket) handlePong(ctx context.Context, message []byte) {
	if w.onPongReceived != nil {
		w.onPongReceived(ctx, message)
	}
}

func (w *Websocket) WaitReady() <-chan struct{} {
	ch, _ := w.waitC.Load().(chan types.Zero)
	if ch != nil {
		return ch
	}
	return types.ClosedChan
}

var ErrWebsocketNotConnected = errors.New("websocket connection is nil")

func (w *Websocket) LogOut(msgType int, message []byte) {
	if w.logOut != nil {
		w.logOut(msgType, message)
	}
}

func (w *Websocket) LogIn(msgType int, message []byte) {
	if w.logIn != nil {
		w.logIn(msgType, message)
	}
}

func (w *Websocket) Write(msgType MessageType, message []byte) error {
	<-w.WaitReady()
	if w.conn == nil {
		return ErrWebsocketNotConnected
	}
	if w.logOut != nil {
		w.logOut(int(msgType), message)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return w.conn.Write(ctx, msgType, message)
}

func (w *Websocket) WriteJSON(v any) error {
	<-w.WaitReady()
	if w.conn == nil {
		return ErrWebsocketNotConnected
	}
	message, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if w.logOut != nil {
		w.logOut(int(websocket.MessageText), message)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return w.conn.Write(ctx, websocket.MessageText, message)
}

func (w *Websocket) Read() (int, []byte, error) {
	<-w.WaitReady()
	if w.conn == nil {
		return 0, nil, ErrWebsocketNotConnected
	}
	msgType, message, err := w.conn.Read(context.Background())
	if err != nil {
		return 0, nil, err
	}
	if w.logIn != nil {
		w.logIn(int(msgType), message)
	}
	return int(msgType), message, nil
}

func (w *Websocket) ReadJSON(v any) error {
	<-w.WaitReady()
	if w.conn == nil {
		return ErrWebsocketNotConnected
	}

	msgType, message, err := w.conn.Read(context.Background())
	if err != nil {
		return err
	}
	if w.logIn != nil {
		w.logIn(int(msgType), message)
	}
	return json.Unmarshal(message, v)
}

func (w *Websocket) Conn() *websocket.Conn {
	<-w.WaitReady()
	return w.conn
}

func (w *Websocket) Done() <-chan struct{} {
	return w.doneC
}

func (w *Websocket) Start() error {
	return w.serve()
}

func (w *Websocket) Stop() {
	ch := w.stopC
	w.stopC = types.ClosedChan
	if ch != types.ClosedChan {
		close(ch)
	}
}

func (w *Websocket) Restart() {
	if w.conn != nil {
		_ = w.conn.Close(websocket.StatusNormalClosure, "")
	}
}
