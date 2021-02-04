package proxy_market

import "time"

// ProxyMarketListRequest:
// Type: [ ipv4, ipv6, all ] Use proxy_market.ProxyType{}
// Page: minimum: 1
// PageSize: set to null to use your default value from the account. Set 0 to show all the records
// Sort:  0 - newest at the top; 1 - oldest at the top
type ProxyMarketListRequest struct {
	Type     string `json:"type"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Sort     int    `json:"sort"`
}

type ProxyMarketListResponse struct {
	Success bool   `json:"success"`
	Balance string `json:"balance"`
	List    struct {
		Error    bool `json:"error"`
		Interval struct {
			From int `json:"from"`
			To   int `json:"to"`
		} `json:"interval"`
		Total    int `json:"total"`
		PageSize int `json:"page_size"`
		Data     []*ProxyMarket `json:"data"`
	} `json:"list"`
}

type ProxyMarket struct {
	ID               int         `json:"id" csv:"id"`
	Login            string      `json:"login" csv:"login"`
	Password         string      `json:"password" csv:"password"`
	TariffID         string      `json:"tariff_id" csv:"tariff_id"`
	Active           int         `json:"active" csv:"active"`
	ExpiredAt        *time.Time  `json:"expired_at" csv:"expired_at"`
	CreatedAt        *time.Time  `json:"created_at" csv:"created_at"`
	UpdatedAt        *time.Time  `json:"updated_at" csv:"updated_at"`
	Comment          *string     `json:"comment" csv:"comment"`
	IP               string      `json:"ip" csv:"ip"`
	IPOut            string      `json:"ip_out" csv:"ip_out"`
	HTTPPort         int         `json:"http_port" csv:"http_port"`
	SocksPort        int         `json:"socks_port" csv:"socks_port"`
	AutoProlongation int         `json:"auto_prolongation" csv:"auto_prolongation"`
	Speed            int         `json:"speed" csv:"speed"`
}

// ProxyMarketBuyRequest
type ProxyMarketBuyRequest struct {
	PurchaseBilling *ProxyMarketBuyPurchaseBilling `json:"PurchaseBilling"`
}

// ProxyMarketBuyPurchaseBilling:
// Count: minimum: 1
// Type: use 100 for ipv4, 101 for ipv6, 102 for ipv4-shared, 105 for avito
// Duration: [7, 14, 30, 60, 90, 180, 365] [WARNING] 7 and 14 are available only for ipv6
// Country: [ru]
// Subnet: [32, 29] [WARNING] available only for ipv6
// Speed: [1, 2, 3] [WARNING] 1 - 1mb/s, 2 - 5mb/s, 3 - 15mb/s (only for ipv6)
type ProxyMarketBuyPurchaseBilling struct {
	Count     int    `json:"count"`
	Type      int    `json:"type"`
	Duration  int    `json:"duration"`
	Country   string `json:"country"`
	Promocode string `json:"promocode"`
	Subnet    int    `json:"subnet"`
	Speed     int    `json:"speed"`
}

type ProxyMarketBuyResponse struct {
	Success bool   `json:"success"`
	Balance int    `json:"balance"`
	Code    string `json:"code"`
}