package runes

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CreateRunesListingEndpoint = "/api/v5/mktplace/nft/runes/make-order"
	CreateRunesListingMethod   = "GET"
)

type CreateRunesListingParams struct {
	RunesId       string `json:"runesId"`       // runes 代币的唯一标识，例如 "840000:3"，必传，字符串类型
	WalletAddress string `json:"walletAddress"` // 持有 runesId 铭文的钱包地址，必传，字符串类型
	Utxo          string `json:"utxo"`          // rune 代币所在的 UTXO，格式 "txHash:vout"，例如 "d578a0967605257f75be625cbdc2506f2a52f9135f56f302badab6a3da54e0d4:0"，必传，字符串类型
	UnitPrice     string `json:"unitPrice"`     // 铭文挂单单价，单位为聪，BigDecimal 类型，使用字符串表示，必传
	TotalPrice    string `json:"totalPrice"`    // 铭文挂单总价，单位为 BTC，BigDecimal 类型，使用字符串表示，必传
	Psbt          string `json:"psbt"`          // 签名后的挂单 PSBT，base64 编码，必传，字符串类型
}

type CreateRunesListingResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func CreateRunesListing(client *common.Client, params *CreateRunesListingParams) (*CreateRunesListingResponse, error) {
	response := &CreateRunesListingResponse{}
	err := client.DoRequest(CreateRunesListingMethod, CreateRunesListingEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
