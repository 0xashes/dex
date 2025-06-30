package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	LiquiditySourcesEndpoint = "/api/v5/dex/aggregator/get-liquidity" // Get liquidity sources available for swap in the OKX aggregation protocol
	LiquiditySourcesMethod   = "GET"
)

type LiquiditySourcesParams struct {
	ChainIndex string `json:"chainIndex"`
}

type LiquiditySourcesResponse struct {
	ID   string `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

// 执行请求
func GetLiquiditySources(client *common.Client, params *LiquiditySourcesParams) (*common.Response[[]LiquiditySourcesResponse], error) {
	response := &common.Response[[]LiquiditySourcesResponse]{}
	err := client.DoRequest(LiquiditySourcesMethod, LiquiditySourcesEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
