package okx

import (
	"context"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// GetAccountBalanceService 用于查询欧易账户信息
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-balance
type GetAccountBalanceService struct {
	c *Client
}

func (c *Client) NewGetAccountBalanceService() *GetAccountBalanceService {
	return &GetAccountBalanceService{c: c}
}

type AccountBalance struct {
	TotalEq            Float64                `json:"totalEq"`            // 美金层面权益
	AdjEq              Float64                `json:"adjEq"`              // 美金层面调整后权益（可用余额）
	Imr                Float64                `json:"imr"`                // 美金层面全仓占用保证金
	Mmr                Float64                `json:"mmr"`                // 美金层面维持保证金
	MgnRatio           Float64                `json:"mgnRatio"`           // 美金层面保证金率
	NotionalUsd        Float64                `json:"notionalUsd"`        // 美金层面持仓名义价值
	NotionalUsdForSwap Float64                `json:"notionalUsdForSwap"` // 美金层面持仓名义价值（永续）
	Upl                Float64                `json:"upl"`                // 账户的未实现盈亏
	UTime              Int64                  `json:"uTime"`              // 账户信息的更新时间
	Details            []AccountBalanceDetail `json:"details"`            // 币种维度账户信息
}

type AccountBalanceDetail struct {
	Ccy       string `json:"ccy"`       // 币种
	Eq        string `json:"eq"`        // 币种总权益
	EqUsd     string `json:"eqUsd"`     // 币种权益美金价值
	CashBal   string `json:"cashBal"`   // 币种余额
	UTime     string `json:"uTime"`     // 更新时间
	AvailBal  string `json:"availBal"`  // 可用余额
	FrozenBal string `json:"frozenBal"` // 冻结余额
	OrdFrozen string `json:"ordFrozen"` // 挂单冻结
	Upl       string `json:"upl"`       // 未实现盈亏
	Imr       string `json:"imr"`       // 币种维度全仓占用保证金
	Mmr       string `json:"mmr"`       // 币种维度全仓维持保证金
}

func (s *GetAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) (*AccountBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []AccountBalance
	if err := StrictDecode(data, &res); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrNoData
	}

	return &res[0], nil
}

// GetAccountPositionsService 查看持仓信息
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-positions
type GetAccountPositionsService struct {
	c *Client
}

func (c *Client) NewGetAccountPositionsService() *GetAccountPositionsService {
	return &GetAccountPositionsService{c: c}
}

type AccountPosition struct {
	InstId   string  `json:"instId"`   // 产品ID
	InstType string  `json:"instType"` // 产品类型
	Pos      Float64 `json:"pos"`      // 持仓数量
	AvgPx    Float64 `json:"avgPx"`    // 开仓均价
	Lever    Int64   `json:"lever"`    // 杠杆倍数
	Upl      Float64 `json:"upl"`      // 未实现盈亏
	LiqPx    Float64 `json:"liqPx"`    // 强平价格
	LastPx   Float64 `json:"last"`     // 最新成交价格
	IdxPx    Float64 `json:"idxPx"`    // 指数价格
	MarkPx   Float64 `json:"markPx"`   // 最新标记价格
	BePx     Float64 `json:"bePx"`     // 盈亏平衡价
	Imr      Float64 `json:"imr"`      // 初始保证金
	Mmr      Float64 `json:"mmr"`      // 维持保证金
	MgnRatio Float64 `json:"mgnRatio"` // 保证金率
	PosId    string  `json:"posId"`    // 持仓ID
	PosSide  string  `json:"posSide"`  // 持仓方向，long：开平仓模式开多，pos为正；short：开平仓模式开空，pos为正；net：买卖模式（交割/永续/期权：pos为正代表开多，pos为负代表开空。币币杠杆时，pos均为正，posCcy为交易货币时，代表开多；posCcy为计价货币时，代表开空。）
	Ctime    Int64   `json:"ctime"`    // 持仓创建时间（Unix时间戳毫秒）
	UTime    Int64   `json:"uTime"`    // 更新时间（Unix时间戳毫秒）
}

func (s *GetAccountPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]AccountPosition, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/positions",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []AccountPosition{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetAssetBalancesService 获取资金账户余额
// https://www.okx.com/docs-v5/zh/#funding-account-rest-api-get-balance
type GetAssetBalancesService struct {
	c *Client
}

func (c *Client) NewGetAssetBalancesService() *GetAssetBalancesService {
	return &GetAssetBalancesService{c: c}
}

type AssetBalance struct {
	Ccy       string `json:"ccy"`       // 币种
	Bal       string `json:"bal"`       // 余额
	FrozenBal string `json:"frozenBal"` // 冻结余额
	AvailBal  string `json:"availBal"`  // 可用余额
}

func (s *GetAssetBalancesService) Do(ctx context.Context, opts ...RequestOption) ([]AssetBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/asset/balances",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []AssetBalance{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SetAccountLeverageService 设置杠杆倍数
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-set-leverage
type SetAccountLeverageService struct {
	c       *Client
	instId  *string
	ccy     *string
	lever   string
	mgnMode string
	posSide *string
}

func (c *Client) NewSetAccountLeverageService(lever string, mgnMode string) *SetAccountLeverageService {
	return &SetAccountLeverageService{c: c, lever: lever, mgnMode: mgnMode}
}

func (s *SetAccountLeverageService) InstId(instId string) *SetAccountLeverageService {
	s.instId = &instId
	return s
}

func (s *SetAccountLeverageService) Ccy(ccy string) *SetAccountLeverageService {
	s.ccy = &ccy
	return s
}

func (s *SetAccountLeverageService) PosSide(posSide string) *SetAccountLeverageService {
	s.posSide = &posSide
	return s
}

func (s *SetAccountLeverageService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/account/set-leverage",
		secType:  secTypeSigned,
	}
	if s.instId != nil {
		r.setData("instId", *s.instId)
	}
	if s.ccy != nil {
		r.setData("ccy", *s.ccy)
	}
	if s.posSide != nil {
		r.setData("posSide", *s.posSide)
	}
	r.setData("lever", s.lever)
	r.setData("mgnMode", s.mgnMode)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
