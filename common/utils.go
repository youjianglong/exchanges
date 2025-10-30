package common

import (
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/youjianglong/exchanges/errorx"
)

// ConvertToFloat64 转换字符串为float64
func ConvertToFloat64(value string) (float64, error) {
	if value == "" || value == "0" || value == "0.0" || value == "0.00" {
		return 0, nil
	}
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return valueFloat, nil
}

// ConvertToInt64 转换字符串为int64
func ConvertToInt64(value string) (int64, error) {
	if value == "" || value == "0" {
		return 0, nil
	}
	return strconv.ParseInt(value, 10, 64)
}

func MustConvertToInt64(value string, logger *slog.Logger, name string) int64 {
	val, err := ConvertToInt64(value)
	if err != nil {
		logger.Error("convert %s to int64: %s", name, err)
		return 0
	}
	return val
}

// BatchConvertToFloat64Map 批量转换字符串为float64
func BatchConvertToFloat64Map(values ...string) (map[string]float64, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("values length must be even")
	}
	if len(values) == 0 {
		return nil, nil
	}
	result := make(map[string]float64)
	for i := 0; i < len(values); i = i + 2 {
		name := values[i]
		valueFloat, err := ConvertToFloat64(values[i+1])
		if err != nil {
			return nil, errorx.Errorf("convert %s to float64: %s", name, err)
		}
		result[name] = valueFloat
	}
	return result, nil
}

// BatchConvertToFloat64 批量转换字符串为float64
func BatchConvertToFloat64(values ...string) ([]float64, error) {
	vals, err := BatchConvertToFloat64Map(values...)
	if err != nil {
		return nil, err
	}
	result := make([]float64, 0, len(vals))
	for i := 0; i < len(values); i = i + 2 {
		result = append(result, vals[values[i]])
	}
	return result, nil
}

// CloneHttpTransport 克隆一个http.Transport
func CloneHttpTransport() *http.Transport {
	t, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		t = &http.Transport{}
	}
	return t.Clone()
}

// NewHttpClient 创建一个http.Client
func NewHttpClient(timeout time.Duration, proxyURL *url.URL) *http.Client {
	transport := CloneHttpTransport()
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

// SplitRange 将一个范围分成多个区间
func SplitRange(startVal int64, endVal int64, step int64) [][]int64 {
	var result [][]int64
	for startVal < endVal {
		end := min(startVal+step, endVal)
		result = append(result, []int64{startVal, end})
		startVal = end
	}
	return result
}

var BaseSymbols = []string{"USDT", "USDC", "USD", "BTC", "BNB", "ETH"}

func StandardizeCoin(coin string) string {
	return strings.ToUpper(coin)
}

// 标准格式化交易对
func StandardizeSymbol(symbol string) string {
	symbol = strings.ToUpper(symbol)
	if strings.Contains(symbol, "-") {
		return symbol
	}
	for _, baseSymbol := range BaseSymbols {
		if symbol == baseSymbol {
			return symbol
		}
		if strings.HasSuffix(symbol, baseSymbol) {
			return symbol[:len(symbol)-len(baseSymbol)] + "-" + baseSymbol
		}
	}
	return symbol
}

func FormatFloat64Percent(value float64, precision int) string {
	if value == 0 {
		return "0"
	}
	s := strconv.FormatFloat(value, 'f', precision, 64)
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		if s[len(s)-1] == '.' {
			s = s[:len(s)-1]
		}
	}
	parts := strings.Split(s, ".")
	size := len(parts[0])
	if size <= 3 {
		return s
	}
	var suffix string
	if len(parts) > 1 {
		suffix = "." + parts[1]
	}
	var result []byte
	start := size % 3
	if start == 0 {
		start = 3
	}
	src := []byte(parts[0])
	pos := 0
	for i := start; i < size; i += 3 {
		result = append(result, src[pos:i]...)
		result = append(result, ',')
		pos = i
	}
	if pos < size {
		result = append(result, src[pos:]...)
	}
	return string(result) + suffix
}

func FormatFloat64(value float64) string {
	return FormatFloat64Percent(value, -1)
}

func IsZeroStr(s string) bool {
	return s == "0" || s == "0.0" || s == "" || strings.TrimRight(s, "0") == "0."
}
