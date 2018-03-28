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

// ERC20FailingMockABI is the input ABI used to generate the binding from.
const ERC20FailingMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20FailingMockBin is the compiled bytecode used for deploying new contracts.
const ERC20FailingMockBin = `6060604052341561000f57600080fd5b6101bb8061001e6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100bf57806323b872dd146100e457806370a0823114610119578063a9059cbb1461007c578063dd62ed3e14610145575b600080fd5b341561008757600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043516602435610173565b604051901515815260200160405180910390f35b34156100ca57600080fd5b6100d261017b565b60405190815260200160405180910390f35b34156100ef57600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610180565b341561012457600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff60043516610189565b341561015057600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff600435811690602435165b600092915050565b600090565b60009392505050565b506000905600a165627a7a7230582053cacf9429ed406a357081ee312217ed85c81cfc3abc13f52510a4930c8afddb0029`

// DeployERC20FailingMock deploys a new Ethereum contract, binding an instance of ERC20FailingMock to it.
func DeployERC20FailingMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20FailingMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20FailingMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20FailingMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20FailingMock{ERC20FailingMockCaller: ERC20FailingMockCaller{contract: contract}, ERC20FailingMockTransactor: ERC20FailingMockTransactor{contract: contract}}, nil
}

// ERC20FailingMock is an auto generated Go binding around an Ethereum contract.
type ERC20FailingMock struct {
	ERC20FailingMockCaller     // Read-only binding to the contract
	ERC20FailingMockTransactor // Write-only binding to the contract
}

// ERC20FailingMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20FailingMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20FailingMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20FailingMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20FailingMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20FailingMockSession struct {
	Contract     *ERC20FailingMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20FailingMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20FailingMockCallerSession struct {
	Contract *ERC20FailingMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20FailingMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20FailingMockTransactorSession struct {
	Contract     *ERC20FailingMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20FailingMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20FailingMockRaw struct {
	Contract *ERC20FailingMock // Generic contract binding to access the raw methods on
}

// ERC20FailingMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20FailingMockCallerRaw struct {
	Contract *ERC20FailingMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20FailingMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20FailingMockTransactorRaw struct {
	Contract *ERC20FailingMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20FailingMock creates a new instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMock(address common.Address, backend bind.ContractBackend) (*ERC20FailingMock, error) {
	contract, err := bindERC20FailingMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMock{ERC20FailingMockCaller: ERC20FailingMockCaller{contract: contract}, ERC20FailingMockTransactor: ERC20FailingMockTransactor{contract: contract}}, nil
}

// NewERC20FailingMockCaller creates a new read-only instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMockCaller(address common.Address, caller bind.ContractCaller) (*ERC20FailingMockCaller, error) {
	contract, err := bindERC20FailingMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockCaller{contract: contract}, nil
}

// NewERC20FailingMockTransactor creates a new write-only instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20FailingMockTransactor, error) {
	contract, err := bindERC20FailingMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockTransactor{contract: contract}, nil
}

// bindERC20FailingMock binds a generic wrapper to an already deployed contract.
func bindERC20FailingMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20FailingMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20FailingMock *ERC20FailingMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20FailingMock.Contract.ERC20FailingMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20FailingMock *ERC20FailingMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.ERC20FailingMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20FailingMock *ERC20FailingMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.ERC20FailingMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20FailingMock *ERC20FailingMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20FailingMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20FailingMock *ERC20FailingMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20FailingMock *ERC20FailingMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20FailingMock.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.Allowance(&_ERC20FailingMock.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.Allowance(&_ERC20FailingMock.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20FailingMock.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.BalanceOf(&_ERC20FailingMock.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.BalanceOf(&_ERC20FailingMock.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20FailingMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockSession) TotalSupply() (*big.Int, error) {
	return _ERC20FailingMock.Contract.TotalSupply(&_ERC20FailingMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20FailingMock.Contract.TotalSupply(&_ERC20FailingMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Approve(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Approve(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Transfer(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Transfer(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.TransferFrom(&_ERC20FailingMock.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.TransferFrom(&_ERC20FailingMock.TransactOpts, arg0, arg1, arg2)
}
