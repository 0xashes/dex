package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiRedaptionCallDataEndpoint = "/api/v5/defi/transaction/redemption"
	DeFiRedaptionCallDataMethod   = "Post"
)

type DeFiRedaptionCallDataParams struct {
	Address          string              `json:"address"`          // 用户钱包地址
	InvestmentId     string              `json:"investmentId"`     // 投资品 ID
	UserInputList    []UserInputParam    `json:"userInputList"`    // 用户输入代币信息
	ExpectOutputList []ExpectOutputParam `json:"expectOutputLIst"` // 用户预期收益信息
	Extra            *string             `json:"extra,omitempty"`  // 具体含义见备注
}

type DeFiRedaptionCallDataResponse DeFiAuthorizationCallDataResponse

func GetDeFiRedaptionCallData(client *common.Client, params *DeFiRedaptionCallDataParams) (*common.OKXOSResponse[DeFiRedaptionCallDataResponse], error) {
	response := &common.OKXOSResponse[DeFiRedaptionCallDataResponse]{}
	err := client.DoRequest(DeFiRedaptionCallDataMethod, DeFiRedaptionCallDataEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
} // 生成赎回交易的调用数据
