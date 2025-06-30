package preTransaction

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	GasPriceEndpoint = "/api/v5/dex/pre-transaction/gas-price" // 获取交易上链 API 支持的链信息。
	GasPriceMethod   = "GET"
)

type GasPriceParams struct {
	ChainIndex string `json:"chainIndex"`
}

type GasPriceResponse struct {
	Normal         string `json:"normal"`
	Min            string `json:"min"`
	Max            string `json:"max"`
	Supporteip1559 bool   `json:"supporteip1559"`

	Eip1559Protocol struct {
		BaseFee            string `json:"baseFee"`
		SuggestBaseFee     string `json:"suggestBaseFee"`
		ProposePriorityFee string `json:"proposePriorityFee"`
		SafePriorityFee    string `json:"safePriorityFee"`
		FastPriorityFee    string `json:"fastPriorityFee"`
	} `json:"eip1559Protocol"`

	PriorityFee struct { // 优先费，只适用于Solana
		ProposePriorityFee string `json:"proposePriorityFee"`
		SafePriorityFee    string `json:"safePriorityFee"`
		FastPriorityFee    string `json:"fastPriorityFee"`
		ExtremePriorityFee string `json:"extremePriorityFee"`
	} `json:"priorityFee"`
}

// 执行请求
func GetGasPrice(client *common.Client, params *GasPriceParams) (*common.Response[[]GasPriceResponse], error) {
	response := &common.Response[[]GasPriceResponse]{}
	err := client.DoRequest(GasPriceMethod, GasPriceEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
