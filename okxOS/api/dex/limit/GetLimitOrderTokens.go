package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LimitOrderTokensEndpoint = "/api/v5/dex/aggregator/all-tokens" // 获取币种列表，也可以指定列表之外的代币在OKDEX交易
	LimitOrderTokensMethod   = "GET"
)

type LimitOrderTokensParams struct {
	ChainId string `json:"chainId"` // 链 ID (如1: Ethereum，更多可查看链 ID 列表)
}

type LimitOrderTokensResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetLimitOrderTokens(client *common.Client, params *LimitOrderTokensParams) (*LimitOrderTokensResponse, error) {
	response := &LimitOrderTokensResponse{}
	err := client.DoRequest(LimitOrderTokensMethod, LimitOrderTokensEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
