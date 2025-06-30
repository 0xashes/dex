package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossDirectTokensEndpoint = "/api/v5/dex/cross-chain/supported/tokens" // Get token list only trade for bridge
	CrossDirectTokensMethod   = "GET"
)

type CrossDirectTokensParams struct {
	ChainId *string `json:"chainId,omitempty"`
}

type CrossDirectTokensResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetCrossDirectTokens(client *common.Client, params *CrossDirectTokensParams) (*CrossDirectTokensResponse, error) {
	response := &CrossDirectTokensResponse{}
	err := client.DoRequest(CrossDirectTokensMethod, CrossDirectTokensEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
