package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	NFTDetailsEndpoint = "/api/v5/mktplace/nft/asset/detail"
	NFTDetailsMethod   = "GET"
)

type NFTDetailsParams struct {
	Chain           string `json:"chain"`           // 链名称
	ContractAddress string `json:"contractAddress"` // 合约地址
	TokenId         string `json:"tokenId"`         // NFT 的 tokenId
}

type NFTDetailsResponse NFTModel

func GetNFTDetails(client *common.Client, params *NFTDetailsParams) (*common.OKXOSResponse[NFTDetailsResponse], error) {
	response := &common.OKXOSResponse[NFTDetailsResponse]{}
	err := client.DoRequest(NFTDetailsMethod, NFTDetailsEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
