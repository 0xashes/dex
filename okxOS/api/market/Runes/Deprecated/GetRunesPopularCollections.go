package runes

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	RunesPopularCollectionsEndpoint = "/api/v5/mktplace/nft/runes/get-hot-collection"
	RunesPopularCollectionsMethod   = "GET"
)

type RunesPopularCollectionsParams struct {
	TimeType *int    `json:"timeType,omitempty"` // 返回 x 个项目的数据，枚举值：1 (24h), 2 (7d), 3 (30d), 4 (all)，默认 24h，可选，整数类型
	Cursor   *string `json:"cursor,omitempty"`   // 查询列表分页游标，默认空，可选，字符串类型
	Limit    *int    `json:"limit,omitempty"`    // 查询列表页大小，默认 10，可选，整数类型
}

type RunesPopularCollectionsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetRunesPopularCollections(client *common.Client, params *RunesPopularCollectionsParams) (*RunesPopularCollectionsResponse, error) {
	response := &RunesPopularCollectionsResponse{}
	err := client.DoRequest(RunesPopularCollectionsMethod, RunesPopularCollectionsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
