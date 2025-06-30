package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	SolanaSwapInstructionsEndpoint = "/api/v5/dex/aggregator/swap-instruction" // 获取 Solana 兑换交易指令
	SolanaSwapInstructionsMethod   = "GET"
)

type SolanaSwapInstructionsParams struct {
	ChainIndex                      string  `json:"chainIndex"`
	Amount                          string  `json:"amount"`
	FromTokenAddress                string  `json:"fromTokenAddress"`
	ToTokenAddress                  string  `json:"toTokenAddress"`
	Slippage                        string  `json:"slippage"`
	AutoSlippage                    *bool   `json:"autoSlippage,omitempty"`
	MaxAutoSlippage                 *string `json:"maxAutoSlippage,omitempty"`
	UserWalletAddress               string  `json:"userWalletAddress"`
	SwapReceiverAddress             *string `json:"swapReceiverAddress,omitempty"`
	FeePercent                      *string `json:"feePercent,omitempty"`
	FromTokenReferrerWalletAddress  *string `json:"fromTokenReferrerWalletAddress,omitempty"`
	ToTokenReferrerWalletAddress    *string `json:"toTokenReferrerWalletAddress,omitempty"`
	PositiveSlippagePercent         *string `json:"positiveSlippagePercent,omitempty"`
	PositiveSlippageFeeAddress      *string `json:"positiveSlippageFeeAddress,omitempty"`
	DexIds                          *string `json:"dexIds,omitempty"`
	DirectRoute                     *bool   `json:"directRoute,omitempty"`
	PriceImpactProtectionPercentage *string `json:"priceImpactProtectionPercentage,omitempty"`
	ComputeUnitPrice                *string `json:"computeUnitPrice,omitempty"`
	ComputeUnitLimit                *string `json:"computeUnitLimit,omitempty"`
}

type SolanaSwapInstructionsResponse struct {
	AddressLookupTableAccount []string `json:"addressLookupTableAccount"`
	InstructionLists          []struct {
		Data     string `json:"data"`
		Accounts []struct {
			IsSigner   bool   `json:"isSigner"`
			IsWritable bool   `json:"isWritable"`
			Pubkey     string `json:"pubkey"`
		} `json:"accounts"`
		ProgramID string `json:"programId"`
	} `json:"instructionLists"`
	RouterResult struct {
		ChainID       string `json:"chainId"`
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
		From             string `json:"from"`
		MinReceiveAmount string `json:"minReceiveAmount"`
		Slippage         string `json:"slippage"`
		To               string `json:"to"`
	} `json:"tx"`
}

// 执行请求
func GetSolanaSwapInstructions(client *common.Client, params *SolanaSwapInstructionsParams) (*common.Response[SwapResponse], error) {
	response := &common.Response[SwapResponse]{}
	err := client.DoRequest(SolanaSwapInstructionsMethod, SolanaSwapInstructionsEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
