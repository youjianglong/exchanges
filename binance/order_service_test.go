package binance

import (
	"context"
	"testing"
	"time"
)

func TestGetSpotOrders(t *testing.T) {
	client := newTestClient()
	orderService := client.NewGetSpotOrdersService()
	orderService.Symbol(*symbol)
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
	orderService := client.NewGetSpotOpenOrdersService()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	orders, err := orderService.Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("orders: %+v", orders)
}

func TestGetPApiSwapOpenOrders(t *testing.T) {
	client := newTestClient()
	orderService := client.NewGetPApiSwapOpenOrdersService()
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

func TestGetFApiSwapOrders(t *testing.T) {
	client := newTestClient()
	orderService := client.NewGetFApiSwapOrdersService()
	orderService.Symbol(*symbol)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	order, err := orderService.Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("order: %+v", order)
}

func TestGetFApiSwapOpenOrders(t *testing.T) {
	client := newTestClient()
	orderService := client.NewGetFApiSwapOpenOrdersService()
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
