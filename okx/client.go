package okx

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/youjianglong/exchanges/common"
)

const (
	apiUrl           = "https://www.okx.com"
	wsPublicMainUrl  = "wss://ws.okx.com:8443/ws/v5/public"
	wsPrivateMainUrl = "wss://ws.okx.com:8443/ws/v5/private"
	wsPublicTestUrl  = "wss://wspap.okx.com:8443/ws/v5/public"
	wsPrivateTestUrl = "wss://wspap.okx.com:8443/ws/v5/private"
	PingTimeout      = 20 * time.Second
	PingDeadline     = 10 * time.Second
)

var (
	ErrNoData   = errors.New("no data")
	PingMessage = []byte("ping")
)

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	APIKey     string
	SecretKey  string
	Passphrase string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	Test       bool
	Logger     *slog.Logger
	TimeOffset int64
	do         doFunc
}

func NewClient(apiKey, secretKey, passphrase string) *Client {
	return NewClientWithHttpClient(apiKey, secretKey, passphrase, false, common.NewHttpClient(16*time.Second, nil))
}

func NewClientWithHttpClient(apiKey, secretKey, passphrase string, test bool, httpClient *http.Client) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		BaseURL:    apiUrl,
		HTTPClient: httpClient,
		Logger:     slog.With("E", "okx"),
		Test:       test,
	}
}

func NewTestClient(apiKey, secretKey, passphrase string) *Client {
	return NewClientWithHttpClient(apiKey, secretKey, passphrase, true, common.NewHttpClient(16*time.Second, nil))
}

func (c *Client) WithHttpClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = common.NewHttpClient(16*time.Second, nil)
	}
	nc := *c
	nc.HTTPClient = httpClient
	return &nc
}

func (c *Client) SetProxyURL(proxyURL *url.URL) {
	transport, ok := c.HTTPClient.Transport.(*http.Transport)
	if !ok {
		transport = common.CloneHttpTransport()
		c.HTTPClient.Transport = transport
	}
	if proxyURL == nil {
		transport.Proxy = nil
	} else {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
}

func (c *Client) debug(format string, v ...any) {
	if c.Debug {
		c.Logger.Error(fmt.Sprintf(format, v...))
	}
}

func currentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) error {
	for _, opt := range opts {
		opt(r)
	}
	err := r.validate()
	if err != nil {
		return err
	}
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	// OKX 签名规则：prehash = timestamp + method + requestPath + (body)
	timestamp := currentTimestamp()
	var bodyStr string
	if len(r.data) > 0 {
		b, _ := json.Marshal(r.data)
		bodyStr = string(b)
	}
	if len(r.query) > 0 {
		r.endpoint = fmt.Sprintf("%s?%s", r.endpoint, r.query.Encode())
		fullURL = fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	}
	prehash := fmt.Sprintf("%s%s%s%s", timestamp, r.method, r.endpoint, bodyStr)
	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	_, err = mac.Write([]byte(prehash))
	if err != nil {
		return err
	}
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyStr != "" {
		header.Set("Content-Type", "application/json")
	}
	header.Set("OK-ACCESS-KEY", c.APIKey)
	header.Set("OK-ACCESS-SIGN", signature)
	header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	header.Set("OK-ACCESS-PASSPHRASE", c.Passphrase)
	if c.Test {
		header.Set("x-simulated-trading", "1")
	}
	r.fullURL = fullURL
	r.header = header
	c.debug("full url: %s", r.fullURL)
	c.debug("request headers: %v", r.header)
	if bodyStr != "" {
		r.body = bytes.NewBufferString(bodyStr)
		c.debug("request body: %s", bodyStr)
	}
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (json.RawMessage, error) {
	err := c.parseRequest(r, opts...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)
		if err := json.Unmarshal(data, apiErr); err != nil {
			c.debug("failed to unmarshal error: %v", err)
		}
		return nil, apiErr
	}

	var common struct {
		Code  string          `json:"code"`
		Msg   string          `json:"msg"`
		Data  json.RawMessage `json:"data"`
		SCode string          `json:"sCode"`
		SMsg  string          `json:"sMsg"`
	}
	if err := json.Unmarshal(data, &common); err != nil {
		return nil, err
	}
	if common.Code != "0" {
		var errMsg string
		if common.SCode != "" {
			errMsg = common.SCode + ": " + common.SMsg
		} else {
			errMsg = common.Msg
		}
		return nil, fmt.Errorf("API error: %s", errMsg)
	}
	// 返回 data 部分的原始 JSON 数据
	return common.Data, nil
}
