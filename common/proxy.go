package common

import (
	"errors"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"

	"github.com/youjianglong/exchanges/errorx"
)

type ProxyURLSetter interface {
	SetProxyURL(proxyUrl *url.URL)
}

type ProxyPool struct {
	mutex     sync.Mutex
	proxies   []*ProxyNode
	logger    *slog.Logger
	current   *ProxyNode
	lastCheck time.Time
	setter    ProxyURLSetter
}

type ProxyNode struct {
	URL         *url.URL
	FailCount   int
	LastChecked time.Time
}

var (
	ErrNoProxyAvailable = errors.New("no proxy available") // 没有可用的代理
	ErrProxy            = errors.New("<ProxyError>")       // 代理不可用
)

// 创建代理池
func NewProxyPool(proxyURLs []string) (*ProxyPool, error) {
	pool := &ProxyPool{}
	pool.proxies = make([]*ProxyNode, 0, len(proxyURLs))
	for _, proxyUrl := range proxyURLs {
		if proxyUrl == "" {
			continue
		}
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, errorx.Errorf("invalid proxy url %s, err: %v", proxyUrl, err)
		}
		node := &ProxyNode{
			URL: proxy,
		}
		pool.proxies = append(pool.proxies, node)
	}
	if len(pool.proxies) > 0 {
		pool.current = pool.proxies[0]
	}
	pool.logger = slog.Default()

	return pool, nil
}

func (p *ProxyPool) WithSetter(setter ProxyURLSetter) *ProxyPool {
	p.setter = setter
	return p
}

func (p *ProxyPool) WithLogger(logger *slog.Logger) *ProxyPool {
	p.logger = logger
	return p
}

func (p *ProxyPool) setProxyURL(proxyUrl *url.URL) {
	if p.setter != nil {
		p.setter.SetProxyURL(proxyUrl)
	}
}

func (p *ProxyPool) Check(ping func() error, checkInterval time.Duration) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.lastCheck.Add(checkInterval).Before(time.Now()) {
		return
	}
	p.lastCheck = time.Now()
	for i := 0; i < len(p.proxies); i++ {
		err := ping()
		if err == nil {
			return
		}
		proxyUrl, err := p.switchProxy()
		if err != nil {
			if errors.Is(err, ErrNoProxyAvailable) {
				p.setProxyURL(nil)
			} else {
				slog.Warn("switch proxy error: " + err.Error())
			}
			return
		}
		p.setProxyURL(proxyUrl)
	}
}

func (p *ProxyPool) GetCurrentProxy() *ProxyNode {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.current
}

func (p *ProxyPool) GetCurrentProxyURL() *url.URL {
	node := p.GetCurrentProxy()
	if node == nil {
		return nil
	}
	return node.URL
}

func (p *ProxyPool) Proxy(*http.Request) (*url.URL, error) {
	return p.GetCurrentProxyURL(), nil
}

// SwitchProxy 切换代理
func (p *ProxyPool) switchProxy() (*url.URL, error) {
	size := len(p.proxies)
	if size == 0 {
		return nil, ErrNoProxyAvailable
	}
	if p.current == nil {
		p.current = p.proxies[0]
	} else {
		p.current.FailCount++
		p.current.LastChecked = time.Now()
	}

	if size == 1 {
		return p.current.URL, nil
	}
	if size == 2 {
		for _, proxy := range p.proxies {
			if proxy == p.current {
				continue
			}
			p.current = proxy
			return proxy.URL, nil
		}
	}

	var pending []*ProxyNode
	for _, proxy := range p.proxies {
		if proxy == p.current {
			continue
		}
		pending = append(pending, proxy)
	}
	sort.Slice(pending, func(i, j int) bool {
		if pending[i].FailCount == pending[j].FailCount {
			return pending[i].LastChecked.Before(pending[j].LastChecked)
		}
		return pending[i].FailCount < pending[j].FailCount
	})
	size = len(pending)
	for i := 0; i < size; i = i + 2 {
		end := min(i+2, size)
		results := p.pingProxies(pending[i:end])
		var available []*PingResult
		for _, result := range results {
			if result.err == nil {
				available = append(available, result)
			} else {
				result.proxy.FailCount++
				result.proxy.LastChecked = time.Now()
			}
		}
		if len(available) == 0 {
			continue
		}
		if len(available) == 1 {
			p.current = available[0].proxy
			return p.current.URL, nil
		}
		sort.Slice(available, func(i, j int) bool {
			return available[i].elapsed < available[j].elapsed
		})
		p.current = available[0].proxy
		return p.current.URL, nil
	}
	slog.Warn("no proxy available", slog.String("current", p.current.URL.String()))
	return p.current.URL, nil
}

func (p *ProxyPool) SwitchProxy() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	proxyUrl, err := p.switchProxy()
	if err != nil {
		return err
	}
	p.setProxyURL(proxyUrl)
	return nil
}

type PingResult struct {
	proxy   *ProxyNode
	elapsed time.Duration
	err     error
}

func (p *ProxyPool) pingProxies(nodes []*ProxyNode) []*PingResult {
	results := make([]*PingResult, len(nodes))
	wg := &sync.WaitGroup{}
	ping := func(i int) {
		defer wg.Done()
		node := nodes[i]
		elapsed, err := TcpPing(node.URL)
		if err != nil {
			node.FailCount++
			node.LastChecked = time.Now()
		}
		results[i] = &PingResult{
			proxy:   node,
			elapsed: elapsed,
			err:     err,
		}
	}
	wg.Add(len(nodes))
	for i := range nodes {
		go ping(i)
	}
	wg.Wait()
	return results
}

func TcpPing(proxyUrl *url.URL) (time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", proxyUrl.Host, 5*time.Second)
	if err != nil {
		return 0, err
	}
	_ = conn.Close()

	return time.Since(start), nil
}

func NewHttpClientWithProxy(proxyUrl *url.URL) (*http.Client, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	return client, nil
}

func CheckProxy(proxyUrl *url.URL, request *http.Request) (time.Duration, error) {
	client, err := NewHttpClientWithProxy(proxyUrl)
	if err != nil {
		return 0, err
	}
	client.Timeout = 10 * time.Second
	start := time.Now()
	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	_ = resp.Body.Close()
	elapsed := time.Since(start)
	return elapsed, nil
}
