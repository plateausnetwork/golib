package web3

import "errors"

var (
	ErrGetBlockNumber = func(err error) error {
		return errors.New("can't get current block height, details: " + err.Error())
	}
	ErrSetAccount = func(err error) error {
		return errors.New("private key is empty, details: " + err.Error())
	}
	ErrGetNonce = func(err error) error {
		return errors.New("getTransactionCount, details: " + err.Error())
	}
	ErrSendRawTransaction = func(err error) error {
		return errors.New("sendRawTransaction, details: " + err.Error())
	}
)
