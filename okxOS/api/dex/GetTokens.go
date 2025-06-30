package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	TokensEndpoint = "/api/v5/dex/aggregator/all-tokens" // 获取币种列表，也可以指定列表之外的代币在OKDEX交易
	TokensMethod   = "GET"
)

type TokensParams struct {
	ChainIndex string `json:"chainIndex"`
}

type TokensResponse struct {
	Decimals             string `json:"decimals"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenLogoUrl         string `json:"tokenLogoUrl"`
	TokenName            string `json:"tokenName"`
	TokenSymbol          string `json:"tokenSymbol"`
}

// 执行请求
func GetTokens(client *common.Client, params *TokensParams) (*common.Response[[]TokensResponse], error) {
	response := &common.Response[[]TokensResponse]{}
	err := client.DoRequest(TokensMethod, TokensEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
