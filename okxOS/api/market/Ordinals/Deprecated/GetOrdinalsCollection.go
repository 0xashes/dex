package orinals

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	GetOrdinalsCollectionEndpoint = "/api/v5/mktplace/nft/ordinals/collections"
	GetOrdinalsCollectionMethod   = "GET"
)

type GetOrdinalsCollectionParams struct {
	Slug    *string `json:"slug,omitempty"`    // 合集 slug 名称，即合集的唯一标识，参数不为空则代表查询指定合集，可选，字符串类型
	Cursor  *string `json:"cursor,omitempty"`  // 指向要检索的页面的游标，可选，字符串类型
	Limit   *string `json:"limit,omitempty"`   // 分页大小，默认值 "100"，最大 "300"，返回最大的合集数，可选，字符串类型
	IsBrc20 *bool   `json:"isBrc20,omitempty"` // 获取全部 BTC NFT 或 Brc20 合集的列表，默认为 Brc20，可选，布尔类型
}

type GetOrdinalsCollectionResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetOrdinalsCollection(client *common.Client, params *GetOrdinalsCollectionParams) (*GetOrdinalsCollectionResponse, error) {
	response := &GetOrdinalsCollectionResponse{}
	err := client.DoRequest(GetOrdinalsCollectionMethod, GetOrdinalsCollectionEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
