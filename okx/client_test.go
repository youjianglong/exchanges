package okx

import (
	"github.com/youjianglong/exchanges/common"
)

func newTestClient() *Client {
	proxyPool, err := common.NewProxyPool([]string{"socks5://127.0.0.1:10808"})
	if err != nil {
		panic(err)
	}
	httpClient, err := common.NewHttpClientWithProxy(proxyPool.GetCurrentProxy().URL)
	if err != nil {
		panic(err)
	}
	client := NewClientWithHttpClient(
		"",
		"",
		"",
		true,
		httpClient)
	client.Debug = true
	return client
}
