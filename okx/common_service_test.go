package okx

import (
	"context"
	"testing"
)

func TestGetInstrumentsService(t *testing.T) {
	c := newTestClient()
	insts, err := c.NewGetInstrumentsService("SPOT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, inst := range insts {
		t.Logf("%#v", inst)
	}
}

func TestGetPriceLimitService(t *testing.T) {
	c := newTestClient()
	priceLimits, err := c.NewGetPriceLimitService("BTC-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, priceLimit := range priceLimits {
		t.Logf("%#v", priceLimit)
	}
}

func TestGetFundingRateService(t *testing.T) {
	c := newTestClient()
	fundingRate, err := c.NewGetFundingRateService("BTC-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, fr := range fundingRate {
		t.Logf("%+v\n", fr)
	}
}

func TestGetMarketTickersService(t *testing.T) {
	c := newTestClient()
	tickers, err := c.NewGetMarketTickersService("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, ticker := range tickers {
		t.Logf("%+v", ticker)
	}
}
