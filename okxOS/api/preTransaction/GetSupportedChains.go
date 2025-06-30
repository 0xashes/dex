package preTransaction

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	SupportedChainsEndpoint = "/api/v5/dex/pre-transaction/supported/chain" // 获取交易上链 API 支持的链信息。
	SupportedChainsMethod   = "GET"
)

type SupportedChainsParams struct{}

type SupportedChainsResponse struct {
	Name       string `json:"name"`
	LogoUrl    string `json:"logoUrl"`
	ShortName  string `json:"shortName"`
	ChainIndex string `json:"chainIndex"`
}

// 执行请求
func GetSupportedChains(client *common.Client, params *SupportedChainsParams) (*common.Response[[]SupportedChainsResponse], error) {
	response := &common.Response[[]SupportedChainsResponse]{}
	err := client.DoRequest(SupportedChainsMethod, SupportedChainsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
