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

// ClaimableABI is the input ABI used to generate the binding from.
const ClaimableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ClaimableBin is the compiled bytecode used for deploying new contracts.
const ClaimableBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556101fe806100306000396000f3006060604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634e71e0c881146100665780638da5cb5b1461007b578063e30c3978146100aa578063f2fde38b146100bd575b600080fd5b341561007157600080fd5b6100796100dc565b005b341561008657600080fd5b61008e61016a565b604051600160a060020a03909116815260200160405180910390f35b34156100b557600080fd5b61008e610179565b34156100c857600080fd5b610079600160a060020a0360043516610188565b60015433600160a060020a039081169116146100f757600080fd5b600154600054600160a060020a0391821691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600154600160a060020a031681565b60005433600160a060020a039081169116146101a357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820446f0e42d79311b8c07cfb52c907de180440feb7b2ec6edb862e18173392dda20029`

// DeployClaimable deploys a new Ethereum contract, binding an instance of Claimable to it.
func DeployClaimable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Claimable, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Claimable{ClaimableCaller: ClaimableCaller{contract: contract}, ClaimableTransactor: ClaimableTransactor{contract: contract}}, nil
}

// Claimable is an auto generated Go binding around an Ethereum contract.
type Claimable struct {
	ClaimableCaller     // Read-only binding to the contract
	ClaimableTransactor // Write-only binding to the contract
}

// ClaimableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimableSession struct {
	Contract     *Claimable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimableCallerSession struct {
	Contract *ClaimableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ClaimableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimableTransactorSession struct {
	Contract     *ClaimableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ClaimableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimableRaw struct {
	Contract *Claimable // Generic contract binding to access the raw methods on
}

// ClaimableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimableCallerRaw struct {
	Contract *ClaimableCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimableTransactorRaw struct {
	Contract *ClaimableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimable creates a new instance of Claimable, bound to a specific deployed contract.
func NewClaimable(address common.Address, backend bind.ContractBackend) (*Claimable, error) {
	contract, err := bindClaimable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claimable{ClaimableCaller: ClaimableCaller{contract: contract}, ClaimableTransactor: ClaimableTransactor{contract: contract}}, nil
}

// NewClaimableCaller creates a new read-only instance of Claimable, bound to a specific deployed contract.
func NewClaimableCaller(address common.Address, caller bind.ContractCaller) (*ClaimableCaller, error) {
	contract, err := bindClaimable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableCaller{contract: contract}, nil
}

// NewClaimableTransactor creates a new write-only instance of Claimable, bound to a specific deployed contract.
func NewClaimableTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimableTransactor, error) {
	contract, err := bindClaimable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ClaimableTransactor{contract: contract}, nil
}

// bindClaimable binds a generic wrapper to an already deployed contract.
func bindClaimable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimable *ClaimableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Claimable.Contract.ClaimableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimable *ClaimableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimable.Contract.ClaimableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimable *ClaimableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimable.Contract.ClaimableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimable *ClaimableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Claimable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimable *ClaimableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimable *ClaimableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Claimable *ClaimableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Claimable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Claimable *ClaimableSession) Owner() (common.Address, error) {
	return _Claimable.Contract.Owner(&_Claimable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Claimable *ClaimableCallerSession) Owner() (common.Address, error) {
	return _Claimable.Contract.Owner(&_Claimable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Claimable *ClaimableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Claimable.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Claimable *ClaimableSession) PendingOwner() (common.Address, error) {
	return _Claimable.Contract.PendingOwner(&_Claimable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Claimable *ClaimableCallerSession) PendingOwner() (common.Address, error) {
	return _Claimable.Contract.PendingOwner(&_Claimable.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Claimable *ClaimableTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimable.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Claimable *ClaimableSession) ClaimOwnership() (*types.Transaction, error) {
	return _Claimable.Contract.ClaimOwnership(&_Claimable.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Claimable *ClaimableTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Claimable.Contract.ClaimOwnership(&_Claimable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Claimable *ClaimableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Claimable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Claimable *ClaimableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Claimable.Contract.TransferOwnership(&_Claimable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Claimable *ClaimableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Claimable.Contract.TransferOwnership(&_Claimable.TransactOpts, newOwner)
}
