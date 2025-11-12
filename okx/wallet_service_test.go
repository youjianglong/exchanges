package okx

import (
	"context"
	"testing"
	"time"
)

func TestGetCapitalWithdrawHistory(t *testing.T) {
	client := newTestClient()
	service := client.NewGetAssetWithdrawHistoryService()
	now := time.Now().UnixMilli()
	startTime := now - 86400*365*5*1000
	history, err := service.StartTime(startTime).EndTime(now).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("history: %+v", history)
}

func TestGetCapitalDepositHistory(t *testing.T) {
	client := newTestClient()
	service := client.NewGetAssetDepositHistoryService()
	now := time.Now().UnixMilli()
	startTime := now - 86400*365*5*1000
	history, err := service.StartTime(startTime).EndTime(now).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("history: %+v", history)
}
