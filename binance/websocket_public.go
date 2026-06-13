package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	. "github.com/youjianglong/exchanges/common"

	"github.com/youjianglong/exchanges/errorx"

	"github.com/youjianglong/exchanges/ws"
)

// WsPublicBaseService 现货公共行情WebSocket服务
type WsPublicBaseService struct {
	name         string
	baseUrl      string
	httpClient   *http.Client
	logger       *slog.Logger
	logIn        func(int, []byte)
	logOut       func(int, []byte)
	KeepInterval time.Duration
	PongTimeout  time.Duration
	channels     map[string]ws.WsHandler
	channelMu    sync.RWMutex
	ws           *ws.Websocket
	idx          int64
}

func (s *WsPublicBaseService) SetHttpClient(httpClient *http.Client) *WsPublicBaseService {
	s.httpClient = httpClient
	if s.ws != nil {
		s.ws.SetHttpClient(httpClient)
	}
	return s
}

func (s *WsPublicBaseService) SetLogger(logIn func(int, []byte), logOut func(int, []byte)) *WsPublicBaseService {
	s.logIn = logIn
	s.logOut = logOut
	if s.ws != nil {
		s.ws.SetLogger(s.logIn, s.logOut)
	}
	return s
}

func (s *WsPublicBaseService) Subscribe(channel string, handler ws.WsHandler) *WsPublicBaseService {
	s.channelMu.Lock()
	s.channels[channel] = handler
	s.channelMu.Unlock()
	if s.ws != nil {
		s.subscribe(channel)
	}
	return s
}

func (s *WsPublicBaseService) Unsubscribe(channel string) *WsPublicBaseService {
	s.channelMu.Lock()
	delete(s.channels, channel)
	s.channelMu.Unlock()
	if s.ws != nil {
		s.unsubscribe(channel)
	}
	return s
}

func (s *WsPublicBaseService) IncrIdx() int64 {
	return atomic.AddInt64(&s.idx, 1)
}

func (s *WsPublicBaseService) subscribe(channels ...string) {
	data := map[string]any{
		"method": "SUBSCRIBE",
		"params": channels,
		"id":     s.IncrIdx(),
	}
	err := s.ws.WriteJSON(data)
	if err != nil {
		s.logger.Error(fmt.Sprintf("write json: %v", err))
		return
	}
}

func (s *WsPublicBaseService) unsubscribe(channels ...string) {
	data := map[string]any{
		"method": "UNSUBSCRIBE",
		"params": channels,
		"id":     s.IncrIdx(),
	}
	err := s.ws.WriteJSON(data)
	if err != nil {
		s.logger.Error(fmt.Sprintf("write json: %v", err))
		return
	}
}

func (s *WsPublicBaseService) Start() error {
	if s.ws != nil {
		s.ws.Stop()
		s.ws = nil
	}
	s.ws = ws.NewWebsocket(s.baseUrl, s.handleMsg, s.handleError, s.keepAlive).
		SetPrevConnect(s.prevConnect).
		SetHttpClient(s.httpClient).
		SetLogger(s.logIn, s.logOut)
	return s.ws.Start()
}

// prevConnect 设置连接前的操作
func (s *WsPublicBaseService) prevConnect(ws *ws.Websocket) error {
	s.logger.Info("init connect: " + s.name)
	var streams []string
	s.channelMu.RLock()
	for stream := range s.channels {
		streams = append(streams, stream)
	}
	s.channelMu.RUnlock()
	endpoint := fmt.Sprintf("%s%s", s.baseUrl, strings.Join(streams, "/"))
	ws.SetEndpoint(endpoint)
	return nil
}

type combinedEvent struct {
	Stream string          `json:"stream,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`

	Id     int             `json:"id,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func (s *WsPublicBaseService) handleMsg(msg []byte) {
	var event combinedEvent
	err := StrictDecode(msg, &event)
	if err != nil {
		s.logger.Error(fmt.Sprintf("unmarshal message: %v", err))
		return
	}
	if event.Stream != "" {
		s.channelMu.RLock()
		h, ok := s.channels[event.Stream]
		s.channelMu.RUnlock()
		if !ok {
			s.logger.Error(fmt.Sprintf("channel %s not found", event.Stream))
			return
		}
		h(event.Data)
		return
	}
	// if event.Id != 0 {
	// 	s.logger.Info(fmt.Sprintf("id: %d, result: %s", event.Id, string(event.Result)))
	// }
}

func (s *WsPublicBaseService) handleError(err error) {
	fr := errorx.GetFrame(1)
	s.logger.Error(fmt.Sprintf("%v at %s:%d", err, fr.File, fr.Line))
}

func (s *WsPublicBaseService) keepAlive(w *ws.Websocket) {
	w.SetOnPingReceived(func(ctx context.Context, message []byte) bool {
		if s.logIn != nil {
			s.logIn(int(ws.MessagePing), message)
		}
		return true
	})

	w.SetOnPongReceived(func(ctx context.Context, message []byte) {
		if s.logIn != nil {
			s.logIn(int(ws.MessagePong), message)
		}
	})
}

func (s *WsPublicBaseService) Stop() {
	if s.ws != nil {
		s.ws.Stop()
		s.ws = nil
	}
}

func getTypeName[T any]() string {
	var t T
	return reflect.TypeOf(t).Name()
}

func wsHandleWrapper[T any](logger *slog.Logger, handler func(event T)) func(data []byte) {
	name := getTypeName[T]()
	return func(data []byte) {
		var event T
		err := StrictDecode(data, &event)
		if err != nil {
			logger.Error(fmt.Sprintf("unmarshal %s message: %v", name, err))
			return
		}
		handler(event)
	}
}

type WsSpotPublicService struct {
	*WsPublicBaseService
}

func NewWsSpotPublicService(baseUrl string) WsSpotPublicService {
	if baseUrl == "" {
		baseUrl = BaseSpotCombinedMainURL
	}
	s := &WsPublicBaseService{
		name:         "spot_public",
		baseUrl:      baseUrl,
		logger:       slog.Default(),
		channels:     make(map[string]ws.WsHandler),
		KeepInterval: 60 * time.Second,
		PongTimeout:  10 * time.Second,
	}
	return WsSpotPublicService{s}
}

// WsMiniTickerEvent define websocket market mini-ticker statistics event
type WsMiniTickerEvent struct {
	Event       string  `json:"e"` // 事件类型
	Time        Int64   `json:"E"` // 事件时间
	Symbol      string  `json:"s"` // 交易对
	LastPrice   Float64 `json:"c"` // 最新价格
	OpenPrice   Float64 `json:"o"` // 开盘价格
	HighPrice   Float64 `json:"h"` // 最高价格
	LowPrice    Float64 `json:"l"` // 最低价格
	BaseVolume  Float64 `json:"v"` // 成交量
	QuoteVolume Float64 `json:"q"` // 成交额
}

// WsAllMiniTickerEvent define array of websocket market mini-ticker statistics events
type WsAllMiniTickerEvent []*WsMiniTickerEvent

func (s WsSpotPublicService) SubscribeAllMiniTicker(handler func(event WsAllMiniTickerEvent)) {
	s.Subscribe("!miniTicker@arr", func(data []byte) {
		var event WsAllMiniTickerEvent
		err := StrictDecode(data, &event)
		if err != nil {
			s.logger.Error(fmt.Sprintf("unmarshal message: %v, data: %s", err, string(data)))
			return
		}
		handler(event)
	})
}

func (s WsSpotPublicService) SubscribeSymbolsMiniTicker(handler func(event WsMiniTickerEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@miniTicker", strings.ToLower(symbol)), fn)
	}
}

// WsTickerEvent define websocket market statistics event
type WsTickerEvent struct {
	Event              string `json:"e"` // 事件类型
	Time               int64  `json:"E"` // 事件时间
	Symbol             string `json:"s"` // 交易对
	PriceChange        string `json:"p"` // 24小时价格变化
	PriceChangePercent string `json:"P"` // 24小时价格变化（百分比）
	WeightedAvgPrice   string `json:"w"` // 平均价格
	PrevClosePrice     string `json:"x"` // 整整24小时之前，向前数的最后一次成交价格
	LastPrice          string `json:"c"` // 最新成交价格
	CloseQty           string `json:"Q"` // 最新成交交易的成交量
	BidPrice           string `json:"b"` // 目前最高买单价
	BidQty             string `json:"B"` // 目前最高买单价的挂单量
	AskPrice           string `json:"a"` // 目前最低卖单价
	AskQty             string `json:"A"` // 目前最低卖单价的挂单量
	OpenPrice          string `json:"o"` // 开盘价
	HighPrice          string `json:"h"` // 24小时内最高成交价
	LowPrice           string `json:"l"` // 24小时内最低成交价
	BaseVolume         string `json:"v"` // 24小时内成交量
	QuoteVolume        string `json:"q"` // 24小时内成交额
	OpenTime           int64  `json:"O"` // 统计开始时间
	CloseTime          int64  `json:"C"` // 统计结束时间
	FirstID            int64  `json:"F"` // 24小时内第一笔成交交易ID
	LastID             int64  `json:"L"` // 24小时内最后一笔成交交易ID
	Count              int64  `json:"n"` // 24小时内成交数
}

// WsAllTickerEvent define array of websocket market statistics events
type WsAllTickerEvent []*WsTickerEvent

func (s WsSpotPublicService) SubscribeAllTicker(handler func(event WsAllTickerEvent)) {
	s.Subscribe("!ticker@arr", wsHandleWrapper(s.logger, handler))
}

func (s WsSpotPublicService) SubscribeSymbolsTicker(handler func(event WsTickerEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@ticker", strings.ToLower(symbol)), fn)
	}
}

type WsBookTickerEvent struct {
	UpdateID Int64   `json:"u"` // 更新ID
	Symbol   string  `json:"s"` // 交易对
	BidPrice Float64 `json:"b"` // 最高买价
	BidQty   Float64 `json:"B"` // 最高买价挂单量
	AskPrice Float64 `json:"a"` // 最低卖价
	AskQty   Float64 `json:"A"` // 最低卖价挂单量
	Time     Int64   `json:"E"` // 事件时间
}

// SubscribeSymbolsBookTicker 订阅指定交易对的盘口信息
func (s WsSpotPublicService) SubscribeSymbolsBookTicker(handler func(event WsBookTickerEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@bookTicker", strings.ToLower(symbol)), fn)
	}
}

type WsSwapPublicService struct {
	*WsPublicBaseService
}

func NewWsSwapPublicService(baseUrl string) WsSwapPublicService {
	if baseUrl == "" {
		baseUrl = BaseSwapCombinedWsMainURL
	}
	s := &WsPublicBaseService{
		name:         "swap_public",
		baseUrl:      baseUrl,
		logger:       slog.Default(),
		channels:     make(map[string]ws.WsHandler),
		KeepInterval: 60 * time.Second,
		PongTimeout:  10 * time.Second,
	}
	return WsSwapPublicService{s}
}

type MarkPriceEvent struct {
	Event           string  `json:"e"` // 事件类型
	Time            int64   `json:"E"` // 事件时间
	Symbol          string  `json:"s"` // 交易对
	MarkPrice       Float64 `json:"p"` // 标记价格
	IndexPrice      Float64 `json:"i"` // 现货指数价格
	EstimatedSettle Float64 `json:"P"` // 预估结算价格
	FundingRate     Float64 `json:"r"` // 资金费率
	NextFundingTime int64   `json:"T"` // 下一个资金费率时间
}

// SubscribeMarkPrice 订阅标记价格
func (s WsSwapPublicService) SubscribeMarkPrice(handler func(event MarkPriceEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@markPrice", strings.ToLower(symbol)), fn)
	}
}

// SubscribeMarkPrice1s 订阅标记价格
func (s WsSwapPublicService) SubscribeMarkPrice1s(handler func(event MarkPriceEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@markPrice@1s", strings.ToLower(symbol)), fn)
	}
}

type WsAllMarkPriceEvent []*MarkPriceEvent

// SubscribeAllMarkPrice 订阅全市场标记价格
func (s WsSwapPublicService) SubscribeAllMarkPrice(handler func(event WsAllMarkPriceEvent)) {
	s.Subscribe("!markPrice@arr", wsHandleWrapper(s.logger, handler))
}

// SubscribeSymbolsMiniTicker 订阅指定交易对的简易信息
func (s WsSwapPublicService) SubscribeSymbolsMiniTicker(handler func(event WsMiniTickerEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@miniTicker", strings.ToLower(symbol)), fn)
	}
}

// SubscribeAllMiniTicker 订阅全市场简易信息
func (s WsSwapPublicService) SubscribeAllMiniTicker(handler func(event WsAllMiniTickerEvent)) {
	s.Subscribe("!miniTicker@arr", wsHandleWrapper(s.logger, handler))
}

func (s WsSwapPublicService) SubscribeSymbolsTicker(handler func(event WsTickerEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@ticker", strings.ToLower(symbol)), fn)
	}
}

// SubscribeAllTicker 订阅全市场信息
func (s WsSwapPublicService) SubscribeAllTicker(handler func(event WsAllTickerEvent)) {
	s.Subscribe("!ticker@arr", wsHandleWrapper(s.logger, handler))
}

// SubscribeBookTicker 订阅指定交易对的盘口信息
func (s WsSwapPublicService) SubscribeBookTicker(handler func(event WsBookTickerEvent), symbols ...string) {
	fn := wsHandleWrapper(s.logger, handler)
	for _, symbol := range symbols {
		s.Subscribe(fmt.Sprintf("%s@bookTicker", strings.ToLower(symbol)), fn)
	}
}
