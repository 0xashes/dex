package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	NFTOffersEndpoint = "/api/v5/mktplace/nft/markets/listings" // 查询挂单(BuyNow)
	NFTOffersMethod   = "GET"
)

type NFTOffersParams struct {
	Chain             string  `json:"chain"`                       // 链名称，必传，字符串类型
	CollectionAddress *string `json:"collectionAddress,omitempty"` // NFT 合约地址，可选，字符串类型
	TokenId           *string `json:"tokenId,omitempty"`           // NFT 的 tokenId，可选，字符串类型
	Maker             *string `json:"maker,omitempty"`             // 订单发起人钱包地址，可选，字符串类型
	CreateAfter       *string `json:"createAfter,omitempty"`       // 创建订单的起始时间（秒），可选，字符串类型
	CreateBefore      *string `json:"createBefore,omitempty"`      // 创建订单的截止时间（秒），可选，字符串类型
	UpdateAfter       *string `json:"updateAfter,omitempty"`       // 更新订单的起始时间（秒），可选，字符串类型
	UpdateBefore      *string `json:"updateBefore,omitempty"`      // 更新订单的截止时间（秒），可选，字符串类型
	Status            *string `json:"status,omitempty"`            // 订单状态（active、inactive、cancelled、sold），可选，字符串类型
	Platform          *string `json:"platform,omitempty"`          // 目标挂单平台，默认 "okx"，可选，字符串类型
	Sort              *string `json:"sort,omitempty"`              // 排序规则（create_time_desc、update_time_desc、price_desc、price_asc），可选，字符串类型
	Limit             *string `json:"limit,omitempty"`             // 单页条数限制，默认 "50"，可选，字符串类型
	Cursor            *string `json:"cursor,omitempty"`            // 查询指定订单页的游标，可选，字符串类型
}

type NFTOffersResponse struct {
	Cursor string           `json:"cursor"`
	Data   []OKXOrdersModel `json:"data"`
}

func QueryNFTOffers(client *common.Client, params *NFTOffersParams) (*common.OKXOSResponse[NFTOffersResponse], error) {
	response := &common.OKXOSResponse[NFTOffersResponse]{}
	err := client.DoRequest(NFTOffersMethod, NFTOffersEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
