// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// InsecureTargetMockABI is the input ABI used to generate the binding from.
const InsecureTargetMockABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"checkInvariant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InsecureTargetMockBin is the compiled bytecode used for deploying new contracts.
const InsecureTargetMockBin = `60606040523415600e57600080fd5b60988061001c6000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e79487da81146043575b600080fd5b3415604d57600080fd5b60536067565b604051901515815260200160405180910390f35b6000905600a165627a7a72305820826d9e56e0f1136fe0ff3b87a16c39877b3ed674bc385072fa56e1da4854cd650029`

// DeployInsecureTargetMock deploys a new Ethereum contract, binding an instance of InsecureTargetMock to it.
func DeployInsecureTargetMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InsecureTargetMock, error) {
	parsed, err := abi.JSON(strings.NewReader(InsecureTargetMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InsecureTargetMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InsecureTargetMock{InsecureTargetMockCaller: InsecureTargetMockCaller{contract: contract}, InsecureTargetMockTransactor: InsecureTargetMockTransactor{contract: contract}, InsecureTargetMockFilterer: InsecureTargetMockFilterer{contract: contract}}, nil
}

// InsecureTargetMock is an auto generated Go binding around an Ethereum contract.
type InsecureTargetMock struct {
	InsecureTargetMockCaller     // Read-only binding to the contract
	InsecureTargetMockTransactor // Write-only binding to the contract
	InsecureTargetMockFilterer   // Log filterer for contract events
}

// InsecureTargetMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type InsecureTargetMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsecureTargetMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InsecureTargetMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsecureTargetMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InsecureTargetMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsecureTargetMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InsecureTargetMockSession struct {
	Contract     *InsecureTargetMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InsecureTargetMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InsecureTargetMockCallerSession struct {
	Contract *InsecureTargetMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InsecureTargetMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InsecureTargetMockTransactorSession struct {
	Contract     *InsecureTargetMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InsecureTargetMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type InsecureTargetMockRaw struct {
	Contract *InsecureTargetMock // Generic contract binding to access the raw methods on
}

// InsecureTargetMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InsecureTargetMockCallerRaw struct {
	Contract *InsecureTargetMockCaller // Generic read-only contract binding to access the raw methods on
}

// InsecureTargetMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InsecureTargetMockTransactorRaw struct {
	Contract *InsecureTargetMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInsecureTargetMock creates a new instance of InsecureTargetMock, bound to a specific deployed contract.
func NewInsecureTargetMock(address common.Address, backend bind.ContractBackend) (*InsecureTargetMock, error) {
	contract, err := bindInsecureTargetMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetMock{InsecureTargetMockCaller: InsecureTargetMockCaller{contract: contract}, InsecureTargetMockTransactor: InsecureTargetMockTransactor{contract: contract}, InsecureTargetMockFilterer: InsecureTargetMockFilterer{contract: contract}}, nil
}

// NewInsecureTargetMockCaller creates a new read-only instance of InsecureTargetMock, bound to a specific deployed contract.
func NewInsecureTargetMockCaller(address common.Address, caller bind.ContractCaller) (*InsecureTargetMockCaller, error) {
	contract, err := bindInsecureTargetMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetMockCaller{contract: contract}, nil
}

// NewInsecureTargetMockTransactor creates a new write-only instance of InsecureTargetMock, bound to a specific deployed contract.
func NewInsecureTargetMockTransactor(address common.Address, transactor bind.ContractTransactor) (*InsecureTargetMockTransactor, error) {
	contract, err := bindInsecureTargetMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetMockTransactor{contract: contract}, nil
}

// NewInsecureTargetMockFilterer creates a new log filterer instance of InsecureTargetMock, bound to a specific deployed contract.
func NewInsecureTargetMockFilterer(address common.Address, filterer bind.ContractFilterer) (*InsecureTargetMockFilterer, error) {
	contract, err := bindInsecureTargetMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetMockFilterer{contract: contract}, nil
}

// bindInsecureTargetMock binds a generic wrapper to an already deployed contract.
func bindInsecureTargetMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InsecureTargetMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsecureTargetMock *InsecureTargetMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InsecureTargetMock.Contract.InsecureTargetMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsecureTargetMock *InsecureTargetMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetMock.Contract.InsecureTargetMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsecureTargetMock *InsecureTargetMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsecureTargetMock.Contract.InsecureTargetMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsecureTargetMock *InsecureTargetMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InsecureTargetMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsecureTargetMock *InsecureTargetMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsecureTargetMock *InsecureTargetMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsecureTargetMock.Contract.contract.Transact(opts, method, params...)
}

// CheckInvariant is a paid mutator transaction binding the contract method 0xe79487da.
//
// Solidity: function checkInvariant() returns(bool)
func (_InsecureTargetMock *InsecureTargetMockTransactor) CheckInvariant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetMock.contract.Transact(opts, "checkInvariant")
}

// CheckInvariant is a paid mutator transaction binding the contract method 0xe79487da.
//
// Solidity: function checkInvariant() returns(bool)
func (_InsecureTargetMock *InsecureTargetMockSession) CheckInvariant() (*types.Transaction, error) {
	return _InsecureTargetMock.Contract.CheckInvariant(&_InsecureTargetMock.TransactOpts)
}

// CheckInvariant is a paid mutator transaction binding the contract method 0xe79487da.
//
// Solidity: function checkInvariant() returns(bool)
func (_InsecureTargetMock *InsecureTargetMockTransactorSession) CheckInvariant() (*types.Transaction, error) {
	return _InsecureTargetMock.Contract.CheckInvariant(&_InsecureTargetMock.TransactOpts)
}
