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

// ERC223ContractInterfaceABI is the input ABI used to generate the binding from.
const ERC223ContractInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC223ContractInterface is an auto generated Go binding around an Ethereum contract.
type ERC223ContractInterface struct {
	ERC223ContractInterfaceCaller     // Read-only binding to the contract
	ERC223ContractInterfaceTransactor // Write-only binding to the contract
}

// ERC223ContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC223ContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223ContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC223ContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223ContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC223ContractInterfaceSession struct {
	Contract     *ERC223ContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC223ContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC223ContractInterfaceCallerSession struct {
	Contract *ERC223ContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC223ContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC223ContractInterfaceTransactorSession struct {
	Contract     *ERC223ContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC223ContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC223ContractInterfaceRaw struct {
	Contract *ERC223ContractInterface // Generic contract binding to access the raw methods on
}

// ERC223ContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC223ContractInterfaceCallerRaw struct {
	Contract *ERC223ContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC223ContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC223ContractInterfaceTransactorRaw struct {
	Contract *ERC223ContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC223ContractInterface creates a new instance of ERC223ContractInterface, bound to a specific deployed contract.
func NewERC223ContractInterface(address common.Address, backend bind.ContractBackend) (*ERC223ContractInterface, error) {
	contract, err := bindERC223ContractInterface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC223ContractInterface{ERC223ContractInterfaceCaller: ERC223ContractInterfaceCaller{contract: contract}, ERC223ContractInterfaceTransactor: ERC223ContractInterfaceTransactor{contract: contract}}, nil
}

// NewERC223ContractInterfaceCaller creates a new read-only instance of ERC223ContractInterface, bound to a specific deployed contract.
func NewERC223ContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC223ContractInterfaceCaller, error) {
	contract, err := bindERC223ContractInterface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223ContractInterfaceCaller{contract: contract}, nil
}

// NewERC223ContractInterfaceTransactor creates a new write-only instance of ERC223ContractInterface, bound to a specific deployed contract.
func NewERC223ContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC223ContractInterfaceTransactor, error) {
	contract, err := bindERC223ContractInterface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC223ContractInterfaceTransactor{contract: contract}, nil
}

// bindERC223ContractInterface binds a generic wrapper to an already deployed contract.
func bindERC223ContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223ContractInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223ContractInterface *ERC223ContractInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223ContractInterface.Contract.ERC223ContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223ContractInterface *ERC223ContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223ContractInterface.Contract.ERC223ContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223ContractInterface *ERC223ContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223ContractInterface.Contract.ERC223ContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223ContractInterface *ERC223ContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223ContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223ContractInterface *ERC223ContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223ContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223ContractInterface *ERC223ContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223ContractInterface.Contract.contract.Transact(opts, method, params...)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ERC223ContractInterface *ERC223ContractInterfaceTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223ContractInterface.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ERC223ContractInterface *ERC223ContractInterfaceSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223ContractInterface.Contract.TokenFallback(&_ERC223ContractInterface.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ERC223ContractInterface *ERC223ContractInterfaceTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223ContractInterface.Contract.TokenFallback(&_ERC223ContractInterface.TransactOpts, _from, _value, _data)
}
