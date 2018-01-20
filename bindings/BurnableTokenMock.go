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

// BurnableTokenMockABI is the input ABI used to generate the binding from.
const BurnableTokenMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BurnableTokenMockBin is the compiled bytecode used for deploying new contracts.
const BurnableTokenMockBin = `6060604052341561000f57600080fd5b60405160408061037d8339810160405280805191906020018051600160a060020a0390931660009081526001602052604081208490559290925550506103238061005a6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461006657806342966c681461008b57806370a08231146100a3578063a9059cbb146100c2575b600080fd5b341561007157600080fd5b6100796100f8565b60405190815260200160405180910390f35b341561009657600080fd5b6100a16004356100fe565b005b34156100ae57600080fd5b610079600160a060020a03600435166101b9565b34156100cd57600080fd5b6100e4600160a060020a03600435166024356101d4565b604051901515815260200160405180910390f35b60005481565b600160a060020a03331660009081526001602052604081205482111561012357600080fd5b5033600160a060020a03811660009081526001602052604090205461014890836102cf565b600160a060020a03821660009081526001602052604081209190915554610175908363ffffffff6102cf16565b600055600160a060020a0381167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58360405190815260200160405180910390a25050565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a03831615156101eb57600080fd5b600160a060020a03331660009081526001602052604090205482111561021057600080fd5b600160a060020a033316600090815260016020526040902054610239908363ffffffff6102cf16565b600160a060020a03338116600090815260016020526040808220939093559085168152205461026e908363ffffffff6102e116565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156102db57fe5b50900390565b6000828201838110156102f057fe5b93925050505600a165627a7a72305820393630db4396f52f0fcdf88021d8397322cb9dc91de6d34c441e5eed4d571d2f0029`

// DeployBurnableTokenMock deploys a new Ethereum contract, binding an instance of BurnableTokenMock to it.
func DeployBurnableTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *BurnableTokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnableTokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableTokenMock{BurnableTokenMockCaller: BurnableTokenMockCaller{contract: contract}, BurnableTokenMockTransactor: BurnableTokenMockTransactor{contract: contract}}, nil
}

// BurnableTokenMock is an auto generated Go binding around an Ethereum contract.
type BurnableTokenMock struct {
	BurnableTokenMockCaller     // Read-only binding to the contract
	BurnableTokenMockTransactor // Write-only binding to the contract
}

// BurnableTokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenMockSession struct {
	Contract     *BurnableTokenMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BurnableTokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenMockCallerSession struct {
	Contract *BurnableTokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BurnableTokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenMockTransactorSession struct {
	Contract     *BurnableTokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BurnableTokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenMockRaw struct {
	Contract *BurnableTokenMock // Generic contract binding to access the raw methods on
}

// BurnableTokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenMockCallerRaw struct {
	Contract *BurnableTokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenMockTransactorRaw struct {
	Contract *BurnableTokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableTokenMock creates a new instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMock(address common.Address, backend bind.ContractBackend) (*BurnableTokenMock, error) {
	contract, err := bindBurnableTokenMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMock{BurnableTokenMockCaller: BurnableTokenMockCaller{contract: contract}, BurnableTokenMockTransactor: BurnableTokenMockTransactor{contract: contract}}, nil
}

// NewBurnableTokenMockCaller creates a new read-only instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMockCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenMockCaller, error) {
	contract, err := bindBurnableTokenMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockCaller{contract: contract}, nil
}

// NewBurnableTokenMockTransactor creates a new write-only instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenMockTransactor, error) {
	contract, err := bindBurnableTokenMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockTransactor{contract: contract}, nil
}

// bindBurnableTokenMock binds a generic wrapper to an already deployed contract.
func bindBurnableTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableTokenMock *BurnableTokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableTokenMock.Contract.BurnableTokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableTokenMock *BurnableTokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.BurnableTokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableTokenMock *BurnableTokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.BurnableTokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableTokenMock *BurnableTokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableTokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableTokenMock *BurnableTokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableTokenMock *BurnableTokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableTokenMock *BurnableTokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableTokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableTokenMock *BurnableTokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableTokenMock.Contract.BalanceOf(&_BurnableTokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableTokenMock *BurnableTokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableTokenMock.Contract.BalanceOf(&_BurnableTokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableTokenMock *BurnableTokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableTokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableTokenMock *BurnableTokenMockSession) TotalSupply() (*big.Int, error) {
	return _BurnableTokenMock.Contract.TotalSupply(&_BurnableTokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableTokenMock *BurnableTokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableTokenMock.Contract.TotalSupply(&_BurnableTokenMock.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableTokenMock *BurnableTokenMockTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableTokenMock *BurnableTokenMockSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Burn(&_BurnableTokenMock.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableTokenMock *BurnableTokenMockTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Burn(&_BurnableTokenMock.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableTokenMock *BurnableTokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableTokenMock *BurnableTokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Transfer(&_BurnableTokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableTokenMock *BurnableTokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Transfer(&_BurnableTokenMock.TransactOpts, _to, _value)
}
