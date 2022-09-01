package web3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (i web3Impl) GetBlockNumber() (uint64, error) {
	blockNumber, err := i.web3.Eth.GetBlockNumber()
	if err != nil {
		return 0, ErrGetBlockNumber(err)
	}
	return blockNumber, nil
}

func (i web3Impl) SetChainId(chainId int64) {
	i.web3.Eth.SetChainId(chainId)
}

func (i web3Impl) SetAccount(privateKey string) error {
	err := i.web3.Eth.SetAccount(privateKey)
	if err != nil {
		return ErrSetAccount(err)
	}
	return nil
}

func (i web3Impl) GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error) {
	nonce, err := i.web3.Eth.GetNonce(addr, blockNumber)
	if err != nil {
		return 0, ErrGetNonce(err)
	}
	return nonce, nil
}

func (i web3Impl) SendRawTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (hash common.Hash, err error) {
	hash, err = i.web3.Eth.SendRawTransaction(
		to,
		amount,
		gasLimit,
		gasPrice,
		data,
	)
	if err != nil {
		return hash, ErrSendRawTransaction(err)
	}
	return hash, nil
}
