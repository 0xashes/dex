package preTransaction

import (
	"math/big"

	"github.com/0xashes/dex/okxOS/common"
)

// 交易模拟API仅向企业客户提供
const (
	SimulateTransactionEndpoint = "/api/v5/dex/pre-transaction/simulate" // 通过交易信息的预执行，获取预估消耗的 Gaslimit 。
	SimulateTransactionMethod   = "POST"
)

type SimulateTransactionParams struct {
	ChainIndex  string   `json:"chainIndex"`
	FromAddress string   `json:"fromAddress"`
	ToAddress   string   `json:"toAddress"`
	TxAmount    *string  `json:"txAmount,omitempty"` // 交易金额(可从swapResponse.tx.value获取)
	ExtJson     *ExtJson `json:"extJson"`
	PriorityFee *string  `json:"priorityFee,omitempty"` // 优先费（Solana）
	GasPrice    *string  `json:"gasPrice,omitempty"`
}

type SimulateTransactionResponse struct {
	Intention   string `json:"intention"`
	AssetChange []struct {
		AssetType string  `json:"assetType"`
		Name      string  `json:"name"`
		Symbol    string  `json:"symbol"`
		Decimals  big.Int `json:"decimals"`
		Address   string  `json:"address"`
		ImageUrl  string  `json:"imageUrl"`
		RawValue  string  `json:"rawValue"`
	} `json:"assetChange"`
	GasUsed    string `json:"gasUsed"`
	FailReason string `json:"failReason"`
	Risks      []struct {
		Address     string `json:"address"`
		AddressType string `json:"addressType"`
	}
}

// 执行请求
func SimulateTransaction(client *common.Client, params *SimulateTransactionParams) (*common.Response[[]SimulateTransactionResponse], error) {
	response := &common.Response[[]SimulateTransactionResponse]{}
	err := client.DoRequest(SimulateTransactionMethod, SimulateTransactionEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
