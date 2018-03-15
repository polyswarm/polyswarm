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

// DayLimitMockABI is the input ABI used to generate the binding from.
const DayLimitMockABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"resetSpentToday\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSpending\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newLimit\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"attemptSpend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DayLimitMockBin is the compiled bytecode used for deploying new contracts.
const DayLimitMockBin = `6060604052341561000f57600080fd5b60405160208061027883398101604052808051600081905591508190506100416401000000006101db61005082021704565b60025550506000600355610059565b62015180420490565b610210806100686000396000f3006060604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635c52c2f5811461008757806367eeba0c1461009c5780636997bcab146100c15780636b0c932d146100d4578063b20d30a9146100e7578063b349f6c7146100fd578063f059cf2b14610113575b600080fd5b341561009257600080fd5b61009a610126565b005b34156100a757600080fd5b6100af610130565b60405190815260200160405180910390f35b34156100cc57600080fd5b6100af610136565b34156100df57600080fd5b6100af61013c565b34156100f257600080fd5b61009a600435610142565b341561010857600080fd5b61009a60043561014e565b341561011e57600080fd5b6100af61016f565b61012e610175565b565b60005481565b60035481565b60025481565b61014b8161017c565b50565b8061015881610181565b151561016357600080fd5b50600380549091019055565b60015481565b6000600155565b600055565b600060025461018e6101db565b11156101a55760006001556101a16101db565b6002555b600154828101108015906101bf5750600054826001540111155b156101d2575060018054820181556101d6565b5060005b919050565b620151804204905600a165627a7a72305820c352421414ffc9da26013a2908bfba259ec1f1c19dada9d348223f975a76c56b0029`

// DeployDayLimitMock deploys a new Ethereum contract, binding an instance of DayLimitMock to it.
func DeployDayLimitMock(auth *bind.TransactOpts, backend bind.ContractBackend, _value *big.Int) (common.Address, *types.Transaction, *DayLimitMock, error) {
	parsed, err := abi.JSON(strings.NewReader(DayLimitMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DayLimitMockBin), backend, _value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DayLimitMock{DayLimitMockCaller: DayLimitMockCaller{contract: contract}, DayLimitMockTransactor: DayLimitMockTransactor{contract: contract}, DayLimitMockFilterer: DayLimitMockFilterer{contract: contract}}, nil
}

// DayLimitMock is an auto generated Go binding around an Ethereum contract.
type DayLimitMock struct {
	DayLimitMockCaller     // Read-only binding to the contract
	DayLimitMockTransactor // Write-only binding to the contract
	DayLimitMockFilterer   // Log filterer for contract events
}

// DayLimitMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DayLimitMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayLimitMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DayLimitMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayLimitMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DayLimitMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayLimitMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DayLimitMockSession struct {
	Contract     *DayLimitMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DayLimitMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DayLimitMockCallerSession struct {
	Contract *DayLimitMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DayLimitMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DayLimitMockTransactorSession struct {
	Contract     *DayLimitMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DayLimitMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DayLimitMockRaw struct {
	Contract *DayLimitMock // Generic contract binding to access the raw methods on
}

// DayLimitMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DayLimitMockCallerRaw struct {
	Contract *DayLimitMockCaller // Generic read-only contract binding to access the raw methods on
}

// DayLimitMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DayLimitMockTransactorRaw struct {
	Contract *DayLimitMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDayLimitMock creates a new instance of DayLimitMock, bound to a specific deployed contract.
func NewDayLimitMock(address common.Address, backend bind.ContractBackend) (*DayLimitMock, error) {
	contract, err := bindDayLimitMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DayLimitMock{DayLimitMockCaller: DayLimitMockCaller{contract: contract}, DayLimitMockTransactor: DayLimitMockTransactor{contract: contract}, DayLimitMockFilterer: DayLimitMockFilterer{contract: contract}}, nil
}

// NewDayLimitMockCaller creates a new read-only instance of DayLimitMock, bound to a specific deployed contract.
func NewDayLimitMockCaller(address common.Address, caller bind.ContractCaller) (*DayLimitMockCaller, error) {
	contract, err := bindDayLimitMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DayLimitMockCaller{contract: contract}, nil
}

// NewDayLimitMockTransactor creates a new write-only instance of DayLimitMock, bound to a specific deployed contract.
func NewDayLimitMockTransactor(address common.Address, transactor bind.ContractTransactor) (*DayLimitMockTransactor, error) {
	contract, err := bindDayLimitMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DayLimitMockTransactor{contract: contract}, nil
}

// NewDayLimitMockFilterer creates a new log filterer instance of DayLimitMock, bound to a specific deployed contract.
func NewDayLimitMockFilterer(address common.Address, filterer bind.ContractFilterer) (*DayLimitMockFilterer, error) {
	contract, err := bindDayLimitMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DayLimitMockFilterer{contract: contract}, nil
}

// bindDayLimitMock binds a generic wrapper to an already deployed contract.
func bindDayLimitMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DayLimitMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DayLimitMock *DayLimitMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DayLimitMock.Contract.DayLimitMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DayLimitMock *DayLimitMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayLimitMock.Contract.DayLimitMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DayLimitMock *DayLimitMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DayLimitMock.Contract.DayLimitMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DayLimitMock *DayLimitMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DayLimitMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DayLimitMock *DayLimitMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayLimitMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DayLimitMock *DayLimitMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DayLimitMock.Contract.contract.Transact(opts, method, params...)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimitMock.contract.Call(opts, out, "dailyLimit")
	return *ret0, err
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_DayLimitMock *DayLimitMockSession) DailyLimit() (*big.Int, error) {
	return _DayLimitMock.Contract.DailyLimit(&_DayLimitMock.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCallerSession) DailyLimit() (*big.Int, error) {
	return _DayLimitMock.Contract.DailyLimit(&_DayLimitMock.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimitMock.contract.Call(opts, out, "lastDay")
	return *ret0, err
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_DayLimitMock *DayLimitMockSession) LastDay() (*big.Int, error) {
	return _DayLimitMock.Contract.LastDay(&_DayLimitMock.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCallerSession) LastDay() (*big.Int, error) {
	return _DayLimitMock.Contract.LastDay(&_DayLimitMock.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimitMock.contract.Call(opts, out, "spentToday")
	return *ret0, err
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_DayLimitMock *DayLimitMockSession) SpentToday() (*big.Int, error) {
	return _DayLimitMock.Contract.SpentToday(&_DayLimitMock.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCallerSession) SpentToday() (*big.Int, error) {
	return _DayLimitMock.Contract.SpentToday(&_DayLimitMock.CallOpts)
}

// TotalSpending is a free data retrieval call binding the contract method 0x6997bcab.
//
// Solidity: function totalSpending() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCaller) TotalSpending(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimitMock.contract.Call(opts, out, "totalSpending")
	return *ret0, err
}

// TotalSpending is a free data retrieval call binding the contract method 0x6997bcab.
//
// Solidity: function totalSpending() constant returns(uint256)
func (_DayLimitMock *DayLimitMockSession) TotalSpending() (*big.Int, error) {
	return _DayLimitMock.Contract.TotalSpending(&_DayLimitMock.CallOpts)
}

// TotalSpending is a free data retrieval call binding the contract method 0x6997bcab.
//
// Solidity: function totalSpending() constant returns(uint256)
func (_DayLimitMock *DayLimitMockCallerSession) TotalSpending() (*big.Int, error) {
	return _DayLimitMock.Contract.TotalSpending(&_DayLimitMock.CallOpts)
}

// AttemptSpend is a paid mutator transaction binding the contract method 0xb349f6c7.
//
// Solidity: function attemptSpend(_value uint256) returns()
func (_DayLimitMock *DayLimitMockTransactor) AttemptSpend(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DayLimitMock.contract.Transact(opts, "attemptSpend", _value)
}

// AttemptSpend is a paid mutator transaction binding the contract method 0xb349f6c7.
//
// Solidity: function attemptSpend(_value uint256) returns()
func (_DayLimitMock *DayLimitMockSession) AttemptSpend(_value *big.Int) (*types.Transaction, error) {
	return _DayLimitMock.Contract.AttemptSpend(&_DayLimitMock.TransactOpts, _value)
}

// AttemptSpend is a paid mutator transaction binding the contract method 0xb349f6c7.
//
// Solidity: function attemptSpend(_value uint256) returns()
func (_DayLimitMock *DayLimitMockTransactorSession) AttemptSpend(_value *big.Int) (*types.Transaction, error) {
	return _DayLimitMock.Contract.AttemptSpend(&_DayLimitMock.TransactOpts, _value)
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_DayLimitMock *DayLimitMockTransactor) ResetSpentToday(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayLimitMock.contract.Transact(opts, "resetSpentToday")
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_DayLimitMock *DayLimitMockSession) ResetSpentToday() (*types.Transaction, error) {
	return _DayLimitMock.Contract.ResetSpentToday(&_DayLimitMock.TransactOpts)
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_DayLimitMock *DayLimitMockTransactorSession) ResetSpentToday() (*types.Transaction, error) {
	return _DayLimitMock.Contract.ResetSpentToday(&_DayLimitMock.TransactOpts)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(_newLimit uint256) returns()
func (_DayLimitMock *DayLimitMockTransactor) SetDailyLimit(opts *bind.TransactOpts, _newLimit *big.Int) (*types.Transaction, error) {
	return _DayLimitMock.contract.Transact(opts, "setDailyLimit", _newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(_newLimit uint256) returns()
func (_DayLimitMock *DayLimitMockSession) SetDailyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _DayLimitMock.Contract.SetDailyLimit(&_DayLimitMock.TransactOpts, _newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(_newLimit uint256) returns()
func (_DayLimitMock *DayLimitMockTransactorSession) SetDailyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _DayLimitMock.Contract.SetDailyLimit(&_DayLimitMock.TransactOpts, _newLimit)
}
