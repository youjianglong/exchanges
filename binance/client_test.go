package binance

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/youjianglong/exchanges/common"
)

var (
	apiKey     = flag.String("api-key", "", "api key")
	apiSecret  = flag.String("api-secret", "", "api secret")
	proxyURL   = flag.String("proxy-url", "", "proxy url")
	symbol     = flag.String("symbol", "ETHUSDT", "symbol")
	startTime  = flag.Int64("start-time", time.Now().Add(-time.Hour*24).UnixMilli(), "start time")
	endTime    = flag.Int64("end-time", time.Now().UnixMilli(), "end time")
	limit      = flag.Int("limit", 100, "limit")
	incomeType = flag.String("income-type", "", "income type")
)

func newHttpClient() *http.Client {
	if *proxyURL == "" {
		return NewHttpClient(time.Second*10, nil)
	}
	transport := CloneHttpTransport()
	transport.Proxy = func(r *http.Request) (*url.URL, error) {
		return url.Parse(*proxyURL)
	}
	return &http.Client{
		Transport: transport,
	}
}

func readEnvFile(file string) (map[string]string, error) {
	env, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	envs := map[string]string{}
	lines := strings.Split(string(env), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envs[key] = value
		}
	}
	return envs, nil
}

func assignStr(envs map[string]string, key string, value *string) {
	if v, ok := envs[key]; ok {
		*value = v
	}
}

func assignInt64(envs map[string]string, key string, value *int64) {
	if v, ok := envs[key]; ok {
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println("read env", key, ", parseInt64", err)
			return
		}
		*value = vv
	}
}

func assignInt(envs map[string]string, key string, value *int) {
	if v, ok := envs[key]; ok {
		vv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("read env", key, ", parseInt:", err)
			return
		}
		*value = vv
	}
}

func readEnv() {
	envs, err := readEnvFile(".env")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("read env:", err)
		}
		return
	}

	assignStr(envs, "API_KEY", apiKey)
	assignStr(envs, "API_SECRET", apiSecret)
	assignStr(envs, "PROXY_URL", proxyURL)
	assignInt64(envs, "START_TIME", startTime)
	assignInt64(envs, "END_TIME", endTime)
	assignInt(envs, "LIMIT", limit)
	assignStr(envs, "INCOME_TYPE", incomeType)
	assignStr(envs, "SYMBOL", symbol)
}

func newTestClient() *Client {
	readEnv()

	httpClient := NewHttpClient(time.Second*10, nil)

	client := NewClientWithHttpClient(
		*apiKey,
		*apiSecret,
		httpClient)
	client.Debug = true
	slog.SetLogLoggerLevel(slog.LevelDebug)
	return client
}
