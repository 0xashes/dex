package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	NFTCollectionEndpoint = "/api/v5/mktplace/nft/collection/list"
	NFTCollectionMethod   = "GET"
)

type NFTCollectionParams struct {
	Chain  *string `json:"chain,omitempty"`  // 链名称，可选，字符串类型
	Cursor *string `json:"cursor,omitempty"` // 指向要检索的页面的游标，可选，字符串类型
	Limit  *string `json:"limit,omitempty"`  // 分页大小，默认 "300"，最大 "300"，可选，字符串类型
}

type NFTCollectionResponse struct {
	Cursor string            `json:"cursor"`
	Data   []CollectionModel `json:"data"`
}

func GetNFTCollection(client *common.Client, params *NFTCollectionParams) (*common.OKXOSResponse[NFTCollectionResponse], error) {
	response := &common.OKXOSResponse[NFTCollectionResponse]{}
	err := client.DoRequest(NFTCollectionMethod, NFTCollectionEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
