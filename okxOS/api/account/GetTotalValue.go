package account

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	TotalValueEndpoint = "/api/v5/dex/balance/total-value-by-address"
	TotalValueMethod   = "GET"
)

type TotalValueParams struct {
	Address          string  `json:"address"`
	Chains           string  `json:"chains"`
	AssetType        *string `json:"assetType,omitempty"`
	ExcludeRiskToken *bool   `json:"excludeRiskToken,omitempty"`
}

type TotalValueResponse struct {
	TotalValue string `json:"totalValue"` // 以usdt为计价单位
}

func GetTotalValue(client *common.Client, params *TotalValueParams) (*common.Response[[]TotalValueResponse], error) {
	response := &common.Response[[]TotalValueResponse]{}
	err := client.DoRequest(TotalValueMethod, TotalValueEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
