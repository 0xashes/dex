package orinals

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CreateOrdinalsInscriptionListingEndpoint = "/api/v5/mktplace/nft/ordinals/okx/make-orders"
	CreateOrdinalsInscriptionListingMethod   = "GET"
)

type CreateOrdinalsInscriptionListingParams struct {
	NftId         string  `json:"nftId"`             // 挂单 NFT 的 ID，必传，字符串类型
	InscriptionId string  `json:"inscriptionId"`     // 铭文的 ID，必传，字符串类型
	OrderType     string  `json:"orderType"`         // 订单类型，挂单为 2，必传，整型
	UnitPrice     float64 `json:"unitPrice"`         // BRC-20 单价或 NFT 价格，单位：聪，必传，BigDecimal型
	IsBrc20       *bool   `json:"isBrc20,omitempty"` // 是否为 BRC-20 合集，默认为 true，可选，布尔类型
	Psbt          string  `json:"psbt"`              // 交易签名，base64 编码的 PSBT，必传，字符串类型
	TotalPrice    string  `json:"totalPrice"`        // NFT 价格或 BRC-20 总价，单位：聪，必传，BigDecimal型
}

type CreateOrdinalsInscriptionListingResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func CreateOrdinalsInscriptionListing(client *common.Client, params *CreateOrdinalsInscriptionListingParams) (*CreateOrdinalsInscriptionListingResponse, error) {
	response := &CreateOrdinalsInscriptionListingResponse{}
	err := client.DoRequest(CreateOrdinalsInscriptionListingMethod, CreateOrdinalsInscriptionListingEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
