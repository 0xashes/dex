package runes

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	RunesCollectionEndpoint = "/api/v5/mktplace/nft/runes/detail"
	RunesCollectionMethod   = "GET"
)

type RunesCollectionParams struct {
	RunesId string `json:"runesId"` // runesId，支持批量查询，使用‘,’分隔，例如 "840000:3,840000:28"，必传，字符串类型
}

type RunesCollectionResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetRunesCollection(client *common.Client, params *RunesCollectionParams) (*RunesCollectionResponse, error) {
	response := &RunesCollectionResponse{}
	err := client.DoRequest(RunesCollectionMethod, RunesCollectionEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
