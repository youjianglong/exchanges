package okx

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/youjianglong/exchanges/common"
)

type fundingRateInfo struct {
	Value    float64
	CurTime  int64
	NextTime int64
}

type OkxPublicStreamAdapter struct {
	*AdapterWebsocket
	client *Client
}

func NewOkxPublicStreamAdapter(proxyUrl *url.URL) *OkxPublicStreamAdapter {
	s := &OkxPublicStreamAdapter{}
	httpClient := common.NewHttpClient(16*time.Second, proxyUrl)
	s.client = NewClient("", "", "")
	s.client.SetProxyURL(proxyUrl)
	s.AdapterWebsocket = NewAdapterWebsocket(
		wsPublicMainUrl,
		nil,
		slog.With("exchange", "okx"),
		httpClient,
	)

	return s
}

func (s *OkxPublicStreamAdapter) Start() error {
	// if err := s.subscribeSpotTickers(); err != nil {
	// 	return err
	// }
	return s.AdapterWebsocket.Start()
}

func (s *OkxPublicStreamAdapter) Stop() {
	s.AdapterWebsocket.Stop()
}

func (s *OkxPublicStreamAdapter) subscribeSpotTickers(spotSymbols, swapSymbols []string) error {
	args := []any{}
	for _, symbol := range spotSymbols {
		args = append(args, H{"channel": "tickers", "instId": symbol})
	}
	for _, symbol := range swapSymbols {
		args = append(args, H{"channel": "tickers", "instId": symbol + "-SWAP"})
	}
	err := s.Subscribe(args...)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}
	s.RegisterHandler("tickers", s.handleTickers)
	return nil
}

func (s *OkxPublicStreamAdapter) handleTickers(event *wsEvent) {
	var tickers []*MarketTicker
	err := json.Unmarshal(event.Data, &tickers)
	if err != nil {
		s.logger.Error("unmarshal tickers: " + err.Error())
		return
	}
	for _, ticker := range tickers {
		isSwap := ticker.InstType == "SWAP"
		_ = isSwap
	}
}
