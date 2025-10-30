package binance

import (
	"context"
	"encoding/json"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// 枚举定义 https://developers.binance.com/docs/zh-CN/binance-spot-api-docs/enums

type GetSpotOrdersService struct {
	c *Client

	symbol    *string
	orderId   *string
	startTime *int64
	endTime   *int64
	limit     *int
}

func (c *Client) NewGetSpotOrdersService() *GetSpotOrdersService {
	return &GetSpotOrdersService{c: c}
}

func (s *GetSpotOrdersService) Symbol(symbol string) *GetSpotOrdersService {
	s.symbol = &symbol
	return s
}

func (s *GetSpotOrdersService) OrderId(orderId string) *GetSpotOrdersService {
	s.orderId = &orderId
	return s
}

func (s *GetSpotOrdersService) StartTime(startTime int64) *GetSpotOrdersService {
	s.startTime = &startTime
	return s
}

func (s *GetSpotOrdersService) EndTime(endTime int64) *GetSpotOrdersService {
	s.endTime = &endTime
	return s
}

func (s *GetSpotOrdersService) Limit(limit int) *GetSpotOrdersService {
	s.limit = &limit
	return s
}

type SpotOrder struct {
	Symbol                  string `json:"symbol"`                  // 交易对
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrderListId             int64  `json:"orderListId"`             // 订单列表ID
	ClientOrderId           string `json:"clientOrderId"`           // 客户订单ID
	Price                   string `json:"price"`                   // 价格
	OrigQty                 string `json:"origQty"`                 // 原始挂单数量
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累积成交额
	Status                  string `json:"status"`                  // 订单状态
	TimeInForce             string `json:"timeInForce"`             // 时间类型
	Type                    string `json:"type"`                    // 订单类型
	Side                    string `json:"side"`                    // BUY or SELL
	StopPrice               string `json:"stopPrice"`               // 止损价格
	IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
	Time                    int64  `json:"time"`                    // 订单创建时间
	UpdateTime              int64  `json:"updateTime"`              // 订单更新时间
	IsWorking               bool   `json:"isWorking"`               // 是否正在处理
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`       // 原始报价订单数量
	WorkingTime             int64  `json:"workingTime"`             // 处理时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交预防模式
}

func (s *GetSpotOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SpotOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/allOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*SpotOrder
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type GetSpotOpenOrdersService struct {
	c *Client

	symbol *string
}

func (c *Client) NewGetSpotOpenOrdersService() *GetSpotOpenOrdersService {
	return &GetSpotOpenOrdersService{c: c}
}

func (s *GetSpotOpenOrdersService) Symbol(symbol string) *GetSpotOpenOrdersService {
	s.symbol = &symbol
	return s
}

type SpotOpenOrder struct {
	Symbol                  string `json:"symbol"`                  // 交易对
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrderListId             int64  `json:"orderListId"`             // 除非此单是订单列表的一部分, 否则此值为 -1
	ClientOrderId           string `json:"clientOrderId"`           // 客户端订单ID
	Price                   string `json:"price"`                   // 订单价格
	OrigQty                 string `json:"origQty"`                 // 原始数量
	ExecutedQty             string `json:"executedQty"`             // 已执行数量
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`       // 原始报价订单数量
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累计报价数量
	Status                  string `json:"status"`                  // 订单状态
	TimeInForce             string `json:"timeInForce"`             // 有效方式(GTC/IOC/FOK)
	Type                    string `json:"type"`                    // 订单类型
	Side                    string `json:"side"`                    // 订单方向(BUY/SELL)
	StopPrice               string `json:"stopPrice"`               // 止损价
	IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
	Time                    int64  `json:"time"`                    // 订单时间
	UpdateTime              int64  `json:"updateTime"`              // 最后更新时间
	IsWorking               bool   `json:"isWorking"`               // 订单是否生效
	WorkingTime             int64  `json:"workingTime"`             // 工作时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交预防模式
}

func (s *GetSpotOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SpotOpenOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*SpotOpenOrder
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type SwapOrder struct {
	AvgPrice             Float64 `json:"avgPrice"`                // 平均价格
	ClientOrderId        string  `json:"clientOrderId"`           // 客户订单ID
	CumQuote             Float64 `json:"cumQuote"`                // 累积成交额
	ExecutedQty          Float64 `json:"executedQty"`             // 成交数量
	OrderId              Int64   `json:"orderId"`                 // 订单ID
	OrigQty              Float64 `json:"origQty"`                 // 原始挂单数量
	OrigType             string  `json:"origType"`                // 原始订单类型
	Price                Float64 `json:"price"`                   // 价格
	ReduceOnly           bool    `json:"reduceOnly"`              // 是否只减仓
	Side                 string  `json:"side"`                    // 买卖方向 BUY/SELL
	PositionSide         string  `json:"positionSide"`            // 持仓方向
	Status               string  `json:"status"`                  // 订单状态
	ClosePosition        bool    `json:"closePosition"`           // 是否全平仓
	Symbol               string  `json:"symbol"`                  // 交易对
	Time                 Int64   `json:"time"`                    // 订单创建时间
	TimeInForce          string  `json:"timeInForce"`             // 时间类型
	Type                 string  `json:"type"`                    // 订单类型
	UpdateTime           Int64   `json:"updateTime"`              // 订单更新时间
	SelfTradePreventMode string  `json:"selfTradePreventionMode"` // 自成交预防模式
	GoodTillDate         Int64   `json:"goodTillDate"`            // 有效期
	PriceMatch           string  `json:"priceMatch"`              // 价格匹配规则
}

type SwapOrdersGetter interface {
	Symbol(symbol string) SwapOrdersGetter
	OrderId(orderId string) SwapOrdersGetter
	StartTime(startTime int64) SwapOrdersGetter
	EndTime(endTime int64) SwapOrdersGetter
	Limit(limit int) SwapOrdersGetter
	Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error)
}

type SwapOpenOrder struct {
	AvgPrice             string  `json:"avgPrice"`                // 平均成交价
	ClientOrderId        string  `json:"clientOrderId"`           // 用户自定义订单ID
	CumQuote             Float64 `json:"cumQuote"`                // 成交金额
	ExecutedQty          Float64 `json:"executedQty"`             // 已成交数量
	OrderId              Int64   `json:"orderId"`                 // 订单ID
	OrigQty              Float64 `json:"origQty"`                 // 原始委托数量
	OrigType             string  `json:"origType"`                // 原始订单类型
	Price                Float64 `json:"price"`                   // 委托价格
	ReduceOnly           bool    `json:"reduceOnly"`              // 是否仅减仓
	Side                 string  `json:"side"`                    // 买卖方向
	PositionSide         string  `json:"positionSide"`            // 持仓方向
	Status               string  `json:"status"`                  // 订单状态
	Symbol               string  `json:"symbol"`                  // 交易对
	Time                 Int64   `json:"time"`                    // 订单时间
	TimeInForce          string  `json:"timeInForce"`             // 有效方式
	Type                 string  `json:"type"`                    // 订单类型
	UpdateTime           Int64   `json:"updateTime"`              // 更新时间
	SelfTradePreventMode string  `json:"selfTradePreventionMode"` // 自成交预防模式
	GoodTillDate         Int64   `json:"goodTillDate"`            // 订单到期时间
	PriceMatch           string  `json:"priceMatch"`              // 价格匹配模式

	// StopPrice     string `json:"stopPrice"`     // 触发价，对`TRAILING_STOP_MARKET`无效
	// ClosePosition bool   `json:"closePosition"` // 是否条件全平仓
	// ActivatePrice string `json:"activatePrice"` // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	// PriceRate     string `json:"priceRate"`     // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
}

type SwapOpenOrdersGetter interface {
	Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error)
}

type GetPApiSwapOpenOrdersService struct {
	c *Client

	symbol *string
}

func (c *Client) NewGetPApiSwapOpenOrdersService() *GetPApiSwapOpenOrdersService {
	return &GetPApiSwapOpenOrdersService{c: c}
}

func (s *GetPApiSwapOpenOrdersService) Symbol(symbol string) SwapOpenOrdersGetter {
	s.symbol = &symbol
	return s
}

func (s *GetPApiSwapOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/openOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*SwapOpenOrder
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type GetFApiSwapOrdersService struct {
	c *Client

	symbol    *string
	orderId   *string
	startTime *int64
	endTime   *int64
	limit     *int
}

func (c *Client) NewGetFApiSwapOrdersService() *GetFApiSwapOrdersService {
	return &GetFApiSwapOrdersService{c: c}
}

func (s *GetFApiSwapOrdersService) Symbol(symbol string) SwapOrdersGetter {
	s.symbol = &symbol
	return s
}

func (s *GetFApiSwapOrdersService) OrderId(orderId string) SwapOrdersGetter {
	s.orderId = &orderId
	return s
}

func (s *GetFApiSwapOrdersService) StartTime(startTime int64) SwapOrdersGetter {
	s.startTime = &startTime
	return s
}

func (s *GetFApiSwapOrdersService) EndTime(endTime int64) SwapOrdersGetter {
	s.endTime = &endTime
	return s
}

func (s *GetFApiSwapOrdersService) Limit(limit int) SwapOrdersGetter {
	s.limit = &limit
	return s
}

func (s *GetFApiSwapOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*SwapOrder
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type GetFApiSwapOpenOrdersService struct {
	c *Client
}

func (c *Client) NewGetFApiSwapOpenOrdersService() *GetFApiSwapOpenOrdersService {
	return &GetFApiSwapOpenOrdersService{c: c}
}

func (s *GetFApiSwapOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*SwapOpenOrder
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
