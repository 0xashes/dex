package preTransaction

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	GasLimitEndpoint = "/api/v5/dex/pre-transaction/gas-limit" // 通过交易信息的预执行，获取预估消耗的 Gaslimit 。
	GasLimitMethod   = "POST"
)

type GasLimitParams struct {
	ChainIndex  string   `json:"chainIndex"`
	FromAddress string   `json:"fromAddress"`
	ToAddress   string   `json:"toAddress"`
	TxAmount    *string  `json:"txAmount,omitempty"` // 交易金额(可从swapResponse.tx.value获取)
	ExtJson     *ExtJson `json:"extJson,omitempty"`
}

type ExtJson struct {
	InputData string `json:"inputData"`
}

type GasLimitResponse struct {
	GasLimit string `json:"gasLimit"`
}

// 执行请求
func GetGasLimit(client *common.Client, params *GasLimitParams) (*common.Response[[]GasLimitResponse], error) {
	response := &common.Response[[]GasLimitResponse]{}
	err := client.DoRequest(GasLimitMethod, GasLimitEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
