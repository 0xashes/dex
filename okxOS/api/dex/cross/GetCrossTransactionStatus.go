package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CrossTransactionStatusEndpoint = "/api/v5/dex/cross-chain/build-tx" // Generate the data to execute a cross-chain swap
	CrossTransactionStatusMethod   = "GET"
)

type CrossTransactionStatusParams struct {
	Hash    string  `json:"hash"`              // 源链 Hash 地址
	ChainId *string `json:"chainId,omitempty"` // 源链 ID (如1: Ethereum，更多可查看链 ID 列表))
}

type CrossTransactionStatusResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetCrossTransactionStatus(client *common.Client, params *CrossTransactionStatusParams) (*CrossTransactionStatusResponse, error) {
	response := &CrossTransactionStatusResponse{}
	err := client.DoRequest(CrossTransactionStatusMethod, CrossTransactionStatusEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
