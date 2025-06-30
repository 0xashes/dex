package orinals

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	FractalOrdinalsTransactionHistoryEndpoint = "/api/v5/mktplace/nft/fractal-ordinals/trade-history"
	FractalOrdinalsTransactionHistoryMethod   = "GET"
)

type FractalOrdinalsTransactionHistoryParams struct {
	Slug               string  `json:"slug"`                         // 合集 slug 名称，即合集的唯一标识，必传，字符串类型
	Cursor             *string `json:"cursor,omitempty"`             // 指向要检索的页面的游标，可选，字符串类型
	Limit              *string `json:"limit,omitempty"`              // 分页大小，默认值 "10"，最大 "100"，每页返回最多交易历史的条数，可选，字符串类型
	Sort               *string `json:"sort,omitempty"`               // 排序规则，交易历史默认 desc 排序，可选 asc，可选，字符串类型
	IsBrc20            *bool   `json:"isBrc20,omitempty"`            // 获取全部 Fractal NFT 或 Brc20 合集的交易历史，默认返回 BRC-20，可选，布尔类型
	OrderSourceList    *[]int  `json:"orderSourceList,omitempty"`    // 获取某些平台的订单，代号如[34,54]（34:OKX 54:Magic Eden 55:OrdinalsWallet 57:Unisat），可选，整数数组类型
	TradeWalletAddress *string `json:"tradeWalletAddress,omitempty"` // 交易涉及的钱包地址，可选，字符串类型
	Type               *string `json:"type,omitempty"`               // 交易类型，多个以‘,’分割，默认返回 SALE,LIST,TRANSFER,CANCEL_LIST,UPDATE_PRICE，可选，字符串类型
}

type FractalOrdinalsTransactionHistoryResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func GetFractalOrdinalsTransactionHistory(client *common.Client, params *FractalOrdinalsTransactionHistoryParams) (*FractalOrdinalsTransactionHistoryResponse, error) {
	response := &FractalOrdinalsTransactionHistoryResponse{}
	err := client.DoRequest(FractalOrdinalsTransactionHistoryMethod, FractalOrdinalsTransactionHistoryEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
