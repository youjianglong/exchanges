package okx

import (
	"context"
	"net/http"
	"strconv"

	. "github.com/youjianglong/exchanges/common"
)

// GetAssetWithdrawHistoryService 获取提币记录
// https://www.okx.com/docs-v5/zh/#funding-account-rest-api-get-withdrawal-history
type GetAssetWithdrawHistoryService struct {
	c *Client

	after  *string // 在此时间之前
	before *string // 在此时间之后
	state  *string // 状态，2=提币成功
	limit  *int    // 每页条数，最大100
}

func (s *GetAssetWithdrawHistoryService) After(after string) *GetAssetWithdrawHistoryService {
	s.after = &after
	return s
}

func (s *GetAssetWithdrawHistoryService) Before(before string) *GetAssetWithdrawHistoryService {
	s.before = &before
	return s
}

func (s *GetAssetWithdrawHistoryService) StartTime(startTime int64) *GetAssetWithdrawHistoryService {
	// before 表示在此时间之后的记录，即 startTime 之后的记录
	before := strconv.FormatInt(startTime, 10)
	s.before = &before
	return s
}

func (s *GetAssetWithdrawHistoryService) EndTime(endTime int64) *GetAssetWithdrawHistoryService {
	// after 表示在此时间之前的记录，即 endTime 之前的记录
	after := strconv.FormatInt(endTime, 10)
	s.after = &after
	return s
}

func (s *GetAssetWithdrawHistoryService) State(state string) *GetAssetWithdrawHistoryService {
	s.state = &state
	return s
}

func (s *GetAssetWithdrawHistoryService) Limit(limit int) *GetAssetWithdrawHistoryService {
	s.limit = &limit
	return s
}

func (c *Client) NewGetAssetWithdrawHistoryService() *GetAssetWithdrawHistoryService {
	return &GetAssetWithdrawHistoryService{c: c}
}

type AssetWithdrawHistory struct {
	TxId string `json:"txId"` // 提币哈希记录
	Ccy  string `json:"ccy"`  // 币种
	Amt  string `json:"amt"`  // 数量
	Ts   string `json:"ts"`   // 提币时间
}

func (s *GetAssetWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]AssetWithdrawHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/asset/withdrawal-history",
		secType:  secTypeSigned,
	}

	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	if s.state != nil {
		r.setParam("state", *s.state)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []AssetWithdrawHistory{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetAssetDepositHistoryService 获取资金充值历史
// https://www.okx.com/docs-v5/zh/#funding-account-rest-api-get-deposit-history
type GetAssetDepositHistoryService struct {
	c *Client

	after  *string // 在此时间之前
	before *string // 在此时间之后
	state  *string // 状态，2=充值成功
	limit  *int    // 每页条数，最大100
}

func (c *Client) NewGetAssetDepositHistoryService() *GetAssetDepositHistoryService {
	return &GetAssetDepositHistoryService{c: c}
}

func (s *GetAssetDepositHistoryService) After(after string) *GetAssetDepositHistoryService {
	s.after = &after
	return s
}

func (s *GetAssetDepositHistoryService) Before(before string) *GetAssetDepositHistoryService {
	s.before = &before
	return s
}

func (s *GetAssetDepositHistoryService) StartTime(startTime int64) *GetAssetDepositHistoryService {
	// before 表示在此时间之后的记录，即 startTime 之后的记录
	before := strconv.FormatInt(startTime, 10)
	s.before = &before
	return s
}

func (s *GetAssetDepositHistoryService) EndTime(endTime int64) *GetAssetDepositHistoryService {
	// after 表示在此时间之前的记录，即 endTime 之前的记录
	after := strconv.FormatInt(endTime, 10)
	s.after = &after
	return s
}

func (s *GetAssetDepositHistoryService) State(state string) *GetAssetDepositHistoryService {
	s.state = &state
	return s
}

func (s *GetAssetDepositHistoryService) Limit(limit int) *GetAssetDepositHistoryService {
	s.limit = &limit
	return s
}

type AssetDepositHistory struct {
	TxId string `json:"txId"` // 区块转账哈希记录
	Ccy  string `json:"ccy"`  // 币种
	Amt  string `json:"amt"`  // 数量
	Ts   string `json:"ts"`   // 充值时间
}

func (s *GetAssetDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]AssetDepositHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/asset/deposit-history",
		secType:  secTypeSigned,
	}
	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	if s.state != nil {
		r.setParam("state", *s.state)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []AssetDepositHistory{}
	if err = StrictDecode(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
