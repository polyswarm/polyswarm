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

// SecureTargetMockABI is the input ABI used to generate the binding from.
const SecureTargetMockABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"checkInvariant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SecureTargetMockBin is the compiled bytecode used for deploying new contracts.
const SecureTargetMockBin = `60606040523415600e57600080fd5b60988061001c6000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e79487da81146043575b600080fd5b3415604d57600080fd5b60536067565b604051901515815260200160405180910390f35b6001905600a165627a7a7230582043ff927b8781c8275a83b47e8b694bae4dbf172d2d4d1a467dcd93cde627ca9e0029`

// DeploySecureTargetMock deploys a new Ethereum contract, binding an instance of SecureTargetMock to it.
func DeploySecureTargetMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SecureTargetMock, error) {
	parsed, err := abi.JSON(strings.NewReader(SecureTargetMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SecureTargetMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecureTargetMock{SecureTargetMockCaller: SecureTargetMockCaller{contract: contract}, SecureTargetMockTransactor: SecureTargetMockTransactor{contract: contract}}, nil
}

// SecureTargetMock is an auto generated Go binding around an Ethereum contract.
type SecureTargetMock struct {
	SecureTargetMockCaller     // Read-only binding to the contract
	SecureTargetMockTransactor // Write-only binding to the contract
}

// SecureTargetMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecureTargetMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureTargetMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecureTargetMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureTargetMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecureTargetMockSession struct {
	Contract     *SecureTargetMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecureTargetMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecureTargetMockCallerSession struct {
	Contract *SecureTargetMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SecureTargetMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecureTargetMockTransactorSession struct {
	Contract     *SecureTargetMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SecureTargetMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecureTargetMockRaw struct {
	Contract *SecureTargetMock // Generic contract binding to access the raw methods on
}

// SecureTargetMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecureTargetMockCallerRaw struct {
	Contract *SecureTargetMockCaller // Generic read-only contract binding to access the raw methods on
}

// SecureTargetMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecureTargetMockTransactorRaw struct {
	Contract *SecureTargetMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecureTargetMock creates a new instance of SecureTargetMock, bound to a specific deployed contract.
func NewSecureTargetMock(address common.Address, backend bind.ContractBackend) (*SecureTargetMock, error) {
	contract, err := bindSecureTargetMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecureTargetMock{SecureTargetMockCaller: SecureTargetMockCaller{contract: contract}, SecureTargetMockTransactor: SecureTargetMockTransactor{contract: contract}}, nil
}

// NewSecureTargetMockCaller creates a new read-only instance of SecureTargetMock, bound to a specific deployed contract.
func NewSecureTargetMockCaller(address common.Address, caller bind.ContractCaller) (*SecureTargetMockCaller, error) {
	contract, err := bindSecureTargetMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SecureTargetMockCaller{contract: contract}, nil
}

// NewSecureTargetMockTransactor creates a new write-only instance of SecureTargetMock, bound to a specific deployed contract.
func NewSecureTargetMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SecureTargetMockTransactor, error) {
	contract, err := bindSecureTargetMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SecureTargetMockTransactor{contract: contract}, nil
}

// bindSecureTargetMock binds a generic wrapper to an already deployed contract.
func bindSecureTargetMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SecureTargetMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecureTargetMock *SecureTargetMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecureTargetMock.Contract.SecureTargetMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecureTargetMock *SecureTargetMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureTargetMock.Contract.SecureTargetMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecureTargetMock *SecureTargetMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecureTargetMock.Contract.SecureTargetMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecureTargetMock *SecureTargetMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecureTargetMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecureTargetMock *SecureTargetMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureTargetMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecureTargetMock *SecureTargetMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecureTargetMock.Contract.contract.Transact(opts, method, params...)
}

// CheckInvariant is a paid mutator transaction binding the contract method 0xe79487da.
//
// Solidity: function checkInvariant() returns(bool)
func (_SecureTargetMock *SecureTargetMockTransactor) CheckInvariant(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureTargetMock.contract.Transact(opts, "checkInvariant")
}

// CheckInvariant is a paid mutator transaction binding the contract method 0xe79487da.
//
// Solidity: function checkInvariant() returns(bool)
func (_SecureTargetMock *SecureTargetMockSession) CheckInvariant() (*types.Transaction, error) {
	return _SecureTargetMock.Contract.CheckInvariant(&_SecureTargetMock.TransactOpts)
}

// CheckInvariant is a paid mutator transaction binding the contract method 0xe79487da.
//
// Solidity: function checkInvariant() returns(bool)
func (_SecureTargetMock *SecureTargetMockTransactorSession) CheckInvariant() (*types.Transaction, error) {
	return _SecureTargetMock.Contract.CheckInvariant(&_SecureTargetMock.TransactOpts)
}
