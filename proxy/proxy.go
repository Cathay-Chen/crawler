package proxy

import (
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

// ProxyFunc 代理函数
type ProxyFunc func(r *http.Request) (*url.URL, error)

// roundRobinSwitcher 轮询代理
type roundRobinSwitcher struct {
	proxyURLs []*url.URL
	index     uint32
}

// GetProxy 返回一个轮询代理的函数
func (r *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&r.index, 1) - 1
	u := r.proxyURLs[index%uint32(len(r.proxyURLs))]
	return u, nil
}

// RoundRobinProxySwitcher 返回一个轮询代理的函数
func RoundRobinProxySwitcher(ProxyURLs ...string) (ProxyFunc, error) {
	if len(ProxyURLs) < 1 {
		return nil, errors.New("Proxy URL list is empty.")
	}
	urls := make([]*url.URL, len(ProxyURLs))
	for i, u := range ProxyURLs {
		parseU, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		urls[i] = parseU
	}
	return (&roundRobinSwitcher{urls, 0}).GetProxy, nil
}
