package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	NFTsEndpoint = "/api/v5/mktplace/nft/asset/list"
	NFTsMethod   = "GET"
)

type NFTsParams struct {
	Chain           string  `json:"chain"`            // 链名称，必传，字符串类型
	ContractAddress string  `json:"contractAddress"`  // 合约地址，必传，字符串类型
	Cursor          *string `json:"cursor,omitempty"` // 指向要检索的页面的游标，可选，字符串类型
	Limit           *string `json:"limit,omitempty"`  // 分页大小，默认 "300"，最大 "300"，可选，字符串类型
}

type NFTsResponse struct {
	Cursor string     `json:"cursor"`
	Data   []NFTModel `json:"data"`
}

func GetNFTs(client *common.Client, params *NFTsParams) (*common.OKXOSResponse[NFTsResponse], error) {
	response := &common.OKXOSResponse[NFTsResponse]{}
	err := client.DoRequest(NFTsMethod, NFTsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
