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

// IncreasingPriceCrowdsaleABI is the input ABI used to generate the binding from.
const IncreasingPriceCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialRate\",\"type\":\"uint256\"},{\"name\":\"_finalRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// IncreasingPriceCrowdsale is an auto generated Go binding around an Ethereum contract.
type IncreasingPriceCrowdsale struct {
	IncreasingPriceCrowdsaleCaller     // Read-only binding to the contract
	IncreasingPriceCrowdsaleTransactor // Write-only binding to the contract
}

// IncreasingPriceCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncreasingPriceCrowdsaleSession struct {
	Contract     *IncreasingPriceCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncreasingPriceCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncreasingPriceCrowdsaleCallerSession struct {
	Contract *IncreasingPriceCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IncreasingPriceCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncreasingPriceCrowdsaleTransactorSession struct {
	Contract     *IncreasingPriceCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IncreasingPriceCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleRaw struct {
	Contract *IncreasingPriceCrowdsale // Generic contract binding to access the raw methods on
}

// IncreasingPriceCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleCallerRaw struct {
	Contract *IncreasingPriceCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// IncreasingPriceCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleTransactorRaw struct {
	Contract *IncreasingPriceCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncreasingPriceCrowdsale creates a new instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsale(address common.Address, backend bind.ContractBackend) (*IncreasingPriceCrowdsale, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsale{IncreasingPriceCrowdsaleCaller: IncreasingPriceCrowdsaleCaller{contract: contract}, IncreasingPriceCrowdsaleTransactor: IncreasingPriceCrowdsaleTransactor{contract: contract}}, nil
}

// NewIncreasingPriceCrowdsaleCaller creates a new read-only instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*IncreasingPriceCrowdsaleCaller, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleCaller{contract: contract}, nil
}

// NewIncreasingPriceCrowdsaleTransactor creates a new write-only instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*IncreasingPriceCrowdsaleTransactor, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleTransactor{contract: contract}, nil
}

// bindIncreasingPriceCrowdsale binds a generic wrapper to an already deployed contract.
func bindIncreasingPriceCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncreasingPriceCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncreasingPriceCrowdsale.Contract.IncreasingPriceCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.IncreasingPriceCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.IncreasingPriceCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncreasingPriceCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.ClosingTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.ClosingTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) FinalRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "finalRate")
	return *ret0, err
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) FinalRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.FinalRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) FinalRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.FinalRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) GetCurrentRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "getCurrentRate")
	return *ret0, err
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) GetCurrentRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.GetCurrentRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) GetCurrentRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.GetCurrentRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) HasClosed() (bool, error) {
	return _IncreasingPriceCrowdsale.Contract.HasClosed(&_IncreasingPriceCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _IncreasingPriceCrowdsale.Contract.HasClosed(&_IncreasingPriceCrowdsale.CallOpts)
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) InitialRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "initialRate")
	return *ret0, err
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) InitialRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.InitialRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) InitialRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.InitialRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.OpeningTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.OpeningTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Rate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.Rate(&_IncreasingPriceCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.Rate(&_IncreasingPriceCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Token() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Token(&_IncreasingPriceCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) Token() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Token(&_IncreasingPriceCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Wallet() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Wallet(&_IncreasingPriceCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Wallet(&_IncreasingPriceCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.WeiRaised(&_IncreasingPriceCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.WeiRaised(&_IncreasingPriceCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.BuyTokens(&_IncreasingPriceCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.BuyTokens(&_IncreasingPriceCrowdsale.TransactOpts, _beneficiary)
}
