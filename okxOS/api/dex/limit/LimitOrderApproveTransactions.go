package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LimitOrderApproveTransactionsEndpoint = "/api/v5/dex/aggregator/approve-transaction" // Generate the relevant data for calling the contract
	LimitOrderApproveTransactionsMethod   = "GET"
)

type LimitOrderApproveTransactionsParams struct {
	ChainId              string `json:"chainId"`              // 链 ID （如1: Ethereum，更多可查看链 ID 列表）
	TokenContractAddress string `json:"tokenContractAddress"` //币种合约地址 （如0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48）
	ApproveAmount        string `json:"approveAmount"`        // 执行授权的币种数量（数量需包含精度，如授权 1.00 USDT 需输入 1000000，授权 1.00 DAI 需输入 1000000000000000000）
}

type LimitOrderApproveTransactionsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func LimitOrderApproveTransactions(client *common.Client, params *LimitOrderApproveTransactionsParams) (*LimitOrderApproveTransactionsResponse, error) {
	response := &LimitOrderApproveTransactionsResponse{}
	err := client.DoRequest(LimitOrderApproveTransactionsMethod, LimitOrderApproveTransactionsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
