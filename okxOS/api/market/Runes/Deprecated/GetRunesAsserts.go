package runes

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	RunesAssertsEndpoint = "/api/v5/mktplace/nft/runes/get-owned-asserts"
	RunesAssertsMethod   = "GET"
)

type RunesAssertsParams struct {
	RunesId         string  `json:"runesId"`          // 代币的唯一标识，必传，字符串类型
	WalletAddresses string  `json:"walletAddresses"`  // 钱包地址，多个地址以逗号分隔，最大 20 个，必传，字符串类型
	Cursor          *string `json:"cursor,omitempty"` // 指向要检索的排序序号的游标，最大 1000，可选，字符串类型
	Limit           *int    `json:"limit,omitempty"`  // 分页大小，默认 10，最大 100，可选，整数类型
}

type RunesAssertsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetRunesAsserts(client *common.Client, params *RunesAssertsParams) (*RunesAssertsResponse, error) {
	response := &RunesAssertsResponse{}
	err := client.DoRequest(RunesAssertsMethod, RunesAssertsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
