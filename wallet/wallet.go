package wallet

import (
	"fmt"

	"github.com/0xashes/dex/wallet/ethereum"
	"github.com/0xashes/dex/wallet/solana"
	"github.com/tyler-smith/go-bip39"
)

// Wallet 定义通用钱包接口
type Wallet interface {
	DerivePrivateKey(index int) (string, error) // 派生私钥
	GetAddress() string                         // 直接返回存储的地址
	GetPrivateKey() string                      // 直接返回存储的私钥
	GetMnemonic() string                        // 获取助记词
	GetIndex() int                              // 获取索引
}

type Config struct {
	Mnemonic string
	Password string
}

// GenerateMnemonic 生成助记词
func GenerateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", fmt.Errorf("生成熵失败: %v", err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("生成助记词失败: %v", err)
	}
	return mnemonic, nil
}

// NewWallet 工厂函数，根据链类型创建钱包
func (c *Config) Wallet(chainType string, index int) (Wallet, error) {
	// 验证助记词
	if c.Mnemonic == "" {
		return nil, fmt.Errorf("助记词不能为空")
	}
	if !bip39.IsMnemonicValid(c.Mnemonic) {
		return nil, fmt.Errorf("无效的助记词")
	}
	if index < 0 {
		return nil, fmt.Errorf("索引不能为负数")
	}

	switch chainType {
	case "solana":
		return solana.Wallet(c.Mnemonic, c.Password, index)
	case "ethereum":
		return ethereum.Wallet(c.Mnemonic, c.Password, index)
	default:
		return nil, fmt.Errorf("不支持的链类型: %s", chainType)
	}
}
