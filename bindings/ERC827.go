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

// ERC827ABI is the input ABI used to generate the binding from.
const ERC827ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC827 is an auto generated Go binding around an Ethereum contract.
type ERC827 struct {
	ERC827Caller     // Read-only binding to the contract
	ERC827Transactor // Write-only binding to the contract
}

// ERC827Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC827Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC827Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC827Session struct {
	Contract     *ERC827           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC827CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC827CallerSession struct {
	Contract *ERC827Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC827TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC827TransactorSession struct {
	Contract     *ERC827Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC827Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC827Raw struct {
	Contract *ERC827 // Generic contract binding to access the raw methods on
}

// ERC827CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC827CallerRaw struct {
	Contract *ERC827Caller // Generic read-only contract binding to access the raw methods on
}

// ERC827TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC827TransactorRaw struct {
	Contract *ERC827Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC827 creates a new instance of ERC827, bound to a specific deployed contract.
func NewERC827(address common.Address, backend bind.ContractBackend) (*ERC827, error) {
	contract, err := bindERC827(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC827{ERC827Caller: ERC827Caller{contract: contract}, ERC827Transactor: ERC827Transactor{contract: contract}}, nil
}

// NewERC827Caller creates a new read-only instance of ERC827, bound to a specific deployed contract.
func NewERC827Caller(address common.Address, caller bind.ContractCaller) (*ERC827Caller, error) {
	contract, err := bindERC827(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC827Caller{contract: contract}, nil
}

// NewERC827Transactor creates a new write-only instance of ERC827, bound to a specific deployed contract.
func NewERC827Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC827Transactor, error) {
	contract, err := bindERC827(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC827Transactor{contract: contract}, nil
}

// bindERC827 binds a generic wrapper to an already deployed contract.
func bindERC827(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC827ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC827 *ERC827Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC827.Contract.ERC827Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC827 *ERC827Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC827.Contract.ERC827Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC827 *ERC827Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC827.Contract.ERC827Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC827 *ERC827CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC827.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC827 *ERC827TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC827.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC827 *ERC827TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC827.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC827 *ERC827Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC827 *ERC827Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC827.Contract.Allowance(&_ERC827.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC827 *ERC827CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC827.Contract.Allowance(&_ERC827.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC827 *ERC827Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC827 *ERC827Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC827.Contract.BalanceOf(&_ERC827.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC827 *ERC827CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC827.Contract.BalanceOf(&_ERC827.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827 *ERC827Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827 *ERC827Session) TotalSupply() (*big.Int, error) {
	return _ERC827.Contract.TotalSupply(&_ERC827.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827 *ERC827CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC827.Contract.TotalSupply(&_ERC827.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.contract.Transact(opts, "approve", _spender, _value, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827Session) Approve(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.Contract.Approve(&_ERC827.TransactOpts, _spender, _value, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827TransactorSession) Approve(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.Contract.Approve(&_ERC827.TransactOpts, _spender, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827Session) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.Contract.Transfer(&_ERC827.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827TransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.Contract.Transfer(&_ERC827.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.contract.Transact(opts, "transferFrom", _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.Contract.TransferFrom(&_ERC827.TransactOpts, _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827 *ERC827TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827.Contract.TransferFrom(&_ERC827.TransactOpts, _from, _to, _value, _data)
}
