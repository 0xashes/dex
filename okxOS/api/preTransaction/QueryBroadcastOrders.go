package preTransaction

import (
	"github.com/0xashes/dex/okxOS/common"
)

// 广播模拟API仅向企业客户提供
const (
	BroadcastOrdersEndpoint = "/api/v5/dex/post-transaction/orders" // 通过交易信息的预执行，获取预估消耗的 Gaslimit 。
	BroadcastOrdersMethod   = "GET"
)

type BroadcastOrdersParams struct {
	Address    string  `json:"address"`
	ChainIndex string  `json:"chainIndex"`
	TxStatus   *string `json:"txStatus,omitempty"`
	OrderId    *string `json:"orderId,omitempty"`
	Cursor     *string `json:"cursor,omitempty"`
	Limit      *string `json:"limit,omitempty"`
}

type BroadcastOrdersResponse struct {
	Cursor string `json:"cursor"`
	Orders []struct {
		ChainIndex string `json:"chainIndex"`
		Address    string `json:"address"`
		OrderId    string `json:"orderId"`
		TxStatus   string `json:"txStatus"`
		FailReason string `json:"failReason"`
		TxHash     string `json:"txHash"`
	} `json:"orders"`
}

// 执行请求
func QueryBroadcastOrders(client *common.Client, params *BroadcastOrdersParams) (*common.Response[[]BroadcastOrdersResponse], error) {
	response := &common.Response[[]BroadcastOrdersResponse]{}
	err := client.DoRequest(BroadcastOrdersMethod, BroadcastOrdersEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
