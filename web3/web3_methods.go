package web3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (i web3Impl) GetBlockNumber() (uint64, error) {
	blockNumber, err := i.web3.Eth.GetBlockNumber()
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (i web3Impl) SetChainId(chainId int64) error {
	i.web3.Eth.SetChainId(chainId)
	return nil
}

func (i web3Impl) SetAccount(privateKey string) error {
	err := i.web3.Eth.SetAccount(privateKey)
	if err != nil {
		return err
	}
	return nil
}

func (i web3Impl) GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error) {
	nonce, err := i.web3.Eth.GetNonce(addr, blockNumber)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (i web3Impl) SendRawTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (common.Hash, error) {
	var hash common.Hash
	txHash, err := i.web3.Eth.SendRawTransaction(
		to,
		amount,
		gasLimit,
		gasPrice,
		data,
	)
	if err != nil {
		return hash, err
	}
	// => Send approve tx hash  0x837136c8b6f34b519c049d1cf703d3bba47d32f6801c25d83d0113bdc0e6936a
	return txHash, nil
}
