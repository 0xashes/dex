package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossApproveTransactionsEndpoint = "/api/v5/dex/aggregator/approve-transaction" // Generate the relevant data for calling the contract
	CrossApproveTransactionsMethod   = "GET"
)

type CrossApproveTransactionsParams struct {
	FromChainId                     string  `json:"fromChainId"`
	ToChainId                       string  `json:"toChainId"`
	FromTokenAddress                string  `json:"fromTokenAddress"`
	ToTokenAddress                  string  `json:"toTokenAddress"`
	Amount                          string  `json:"amount"`
	Slippage                        string  `json:"slippage"`
	Sort                            *string `json:"sort,omitempty"`
	DexIds                          *string `json:"dexIds,omitempty"`
	UserWalletAddress               string  `json:"userWalletAddress"`
	AllowBridge                     *[]int  `json:"allowBridge,omitempty"`
	DenyBridge                      *[]int  `json:"denyBridge,omitempty"`
	ReceiveAddress                  *string `json:"receiveAddress,omitempty"`
	FeePercent                      *string `json:"feePercent,omitempty"`
	ReferrerAddress                 *string `json:"referrerAddress,omitempty"`
	PriceImpactProtectionPercentage *string `json:"priceImpactProtectionPercentage,omitempty"`
	OnlyBridge                      *string `json:"onlyBridge,omitempty"`
	Memo                            *string `json:"memo,omitempty"`
}

type CrossApproveTransactionsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func CrossApproveTransactions(client *common.Client, params *CrossApproveTransactionsParams) (*CrossApproveTransactionsResponse, error) {
	response := &CrossApproveTransactionsResponse{}
	err := client.DoRequest(CrossApproveTransactionsMethod, CrossApproveTransactionsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
