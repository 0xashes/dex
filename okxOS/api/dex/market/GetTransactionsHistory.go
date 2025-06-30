package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	TransactionsHistoryEndpoint = "/api/v5/dex/market/trades"
	TransactionsHistoryMethod   = "GET"
)

type TransactionsHistoryParams struct {
	ChainIndex           string  `json:"chainIndex"`
	TokenContractAddress string  `json:"tokenContractAddress"`
	After                *string `json:"after,omitempty"` // 请求此id之前的分页内容
	Limit                *string `json:"limit,omitempty"` // 最大500条
}

type TransactionsHistoryResponse struct {
	Id                   string `json:"id"`
	ChainIndex           string `json:"chainIndex"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TxHashUrl            string `json:"txHashUrl"`
	UserAddress          string `json:"userAddress"`
	DexName              string `json:"dexName"`
	PoolLogoUrl          string `json:"poolLogoUrl"`
	Type                 string `json:"type"`
	ChangedTokenInfo     []struct {
		Amount               string `json:"amount"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenContractAddress string `json:"tokenContractAddress"`
	} `json:"changedTokenInfo"`
	Price      string `json:"price"`
	Volume     string `json:"volume"`
	Time       string `json:"time"`
	IsFiltered string `json:"isFiltered"`
}

func GetTransactionsHistory(client *common.Client, params *TransactionsHistoryParams) (*common.Response[[]TransactionsHistoryResponse], error) {
	response := &common.Response[[]TransactionsHistoryResponse]{}
	err := client.DoRequest(TransactionsHistoryMethod, TransactionsHistoryEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
