package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

// 参数老错误
const (
	DeFiRedaptionInfoEndpoint = "/api/v5/defi/calculator/redeem-info"
	DeFiRedaptionInfoMethod   = "POST"
)

type DeFiRedaptionInfoParams struct {
	Address            *string `json:"address,omitempty"`
	InputTokenAmount   string  `json:"inputTokenAmount"`
	InvestmentCategory *string `json:"investmentCategory,omitempty"` // 认购类型：(0：默认类型；1：BRC-20)
	InvestmentId       string  `json:"investmentId"`
	IsSingle           *bool   `json:"isSingle,omitempty"` // 判断是否为单币投资品：是：单币；否：多币
	OutputTokenAddress string  `json:"outputTokenAddress"` // 赎回代币的智能合约地址
	Slippage           *string `json:"slippage,omitempty"` // 滑点：默认为 1%
}

type DeFiRedaptionInfoResponse struct {
	IsAllowRedeem       bool               `json:"isAllowRedeem"`
	EstimateGasFee      string             `json:"estimateGasFee"`
	IsSwapInvest        bool               `json:"isSwapInvest"`
	InputCurrencyAmount string             `json:"inputCurrencyAmount"`
	ReceiveTokenInfo    InvestTokenModel   `json:"receiveTokenInfo"`
	SwapFromTokenList   []InvestTokenModel `json:"swapFromTokenList"`
	MySupply            InvestTokenModel   `json:"mySupply"`
	ApproveStatusList   []ApproveInfoModel `json:"approveStatusList"`
}

func GetDeFiRedaptionInfo(client *common.Client, params *DeFiRedaptionInfoParams) (*common.OKXOSResponse[DeFiRedaptionInfoResponse], error) {
	response := &common.OKXOSResponse[DeFiRedaptionInfoResponse]{}
	err := client.DoRequest(DeFiRedaptionInfoMethod, DeFiRedaptionInfoEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
