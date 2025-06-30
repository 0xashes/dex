package defi

import (
	"github.com/0xashes/dex/okxOS/common"
)

const (
	DeFiPositionDetailsBasedProtocolEndpoint = "/api/v5/defi/user/asset/platform/detail"
	DeFiPositionDetailsBasedProtocolMethod   = "POST"
)

type DeFiPositionDetailsBasedProtocolParams struct {
	AnalysisPlatformId int64           `json:"analysisPlatformId"` // 等于AnalysisPlatformID
	AccountIdInfoList  []AccountIdInfo `json:"accountIdInfoList"`
}

type AccountIdInfo struct {
	WalletAddressList []WalletAddressParam `json:"walletAddressList"`
}

type DeFiPositionDetailsBasedProtocolResponse struct {
	WalletIdPlatformDetailList []WalletIdPlatformDetail `json:"walletIdPlatformDetailList"`
	PlatformName               string                   `json:"platformName"`
	AnalysisPlatformId         int64                    `json:"analysisPlatformId"`
	PlatformLogo               string                   `json:"platformLogo"`
}

type WalletIdPlatformDetail struct {
	NetworkHoldVoList []struct {
		Network                  string                    `json:"network"`
		ChainId                  int64                     `json:"chainId"`
		InvestTokenBalanceVoList []InvestTokenBalanceModel `json:"investTokenBalanceVoList"`
		AvailableRewards         interface{}               `json:"availableRewards"`
		AirDropRewardInfo        interface{}               `json:"airDropRewardInfo"`
	} `json:"networkHoldVoList"`
	AccountId string `json:"accountId"`
}

func GetDeFiPositionDetailsBasedProtocol(client *common.Client, params *DeFiPositionDetailsBasedProtocolParams) (*common.OKXOSResponse[DeFiPositionDetailsBasedProtocolResponse], error) {
	response := &common.OKXOSResponse[DeFiPositionDetailsBasedProtocolResponse]{}
	err := client.DoRequest(DeFiPositionDetailsBasedProtocolMethod, DeFiPositionDetailsBasedProtocolEndpoint, params, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
