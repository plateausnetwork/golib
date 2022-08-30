package web3

import (
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
)

func (i web3Impl) NewWeb3(rpcProviderURL string) (*web3.Web3, error) {
	web3, err := web3.NewWeb3(rpcProviderURL)
	if err != nil {
		return nil, err
	}
	return web3, nil
}

func (i web3Impl) GetBlockNumber(rpcProviderURL string) uint64 {
	web3, err := i.NewWeb3(rpcProviderURL)
	if err != nil {
		panic(err)
	}
	blockNumber, err := web3.Eth.GetBlockNumber()
	if err != nil {
		panic(err)
	}
	return blockNumber
}

func (i web3Impl) SetChainId(rpcProviderURL string, chainId int64) error {
	web3, err := i.NewWeb3(rpcProviderURL)
	if err != nil {
		return err
	}
	web3.Eth.SetChainId(chainId)
	return nil
}

func (i web3Impl) SetAccount(rpcProviderURL string, privateKey string) error {
	web3, err := i.NewWeb3(rpcProviderURL)
	if err != nil {
		return err
	}
	err = web3.Eth.SetAccount(privateKey)
	if err != nil {
		return err
	}
	return nil
}

func (i web3Impl) GetNonce(rpcProviderURL string, addr common.Address, blockNumber *big.Int) (uint64, error) {
	web3, err := i.NewWeb3(rpcProviderURL)
	if err != nil {
		panic(err)
	}
	nonce, err := web3.Eth.GetNonce(addr, blockNumber)
	if err != nil {
		panic(err)
	}
	return nonce, nil
}

func (i web3Impl) SendRawTransaction(tokenAddr string, rpcProviderURL string, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (common.Hash, error) {
	web3, err := i.NewWeb3(rpcProviderURL)
	if err != nil {
		panic(err)
	}
	txHash, err := web3.Eth.SendRawTransaction(
		common.HexToAddress(tokenAddr),
		big.NewInt(0),
		gasLimit,
		web3.Utils.ToGWei(1),
		data,
	)
	if err != nil {
		panic(err)
	}
	// => Send approve tx hash  0x837136c8b6f34b519c049d1cf703d3bba47d32f6801c25d83d0113bdc0e6936a
	return txHash, nil
}
