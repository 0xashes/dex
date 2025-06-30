package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiBonusClamingCallDataEndpoint = "/api/v5/defi/transaction/bonus"
	DeFiBonusClamingCallDataMethod   = "POST"
)

type DeFiBonusClamingCallDataParams struct {
	Address          string              `json:"address"`          // 用户钱包地址
	InvestmentId     string              `json:"investmentId"`     // 投资品 ID
	UserInputList    []UserInputParam    `json:"userInputList"`    // 用户输入代币信息
	ExpectOutputList []ExpectOutputParam `json:"expectOutputLIst"` // 用户预期收益信息
	Extra            *string             `json:"extra,omitempty"`  // 具体含义见备注
}

type DeFiBonusClamingCallDataResponse DeFiAuthorizationCallDataResponse

func GetDeFiBonusClamingCallData(client *common.Client, params *DeFiBonusClamingCallDataParams) (*common.OKXOSResponse[DeFiBonusClamingCallDataResponse], error) {
	response := &common.OKXOSResponse[DeFiBonusClamingCallDataResponse]{}
	err := client.DoRequest(DeFiBonusClamingCallDataMethod, DeFiBonusClamingCallDataEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
} // 生成领取奖励金交易的调用数据
