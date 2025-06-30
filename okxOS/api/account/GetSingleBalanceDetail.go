package account

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	SingleBalanceDetailEndpoint = "/api/v5/dex/balance/token-balances-by-address"
	SingleBalanceDetailMethod   = "POST"
)

type SingleBalanceDetailParams struct {
	Address                string      `json:"address"`
	TokenContractAddresses []TokenInfo `json:"tokenContractAddresses"`
	ChainIndex             *string     `json:"chainIndex,omitempty"`
	ExcludeRiskToken       *bool       `json:"excludeRiskToken,omitempty"`
}

type TokenInfo struct {
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

type SingleBalanceDetailResponse struct {
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

func GetSingleBalanceDetail(client *common.Client, params *SingleBalanceDetailParams) (*common.Response[[]SingleBalanceDetailResponse], error) {
	response := &common.Response[[]SingleBalanceDetailResponse]{}
	err := client.DoRequest(SingleBalanceDetailMethod, SingleBalanceDetailEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
