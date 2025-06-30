package dex

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	CreateLimitOrderEndpoint = "/api/v5/dex/aggregator/limit-order/save-order" // Create a DEX limit order
	CreateLimitOrderMethod   = "GET"
)

type CreateLimitOrderParams struct {
	OrderHash string `json:"orderHash"` // 订单hash，必传，字符串类型
	Signature string `json:"signature"` // 签名，必传，字符串类型
	ChainId   string `json:"chainId"`   // 链ID（如 "1" 表示 Ethereum），必传，字符串类型
	Data      struct {
		Salt          string `json:"salt"`          // 随机数盐，作为幂等标识，必传，字符串类型（时间戳秒值）
		MakingAmount  string `json:"makingAmount"`  // 卖出币数量，含精度，必传，字符串类型
		TakingAmount  string `json:"takingAmount"`  // 买入币数量，含精度，必传，字符串类型
		MakerToken    string `json:"makerToken"`    // 卖出币种合约地址，必传，字符串类型
		TakerToken    string `json:"takerToken"`    // 买入币种合约地址，必传，字符串类型
		Maker         string `json:"maker"`         // 创建订单的钱包地址，必传，字符串类型
		DeadLine      string `json:"deadLine"`      // 订单生效截止时间（时间戳秒值），必传，字符串类型
		AllowedSender string `json:"allowedSender"` // 允许的交易发送者地址，必传，字符串类型
		Receiver      string `json:"receiver"`      // 资产接收地址，必传，字符串类型
		MinReturn     string `json:"minReturn"`     // 买入币最小获得数量，含精度，必传，字符串类型
		PartiallyAble bool   `json:"partiallyAble"` // 是否支持部分成交，必传，布尔类型
	} `json:"data"` // 订单的数据对象，必传，嵌套对象
}

type CreateLimitOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 执行请求
func CreateLimitOrder(client *common.Client, params *CreateLimitOrderParams) (*CreateLimitOrderResponse, error) {
	response := &CreateLimitOrderResponse{}
	err := client.DoRequest(CreateLimitOrderMethod, CreateLimitOrderEndpoint, params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
