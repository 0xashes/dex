package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiSubcriptionInfoEndpoint = "/api/v5/defi/calculator/subscribe-info"
	DeFiSubcriptionInfoMethod   = "POST"
)

type DeFiSubcriptionInfoParams struct {
	Address            *string `json:"address,omitempty"`
	InputAmount        string  `json:"inputAmount"`
	InvestmentCategory *string `json:"investmentCategory,omitempty"` // 认购类型：(0：默认类型；1：BRC-20)
	InvestmentId       string  `json:"investmentId"`
	InputTokenAddress  string  `json:"inputTokenAddress"`  // 认购代币的智能合约地址
	IsSingle           *bool   `json:"isSingle,omitempty"` // 判断是否为单币投资品：是：单币；否：多币
	Slippage           *string `json:"slippage,omitempty"` // 滑点：默认为 1%
}

type DeFiSubcriptionInfoResponse struct {
	InvestWithTokenList []InvestTokenModel `json:"investWithTokenList"`
	GainsTokenList      []InvestTokenModel `json:"gainsTokenList"`
	ReceiveTokenInfo    InvestTokenModel   `json:"receiveTokenInfo"`
	ApproveStatusList   []ApproveInfoModel `json:"approveStatusList"`
	IsSwapInvest        bool               `json:"isSwapInvest"`
	EstimateGasFee      string             `json:"estimateGasFee"`
	IsAllowSubscribe    bool               `json:"isAllowSubscribe"`
	ValidatorName       string             `json:"validatorName"`
}

type InvestTokenModel struct {
	CoinAmount     string `json:"coinAmount"`
	CurrencyAmount string `json:"currencyAmount"`
	TokenSymbol    string `json:"tokenSymbol"`
	TokenName      string `json:"tokenName"`
	TokenAddress   string `json:"tokenAddress"`
	TokenPrecision string `json:"tokenPrecision"`
	IsBaseToken    bool   `json:"isBaseToken"`
	Network        string `json:"network"`
	ChainID        string `json:"chainId"`
	DataType       string `json:"dataType"`
}

type ApproveInfoModel struct {
	TokenSymbol    string `json:"tokenSymbol"`
	TokenAddress   string `json:"tokenAddress"`
	TokenPrecision string `json:"tokenPrecision"`
	IsNeedApprove  bool   `json:"isNeedApprove"`
	ApproveAddress string `json:"approveAddress"`
	Network        string `json:"network"`
	ChainID        string `json:"chainId"`
	OrderType      string `json:"orderType"`
}

func GetDeFiSubcriptionInfo(client *common.Client, params *DeFiSubcriptionInfoParams) (*common.OKXOSResponse[DeFiSubcriptionInfoResponse], error) {
	response := &common.OKXOSResponse[DeFiSubcriptionInfoResponse]{}
	err := client.DoRequest(DeFiSubcriptionInfoMethod, DeFiSubcriptionInfoEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
