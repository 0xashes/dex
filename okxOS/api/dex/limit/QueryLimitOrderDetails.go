package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LimitOrderDetailsEndpoint = "/api/v5/dex/aggregator/limit-order/detail" // Query detail of DEX limit order
	LimitOrderDetailsMethod   = "GET"
)

type LimitOrderDetailsParams struct {
	OrderHash string `json:"orderHash"` //	订单 hash
}

type LimitOrderDetailsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func QueryLimitOrderDetails(client *common.Client, params *LimitOrderDetailsParams) (*LimitOrderDetailsResponse, error) {
	response := &LimitOrderDetailsResponse{}
	err := client.DoRequest(LimitOrderDetailsMethod, LimitOrderDetailsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
