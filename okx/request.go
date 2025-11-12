package okx

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type secType int

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned
)

type params map[string]interface{}

// request 定义一个 API 请求
type request struct {
	method   string
	endpoint string
	query    url.Values
	data     map[string]any
	secType  secType
	header   http.Header
	body     io.Reader
	fullURL  string
}

func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setParams(m params) *request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

func (r *request) setData(key string, value any) *request {
	if r.data == nil {
		r.data = make(map[string]any)
	}
	r.data[key] = value
	return r
}

func (r *request) setDatas(m map[string]any) *request {
	for k, v := range m {
		r.setData(k, v)
	}
	return r
}

func (r *request) validate() error {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.data == nil {
		r.data = make(map[string]any)
	}
	return nil
}

type RequestOption func(*request)

func WithHeader(key, value string, replace bool) RequestOption {
	return func(r *request) {
		if r.header == nil {
			r.header = http.Header{}
		}
		if replace {
			r.header.Set(key, value)
		} else {
			r.header.Add(key, value)
		}
	}
}

func WithHeaders(header http.Header) RequestOption {
	return func(r *request) {
		r.header = header.Clone()
	}
}
