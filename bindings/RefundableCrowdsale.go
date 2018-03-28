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

// RefundableCrowdsaleABI is the input ABI used to generate the binding from.
const RefundableCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_goal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// RefundableCrowdsale is an auto generated Go binding around an Ethereum contract.
type RefundableCrowdsale struct {
	RefundableCrowdsaleCaller     // Read-only binding to the contract
	RefundableCrowdsaleTransactor // Write-only binding to the contract
}

// RefundableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundableCrowdsaleSession struct {
	Contract     *RefundableCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RefundableCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundableCrowdsaleCallerSession struct {
	Contract *RefundableCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RefundableCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundableCrowdsaleTransactorSession struct {
	Contract     *RefundableCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RefundableCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundableCrowdsaleRaw struct {
	Contract *RefundableCrowdsale // Generic contract binding to access the raw methods on
}

// RefundableCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundableCrowdsaleCallerRaw struct {
	Contract *RefundableCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// RefundableCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundableCrowdsaleTransactorRaw struct {
	Contract *RefundableCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundableCrowdsale creates a new instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsale(address common.Address, backend bind.ContractBackend) (*RefundableCrowdsale, error) {
	contract, err := bindRefundableCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsale{RefundableCrowdsaleCaller: RefundableCrowdsaleCaller{contract: contract}, RefundableCrowdsaleTransactor: RefundableCrowdsaleTransactor{contract: contract}}, nil
}

// NewRefundableCrowdsaleCaller creates a new read-only instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*RefundableCrowdsaleCaller, error) {
	contract, err := bindRefundableCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleCaller{contract: contract}, nil
}

// NewRefundableCrowdsaleTransactor creates a new write-only instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundableCrowdsaleTransactor, error) {
	contract, err := bindRefundableCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleTransactor{contract: contract}, nil
}

// bindRefundableCrowdsale binds a generic wrapper to an already deployed contract.
func bindRefundableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundableCrowdsale *RefundableCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RefundableCrowdsale.Contract.RefundableCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundableCrowdsale *RefundableCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.RefundableCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundableCrowdsale *RefundableCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.RefundableCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundableCrowdsale *RefundableCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RefundableCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.ClosingTime(&_RefundableCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.ClosingTime(&_RefundableCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "goal")
	return *ret0, err
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Goal() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Goal(&_RefundableCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Goal() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Goal(&_RefundableCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "goalReached")
	return *ret0, err
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) GoalReached() (bool, error) {
	return _RefundableCrowdsale.Contract.GoalReached(&_RefundableCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) GoalReached() (bool, error) {
	return _RefundableCrowdsale.Contract.GoalReached(&_RefundableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) HasClosed() (bool, error) {
	return _RefundableCrowdsale.Contract.HasClosed(&_RefundableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _RefundableCrowdsale.Contract.HasClosed(&_RefundableCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) IsFinalized() (bool, error) {
	return _RefundableCrowdsale.Contract.IsFinalized(&_RefundableCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) IsFinalized() (bool, error) {
	return _RefundableCrowdsale.Contract.IsFinalized(&_RefundableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.OpeningTime(&_RefundableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.OpeningTime(&_RefundableCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Owner() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Owner(&_RefundableCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Owner(&_RefundableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Rate() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Rate(&_RefundableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Rate(&_RefundableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Token() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Token(&_RefundableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Token() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Token(&_RefundableCrowdsale.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Vault() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Vault(&_RefundableCrowdsale.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Vault() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Vault(&_RefundableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Wallet() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Wallet(&_RefundableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Wallet(&_RefundableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.WeiRaised(&_RefundableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.WeiRaised(&_RefundableCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.BuyTokens(&_RefundableCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.BuyTokens(&_RefundableCrowdsale.TransactOpts, _beneficiary)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) ClaimRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "claimRefund")
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) ClaimRefund() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.ClaimRefund(&_RefundableCrowdsale.TransactOpts)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) ClaimRefund() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.ClaimRefund(&_RefundableCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Finalize(&_RefundableCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Finalize(&_RefundableCrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.TransferOwnership(&_RefundableCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.TransferOwnership(&_RefundableCrowdsale.TransactOpts, newOwner)
}
