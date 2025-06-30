package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossSupportedChainsEndpoint = "/api/v5/dex/cross-chain/supported/chain" // Get information from chains that support cross-chain transactions
	CrossSupportedChainsMethod   = "GET"
)

type CrossSupportedChainsParams struct {
	ChainId *string `json:"chainId,omitempty"`
}

type CrossSupportedChainsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetCrossSupportedChains(client *common.Client, params *CrossSupportedChainsParams) (*CrossSupportedChainsResponse, error) {
	response := &CrossSupportedChainsResponse{}
	err := client.DoRequest(CrossSupportedChainsMethod, CrossSupportedChainsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
