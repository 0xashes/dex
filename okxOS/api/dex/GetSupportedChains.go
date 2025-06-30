package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	SupportedChainsEndpoint = "/api/v5/dex/aggregator/supported/chain"
	SupportedChainsMethod   = "GET"
)

type SupportedChainsParams struct {
	ChainIndex *string `json:"chainIndex,omitempty"`
}

type SupportedChainsResponse struct {
	ChainId                int    `json:"chainId"`
	ChainIndex             int    `json:"chainIndex"`
	ChainName              string `json:"chainName"`
	DexTokenApproveAddress string `json:"dexTokenApproveAddress"`
}

type ChainInfo struct{}

func GetSupportedChains(client *common.Client, params *SupportedChainsParams) (*common.Response[[]SupportedChainsResponse], error) {
	response := &common.Response[[]SupportedChainsResponse]{}
	err := client.DoRequest(SupportedChainsMethod, SupportedChainsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
