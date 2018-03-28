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

// DetailedERC20ABI is the input ABI used to generate the binding from.
const DetailedERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// DetailedERC20 is an auto generated Go binding around an Ethereum contract.
type DetailedERC20 struct {
	DetailedERC20Caller     // Read-only binding to the contract
	DetailedERC20Transactor // Write-only binding to the contract
}

// DetailedERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type DetailedERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DetailedERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DetailedERC20Session struct {
	Contract     *DetailedERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DetailedERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DetailedERC20CallerSession struct {
	Contract *DetailedERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DetailedERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DetailedERC20TransactorSession struct {
	Contract     *DetailedERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DetailedERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type DetailedERC20Raw struct {
	Contract *DetailedERC20 // Generic contract binding to access the raw methods on
}

// DetailedERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DetailedERC20CallerRaw struct {
	Contract *DetailedERC20Caller // Generic read-only contract binding to access the raw methods on
}

// DetailedERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DetailedERC20TransactorRaw struct {
	Contract *DetailedERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDetailedERC20 creates a new instance of DetailedERC20, bound to a specific deployed contract.
func NewDetailedERC20(address common.Address, backend bind.ContractBackend) (*DetailedERC20, error) {
	contract, err := bindDetailedERC20(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DetailedERC20{DetailedERC20Caller: DetailedERC20Caller{contract: contract}, DetailedERC20Transactor: DetailedERC20Transactor{contract: contract}}, nil
}

// NewDetailedERC20Caller creates a new read-only instance of DetailedERC20, bound to a specific deployed contract.
func NewDetailedERC20Caller(address common.Address, caller bind.ContractCaller) (*DetailedERC20Caller, error) {
	contract, err := bindDetailedERC20(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DetailedERC20Caller{contract: contract}, nil
}

// NewDetailedERC20Transactor creates a new write-only instance of DetailedERC20, bound to a specific deployed contract.
func NewDetailedERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*DetailedERC20Transactor, error) {
	contract, err := bindDetailedERC20(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DetailedERC20Transactor{contract: contract}, nil
}

// bindDetailedERC20 binds a generic wrapper to an already deployed contract.
func bindDetailedERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DetailedERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DetailedERC20 *DetailedERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DetailedERC20.Contract.DetailedERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DetailedERC20 *DetailedERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DetailedERC20.Contract.DetailedERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DetailedERC20 *DetailedERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DetailedERC20.Contract.DetailedERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DetailedERC20 *DetailedERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DetailedERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DetailedERC20 *DetailedERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DetailedERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DetailedERC20 *DetailedERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DetailedERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_DetailedERC20 *DetailedERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_DetailedERC20 *DetailedERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DetailedERC20.Contract.Allowance(&_DetailedERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_DetailedERC20 *DetailedERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DetailedERC20.Contract.Allowance(&_DetailedERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_DetailedERC20 *DetailedERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_DetailedERC20 *DetailedERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _DetailedERC20.Contract.BalanceOf(&_DetailedERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_DetailedERC20 *DetailedERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _DetailedERC20.Contract.BalanceOf(&_DetailedERC20.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedERC20 *DetailedERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DetailedERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedERC20 *DetailedERC20Session) Decimals() (uint8, error) {
	return _DetailedERC20.Contract.Decimals(&_DetailedERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedERC20 *DetailedERC20CallerSession) Decimals() (uint8, error) {
	return _DetailedERC20.Contract.Decimals(&_DetailedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedERC20 *DetailedERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DetailedERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedERC20 *DetailedERC20Session) Name() (string, error) {
	return _DetailedERC20.Contract.Name(&_DetailedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedERC20 *DetailedERC20CallerSession) Name() (string, error) {
	return _DetailedERC20.Contract.Name(&_DetailedERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedERC20 *DetailedERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DetailedERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedERC20 *DetailedERC20Session) Symbol() (string, error) {
	return _DetailedERC20.Contract.Symbol(&_DetailedERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedERC20 *DetailedERC20CallerSession) Symbol() (string, error) {
	return _DetailedERC20.Contract.Symbol(&_DetailedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedERC20 *DetailedERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedERC20 *DetailedERC20Session) TotalSupply() (*big.Int, error) {
	return _DetailedERC20.Contract.TotalSupply(&_DetailedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedERC20 *DetailedERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _DetailedERC20.Contract.TotalSupply(&_DetailedERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.Contract.Approve(&_DetailedERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.Contract.Approve(&_DetailedERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.Contract.Transfer(&_DetailedERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.Contract.Transfer(&_DetailedERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.Contract.TransferFrom(&_DetailedERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_DetailedERC20 *DetailedERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20.Contract.TransferFrom(&_DetailedERC20.TransactOpts, from, to, value)
}
