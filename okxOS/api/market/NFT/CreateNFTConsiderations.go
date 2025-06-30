package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CreateNFTConsiderationsEndpoint = "/api/v5/mktplace/nft/markets/buy"
	CreateNFTConsiderationsMethod   = "POST"
)

type CreateNFTConsiderationsParams struct {
	Chain         string              `json:"chain"`
	WalletAddress string              `json:"walletAddress"`
	Items         []ConsiderationItem `json:"items"`
}

type ConsiderationItem struct {
	OrderId   string `json:"orderId"`   // 要购买的订单 ID
	TakeCount int    `json:"takeCount"` // 购买数量
}

type CreateNFTConsiderationsResponse struct {
	Errors interface{}      `json:"errors"`
	Steps  []OrderStepModel `json:"steps"`
}

// Deprecated
func CreateNFTConsiderations(client *common.Client, params *CreateNFTConsiderationsParams) (*common.OKXOSResponse[CreateNFTConsiderationsResponse], error) {
	response := &common.OKXOSResponse[CreateNFTConsiderationsResponse]{}
	err := client.DoRequest(CreateNFTConsiderationsMethod, CreateNFTConsiderationsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
