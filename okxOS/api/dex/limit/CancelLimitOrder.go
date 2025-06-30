package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CancelLimitOrderEndpoint = "/api/v5/dex/aggregator/limit-order/cancel/calldata" // Get calldata of cancel limit order
	CancelLimitOrderMethod   = "GET"
)

type CancelLimitOrderParams struct {
	OrderHash string `json:"orderHash"` //	订单 hash
}

type CancelLimitOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func CancelLimitOrder(client *common.Client, params *CancelLimitOrderParams) (*CancelLimitOrderResponse, error) {
	response := &CancelLimitOrderResponse{}
	err := client.DoRequest(CancelLimitOrderMethod, CancelLimitOrderEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
