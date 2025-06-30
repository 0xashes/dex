package account

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	BalancesDetailEndpoint = "/api/v5/dex/balance/all-token-balances-by-address"
	BalancesDetailMethod   = "GET"
)

type BalancesDetailParams struct {
	Address          string `json:"address"`
	Chains           string `json:"chains"`
	ExcludeRiskToken *bool  `json:"excludeRiskToken,omitempty"`
}

type BalancesDetailResponse struct {
	TokenAssets []struct {
		ChainIndex           string `json:"chainIndex"`
		TokenContractAddress string `json:"tokenContractAddress"`
		Address              string `json:"address"`
		Symbol               string `json:"symbol"`
		Balance              string `json:"balance"` // 币本位代币数量
		RawBalance           string `json:"rawBalance"`
		TokenPrice           string `json:"tokenPrice"`
		IsRiskToken          bool   `json:"isRiskToken"`
	} `json:"tokenAssets"`
}

func GetBalancesDetail(client *common.Client, params *BalancesDetailParams) (*common.Response[[]BalancesDetailResponse], error) {
	response := &common.Response[[]BalancesDetailResponse]{}
	err := client.DoRequest(BalancesDetailMethod, BalancesDetailEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
