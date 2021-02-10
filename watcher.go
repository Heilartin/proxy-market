package proxy_market

import "time"

// StartWatcher: updatedTime in milliseconds
func (c *ProxyMarketClient) StartWatcher(updatedTime time.Duration)  {
	go func() {
		for {
			res, err := c.GetProxyListAllByNewest()
			if err != nil {
				time.Sleep(updatedTime * time.Millisecond)
				continue
			}
			c.Proxies = res.List.Data
			time.Sleep(updatedTime * time.Millisecond)
		}
	}()
}
