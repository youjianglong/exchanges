package binance

import (
	"context"
	"encoding/json"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

type FApiGetAccountBalanceService struct {
	c *Client
}

func (c *Client) NewFApiGetAccountBalanceService() *FApiGetAccountBalanceService {
	return &FApiGetAccountBalanceService{c: c}
}

type FApiAccountBalance struct {
	AccountAlias       string  `json:"accountAlias"`       // 账户唯一识别码
	Asset              string  `json:"asset"`              // 资产
	Balance            Float64 `json:"balance"`            // 总余额
	CrossWalletBalance Float64 `json:"crossWalletBalance"` // 全仓余额
	CrossUnPnl         Float64 `json:"crossUnPnl"`         // 全仓持仓未实现盈亏
	AvailableBalance   Float64 `json:"availableBalance"`   // 下单可用余额
	MaxWithdrawAmount  string  `json:"maxWithdrawAmount"`  // 最大可转出余额
	MarginAvailable    bool    `json:"marginAvailable"`    // 是否可用作联合保证金
	UpdateTime         int64   `json:"updateTime"`         // 更新时间
}

func (s *FApiGetAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]*FApiAccountBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v3/balance",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*FApiAccountBalance
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Position struct {
	Symbol           string  `json:"symbol"`           // 交易对
	PositionAmt      Float64 `json:"positionAmt"`      // 头寸数量，符号代表多空方向, 正数为多，负数为空
	PositionSide     string  `json:"positionSide"`     // 持仓方向
	EntryPrice       Float64 `json:"entryPrice"`       // 开仓均价
	Leverage         Int64   `json:"leverage"`         // 杠杆倍数
	UnRealizedProfit Float64 `json:"unRealizedProfit"` // 持仓未实现盈亏
	InitialMargin    Float64 `json:"initialMargin"`    // 初始保证金
	MaintMargin      Float64 `json:"maintMargin"`      // 维持保证金
	MarkPrice        Float64 `json:"markPrice"`        // 当前标记价格
	Notional         Float64 `json:"notional"`         // 名义价值
	LiquidationPrice Float64 `json:"liquidationPrice"` // 预估强平价格
	UpdateTime       Int64   `json:"updateTime"`       // 更新时间
}

type PositionsGetter interface {
	Do(ctx context.Context, opts ...RequestOption) ([]*Position, error)
}

type FApiGetPositionsService struct {
	c *Client
}

func (c *Client) NewFApiGetPositionsService() *FApiGetPositionsService {
	return &FApiGetPositionsService{c: c}
}

func (s *FApiGetPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*Position, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v3/positionRisk",
		secType:  secTypeSigned,
		baseURL:  &s.c.FApiBaseURL,
	}
	resp, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var risks []*Position
	err = json.Unmarshal(resp, &risks)
	if err != nil {
		return nil, err
	}
	return risks, nil
}

type FApiGetIncomeService struct {
	c *Client

	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	page       *int
	limit      *int
}

func (c *Client) NewFApiGetIncomeService() *FApiGetIncomeService {
	return &FApiGetIncomeService{c: c}
}

func (s *FApiGetIncomeService) Symbol(symbol string) *FApiGetIncomeService {
	s.symbol = &symbol
	return s
}

func (s *FApiGetIncomeService) IncomeType(incomeType string) *FApiGetIncomeService {
	s.incomeType = &incomeType
	return s
}

func (s *FApiGetIncomeService) StartTime(startTime int64) *FApiGetIncomeService {
	s.startTime = &startTime
	return s
}

func (s *FApiGetIncomeService) EndTime(endTime int64) *FApiGetIncomeService {
	s.endTime = &endTime
	return s
}

func (s *FApiGetIncomeService) Page(page int) *FApiGetIncomeService {
	s.page = &page
	return s
}

func (s *FApiGetIncomeService) Limit(limit int) *FApiGetIncomeService {
	s.limit = &limit
	return s
}

func (s *FApiGetIncomeService) Do(ctx context.Context, opts ...RequestOption) ([]Income, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/income",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.incomeType != nil {
		r.setParam("incomeType", *s.incomeType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []Income
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
