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

// CappedCrowdsaleABI is the input ABI used to generate the binding from.
const CappedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// CappedCrowdsale is an auto generated Go binding around an Ethereum contract.
type CappedCrowdsale struct {
	CappedCrowdsaleCaller     // Read-only binding to the contract
	CappedCrowdsaleTransactor // Write-only binding to the contract
}

// CappedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedCrowdsaleSession struct {
	Contract     *CappedCrowdsale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CappedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedCrowdsaleCallerSession struct {
	Contract *CappedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CappedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedCrowdsaleTransactorSession struct {
	Contract     *CappedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CappedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedCrowdsaleRaw struct {
	Contract *CappedCrowdsale // Generic contract binding to access the raw methods on
}

// CappedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedCrowdsaleCallerRaw struct {
	Contract *CappedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CappedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedCrowdsaleTransactorRaw struct {
	Contract *CappedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedCrowdsale creates a new instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsale(address common.Address, backend bind.ContractBackend) (*CappedCrowdsale, error) {
	contract, err := bindCappedCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsale{CappedCrowdsaleCaller: CappedCrowdsaleCaller{contract: contract}, CappedCrowdsaleTransactor: CappedCrowdsaleTransactor{contract: contract}}, nil
}

// NewCappedCrowdsaleCaller creates a new read-only instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CappedCrowdsaleCaller, error) {
	contract, err := bindCappedCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleCaller{contract: contract}, nil
}

// NewCappedCrowdsaleTransactor creates a new write-only instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedCrowdsaleTransactor, error) {
	contract, err := bindCappedCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleTransactor{contract: contract}, nil
}

// bindCappedCrowdsale binds a generic wrapper to an already deployed contract.
func bindCappedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsale.Contract.CappedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.CappedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.CappedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsale *CappedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsale *CappedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsale *CappedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) Cap() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Cap(&_CappedCrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Cap(&_CappedCrowdsale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) EndTime() (*big.Int, error) {
	return _CappedCrowdsale.Contract.EndTime(&_CappedCrowdsale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) EndTime() (*big.Int, error) {
	return _CappedCrowdsale.Contract.EndTime(&_CappedCrowdsale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCaller) HasEnded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "hasEnded")
	return *ret0, err
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleSession) HasEnded() (bool, error) {
	return _CappedCrowdsale.Contract.HasEnded(&_CappedCrowdsale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) HasEnded() (bool, error) {
	return _CappedCrowdsale.Contract.HasEnded(&_CappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) Rate() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Rate(&_CappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Rate(&_CappedCrowdsale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) StartTime() (*big.Int, error) {
	return _CappedCrowdsale.Contract.StartTime(&_CappedCrowdsale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) StartTime() (*big.Int, error) {
	return _CappedCrowdsale.Contract.StartTime(&_CappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Token() (common.Address, error) {
	return _CappedCrowdsale.Contract.Token(&_CappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _CappedCrowdsale.Contract.Token(&_CappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Wallet() (common.Address, error) {
	return _CappedCrowdsale.Contract.Wallet(&_CappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _CappedCrowdsale.Contract.Wallet(&_CappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsale.Contract.WeiRaised(&_CappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsale.Contract.WeiRaised(&_CappedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.BuyTokens(&_CappedCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.BuyTokens(&_CappedCrowdsale.TransactOpts, beneficiary)
}
