package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiPositionDetailsBasedProductEndpoint = "/api/v5/defi/user/investment/asset-detail"
	DeFiPositionDetailsBasedProductMethod   = "POST"
)

type DeFiPositionDetailsBasedProductParams struct {
	Address          string  `json:"address"`                    // 用户钱包地址
	ChainId          *string `json:"chainId,omitempty"`          // 公链 ID
	InvestmentId     string  `json:"investmentId"`               // investmentId 和 poolId 不可同时为空
	PoolId           *string `json:"poolId,omitempty"`           // BRC-20 类型投资品所需 ID (investmentId 和 poolId 不可同时为空)
	FarmInvestmentId *string `json:"farmInvestmentId,omitempty"` // 挖矿类型投资品 ID，用以检测是否可用于挖矿 (仅 Uni V3 类型投资品需此条件)
}

type DeFiPositionDetailsBasedProductResponse []InvestTokenBalanceModel

type InvestTokenBalanceModel struct {
	InvestmentName      string                     `json:"investmentName"`
	InvestmentId        string                     `json:"investmentId"`
	SourceInvestmentId  string                     `json:"sourceInvestmentId"`
	InvestType          int64                      `json:"investType"`
	InvestName          string                     `json:"investName"`
	AssetsTokenList     []AssetsTokenListModel     `json:"assetsTokenList"`
	RewardDefiTokenInfo []RewardDefiTokenInfoModel `json:"rewardDefiTokenInfo"`
	TotalValue          string                     `json:"totalValue"`
}

type AssetsTokenListModel struct {
	TokenSymbol    string `json:"tokenSymbol"`
	TokenLogo      string `json:"tokenLogo"`
	CoinAmount     string `json:"coinAmount"`
	CurrencyAmount string `json:"currencyAmount"`
	TokenPrecision int64  `json:"tokenPrecision"`
	TokenAddress   string `json:"tokenAddress"`
	Network        string `json:"network"`
}

type RewardDefiTokenInfoModel struct {
	BaseDefiTokenInfos []BaseDefiTokenInfoModel `json:"baseDefiTokenInfos"`
	ButtonType         string                   `json:"buttonType"`
	RewardType         int64                    `json:"rewardType"`
}

type BaseDefiTokenInfoModel struct {
	TokenSymbol    string `json:"tokenSymbol"`
	TokenAddress   string `json:"tokenAddress"`
	Network        string `json:"network"`
	TokenPrecision int64  `json:"tokenPrecision"`
	CoinAmount     string `json:"coinAmount"`
	CurrencyAmount string `json:"currencyAmount"`
	ButtonType     string `json:"buttonType"`
}

func GetDeFiPositionDetailsBasedProduct(client *common.Client, params *DeFiPositionDetailsBasedProductParams) (*common.OKXOSResponse[DeFiPositionDetailsBasedProductResponse], error) {
	response := &common.OKXOSResponse[DeFiPositionDetailsBasedProductResponse]{}
	err := client.DoRequest(DeFiPositionDetailsBasedProductMethod, DeFiPositionDetailsBasedProductEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
