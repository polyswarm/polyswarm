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

// SafeMathMockABI is the input ABI used to generate the binding from.
const SafeMathMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"multiply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"subtract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"result\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SafeMathMockBin is the compiled bytecode used for deploying new contracts.
const SafeMathMockBin = `6060604052341561000f57600080fd5b6101868061001e6000396000f3006060604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663165c4a1681146100665780633ef5e44514610081578063653721471461009a578063771602f7146100bf575b600080fd5b341561007157600080fd5b61007f6004356024356100d8565b005b341561008c57600080fd5b61007f6004356024356100e9565b34156100a557600080fd5b6100ad6100f3565b60405190815260200160405180910390f35b34156100ca57600080fd5b61007f6004356024356100f9565b6100e28282610103565b6000555050565b6100e28282610139565b60005481565b6100e2828261014b565b6000808315156101165760009150610132565b5082820282848281151561012657fe5b041461012e57fe5b8091505b5092915050565b60008282111561014557fe5b50900390565b60008282018381101561012e57fe00a165627a7a7230582014ceccff73051237d687a6835c09cfe3f463725177cd261c1755f8f636eb315f0029`

// DeploySafeMathMock deploys a new Ethereum contract, binding an instance of SafeMathMock to it.
func DeploySafeMathMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathMock, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathMock{SafeMathMockCaller: SafeMathMockCaller{contract: contract}, SafeMathMockTransactor: SafeMathMockTransactor{contract: contract}}, nil
}

// SafeMathMock is an auto generated Go binding around an Ethereum contract.
type SafeMathMock struct {
	SafeMathMockCaller     // Read-only binding to the contract
	SafeMathMockTransactor // Write-only binding to the contract
}

// SafeMathMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathMockSession struct {
	Contract     *SafeMathMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathMockCallerSession struct {
	Contract *SafeMathMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SafeMathMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathMockTransactorSession struct {
	Contract     *SafeMathMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SafeMathMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathMockRaw struct {
	Contract *SafeMathMock // Generic contract binding to access the raw methods on
}

// SafeMathMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathMockCallerRaw struct {
	Contract *SafeMathMockCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathMockTransactorRaw struct {
	Contract *SafeMathMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathMock creates a new instance of SafeMathMock, bound to a specific deployed contract.
func NewSafeMathMock(address common.Address, backend bind.ContractBackend) (*SafeMathMock, error) {
	contract, err := bindSafeMathMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathMock{SafeMathMockCaller: SafeMathMockCaller{contract: contract}, SafeMathMockTransactor: SafeMathMockTransactor{contract: contract}}, nil
}

// NewSafeMathMockCaller creates a new read-only instance of SafeMathMock, bound to a specific deployed contract.
func NewSafeMathMockCaller(address common.Address, caller bind.ContractCaller) (*SafeMathMockCaller, error) {
	contract, err := bindSafeMathMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathMockCaller{contract: contract}, nil
}

// NewSafeMathMockTransactor creates a new write-only instance of SafeMathMock, bound to a specific deployed contract.
func NewSafeMathMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathMockTransactor, error) {
	contract, err := bindSafeMathMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SafeMathMockTransactor{contract: contract}, nil
}

// bindSafeMathMock binds a generic wrapper to an already deployed contract.
func bindSafeMathMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathMock *SafeMathMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathMock.Contract.SafeMathMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathMock *SafeMathMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathMock.Contract.SafeMathMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathMock *SafeMathMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathMock.Contract.SafeMathMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathMock *SafeMathMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMathMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathMock *SafeMathMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathMock *SafeMathMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathMock.Contract.contract.Transact(opts, method, params...)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint256)
func (_SafeMathMock *SafeMathMockCaller) Result(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMathMock.contract.Call(opts, out, "result")
	return *ret0, err
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint256)
func (_SafeMathMock *SafeMathMockSession) Result() (*big.Int, error) {
	return _SafeMathMock.Contract.Result(&_SafeMathMock.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint256)
func (_SafeMathMock *SafeMathMockCallerSession) Result() (*big.Int, error) {
	return _SafeMathMock.Contract.Result(&_SafeMathMock.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockTransactor) Add(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.contract.Transact(opts, "add", a, b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.Contract.Add(&_SafeMathMock.TransactOpts, a, b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockTransactorSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.Contract.Add(&_SafeMathMock.TransactOpts, a, b)
}

// Multiply is a paid mutator transaction binding the contract method 0x165c4a16.
//
// Solidity: function multiply(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockTransactor) Multiply(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.contract.Transact(opts, "multiply", a, b)
}

// Multiply is a paid mutator transaction binding the contract method 0x165c4a16.
//
// Solidity: function multiply(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockSession) Multiply(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.Contract.Multiply(&_SafeMathMock.TransactOpts, a, b)
}

// Multiply is a paid mutator transaction binding the contract method 0x165c4a16.
//
// Solidity: function multiply(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockTransactorSession) Multiply(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.Contract.Multiply(&_SafeMathMock.TransactOpts, a, b)
}

// Subtract is a paid mutator transaction binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockTransactor) Subtract(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.contract.Transact(opts, "subtract", a, b)
}

// Subtract is a paid mutator transaction binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockSession) Subtract(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.Contract.Subtract(&_SafeMathMock.TransactOpts, a, b)
}

// Subtract is a paid mutator transaction binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(a uint256, b uint256) returns()
func (_SafeMathMock *SafeMathMockTransactorSession) Subtract(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _SafeMathMock.Contract.Subtract(&_SafeMathMock.TransactOpts, a, b)
}
