package proxy_market

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alok87/goutils/pkg/random"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"net/http"
)

func (c *ProxyMarketClient) doReq(method, path string, data []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ApiUrl, path, c.ApiKey)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProxyMarketClient) GetProxyListByCustom(reqBody *ProxyMarketListRequest) (*ProxyMarketListResponse, error){
	path := "list"
	breq, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.doReq("POST", path, breq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *ProxyMarketListResponse
	err = json.Unmarshal(br, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ProxyMarketClient) GetProxyListAllByNewest() (*ProxyMarketListResponse, error)  {
	path := "list"
	reqBody := ProxyMarketListRequest{
		Type:     "all",
		Page:     1,
		PageSize: 100000,
		Sort:     0,
	}
	breq, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.doReq("POST", path, breq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *ProxyMarketListResponse
	err = json.Unmarshal(br, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ProxyMarketClient) GetProxyListAllByOldest() (*ProxyMarketListResponse, error)  {
	path := "list"
	reqBody := ProxyMarketListRequest{
		Type:     "all",
		Page:     1,
		PageSize: 100000,
		Sort:     1,
	}
	breq, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.doReq("POST", path, breq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *ProxyMarketListResponse
	err = json.Unmarshal(br, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ProxyMarketClient) BuyProxyMarketCustom(reqBody *ProxyMarketBuyRequest) (*ProxyMarketBuyResponse, error) {
	path := "buy-proxy"
	breq, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.doReq("POST", path, breq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *ProxyMarketBuyResponse
	err = json.Unmarshal(br, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// BuyProxyMarketTypeIPV4Shared:
// Покупка proxies IPV4 Shared
func (c *ProxyMarketClient) BuyProxyMarketTypeIPV4Shared(count int, duration int) (*ProxyMarketBuyResponse, error) {
	path := "buy-proxy"
	reqBody := ProxyMarketBuyRequest{PurchaseBilling: &ProxyMarketBuyPurchaseBilling{
		Count:     count,
		Type:      102,
		Duration:  duration,
		Country:   "ru",
		Promocode: "",
		Subnet:    0,
		Speed:     3,
	}}
	breq, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.doReq("POST", path, breq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *ProxyMarketBuyResponse
	err = json.Unmarshal(br, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetRandomProxyFromAllList:
// If refresh key == true, first action - update proxy slice, two action - random
// if refresh key == false, getting fast proxy from proxies slice
func (c *ProxyMarketClient) GetRandomProxyFromAllList(refresh bool) (*ProxyMarket, error)  {
	if len(c.Proxies) == 0 || refresh == true {
		prxResp, err := c.GetProxyListAllByNewest()
		if err != nil {
			return nil, err
		}
		c.Proxies = prxResp.List.Data
	}
	rInt := random.RangeInt(0, len(c.Proxies) - 1, 1)
	return c.Proxies[rInt[0]], nil
}

func (c *ProxyMarketClient) GetProxyCSVFile() (string, error)  {
	resp, err := c.GetProxyListAllByNewest()
	if err != nil {
		return "", err
	}
	var proxies []*ProxyMarket
	for _, v := range resp.List.Data {
		proxies = append(proxies, v)
	}
	csvContent, err := gocsv.MarshalString(&proxies)
	if err != nil {
		return "", err
	}
	return csvContent, nil
}