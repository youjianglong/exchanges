package binance

import (
	"context"
	"testing"
	"time"
)

func TestGetWalletBalance(t *testing.T) {
	client := newTestClient()
	service := client.NewGetWalletBalanceService()
	service.QuoteAsset("USDT")
	assets, err := service.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("assets: %+v", assets)
}

func TestGetUserAsset(t *testing.T) {
	client := newTestClient()
	service := client.NewGetUserAssetService()
	assets, err := service.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("assets: %+v", assets)
}

func TestGetCapitalWithdrawHistory(t *testing.T) {
	client := newTestClient()
	service := client.NewGetCapitalWithdrawHistoryService()
	history, err := service.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("history: %+v", history)
}

func TestGetCapitalDepositHistory(t *testing.T) {
	client := newTestClient()
	service := client.NewGetCapitalDepositHistoryService()
	ti := time.Now()
	ti = ti.Add(time.Hour * 24 * -90)
	service.StartTime(ti.UnixMilli())
	service.EndTime(ti.Add(time.Hour * 24 * 89).UnixMilli())
	history, err := service.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("history: %+v", history)
}
