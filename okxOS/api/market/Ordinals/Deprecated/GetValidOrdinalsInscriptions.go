package orinals

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	ValidOrdinalsInscriptionsEndpoint = "/api/v5/mktplace/nft/ordinals/get-valid-inscriptions"
	ValidOrdinalsInscriptionsMethod   = "GET"
)

type ValidOrdinalsInscriptionsParams struct {
	Slug          string  `json:"slug"`              // 合集 slug 名称，即合集的唯一标识，必传，字符串类型
	Cursor        *string `json:"cursor,omitempty"`  // 指向要检索的页面的游标，可选，字符串类型
	Limit         *string `json:"limit,omitempty"`   // 分页大小，默认 "10"，最大 "100"，可选，字符串类型
	Sort          *string `json:"sort,omitempty"`    // 挂单排序规则，默认 "listing_time_desc"，可选，字符串类型
	IsBrc20       *bool   `json:"isBrc20,omitempty"` // 是否获取 BRC-20 合集，默认 true，可选，布尔类型
	WalletAddress string  `json:"walletAddress"`     // 获取铭文的钱包地址，必传，字符串类型
}

type ValidOrdinalsInscriptionsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetValidOrdinalsInscriptions(client *common.Client, params *ValidOrdinalsInscriptionsParams) (*ValidOrdinalsInscriptionsResponse, error) {
	response := &ValidOrdinalsInscriptionsResponse{}
	err := client.DoRequest(ValidOrdinalsInscriptionsMethod, ValidOrdinalsInscriptionsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
