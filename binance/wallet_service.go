package binance

import (
	"context"
	"net/http"
	"net/url"

	. "github.com/youjianglong/exchanges/common"
)

// GetWalletBalanceService 获取钱包余额
type GetWalletBalanceService struct {
	c *Client

	quoteAsset *string
}

func (c *Client) NewGetWalletBalanceService() *GetWalletBalanceService {
	return &GetWalletBalanceService{c: c}
}

func (s *GetWalletBalanceService) QuoteAsset(quoteAsset string) *GetWalletBalanceService {
	s.quoteAsset = &quoteAsset
	return s
}

// WalletBalance 钱包余额
type WalletBalance struct {
	Activate   bool    `json:"activate"`
	Balance    Float64 `json:"balance"`
	WalletName string  `json:"walletName"`
}

func (s *GetWalletBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]WalletBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/wallet/balance",
		secType:  secTypeSigned,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.quoteAsset != nil {
		query := url.Values{}
		query.Set("quoteAsset", *s.quoteAsset)
		r.query = query
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []WalletBalance{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type GetUserAssetService struct {
	c *Client
}

func (c *Client) NewGetUserAssetService() *GetUserAssetService {
	return &GetUserAssetService{c: c}
}

type UserAsset struct {
	Asset       string  `json:"asset"`       // 资产名称
	Free        Float64 `json:"free"`        // 可用余额
	Locked      Float64 `json:"locked"`      // 锁定余额
	Freeze      Float64 `json:"freeze"`      // 冻结余额
	Withdrawing Float64 `json:"withdrawing"` // 提现中金额
	// Ipoable      Float64 `json:"ipoable"`      // IPO可用余额
	// BtcValuation Float64 `json:"btcValuation"` // BTC估值
}

func (s *GetUserAssetService) Do(ctx context.Context, opts ...RequestOption) ([]UserAsset, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v3/asset/getUserAsset",
		secType:  secTypeSigned,
		baseURL:  &s.c.ApiBaseURL,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []UserAsset{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// CapitalWithdrawHistoryService 获取资金提现历史
type CapitalWithdrawHistoryService struct {
	c *Client

	coin            *string
	withdrawOrderId *string
	status          *int
	offset          *int
	limit           *int
	idList          *[]string
	startTime       *int64
	endTime         *int64
}

func (s *CapitalWithdrawHistoryService) Coin(coin string) *CapitalWithdrawHistoryService {
	s.coin = &coin
	return s
}

func (s *CapitalWithdrawHistoryService) WithdrawOrderId(withdrawOrderId string) *CapitalWithdrawHistoryService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

func (s *CapitalWithdrawHistoryService) Status(status int) *CapitalWithdrawHistoryService {
	s.status = &status
	return s
}

func (s *CapitalWithdrawHistoryService) Offset(offset int) *CapitalWithdrawHistoryService {
	s.offset = &offset
	return s
}

func (s *CapitalWithdrawHistoryService) Limit(limit int) *CapitalWithdrawHistoryService {
	s.limit = &limit
	return s
}

func (s *CapitalWithdrawHistoryService) IdList(idList []string) *CapitalWithdrawHistoryService {
	s.idList = &idList
	return s
}

func (s *CapitalWithdrawHistoryService) StartTime(startTime int64) *CapitalWithdrawHistoryService {
	s.startTime = &startTime
	return s
}

func (s *CapitalWithdrawHistoryService) EndTime(endTime int64) *CapitalWithdrawHistoryService {
	s.endTime = &endTime
	return s
}

func (c *Client) NewGetCapitalWithdrawHistoryService() *CapitalWithdrawHistoryService {
	return &CapitalWithdrawHistoryService{c: c}
}

type CapitalWithdrawHistory struct {
	ID              string `json:"id"`              // 该笔提现在币安的id
	Amount          string `json:"amount"`          // 提现转出金额
	TransactionFee  string `json:"transactionFee"`  // 手续费
	Coin            string `json:"coin"`            // 币种
	Status          int    `json:"status"`          // 状态
	Address         string `json:"address"`         // 提现地址
	TxID            string `json:"txId"`            // 提现交易id
	ApplyTime       string `json:"applyTime"`       // UTC 时间
	Network         string `json:"network"`         // 网络
	TransferType    int    `json:"transferType"`    // 1: 站内转账, 0: 站外转账
	WithdrawOrderID string `json:"withdrawOrderId"` // 自定义ID, 如果没有则不返回该字段
	Info            string `json:"info"`            // 提币失败原因
	ConfirmNo       int    `json:"confirmNo"`       // 提现确认数
	WalletType      int    `json:"walletType"`      // 1: 资金钱包 0:现货钱包
	TxKey           string `json:"txKey"`           // 交易key
	CompleteTime    string `json:"completeTime"`    // 提现完成，成功下账时间(UTC)
}

func (s *CapitalWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]CapitalWithdrawHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/withdraw/history",
		secType:  secTypeSigned,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.withdrawOrderId != nil {
		r.setParam("withdrawOrderId", *s.withdrawOrderId)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.idList != nil {
		r.setParam("idList", *s.idList)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []CapitalWithdrawHistory{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// CapitalDepositHistoryService 获取资金充值历史
type CapitalDepositHistoryService struct {
	c *Client

	includeSource *bool   // 默认 false，如果为true时会返回sourceAddress字段
	coin          *string // 币种
	status        *int    // 状态
	offset        *int    // 偏移量
	limit         *int    // 限制
	startTime     *int64  // 开始时间
	endTime       *int64  // 结束时间
	txId          *string // 交易ID
}

func (c *Client) NewGetCapitalDepositHistoryService() *CapitalDepositHistoryService {
	return &CapitalDepositHistoryService{c: c}
}

func (s *CapitalDepositHistoryService) IncludeSource(includeSource bool) *CapitalDepositHistoryService {
	s.includeSource = &includeSource
	return s
}

func (s *CapitalDepositHistoryService) Coin(coin string) *CapitalDepositHistoryService {
	s.coin = &coin
	return s
}

func (s *CapitalDepositHistoryService) Status(status int) *CapitalDepositHistoryService {
	s.status = &status
	return s
}

func (s *CapitalDepositHistoryService) Offset(offset int) *CapitalDepositHistoryService {
	s.offset = &offset
	return s
}

func (s *CapitalDepositHistoryService) Limit(limit int) *CapitalDepositHistoryService {
	s.limit = &limit
	return s
}

func (s *CapitalDepositHistoryService) StartTime(startTime int64) *CapitalDepositHistoryService {
	s.startTime = &startTime
	return s
}

func (s *CapitalDepositHistoryService) EndTime(endTime int64) *CapitalDepositHistoryService {
	s.endTime = &endTime
	return s
}

func (s *CapitalDepositHistoryService) TxId(txId string) *CapitalDepositHistoryService {
	s.txId = &txId
	return s
}

type CapitalDepositHistory struct {
	ID            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxID          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int    `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int    `json:"unlockConfirm"`
	WalletType    int    `json:"walletType"`
}

func (s *CapitalDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]CapitalDepositHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
		baseURL:  &s.c.ApiBaseURL,
	}
	if s.includeSource != nil {
		r.setParam("includeSource", *s.includeSource)
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []CapitalDepositHistory{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
