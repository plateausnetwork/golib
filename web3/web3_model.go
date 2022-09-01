package web3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type (
	Options struct {
		RpcProvider string
	}
	GetNonceOptions struct {
		Addr        common.Address
		BlockNumber *big.Int
	}
	SendRawTransactionOptions struct {
		To       common.Address
		Data     []byte
		Amount   *big.Int
		GasPrice *big.Int
		GasLimit uint64
	}
)
