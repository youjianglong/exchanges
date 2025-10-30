package binance

import (
	"context"
	"net/http"
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
