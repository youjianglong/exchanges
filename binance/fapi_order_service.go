package binance

import (
	"context"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api
type FApiOrderService struct {
	c *Client

	symbol                  string
	side                    string
	positionSide            *string
	orderType               string // type
	timeInForce             *string
	quantity                *string
	reduceOnly              *bool
	price                   *string
	newClientOrderId        *string
	stopPrice               *string
	closePosition           *string
	activationPrice         *string
	callbackRate            *string
	workingType             *string
	priceProtect            *bool
	newOrderRespType        *string
	priceMatch              *string
	selfTradePreventionMode *string
	goodTillDate            *int64
}

func (c *Client) NewFApiOrderService(symbol string, side string, orderType string) *FApiOrderService {
	return &FApiOrderService{c: c, symbol: symbol, side: side, orderType: orderType}
}

func (s *FApiOrderService) PositionSide(positionSide string) *FApiOrderService {
	if positionSide != "" {
		s.positionSide = &positionSide
	}
	return s
}

func (s *FApiOrderService) OrderType(orderType string) *FApiOrderService {
	s.orderType = orderType
	return s
}

func (s *FApiOrderService) TimeInForce(timeInForce string) *FApiOrderService {
	s.timeInForce = &timeInForce
	return s
}

func (s *FApiOrderService) Quantity(quantity string) *FApiOrderService {
	s.quantity = &quantity
	return s
}

func (s *FApiOrderService) ReduceOnly(reduceOnly bool) *FApiOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *FApiOrderService) Price(price string) *FApiOrderService {
	s.price = &price
	return s
}

func (s *FApiOrderService) NewClientOrderId(newClientOrderId string) *FApiOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *FApiOrderService) StopPrice(stopPrice string) *FApiOrderService {
	s.stopPrice = &stopPrice
	return s
}

func (s *FApiOrderService) ClosePosition(closePosition string) *FApiOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *FApiOrderService) ActivationPrice(activationPrice string) *FApiOrderService {
	s.activationPrice = &activationPrice
	return s
}

func (s *FApiOrderService) CallbackRate(callbackRate string) *FApiOrderService {
	s.callbackRate = &callbackRate
	return s
}

func (s *FApiOrderService) WorkingType(workingType string) *FApiOrderService {
	s.workingType = &workingType
	return s
}

func (s *FApiOrderService) PriceProtect(priceProtect bool) *FApiOrderService {
	s.priceProtect = &priceProtect
	return s
}

func (s *FApiOrderService) NewOrderRespType(newOrderRespType string) *FApiOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *FApiOrderService) PriceMatch(priceMatch string) *FApiOrderService {
	s.priceMatch = &priceMatch
	return s
}

func (s *FApiOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *FApiOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *FApiOrderService) GoodTillDate(goodTillDate int64) *FApiOrderService {
	s.goodTillDate = &goodTillDate
	return s
}

func (s *FApiOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	r.setFormParam("symbol", s.symbol)
	r.setFormParam("side", s.side)
	r.setFormParam("type", s.orderType)
	if s.positionSide != nil {
		r.setFormParam("positionSide", *s.positionSide)
	}
	if s.timeInForce != nil {
		r.setFormParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setFormParam("quantity", *s.quantity)
	}
	if s.reduceOnly != nil {
		r.setFormParam("reduceOnly", *s.reduceOnly)
	}
	if s.price != nil {
		r.setFormParam("price", *s.price)
	}
	if s.newClientOrderId != nil {
		r.setFormParam("newClientOrderId", *s.newClientOrderId)
	}
	if s.stopPrice != nil {
		r.setFormParam("stopPrice", *s.stopPrice)
	}
	if s.closePosition != nil {
		r.setFormParam("closePosition", *s.closePosition)
	}
	if s.activationPrice != nil {
		r.setFormParam("activationPrice", *s.activationPrice)
	}
	if s.callbackRate != nil {
		r.setFormParam("callbackRate", *s.callbackRate)
	}
	if s.workingType != nil {
		r.setFormParam("workingType", *s.workingType)
	}
	if s.priceProtect != nil {
		r.setFormParam("priceProtect", *s.priceProtect)
	}
	if s.newOrderRespType != nil {
		r.setFormParam("newOrderRespType", *s.newOrderRespType)
	}
	if s.priceMatch != nil {
		r.setFormParam("priceMatch", *s.priceMatch)
	}
	if s.selfTradePreventionMode != nil {
		r.setFormParam("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.goodTillDate != nil {
		r.setFormParam("goodTillDate", *s.goodTillDate)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var order SwapOrder
	err = StrictDecode(resp, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// 取消订单
// https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api
type FApiCancelOrderService struct {
	c *Client

	symbol            string
	orderId           *string
	origClientOrderId *string
}

func (c *Client) NewFApiCancelOrderService(symbol string) *FApiCancelOrderService {
	return &FApiCancelOrderService{c: c, symbol: symbol}
}

func (s *FApiCancelOrderService) OrderId(orderId string) *FApiCancelOrderService {
	s.orderId = &orderId
	return s
}

func (s *FApiCancelOrderService) OrigClientOrderId(origClientOrderId string) *FApiCancelOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *FApiCancelOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
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
	var order SwapOrder
	err = StrictDecode(resp, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// 查询订单
// https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/Query-Order
type FApiGetOrderService struct {
	c *Client

	symbol            string
	orderId           *string
	origClientOrderId *string
}

func (c *Client) NewFApiGetOrderService(symbol string) *FApiGetOrderService {
	return &FApiGetOrderService{c: c, symbol: symbol}
}

func (s *FApiGetOrderService) OrderId(orderId string) *FApiGetOrderService {
	s.orderId = &orderId
	return s
}

func (s *FApiGetOrderService) OrigClientOrderId(origClientOrderId string) *FApiGetOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *FApiGetOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
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
	var order SwapOrder
	err = StrictDecode(resp, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// 查询所有订单
// https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/All-Orders
type FApiGetAllOrdersService struct {
	c *Client

	symbol    string
	orderId   *string
	startTime *int64
	endTime   *int64
	limit     *int
}

func (c *Client) NewFApiGetAllOrdersService(symbol string) *FApiGetAllOrdersService {
	return &FApiGetAllOrdersService{c: c, symbol: symbol}
}

func (s *FApiGetAllOrdersService) OrderId(orderId string) *FApiGetAllOrdersService {
	s.orderId = &orderId
	return s
}

func (s *FApiGetAllOrdersService) StartTime(startTime int64) *FApiGetAllOrdersService {
	s.startTime = &startTime
	return s
}

func (s *FApiGetAllOrdersService) EndTime(endTime int64) *FApiGetAllOrdersService {
	s.endTime = &endTime
	return s
}

func (s *FApiGetAllOrdersService) Limit(limit int) *FApiGetAllOrdersService {
	s.limit = &limit
	return s
}

func (s *FApiGetAllOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
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

// 查询当前所有挂单
// https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/Current-All-Open-Orders
type FApiGetAllOpenOrdersService struct {
	c *Client

	symbol *string
}

func (c *Client) NewFApiGetAllOpenOrdersService() *FApiGetAllOpenOrdersService {
	return &FApiGetAllOpenOrdersService{c: c}
}

func (s *FApiGetAllOpenOrdersService) Symbol(symbol string) *FApiGetAllOpenOrdersService {
	s.symbol = &symbol
	return s
}

func (s *FApiGetAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openOrders",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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

// 调整开仓杠杆
// https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/Change-Initial-Leverage
type FApiChangeLeverageService struct {
	c *Client

	symbol   string
	leverage int64
}

func (c *Client) NewFApiChangeLeverageService(symbol string, leverage int64) *FApiChangeLeverageService {
	return &FApiChangeLeverageService{c: c, symbol: symbol, leverage: leverage}
}

func (s *FApiChangeLeverageService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/leverage",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	r.setFormParam("symbol", s.symbol)
	r.setFormParam("leverage", s.leverage)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

type FApiGetUserTradesService struct {
	c *Client

	symbol    string
	orderId   *string
	startTime *int64
	endTime   *int64
	fromId    *int64
	limit     *int
}

func (c *Client) NewFApiGetUserTradesService(symbol string) *FApiGetUserTradesService {
	return &FApiGetUserTradesService{c: c, symbol: symbol}
}

func (s *FApiGetUserTradesService) OrderId(orderId string) *FApiGetUserTradesService {
	s.orderId = &orderId
	return s
}

func (s *FApiGetUserTradesService) StartTime(startTime int64) *FApiGetUserTradesService {
	s.startTime = &startTime
	return s
}

func (s *FApiGetUserTradesService) EndTime(endTime int64) *FApiGetUserTradesService {
	s.endTime = &endTime
	return s
}

func (s *FApiGetUserTradesService) FromId(fromId int64) *FApiGetUserTradesService {
	s.fromId = &fromId
	return s
}

func (s *FApiGetUserTradesService) Limit(limit int) *FApiGetUserTradesService {
	s.limit = &limit
	return s
}

func (s *FApiGetUserTradesService) Do(ctx context.Context, opts ...RequestOption) ([]*TradeRecord, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/userTrades",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
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
