package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossBestRouteInformationEndpoint = "/api/v5/dex/cross-chain/quote" // Find the best route for a cross-chain swap through OKX’s DEX cross-chain aggregator
	CrossBestRouteInformationMethod   = "GET"
)

type CrossBestRouteInformationParams struct {
	FromChainId                     string  `json:"fromChainId"`
	ToChainId                       string  `json:"toChainId"`
	FromTokenAddress                string  `json:"fromTokenAddress"`
	ToTokenAddress                  string  `json:"toTokenAddress"`
	Amount                          string  `json:"amount"`
	Slippage                        string  `json:"slippage"`
	Sort                            *int    `json:"sort,omitempty"`
	DexIds                          *string `json:"dexIds,omitempty"`
	FeePercent                      *string `json:"feePercent,omitempty"`
	AllowBridge                     *[]int  `json:"allowBridge,omitempty"`
	DenyBridge                      *[]int  `json:"denyBridge,omitempty"`
	PriceImpactProtectionPercentage *string `json:"priceImpactProtectionPercentage,omitempty"`
}

type CrossBestRouteInformationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetCrossBestRouteInformation(client *common.Client, params *CrossBestRouteInformationParams) (*CrossBestRouteInformationResponse, error) {
	response := &CrossBestRouteInformationResponse{}
	err := client.DoRequest(CrossBestRouteInformationMethod, CrossBestRouteInformationEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
