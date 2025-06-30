package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	SwapEndpoint = "/api/v5/dex/aggregator/swap" // Generate the data for calling the OKX DEX router to swap
	SwapMethod   = "GET"
)

type SwapParams struct {
	ChainIndex                      string  `json:"chainIndex"`
	Amount                          string  `json:"amount"`
	SwapMode                        string  `json:"swapMode"`
	FromTokenAddress                string  `json:"fromTokenAddress"`
	ToTokenAddress                  string  `json:"toTokenAddress"`
	Slippage                        string  `json:"slippage"`
	UserWalletAddress               string  `json:"userWalletAddress"`
	SwapReceiverAddress             *string `json:"swapReceiverAddress,omitempty"`
	FeePercent                      *string `json:"feePercent,omitempty"`
	FromTokenReferrerWalletAddress  *string `json:"fromTokenReferrerWalletAddress,omitempty"`
	ToTokenReferrerWalletAddress    *string `json:"toTokenReferrerWalletAddress,omitempty"`
	PositiveSlippagePercent         *string `json:"positiveSlippagePercent,omitempty"`
	PositiveSlippageFeeAddress      *string `json:"positiveSlippageFeeAddress,omitempty"`
	Gaslimit                        *string `json:"gaslimit,omitempty"`
	GasLevel                        *string `json:"gasLevel,omitempty"`
	DexIds                          *string `json:"dexIds,omitempty"`
	DirectRoute                     *bool   `json:"directRoute,omitempty"`
	PriceImpactProtectionPercentage *string `json:"priceImpactProtectionPercentage,omitempty"`
	CallDataMemo                    *string `json:"callDataMemo,omitempty"`
	ComputeUnitPrice                *string `json:"computeUnitPrice,omitempty"`
	ComputeUnitLimit                *string `json:"computeUnitLimit,omitempty"`
	AutoSlippage                    *bool   `json:"autoSlippage,omitempty"`
	MaxAutoSlippage                 *string `json:"maxAutoSlippage,omitempty"`
}

type SwapResponse struct {
	RouterResult struct {
		ChainID       string `json:"chainId"`
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
	} `json:"routerResult"`
	Tx struct {
		Data                 string   `json:"data"`
		From                 string   `json:"from"`
		Gas                  string   `json:"gas"`
		GasPrice             string   `json:"gasPrice"`
		MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
		MaxSpendAmount       string   `json:"maxSpendAmount"`
		MinReceiveAmount     string   `json:"minReceiveAmount"`
		SignatureData        []string `json:"signatureData"`
		To                   string   `json:"to"`
		Value                string   `json:"value"`
	} `json:"tx"`
}

// 执行请求
func Swap(client *common.Client, params *SwapParams) (*common.Response[[]SwapResponse], error) {
	response := &common.Response[[]SwapResponse]{}
	err := client.DoRequest(SwapMethod, SwapEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
