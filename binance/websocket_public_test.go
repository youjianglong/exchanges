package binance

import (
	"testing"
	"time"
)

func TestWsSpotPublicService(t *testing.T) {
	ws := NewWsSpotPublicService("")
	ws.SetHttpClient(newHttpClient())
	ws.SubscribeSymbolsMiniTicker(func(event WsMiniTickerEvent) {
		t.Log(event)
	}, "BTCUSDT", "ETHUSDT")
	err := ws.Start()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 30)
	ws.ws.Restart()
	time.Sleep(time.Second * 60)
}

func TestSwapPublicService(t *testing.T) {
	ws := NewWsSwapPublicService("")
	ws.SetHttpClient(newHttpClient())
	ws.SubscribeMarkPrice(func(event MarkPriceEvent) {
		t.Logf("%+v", event)
	}, "BTCUSDT", "ETHUSDT")
	err := ws.Start()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 30)
	ws.ws.Restart()
	time.Sleep(time.Second * 60)
}
