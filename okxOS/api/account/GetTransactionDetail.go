package account

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	TransactionDetailEndpoint = "/api/v5/dex/post-transaction/transaction-detail-by-txhash"
	TransactionDetailMethod   = "GET"
)

type TransactionDetailParams struct {
	ChainIndex string  `json:"chainIndex"`
	TxHash     string  `json:"txHash"`
	Itype      *string `json:"itype,omitempty"`
}

type TransactionDetailResponse struct {
	ChainIndex  string `json:"chainIndex"`
	Height      string `json:"height"`
	TxTime      string `json:"txTime"`
	TxHash      string `json:"txHash"`
	TxStatus    string `json:"txStatus"`
	GasLimit    string `json:"gasLimit"`
	GasUsed     string `json:"gasUsed"`
	GasPrice    string `json:"gasPrice"`
	TxFee       string `json:"txFee"`
	Nonce       string `json:"nonce"`
	Amount      string `json:"amount"`
	Symbol      string `json:"symbol"`
	MethodId    string `json:"methodId"`
	FromDetails []struct {
		Address      string `json:"address"`
		VinIndex     string `json:"vinIndex"`
		PreVoutIndex string `json:"preVoutIndex"`
		TxHash       string `json:"txHash"`
		IsContract   bool   `json:"isContract"`
		Amount       string `json:"amount"`
	} `json:"fromDetails"`
	ToDetails []struct {
		Address    string `json:"address"`
		VoutIndex  string `json:"voutIndex"`
		IsContract bool   `json:"isContract"`
		Amount     string `json:"amount"`
	} `json:"toDetails"`
	InternalTransactionDetails []struct {
		From           string `json:"from"`
		To             string `json:"to"`
		IsFromContract bool   `json:"isFromContract"`
		IsToContract   bool   `json:"isToContract"`
		Amount         string `json:"amount"`
		TxStatus       string `json:"txStatus"`
	} `json:"internalTransactionDetails"`
	TokenTransferDetails []struct {
		From                 string `json:"from"`
		To                   string `json:"to"`
		IsFromContract       bool   `json:"isFromContract"`
		IsToContract         bool   `json:"isToContract"`
		TokenContractAddress string `json:"tokenContractAddress"`
		Symbol               string `json:"symbol"`
		Amount               string `json:"amount"`
	} `json:"tokenTransferDetails"`
	L1OriginHash string `json:"l1OriginHash"`
}

func GetTransactionDetail(client *common.Client, params *TransactionDetailParams) (*common.Response[[]TransactionDetailResponse], error) {
	response := &common.Response[[]TransactionDetailResponse]{}
	err := client.DoRequest(TransactionDetailMethod, TransactionDetailEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
