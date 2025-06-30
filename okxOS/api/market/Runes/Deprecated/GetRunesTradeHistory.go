package runes

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	RunesTradeHistoryEndpoint = "/api/v5/mktplace/nft/runes/trade-history"
	RunesTradeHistoryMethod   = "GET"
)

type RunesTradeHistoryParams struct {
	RunesIds  *string `json:"runesIds,omitempty"`  // 代币唯一标识，多个以逗号分隔，不超过20个，空串查询所有代币，可选，字符串类型
	Cursor    *string `json:"cursor,omitempty"`    // 指向要检索的排序序号的游标，可选，字符串类型
	Limit     *int    `json:"limit,omitempty"`     // 分页大小，默认 10，最大 100，可选，整数类型
	StartTime *int64  `json:"startTime,omitempty"` // 交易动态的起始时间（单位：秒），可选，长整型
	EndTime   *int64  `json:"endTime,omitempty"`   // 交易动态的截止时间（单位：秒），可选，长整型
}

type RunesTradeHistoryResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetRunesTradeHistory(client *common.Client, params *RunesTradeHistoryParams) (*RunesTradeHistoryResponse, error) {
	response := &RunesTradeHistoryResponse{}
	err := client.DoRequest(RunesTradeHistoryMethod, RunesTradeHistoryEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
