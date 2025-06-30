package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	NFTCollectionDetailsEndpoint = "/api/v5/mktplace/nft/collection/detail"
	NFTCollectionDetailsMethod   = "GET"
)

type NFTCollectionDetailsParams struct {
	Slug string `json:"slug"` // 合集 slug 名称，即合集的唯一标识，必传，字符串类型
}

type NFTCollectionDetailsResponse CollectionModel

func GetNFTCollectionDetails(client *common.Client, params *NFTCollectionDetailsParams) (*common.OKXOSResponse[NFTCollectionDetailsResponse], error) {
	response := &common.OKXOSResponse[NFTCollectionDetailsResponse]{}
	err := client.DoRequest(NFTCollectionDetailsMethod, NFTCollectionDetailsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
