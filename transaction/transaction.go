package transaction

import (
	"context"
	"errors"

	"github.com/0xashes/dex/transaction/ethereum"
	"github.com/0xashes/dex/transaction/solana"
	"github.com/0xashes/dex/wallet"
)

type Transaction interface {
	Send(ctx context.Context, to string, amount string) (txHash string, err error)
	Swap(ctx context.Context, inputToken, outputToken, amount, slippage string) (txHash string, err error)
	PlaceLimitOrder(ctx context.Context, token, amount, price string, isBuy bool) (orderID string, err error)
	AddLiquidity(ctx context.Context, tokenA, tokenB, amountA, amountB, slippage string) (txHash string, err error)
	Stake(ctx context.Context, amount string) (txHash string, err error)
}

func NewTransaction(chain, platform string, w wallet.Wallet) (Transaction, error) {
	switch chain {
	case "solana":
		if platform == "pancakeswap" {
			return nil, errors.New("pancakeswap does not support solana")
		}
		return solana.Transaction(platform, w)
	case "ethereum":
		return ethereum.Transaction(platform, w)
	default:
		return nil, errors.New("unsupported chain")
	}
}
