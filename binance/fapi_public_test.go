package binance

import (
	"context"
	"testing"
	"time"
)

func TestFApiGetKLines(t *testing.T) {
	client := newTestClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()
	klines, err := client.NewFApiGetKLinesService("BTCUSDT", "1s").StartTime(now.Add(-1 * time.Minute).UnixMilli()).EndTime(now.UnixMilli()).Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, kline := range klines {
		t.Logf("%+v", kline)
	}
}
