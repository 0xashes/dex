package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossSwapEndpoint = "/api/v5/dex/cross-chain/status" // Check the final status of the cross-chain swap according to txhash
	CrossSwapMethod   = "GET"
)

type CrossSwapParams struct {
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

type CrossSwapResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	// Data []ChainInfo `json:"data"`
}

// 执行请求
func CrossSwap(client *common.Client, params *CrossSwapParams) (*CrossSwapResponse, error) {
	response := &CrossSwapResponse{}
	err := client.DoRequest(CrossSwapMethod, CrossSwapEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
