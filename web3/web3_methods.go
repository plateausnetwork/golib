package web3

import (
	"github.com/chenzhijie/go-web3"
)

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
