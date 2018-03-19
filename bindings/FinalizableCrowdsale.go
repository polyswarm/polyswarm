// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// FinalizableCrowdsaleABI is the input ABI used to generate the binding from.
const FinalizableCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// FinalizableCrowdsale is an auto generated Go binding around an Ethereum contract.
type FinalizableCrowdsale struct {
	FinalizableCrowdsaleCaller     // Read-only binding to the contract
	FinalizableCrowdsaleTransactor // Write-only binding to the contract
	FinalizableCrowdsaleFilterer   // Log filterer for contract events
}

// FinalizableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizableCrowdsaleFilterer struct {
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
	contract, err := bindFinalizableCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsale{FinalizableCrowdsaleCaller: FinalizableCrowdsaleCaller{contract: contract}, FinalizableCrowdsaleTransactor: FinalizableCrowdsaleTransactor{contract: contract}, FinalizableCrowdsaleFilterer: FinalizableCrowdsaleFilterer{contract: contract}}, nil
}

// NewFinalizableCrowdsaleCaller creates a new read-only instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*FinalizableCrowdsaleCaller, error) {
	contract, err := bindFinalizableCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleCaller{contract: contract}, nil
}

// NewFinalizableCrowdsaleTransactor creates a new write-only instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizableCrowdsaleTransactor, error) {
	contract, err := bindFinalizableCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleTransactor{contract: contract}, nil
}

// NewFinalizableCrowdsaleFilterer creates a new log filterer instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizableCrowdsaleFilterer, error) {
	contract, err := bindFinalizableCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleFilterer{contract: contract}, nil
}

// bindFinalizableCrowdsale binds a generic wrapper to an already deployed contract.
func bindFinalizableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// FinalizableCrowdsaleFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleFinalizedIterator struct {
	Event *FinalizableCrowdsaleFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FinalizableCrowdsaleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FinalizableCrowdsaleFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FinalizableCrowdsaleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleFinalized represents a Finalized event raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) FilterFinalized(opts *bind.FilterOpts) (*FinalizableCrowdsaleFinalizedIterator, error) {

	logs, sub, err := _FinalizableCrowdsale.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleFinalizedIterator{contract: _FinalizableCrowdsale.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleFinalized) (event.Subscription, error) {

	logs, sub, err := _FinalizableCrowdsale.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleFinalized)
				if err := _FinalizableCrowdsale.contract.UnpackLog(event, "Finalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FinalizableCrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleOwnershipTransferredIterator struct {
	Event *FinalizableCrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FinalizableCrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FinalizableCrowdsaleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FinalizableCrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinalizableCrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalizableCrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleOwnershipTransferredIterator{contract: _FinalizableCrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalizableCrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleOwnershipTransferred)
				if err := _FinalizableCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FinalizableCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleTokenPurchaseIterator struct {
	Event *FinalizableCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FinalizableCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleTokenPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FinalizableCrowdsaleTokenPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FinalizableCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleTokenPurchase represents a TokenPurchase event raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*FinalizableCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleTokenPurchaseIterator{contract: _FinalizableCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleTokenPurchase)
				if err := _FinalizableCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
