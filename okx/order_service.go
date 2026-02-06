package okx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// GetOrderHistoryService 获取历史订单记录（近三个月）
// https://www.okx.com/docs-v5/zh/#order-book-trading-trade-get-order-list
type GetOrderHistoryService struct {
	c *Client

	archive  bool    // 是否为归档订单，true：归档订单 false：非归档订单
	instType *string // 产品类型，SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	instId   *string // 产品ID
	state    *string // 订单状态，canceled：撤单成功 filled：完全成交
	after    *string // 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
	before   *string // 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
	begin    *int64  // 开始时间，毫秒时间戳
	end      *int64  // 结束时间，毫秒时间戳
	limit    *int    // 每页条数，最大100
}

func (c *Client) NewGetOrderHistoryService() *GetOrderHistoryService {
	return &GetOrderHistoryService{c: c}
}

func (s *GetOrderHistoryService) InstType(instType string) *GetOrderHistoryService {
	s.instType = &instType
	return s
}

func (s *GetOrderHistoryService) InstId(instId string) *GetOrderHistoryService {
	s.instId = &instId
	return s
}

func (s *GetOrderHistoryService) State(state string) *GetOrderHistoryService {
	s.state = &state
	return s
}

func (s *GetOrderHistoryService) Archive(archive bool) *GetOrderHistoryService {
	s.archive = archive
	return s
}

func (s *GetOrderHistoryService) After(after string) *GetOrderHistoryService {
	s.after = &after
	return s
}

func (s *GetOrderHistoryService) Before(before string) *GetOrderHistoryService {
	s.before = &before
	return s
}

func (s *GetOrderHistoryService) Begin(begin int64) *GetOrderHistoryService {
	s.begin = &begin
	return s
}

func (s *GetOrderHistoryService) End(end int64) *GetOrderHistoryService {
	s.end = &end
	return s
}

func (s *GetOrderHistoryService) Limit(limit int) *GetOrderHistoryService {
	s.limit = &limit
	return s
}

type Order struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	OrdId     string `json:"ordId"`     // 订单ID
	OrdType   string `json:"ordType"`   // 订单类型
	Side      string `json:"side"`      // 交易方向
	PosSide   string `json:"posSide"`   // 持仓方向
	AccFillSz string `json:"accFillSz"` // 累计成交数量
	AvgPx     string `json:"avgPx"`     // 成交均价
	State     string `json:"state"`     // 订单状态
	Lever     string `json:"lever"`     // 杠杆倍数
	FeeCcy    string `json:"feeCcy"`    // 手续费币种
	Fee       string `json:"fee"`       // 手续费
	CTime     string `json:"cTime"`     // 创建时间
	FillTime  string `json:"fillTime"`  // 最新成交时间
}

func (s *GetOrderHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]*Order, error) {
	var endpoint string
	if s.archive {
		endpoint = "/api/v5/trade/orders-history-archive"
	} else {
		endpoint = "/api/v5/trade/orders-history"
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	// 设置产品类型
	if s.instType != nil {
		r.setParam("instType", *s.instType)
	}
	if s.instId != nil {
		if *s.instType == "SWAP" {
			newValue := *s.instId + "-SWAP"
			s.instId = &newValue
		}
		r.setParam("instId", *s.instId)
	}
	if s.state != nil {
		r.setParam("state", *s.state)
	}
	if s.begin != nil {
		r.setParam("begin", *s.begin)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*Order
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type GetOrdersPendingService struct {
	c *Client

	instType *string // 产品类型，SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	instId   *string // 产品ID
	ordType  *string // 订单类型，limit：限价单 market：市价单 post_only：只做maker市价单 fok：全额成交或立即取消市价单 ioc：立即成交并取消剩余
	state    *string // 订单状态，live：未成交 partially_filled：部分成交
	after    *string // 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
	before   *string // 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
	limit    *int    // 每页条数，最大100
}

func (c *Client) NewGetOrdersPendingService() *GetOrdersPendingService {
	return &GetOrdersPendingService{c: c}
}

func (s *GetOrdersPendingService) InstType(instType string) *GetOrdersPendingService {
	s.instType = &instType
	return s
}

func (s *GetOrdersPendingService) InstId(instId string) *GetOrdersPendingService {
	s.instId = &instId
	return s
}

func (s *GetOrdersPendingService) OrdType(ordType string) *GetOrdersPendingService {
	s.ordType = &ordType
	return s
}

func (s *GetOrdersPendingService) State(state string) *GetOrdersPendingService {
	s.state = &state
	return s
}

func (s *GetOrdersPendingService) After(after string) *GetOrdersPendingService {
	s.after = &after
	return s
}

func (s *GetOrdersPendingService) Before(before string) *GetOrdersPendingService {
	s.before = &before
	return s
}

func (s *GetOrdersPendingService) Limit(limit int) *GetOrdersPendingService {
	s.limit = &limit
	return s
}

type PendingOrder struct {
	InstType string `json:"instType"` // 产品类型
	InstId   string `json:"instId"`   // 产品ID
	OrdId    string `json:"ordId"`    // 订单ID
	OrdType  string `json:"ordType"`  // 订单类型
	Side     string `json:"side"`     // 交易方向
	PosSide  string `json:"posSide"`  // 持仓方向
	Px       string `json:"px"`       // 委托价格
	Sz       string `json:"sz"`       // 委托数量
	State    string `json:"state"`    // 订单状态
	CTime    string `json:"cTime"`    // 创建时间
	UTime    string `json:"uTime"`    // 更新时间
}

func (s *GetOrdersPendingService) Do(ctx context.Context, opts ...RequestOption) ([]*PendingOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/orders-pending",
		secType:  secTypeSigned,
	}
	if s.instType != nil {
		r.setParam("instType", *s.instType)
	}
	if s.instId != nil {
		if *s.instType == "SWAP" {
			newValue := *s.instId + "-SWAP"
			s.instId = &newValue
		}
		r.setParam("instId", *s.instId)
	}
	if s.ordType != nil {
		r.setParam("ordType", *s.ordType)
	}
	if s.state != nil {
		r.setParam("state", *s.state)
	}
	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var orders []*PendingOrder
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type TradeOrderService struct {
	c          *Client
	instId     string  // 产品ID，如 BTC-USDT-SWAP
	tdMode     string  // 交易模式
	side       string  // 订单方向 buy: 买，sell: 卖
	ordType    string  // 订单类型，limit: 限价单，market: 市价单
	sz         string  // 委托数量
	px         *string // 委托价格
	posSide    *string // 持仓方向，long: 多头，short: 空头
	clOrdId    *string // 客户自定义订单ID
	reduceOnly *bool   // 是否只减仓，true: 只减仓，false: 不减仓
}

func (c *Client) NewTradeOrderService(instId string, tdMode string, side string, ordType string, sz string) *TradeOrderService {
	return &TradeOrderService{c: c, instId: instId, tdMode: tdMode, side: side, ordType: ordType, sz: sz}
}

func (s *TradeOrderService) PosSide(posSide string) *TradeOrderService {
	s.posSide = &posSide
	return s
}

func (s *TradeOrderService) Px(px string) *TradeOrderService {
	s.px = &px
	return s
}

func (s *TradeOrderService) ClOrdId(clOrdId string) *TradeOrderService {
	s.clOrdId = &clOrdId
	return s
}

func (s *TradeOrderService) ReduceOnly(reduceOnly bool) *TradeOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

type OrderResult struct {
	OrdId   string `json:"ordId"`   // 订单ID
	ClOrdId string `json:"clOrdId"` // 客户自定义订单ID
	Tag     string `json:"tag"`     // 订单标签
	Ts      Int64  `json:"ts"`      // 订单创建时间
	SCode   string `json:"sCode"`   // 状态码
	SMsg    string `json:"sMsg"`    // 状态信息
}

func (s *TradeOrderService) Do(ctx context.Context, opts ...RequestOption) (*OrderResult, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/order",
		secType:  secTypeSigned,
	}
	r.setData("instId", s.instId)
	r.setData("tdMode", s.tdMode)
	r.setData("side", s.side)
	r.setData("ordType", s.ordType)
	r.setData("sz", s.sz)
	if s.px != nil {
		r.setData("px", *s.px)
	}
	if s.posSide != nil {
		r.setData("posSide", *s.posSide)
	}
	if s.clOrdId != nil {
		r.setData("clOrdId", *s.clOrdId)
	}
	if s.reduceOnly != nil {
		r.setData("reduceOnly", *s.reduceOnly)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var results []OrderResult
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no order found")
	}
	return &results[0], nil
}

type GetTradeFillService struct {
	c *Client

	instType   *string // 产品类型
	instFamily *string // 产品家族
	instId     *string // 产品ID
	ordId      *string // 订单ID
	subType    *string // 子类型
	after      *string // 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
	before     *string // 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
	begin      *string // 筛选的开始时间戳 ts，Unix 时间戳为毫秒数格式，如 1597026383085
	end        *string // 筛选的结束时间戳 ts，Unix 时间戳为毫秒数格式，如 1597027383085
	limit      *string // 返回结果的数量，最大为100，默认100条
}

func (c *Client) NewGetTradeFillService() *GetTradeFillService {
	return &GetTradeFillService{c: c}
}

func (s *GetTradeFillService) InstType(instType string) *GetTradeFillService {
	s.instType = &instType
	return s
}

func (s *GetTradeFillService) InstFamily(instFamily string) *GetTradeFillService {
	s.instFamily = &instFamily
	return s
}

func (s *GetTradeFillService) InstId(instId string) *GetTradeFillService {
	s.instId = &instId
	return s
}

func (s *GetTradeFillService) OrdId(ordId string) *GetTradeFillService {
	s.ordId = &ordId
	return s
}

func (s *GetTradeFillService) SubType(subType string) *GetTradeFillService {
	s.subType = &subType
	return s
}

func (s *GetTradeFillService) After(after string) *GetTradeFillService {
	s.after = &after
	return s
}

func (s *GetTradeFillService) Before(before string) *GetTradeFillService {
	s.before = &before
	return s
}

func (s *GetTradeFillService) Begin(begin string) *GetTradeFillService {
	s.begin = &begin
	return s
}

func (s *GetTradeFillService) End(end string) *GetTradeFillService {
	s.end = &end
	return s
}

func (s *GetTradeFillService) Limit(limit string) *GetTradeFillService {
	s.limit = &limit
	return s
}

type TradeFill struct {
	Side          string  `json:"side"`
	FillSz        Float64 `json:"fillSz"`
	FillPx        Float64 `json:"fillPx"`
	FillPxVol     string  `json:"fillPxVol"`
	FillFwdPx     string  `json:"fillFwdPx"`
	Fee           Float64 `json:"fee"`
	FillPnl       Float64 `json:"fillPnl"`
	OrdId         string  `json:"ordId"`
	FeeRate       string  `json:"feeRate"`
	InstType      string  `json:"instType"`
	FillPxUsd     string  `json:"fillPxUsd"`
	InstId        string  `json:"instId"`
	ClOrdId       string  `json:"clOrdId"`
	PosSide       string  `json:"posSide"`
	BillId        string  `json:"billId"`
	SubType       string  `json:"subType"`
	FillMarkVol   string  `json:"fillMarkVol"`
	Tag           string  `json:"tag"`
	FillTime      Int64   `json:"fillTime"`
	ExecType      string  `json:"execType"`
	FillIdxPx     string  `json:"fillIdxPx"`
	TradeId       string  `json:"tradeId"`
	FillMarkPx    string  `json:"fillMarkPx"`
	FeeCcy        string  `json:"feeCcy"`
	Ts            Int64   `json:"ts"`
	TradeQuoteCcy string  `json:"tradeQuoteCcy"`
}

func (s *GetTradeFillService) Do(ctx context.Context, opts ...RequestOption) ([]TradeFill, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/fills",
		secType:  secTypeSigned,
	}
	if s.instType != nil {
		r.setParam("instType", *s.instType)
	}
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	if s.ordId != nil {
		r.setParam("ordId", *s.ordId)
	}
	if s.subType != nil {
		r.setParam("subType", *s.subType)
	}
	if s.begin != nil {
		r.setParam("begin", *s.begin)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var fills []TradeFill
	err = json.Unmarshal(resp, &fills)
	if err != nil {
		return nil, err
	}
	return fills, nil
}
