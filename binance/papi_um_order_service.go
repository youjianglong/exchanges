package binance

import (
	"context"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

type PApiUmOrderService struct {
	c *Client

	symbol                  string
	side                    string
	orderType               string
	positionSide            *string
	timeInForce             *string
	quantity                *string
	reduceOnly              *bool
	price                   *string
	newClientOrderId        *string
	newOrderRespType        *string
	priceMatch              *string
	selfTradePreventionMode *string
	goodTillDate            *int64
}

func (c *Client) NewPApiUmOrderService(symbol string, side string, orderType string) *PApiUmOrderService {
	return &PApiUmOrderService{c: c, symbol: symbol, side: side, orderType: orderType}
}

func (s *PApiUmOrderService) PositionSide(positionSide string) *PApiUmOrderService {
	if positionSide != "" {
		s.positionSide = &positionSide
	}
	return s
}

func (s *PApiUmOrderService) OrderType(orderType string) *PApiUmOrderService {
	s.orderType = orderType
	return s
}

func (s *PApiUmOrderService) TimeInForce(timeInForce string) *PApiUmOrderService {
	s.timeInForce = &timeInForce
	return s
}

func (s *PApiUmOrderService) Quantity(quantity string) *PApiUmOrderService {
	s.quantity = &quantity
	return s
}

func (s *PApiUmOrderService) ReduceOnly(reduceOnly bool) *PApiUmOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *PApiUmOrderService) Price(price string) *PApiUmOrderService {
	s.price = &price
	return s
}

func (s *PApiUmOrderService) NewClientOrderId(newClientOrderId string) *PApiUmOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *PApiUmOrderService) NewOrderRespType(newOrderRespType string) *PApiUmOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *PApiUmOrderService) PriceMatch(priceMatch string) *PApiUmOrderService {
	s.priceMatch = &priceMatch
	return s
}

func (s *PApiUmOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *PApiUmOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *PApiUmOrderService) GoodTillDate(goodTillDate int64) *PApiUmOrderService {
	s.goodTillDate = &goodTillDate
	return s
}

func (s *PApiUmOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	if s.positionSide != nil {
		r.setParam("positionSide", *s.positionSide)
	}
	r.setParam("type", s.orderType)
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setParam("quantity", *s.quantity)
	}
	if s.reduceOnly != nil {
		r.setParam("reduceOnly", *s.reduceOnly)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.newClientOrderId != nil {
		r.setParam("newClientOrderId", *s.newClientOrderId)
	}
	if s.newOrderRespType != nil {
		r.setParam("newOrderRespType", *s.newOrderRespType)
	}
	if s.priceMatch != nil {
		r.setParam("priceMatch", *s.priceMatch)
	}
	if s.selfTradePreventionMode != nil {
		r.setParam("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.goodTillDate != nil {
		r.setParam("goodTillDate", *s.goodTillDate)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var order *SwapOrder
	err = StrictDecode(resp, &order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

type PApiUmCancelOrderService struct {
	c *Client

	symbol            string
	orderId           *string
	origClientOrderId *string
}

func (c *Client) NewPApiUmCancelOrderService() *PApiUmCancelOrderService {
	return &PApiUmCancelOrderService{c: c}
}

func (s *PApiUmCancelOrderService) Symbol(symbol string) *PApiUmCancelOrderService {
	s.symbol = symbol
	return s
}

func (s *PApiUmCancelOrderService) OrderId(orderId string) *PApiUmCancelOrderService {
	s.orderId = &orderId
	return s
}

func (s *PApiUmCancelOrderService) OrigClientOrderId(origClientOrderId string) *PApiUmCancelOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *PApiUmCancelOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.origClientOrderId != nil {
		r.setParam("origClientOrderId", *s.origClientOrderId)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var order *SwapOrder
	err = StrictDecode(resp, &order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

type PApiUmCancelAllOpenOrdersService struct {
	c *Client

	symbol string
}

func (c *Client) NewPApiUmCancelAllOpenOrdersService(symbol string) *PApiUmCancelAllOpenOrdersService {
	return &PApiUmCancelAllOpenOrdersService{c: c, symbol: symbol}
}

func (s *PApiUmCancelAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/allOpenOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

type PApiGetOpenOrdersService struct {
	c *Client

	symbol string
}

func (c *Client) NewPApiGetOpenOrdersService() *PApiGetOpenOrdersService {
	return &PApiGetOpenOrdersService{c: c}
}

func (s *PApiGetOpenOrdersService) Symbol(symbol string) *PApiGetOpenOrdersService {
	s.symbol = symbol
	return s
}

func (s *PApiGetOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/openOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*SwapOpenOrder
	err = StrictDecode(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type PApiGetAllOrdersService struct {
	c *Client

	symbol    string
	orderId   *string
	startTime *int64
	endTime   *int64
	limit     *int
}

func (c *Client) NewPApiGetAllOrdersService(symbol string) *PApiGetAllOrdersService {
	return &PApiGetAllOrdersService{c: c, symbol: symbol}
}

func (s *PApiGetAllOrdersService) OrderId(orderId string) *PApiGetAllOrdersService {
	s.orderId = &orderId
	return s
}

func (s *PApiGetAllOrdersService) StartTime(startTime int64) *PApiGetAllOrdersService {
	s.startTime = &startTime
	return s
}

func (s *PApiGetAllOrdersService) EndTime(endTime int64) *PApiGetAllOrdersService {
	s.endTime = &endTime
	return s
}

func (s *PApiGetAllOrdersService) Limit(limit int) *PApiGetAllOrdersService {
	s.limit = &limit
	return s
}

func (s *PApiGetAllOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/allOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
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
	err = StrictDecode(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type PApiUmGetUserTradesService struct {
	c *Client

	symbol    string
	startTime *int64
	endTime   *int64
	fromId    *int64
	limit     *int
}

type TradeRecord struct {
	Symbol          string  `json:"symbol"`
	ID              Int64   `json:"id"`
	OrderID         Int64   `json:"orderId"`
	Side            string  `json:"side"`
	Price           Float64 `json:"price"`
	Qty             Float64 `json:"qty"`
	RealizedPnl     Float64 `json:"realizedPnl"`
	QuoteQty        Float64 `json:"quoteQty"`
	Commission      Float64 `json:"commission"`
	CommissionAsset string  `json:"commissionAsset"`
	Time            Int64   `json:"time"`
	Buyer           bool    `json:"buyer"`
	Maker           bool    `json:"maker"`
	PositionSide    string  `json:"positionSide"`
}

func (c *Client) NewPApiUmGetUserTradesService(symbol string) *PApiUmGetUserTradesService {
	return &PApiUmGetUserTradesService{c: c, symbol: symbol}
}

func (s *PApiUmGetUserTradesService) StartTime(startTime int64) *PApiUmGetUserTradesService {
	s.startTime = &startTime
	return s
}

func (s *PApiUmGetUserTradesService) EndTime(endTime int64) *PApiUmGetUserTradesService {
	s.endTime = &endTime
	return s
}

func (s *PApiUmGetUserTradesService) FromId(fromId int64) *PApiUmGetUserTradesService {
	if fromId > 0 {
		s.fromId = &fromId
	}
	return s
}

func (s *PApiUmGetUserTradesService) Limit(limit int) *PApiUmGetUserTradesService {
	if limit > 0 {
		s.limit = &limit
	}
	return s
}

func (s *PApiUmGetUserTradesService) Do(ctx context.Context, opts ...RequestOption) ([]*TradeRecord, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/userTrades",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var trades []*TradeRecord
	err = StrictDecode(resp, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}

type PApiUmGetOrderService struct {
	c *Client

	symbol            string
	orderId           string
	origClientOrderId string
}

func (c *Client) NewPApiUmGetOrderService(symbol string) *PApiUmGetOrderService {
	return &PApiUmGetOrderService{c: c, symbol: symbol}
}

func (s *PApiUmGetOrderService) OrderId(orderId string) *PApiUmGetOrderService {
	s.orderId = orderId
	return s
}

func (s *PApiUmGetOrderService) OrigClientOrderId(origClientOrderId string) *PApiUmGetOrderService {
	s.origClientOrderId = origClientOrderId
	return s
}

func (s *PApiUmGetOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	if s.orderId != "" {
		r.setParam("orderId", s.orderId)
	}
	if s.origClientOrderId != "" {
		r.setParam("origClientOrderId", s.origClientOrderId)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var order *SwapOrder
	err = StrictDecode(resp, &order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
