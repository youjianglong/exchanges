package binance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// GetSpotTickerPriceService 获取现货交易对价格
type GetSpotTickerPriceService struct {
	c       *Client
	symbol  *string
	symbols []string
}

func (c *Client) NewGetSpotTickerPriceService() *GetSpotTickerPriceService {
	return &GetSpotTickerPriceService{c: c}
}

func (s *GetSpotTickerPriceService) Symbol(symbol string) *GetSpotTickerPriceService {
	s.symbol = &symbol
	return s
}

func (s *GetSpotTickerPriceService) Symbols(symbols []string) *GetSpotTickerPriceService {
	s.symbols = symbols
	return s
}

func (s *GetSpotTickerPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*TickerPrice, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/price",
		secType:  secTypeNone,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if len(s.symbols) > 0 {
		data, _ := json.Marshal(s.symbols)
		r.setParam("symbols", string(data))
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

type GetSpotTicker24HService struct {
	c       *Client
	symbol  *string
	symbols []string
}

func (c *Client) NewGetSpotTicker24HService() *GetSpotTicker24HService {
	return &GetSpotTicker24HService{c: c}
}

func (s *GetSpotTicker24HService) Symbol(symbol string) *GetSpotTicker24HService {
	s.symbol = &symbol
	return s
}

func (s *GetSpotTicker24HService) Symbols(symbols []string) *GetSpotTicker24HService {
	s.symbols = symbols
	return s
}

type SpotTicker24H struct {
	Symbol      string  `json:"symbol"`      // 交易对
	OpenPrice   Float64 `json:"openPrice"`   // 间隔开盘价
	HighPrice   Float64 `json:"highPrice"`   // 间隔最高价
	LowPrice    Float64 `json:"lowPrice"`    // 间隔最低价
	LastPrice   Float64 `json:"lastPrice"`   // 间隔收盘价
	Volume      Float64 `json:"volume"`      // 总交易量
	QuoteVolume Float64 `json:"quoteVolume"` // 总交易额
	OpenTime    Int64   `json:"openTime"`    // ticker间隔的开始时间
	CloseTime   Int64   `json:"closeTime"`   // ticker间隔的结束时间
	FirstId     Int64   `json:"firstId"`     // 统计时间内的第一笔trade id
	LastId      Int64   `json:"lastId"`      // 统计时间内的最后一笔trade id
	Count       Int64   `json:"count"`       // 统计时间内交易笔数
}

func (s *GetSpotTicker24HService) Do(ctx context.Context, opts ...RequestOption) ([]*SpotTicker24H, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/24hr",
		secType:  secTypeNone,
		baseURL:  &s.c.ApiBaseURL,
	}
	r.setParam("type", "MINI")
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if len(s.symbols) > 0 {
		data, _ := json.Marshal(s.symbols)
		r.setParam("symbols", string(data))
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.symbol != nil {
		var res SpotTicker24H
		if err := StrictDecode(data, &res); err != nil {
			return nil, err
		}
		return []*SpotTicker24H{&res}, nil
	}
	var res []*SpotTicker24H
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetSpotSymbolInfosService 获取现货交易对信息
type GetSpotSymbolInfosService struct {
	c *Client

	status *string
}

type SpotSymbolInfos struct {
	Symbols []struct {
		Symbol     string            `json:"symbol"`     // 交易对
		Status     string            `json:"status"`     // 状态
		QuoteAsset string            `json:"quoteAsset"` // 报价币种
		Filters    []json.RawMessage `json:"filters"`    // 过滤器
	} `json:"symbols"`
}

func (c *Client) NewGetSpotSymbolInfosService() *GetSpotSymbolInfosService {
	return &GetSpotSymbolInfosService{c: c}
}

func (s *GetSpotSymbolInfosService) Status(status string) *GetSpotSymbolInfosService {
	s.status = &status
	return s
}

func (s *GetSpotSymbolInfosService) Do(ctx context.Context, opts ...RequestOption) ([]*SymbolInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/exchangeInfo",
		secType:  secTypeNone,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.status != nil {
		r.setParam("symbolStatus", *s.status)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res SpotSymbolInfos
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	var symbols []*SymbolInfo
	for _, s := range res.Symbols {
		if s.Status != "TRADING" {
			continue
		}
		info := &SymbolInfo{
			Symbol:     s.Symbol,
			QuoteAsset: s.QuoteAsset,
		}
		for _, f := range s.Filters {
			if !bytes.Contains(f, lotSizeKey) {
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
