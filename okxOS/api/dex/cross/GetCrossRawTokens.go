package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossRawTokensEndpoint = "/api/v5/dex/aggregator/all-tokens" // 获取币种列表，也可以指定列表之外的代币在OKDEX交易
	CrossRawTokensMethod   = "GET"
)

type CrossRawTokensParams struct {
	ChainId string `json:"chainId"`
}

type CrossRawTokensResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetCrossRawTokens(client *common.Client, params *CrossRawTokensParams) (*CrossRawTokensResponse, error) {
	response := &CrossRawTokensResponse{}
	err := client.DoRequest(CrossRawTokensMethod, CrossRawTokensEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
