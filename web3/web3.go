//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package web3

import (
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
)

type Web3 interface {
	NewWeb3(rpcProviderURL string) (*web3.Web3, error)
	GetBlockNumber(rpcProviderURL string) uint64
	SetChainId(rpcProviderURL string, chainId int64) error
	SetAccount(rpcProviderURL string, privateKey string) error
	GetNonce(rpcProviderURL string, addr common.Address, blockNumber *big.Int) (uint64, error)
	SendRawTransaction(tokenAddr string, rpcProviderURL string, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (common.Hash, error)
}

type web3Impl struct {
	rpcProvider string
}

type Options struct {
	RpcProvider string
}

func New(opt Options) Web3 {
	return &web3Impl{
		rpcProvider: opt.RpcProvider,
	}
}
