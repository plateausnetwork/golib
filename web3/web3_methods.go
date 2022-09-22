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

func (i web3Impl) GetNonce(addr common.Address, blockNumber *big.Int) (uint64, error) {
	nonce, err := i.web3.Eth.GetNonce(addr, blockNumber)
	if err != nil {
		return 0, ErrGetNonce(err)
	}
	return nonce, nil
}

func (i web3Impl) NewContract(abiString string, contractAddr ...string) (Contract, error) {
	contract, err := i.web3.Eth.NewContract(abiString, contractAddr...)
	if err != nil {
		return nil, ErrCreateContract(err)
	}
	return contract, nil
}

func (i web3Impl) SendRawTransaction(
	to common.Address, data []byte, amount *big.Int, gasPrice *big.Int, gasLimit uint64,
) (common.Hash, error) {
	hash, err := i.web3.Eth.SendRawTransaction(to, amount, gasLimit, gasPrice, data)
	if err != nil {
		return hash, ErrSendRawTransaction(err)
	}
	return hash, nil
}

func (i web3Impl) SetAccount(privateKey string) error {
	if err := i.web3.Eth.SetAccount(privateKey); err != nil {
		return ErrSetAccount(err)
	}
	return nil
}

func (i web3Impl) SetChainId(chainId int64) {
	i.web3.Eth.SetChainId(chainId)
}
