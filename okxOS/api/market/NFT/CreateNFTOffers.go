package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CreateNFTOffersEndpoint = "/api/v5/mktplace/nft/markets/create-listing" // 创建挂单
	CreateNFTOffersMethod   = "POST"
)

type CreateNFTOffersParams struct {
	Chain             string `json:"chain"`
	WalletAddress     string `json:"walletAddress"`
	CollectionAddress string `json:"collectionAddress"`
	TokenId           string `json:"tokenId"`
	Price             string `json:"price"`
	CurrencyAddress   string `json:"currencyAddress"` // https://web3.okx.com/zh-hans/build/dev-docs/dex-api/dex-aggregation-faq
	Count             int    `json:"count"`
	ValidTime         string `json:"validTime"` // 挂单截止时间戳（秒），如 "2200000000"，必传，字符串类型
	Platform          string `json:"platform"`
}

type CreateNFTOffersResponse struct {
	Errors interface{}      `json:"errors"`
	Orders []OrderModel     `json:"orders"`
	Steps  []OrderStepModel `json:"steps"`
}

// Deprecated
func CreateNFTOffers(client *common.Client, params *CreateNFTOffersParams) (*common.OKXOSResponse[CreateNFTOffersResponse], error) {
	response := &common.OKXOSResponse[CreateNFTOffersResponse]{}
	err := client.DoRequest(CreateNFTOffersMethod, CreateNFTOffersEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
