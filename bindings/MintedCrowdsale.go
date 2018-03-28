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

// MintedCrowdsaleABI is the input ABI used to generate the binding from.
const MintedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// MintedCrowdsale is an auto generated Go binding around an Ethereum contract.
type MintedCrowdsale struct {
	MintedCrowdsaleCaller     // Read-only binding to the contract
	MintedCrowdsaleTransactor // Write-only binding to the contract
}

// MintedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintedCrowdsaleSession struct {
	Contract     *MintedCrowdsale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintedCrowdsaleCallerSession struct {
	Contract *MintedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MintedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintedCrowdsaleTransactorSession struct {
	Contract     *MintedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MintedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintedCrowdsaleRaw struct {
	Contract *MintedCrowdsale // Generic contract binding to access the raw methods on
}

// MintedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintedCrowdsaleCallerRaw struct {
	Contract *MintedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// MintedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintedCrowdsaleTransactorRaw struct {
	Contract *MintedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintedCrowdsale creates a new instance of MintedCrowdsale, bound to a specific deployed contract.
func NewMintedCrowdsale(address common.Address, backend bind.ContractBackend) (*MintedCrowdsale, error) {
	contract, err := bindMintedCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsale{MintedCrowdsaleCaller: MintedCrowdsaleCaller{contract: contract}, MintedCrowdsaleTransactor: MintedCrowdsaleTransactor{contract: contract}}, nil
}

// NewMintedCrowdsaleCaller creates a new read-only instance of MintedCrowdsale, bound to a specific deployed contract.
func NewMintedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*MintedCrowdsaleCaller, error) {
	contract, err := bindMintedCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleCaller{contract: contract}, nil
}

// NewMintedCrowdsaleTransactor creates a new write-only instance of MintedCrowdsale, bound to a specific deployed contract.
func NewMintedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*MintedCrowdsaleTransactor, error) {
	contract, err := bindMintedCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleTransactor{contract: contract}, nil
}

// bindMintedCrowdsale binds a generic wrapper to an already deployed contract.
func bindMintedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintedCrowdsale *MintedCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintedCrowdsale.Contract.MintedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintedCrowdsale *MintedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintedCrowdsale.Contract.MintedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintedCrowdsale *MintedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintedCrowdsale.Contract.MintedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintedCrowdsale *MintedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintedCrowdsale *MintedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintedCrowdsale *MintedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_MintedCrowdsale *MintedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintedCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_MintedCrowdsale *MintedCrowdsaleSession) Rate() (*big.Int, error) {
	return _MintedCrowdsale.Contract.Rate(&_MintedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_MintedCrowdsale *MintedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _MintedCrowdsale.Contract.Rate(&_MintedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MintedCrowdsale *MintedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintedCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MintedCrowdsale *MintedCrowdsaleSession) Token() (common.Address, error) {
	return _MintedCrowdsale.Contract.Token(&_MintedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MintedCrowdsale *MintedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _MintedCrowdsale.Contract.Token(&_MintedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_MintedCrowdsale *MintedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintedCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_MintedCrowdsale *MintedCrowdsaleSession) Wallet() (common.Address, error) {
	return _MintedCrowdsale.Contract.Wallet(&_MintedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_MintedCrowdsale *MintedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _MintedCrowdsale.Contract.Wallet(&_MintedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_MintedCrowdsale *MintedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintedCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_MintedCrowdsale *MintedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _MintedCrowdsale.Contract.WeiRaised(&_MintedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_MintedCrowdsale *MintedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _MintedCrowdsale.Contract.WeiRaised(&_MintedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_MintedCrowdsale *MintedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _MintedCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_MintedCrowdsale *MintedCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _MintedCrowdsale.Contract.BuyTokens(&_MintedCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_MintedCrowdsale *MintedCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _MintedCrowdsale.Contract.BuyTokens(&_MintedCrowdsale.TransactOpts, _beneficiary)
}
