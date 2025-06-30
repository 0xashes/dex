package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiBalanceEndpoint = "/api/v5/defi/user/balance-list"
	DeFiBalanceMethod   = "POST"
)

type DeFiBalanceParams struct {
	ChainId          string   `json:"chainId"`          // 公链 ID
	Address          string   `json:"address"`          // 用户钱包地址
	TokenAddressList []string `json:"tokenAddressList"` // 代币智能合约列表 (主网币合约地址需要提前约定)
}

type DeFiBalanceResponse []struct {
	TokenSymbol    string `json:"tokenSymbol"`
	TokenName      string `json:"tokenName"`
	TokenLogo      string `json:"tokenLogo"`
	TokenAddress   string `json:"tokenAddress"`
	Network        string `json:"network"`
	ChainId        string `json:"chainId"`
	TokenPrecision string `json:"tokenPrecision"`
	IsBaseToken    bool   `json:"isBaseToken"`
	CoinAmount     string `json:"coinAmount"`
	CurrencyAmount string `json:"currencyAmount"`
	BrowserUrl     string `json:"browserUrl"`
}

func GetDeFiBalance(client *common.Client, params *DeFiBalanceParams) (*common.OKXOSResponse[DeFiBalanceResponse], error) {
	response := &common.OKXOSResponse[DeFiBalanceResponse]{}
	err := client.DoRequest(DeFiBalanceMethod, DeFiBalanceEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
