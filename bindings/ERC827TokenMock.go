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

// ERC827TokenMockABI is the input ABI used to generate the binding from.
const ERC827TokenMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC827TokenMockBin is the compiled bytecode used for deploying new contracts.
const ERC827TokenMockBin = `6060604052341561000f57600080fd5b604051604080610b888339810160405280805191906020018051600160a060020a0390931660009081526020819052604090208390555050600155610b2f806100596000396000f3006060604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100c957806316ca3b63146100ff57806318160ddd1461016457806323b872dd146101895780635c17f9f4146101b1578063661884631461021657806370a08231146102385780637272ad4914610257578063a9059cbb146102bc578063ab67aa58146102de578063be45fd621461034a578063d73dd623146103af578063dd62ed3e146103d1575b600080fd5b34156100d457600080fd5b6100eb600160a060020a03600435166024356103f6565b604051901515815260200160405180910390f35b341561010a57600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061046295505050505050565b341561016f57600080fd5b61017761051c565b60405190815260200160405180910390f35b341561019457600080fd5b6100eb600160a060020a0360043581169060243516604435610522565b34156101bc57600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506106a295505050505050565b341561022157600080fd5b6100eb600160a060020a03600435166024356106cf565b341561024357600080fd5b610177600160a060020a03600435166107c9565b341561026257600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506107e495505050505050565b34156102c757600080fd5b6100eb600160a060020a0360043516602435610811565b34156102e957600080fd5b6100eb600160a060020a036004803582169160248035909116916044359160849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061092395505050505050565b341561035557600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109df95505050505050565b34156103ba57600080fd5b6100eb600160a060020a0360043516602435610a0c565b34156103dc57600080fd5b610177600160a060020a0360043581169060243516610ab0565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b600030600160a060020a031684600160a060020a03161415151561048557600080fd5b61048f8484610a0c565b5083600160a060020a03168260405180828051906020019080838360005b838110156104c55780820151838201526020016104ad565b50505050905090810190601f1680156104f25780820380516001836020036101000a031916815260200191505b509150506000604051808303816000865af1915050151561051257600080fd5b5060019392505050565b60015490565b6000600160a060020a038316151561053957600080fd5b600160a060020a03841660009081526020819052604090205482111561055e57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561059157600080fd5b600160a060020a0384166000908152602081905260409020546105ba908363ffffffff610adb16565b600160a060020a0380861660009081526020819052604080822093909355908516815220546105ef908363ffffffff610aed16565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610635908363ffffffff610adb16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600030600160a060020a031684600160a060020a0316141515156106c557600080fd5b61048f84846103f6565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561072c57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610763565b61073c818463ffffffff610adb16565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b600030600160a060020a031684600160a060020a03161415151561080757600080fd5b61048f84846106cf565b6000600160a060020a038316151561082857600080fd5b600160a060020a03331660009081526020819052604090205482111561084d57600080fd5b600160a060020a033316600090815260208190526040902054610876908363ffffffff610adb16565b600160a060020a0333811660009081526020819052604080822093909355908516815220546108ab908363ffffffff610aed16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600030600160a060020a031684600160a060020a03161415151561094657600080fd5b610951858585610522565b5083600160a060020a03168260405180828051906020019080838360005b8381101561098757808201518382015260200161096f565b50505050905090810190601f1680156109b45780820380516001836020036101000a031916815260200191505b509150506000604051808303816000865af191505015156109d457600080fd5b506001949350505050565b600030600160a060020a031684600160a060020a031614151515610a0257600080fd5b61048f8484610811565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610a44908363ffffffff610aed16565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600082821115610ae757fe5b50900390565b600082820183811015610afc57fe5b93925050505600a165627a7a723058206b6b7867f09b38ee7b1650d5fb8072c7016b2893193542c0d114f128e49202f40029`

// DeployERC827TokenMock deploys a new Ethereum contract, binding an instance of ERC827TokenMock to it.
func DeployERC827TokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *ERC827TokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC827TokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC827TokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC827TokenMock{ERC827TokenMockCaller: ERC827TokenMockCaller{contract: contract}, ERC827TokenMockTransactor: ERC827TokenMockTransactor{contract: contract}}, nil
}

// ERC827TokenMock is an auto generated Go binding around an Ethereum contract.
type ERC827TokenMock struct {
	ERC827TokenMockCaller     // Read-only binding to the contract
	ERC827TokenMockTransactor // Write-only binding to the contract
}

// ERC827TokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC827TokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827TokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC827TokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827TokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC827TokenMockSession struct {
	Contract     *ERC827TokenMock  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC827TokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC827TokenMockCallerSession struct {
	Contract *ERC827TokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC827TokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC827TokenMockTransactorSession struct {
	Contract     *ERC827TokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC827TokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC827TokenMockRaw struct {
	Contract *ERC827TokenMock // Generic contract binding to access the raw methods on
}

// ERC827TokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC827TokenMockCallerRaw struct {
	Contract *ERC827TokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC827TokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC827TokenMockTransactorRaw struct {
	Contract *ERC827TokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC827TokenMock creates a new instance of ERC827TokenMock, bound to a specific deployed contract.
func NewERC827TokenMock(address common.Address, backend bind.ContractBackend) (*ERC827TokenMock, error) {
	contract, err := bindERC827TokenMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenMock{ERC827TokenMockCaller: ERC827TokenMockCaller{contract: contract}, ERC827TokenMockTransactor: ERC827TokenMockTransactor{contract: contract}}, nil
}

// NewERC827TokenMockCaller creates a new read-only instance of ERC827TokenMock, bound to a specific deployed contract.
func NewERC827TokenMockCaller(address common.Address, caller bind.ContractCaller) (*ERC827TokenMockCaller, error) {
	contract, err := bindERC827TokenMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenMockCaller{contract: contract}, nil
}

// NewERC827TokenMockTransactor creates a new write-only instance of ERC827TokenMock, bound to a specific deployed contract.
func NewERC827TokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC827TokenMockTransactor, error) {
	contract, err := bindERC827TokenMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenMockTransactor{contract: contract}, nil
}

// bindERC827TokenMock binds a generic wrapper to an already deployed contract.
func bindERC827TokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC827TokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC827TokenMock *ERC827TokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC827TokenMock.Contract.ERC827TokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC827TokenMock *ERC827TokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.ERC827TokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC827TokenMock *ERC827TokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.ERC827TokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC827TokenMock *ERC827TokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC827TokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC827TokenMock *ERC827TokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC827TokenMock *ERC827TokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC827TokenMock *ERC827TokenMockCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827TokenMock.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC827TokenMock *ERC827TokenMockSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC827TokenMock.Contract.Allowance(&_ERC827TokenMock.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC827TokenMock *ERC827TokenMockCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC827TokenMock.Contract.Allowance(&_ERC827TokenMock.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC827TokenMock *ERC827TokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827TokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC827TokenMock *ERC827TokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC827TokenMock.Contract.BalanceOf(&_ERC827TokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC827TokenMock *ERC827TokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC827TokenMock.Contract.BalanceOf(&_ERC827TokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827TokenMock *ERC827TokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827TokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827TokenMock *ERC827TokenMockSession) TotalSupply() (*big.Int, error) {
	return _ERC827TokenMock.Contract.TotalSupply(&_ERC827TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827TokenMock *ERC827TokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC827TokenMock.Contract.TotalSupply(&_ERC827TokenMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.contract.Transact(opts, "approve", _spender, _value, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockSession) Approve(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.Approve(&_ERC827TokenMock.TransactOpts, _spender, _value, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactorSession) Approve(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.Approve(&_ERC827TokenMock.TransactOpts, _spender, _value, _data)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x7272ad49.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue, _data)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x7272ad49.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.DecreaseApproval(&_ERC827TokenMock.TransactOpts, _spender, _subtractedValue, _data)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x7272ad49.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.DecreaseApproval(&_ERC827TokenMock.TransactOpts, _spender, _subtractedValue, _data)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC827TokenMock.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.IncreaseApproval(&_ERC827TokenMock.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.IncreaseApproval(&_ERC827TokenMock.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.Transfer(&_ERC827TokenMock.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.Transfer(&_ERC827TokenMock.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.contract.Transact(opts, "transferFrom", _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.TransferFrom(&_ERC827TokenMock.TransactOpts, _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827TokenMock *ERC827TokenMockTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827TokenMock.Contract.TransferFrom(&_ERC827TokenMock.TransactOpts, _from, _to, _value, _data)
}
