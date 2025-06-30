package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiProtocolsEndpoint = "/api/v5/defi/explore/protocol/list" // 香港IP不能用
	DeFiProtocolsMethod   = "GET"
)

type DeFiProtocolsParams struct {
	PlatformId   *string `json:"platformId,omitempty"`   // 平台 ID
	PlatformName *string `json:"platformName,omitempty"` // 平台官方名称，可选，字符串类型
}

type DeFiProtocolsResponse struct {
	PlatformID               int64                  `json:"platformId"`
	PlatformName             string                 `json:"platformName"`
	PlatformWebSite          string                 `json:"platformWebSite"`
	InvestmentApiUrlPattern  string                 `json:"investmentApiUrlPattern"`
	InvestmentPageUrlPattern string                 `json:"investmentPageUrlPattern"`
	PlatformMinInfos         []PlatformMinInfoModel `json:"platformMinInfos"`
}

type PlatformMinInfoModel struct {
	InvestmentID string `json:"investmentId"`
	ProtocolID   string `json:"protocolId"`
	Network      string `json:"network"`
	ChainID      string `json:"chainId"`
}

func GetDeFiProtocols(client *common.Client, params *DeFiProtocolsParams) (*common.OKXOSResponse[[]DeFiProtocolsResponse], error) {
	response := &common.OKXOSResponse[[]DeFiProtocolsResponse]{}
	err := client.DoRequest(DeFiProtocolsMethod, DeFiProtocolsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
} // 获取okx支持的协议和投资产品的简要信息
