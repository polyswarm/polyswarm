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

// BasicTokenMockABI is the input ABI used to generate the binding from.
const BasicTokenMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenMockBin is the compiled bytecode used for deploying new contracts.
const BasicTokenMockBin = `6060604052341561000f57600080fd5b60405160408061029f8339810160405280805191906020018051600160a060020a0390931660009081526001602052604081208490559290925550506102458061005a6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610080578063a9059cbb1461009f575b600080fd5b341561006657600080fd5b61006e6100d5565b60405190815260200160405180910390f35b341561008b57600080fd5b61006e600160a060020a03600435166100db565b34156100aa57600080fd5b6100c1600160a060020a03600435166024356100f6565b604051901515815260200160405180910390f35b60005481565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a038316151561010d57600080fd5b600160a060020a03331660009081526001602052604090205482111561013257600080fd5b600160a060020a03331660009081526001602052604090205461015b908363ffffffff6101f116565b600160a060020a033381166000908152600160205260408082209390935590851681522054610190908363ffffffff61020316565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156101fd57fe5b50900390565b60008282018381101561021257fe5b93925050505600a165627a7a72305820ca2bce4b4eb02af90dfc8643094e0be146fdd3c018fa3d8110c89483f05bf0670029`

// DeployBasicTokenMock deploys a new Ethereum contract, binding an instance of BasicTokenMock to it.
func DeployBasicTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *BasicTokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicTokenMock{BasicTokenMockCaller: BasicTokenMockCaller{contract: contract}, BasicTokenMockTransactor: BasicTokenMockTransactor{contract: contract}}, nil
}

// BasicTokenMock is an auto generated Go binding around an Ethereum contract.
type BasicTokenMock struct {
	BasicTokenMockCaller     // Read-only binding to the contract
	BasicTokenMockTransactor // Write-only binding to the contract
}

// BasicTokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenMockSession struct {
	Contract     *BasicTokenMock   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenMockCallerSession struct {
	Contract *BasicTokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BasicTokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenMockTransactorSession struct {
	Contract     *BasicTokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BasicTokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenMockRaw struct {
	Contract *BasicTokenMock // Generic contract binding to access the raw methods on
}

// BasicTokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenMockCallerRaw struct {
	Contract *BasicTokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenMockTransactorRaw struct {
	Contract *BasicTokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicTokenMock creates a new instance of BasicTokenMock, bound to a specific deployed contract.
func NewBasicTokenMock(address common.Address, backend bind.ContractBackend) (*BasicTokenMock, error) {
	contract, err := bindBasicTokenMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicTokenMock{BasicTokenMockCaller: BasicTokenMockCaller{contract: contract}, BasicTokenMockTransactor: BasicTokenMockTransactor{contract: contract}}, nil
}

// NewBasicTokenMockCaller creates a new read-only instance of BasicTokenMock, bound to a specific deployed contract.
func NewBasicTokenMockCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenMockCaller, error) {
	contract, err := bindBasicTokenMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenMockCaller{contract: contract}, nil
}

// NewBasicTokenMockTransactor creates a new write-only instance of BasicTokenMock, bound to a specific deployed contract.
func NewBasicTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenMockTransactor, error) {
	contract, err := bindBasicTokenMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BasicTokenMockTransactor{contract: contract}, nil
}

// bindBasicTokenMock binds a generic wrapper to an already deployed contract.
func bindBasicTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicTokenMock *BasicTokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicTokenMock.Contract.BasicTokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicTokenMock *BasicTokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicTokenMock.Contract.BasicTokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicTokenMock *BasicTokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicTokenMock.Contract.BasicTokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicTokenMock *BasicTokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicTokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicTokenMock *BasicTokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicTokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicTokenMock *BasicTokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicTokenMock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicTokenMock *BasicTokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicTokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicTokenMock *BasicTokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicTokenMock.Contract.BalanceOf(&_BasicTokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicTokenMock *BasicTokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicTokenMock.Contract.BalanceOf(&_BasicTokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicTokenMock *BasicTokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicTokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicTokenMock *BasicTokenMockSession) TotalSupply() (*big.Int, error) {
	return _BasicTokenMock.Contract.TotalSupply(&_BasicTokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicTokenMock *BasicTokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicTokenMock.Contract.TotalSupply(&_BasicTokenMock.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicTokenMock *BasicTokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicTokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicTokenMock *BasicTokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicTokenMock.Contract.Transfer(&_BasicTokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicTokenMock *BasicTokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicTokenMock.Contract.Transfer(&_BasicTokenMock.TransactOpts, _to, _value)
}
