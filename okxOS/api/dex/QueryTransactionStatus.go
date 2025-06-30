package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	TransactionStatusEndpoint = "/api/v5/dex/aggregator/history"
	TransactionStatusMethod   = "GET"
)

type TransactionStatusParams struct {
	ChainIndex      string `json:"chainIndex"`
	TxHash          string `json:"txHash"`
	IsFromMyProject *bool  `json:"isFromMyProject,omitempty"`
}

type TransactionStatusResponse struct {
	ChainId          string `json:"chainId"`
	DexRouter        string `json:"dexRouter"`
	ErrorMsg         string `json:"errorMsg"`
	FromAddress      string `json:"fromAddress"`
	FromTokenDetails struct {
		Amount       string `json:"amount"`
		Symbol       string `json:"symbol"`
		TokenAddress string `json:"tokenAddress"`
	} `json:"fromTokenDetails"`
	GasLimit       string `json:"gasLimit"`
	GasPrice       string `json:"gasPrice"`
	GasUsed        string `json:"gasUsed"`
	Height         string `json:"height"`
	ReferralAmount string `json:"referralAmount"`
	Status         string `json:"status"`
	ToAddress      string `json:"toAddress"`
	ToTokenDetails struct {
		Amount       string `json:"amount"`
		Symbol       string `json:"symbol"`
		TokenAddress string `json:"tokenAddress"`
	} `json:"toTokenDetails"`
	TxFee  string `json:"txFee"`
	TxHash string `json:"txHash"`
	TxTime string `json:"txTime"`
	TxType string `json:"txType"`
}

// 执行请求
func QueryTransactionStatus(client *common.Client, params *TransactionStatusParams) (*TransactionStatusResponse, error) {
	response := &TransactionStatusResponse{}
	err := client.DoRequest(TransactionStatusMethod, TransactionStatusEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
