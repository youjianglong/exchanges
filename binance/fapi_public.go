package binance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/youjianglong/exchanges/common"
)

// GetPremiumIndexService 获取资金费率和指数价格
type GetPremiumIndexService struct {
	c      *Client
	symbol *string
}

func (c *Client) NewGetPremiumIndexService() *GetPremiumIndexService {
	return &GetPremiumIndexService{c: c}
}

func (s *GetPremiumIndexService) Symbol(symbol string) *GetPremiumIndexService {
	s.symbol = &symbol
	return s
}

type PremiumIndex struct {
	Symbol          string  `json:"symbol"`               // 交易对
	MarkPrice       Float64 `json:"markPrice"`            // 标记价格
	IndexPrice      Float64 `json:"indexPrice"`           // 指数价格
	EstSettlePrice  string  `json:"estimatedSettlePrice"` // 预估结算价,仅在交割开始前最后一小时有意义
	LastFundingRate Float64 `json:"lastFundingRate"`      // 最近更新的资金费率
	InterestRate    string  `json:"interestRate"`         // 标的资产基础利率
	NextFundingTime Int64   `json:"nextFundingTime"`      // 下次资金费时间
	Time            Int64   `json:"time"`                 // 更新时间
}

func (s *GetPremiumIndexService) Do(ctx context.Context, opts ...RequestOption) ([]*PremiumIndex, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/premiumIndex",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.symbol == nil {
		var res []*PremiumIndex
		if err := StrictDecode(data, &res); err != nil {
			return nil, err
		}
		return res, nil
	}
	var res PremiumIndex
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return []*PremiumIndex{&res}, nil
}

type TickerPrice struct {
	Symbol string  `json:"symbol"` // 交易对
	Price  Float64 `json:"price"`  // 价格
	Time   Int64   `json:"time"`   // 更新时间
}

// GetFApiSymbolInfosService 获取永续合约交易对信息
type GetFApiSymbolInfosService struct {
	c *Client
}

type FApiSymbolInfos struct {
	Symbols []struct {
		Symbol              string            `json:"symbol"`              // 交易对
		QuoteAsset          string            `json:"quoteAsset"`          // 报价币种
		ContractType        string            `json:"contractType"`        // 合约类型
		Status              string            `json:"status"`              // 状态
		PricePrecision      int               `json:"pricePrecision"`      // 价格精度
		QuantityPrecision   int               `json:"quantityPrecision"`   // 数量精度
		BaseAssetPrecision  int               `json:"baseAssetPrecision"`  // 基础货币精度
		QuoteAssetPrecision int               `json:"quoteAssetPrecision"` // 报价货币精度
		Filters             []json.RawMessage `json:"filters"`             // 过滤器
	} `json:"symbols"`
}

func (c *Client) NewGetFApiSymbolInfosService() *GetFApiSymbolInfosService {
	return &GetFApiSymbolInfosService{c: c}
}

func (s *GetFApiSymbolInfosService) Do(ctx context.Context, opts ...RequestOption) ([]*SymbolInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/exchangeInfo",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res FApiSymbolInfos
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	var symbols []*SymbolInfo
	for _, s := range res.Symbols {
		if s.ContractType != "PERPETUAL" {
			continue
		}
		if s.Status != "TRADING" {
			continue
		}
		info := &SymbolInfo{
			Symbol:              s.Symbol,
			QuoteAsset:          s.QuoteAsset,
			PricePrecision:      s.PricePrecision,
			QuantityPrecision:   s.QuantityPrecision,
			BaseAssetPrecision:  s.BaseAssetPrecision,
			QuoteAssetPrecision: s.QuoteAssetPrecision,
		}
		for _, f := range s.Filters {
			if !bytes.Contains(f, lotSizeKey) && !bytes.Contains(f, minValueKey) {
				continue
			}
			var lotSize LimitFilter
			if err := StrictDecode(f, &lotSize); err != nil {
				return nil, fmt.Errorf("unmarshal lotSize %s, %s: %w", s.Symbol, string(f), err)
			}
			switch lotSize.FilterType {
			case "LOT_SIZE":
				info.LotSize = LotLimit{
					MinQty:   lotSize.MinQty,
					MaxQty:   lotSize.MaxQty,
					StepSize: lotSize.StepSize,
				}
			case "MARKET_LOT_SIZE":
				info.MarketLotSize = LotLimit{
					MinQty:   lotSize.MinQty,
					MaxQty:   lotSize.MaxQty,
					StepSize: lotSize.StepSize,
				}
			case "MIN_NOTIONAL":
				info.MinValue = lotSize.Notional
			}
		}
		symbols = append(symbols, info)
	}
	return symbols, nil
}

// GetFApiTickerPriceService 获取永续合约交易对价格
type GetFApiTickerPriceService struct {
	c      *Client
	symbol *string
}

func (c *Client) NewGetFApiTickerPriceService() *GetFApiTickerPriceService {
	return &GetFApiTickerPriceService{c: c}
}

func (s *GetFApiTickerPriceService) Symbol(symbol string) *GetFApiTickerPriceService {
	s.symbol = &symbol
	return s
}

func (s *GetFApiTickerPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*TickerPrice, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/ticker/price",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.symbol != nil {
		var res TickerPrice
		if err := StrictDecode(data, &res); err != nil {
			return nil, err
		}
		return []*TickerPrice{&res}, nil
	}
	var res []*TickerPrice
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type GetFApiTicker24HService struct {
	c *Client

	symbol *string
}

func (c *Client) NewGetFApiTicker24HService() *GetFApiTicker24HService {
	return &GetFApiTicker24HService{c: c}
}

func (s *GetFApiTicker24HService) Symbol(symbol string) *GetFApiTicker24HService {
	s.symbol = &symbol
	return s
}

type FApiTicker24H struct {
	Symbol             string  `json:"symbol"`
	PriceChange        Float64 `json:"priceChange"`        // 24小时价格变动
	PriceChangePercent Float64 `json:"priceChangePercent"` // 24小时价格变动百分比
	WeightedAvgPrice   Float64 `json:"weightedAvgPrice"`   // 加权平均价
	LastPrice          Float64 `json:"lastPrice"`          // 最近一次成交价
	LastQty            Float64 `json:"lastQty"`            // 最近一次成交额
	OpenPrice          Float64 `json:"openPrice"`          // 24小时内第一次成交的价格
	HighPrice          Float64 `json:"highPrice"`          // 24小时最高价
	LowPrice           Float64 `json:"lowPrice"`           // 24小时最低价
	Volume             Float64 `json:"volume"`             // 24小时成交量
	QuoteVolume        Float64 `json:"quoteVolume"`        // 24小时成交额
	OpenTime           Int64   `json:"openTime"`           // 24小时内，第一笔交易的发生时间
	CloseTime          Int64   `json:"closeTime"`          // 24小时内，最后一笔交易的发生时间
	FirstId            Int64   `json:"firstId"`            // 首笔成交id
	LastId             Int64   `json:"lastId"`             // 末笔成交id
	Count              Int64   `json:"count"`              // 成交笔数
}

func (s *GetFApiTicker24HService) Do(ctx context.Context, opts ...RequestOption) ([]*FApiTicker24H, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/ticker/24hr",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.symbol != nil {
		var res FApiTicker24H
		if err := StrictDecode(data, &res); err != nil {
			return nil, err
		}
		return []*FApiTicker24H{&res}, nil
	}
	var res []*FApiTicker24H
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type FApiGetDepthService struct {
	c *Client

	symbol string
	limit  *int
}

func (c *Client) NewFApiGetDepthService(symbol string) *FApiGetDepthService {
	return &FApiGetDepthService{c: c, symbol: symbol}
}

func (s *FApiGetDepthService) Limit(limit int) *FApiGetDepthService {
	s.limit = &limit
	return s
}

type Depth struct {
	LastUpdateId Int64        `json:"lastUpdateId"` // 最后更新ID
	EventTime    Int64        `json:"E"`            // 事件时间
	UpdateTime   Int64        `json:"T"`            // 撮合更新时间
	Asks         [][2]Float64 `json:"asks"`         // 卖盘
	Bids         [][2]Float64 `json:"bids"`         // 买盘
}

func (s *FApiGetDepthService) Do(ctx context.Context, opts ...RequestOption) (*Depth, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/depth",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", strconv.Itoa(*s.limit))
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res Depth
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type FApiGetKLinesService struct {
	c *Client

	symbol    string
	interval  string
	startTime *int64
	endTime   *int64
	limit     *int
}

func (c *Client) NewFApiGetKLinesService(symbol, interval string) *FApiGetKLinesService {
	return &FApiGetKLinesService{c: c, symbol: symbol, interval: interval}
}

func (s *FApiGetKLinesService) StartTime(startTime int64) *FApiGetKLinesService {
	s.startTime = &startTime
	return s
}

func (s *FApiGetKLinesService) EndTime(endTime int64) *FApiGetKLinesService {
	s.endTime = &endTime
	return s
}

func (s *FApiGetKLinesService) Limit(limit int) *FApiGetKLinesService {
	s.limit = &limit
	return s
}

func (s *FApiGetKLinesService) Do(ctx context.Context, opts ...RequestOption) ([][]Mixed, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/klines",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res [][]Mixed
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type FApiGetBookTickerService struct {
	c      *Client
	symbol *string
}

type FApiBookTicker struct {
	Symbol   string  `json:"symbol"`   // 交易对
	BidPrice Float64 `json:"bidPrice"` // 最高买价
	BidQty   Float64 `json:"bidQty"`   // 最高买价挂单量
	AskPrice Float64 `json:"askPrice"` // 最低卖价
	AskQty   Float64 `json:"askQty"`   // 最低卖价挂单量
	Time     Int64   `json:"time"`     // 更新时间
}

func (c *Client) NewFApiGetBookTickerService() *FApiGetBookTickerService {
	return &FApiGetBookTickerService{c: c}
}

func (s *FApiGetBookTickerService) Symbol(symbol string) *FApiGetBookTickerService {
	s.symbol = &symbol
	return s
}

func (s *FApiGetBookTickerService) Do(ctx context.Context, opts ...RequestOption) ([]*FApiBookTicker, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/ticker/bookTicker",
		secType:  secTypeNone,
		baseURL:  &s.c.FApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.symbol != nil {
		var res FApiBookTicker
		if err := StrictDecode(data, &res); err != nil {
			return nil, err
		}
		return []*FApiBookTicker{&res}, nil
	}
	var res []*FApiBookTicker
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
