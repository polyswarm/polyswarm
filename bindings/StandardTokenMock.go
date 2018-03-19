// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// StandardTokenMockABI is the input ABI used to generate the binding from.
const StandardTokenMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenMockBin is the compiled bytecode used for deploying new contracts.
const StandardTokenMockBin = `6060604052341561000f57600080fd5b6040516040806107548339810160405280805191906020018051600160a060020a03909316600090815260208190526040902083905550506001556106fb806100596000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100c857806323b872dd146100ed578063661884631461011557806370a0823114610137578063a9059cbb14610156578063d73dd62314610178578063dd62ed3e1461019a575b600080fd5b341561009d57600080fd5b6100b4600160a060020a03600435166024356101bf565b604051901515815260200160405180910390f35b34156100d357600080fd5b6100db61022b565b60405190815260200160405180910390f35b34156100f857600080fd5b6100b4600160a060020a0360043581169060243516604435610231565b341561012057600080fd5b6100b4600160a060020a03600435166024356103b1565b341561014257600080fd5b6100db600160a060020a03600435166104ab565b341561016157600080fd5b6100b4600160a060020a03600435166024356104c6565b341561018357600080fd5b6100b4600160a060020a03600435166024356105d8565b34156101a557600080fd5b6100db600160a060020a036004358116906024351661067c565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b6000600160a060020a038316151561024857600080fd5b600160a060020a03841660009081526020819052604090205482111561026d57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156102a057600080fd5b600160a060020a0384166000908152602081905260409020546102c9908363ffffffff6106a716565b600160a060020a0380861660009081526020819052604080822093909355908516815220546102fe908363ffffffff6106b916565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610344908363ffffffff6106a716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561040e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610445565b61041e818463ffffffff6106a716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a03831615156104dd57600080fd5b600160a060020a03331660009081526020819052604090205482111561050257600080fd5b600160a060020a03331660009081526020819052604090205461052b908363ffffffff6106a716565b600160a060020a033381166000908152602081905260408082209390935590851681522054610560908363ffffffff6106b916565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610610908363ffffffff6106b916565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b6000828211156106b357fe5b50900390565b6000828201838110156106c857fe5b93925050505600a165627a7a72305820b8ec953c50ad8ecf75d9227e04f95465043f9d6abc8aec5eb081a8ee1fcbf2f80029`

// DeployStandardTokenMock deploys a new Ethereum contract, binding an instance of StandardTokenMock to it.
func DeployStandardTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *StandardTokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardTokenMock{StandardTokenMockCaller: StandardTokenMockCaller{contract: contract}, StandardTokenMockTransactor: StandardTokenMockTransactor{contract: contract}, StandardTokenMockFilterer: StandardTokenMockFilterer{contract: contract}}, nil
}

// StandardTokenMock is an auto generated Go binding around an Ethereum contract.
type StandardTokenMock struct {
	StandardTokenMockCaller     // Read-only binding to the contract
	StandardTokenMockTransactor // Write-only binding to the contract
	StandardTokenMockFilterer   // Log filterer for contract events
}

// StandardTokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenMockSession struct {
	Contract     *StandardTokenMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StandardTokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenMockCallerSession struct {
	Contract *StandardTokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StandardTokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenMockTransactorSession struct {
	Contract     *StandardTokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StandardTokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenMockRaw struct {
	Contract *StandardTokenMock // Generic contract binding to access the raw methods on
}

// StandardTokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenMockCallerRaw struct {
	Contract *StandardTokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenMockTransactorRaw struct {
	Contract *StandardTokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardTokenMock creates a new instance of StandardTokenMock, bound to a specific deployed contract.
func NewStandardTokenMock(address common.Address, backend bind.ContractBackend) (*StandardTokenMock, error) {
	contract, err := bindStandardTokenMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardTokenMock{StandardTokenMockCaller: StandardTokenMockCaller{contract: contract}, StandardTokenMockTransactor: StandardTokenMockTransactor{contract: contract}, StandardTokenMockFilterer: StandardTokenMockFilterer{contract: contract}}, nil
}

// NewStandardTokenMockCaller creates a new read-only instance of StandardTokenMock, bound to a specific deployed contract.
func NewStandardTokenMockCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenMockCaller, error) {
	contract, err := bindStandardTokenMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenMockCaller{contract: contract}, nil
}

// NewStandardTokenMockTransactor creates a new write-only instance of StandardTokenMock, bound to a specific deployed contract.
func NewStandardTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenMockTransactor, error) {
	contract, err := bindStandardTokenMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenMockTransactor{contract: contract}, nil
}

// NewStandardTokenMockFilterer creates a new log filterer instance of StandardTokenMock, bound to a specific deployed contract.
func NewStandardTokenMockFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenMockFilterer, error) {
	contract, err := bindStandardTokenMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenMockFilterer{contract: contract}, nil
}

// bindStandardTokenMock binds a generic wrapper to an already deployed contract.
func bindStandardTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardTokenMock *StandardTokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardTokenMock.Contract.StandardTokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardTokenMock *StandardTokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.StandardTokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardTokenMock *StandardTokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.StandardTokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardTokenMock *StandardTokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardTokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardTokenMock *StandardTokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardTokenMock *StandardTokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardTokenMock *StandardTokenMockCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardTokenMock.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardTokenMock *StandardTokenMockSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardTokenMock.Contract.Allowance(&_StandardTokenMock.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardTokenMock *StandardTokenMockCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardTokenMock.Contract.Allowance(&_StandardTokenMock.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardTokenMock *StandardTokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardTokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardTokenMock *StandardTokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardTokenMock.Contract.BalanceOf(&_StandardTokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardTokenMock *StandardTokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardTokenMock.Contract.BalanceOf(&_StandardTokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardTokenMock *StandardTokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardTokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardTokenMock *StandardTokenMockSession) TotalSupply() (*big.Int, error) {
	return _StandardTokenMock.Contract.TotalSupply(&_StandardTokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardTokenMock *StandardTokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardTokenMock.Contract.TotalSupply(&_StandardTokenMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.Approve(&_StandardTokenMock.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.Approve(&_StandardTokenMock.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.DecreaseApproval(&_StandardTokenMock.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.DecreaseApproval(&_StandardTokenMock.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.IncreaseApproval(&_StandardTokenMock.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.IncreaseApproval(&_StandardTokenMock.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.Transfer(&_StandardTokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.Transfer(&_StandardTokenMock.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.TransferFrom(&_StandardTokenMock.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardTokenMock *StandardTokenMockTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardTokenMock.Contract.TransferFrom(&_StandardTokenMock.TransactOpts, _from, _to, _value)
}

// StandardTokenMockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardTokenMock contract.
type StandardTokenMockApprovalIterator struct {
	Event *StandardTokenMockApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenMockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenMockApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenMockApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenMockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenMockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenMockApproval represents a Approval event raised by the StandardTokenMock contract.
type StandardTokenMockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardTokenMock *StandardTokenMockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenMockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardTokenMock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenMockApprovalIterator{contract: _StandardTokenMock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardTokenMock *StandardTokenMockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenMockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardTokenMock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenMockApproval)
				if err := _StandardTokenMock.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardTokenMock contract.
type StandardTokenMockTransferIterator struct {
	Event *StandardTokenMockTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenMockTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenMockTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenMockTransfer represents a Transfer event raised by the StandardTokenMock contract.
type StandardTokenMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_StandardTokenMock *StandardTokenMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardTokenMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenMockTransferIterator{contract: _StandardTokenMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_StandardTokenMock *StandardTokenMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardTokenMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenMockTransfer)
				if err := _StandardTokenMock.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
