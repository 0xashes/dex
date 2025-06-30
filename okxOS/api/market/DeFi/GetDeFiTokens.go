package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiTokensEndpoint = "/api/v5/defi/explore/token/list"
	DeFiTokensMethod   = "GET"
)

type DeFiTokensParams struct {
	TokenAddress *string `json:"tokenAddress,omitempty"`
	ChainId      *string `json:"chainId,omitempty"`
}

type DeFiTokensResponse struct {
	Symbol     string      `json:"symbol"`
	TokenInfos []TokenInfo `json:"tokenInfos"`
}

type TokenInfo struct {
	TokenID      string `json:"tokenId"`
	TokenSymbol  string `json:"tokenSymbol"`
	Network      string `json:"network"`
	LogoURL      string `json:"logoUrl"`
	TokenAddress string `json:"tokenAddress"`
}

func GetDeFiTokens(client *common.Client, params *DeFiTokensParams) (*common.OKXOSResponse[[]DeFiTokensResponse], error) {
	response := &common.OKXOSResponse[[]DeFiTokensResponse]{}
	err := client.DoRequest(DeFiTokensMethod, DeFiTokensEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
