package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiProductDetailsEndpoint = "/api/v5/defi/explore/product/detail"
	DeFiProductDetailsMethod   = "GET"
)

type DeFiProductDetailsParams struct {
	InvestmentId       string  `json:"investmentId"`                 // 投资品 ID
	InvestmentCategory *string `json:"investmentCategory,omitempty"` // 认购类型：(例如，0：默认类型；1：BRC-20）
}

type DeFiProductDetailsResponse InvestmentModel

func GetDeFiProductsDetails(client *common.Client, params *DeFiProductDetailsParams) (*common.OKXOSResponse[DeFiProductDetailsResponse], error) {
	response := &common.OKXOSResponse[DeFiProductDetailsResponse]{}
	err := client.DoRequest(DeFiProductDetailsMethod, DeFiProductDetailsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
