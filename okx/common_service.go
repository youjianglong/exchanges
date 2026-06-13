package okx

import (
	"context"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

// PingService 用于检测欧易服务器连通性
type PingService struct {
	c *Client
}

func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

func (s *PingService) Do(ctx context.Context) error {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/system/status",
	}
	_, err := s.c.callAPI(ctx, r)
	return err
}

type GetMarketIndexTickersService struct {
	c        *Client
	quoteCcy *string
	instId   *string
}

func (c *Client) NewGetMarketIndexTickersService() *GetMarketIndexTickersService {
	return &GetMarketIndexTickersService{c: c}
}

func (s *GetMarketIndexTickersService) QuoteCcy(quoteCcy string) *GetMarketIndexTickersService {
	s.quoteCcy = &quoteCcy
	return s
}

func (s *GetMarketIndexTickersService) InstId(instId string) *GetMarketIndexTickersService {
	s.instId = &instId
	return s
}

type MarketIndexTicker struct {
	InstId  string `json:"instId"`  // 指数
	IdxPx   string `json:"idxPx"`   // 最新指数价格
	High24h string `json:"high24h"` // 24小时最高价格
	SodUtc0 string `json:"sodUtc0"` // UTC 0 时开盘价
	Open24h string `json:"open24h"` // 24小时开盘价格
	Low24h  string `json:"low24h"`  // 24小时最低价格
	SodUtc8 string `json:"sodUtc8"` // UTC+8 时开盘价
	Ts      string `json:"ts"`      // 数据产生时间，Unix时间戳的毫秒数格式，如 1597026383085
}

func (s *GetMarketIndexTickersService) Do(ctx context.Context, opts ...RequestOption) ([]*MarketIndexTicker, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/index-tickers",
	}
	if s.quoteCcy != nil {
		r.setParam("quoteCcy", *s.quoteCcy)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var tickers []*MarketIndexTicker
	err = StrictDecode(data, &tickers)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}

type GetInstrumentsService struct {
	c          *Client
	instType   string
	uly        *string
	instFamily *string
	instId     *string
}

func (c *Client) NewGetInstrumentsService(instType string) *GetInstrumentsService {
	return &GetInstrumentsService{c: c, instType: instType}
}

func (s *GetInstrumentsService) Uly(uly string) *GetInstrumentsService {
	s.uly = &uly
	return s
}

func (s *GetInstrumentsService) InstFamily(instFamily string) *GetInstrumentsService {
	s.instFamily = &instFamily
	return s
}

func (s *GetInstrumentsService) InstId(instId string) *GetInstrumentsService {
	s.instId = &instId
	return s
}

type Instrument struct {
	Alias            string  `json:"alias"`            // 合约日期别名，this_week：本周,next_week：次周,this_month：本月,next_month：次月,quarter：季度,next_quarter：次季度,third_quarter：第三季度,仅适用于交割,不建议使用，用户应通过 expTime 字段获取合约的交割日期
	AuctionEndTime   string  `json:"auctionEndTime"`   // 集合竞价结束时间，Unix时间戳的毫秒数格式，如 1597026383085,仅适用于通过集合竞价方式上线的币币，其余情况返回""
	BaseCcy          string  `json:"baseCcy"`          // 交易货币币种，如 BTC-USDT 中的 BTC ，仅适用于币币/币币杠杆
	Category         string  `json:"category"`         // 币种类别（已废弃）
	CtMult           string  `json:"ctMult"`           // 合约乘数，仅适用于交割/永续/期权
	CtType           string  `json:"ctType"`           // 合约类型,linear：正向合约,inverse：反向合约,仅适用于交割/永续
	CtVal            Float64 `json:"ctVal"`            // 合约面值，仅适用于交割/永续/期权
	CtValCcy         string  `json:"ctValCcy"`         // 合约面值计价币种，仅适用于交割/永续/期权
	ExpTime          string  `json:"expTime"`          // 产品下线时间,适用于币币/杠杆/交割/永续/期权，对于 交割/期权，为交割/行权日期；亦可以为产品下线时间，有变动就会推送。
	FutureSettlement bool    `json:"futureSettlement"` // 交割合约是否支持每日结算,适用于全仓交割
	InstFamily       string  `json:"instFamily"`       // 交易品种，如 BTC-USD，仅适用于杠杆/交割/永续/期权
	InstId           string  `json:"instId"`           // 产品id， 如 BTC-USDT
	InstType         string  `json:"instType"`         // 产品类型
	Lever            Int64   `json:"lever"`            // 该instId支持的最大杠杆倍数，不适用于币币、期权
	ListTime         string  `json:"listTime"`         // 上线时间，Unix时间戳的毫秒数格式，如 1597026383085
	LotSz            Float64 `json:"lotSz"`            // 下单数量精度，合约的数量单位是张，现货的数量单位是交易货币
	MaxIcebergSz     string  `json:"maxIcebergSz"`     // 冰山委托的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币
	MaxLmtAmt        Float64 `json:"maxLmtAmt"`        // 限价单的单笔最大美元价值
	MaxLmtSz         Float64 `json:"maxLmtSz"`         // 限价单的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币
	MaxMktAmt        Float64 `json:"maxMktAmt"`        // 市价单的单笔最大美元价值，仅适用于币币/币币杠杆
	MaxMktSz         Float64 `json:"maxMktSz"`         // 市价单的单笔最大委托数量，合约的数量单位是张，现货的数量单位是USDT
	MaxStopSz        Float64 `json:"maxStopSz"`        // 止盈止损市价委托的单笔最大委托数量，合约的数量单位是张，现货的数量单位是USDT
	MaxTriggerSz     Float64 `json:"maxTriggerSz"`     // 计划委托委托的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币
	MaxTwapSz        Float64 `json:"maxTwapSz"`        // 时间加权单的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币。单笔最小委托数量为 minSz*2
	MinSz            Float64 `json:"minSz"`            // 最小下单数量，合约的数量单位是张，现货的数量单位是交易货币
	OptType          string  `json:"optType"`          // 期权类型，C或P 仅适用于期权
	QuoteCcy         string  `json:"quoteCcy"`         // 计价货币币种，如 BTC-USDT 中的USDT ，仅适用于币币/币币杠杆
	SettleCcy        string  `json:"settleCcy"`        // 盈亏结算和保证金币种，如 BTC 仅适用于交割/永续/期权
	State            string  `json:"state"`            // 产品状态，live：交易中,suspend：暂停中,preopen：预上线，交割和期权合约轮转生成到开始交易；部分交易产品上线前,test：测试中（测试产品，不可交易）
	RuleType         string  `json:"ruleType"`         // 交易规则类型，normal：普通交易,pre_market：盘前交易
	Stk              string  `json:"stk"`              // 行权价格，仅适用于期权
	TickSz           Float64 `json:"tickSz"`           // 下单价格精度，如 0.0001，对于期权来说，是梯度中的最小下单价格精度，如果想要获取期权价格梯度，请使用"获取期权价格梯度"接口
	Uly              string  `json:"uly"`              // 标的指数，如 BTC-USD，仅适用于杠杆/交割/永续/期权
}

func (s *GetInstrumentsService) Do(ctx context.Context, opts ...RequestOption) ([]*Instrument, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/instruments",
	}
	r.setParam("instType", s.instType)
	if s.uly != nil {
		r.setParam("uly", *s.uly)
	}
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var instruments []*Instrument
	err = StrictDecode(data, &instruments)
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

type GetMarketTickersService struct {
	c          *Client
	instType   string
	uly        *string
	instFamily *string
}

func (c *Client) NewGetMarketTickersService(instType string) *GetMarketTickersService {
	return &GetMarketTickersService{c: c, instType: instType}
}

func (s *GetMarketTickersService) Uly(uly string) *GetMarketTickersService {
	s.uly = &uly
	return s
}

func (s *GetMarketTickersService) InstFamily(instFamily string) *GetMarketTickersService {
	s.instFamily = &instFamily
	return s
}

type MarketTicker struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Last      string `json:"last"`      // 最新成交价
	LastSz    string `json:"lastSz"`    // 最新成交的数量，0 代表没有成交量
	AskPx     string `json:"askPx"`     // 卖一价
	AskSz     string `json:"askSz"`     // 卖一价的挂单数数量
	BidPx     string `json:"bidPx"`     // 买一价
	BidSz     string `json:"bidSz"`     // 买一价的挂单数量
	Open24h   string `json:"open24h"`   // 24小时开盘价
	High24h   string `json:"high24h"`   // 24小时最高价
	Low24h    string `json:"low24h"`    // 24小时最低价
	VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位
	SodUtc0   string `json:"sodUtc0"`   // UTC 0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
	Ts        string `json:"ts"`        // ticker数据产生时间，Unix时间戳的毫秒数格式
}

func (s *GetMarketTickersService) Do(ctx context.Context, opts ...RequestOption) ([]*MarketTicker, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/tickers",
	}
	r.setParam("instType", s.instType)
	if s.uly != nil {
		r.setParam("uly", *s.uly)
	}
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var tickers []*MarketTicker
	err = StrictDecode(data, &tickers)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}

type GetFundingRateService struct {
	c      *Client
	instId string
}

func (c *Client) NewGetFundingRateService(instId string) *GetFundingRateService {
	return &GetFundingRateService{c: c, instId: instId}
}

type FundingRate struct {
	InstType        string `json:"instType"`        // 产品类型 SWAP：永续合约
	InstId          string `json:"instId"`          // 产品ID，如BTC-USD-SWAP
	Method          string `json:"method"`          // 资金费收取逻辑
	FormulaType     string `json:"formulaType"`     // 公式类型
	FundingRate     string `json:"fundingRate"`     // 资金费率
	NextFundingRate string `json:"nextFundingRate"` // 下一期预测资金费率
	FundingTime     string `json:"fundingTime"`     // 资金费时间
	NextFundingTime string `json:"nextFundingTime"` // 下一期资金费时间
	MinFundingRate  string `json:"minFundingRate"`  // 下一期的预测资金费率下限
	MaxFundingRate  string `json:"maxFundingRate"`  // 下一期的预测资金费率上限
	InterestRate    string `json:"interestRate"`    // 利率
	ImpactValue     string `json:"impactValue"`     // 深度加权金额（计价币数量）
	SettState       string `json:"settState"`       // 资金费率结算状态
	SettFundingRate string `json:"settFundingRate"` // 结算资金费率
	Premium         string `json:"premium"`         // 溢价，为合约的中间价和指数价格的差异
	Ts              string `json:"ts"`              // 数据更新时间
}

func (s *GetFundingRateService) Do(ctx context.Context, opts ...RequestOption) ([]*FundingRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/funding-rate",
	}
	r.setParam("instId", s.instId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var fundingRates []*FundingRate
	err = StrictDecode(data, &fundingRates)
	if err != nil {
		return nil, err
	}
	return fundingRates, nil
}

type GetPriceLimitService struct {
	c      *Client
	instId string
}

func (c *Client) NewGetPriceLimitService(instId string) *GetPriceLimitService {
	return &GetPriceLimitService{c: c, instId: instId}
}

type PriceLimit struct {
	InstType string `json:"instType"` // 产品类型 SPOT：币币 MARGIN：杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId   string `json:"instId"`   // 产品ID，如 BTC-USDT-SWAP
	BuyLmt   string `json:"buyLmt"`   // 最高买价，当enabled为false时，返回""
	SellLmt  string `json:"sellLmt"`  // 最低卖价，当enabled为false时，返回""
	Ts       string `json:"ts"`       // 限价数据更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	Enabled  bool   `json:"enabled"`  // 限价是否生效 true：限价生效 false：限价不生效
}

func (s *GetPriceLimitService) Do(ctx context.Context, opts ...RequestOption) ([]*PriceLimit, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/price-limit",
	}
	r.setParam("instId", s.instId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var priceLimits []*PriceLimit
	err = StrictDecode(data, &priceLimits)
	if err != nil {
		return nil, err
	}
	return priceLimits, nil
}

type GetMarkPriceService struct {
	c          *Client
	instType   string
	instFamily *string
	instId     *string
}

func (c *Client) NewGetMarkPriceService(instType string) *GetMarkPriceService {
	return &GetMarkPriceService{c: c, instType: instType}
}

func (s *GetMarkPriceService) InstFamily(instFamily string) *GetMarkPriceService {
	s.instFamily = &instFamily
	return s
}

func (s *GetMarkPriceService) InstId(instId string) *GetMarkPriceService {
	s.instId = &instId
	return s
}

type MarkPrice struct {
	InstType string  `json:"instType"` // 产品类型 SPOT：币币 MARGIN：杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId   string  `json:"instId"`   // 产品ID，如 BTC-USDT-SWAP
	MarkPx   Float64 `json:"markPx"`   // 标记价格
	Ts       Int64   `json:"ts"`       // 标记价格更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}

func (s *GetMarkPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*MarkPrice, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/mark-price",
	}
	r.setParam("instType", s.instType)
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var markPrices []*MarkPrice
	err = StrictDecode(data, &markPrices)
	if err != nil {
		return nil, err
	}
	return markPrices, nil
}
