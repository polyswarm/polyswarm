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

// FinalizableCrowdsaleABI is the input ABI used to generate the binding from.
const FinalizableCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// FinalizableCrowdsale is an auto generated Go binding around an Ethereum contract.
type FinalizableCrowdsale struct {
	FinalizableCrowdsaleCaller     // Read-only binding to the contract
	FinalizableCrowdsaleTransactor // Write-only binding to the contract
}

// FinalizableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizableCrowdsaleSession struct {
	Contract     *FinalizableCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizableCrowdsaleCallerSession struct {
	Contract *FinalizableCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FinalizableCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizableCrowdsaleTransactorSession struct {
	Contract     *FinalizableCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizableCrowdsaleRaw struct {
	Contract *FinalizableCrowdsale // Generic contract binding to access the raw methods on
}

// FinalizableCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleCallerRaw struct {
	Contract *FinalizableCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizableCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleTransactorRaw struct {
	Contract *FinalizableCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizableCrowdsale creates a new instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsale(address common.Address, backend bind.ContractBackend) (*FinalizableCrowdsale, error) {
	contract, err := bindFinalizableCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsale{FinalizableCrowdsaleCaller: FinalizableCrowdsaleCaller{contract: contract}, FinalizableCrowdsaleTransactor: FinalizableCrowdsaleTransactor{contract: contract}}, nil
}

// NewFinalizableCrowdsaleCaller creates a new read-only instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*FinalizableCrowdsaleCaller, error) {
	contract, err := bindFinalizableCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleCaller{contract: contract}, nil
}

// NewFinalizableCrowdsaleTransactor creates a new write-only instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizableCrowdsaleTransactor, error) {
	contract, err := bindFinalizableCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleTransactor{contract: contract}, nil
}

// bindFinalizableCrowdsale binds a generic wrapper to an already deployed contract.
func bindFinalizableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsale *FinalizableCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsale.Contract.FinalizableCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsale *FinalizableCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.FinalizableCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsale *FinalizableCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.FinalizableCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.ClosingTime(&_FinalizableCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.ClosingTime(&_FinalizableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) HasClosed() (bool, error) {
	return _FinalizableCrowdsale.Contract.HasClosed(&_FinalizableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _FinalizableCrowdsale.Contract.HasClosed(&_FinalizableCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) IsFinalized() (bool, error) {
	return _FinalizableCrowdsale.Contract.IsFinalized(&_FinalizableCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) IsFinalized() (bool, error) {
	return _FinalizableCrowdsale.Contract.IsFinalized(&_FinalizableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.OpeningTime(&_FinalizableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.OpeningTime(&_FinalizableCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Owner() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Owner(&_FinalizableCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Owner(&_FinalizableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.Rate(&_FinalizableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.Rate(&_FinalizableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Token() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Token(&_FinalizableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Token() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Token(&_FinalizableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Wallet(&_FinalizableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Wallet(&_FinalizableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.WeiRaised(&_FinalizableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.WeiRaised(&_FinalizableCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.BuyTokens(&_FinalizableCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.BuyTokens(&_FinalizableCrowdsale.TransactOpts, _beneficiary)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Finalize(&_FinalizableCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Finalize(&_FinalizableCrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.TransferOwnership(&_FinalizableCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.TransferOwnership(&_FinalizableCrowdsale.TransactOpts, newOwner)
}
