package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LastPriceEndpoint = "/api/v5/dex/market/price"
	LastPriceMethod   = "POST"
)

type LastPriceParam struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

type LastPriceParams []LastPriceParam

type LastPriceResponse struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
	Time                 string `json:"time"`
	Price                string `json:"price"`
}

func GetLastPrice(client *common.Client, params *LastPriceParams) (*common.Response[[]LastPriceResponse], error) {
	response := &common.Response[[]LastPriceResponse]{}
	err := client.DoRequest(LastPriceMethod, LastPriceEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
