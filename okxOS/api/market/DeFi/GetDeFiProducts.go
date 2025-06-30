package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiProductsEndpoint = "/api/v5/defi/explore/product/list"
	DeFiProductsMethod   = "POST"
)

type DeFiProductsParams struct {
	SimplifyInvestType string     `json:"simplifyInvestType"`    // 查询投资品类型 (100：稳定币；101：单币赚币；102：多币赚币；103：机枪池)
	Network            string     `json:"network"`               // 链名缩写
	PoolVersion        *string    `json:"poolVersion,omitempty"` // 区分 V2 和 V3 类型流动池，默认为 V2
	PlatformIds        *[]string  `json:"platformIds,omitempty"` // 获取：/api/v5/defi/explore/protocol/list
	Sort               *SortModel `json:"sort,omitempty"`
	TokenIds           *[]string  `json:"tokenIds,omitempty"`
	Offset             *string    `json:"offset,omitempty"`
	Limit              string     `json:"limit"`
}

type SortModel struct {
	Orders *[]OrdedrModel `json:"Orders,omitempty"`
}

type OrdedrModel struct {
	Direction *string `json:"direction,omitempty"` // 排序方向：(升序 - ASC、降序 - DESC)
	Property  *string `json:"property,omitempty"`  // 排序的属性名（锁仓量 - TVL、APY利率 - RATE）
}

type DeFiProductsResponse struct {
	Investments []InvestmentModel `json:"investments"`
}

type InvestmentModel struct {
	InvestmentID    string                 `json:"investmentId"`
	InvestmentName  string                 `json:"investmentName"`
	ChainID         string                 `json:"chainId"`
	Rate            string                 `json:"rate"`
	InvestType      string                 `json:"investType"`
	PlatformName    string                 `json:"platformName"`
	PlatformID      string                 `json:"platformId"`
	RateType        string                 `json:"rateType"`
	TVL             string                 `json:"tvl"`
	UnderlyingToken []UnderlyingTokenModel `json:"underlyingToken"` // 质押代币
	Total           string                 `json:"total"`
	// 细节
	AnalysisPlatformId string                 `json:"analysisPlatformId"`
	IsInvestavle       string                 `json:"isInvestavle"`
	UtilizationRate    string                 `json:"utilizationRate"`
	EarnedToken        []UnderlyingTokenModel `json:"earnedToken"`
	LpToken            []UnderlyingTokenModel `json:"lpToken"`
}

type UnderlyingTokenModel struct {
	IsBaseToken   bool   `json:"isBaseToken"`
	TokenContract string `json:"tokenContract"`
	TokenSymbol   string `json:"tokenSymbol"`
}

func GetDeFiProducts(client *common.Client, params *DeFiProductsParams) (*common.OKXOSResponse[DeFiProductsResponse], error) {
	response := &common.OKXOSResponse[DeFiProductsResponse]{}
	err := client.DoRequest(DeFiProductsMethod, DeFiProductsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
