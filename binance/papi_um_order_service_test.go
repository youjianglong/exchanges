package binance

import (
	"context"
	"fmt"
	"os"
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

func writeCsvRow(fp *os.File, values ...any) {
	for i, v := range values {
		if i > 0 {
			fp.WriteString(",")
		}
		fmt.Fprintf(fp, "%v", v)
	}
	fp.WriteString("\n")
}

func TestExportPApiUmTrades(t *testing.T) {
	client := newTestClient()
	limit := 1000
	tradeService := client.NewPApiUmGetUserTradesService(*symbol).Limit(limit)
	fp, err := os.Create(*symbol + "-trades.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()
	writeCsvRow(fp, "OrderID", "TradeID", "Price", "Qty", "RealizedPnl", "Commission", "Time")
	start := *startTime
	for start < *endTime {
		end := start + 6*60*60*1000 // 6 hours
		tradeService = tradeService.StartTime(start).EndTime(end)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trades, err := tradeService.Do(ctx)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
		for _, trade := range trades {
			writeCsvRow(fp, trade.OrderID, trade.ID, trade.Price, trade.Qty, trade.RealizedPnl, trade.Commission, time.UnixMilli(trade.Time.Value()).Format(time.DateTime))
		}
		start = end
	}
}
