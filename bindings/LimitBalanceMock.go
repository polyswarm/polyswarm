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

// LimitBalanceMockABI is the input ABI used to generate the binding from.
const LimitBalanceMockABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"limitedDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LimitBalanceMockBin is the compiled bytecode used for deploying new contracts.
const LimitBalanceMockBin = `60606040526103e860005560d0806100186000396000f30060606040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166386f3d0cd8114604d578063a4d66daf146055575b600080fd5b60536077565b005b3415605f57600080fd5b6065609e565b60405190815260200160405180910390f35b60005473ffffffffffffffffffffffffffffffffffffffff3016311115609c57600080fd5b565b600054815600a165627a7a72305820ab576179c39768f100ebf67c1af0ab3120689457336d90da5b3db4108c3cb8870029`

// DeployLimitBalanceMock deploys a new Ethereum contract, binding an instance of LimitBalanceMock to it.
func DeployLimitBalanceMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LimitBalanceMock, error) {
	parsed, err := abi.JSON(strings.NewReader(LimitBalanceMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LimitBalanceMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LimitBalanceMock{LimitBalanceMockCaller: LimitBalanceMockCaller{contract: contract}, LimitBalanceMockTransactor: LimitBalanceMockTransactor{contract: contract}, LimitBalanceMockFilterer: LimitBalanceMockFilterer{contract: contract}}, nil
}

// LimitBalanceMock is an auto generated Go binding around an Ethereum contract.
type LimitBalanceMock struct {
	LimitBalanceMockCaller     // Read-only binding to the contract
	LimitBalanceMockTransactor // Write-only binding to the contract
	LimitBalanceMockFilterer   // Log filterer for contract events
}

// LimitBalanceMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitBalanceMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitBalanceMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitBalanceMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitBalanceMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimitBalanceMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitBalanceMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitBalanceMockSession struct {
	Contract     *LimitBalanceMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LimitBalanceMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitBalanceMockCallerSession struct {
	Contract *LimitBalanceMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LimitBalanceMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitBalanceMockTransactorSession struct {
	Contract     *LimitBalanceMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LimitBalanceMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitBalanceMockRaw struct {
	Contract *LimitBalanceMock // Generic contract binding to access the raw methods on
}

// LimitBalanceMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitBalanceMockCallerRaw struct {
	Contract *LimitBalanceMockCaller // Generic read-only contract binding to access the raw methods on
}

// LimitBalanceMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitBalanceMockTransactorRaw struct {
	Contract *LimitBalanceMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitBalanceMock creates a new instance of LimitBalanceMock, bound to a specific deployed contract.
func NewLimitBalanceMock(address common.Address, backend bind.ContractBackend) (*LimitBalanceMock, error) {
	contract, err := bindLimitBalanceMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitBalanceMock{LimitBalanceMockCaller: LimitBalanceMockCaller{contract: contract}, LimitBalanceMockTransactor: LimitBalanceMockTransactor{contract: contract}, LimitBalanceMockFilterer: LimitBalanceMockFilterer{contract: contract}}, nil
}

// NewLimitBalanceMockCaller creates a new read-only instance of LimitBalanceMock, bound to a specific deployed contract.
func NewLimitBalanceMockCaller(address common.Address, caller bind.ContractCaller) (*LimitBalanceMockCaller, error) {
	contract, err := bindLimitBalanceMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimitBalanceMockCaller{contract: contract}, nil
}

// NewLimitBalanceMockTransactor creates a new write-only instance of LimitBalanceMock, bound to a specific deployed contract.
func NewLimitBalanceMockTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitBalanceMockTransactor, error) {
	contract, err := bindLimitBalanceMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimitBalanceMockTransactor{contract: contract}, nil
}

// NewLimitBalanceMockFilterer creates a new log filterer instance of LimitBalanceMock, bound to a specific deployed contract.
func NewLimitBalanceMockFilterer(address common.Address, filterer bind.ContractFilterer) (*LimitBalanceMockFilterer, error) {
	contract, err := bindLimitBalanceMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimitBalanceMockFilterer{contract: contract}, nil
}

// bindLimitBalanceMock binds a generic wrapper to an already deployed contract.
func bindLimitBalanceMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LimitBalanceMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitBalanceMock *LimitBalanceMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LimitBalanceMock.Contract.LimitBalanceMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitBalanceMock *LimitBalanceMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitBalanceMock.Contract.LimitBalanceMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitBalanceMock *LimitBalanceMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitBalanceMock.Contract.LimitBalanceMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitBalanceMock *LimitBalanceMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LimitBalanceMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitBalanceMock *LimitBalanceMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitBalanceMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitBalanceMock *LimitBalanceMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitBalanceMock.Contract.contract.Transact(opts, method, params...)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint256)
func (_LimitBalanceMock *LimitBalanceMockCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LimitBalanceMock.contract.Call(opts, out, "limit")
	return *ret0, err
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint256)
func (_LimitBalanceMock *LimitBalanceMockSession) Limit() (*big.Int, error) {
	return _LimitBalanceMock.Contract.Limit(&_LimitBalanceMock.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint256)
func (_LimitBalanceMock *LimitBalanceMockCallerSession) Limit() (*big.Int, error) {
	return _LimitBalanceMock.Contract.Limit(&_LimitBalanceMock.CallOpts)
}

// LimitedDeposit is a paid mutator transaction binding the contract method 0x86f3d0cd.
//
// Solidity: function limitedDeposit() returns()
func (_LimitBalanceMock *LimitBalanceMockTransactor) LimitedDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitBalanceMock.contract.Transact(opts, "limitedDeposit")
}

// LimitedDeposit is a paid mutator transaction binding the contract method 0x86f3d0cd.
//
// Solidity: function limitedDeposit() returns()
func (_LimitBalanceMock *LimitBalanceMockSession) LimitedDeposit() (*types.Transaction, error) {
	return _LimitBalanceMock.Contract.LimitedDeposit(&_LimitBalanceMock.TransactOpts)
}

// LimitedDeposit is a paid mutator transaction binding the contract method 0x86f3d0cd.
//
// Solidity: function limitedDeposit() returns()
func (_LimitBalanceMock *LimitBalanceMockTransactorSession) LimitedDeposit() (*types.Transaction, error) {
	return _LimitBalanceMock.Contract.LimitedDeposit(&_LimitBalanceMock.TransactOpts)
}
