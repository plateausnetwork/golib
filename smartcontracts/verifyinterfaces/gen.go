// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifyinterfaces

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VerifyinterfacesMetaData contains all meta data concerning the Verifyinterfaces contract.
var VerifyinterfacesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"_interfaceIds\",\"type\":\"bytes4[]\"}],\"name\":\"getSupportedInterfaces\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"_interfaceIds\",\"type\":\"bytes4[]\"}],\"name\":\"supportsAllInterfaces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VerifyinterfacesABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyinterfacesMetaData.ABI instead.
var VerifyinterfacesABI = VerifyinterfacesMetaData.ABI

// Verifyinterfaces is an auto generated Go binding around an Ethereum contract.
type Verifyinterfaces struct {
	VerifyinterfacesCaller     // Read-only binding to the contract
	VerifyinterfacesTransactor // Write-only binding to the contract
	VerifyinterfacesFilterer   // Log filterer for contract events
}

// VerifyinterfacesCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyinterfacesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyinterfacesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyinterfacesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyinterfacesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyinterfacesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyinterfacesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyinterfacesSession struct {
	Contract     *Verifyinterfaces // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyinterfacesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyinterfacesCallerSession struct {
	Contract *VerifyinterfacesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VerifyinterfacesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyinterfacesTransactorSession struct {
	Contract     *VerifyinterfacesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VerifyinterfacesRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyinterfacesRaw struct {
	Contract *Verifyinterfaces // Generic contract binding to access the raw methods on
}

// VerifyinterfacesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyinterfacesCallerRaw struct {
	Contract *VerifyinterfacesCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyinterfacesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyinterfacesTransactorRaw struct {
	Contract *VerifyinterfacesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyinterfaces creates a new instance of Verifyinterfaces, bound to a specific deployed contract.
func NewVerifyinterfaces(address common.Address, backend bind.ContractBackend) (*Verifyinterfaces, error) {
	contract, err := bindVerifyinterfaces(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifyinterfaces{VerifyinterfacesCaller: VerifyinterfacesCaller{contract: contract}, VerifyinterfacesTransactor: VerifyinterfacesTransactor{contract: contract}, VerifyinterfacesFilterer: VerifyinterfacesFilterer{contract: contract}}, nil
}

// NewVerifyinterfacesCaller creates a new read-only instance of Verifyinterfaces, bound to a specific deployed contract.
func NewVerifyinterfacesCaller(address common.Address, caller bind.ContractCaller) (*VerifyinterfacesCaller, error) {
	contract, err := bindVerifyinterfaces(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyinterfacesCaller{contract: contract}, nil
}

// NewVerifyinterfacesTransactor creates a new write-only instance of Verifyinterfaces, bound to a specific deployed contract.
func NewVerifyinterfacesTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyinterfacesTransactor, error) {
	contract, err := bindVerifyinterfaces(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyinterfacesTransactor{contract: contract}, nil
}

// NewVerifyinterfacesFilterer creates a new log filterer instance of Verifyinterfaces, bound to a specific deployed contract.
func NewVerifyinterfacesFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyinterfacesFilterer, error) {
	contract, err := bindVerifyinterfaces(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyinterfacesFilterer{contract: contract}, nil
}

// bindVerifyinterfaces binds a generic wrapper to an already deployed contract.
func bindVerifyinterfaces(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifyinterfacesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifyinterfaces *VerifyinterfacesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifyinterfaces.Contract.VerifyinterfacesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifyinterfaces *VerifyinterfacesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifyinterfaces.Contract.VerifyinterfacesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifyinterfaces *VerifyinterfacesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifyinterfaces.Contract.VerifyinterfacesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifyinterfaces *VerifyinterfacesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifyinterfaces.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifyinterfaces *VerifyinterfacesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifyinterfaces.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifyinterfaces *VerifyinterfacesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifyinterfaces.Contract.contract.Transact(opts, method, params...)
}

// GetSupportedInterfaces is a free data retrieval call binding the contract method 0x77e6b4cc.
//
// Solidity: function getSupportedInterfaces(address _contract, bytes4[] _interfaceIds) view returns(bool[])
func (_Verifyinterfaces *VerifyinterfacesCaller) GetSupportedInterfaces(opts *bind.CallOpts, _contract common.Address, _interfaceIds [][4]byte) ([]bool, error) {
	var out []interface{}
	err := _Verifyinterfaces.contract.Call(opts, &out, "getSupportedInterfaces", _contract, _interfaceIds)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetSupportedInterfaces is a free data retrieval call binding the contract method 0x77e6b4cc.
//
// Solidity: function getSupportedInterfaces(address _contract, bytes4[] _interfaceIds) view returns(bool[])
func (_Verifyinterfaces *VerifyinterfacesSession) GetSupportedInterfaces(_contract common.Address, _interfaceIds [][4]byte) ([]bool, error) {
	return _Verifyinterfaces.Contract.GetSupportedInterfaces(&_Verifyinterfaces.CallOpts, _contract, _interfaceIds)
}

// GetSupportedInterfaces is a free data retrieval call binding the contract method 0x77e6b4cc.
//
// Solidity: function getSupportedInterfaces(address _contract, bytes4[] _interfaceIds) view returns(bool[])
func (_Verifyinterfaces *VerifyinterfacesCallerSession) GetSupportedInterfaces(_contract common.Address, _interfaceIds [][4]byte) ([]bool, error) {
	return _Verifyinterfaces.Contract.GetSupportedInterfaces(&_Verifyinterfaces.CallOpts, _contract, _interfaceIds)
}

// SupportsAllInterfaces is a free data retrieval call binding the contract method 0x4b9dd904.
//
// Solidity: function supportsAllInterfaces(address _contract, bytes4[] _interfaceIds) view returns(bool)
func (_Verifyinterfaces *VerifyinterfacesCaller) SupportsAllInterfaces(opts *bind.CallOpts, _contract common.Address, _interfaceIds [][4]byte) (bool, error) {
	var out []interface{}
	err := _Verifyinterfaces.contract.Call(opts, &out, "supportsAllInterfaces", _contract, _interfaceIds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAllInterfaces is a free data retrieval call binding the contract method 0x4b9dd904.
//
// Solidity: function supportsAllInterfaces(address _contract, bytes4[] _interfaceIds) view returns(bool)
func (_Verifyinterfaces *VerifyinterfacesSession) SupportsAllInterfaces(_contract common.Address, _interfaceIds [][4]byte) (bool, error) {
	return _Verifyinterfaces.Contract.SupportsAllInterfaces(&_Verifyinterfaces.CallOpts, _contract, _interfaceIds)
}

// SupportsAllInterfaces is a free data retrieval call binding the contract method 0x4b9dd904.
//
// Solidity: function supportsAllInterfaces(address _contract, bytes4[] _interfaceIds) view returns(bool)
func (_Verifyinterfaces *VerifyinterfacesCallerSession) SupportsAllInterfaces(_contract common.Address, _interfaceIds [][4]byte) (bool, error) {
	return _Verifyinterfaces.Contract.SupportsAllInterfaces(&_Verifyinterfaces.CallOpts, _contract, _interfaceIds)
}

// SupportsInterface is a free data retrieval call binding the contract method 0xd9057007.
//
// Solidity: function supportsInterface(address _contract, bytes4 _interfaceId) view returns(bool)
func (_Verifyinterfaces *VerifyinterfacesCaller) SupportsInterface(opts *bind.CallOpts, _contract common.Address, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Verifyinterfaces.contract.Call(opts, &out, "supportsInterface", _contract, _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0xd9057007.
//
// Solidity: function supportsInterface(address _contract, bytes4 _interfaceId) view returns(bool)
func (_Verifyinterfaces *VerifyinterfacesSession) SupportsInterface(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Verifyinterfaces.Contract.SupportsInterface(&_Verifyinterfaces.CallOpts, _contract, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0xd9057007.
//
// Solidity: function supportsInterface(address _contract, bytes4 _interfaceId) view returns(bool)
func (_Verifyinterfaces *VerifyinterfacesCallerSession) SupportsInterface(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Verifyinterfaces.Contract.SupportsInterface(&_Verifyinterfaces.CallOpts, _contract, _interfaceId)
}
