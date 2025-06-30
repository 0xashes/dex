package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	ComprehensivePriceEndpoint = "/api/v5/dex/index/current-price" // 多个第三方数据源的价格,计算出来的综合币价
	ComprehensivePriceMethod   = "POST"
)

type ComprehensivePriceParam struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
}
type ComprehensivePriceParams []ComprehensivePriceParam // 最多批量查询100各代币

type ComprehensivePriceResponse struct {
	Price                string `json:"price"`
	Time                 string `json:"time"`
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

func GetComprehensivePrice(client *common.Client, params *ComprehensivePriceParams) (*common.Response[[]ComprehensivePriceResponse], error) {
	response := &common.Response[[]ComprehensivePriceResponse]{}
	err := client.DoRequest(ComprehensivePriceMethod, ComprehensivePriceEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
