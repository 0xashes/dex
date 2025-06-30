package solana

import (
	"crypto/ed25519"
	"fmt"

	"github.com/anyproto/go-slip10"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip39"
)

// SolanaWallet Solana 钱包结构体
type SolanaWallet struct {
	mnemonic   string
	password   string
	index      int
	privateKey string
	address    string
}

// Wallet 创建 Solana 钱包
func Wallet(mnemonic, password string, index int) (*SolanaWallet, error) {

	w := &SolanaWallet{
		mnemonic: mnemonic,
		password: password,
		index:    index,
	}

	privateKey, err := w.derivePrivateKey(index)
	if err != nil {
		return nil, fmt.Errorf("生成 Solana 私钥失败: %v", err)
	}
	w.privateKey = privateKey

	address, err := w.generateAddress(privateKey)
	if err != nil {
		return nil, fmt.Errorf("生成 Solana 地址失败: %v", err)
	}
	w.address = address

	return w, nil
}

// GetDerivedPath 获取 Solana 派生路径
func GetDerivedPath(index int) string {
	return fmt.Sprintf("m/44'/501'/%d'/0'", index)
}

// DerivePrivateKey 派生私钥
func (w *SolanaWallet) DerivePrivateKey(index int) (string, error) {
	if w.privateKey != "" && w.index == index {
		return w.privateKey, nil // 缓存结果
	}
	return w.derivePrivateKey(index)
}

// derivePrivateKey 内部方法，实际执行私钥派生
func (w *SolanaWallet) derivePrivateKey(index int) (string, error) {
	if !bip39.IsMnemonicValid(w.mnemonic) {
		return "", fmt.Errorf("无效的助记词")
	}
	seed := bip39.NewSeed(w.mnemonic, w.password)
	node, err := slip10.DeriveForPath(GetDerivedPath(index), seed)
	if err != nil {
		return "", fmt.Errorf("派生失败: %v", err)
	}
	_, prv := node.Keypair()

	return base58.Encode(prv), nil
}

// generateAddress 从私钥生成地址（内部方法）
func (w *SolanaWallet) generateAddress(privateKey string) (string, error) {
	prv, err := base58.Decode(privateKey)
	if err != nil {
		return "", fmt.Errorf("解码私钥失败: %v", err)
	}
	pub := ed25519.NewKeyFromSeed(prv[:32]).Public().(ed25519.PublicKey)

	return base58.Encode(pub), nil
}

func (w *SolanaWallet) GetAddress() string    { return w.address }
func (w *SolanaWallet) GetPrivateKey() string { return w.privateKey }
func (w *SolanaWallet) GetMnemonic() string   { return w.mnemonic }
func (w *SolanaWallet) GetIndex() int         { return w.index }
