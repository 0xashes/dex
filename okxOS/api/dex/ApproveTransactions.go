package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	ApproveTransactionsEndpoint = "/api/v5/dex/aggregator/approve-transaction" // 交易授权
	ApproveTransactionsMethod   = "GET"
)

type ApproveTransactionsParams struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
	ApproveAmount        string `json:"approveAmount"`
}

type ApproveTransactionsResponse struct {
	Data               string `json:"data"`
	DexContractAddress string `json:"dexContractAddress"`
	GasLimit           string `json:"gasLimit"`
	GasPrice           string `json:"gasPrice"`
}

// 执行请求
func ApproveTransactions(client *common.Client, params *ApproveTransactionsParams) (*common.Response[[]ApproveTransactionsResponse], error) {
	response := &common.Response[[]ApproveTransactionsResponse]{}
	err := client.DoRequest(ApproveTransactionsMethod, ApproveTransactionsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
