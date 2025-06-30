package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CancelNFTOrdersEndpoint = "/api/v5/mktplace/nft/markets/cancel-listing"
	CancelNFTOrdersMethod   = "POST"
)

type CancelNFTOrdersParams struct {
	Chain         string   `json:"chain"`
	WalletAddress string   `json:"walletAddress"`
	OrderIds      []string `json:"orderIds"`
}

type CancelNFTOrdersResponse struct {
	Errors interface{}      `json:"errors"`
	Steps  []OrderStepModel `json:"steps"`
}

// Deprecated
func CancelNFTOrders(client *common.Client, params *CancelNFTOrdersParams) (*common.OKXOSResponse[CancelNFTOrdersResponse], error) {
	response := &common.OKXOSResponse[CancelNFTOrdersResponse]{}
	err := client.DoRequest(CancelNFTOrdersMethod, CancelNFTOrdersEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
