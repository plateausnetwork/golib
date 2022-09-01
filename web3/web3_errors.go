package web3

import "errors"

var (
	ErrGetBlockNumber = func(err error) error {
		return errors.New("could not get block number, details: " + err.Error())
	}
	ErrSetAccount = func(err error) error {
		return errors.New("could not set account, details: " + err.Error())
	}
	ErrGetNonce = func(err error) error {
		return errors.New("could not get nonce, details: " + err.Error())
	}
	ErrSendRawTransaction = func(err error) error {
		return errors.New("could not send raw transaction, details: " + err.Error())
	}
)
