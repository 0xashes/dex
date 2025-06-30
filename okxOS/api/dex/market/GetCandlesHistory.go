package market

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CandlesHistoryEndpoint = "/api/v5/dex/market/candles"
	CandlesHistoryMethod   = "GET"
)

type CandlesHistoryParams struct {
	ChainIndex           string  `json:"chainIndex"`
	TokenContractAddress string  `json:"tokenContractAddress"`
	After                *string `json:"after,omitempty"` // 请求此id之前的分页内容
	Before               *string `json:"before,omitempty"`
	Bar                  *string `json:"bar,omitempty"`   // 时间粒度，默认值1m 如 [1s/1m/3m/5m/15m/30m/1H/2H/4H]; 香港时间开盘价k线：[6H/12H/1D/1W/1M/3M]; UTC时间开盘价k线：[/6Hutc/12Hutc/1Dutc/1Wutc/1Mutc/3Mutc]
	Limit                *string `json:"limit,omitempty"` // 最大500条
}

type CandlesHistoryResponse []string // 返回顺序分别是[ts,o,h,l,c,vol,volUsd,confirm]

func GetCandlesHistory(client *common.Client, params *CandlesHistoryParams) (*common.Response[[]CandlesHistoryResponse], error) {
	response := &common.Response[[]CandlesHistoryResponse]{}
	err := client.DoRequest(CandlesHistoryMethod, CandlesHistoryEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
