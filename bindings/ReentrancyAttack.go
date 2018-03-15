// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ReentrancyAttackABI is the input ABI used to generate the binding from.
const ReentrancyAttackABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes4\"}],\"name\":\"callSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReentrancyAttackBin is the compiled bytecode used for deploying new contracts.
const ReentrancyAttackBin = `6060604052341561000f57600080fd5b60df8061001d6000396000f30060606040526004361060255763ffffffff60e060020a6000350416630a2df1ed8114602a575b600080fd5b3415603457600080fd5b605f7fffffffff00000000000000000000000000000000000000000000000000000000600435166061565b005b73ffffffffffffffffffffffffffffffffffffffff331660e060020a82046040518163ffffffff1660e060020a0281526004016000604051808303816000875af192505050151560b057600080fd5b505600a165627a7a723058203078ab7c7079b4857d026694fc4204e1f94b0c5f247cb877cceca737bdd30f1f0029`

// DeployReentrancyAttack deploys a new Ethereum contract, binding an instance of ReentrancyAttack to it.
func DeployReentrancyAttack(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyAttack, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyAttackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancyAttackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyAttack{ReentrancyAttackCaller: ReentrancyAttackCaller{contract: contract}, ReentrancyAttackTransactor: ReentrancyAttackTransactor{contract: contract}, ReentrancyAttackFilterer: ReentrancyAttackFilterer{contract: contract}}, nil
}

// ReentrancyAttack is an auto generated Go binding around an Ethereum contract.
type ReentrancyAttack struct {
	ReentrancyAttackCaller     // Read-only binding to the contract
	ReentrancyAttackTransactor // Write-only binding to the contract
	ReentrancyAttackFilterer   // Log filterer for contract events
}

// ReentrancyAttackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyAttackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyAttackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyAttackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyAttackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyAttackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyAttackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyAttackSession struct {
	Contract     *ReentrancyAttack // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyAttackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyAttackCallerSession struct {
	Contract *ReentrancyAttackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ReentrancyAttackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyAttackTransactorSession struct {
	Contract     *ReentrancyAttackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyAttackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyAttackRaw struct {
	Contract *ReentrancyAttack // Generic contract binding to access the raw methods on
}

// ReentrancyAttackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyAttackCallerRaw struct {
	Contract *ReentrancyAttackCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyAttackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyAttackTransactorRaw struct {
	Contract *ReentrancyAttackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyAttack creates a new instance of ReentrancyAttack, bound to a specific deployed contract.
func NewReentrancyAttack(address common.Address, backend bind.ContractBackend) (*ReentrancyAttack, error) {
	contract, err := bindReentrancyAttack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyAttack{ReentrancyAttackCaller: ReentrancyAttackCaller{contract: contract}, ReentrancyAttackTransactor: ReentrancyAttackTransactor{contract: contract}, ReentrancyAttackFilterer: ReentrancyAttackFilterer{contract: contract}}, nil
}

// NewReentrancyAttackCaller creates a new read-only instance of ReentrancyAttack, bound to a specific deployed contract.
func NewReentrancyAttackCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyAttackCaller, error) {
	contract, err := bindReentrancyAttack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyAttackCaller{contract: contract}, nil
}

// NewReentrancyAttackTransactor creates a new write-only instance of ReentrancyAttack, bound to a specific deployed contract.
func NewReentrancyAttackTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyAttackTransactor, error) {
	contract, err := bindReentrancyAttack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyAttackTransactor{contract: contract}, nil
}

// NewReentrancyAttackFilterer creates a new log filterer instance of ReentrancyAttack, bound to a specific deployed contract.
func NewReentrancyAttackFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyAttackFilterer, error) {
	contract, err := bindReentrancyAttack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyAttackFilterer{contract: contract}, nil
}

// bindReentrancyAttack binds a generic wrapper to an already deployed contract.
func bindReentrancyAttack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyAttackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyAttack *ReentrancyAttackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyAttack.Contract.ReentrancyAttackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyAttack *ReentrancyAttackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyAttack.Contract.ReentrancyAttackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyAttack *ReentrancyAttackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyAttack.Contract.ReentrancyAttackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyAttack *ReentrancyAttackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyAttack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyAttack *ReentrancyAttackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyAttack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyAttack *ReentrancyAttackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyAttack.Contract.contract.Transact(opts, method, params...)
}

// CallSender is a paid mutator transaction binding the contract method 0x0a2df1ed.
//
// Solidity: function callSender(data bytes4) returns()
func (_ReentrancyAttack *ReentrancyAttackTransactor) CallSender(opts *bind.TransactOpts, data [4]byte) (*types.Transaction, error) {
	return _ReentrancyAttack.contract.Transact(opts, "callSender", data)
}

// CallSender is a paid mutator transaction binding the contract method 0x0a2df1ed.
//
// Solidity: function callSender(data bytes4) returns()
func (_ReentrancyAttack *ReentrancyAttackSession) CallSender(data [4]byte) (*types.Transaction, error) {
	return _ReentrancyAttack.Contract.CallSender(&_ReentrancyAttack.TransactOpts, data)
}

// CallSender is a paid mutator transaction binding the contract method 0x0a2df1ed.
//
// Solidity: function callSender(data bytes4) returns()
func (_ReentrancyAttack *ReentrancyAttackTransactorSession) CallSender(data [4]byte) (*types.Transaction, error) {
	return _ReentrancyAttack.Contract.CallSender(&_ReentrancyAttack.TransactOpts, data)
}
