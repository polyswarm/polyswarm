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

// TimedCrowdsaleABI is the input ABI used to generate the binding from.
const TimedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// TimedCrowdsale is an auto generated Go binding around an Ethereum contract.
type TimedCrowdsale struct {
	TimedCrowdsaleCaller     // Read-only binding to the contract
	TimedCrowdsaleTransactor // Write-only binding to the contract
}

// TimedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedCrowdsaleSession struct {
	Contract     *TimedCrowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedCrowdsaleCallerSession struct {
	Contract *TimedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TimedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedCrowdsaleTransactorSession struct {
	Contract     *TimedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TimedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedCrowdsaleRaw struct {
	Contract *TimedCrowdsale // Generic contract binding to access the raw methods on
}

// TimedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedCrowdsaleCallerRaw struct {
	Contract *TimedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// TimedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedCrowdsaleTransactorRaw struct {
	Contract *TimedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimedCrowdsale creates a new instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsale(address common.Address, backend bind.ContractBackend) (*TimedCrowdsale, error) {
	contract, err := bindTimedCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsale{TimedCrowdsaleCaller: TimedCrowdsaleCaller{contract: contract}, TimedCrowdsaleTransactor: TimedCrowdsaleTransactor{contract: contract}}, nil
}

// NewTimedCrowdsaleCaller creates a new read-only instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*TimedCrowdsaleCaller, error) {
	contract, err := bindTimedCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleCaller{contract: contract}, nil
}

// NewTimedCrowdsaleTransactor creates a new write-only instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedCrowdsaleTransactor, error) {
	contract, err := bindTimedCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTransactor{contract: contract}, nil
}

// bindTimedCrowdsale binds a generic wrapper to an already deployed contract.
func bindTimedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsale.Contract.TimedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TimedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TimedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsale *TimedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsale *TimedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsale *TimedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.ClosingTime(&_TimedCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.ClosingTime(&_TimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleSession) HasClosed() (bool, error) {
	return _TimedCrowdsale.Contract.HasClosed(&_TimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _TimedCrowdsale.Contract.HasClosed(&_TimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.OpeningTime(&_TimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.OpeningTime(&_TimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) Rate() (*big.Int, error) {
	return _TimedCrowdsale.Contract.Rate(&_TimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _TimedCrowdsale.Contract.Rate(&_TimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Token() (common.Address, error) {
	return _TimedCrowdsale.Contract.Token(&_TimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _TimedCrowdsale.Contract.Token(&_TimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Wallet() (common.Address, error) {
	return _TimedCrowdsale.Contract.Wallet(&_TimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _TimedCrowdsale.Contract.Wallet(&_TimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsale.Contract.WeiRaised(&_TimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsale.Contract.WeiRaised(&_TimedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.BuyTokens(&_TimedCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.BuyTokens(&_TimedCrowdsale.TransactOpts, _beneficiary)
}
