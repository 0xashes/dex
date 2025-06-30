package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	SupportedChainsEndpoint = "/api/v5/dex/market/supported/chain"
	SupportedChainsMethod   = "GET"
)

type SupportedChainsParams struct {
	ChainIndex *string `json:"chainIndex,omitempty"`
}

type SupportedChainsResponse struct {
	ChainIndex  string `json:"chainIndex"`
	ChainName   string `json:"chainName"`
	ChainSymbol string `json:"chainSymbol"`
}

func GetSupportedChains(client *common.Client, params *SupportedChainsParams) (*common.Response[[]SupportedChainsResponse], error) {
	response := &common.Response[[]SupportedChainsResponse]{}
	err := client.DoRequest(SupportedChainsMethod, SupportedChainsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
