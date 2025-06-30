package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LimitOrderSupportedChainsEndpoint = "/api/v5/dex/aggregator/limit-order/chain" // Obtain information about chain IDs
	LimitOrderSupportedChainsMethod   = "GET"
)

type LimitOrderSupportedChainsParams struct {
	ChainId *int `json:"chainId,omitempty"` // 链 ID （如1: Ethereum，更多可查看链 ID 列表）
}

type LimitOrderSupportedChainsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetLimitOrderSupportedChains(client *common.Client, params *LimitOrderSupportedChainsParams) (*LimitOrderSupportedChainsResponse, error) {
	response := &LimitOrderSupportedChainsResponse{}
	err := client.DoRequest(LimitOrderSupportedChainsMethod, LimitOrderSupportedChainsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
