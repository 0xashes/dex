package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossSupportedBridgesEndpoint = "/api/v5/dex/cross-chain/supported/bridges" // Get information of the cross-chain bridges supported by OKX’s DEX cross-chain aggregator protocole
	CrossSupportedBridgesMethod   = "GET"
)

type CrossSupportedBridgesParams struct {
	ChainId *string `json:"chainId,omitempty"`
}

type CrossSupportedBridgesResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetCrossSupportedBridges(client *common.Client, params *CrossSupportedBridgesParams) (*CrossSupportedBridgesResponse, error) {
	response := &CrossSupportedBridgesResponse{}
	err := client.DoRequest(CrossSupportedBridgesMethod, CrossSupportedBridgesEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
