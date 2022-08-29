package web3

import (
	"encoding/hex"
	"math/big"
	"net/rpc"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Contract struct {
	abi      abi.ABI
	addr     common.Address
	provider *rpc.Client
}

func (i web3Impl) GetBlockNumber() uint64 {
	// change to your rpc provider
	var rpcProvider = "https://rpc.flashbots.net"
	web3, err := web3.NewWeb3(rpcProvider)
	if err != nil {
		panic(err)
	}
	blockNumber, err := web3.Eth.GetBlockNumber()
	if err != nil {
		panic(err)
	}
	return blockNumber
}

func (i web3Impl) NewWeb3() *web3.Web3 {
	// change to your rpc provider
	var rpcProviderURL = "https://rpc.flashbots.net"
	web3, err := web3.NewWeb3(rpcProviderURL)
	if err != nil {
		panic(err)
	}
	return web3
}

func (i web3Impl) SetChainId(chainId int64) {
	var rpcProvider = "https://rpc.flashbots.net"
	web3, err := web3.NewWeb3(rpcProvider)
	if err != nil {
		panic(err)
	}
	web3.Eth.SetChainId(1)
}

func (i web3Impl) SetAccount(privateKey string) error {
	var rpcProvider = "https://rpc.flashbots.net"
	pv, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	web3, err := web3.NewWeb3(rpcProvider)
	if err != nil {
		panic(err)
	}
	privateKey = hex.EncodeToString(crypto.FromECDSA(pv))
	err = web3.Eth.SetAccount(privateKey)
	if err != nil {
		panic(err)
	}
	return nil
}

func (i web3Impl) GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error) {
	var rpcProvider = "https://rpc.flashbots.net"
	web3, err := web3.NewWeb3(rpcProvider)
	if err != nil {
		panic(err)
	}
	nonce, err := web3.Eth.GetNonce(web3.Eth.Address(), nil)
	if err != nil {
		panic(err)
	}
	return nonce, nil
}

func (i web3Impl) SendRawTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (common.Hash, error) {
	var rpcProvider = "https://rpc.flashbots.net"
	var tokenAddr = "0xBB0E17EF65F82Ab018d8EDd776e8DD940327B28b" // AXS
	web3, err := web3.NewWeb3(rpcProvider)
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
