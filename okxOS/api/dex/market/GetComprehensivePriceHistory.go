package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	ComprehensivePriceHistoryEndpoint = "/api/v5/dex/index/historical-price" // 多个第三方数据源的价格,计算出来的历史综合币价
	ComprehensivePriceHistoryMethod   = "GET"
)

type ComprehensivePriceHistoryParams struct {
	ChainIndex           string  `json:"chainIndex"`
	TokenContractAddress *string `json:"tokenContractAddress,omitempty"`
	Limit                *string `json:"limit,omitempty"`
	Cursor               *string `json:"cursor,omitempty"`
	Begin                *string `json:"begin,omitempty"`
	End                  *string `json:"end,omitempty"`
	Period               *string `json:"period,omitempty"` // 1m/5m/30m/1h/1d
} // 只支持单个代币

type ComprehensivePriceHistoryResponse struct {
	Prices struct {
		Time       string `json:"time"`
		ChainIndex string `json:"chainIndex"`
	} `json:"prices"`
	Cursor string `json:"cursor"`
}

func GetComprehensivePriceHistory(client *common.Client, params *ComprehensivePriceHistoryParams) (*common.Response[[]ComprehensivePriceHistoryResponse], error) {
	response := &common.Response[[]ComprehensivePriceHistoryResponse]{}
	err := client.DoRequest(ComprehensivePriceHistoryMethod, ComprehensivePriceHistoryEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
