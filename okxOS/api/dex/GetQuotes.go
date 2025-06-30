package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	QuotesEndpoint = "/api/v5/dex/aggregator/quote" // 获取兑换价格
	QuotesMethod   = "GET"
)

type QuotesParams struct {
	ChainIndex                      string  `json:"chainIndex"`
	Amount                          string  `json:"amount"`
	SwapMode                        string  `json:"swapMode"`
	FromTokenAddress                string  `json:"fromTokenAddress"`
	ToTokenAddress                  string  `json:"toTokenAddress"`
	DexIds                          *string `json:"dexIds,omitempty"`
	DirectRoute                     *bool   `json:"directRoute,omitempty"`
	PriceImpactProtectionPercentage *string `json:"priceImpactProtectionPercentage,omitempty"`
	FeePercent                      *string `json:"feePercent,omitempty"`
}

type QuoteResponse struct {
	ChainIndex    string `json:"chainIndex"`
	DexRouterList []struct {
		Router        string `json:"router"`
		RouterPercent string `json:"routerPercent"`

		SubRouterList []struct {
			DexProtocol []struct {
				DexName string `json:"dexName"`
				Percent string `json:"percent"`
			} `json:"dexProtocol"`
			FromToken struct {
				Decimal              string `json:"decimal"`
				IsHoneyPot           bool   `json:"isHoneyPot"`
				TaxRate              string `json:"taxRate"`
				TokenContractAddress string `json:"tokenContractAddress"`
				TokenSymbol          string `json:"tokenSymbol"`
				TokenUnitPrice       string `json:"tokenUnitPrice"`
			} `json:"fromToken"`
			ToToken struct {
				Decimal              string `json:"decimal"`
				IsHoneyPot           bool   `json:"isHoneyPot"`
				TaxRate              string `json:"taxRate"`
				TokenContractAddress string `json:"tokenContractAddress"`
				TokenSymbol          string `json:"tokenSymbol"`
				TokenUnitPrice       string `json:"tokenUnitPrice"`
			} `json:"toToken"`
		} `json:"subRouterList"`
	} `json:"dexRouterList"`
	EstimateGasFee string `json:"estimateGasFee"`
	FromToken      struct {
		Decimal              string `json:"decimal"`
		IsHoneyPot           bool   `json:"isHoneyPot"`
		TaxRate              string `json:"taxRate"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenUnitPrice       string `json:"tokenUnitPrice"`
	} `json:"fromToken"`
	FromTokenAmount       string `json:"fromTokenAmount"`
	OriginToTokenAmount   string `json:"originToTokenAmount"`
	PriceImpactPercentage string `json:"priceImpactPercentage"`
	QuoteCompareList      []struct {
		AmountOut string `json:"amountOut"`
		DexLogo   string `json:"dexLogo"`
		DexName   string `json:"dexName"`
		TradeFee  string `json:"tradeFee"`
	} `json:"quoteCompareList"`
	SwapMode string `json:"swapMode"`
	ToToken  struct {
		Decimal              string `json:"decimal"`
		IsHoneyPot           bool   `json:"isHoneyPot"`
		TaxRate              string `json:"taxRate"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenUnitPrice       string `json:"tokenUnitPrice"`
	} `json:"toToken"`
	ToTokenAmount string `json:"toTokenAmount"`
	TradeFee      string `json:"tradeFee"`
}

// 执行请求
func GetQuotes(client *common.Client, params *QuotesParams) (*common.Response[[]QuoteResponse], error) {
	response := &common.Response[[]QuoteResponse]{}
	err := client.DoRequest(QuotesMethod, QuotesEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
