package orinals

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	QueryOrdinalsListingEndpoint = "/api/v5/mktplace/nft/ordinals/listings"
	QueryOrdinalsListingMethod   = "GET"
)

type QueryOrdinalsListingParams struct {
	Slug    *string `json:"slug,omitempty"`    // 合集 slug 名称，即合集的唯一标识，可选，字符串类型
	Cursor  *string `json:"cursor,omitempty"`  // 指向要检索的页面的游标，可选，字符串类型
	Limit   *string `json:"limit,omitempty"`   // 分页大小，默认值 "10"，最大 "100"，返回最大的订单数，可选，字符串类型
	Sort    *string `json:"sort,omitempty"`    // 挂单排序规则：listing_time_desc, listing_time_asc, price_desc, price_asc, unit_price_desc, unit_price_asc，默认 listing_time_desc，可选，字符串类型
	IsBrc20 *bool   `json:"isBrc20,omitempty"` // 获取全部 BTC NFT 或 Brc20 合集的订单，默认为 Brc20，可选，布尔类型
}

type QueryOrdinalsListingResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func QueryOrdinalsListing(client *common.Client, params *QueryOrdinalsListingParams) (*QueryOrdinalsListingResponse, error) {
	response := &QueryOrdinalsListingResponse{}
	err := client.DoRequest(QueryOrdinalsListingMethod, QueryOrdinalsListingEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
