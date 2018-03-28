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

// PostDeliveryCrowdsaleABI is the input ABI used to generate the binding from.
const PostDeliveryCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// PostDeliveryCrowdsale is an auto generated Go binding around an Ethereum contract.
type PostDeliveryCrowdsale struct {
	PostDeliveryCrowdsaleCaller     // Read-only binding to the contract
	PostDeliveryCrowdsaleTransactor // Write-only binding to the contract
}

// PostDeliveryCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostDeliveryCrowdsaleSession struct {
	Contract     *PostDeliveryCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PostDeliveryCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostDeliveryCrowdsaleCallerSession struct {
	Contract *PostDeliveryCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PostDeliveryCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostDeliveryCrowdsaleTransactorSession struct {
	Contract     *PostDeliveryCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PostDeliveryCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleRaw struct {
	Contract *PostDeliveryCrowdsale // Generic contract binding to access the raw methods on
}

// PostDeliveryCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleCallerRaw struct {
	Contract *PostDeliveryCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// PostDeliveryCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleTransactorRaw struct {
	Contract *PostDeliveryCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPostDeliveryCrowdsale creates a new instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsale(address common.Address, backend bind.ContractBackend) (*PostDeliveryCrowdsale, error) {
	contract, err := bindPostDeliveryCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsale{PostDeliveryCrowdsaleCaller: PostDeliveryCrowdsaleCaller{contract: contract}, PostDeliveryCrowdsaleTransactor: PostDeliveryCrowdsaleTransactor{contract: contract}}, nil
}

// NewPostDeliveryCrowdsaleCaller creates a new read-only instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*PostDeliveryCrowdsaleCaller, error) {
	contract, err := bindPostDeliveryCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleCaller{contract: contract}, nil
}

// NewPostDeliveryCrowdsaleTransactor creates a new write-only instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*PostDeliveryCrowdsaleTransactor, error) {
	contract, err := bindPostDeliveryCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleTransactor{contract: contract}, nil
}

// bindPostDeliveryCrowdsale binds a generic wrapper to an already deployed contract.
func bindPostDeliveryCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PostDeliveryCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PostDeliveryCrowdsale.Contract.PostDeliveryCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.PostDeliveryCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.PostDeliveryCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PostDeliveryCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.Balances(&_PostDeliveryCrowdsale.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.Balances(&_PostDeliveryCrowdsale.CallOpts, arg0)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.ClosingTime(&_PostDeliveryCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.ClosingTime(&_PostDeliveryCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) HasClosed() (bool, error) {
	return _PostDeliveryCrowdsale.Contract.HasClosed(&_PostDeliveryCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _PostDeliveryCrowdsale.Contract.HasClosed(&_PostDeliveryCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.OpeningTime(&_PostDeliveryCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.OpeningTime(&_PostDeliveryCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Rate() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.Rate(&_PostDeliveryCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.Rate(&_PostDeliveryCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Token() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Token(&_PostDeliveryCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Token() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Token(&_PostDeliveryCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Wallet() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Wallet(&_PostDeliveryCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Wallet(&_PostDeliveryCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.WeiRaised(&_PostDeliveryCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.WeiRaised(&_PostDeliveryCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.BuyTokens(&_PostDeliveryCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.BuyTokens(&_PostDeliveryCrowdsale.TransactOpts, _beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactor) WithdrawTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.contract.Transact(opts, "withdrawTokens")
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) WithdrawTokens() (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.WithdrawTokens(&_PostDeliveryCrowdsale.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorSession) WithdrawTokens() (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.WithdrawTokens(&_PostDeliveryCrowdsale.TransactOpts)
}
