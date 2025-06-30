package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiRedaptonHistoryEndpoint = "/api/v5/defi/user/investment/unstake-list"
	DeFiRedaptonHistoryMethod   = "GET"
)

type DeFiRedaptonHistoryParams struct {
	InvestmentId string `json:"investmentId"` // 投资品 ID (这里获取)
	UserAddress  string `json:"userAddress"`  // 用户钱包地址
}

type DeFiRedaptonHistoryResponse struct {
	UserAddress      string `json:"userAddress"`
	InvestmentID     int    `json:"investmentId"`
	CanClaimAll      bool   `json:"canClaimAll"`
	CoinAmount       string `json:"coinAmount"`
	CurrencyAmount   string `json:"currencyAmount"`
	RewardTokenInfos []struct {
		RewardType        string `json:"rewardType"`
		ClaimIndex        string `json:"claimIndex"`
		TokenSymbol       string `json:"tokenSymbol"`
		TokenLogo         string `json:"tokenLogo"`
		TokenAddress      string `json:"tokenAddress"`
		Network           string `json:"network"`
		TokenPrecision    string `json:"tokenPrecision"`
		CoinAmount        string `json:"coinAmount"`
		CurrencyAmount    string `json:"currencyAmount"`
		RewardDescription string `json:"rewardDescription"`
		RewardTip         string `json:"rewardTip"`
	} `json:"rewardTokenInfos"`
}

func GetDeFiRedaptonHistory(client *common.Client, params *DeFiRedaptonHistoryParams) (*common.OKXOSResponse[DeFiRedaptonHistoryResponse], error) {
	response := &common.OKXOSResponse[DeFiRedaptonHistoryResponse]{}
	err := client.DoRequest(DeFiRedaptonHistoryMethod, DeFiRedaptonHistoryEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
