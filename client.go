package proxy_market

import (
	"net/http"
	"time"
)

const (
	BaseProxyMarketUrl = "https://proxy.market/dev-api"
	ProxyTypeIPV4	   = "ipv4"
	ProxyTypeIPV6	   = "ipv6"
	ProxyTypeAll 	   = "all"
)

type ProxyMarketClient struct {
	ApiUrl		string
	ApiKey  	string
	Client  	*http.Client
	Proxies     []*ProxyMarket
}


func NewProxyMarketClient(apiKey string) *ProxyMarketClient  {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	pmc := ProxyMarketClient{
		ApiUrl: BaseProxyMarketUrl,
		ApiKey: apiKey,
		Client: &c,
	}
	return &pmc
}
