package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LastPriceInfoEndpoint = "/api/v5/dex/market/price-info"
	LastPriceInfoMethod   = "POST"
)

type LastPriceInfoParam struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

type LastPriceInfoParams []LastPriceInfoParam

type LastPriceInfoResponse struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
	Time                 string `json:"time"`
	Price                string `json:"price"`
	MarketCap            string `json:"marketCap"`
	PriceChange5M        string `json:"priceChange5M"`
	PriceChange1H        string `json:"priceChange1H"`
	PriceChange4H        string `json:"priceChange4H"`
	PriceChange24H       string `json:"priceChange24H"`
	Volume5M             string `json:"volume5M"`
	Volume1H             string `json:"volume1H"`
	Volume4H             string `json:"volume4H"`
	Volume24H            string `json:"volume24H"`
}

func GetLastPriceInfo(client *common.Client, params *LastPriceInfoParams) (*common.Response[[]LastPriceInfoResponse], error) {
	response := &common.Response[[]LastPriceInfoResponse]{}
	err := client.DoRequest(LastPriceInfoMethod, LastPriceInfoEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
