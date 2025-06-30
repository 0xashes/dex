package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// EthereumWallet Ethereum 钱包结构体
type EthereumWallet struct {
	mnemonic   string
	password   string
	index      int
	privateKey string
	address    string
}

// Wallet 创建 Ethereum 钱包
func Wallet(mnemonic, password string, index int) (*EthereumWallet, error) {
	w := &EthereumWallet{
		mnemonic: mnemonic,
		password: password,
		index:    index,
	}

	privateKey, err := w.derivePrivateKey(index)
	if err != nil {
		return nil, fmt.Errorf("生成 Ethereum 私钥失败: %v", err)
	}
	w.privateKey = privateKey

	address, err := w.generateAddress(privateKey)
	if err != nil {
		return nil, fmt.Errorf("生成 Ethereum 地址失败: %v", err)
	}
	w.address = address

	return w, nil
}

// GetDerivedPath 获取 Ethereum 派生路径
func GetDerivedPath(index int) string {
	return fmt.Sprintf("m/44'/60'/0'/0/%d", index)
}

// DerivePrivateKey 派生私钥
func (w *EthereumWallet) DerivePrivateKey(index int) (string, error) {
	if w.privateKey != "" && w.index == index {
		return w.privateKey, nil
	}
	return w.derivePrivateKey(index)
}

// derivePrivateKey 内部方法，实际执行私钥派生
func (w *EthereumWallet) derivePrivateKey(index int) (string, error) {
	if !bip39.IsMnemonicValid(w.mnemonic) {
		return "", fmt.Errorf("无效的助记词")
	}
	seed := bip39.NewSeed(w.mnemonic, w.password)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", fmt.Errorf("生成种子失败: %v", err)
	}
	purpose, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	coinType, _ := purpose.NewChildKey(bip32.FirstHardenedChild + 60)
	account, _ := coinType.NewChildKey(bip32.FirstHardenedChild + 0)
	change, _ := account.NewChildKey(0)
	addressIndexKey, _ := change.NewChildKey(uint32(index))

	return hex.EncodeToString(addressIndexKey.Key), nil
}

// generateAddress 从私钥生成地址（内部方法）
func (w *EthereumWallet) generateAddress(privateKey string) (string, error) {
	prv, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("解码私钥失败: %v", err)
	}
	prvECDSA, err := crypto.ToECDSA(prv)
	if err != nil {
		return "", fmt.Errorf("私钥转换失败: %v", err)
	}

	publicKey := prvECDSA.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey)

	return address.Hex(), nil
}

func (w *EthereumWallet) GetAddress() string    { return w.address }
func (w *EthereumWallet) GetPrivateKey() string { return w.privateKey }
func (w *EthereumWallet) GetMnemonic() string   { return w.mnemonic }
func (w *EthereumWallet) GetIndex() int         { return w.index }
