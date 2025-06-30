package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiPositionEndpoint = "/api/v5/defi/user/asset/platform/list"
	DeFiPositionMethod   = "POST"
)

type DeFiPositionParams struct {
	WalletAddressList []WalletAddressParam `json:"walletAddressList"` // 用户钱包地址列表
}

type WalletAddressParam struct {
	ChainId       string `json:"chainId"`       // 公链 ID
	WalletAddress string `json:"walletAddress"` // 用户钱包地址
}

type DeFiPositionResponse struct {
	WalletIDPlatformList []WalletPlatformModel `json:"walletIdPlatformList"`
}

type WalletPlatformModel struct {
	PlatformList []struct {
		PlatformList         string `json:"platformList"`
		PlatformName         string `json:"platformName"`       // 平台名称
		AnalysisPlatformID   int64  `json:"analysisPlatformId"` // 协议ID
		PlatformLogo         string `json:"platformLogo"`       // 平台图标
		PlatformColor        string `json:"platformColor"`
		CurrencyAmount       string `json:"currencyAmount"`
		IsSupportInvest      bool   `json:"isSupportInvest"` // OKX DEFI 是否支持
		BonusTag             int64  `json:"bonusTag"`
		PlatformUrl          string `json:"platformUrl"`
		NetworkBalanceVoList []struct {
			Network        string `json:"network"`        // 网络名称
			NetworkLogo    string `json:"networkLogo"`    // 网络图标
			ChainID        int64  `json:"chainId"`        // 链 ID
			CurrencyAmount string `json:"currencyAmount"` // 美元持仓金额
		} `json:"networkBalanceVoList"` // 不同网络下的持仓列表
	} `json:"platformList"`
}

func GetDeFiPosition(client *common.Client, params *DeFiPositionParams) (*common.OKXOSResponse[DeFiPositionResponse], error) {
	response := &common.OKXOSResponse[DeFiPositionResponse]{}
	err := client.DoRequest(DeFiPositionMethod, DeFiPositionEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
