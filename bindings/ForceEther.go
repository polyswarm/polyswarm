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

// ForceEtherABI is the input ABI used to generate the binding from.
const ForceEtherABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// ForceEtherBin is the compiled bytecode used for deploying new contracts.
const ForceEtherBin = `606060405260b3806100126000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f5074f4181146043575b600080fd5b3415604d57600080fd5b606c73ffffffffffffffffffffffffffffffffffffffff60043516606e565b005b8073ffffffffffffffffffffffffffffffffffffffff16ff00a165627a7a723058204dfc2b602a2446dcdd299282317a472a260dbc17aa0a0ddd15f9fc530318bca10029`

// DeployForceEther deploys a new Ethereum contract, binding an instance of ForceEther to it.
func DeployForceEther(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ForceEther, error) {
	parsed, err := abi.JSON(strings.NewReader(ForceEtherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ForceEtherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForceEther{ForceEtherCaller: ForceEtherCaller{contract: contract}, ForceEtherTransactor: ForceEtherTransactor{contract: contract}}, nil
}

// ForceEther is an auto generated Go binding around an Ethereum contract.
type ForceEther struct {
	ForceEtherCaller     // Read-only binding to the contract
	ForceEtherTransactor // Write-only binding to the contract
}

// ForceEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForceEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForceEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForceEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForceEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForceEtherSession struct {
	Contract     *ForceEther       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForceEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForceEtherCallerSession struct {
	Contract *ForceEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ForceEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForceEtherTransactorSession struct {
	Contract     *ForceEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ForceEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForceEtherRaw struct {
	Contract *ForceEther // Generic contract binding to access the raw methods on
}

// ForceEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForceEtherCallerRaw struct {
	Contract *ForceEtherCaller // Generic read-only contract binding to access the raw methods on
}

// ForceEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForceEtherTransactorRaw struct {
	Contract *ForceEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForceEther creates a new instance of ForceEther, bound to a specific deployed contract.
func NewForceEther(address common.Address, backend bind.ContractBackend) (*ForceEther, error) {
	contract, err := bindForceEther(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForceEther{ForceEtherCaller: ForceEtherCaller{contract: contract}, ForceEtherTransactor: ForceEtherTransactor{contract: contract}}, nil
}

// NewForceEtherCaller creates a new read-only instance of ForceEther, bound to a specific deployed contract.
func NewForceEtherCaller(address common.Address, caller bind.ContractCaller) (*ForceEtherCaller, error) {
	contract, err := bindForceEther(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ForceEtherCaller{contract: contract}, nil
}

// NewForceEtherTransactor creates a new write-only instance of ForceEther, bound to a specific deployed contract.
func NewForceEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*ForceEtherTransactor, error) {
	contract, err := bindForceEther(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ForceEtherTransactor{contract: contract}, nil
}

// bindForceEther binds a generic wrapper to an already deployed contract.
func bindForceEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForceEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForceEther *ForceEtherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ForceEther.Contract.ForceEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForceEther *ForceEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForceEther.Contract.ForceEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForceEther *ForceEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForceEther.Contract.ForceEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForceEther *ForceEtherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ForceEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForceEther *ForceEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForceEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForceEther *ForceEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForceEther.Contract.contract.Transact(opts, method, params...)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_ForceEther *ForceEtherTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _ForceEther.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_ForceEther *ForceEtherSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _ForceEther.Contract.DestroyAndSend(&_ForceEther.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_ForceEther *ForceEtherTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _ForceEther.Contract.DestroyAndSend(&_ForceEther.TransactOpts, _recipient)
}
