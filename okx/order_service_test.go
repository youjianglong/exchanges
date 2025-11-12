package okx

import (
	"context"
	"testing"
	"time"
)

func TestGetSpotOrder(t *testing.T) {
	client := newTestClient()
	orderService := client.NewGetOrderHistoryService()
	orderService.InstType("SPOT")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	order, err := orderService.Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("order: %+v", order)
}

func TestGetSpotOpenOrders(t *testing.T) {
	client := newTestClient()
	service := client.NewGetOrdersPendingService().InstType("SPOT").Limit(100)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	orders, err := service.Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, order := range orders {
		t.Logf("order: %+v", order)
	}
}
