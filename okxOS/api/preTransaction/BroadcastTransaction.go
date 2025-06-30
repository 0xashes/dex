package preTransaction

import (
	"github.com/0xashes/dex/okxOS/common"
)

// 广播模拟API仅向企业客户提供
const (
	BroadcastTransactionEndpoint = "/api/v5/dex/pre-transaction/broadcast-transaction" // 通过交易信息的预执行，获取预估消耗的 Gaslimit 。
	BroadcastTransactionMethod   = "POST"
)

type BroadcastTransactionParams struct {
	SignedTx   string `json:"signedTx"`
	ChainIndex string `json:"chainIndex"`
	Address    string `json:"address"`
}

type BroadcastTransactionResponse struct {
	OrderId string `json:"orderId"`
}

// 执行请求
func BroadcastTransaction(client *common.Client, params *BroadcastTransactionParams) (*common.Response[[]BroadcastTransactionResponse], error) {
	response := &common.Response[[]BroadcastTransactionResponse]{}
	err := client.DoRequest(BroadcastTransactionMethod, BroadcastTransactionMethod, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
