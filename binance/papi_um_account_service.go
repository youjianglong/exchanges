package binance

import (
	"context"
	"encoding/json"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// 获取U本位账户信息
// https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/account/Account-Information
type GetUMAccountService struct {
	c *Client
}

func (c *Client) NewGetUMAccountService() *GetUMAccountService {
	return &GetUMAccountService{c: c}
}

type UMAccount struct {
	UniMMR                   Float64 `json:"uniMMR"`                   // 统一账户维持保证金率
	AccountEquity            Float64 `json:"accountEquity"`            // 账户权益（美元）
	ActualEquity             Float64 `json:"actualEquity"`             // 实际权益（不包含抵押）
	AccountInitialMargin     Float64 `json:"accountInitialMargin"`     // 账户初始保证金
	AccountMaintMargin       Float64 `json:"accountMaintMargin"`       // 账户维持保证金（美元）
	AccountStatus            string  `json:"accountStatus"`            // 账户状态："NORMAL", "MARGIN_CALL", "SUPPLY_MARGIN", "REDUCE_ONLY", "ACTIVE_LIQUIDATION", "FORCE_LIQUIDATION", "BANKRUPTED"
	VirtualMaxWithdrawAmount Float64 `json:"virtualMaxWithdrawAmount"` // 虚拟最大提现金额
	TotalAvailableBalance    Float64 `json:"totalAvailableBalance"`    // 最高可转出金额（美元）
	TotalMarginOpenLoss      Float64 `json:"totalMarginOpenLoss"`      // 美元保证金未结订单
	UpdateTime               int64   `json:"updateTime"`               // 更新时间
}

func (s *GetUMAccountService) Do(ctx context.Context, opts ...RequestOption) (*UMAccount, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(UMAccount)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// 获取U本位账户详情
// https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/account/Get-UM-Account-Detail
type GetUMAccountDetailService struct {
	c *Client
}

func (c *Client) NewGetUMAccountDetailService() *GetUMAccountDetailService {
	return &GetUMAccountDetailService{c: c}
}

type UMAsset struct {
	Asset                  string  `json:"asset"`                  // 资产
	CrossWalletBalance     Float64 `json:"crossWalletBalance"`     // 全仓余额
	CrossUnPnl             Float64 `json:"crossUnPnl"`             // 全仓未实现盈亏
	MaintMargin            Float64 `json:"maintMargin"`            // 维持保证金
	InitialMargin          Float64 `json:"initialMargin"`          // 初始保证金
	PositionInitialMargin  Float64 `json:"positionInitialMargin"`  // 持仓初始保证金
	OpenOrderInitialMargin Float64 `json:"openOrderInitialMargin"` // 挂单初始保证金
	UpdateTime             Int64   `json:"updateTime"`             // 更新时间
}

type UMAccountDetail struct {
	Assets    []UMAsset    `json:"assets"`
	Positions []UMPosition `json:"positions"`
}

func (s *GetUMAccountDetailService) Do(ctx context.Context, opts ...RequestOption) (*UMAccountDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(UMAccountDetail)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetPApiAccountBalanceService 获取账户余额
// https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/account
type GetPApiAccountBalanceService struct {
	c     *Client
	asset *string
}

func (c *Client) NewGetPApiAccountBalanceService() *GetPApiAccountBalanceService {
	return &GetPApiAccountBalanceService{c: c}
}

func (s *GetPApiAccountBalanceService) Asset(asset string) *GetPApiAccountBalanceService {
	s.asset = &asset
	return s
}

type PApiAccountBalance struct {
	Asset               string  `json:"asset"`               // 资产
	TotalWalletBalance  string  `json:"totalWalletBalance"`  // 钱包余额 = 全仓杠杆未锁定 + 全仓杠杆锁定 + u本位合约钱包余额 + 币本位合约钱包余额
	CrossMarginAsset    Float64 `json:"crossMarginAsset"`    // 全仓资产 = 全仓杠杆未锁定 + 全仓杠杆锁定
	CrossMarginBorrowed string  `json:"crossMarginBorrowed"` // 全仓杠杆借贷
	CrossMarginFree     Float64 `json:"crossMarginFree"`     // 全仓杠杆未锁定
	CrossMarginInterest string  `json:"crossMarginInterest"` // 全仓杠杆利息
	CrossMarginLocked   Float64 `json:"crossMarginLocked"`   // 全仓杠杆锁定
	UmWalletBalance     Float64 `json:"umWalletBalance"`     // u本位合约钱包余额
	UmUnrealizedPNL     string  `json:"umUnrealizedPNL"`     // u本位未实现盈亏
	CmWalletBalance     Float64 `json:"cmWalletBalance"`     // 币本位合约钱包余额
	CmUnrealizedPNL     string  `json:"cmUnrealizedPNL"`     // 币本位未实现盈亏
	UpdateTime          int64   `json:"updateTime"`          // 更新时间
	NegativeBalance     string  `json:"negativeBalance"`     // 负余额
}

func (s *GetPApiAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]*PApiAccountBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
		baseURL:  &s.c.PApiBaseURL,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*PApiAccountBalance
	if s.asset != nil {
		r := new(PApiAccountBalance)
		err = json.Unmarshal(data, r)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	} else {
		err = json.Unmarshal(data, &res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

type PapiUMGetPositionRiskService struct {
	c *Client

	symbol *string
}

func (c *Client) NewPapiUMGetPositionRiskService() *PapiUMGetPositionRiskService {
	return &PapiUMGetPositionRiskService{c: c}
}

func (s *PapiUMGetPositionRiskService) Symbol(symbol string) *PapiUMGetPositionRiskService {
	s.symbol = &symbol
	return s
}

type UMPosition struct {
	Symbol                 string  `json:"symbol"`                 // 交易对
	InitialMargin          Float64 `json:"initialMargin"`          // 当前标记价格下的初始保证金
	MaintMargin            Float64 `json:"maintMargin"`            // 维持保证金
	UnrealizedProfit       Float64 `json:"unrealizedProfit"`       // 未实现盈亏
	UnRealizedProfit       Float64 `json:"unRealizedProfit"`       // 未实现盈亏
	PositionInitialMargin  string  `json:"positionInitialMargin"`  // 当前标记价格下持仓所需初始保证金
	OpenOrderInitialMargin string  `json:"openOrderInitialMargin"` // 当前标记价格下挂单所需初始保证金
	Leverage               Int64   `json:"leverage"`               // 当前杠杆倍数
	EntryPrice             Float64 `json:"entryPrice"`             // 平均入场价格
	MarkPrice              Float64 `json:"markPrice"`              // 当前标记价格
	LiquidationPrice       Float64 `json:"liquidationPrice"`       // 预估强平价格
	MaxNotional            string  `json:"maxNotional"`            // 当前杠杆下最大可用名义价值
	MaxNotionalValue       Float64 `json:"maxNotionalValue"`       // 当前杠杆下最大可用名义价值
	Notional               Float64 `json:"notional"`               // 名义价值
	BidNotional            string  `json:"bidNotional"`            // 买单名义价值(忽略)
	AskNotional            string  `json:"askNotional"`            // 卖单名义价值(忽略)
	PositionSide           string  `json:"positionSide"`           // 持仓方向
	PositionAmt            Float64 `json:"positionAmt"`            // 持仓数量
	UpdateTime             Int64   `json:"updateTime"`             // 最后更新时间
}

func (s *PapiUMGetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) ([]UMPosition, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/positionRisk",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res []UMPosition
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type Income struct {
	Symbol     string  `json:"symbol"`     // 交易对，仅针对涉及交易对的资金流
	IncomeType string  `json:"incomeType"` // 资金流类型
	Income     Float64 `json:"income"`     // 资金流数量，正数代表流入，负数代表流出
	Asset      string  `json:"asset"`      // 资产内容
	Info       string  `json:"info"`       // 备注信息，取决于流水类型
	Time       Int64   `json:"time"`       // 时间
	TranId     Int64   `json:"tranId"`     // 划转ID
	TradeId    Int64   `json:"tradeId"`    // 引起流水产生的原始交易ID
}

type PApiUmGetIncomeService struct {
	c *Client

	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	limit      *int
}

func (c *Client) NewPApiUmGetIncomeService() *PApiUmGetIncomeService {
	return &PApiUmGetIncomeService{c: c}
}

func (s *PApiUmGetIncomeService) Symbol(symbol string) *PApiUmGetIncomeService {
	s.symbol = &symbol
	return s
}

func (s *PApiUmGetIncomeService) IncomeType(incomeType string) *PApiUmGetIncomeService {
	s.incomeType = &incomeType
	return s
}

func (s *PApiUmGetIncomeService) StartTime(startTime int64) *PApiUmGetIncomeService {
	s.startTime = &startTime
	return s
}

func (s *PApiUmGetIncomeService) EndTime(endTime int64) *PApiUmGetIncomeService {
	s.endTime = &endTime
	return s
}

func (s *PApiUmGetIncomeService) Limit(limit int) *PApiUmGetIncomeService {
	s.limit = &limit
	return s
}

func (s *PApiUmGetIncomeService) Do(ctx context.Context, opts ...RequestOption) ([]Income, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/income",
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

type PApiUmChangeLeverageService struct {
	c *Client

	symbol   string
	leverage int64
}

func (c *Client) NewPApiUmChangeLeverageService(symbol string, leverage int64) *PApiUmChangeLeverageService {
	return &PApiUmChangeLeverageService{c: c, symbol: symbol, leverage: leverage}
}

func (s *PApiUmChangeLeverageService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("leverage", s.leverage)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

type PApiGetUMPositionsService struct {
	uaSvc *GetUMAccountDetailService
}

func (c *Client) NewPApiGetUMPositionsService() *PApiGetUMPositionsService {
	return &PApiGetUMPositionsService{uaSvc: c.NewGetUMAccountDetailService()}
}

func (s *PApiGetUMPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*Position, error) {
	detail, err := s.uaSvc.Do(ctx, opts...)
	if err != nil {
		return nil, err
	}
	var positions []*Position
	for _, p := range detail.Positions {
		if p.PositionAmt.Value() == 0 {
			continue
		}
		positions = append(positions, &Position{
			Symbol:           p.Symbol,
			PositionSide:     p.PositionSide,
			PositionAmt:      p.PositionAmt,
			EntryPrice:       p.EntryPrice,
			Leverage:         p.Leverage,
			UnRealizedProfit: p.UnrealizedProfit,
			InitialMargin:    p.InitialMargin,
			MaintMargin:      p.MaintMargin,
			MarkPrice:        p.MarkPrice,
			Notional:         p.Notional,
			LiquidationPrice: p.LiquidationPrice,
			UpdateTime:       p.UpdateTime,
		})
	}
	return positions, nil
}
