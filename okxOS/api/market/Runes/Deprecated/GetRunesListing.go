package runes

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	RunesListingEndpoint = "/api/v5/mktplace/nft/runes/get-runes-order-list"
	RunesListingMethod   = "GET"
)

type RunesListingParams struct {
	RunesId string  `json:"runesId"`          // runes 代币的唯一标识，必传，字符串类型
	Cursor  *string `json:"cursor,omitempty"` // 指向要检索的排序序号的游标，最大 1000，可选，字符串类型
	Limit   *int    `json:"limit,omitempty"`  // 分页大小，默认 10，最大 100，可选，整数类型
	SortBy  *string `json:"sortBy,omitempty"` // 订单排序规则，默认 "unitPriceAsc"，枚举值：unitPriceAsc、unitPriceDesc、totalPriceAsc、totalPriceDesc、listedTimeAsc、listedTimeDesc，可选，字符串类型
}

type RunesListingResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetRunesListing(client *common.Client, params *RunesListingParams) (*RunesListingResponse, error) {
	response := &RunesListingResponse{}
	err := client.DoRequest(RunesListingMethod, RunesListingEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
