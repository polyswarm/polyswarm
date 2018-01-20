// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC23ContractInterfaceABI is the input ABI used to generate the binding from.
const ERC23ContractInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC23ContractInterface is an auto generated Go binding around an Ethereum contract.
type ERC23ContractInterface struct {
	ERC23ContractInterfaceCaller     // Read-only binding to the contract
	ERC23ContractInterfaceTransactor // Write-only binding to the contract
}

// ERC23ContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC23ContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC23ContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC23ContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC23ContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC23ContractInterfaceSession struct {
	Contract     *ERC23ContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC23ContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC23ContractInterfaceCallerSession struct {
	Contract *ERC23ContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ERC23ContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC23ContractInterfaceTransactorSession struct {
	Contract     *ERC23ContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ERC23ContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC23ContractInterfaceRaw struct {
	Contract *ERC23ContractInterface // Generic contract binding to access the raw methods on
}

// ERC23ContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC23ContractInterfaceCallerRaw struct {
	Contract *ERC23ContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC23ContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC23ContractInterfaceTransactorRaw struct {
	Contract *ERC23ContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC23ContractInterface creates a new instance of ERC23ContractInterface, bound to a specific deployed contract.
func NewERC23ContractInterface(address common.Address, backend bind.ContractBackend) (*ERC23ContractInterface, error) {
	contract, err := bindERC23ContractInterface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC23ContractInterface{ERC23ContractInterfaceCaller: ERC23ContractInterfaceCaller{contract: contract}, ERC23ContractInterfaceTransactor: ERC23ContractInterfaceTransactor{contract: contract}}, nil
}

// NewERC23ContractInterfaceCaller creates a new read-only instance of ERC23ContractInterface, bound to a specific deployed contract.
func NewERC23ContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC23ContractInterfaceCaller, error) {
	contract, err := bindERC23ContractInterface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC23ContractInterfaceCaller{contract: contract}, nil
}

// NewERC23ContractInterfaceTransactor creates a new write-only instance of ERC23ContractInterface, bound to a specific deployed contract.
func NewERC23ContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC23ContractInterfaceTransactor, error) {
	contract, err := bindERC23ContractInterface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC23ContractInterfaceTransactor{contract: contract}, nil
}

// bindERC23ContractInterface binds a generic wrapper to an already deployed contract.
func bindERC23ContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC23ContractInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC23ContractInterface *ERC23ContractInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC23ContractInterface.Contract.ERC23ContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC23ContractInterface *ERC23ContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC23ContractInterface.Contract.ERC23ContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC23ContractInterface *ERC23ContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC23ContractInterface.Contract.ERC23ContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC23ContractInterface *ERC23ContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC23ContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC23ContractInterface *ERC23ContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC23ContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC23ContractInterface *ERC23ContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC23ContractInterface.Contract.contract.Transact(opts, method, params...)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ERC23ContractInterface *ERC23ContractInterfaceTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC23ContractInterface.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ERC23ContractInterface *ERC23ContractInterfaceSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC23ContractInterface.Contract.TokenFallback(&_ERC23ContractInterface.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ERC23ContractInterface *ERC23ContractInterfaceTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC23ContractInterface.Contract.TokenFallback(&_ERC23ContractInterface.TransactOpts, _from, _value, _data)
}
