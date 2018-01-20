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

// ERC20SucceedingMockABI is the input ABI used to generate the binding from.
const ERC20SucceedingMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20SucceedingMockBin is the compiled bytecode used for deploying new contracts.
const ERC20SucceedingMockBin = `6060604052341561000f57600080fd5b6101c88061001e6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100bf57806323b872dd146100e457806370a0823114610119578063a9059cbb1461007c578063dd62ed3e14610145575b600080fd5b341561008757600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043516602435610177565b604051901515815260200160405180910390f35b34156100ca57600080fd5b6100d261017f565b60405190815260200160405180910390f35b34156100ef57600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610185565b341561012457600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff6004351661018e565b341561015057600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff60043581169060243516610194565b600192915050565b60005481565b60019392505050565b50600090565b6000929150505600a165627a7a7230582088e502b8569533604a725ed3712f8722558e2ab624003b6e87aa2f25ae631f370029`

// DeployERC20SucceedingMock deploys a new Ethereum contract, binding an instance of ERC20SucceedingMock to it.
func DeployERC20SucceedingMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20SucceedingMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SucceedingMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20SucceedingMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20SucceedingMock{ERC20SucceedingMockCaller: ERC20SucceedingMockCaller{contract: contract}, ERC20SucceedingMockTransactor: ERC20SucceedingMockTransactor{contract: contract}}, nil
}

// ERC20SucceedingMock is an auto generated Go binding around an Ethereum contract.
type ERC20SucceedingMock struct {
	ERC20SucceedingMockCaller     // Read-only binding to the contract
	ERC20SucceedingMockTransactor // Write-only binding to the contract
}

// ERC20SucceedingMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SucceedingMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SucceedingMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SucceedingMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SucceedingMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SucceedingMockSession struct {
	Contract     *ERC20SucceedingMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20SucceedingMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SucceedingMockCallerSession struct {
	Contract *ERC20SucceedingMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ERC20SucceedingMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SucceedingMockTransactorSession struct {
	Contract     *ERC20SucceedingMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ERC20SucceedingMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SucceedingMockRaw struct {
	Contract *ERC20SucceedingMock // Generic contract binding to access the raw methods on
}

// ERC20SucceedingMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SucceedingMockCallerRaw struct {
	Contract *ERC20SucceedingMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SucceedingMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SucceedingMockTransactorRaw struct {
	Contract *ERC20SucceedingMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20SucceedingMock creates a new instance of ERC20SucceedingMock, bound to a specific deployed contract.
func NewERC20SucceedingMock(address common.Address, backend bind.ContractBackend) (*ERC20SucceedingMock, error) {
	contract, err := bindERC20SucceedingMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20SucceedingMock{ERC20SucceedingMockCaller: ERC20SucceedingMockCaller{contract: contract}, ERC20SucceedingMockTransactor: ERC20SucceedingMockTransactor{contract: contract}}, nil
}

// NewERC20SucceedingMockCaller creates a new read-only instance of ERC20SucceedingMock, bound to a specific deployed contract.
func NewERC20SucceedingMockCaller(address common.Address, caller bind.ContractCaller) (*ERC20SucceedingMockCaller, error) {
	contract, err := bindERC20SucceedingMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SucceedingMockCaller{contract: contract}, nil
}

// NewERC20SucceedingMockTransactor creates a new write-only instance of ERC20SucceedingMock, bound to a specific deployed contract.
func NewERC20SucceedingMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SucceedingMockTransactor, error) {
	contract, err := bindERC20SucceedingMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20SucceedingMockTransactor{contract: contract}, nil
}

// bindERC20SucceedingMock binds a generic wrapper to an already deployed contract.
func bindERC20SucceedingMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SucceedingMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20SucceedingMock *ERC20SucceedingMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20SucceedingMock.Contract.ERC20SucceedingMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20SucceedingMock *ERC20SucceedingMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.ERC20SucceedingMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20SucceedingMock *ERC20SucceedingMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.ERC20SucceedingMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20SucceedingMock *ERC20SucceedingMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20SucceedingMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20SucceedingMock.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20SucceedingMock.Contract.Allowance(&_ERC20SucceedingMock.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20SucceedingMock.Contract.Allowance(&_ERC20SucceedingMock.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20SucceedingMock.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20SucceedingMock.Contract.BalanceOf(&_ERC20SucceedingMock.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20SucceedingMock.Contract.BalanceOf(&_ERC20SucceedingMock.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20SucceedingMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockSession) TotalSupply() (*big.Int, error) {
	return _ERC20SucceedingMock.Contract.TotalSupply(&_ERC20SucceedingMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20SucceedingMock *ERC20SucceedingMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20SucceedingMock.Contract.TotalSupply(&_ERC20SucceedingMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.Approve(&_ERC20SucceedingMock.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.Approve(&_ERC20SucceedingMock.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.Transfer(&_ERC20SucceedingMock.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.Transfer(&_ERC20SucceedingMock.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.TransferFrom(&_ERC20SucceedingMock.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20SucceedingMock *ERC20SucceedingMockTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20SucceedingMock.Contract.TransferFrom(&_ERC20SucceedingMock.TransactOpts, arg0, arg1, arg2)
}
