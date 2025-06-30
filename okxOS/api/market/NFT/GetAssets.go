package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	AssetsEndpoint = "/api/v5/mktplace/nft/owner/asset-list" // 查询地址下的NFT资产
	AssetsMethod   = "GET"
)

type AssetsParams struct {
	Chain           string  `json:"chain"`                     // 链名称，必传，字符串类型
	ContractAddress *string `json:"contractAddress,omitempty"` // 合约地址，可选，字符串类型
	OwnerAddress    string  `json:"ownerAddress"`              // 资产所有者的地址，必传，字符串类型
	Cursor          *string `json:"cursor,omitempty"`          // 指向要检索的页面的游标，可选，字符串类型
	Limit           *string `json:"limit,omitempty"`           // 分页大小，默认 "10"，最大 "100"，可选，字符串类型
}

type AssetsResponse struct {
	Cursor string           `json:"cursor"`
	Data   []NFTAssetsModel `json:"data"`
}

func GetAssets(client *common.Client, params *AssetsParams) (*common.OKXOSResponse[AssetsResponse], error) {
	response := &common.OKXOSResponse[AssetsResponse]{}
	err := client.DoRequest(AssetsMethod, AssetsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
