package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LimitOrderListEndpoint = "/api/v5/dex/aggregator/limit-order/all" // Query list of DEX limit order
	LimitOrderListMethod   = "GET"
)

type LimitOrderListParams struct {
	ChainId    string  `json:"chainId"`              // 链ID（如 "1" 表示 Ethereum），必传，字符串类型
	Page       *string `json:"page,omitempty"`       // 页码数，默认 "1"，可选，字符串类型
	Limit      *string `json:"limit,omitempty"`      // 每页条数，默认 "100"，最大 "500"，可选，字符串类型
	Statuses   *string `json:"statuses,omitempty"`   // 状态数组，用于筛选订单状态，可选，字符串类型
	TakerAsset *string `json:"takerAsset,omitempty"` // 成交钱包地址，可选，字符串类型
	MakerAsset *string `json:"makerAsset,omitempty"` // 用户下单钱包地址，可选，字符串类型
}

type LimitOrderListResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func QueryLimitOrderList(client *common.Client, params *LimitOrderListParams) (*LimitOrderListResponse, error) {
	response := &LimitOrderListResponse{}
	err := client.DoRequest(LimitOrderListMethod, LimitOrderListEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
