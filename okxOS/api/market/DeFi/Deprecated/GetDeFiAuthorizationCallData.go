package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiAuthorizationCallDataEndpoint = "/api/v5/defi/transaction/authorization"
	DeFiAuthorizationCallDataMethod   = "POST"
)

type DeFiAuthorizationCallDataParams struct {
	Address          string              `json:"address"`                    // 用户钱包地址
	InvestmentId     string              `json:"investmentId"`               // 投资品 ID
	Type             int                 `json:"type"`                       // 交易类型 (例如，3：认购授权；4：赎回授权；5：领取授权)
	UserInputList    []UserInputParam    `json:"userInputList"`              // 用户输入代币信息
	ExpectOutputList []ExpectOutputParam `json:"expectOutputLIst,omitempty"` // 用户预期收益信息
}

type DeFiAuthorizationCallDataResponse struct {
	DataList []struct {
		From           string      `json:"from"`
		To             string      `json:"to"`
		Value          string      `json:"value"`
		SerializedData string      `json:"serializedData"`
		OriginalData   interface{} `json:"originalData"`
		CallDataType   string      `json:"callDataType"`
		SignatureData  string      `json:"signatureData"`
	}
}

func GetDeFiAuthorizationCallData(client *common.Client, params *DeFiAuthorizationCallDataParams) (*common.OKXOSResponse[DeFiAuthorizationCallDataResponse], error) {
	response := &common.OKXOSResponse[DeFiAuthorizationCallDataResponse]{}
	err := client.DoRequest(DeFiAuthorizationCallDataMethod, DeFiAuthorizationCallDataEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
} // 生成交易前授权的调用数据
