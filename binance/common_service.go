package binance

import (
	"context"
	"net/http"

	. "github.com/youjianglong/exchanges/common"
)

type PingService struct {
	c *Client
}

func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

func (s *PingService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/ping",
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

type PostService struct {
	c *Client

	baseURL  *string
	endpoint string
	params   params
}

func (c *Client) NewPostService(baseURL *string, endpoint string, params params) *PostService {
	return &PostService{c: c, baseURL: baseURL, endpoint: endpoint, params: params}
}

func (s *PostService) Do(ctx context.Context, dest any, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: s.endpoint,
		secType:  secTypeSigned,
		baseURL:  s.baseURL,
	}
	r.setFormParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return StrictDecode(data, dest)
}

type GetService struct {
	c *Client

	baseURL  *string
	endpoint string
	params   params
}

func (c *Client) NewGetService(baseURL *string, endpoint string, params params) *GetService {
	return &GetService{c: c, baseURL: baseURL, endpoint: endpoint, params: params}
}

func (s *GetService) Do(ctx context.Context, dest any, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodGet,
		endpoint: s.endpoint,
		secType:  secTypeSigned,
		baseURL:  s.baseURL,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return StrictDecode(data, dest)
}

func (c *Client) NewFApiPostService(endpoint string, params params) *PostService {
	return &PostService{c: c, baseURL: &c.FApiBaseURL, endpoint: endpoint, params: params}
}

func (c *Client) NewFApiGetService(endpoint string, params params) *GetService {
	return &GetService{c: c, baseURL: &c.FApiBaseURL, endpoint: endpoint, params: params}
}

func (c *Client) NewPApiPostService(endpoint string, params params) *PostService {
	return &PostService{c: c, baseURL: &c.PApiBaseURL, endpoint: endpoint, params: params}
}

func (c *Client) NewPApiGetService(endpoint string, params params) *GetService {
	return &GetService{c: c, baseURL: &c.PApiBaseURL, endpoint: endpoint, params: params}
}

func (c *Client) NewSpotPostService(endpoint string, params params) *PostService {
	return &PostService{c: c, baseURL: &c.ApiBaseURL, endpoint: endpoint, params: params}
}

func (c *Client) NewSpotGetService(endpoint string, params params) *GetService {
	return &GetService{c: c, baseURL: &c.ApiBaseURL, endpoint: endpoint, params: params}
}
