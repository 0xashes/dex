package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/0xashes/dex/okxOS/api/account"

	"github.com/0xashes/dex/okxOS/common"
)

func TestRequest(t *testing.T) {
	client := common.NewClient()
	params := &account.SupportedChainsParams{}

	// params := &defiAccount.DeFiRedaptonHistoryParams{
	// 	InvestmentId: "7022",
	// 	UserAddress:  "0xe844f3de70771f878b3e1f8ab1111bc7fe7140ee",
	// }

	response, err := account.GetSupportedChains(client, params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("response.Code: %v\n", response.Code)
	fmt.Printf("response.Msg: %v\n", response.Msg)
	fmt.Printf("response.Data: %+v\n", response.Data)
}

//	defiAccount.DeFiPositionDetailsBasedProtocolParams{
//			AccountIdInfoList: []defiAccount.AccountIdInfo{
//				{
//					WalletAddressList: []defiAccount.WalletAddressParam{
//						{
//							ChainId:       "56",
//							WalletAddress: "0xe844f3de70771f878b3e1f8ab1111bc7fe7140ee",
//						},
//					},
//				},
//			},
//			AnalysisPlatformId: 114329,
//		}

func TestRequestRaw(t *testing.T) {
	test_url := account.SupportedChainsEndpoint
	test_method := account.SupportedChainsMethod

	test_url_params := account.SupportedChainsParams{}

	client := common.NewClient()
	raw, err := client.RequestRaw(test_method, test_url, test_url_params)
	// fmt.Println("Raw Response:", string(raw))
	if err != nil {
		log.Fatal(err)
	}

	// 使用 map[string]interface{} 解析 JSON
	var result map[string]interface{}
	err = json.Unmarshal(raw, &result)
	if err != nil {
		log.Fatalf("解析 JSON 失败: %v", err)
	}

	// 打印整个 JSON 结构
	// fmt.Printf("解析结果: %+v\n", result)

	// 格式化输出 JSON，方便阅读
	formattedJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("格式化 JSON 失败: %v", err)
	}
	fmt.Printf("格式化 JSON:\n%s\n", string(formattedJSON))

}
