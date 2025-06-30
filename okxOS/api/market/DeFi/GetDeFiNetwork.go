package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiNetworkEndpoint = "/api/v5/defi/explore/network-list"
	DeFiNetworkMethod   = "GET"
)

type DeFiNetworkParams struct {
	Network *string `json:"network,omitempty"` // 网络名称
	ChainId *string `json:"chainId,omitempty"` // 公链 ID
}

type DeFiNetworkResponse struct {
	Symbol  string `json:"symbol"`
	ChainId string `json:"chainId"`
}

func GetDeFiNetwork(client *common.Client, params *DeFiNetworkParams) (*common.OKXOSResponse[[]DeFiNetworkResponse], error) {
	response := &common.OKXOSResponse[[]DeFiNetworkResponse]{}
	err := client.DoRequest(DeFiNetworkMethod, DeFiNetworkEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
