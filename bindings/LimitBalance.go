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

// LimitBalanceABI is the input ABI used to generate the binding from.
const LimitBalanceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// LimitBalanceBin is the compiled bytecode used for deploying new contracts.
const LimitBalanceBin = `6060604052341561000f57600080fd5b6040516020806100cc8339810160405280805160005550506097806100356000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663a4d66daf81146043575b600080fd5b3415604d57600080fd5b60536065565b60405190815260200160405180910390f35b600054815600a165627a7a72305820f887a5ffff3cf6016b575ed193478abe7db653bd1f76b7ece86812ccef4592740029`

// DeployLimitBalance deploys a new Ethereum contract, binding an instance of LimitBalance to it.
func DeployLimitBalance(auth *bind.TransactOpts, backend bind.ContractBackend, _limit *big.Int) (common.Address, *types.Transaction, *LimitBalance, error) {
	parsed, err := abi.JSON(strings.NewReader(LimitBalanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LimitBalanceBin), backend, _limit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LimitBalance{LimitBalanceCaller: LimitBalanceCaller{contract: contract}, LimitBalanceTransactor: LimitBalanceTransactor{contract: contract}}, nil
}

// LimitBalance is an auto generated Go binding around an Ethereum contract.
type LimitBalance struct {
	LimitBalanceCaller     // Read-only binding to the contract
	LimitBalanceTransactor // Write-only binding to the contract
}

// LimitBalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitBalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitBalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitBalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitBalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitBalanceSession struct {
	Contract     *LimitBalance     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LimitBalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitBalanceCallerSession struct {
	Contract *LimitBalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LimitBalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitBalanceTransactorSession struct {
	Contract     *LimitBalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LimitBalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitBalanceRaw struct {
	Contract *LimitBalance // Generic contract binding to access the raw methods on
}

// LimitBalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitBalanceCallerRaw struct {
	Contract *LimitBalanceCaller // Generic read-only contract binding to access the raw methods on
}

// LimitBalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitBalanceTransactorRaw struct {
	Contract *LimitBalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitBalance creates a new instance of LimitBalance, bound to a specific deployed contract.
func NewLimitBalance(address common.Address, backend bind.ContractBackend) (*LimitBalance, error) {
	contract, err := bindLimitBalance(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitBalance{LimitBalanceCaller: LimitBalanceCaller{contract: contract}, LimitBalanceTransactor: LimitBalanceTransactor{contract: contract}}, nil
}

// NewLimitBalanceCaller creates a new read-only instance of LimitBalance, bound to a specific deployed contract.
func NewLimitBalanceCaller(address common.Address, caller bind.ContractCaller) (*LimitBalanceCaller, error) {
	contract, err := bindLimitBalance(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &LimitBalanceCaller{contract: contract}, nil
}

// NewLimitBalanceTransactor creates a new write-only instance of LimitBalance, bound to a specific deployed contract.
func NewLimitBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitBalanceTransactor, error) {
	contract, err := bindLimitBalance(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &LimitBalanceTransactor{contract: contract}, nil
}

// bindLimitBalance binds a generic wrapper to an already deployed contract.
func bindLimitBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LimitBalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitBalance *LimitBalanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LimitBalance.Contract.LimitBalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitBalance *LimitBalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitBalance.Contract.LimitBalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitBalance *LimitBalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitBalance.Contract.LimitBalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitBalance *LimitBalanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LimitBalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitBalance *LimitBalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitBalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitBalance *LimitBalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitBalance.Contract.contract.Transact(opts, method, params...)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint256)
func (_LimitBalance *LimitBalanceCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LimitBalance.contract.Call(opts, out, "limit")
	return *ret0, err
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint256)
func (_LimitBalance *LimitBalanceSession) Limit() (*big.Int, error) {
	return _LimitBalance.Contract.Limit(&_LimitBalance.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint256)
func (_LimitBalance *LimitBalanceCallerSession) Limit() (*big.Int, error) {
	return _LimitBalance.Contract.Limit(&_LimitBalance.CallOpts)
}
