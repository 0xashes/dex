package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiSubcriptionCallDataEndpoint = "/api/v5/defi/transaction/subscription"
	DeFiSubcriptionCallDataMethod   = "POST"
)

type DeFiSubcriptionCallDataParams struct {
	Address          string              `json:"address"`
	InvestmentId     string              `json:"investmentId"`
	UserInputList    []UserInputParam    `json:"userInputList"`    // 用户输入代币信息
	ExpectOutputList []ExpectOutputParam `json:"expectOutputLIst"` // 用户预期收益信息
	Extra            string              `json:"extra,omitempty"`
}

type UserInputParam struct {
	ChainId      *string `json:"chainId,omitempty"`      // 区块链 ID
	CoinAmount   string  `json:"coinAmount"`             // 认购数量
	TokenAddress *string `json:"tokenAddress,omitempty"` // 认购代币的智能合约地址
}

type ExpectOutputParam struct {
	ChainId      *string `json:"chainId,omitempty"`      // 区块链 ID
	CoinAmount   string  `json:"coinAmount"`             // 认购数量
	TokenAddress *string `json:"tokenAddress,omitempty"` // 认购代币的智能合约地址
}

type DeFiSubcriptionCallDataResponse DeFiAuthorizationCallDataResponse

func GetDeFiSubcriptionCallData(client *common.Client, params *DeFiSubcriptionCallDataParams) (*common.OKXOSResponse[DeFiSubcriptionCallDataResponse], error) {
	response := &common.OKXOSResponse[DeFiSubcriptionCallDataResponse]{}
	err := client.DoRequest(DeFiSubcriptionCallDataMethod, DeFiSubcriptionCallDataEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
} // 生成申购交易的调用数据
