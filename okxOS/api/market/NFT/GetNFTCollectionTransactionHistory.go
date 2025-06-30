package nft

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	NFTCollectionTransactionHistoryEndpoint = "/api/v5/mktplace/nft/markets/trades"
	NFTCollectionTransactionHistoryMethod   = "GET"
)

type NFTCollectionTransactionHistoryParams struct {
	Chain             string  `json:"chain"`
	CollectionAddress string  `json:"collectionAddress"`
	Platform          *string `json:"platform,omitempty"`
	Limit             *string `json:"limit,omitempty"`
	TokenId           *string `json:"tokenId,omitempty"`
	Cursor            *string `json:"cursor,omitempty"`
	StartTime         *string `json:"startTime,omitempty"`
	EndTime           *string `json:"endTime,omitempty"`
}

type NFTCollectionTransactionHistoryResponse struct {
	Cursor string `json:"cursor"`
	Data   []struct {
		Amount            int     `json:"amount"`
		Chain             string  `json:"chain"`
		CollectionAddress string  `json:"collectionAddress"`
		CurrencyAddress   string  `json:"currencyAddress"`
		From              string  `json:"from"`
		Platform          string  `json:"platform"`
		Price             float64 `json:"price"`
		Timestamp         int64   `json:"timestamp"`
		To                string  `json:"to"`
		TokenID           string  `json:"tokenId"`
		TxHash            string  `json:"txHash"`
	} `json:"data"`
	Next  bool `json:"next"`
	Total int  `json:"total"`
}

func GetNFTCollectionTransactionHistory(client *common.Client, params *NFTCollectionTransactionHistoryParams) (*common.OKXOSResponse[NFTCollectionTransactionHistoryResponse], error) {
	response := &common.OKXOSResponse[NFTCollectionTransactionHistoryResponse]{}
	err := client.DoRequest(NFTCollectionTransactionHistoryMethod, NFTCollectionTransactionHistoryEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
