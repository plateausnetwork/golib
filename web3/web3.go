//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package web3

import (
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type (
	Web3 interface {
		GetBlockNumber() (uint64, error)
		SetChainId(chainId int64)
		SetAccount(privateKey string) error
		GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error)
		SendRawTransaction(to common.Address, data []byte, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (common.Hash, error)
		NewContract(abiString string, contractAddr ...string) (Contract, error)
	}
	Contract interface {
		AllMethods() []string
		Methods(methodName string) abi.Method
		Address() common.Address
		Call(methodName string, args ...interface{}) (interface{}, error)
		CallWithMultiReturns(methodName string, args ...interface{}) ([]interface{}, error)
		CallWithFromAndValue(methodName string, from common.Address, value *big.Int, args ...interface{}) ([]interface{}, error)
		EncodeABI(methodName string, args ...interface{}) ([]byte, error)
	}
	Options struct {
		RpcProvider string
	}
	web3Impl struct {
		web3 *web3.Web3
	}
)

func New(opt Options) (Web3, error) {
	web3, err := web3.NewWeb3(opt.RpcProvider)
	if err != nil {
		return nil, err
	}
	return web3Impl{web3: web3}, nil
}
