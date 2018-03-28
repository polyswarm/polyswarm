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

// ERC223TokenMockABI is the input ABI used to generate the binding from.
const ERC223TokenMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferERC223\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC223TokenMockBin is the compiled bytecode used for deploying new contracts.
const ERC223TokenMockBin = `6060604052341561000f57600080fd5b6040516040806104458339810160405280805191906020018051600160a060020a03909316600090815260208190526040902083905550506001556103ec806100596000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166303bad56e811461006657806318160ddd146100df57806370a0823114610104578063a9059cbb14610123575b600080fd5b341561007157600080fd5b6100cb60048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061014595505050505050565b604051901515815260200160405180910390f35b34156100ea57600080fd5b6100f2610265565b60405190815260200160405180910390f35b341561010f57600080fd5b6100f2600160a060020a036004351661026b565b341561012e57600080fd5b6100cb600160a060020a0360043516602435610286565b60008060006101548686610286565b50853b151991508115610259575084600160a060020a03811663c0ee0b8a3387876040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101fb5780820151838201526020016101e3565b50505050905090810190601f1680156102285780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b151561024857600080fd5b5af1151561025557600080fd5b5050505b50600195945050505050565b60015490565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a038316151561029d57600080fd5b600160a060020a0333166000908152602081905260409020548211156102c257600080fd5b600160a060020a0333166000908152602081905260409020546102eb908363ffffffff61039816565b600160a060020a033381166000908152602081905260408082209390935590851681522054610320908363ffffffff6103aa16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b6000828211156103a457fe5b50900390565b6000828201838110156103b957fe5b93925050505600a165627a7a72305820e674dc1763179755ae7e19c6c113f7ffffc89e2e0906b6318d3749e6bd5b53e50029`

// DeployERC223TokenMock deploys a new Ethereum contract, binding an instance of ERC223TokenMock to it.
func DeployERC223TokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *ERC223TokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223TokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC223TokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC223TokenMock{ERC223TokenMockCaller: ERC223TokenMockCaller{contract: contract}, ERC223TokenMockTransactor: ERC223TokenMockTransactor{contract: contract}}, nil
}

// ERC223TokenMock is an auto generated Go binding around an Ethereum contract.
type ERC223TokenMock struct {
	ERC223TokenMockCaller     // Read-only binding to the contract
	ERC223TokenMockTransactor // Write-only binding to the contract
}

// ERC223TokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC223TokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC223TokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC223TokenMockSession struct {
	Contract     *ERC223TokenMock  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223TokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC223TokenMockCallerSession struct {
	Contract *ERC223TokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC223TokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC223TokenMockTransactorSession struct {
	Contract     *ERC223TokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC223TokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC223TokenMockRaw struct {
	Contract *ERC223TokenMock // Generic contract binding to access the raw methods on
}

// ERC223TokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC223TokenMockCallerRaw struct {
	Contract *ERC223TokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC223TokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC223TokenMockTransactorRaw struct {
	Contract *ERC223TokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC223TokenMock creates a new instance of ERC223TokenMock, bound to a specific deployed contract.
func NewERC223TokenMock(address common.Address, backend bind.ContractBackend) (*ERC223TokenMock, error) {
	contract, err := bindERC223TokenMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenMock{ERC223TokenMockCaller: ERC223TokenMockCaller{contract: contract}, ERC223TokenMockTransactor: ERC223TokenMockTransactor{contract: contract}}, nil
}

// NewERC223TokenMockCaller creates a new read-only instance of ERC223TokenMock, bound to a specific deployed contract.
func NewERC223TokenMockCaller(address common.Address, caller bind.ContractCaller) (*ERC223TokenMockCaller, error) {
	contract, err := bindERC223TokenMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenMockCaller{contract: contract}, nil
}

// NewERC223TokenMockTransactor creates a new write-only instance of ERC223TokenMock, bound to a specific deployed contract.
func NewERC223TokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC223TokenMockTransactor, error) {
	contract, err := bindERC223TokenMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenMockTransactor{contract: contract}, nil
}

// bindERC223TokenMock binds a generic wrapper to an already deployed contract.
func bindERC223TokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223TokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223TokenMock *ERC223TokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223TokenMock.Contract.ERC223TokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223TokenMock *ERC223TokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.ERC223TokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223TokenMock *ERC223TokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.ERC223TokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223TokenMock *ERC223TokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223TokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223TokenMock *ERC223TokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223TokenMock *ERC223TokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC223TokenMock *ERC223TokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223TokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC223TokenMock *ERC223TokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC223TokenMock.Contract.BalanceOf(&_ERC223TokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC223TokenMock *ERC223TokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC223TokenMock.Contract.BalanceOf(&_ERC223TokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC223TokenMock *ERC223TokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223TokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC223TokenMock *ERC223TokenMockSession) TotalSupply() (*big.Int, error) {
	return _ERC223TokenMock.Contract.TotalSupply(&_ERC223TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC223TokenMock *ERC223TokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC223TokenMock.Contract.TotalSupply(&_ERC223TokenMock.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC223TokenMock *ERC223TokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC223TokenMock *ERC223TokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.Transfer(&_ERC223TokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC223TokenMock *ERC223TokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.Transfer(&_ERC223TokenMock.TransactOpts, _to, _value)
}

// TransferERC223 is a paid mutator transaction binding the contract method 0x03bad56e.
//
// Solidity: function transferERC223(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC223TokenMock *ERC223TokenMockTransactor) TransferERC223(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223TokenMock.contract.Transact(opts, "transferERC223", _to, _value, _data)
}

// TransferERC223 is a paid mutator transaction binding the contract method 0x03bad56e.
//
// Solidity: function transferERC223(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC223TokenMock *ERC223TokenMockSession) TransferERC223(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.TransferERC223(&_ERC223TokenMock.TransactOpts, _to, _value, _data)
}

// TransferERC223 is a paid mutator transaction binding the contract method 0x03bad56e.
//
// Solidity: function transferERC223(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC223TokenMock *ERC223TokenMockTransactorSession) TransferERC223(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC223TokenMock.Contract.TransferERC223(&_ERC223TokenMock.TransactOpts, _to, _value, _data)
}
