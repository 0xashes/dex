package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossDirectTokensPairsEndpoint = "/api/v5/dex/cross-chain/supported/bridge-tokens-pairs" // Get the list of tokens pairs supported by chains
	CrossDirectTokensPairsMethod   = "GET"
)

type CrossDirectTokensPairsParams struct {
	FromChainId string `json:"fromChainId"`
}

type CrossDirectTokensPairsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func CrossDirectTokensPairsTokens(client *common.Client, params *CrossDirectTokensPairsParams) (*CrossDirectTokensPairsResponse, error) {
	response := &CrossDirectTokensPairsResponse{}
	err := client.DoRequest(CrossDirectTokensPairsMethod, CrossDirectTokensPairsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
