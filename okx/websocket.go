package okx

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/youjianglong/exchanges/okx/common"
	"github.com/youjianglong/exchanges/ws"

	"github.com/youjianglong/exchanges/errorx"
	"github.com/youjianglong/exchanges/mapx"
)

const (
	OpSubscribe   = "subscribe"
	OpUnsubscribe = "unsubscribe"
	OpLogin       = "login"

	EventSubscribe   = OpSubscribe
	EventUnsubscribe = OpUnsubscribe
	EventError       = "error"
)

type OperateRequest struct {
	Op   string `json:"op"`
	Args []any  `json:"args"`
}

type OperatePublicArg struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type H = map[string]any

type ArgsLogin struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

func NewArgsLoginFromAuth(auth *common.Auth) *ArgsLogin {
	signature := auth.Signature("GET", "/users/self/verify", "", true)
	return &ArgsLogin{
		ApiKey:     auth.ApiKey,
		Passphrase: auth.Passphrase,
		Sign:       signature.Build(),
		Timestamp:  signature.Timestamp,
	}
}

type Response struct {
	Event string `json:"event"`
	Arg   any    `json:"arg"`
	Code  string `json:"code"`
	Msg   string `json:"msg"`
}

func (resp Response) Error() error {
	if !resp.IsError() {
		return nil
	}
	return fmt.Errorf("code: %s, msg: %s", resp.Code, resp.Msg)
}

func (resp Response) IsError() bool {
	return resp.Event == EventError
}

type AdapterWebsocket struct {
	auth       *common.Auth
	handlers   *mapx.Map[string, func(*wsEvent)]
	logger     *slog.Logger
	ws         *ws.Websocket
	subscribes []any
	started    bool
}

func NewAdapterWebsocket(wsUrl string, auth *common.Auth, logger *slog.Logger, httpClient *http.Client) *AdapterWebsocket {
	a := &AdapterWebsocket{
		auth:     auth,
		handlers: mapx.New[string, func(*wsEvent)](),
		logger:   logger,
	}
	ws := ws.NewWebsocket(wsUrl, a.handleMsg, a.handleError, a.handlePing).
		SetStartConnect(a.startConnect)
	if httpClient != nil {
		ws.SetHttpClient(httpClient)
	}
	a.ws = ws
	return a
}

func (a *AdapterWebsocket) SetLogger(logIn func(msgType int, msg []byte), logOut func(msgType int, msg []byte)) *AdapterWebsocket {
	a.ws.SetLogger(logIn, logOut)
	return a
}

type wsEvent struct {
	Event  string `json:"event,omitempty"`
	ConnId string `json:"connId,omitempty"`
	Arg    struct {
		Channel  string `json:"channel"`
		InstType string `json:"instType,omitempty"`
		InstId   string `json:"instId,omitempty"`
		Ccy      string `json:"ccy,omitempty"`
	} `json:"arg"`
	Code      string          `json:"code,omitempty"`
	Msg       string          `json:"msg,omitempty"`
	EventType string          `json:"eventType,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

func (a *AdapterWebsocket) handleMsg(msg []byte) {
	var event wsEvent
	err := json.Unmarshal(msg, &event)
	if err != nil {
		a.logger.Error("unmarshal error: " + err.Error())
		return
	}
	if event.Event != "" {
		return
	}
	handler, ok := a.handlers.Get(event.Arg.Channel)
	if !ok {
		a.logger.Warn("handler not found: " + event.Arg.Channel)
		return
	}
	handler(&event)
}

func (a *AdapterWebsocket) handleError(err error) {
	fr := errorx.GetFrame(1)
	a.logger.Error(fmt.Sprintf("%v at %s:%d", err, fr.File, fr.Line))
}

func (a *AdapterWebsocket) handlePing(conn *ws.Websocket) {
	go func() {
		ticker := time.NewTicker(PingTimeout)
		defer ticker.Stop()
		for {
			select {
			case <-a.ws.Done():
				return
			case <-ticker.C:
				a.ws.LogOut(int(ws.MessagePing), PingMessage)
				if err := conn.Write(ws.MessagePing, PingMessage); err != nil {
					a.logger.Error("ping error: " + err.Error())
					conn.Stop()
					return
				}
			}
		}
	}()
}

func (a *AdapterWebsocket) RegisterHandler(channel string, handler func(*wsEvent)) *AdapterWebsocket {
	a.handlers.Set(channel, handler)
	return a
}

func (a *AdapterWebsocket) RegisterHandlers(handlers map[string]func(*wsEvent)) *AdapterWebsocket {
	for channel, handler := range handlers {
		a.RegisterHandler(channel, handler)
	}
	return a
}

func (a *AdapterWebsocket) Send(op string, args ...any) error {
	req := OperateRequest{
		Op:   op,
		Args: args,
	}
	return a.ws.WriteJSON(req)
}

func (a *AdapterWebsocket) Operate(op string, args ...any) error {
	req := OperateRequest{
		Op:   op,
		Args: args,
	}
	err := a.ws.WriteJSON(req)
	if err != nil {
		return err
	}
	var resp Response
	err = a.ws.ReadJSON(&resp)
	if err != nil {
		return err
	}
	return resp.Error()
}

func (a *AdapterWebsocket) Login() error {
	return a.Operate(OpLogin, NewArgsLoginFromAuth(a.auth))
}

func (a *AdapterWebsocket) Subscribe(args ...any) error {
	a.subscribes = append(a.subscribes, args...)
	if a.started {
		return a.Operate(OpSubscribe, args...)
	}
	return nil
}

func (a *AdapterWebsocket) startConnect(ws *ws.Websocket, conn *ws.Conn) error {
	a.started = true
	if a.auth != nil {
		err := a.Login()
		if err != nil {
			return err
		}
	}
	if len(a.subscribes) > 0 {
		err := a.Send(OpSubscribe, a.subscribes...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AdapterWebsocket) Start() error {
	return a.ws.Start()
}

func (a *AdapterWebsocket) Stop() {
	a.started = false
	a.ws.Stop()
}
