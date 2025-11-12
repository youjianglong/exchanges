package okx

import (
	"fmt"

	"github.com/youjianglong/exchanges/common"
)

// APIError 定义 OKX API 错误，当响应状态码为 4xx 或 5xx 时返回
type APIError struct {
	Code     int64  `json:"code,string"`
	Message  string `json:"msg"`
	Response []byte `json:"-"`
}

func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> rsp=%s", string(e.Response))
}

func (e APIError) IsValid() bool {
	return e.Code != 0 || e.Message != ""
}

// IsAPIError 检查 error 是否为 APIError 类型
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}

// AsProxyError convert APIError to ProxyError
func AsProxyError(e error) (error, bool) {
	if e == nil {
		return nil, false
	}
	ae, ok := e.(*APIError)
	if !ok {
		return e, false
	}
	if ae.Code == 50110 {
		return fmt.Errorf("%w %w", common.ErrProxy, ae), true
	}
	return ae, false
}

// ConvertProxyError convert error to ProxyError
func ConvertProxyError(e error) error {
	if e == nil {
		return nil
	}
	ae, ok := AsProxyError(e)
	if !ok {
		return e
	}
	return ae
}
