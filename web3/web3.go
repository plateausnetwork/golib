//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package web3

import (
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
)

type Web3 interface {
	GetBlockNumber() (uint64, error)
	SetChainId(chainId int64)
	SetAccount(privateKey string) error
	GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error)
	SendRawTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (common.Hash, error)
}

type web3Impl struct {
	web3 *web3.Web3
}

type Options struct {
	RpcProvider string
}

func New(opt Options) (Web3, error) {
	web3, err := web3.NewWeb3(opt.RpcProvider)
	if err != nil {
		return nil, err
	}
	return web3Impl{web3: web3}, nil
}
