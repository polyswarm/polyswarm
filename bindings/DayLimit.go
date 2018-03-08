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

// DayLimitABI is the input ABI used to generate the binding from.
const DayLimitABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DayLimitBin is the compiled bytecode used for deploying new contracts.
const DayLimitBin = `6060604052341561000f57600080fd5b604051602080610140833981016040528080516000819055915061004090506401000000006100ab61004982021704565b60025550610052565b62015180420490565b60e0806100606000396000f30060606040526004361060525763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166367eeba0c811460575780636b0c932d146079578063f059cf2b146089575b600080fd5b3415606157600080fd5b60676099565b60405190815260200160405180910390f35b3415608357600080fd5b6067609f565b3415609357600080fd5b606760a5565b60005481565b60025481565b60015481565b620151804204905600a165627a7a723058200090a9f278695e0aab84a33059524cf84b7a889615082dc62e881b4044ff51720029`

// DeployDayLimit deploys a new Ethereum contract, binding an instance of DayLimit to it.
func DeployDayLimit(auth *bind.TransactOpts, backend bind.ContractBackend, _limit *big.Int) (common.Address, *types.Transaction, *DayLimit, error) {
	parsed, err := abi.JSON(strings.NewReader(DayLimitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DayLimitBin), backend, _limit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DayLimit{DayLimitCaller: DayLimitCaller{contract: contract}, DayLimitTransactor: DayLimitTransactor{contract: contract}, DayLimitFilterer: DayLimitFilterer{contract: contract}}, nil
}

// DayLimit is an auto generated Go binding around an Ethereum contract.
type DayLimit struct {
	DayLimitCaller     // Read-only binding to the contract
	DayLimitTransactor // Write-only binding to the contract
	DayLimitFilterer   // Log filterer for contract events
}

// DayLimitCaller is an auto generated read-only Go binding around an Ethereum contract.
type DayLimitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayLimitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DayLimitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayLimitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DayLimitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayLimitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DayLimitSession struct {
	Contract     *DayLimit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DayLimitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DayLimitCallerSession struct {
	Contract *DayLimitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DayLimitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DayLimitTransactorSession struct {
	Contract     *DayLimitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DayLimitRaw is an auto generated low-level Go binding around an Ethereum contract.
type DayLimitRaw struct {
	Contract *DayLimit // Generic contract binding to access the raw methods on
}

// DayLimitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DayLimitCallerRaw struct {
	Contract *DayLimitCaller // Generic read-only contract binding to access the raw methods on
}

// DayLimitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DayLimitTransactorRaw struct {
	Contract *DayLimitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDayLimit creates a new instance of DayLimit, bound to a specific deployed contract.
func NewDayLimit(address common.Address, backend bind.ContractBackend) (*DayLimit, error) {
	contract, err := bindDayLimit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DayLimit{DayLimitCaller: DayLimitCaller{contract: contract}, DayLimitTransactor: DayLimitTransactor{contract: contract}, DayLimitFilterer: DayLimitFilterer{contract: contract}}, nil
}

// NewDayLimitCaller creates a new read-only instance of DayLimit, bound to a specific deployed contract.
func NewDayLimitCaller(address common.Address, caller bind.ContractCaller) (*DayLimitCaller, error) {
	contract, err := bindDayLimit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DayLimitCaller{contract: contract}, nil
}

// NewDayLimitTransactor creates a new write-only instance of DayLimit, bound to a specific deployed contract.
func NewDayLimitTransactor(address common.Address, transactor bind.ContractTransactor) (*DayLimitTransactor, error) {
	contract, err := bindDayLimit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DayLimitTransactor{contract: contract}, nil
}

// NewDayLimitFilterer creates a new log filterer instance of DayLimit, bound to a specific deployed contract.
func NewDayLimitFilterer(address common.Address, filterer bind.ContractFilterer) (*DayLimitFilterer, error) {
	contract, err := bindDayLimit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DayLimitFilterer{contract: contract}, nil
}

// bindDayLimit binds a generic wrapper to an already deployed contract.
func bindDayLimit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DayLimitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DayLimit *DayLimitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DayLimit.Contract.DayLimitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DayLimit *DayLimitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayLimit.Contract.DayLimitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DayLimit *DayLimitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DayLimit.Contract.DayLimitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DayLimit *DayLimitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DayLimit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DayLimit *DayLimitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayLimit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DayLimit *DayLimitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DayLimit.Contract.contract.Transact(opts, method, params...)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_DayLimit *DayLimitCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimit.contract.Call(opts, out, "dailyLimit")
	return *ret0, err
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_DayLimit *DayLimitSession) DailyLimit() (*big.Int, error) {
	return _DayLimit.Contract.DailyLimit(&_DayLimit.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_DayLimit *DayLimitCallerSession) DailyLimit() (*big.Int, error) {
	return _DayLimit.Contract.DailyLimit(&_DayLimit.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_DayLimit *DayLimitCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimit.contract.Call(opts, out, "lastDay")
	return *ret0, err
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_DayLimit *DayLimitSession) LastDay() (*big.Int, error) {
	return _DayLimit.Contract.LastDay(&_DayLimit.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_DayLimit *DayLimitCallerSession) LastDay() (*big.Int, error) {
	return _DayLimit.Contract.LastDay(&_DayLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_DayLimit *DayLimitCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DayLimit.contract.Call(opts, out, "spentToday")
	return *ret0, err
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_DayLimit *DayLimitSession) SpentToday() (*big.Int, error) {
	return _DayLimit.Contract.SpentToday(&_DayLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_DayLimit *DayLimitCallerSession) SpentToday() (*big.Int, error) {
	return _DayLimit.Contract.SpentToday(&_DayLimit.CallOpts)
}
