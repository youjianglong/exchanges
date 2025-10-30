package binance

import (
	"context"
	"testing"
	"time"
)

func TestPApiUmGetAllOrders(t *testing.T) {
	client := newTestClient()
	orderService := client.NewPApiGetAllOrdersService(*symbol).StartTime(*startTime).EndTime(*endTime).Limit(*limit)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	orders, err := orderService.Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, order := range orders {
		t.Logf("order: %+v", order)
	}
}

func TestPApiUmGetTrades(t *testing.T) {
	client := newTestClient()
	tradeService := client.NewPApiUmGetUserTradesService(*symbol).StartTime(*startTime).EndTime(*endTime).Limit(*limit)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	trades, err := tradeService.Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, trade := range trades {
		t.Logf("  %+v", trade)
	}
}
