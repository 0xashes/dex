package account

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	TransactionsHistoryEndpoint = "/api/v5/dex/post-transaction/transactions-by-address"
	TransactionsHistoryMethod   = "GET"
)

type TransactionsHistoryParams struct {
	Address              string  `json:"address"`
	Chains               string  `json:"chains"`
	TokenContractAddress *string `json:"tokenContractAddress,omitempty"`
	Begin                *string `json:"begin,omitempty"`
	End                  *string `json:"end,omitempty"`
	Cursor               *string `json:"cursor,omitempty"`
	Limit                *string `json:"limit,omitempty"`
}

type TransactionsHistoryResponse struct {
	Transactions []struct {
		ChainIndex string `json:"chainIndex"`
		TxHash     string `json:"txHash"`
		Itype      string `json:"itype"`    // 交易层级类型
		MethodId   string `json:"methodId"` // 合约调用函数
		Nonce      string `json:"nonce"`    // 发起者地址发起的第几笔交易
		TxTime     string `json:"txTime"`   // 交易时间，UNIX时间戳毫秒级
		From       []struct {
			Address string `json:"address"`
			Amount  string `json:"amount"`
		} `json:"from"`
		To []struct {
			Address string `json:"address"`
			Amount  string `json:"amount"`
		} `json:"to"`
		TokenContractAddress string `json:"tokenContractAddress"`
		Amount               string `json:"amount"`
		Symbol               string `json:"symbol"`
		TxFee                string `json:"txFee"`
		TxStatus             string `json:"txStatus"`
		HitBlacklist         bool   `json:"hitBlacklist"`
	} `json:"transactions"`
	Cursor string `json:"cursor"`
}

func GetTransactionsHistory(client *common.Client, params *TransactionsHistoryParams) (*common.Response[[]TransactionsHistoryResponse], error) {
	response := &common.Response[[]TransactionsHistoryResponse]{}
	err := client.DoRequest(TransactionsHistoryMethod, TransactionsHistoryEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
