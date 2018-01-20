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

// HasNoContractsABI is the input ABI used to generate the binding from.
const HasNoContractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"reclaimContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoContractsBin is the compiled bytecode used for deploying new contracts.
const HasNoContractsBin = `606060405260008054600160a060020a033316600160a060020a0319909116179055610244806100306000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632aed7f3f811461005b5780638da5cb5b1461007c578063f2fde38b146100ab575b600080fd5b341561006657600080fd5b61007a600160a060020a03600435166100ca565b005b341561008757600080fd5b61008f61016e565b604051600160a060020a03909116815260200160405180910390f35b34156100b657600080fd5b61007a600160a060020a036004351661017d565b6000805433600160a060020a039081169116146100e657600080fd5b506000548190600160a060020a038083169163f2fde38b91166040517c010000000000000000000000000000000000000000000000000000000063ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561015657600080fd5b6102c65a03f1151561016757600080fd5b5050505050565b600054600160a060020a031681565b60005433600160a060020a0390811691161461019857600080fd5b600160a060020a03811615156101ad57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820a9da93a76cf2c918d8d46dd4d46854d538e6f9cc8d36228f7db29a085927fc780029`

// DeployHasNoContracts deploys a new Ethereum contract, binding an instance of HasNoContracts to it.
func DeployHasNoContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoContracts{HasNoContractsCaller: HasNoContractsCaller{contract: contract}, HasNoContractsTransactor: HasNoContractsTransactor{contract: contract}}, nil
}

// HasNoContracts is an auto generated Go binding around an Ethereum contract.
type HasNoContracts struct {
	HasNoContractsCaller     // Read-only binding to the contract
	HasNoContractsTransactor // Write-only binding to the contract
}

// HasNoContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoContractsSession struct {
	Contract     *HasNoContracts   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoContractsCallerSession struct {
	Contract *HasNoContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HasNoContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoContractsTransactorSession struct {
	Contract     *HasNoContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HasNoContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoContractsRaw struct {
	Contract *HasNoContracts // Generic contract binding to access the raw methods on
}

// HasNoContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoContractsCallerRaw struct {
	Contract *HasNoContractsCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoContractsTransactorRaw struct {
	Contract *HasNoContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoContracts creates a new instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContracts(address common.Address, backend bind.ContractBackend) (*HasNoContracts, error) {
	contract, err := bindHasNoContracts(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoContracts{HasNoContractsCaller: HasNoContractsCaller{contract: contract}, HasNoContractsTransactor: HasNoContractsTransactor{contract: contract}}, nil
}

// NewHasNoContractsCaller creates a new read-only instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContractsCaller(address common.Address, caller bind.ContractCaller) (*HasNoContractsCaller, error) {
	contract, err := bindHasNoContracts(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoContractsCaller{contract: contract}, nil
}

// NewHasNoContractsTransactor creates a new write-only instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoContractsTransactor, error) {
	contract, err := bindHasNoContracts(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &HasNoContractsTransactor{contract: contract}, nil
}

// bindHasNoContracts binds a generic wrapper to an already deployed contract.
func bindHasNoContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoContracts *HasNoContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoContracts.Contract.HasNoContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoContracts *HasNoContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoContracts.Contract.HasNoContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoContracts *HasNoContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoContracts.Contract.HasNoContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoContracts *HasNoContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoContracts *HasNoContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoContracts *HasNoContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoContracts.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoContracts *HasNoContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoContracts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoContracts *HasNoContractsSession) Owner() (common.Address, error) {
	return _HasNoContracts.Contract.Owner(&_HasNoContracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoContracts *HasNoContractsCallerSession) Owner() (common.Address, error) {
	return _HasNoContracts.Contract.Owner(&_HasNoContracts.CallOpts)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_HasNoContracts *HasNoContractsTransactor) ReclaimContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _HasNoContracts.contract.Transact(opts, "reclaimContract", contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_HasNoContracts *HasNoContractsSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.ReclaimContract(&_HasNoContracts.TransactOpts, contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_HasNoContracts *HasNoContractsTransactorSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.ReclaimContract(&_HasNoContracts.TransactOpts, contractAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoContracts *HasNoContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoContracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoContracts *HasNoContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.TransferOwnership(&_HasNoContracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoContracts *HasNoContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.TransferOwnership(&_HasNoContracts.TransactOpts, newOwner)
}
