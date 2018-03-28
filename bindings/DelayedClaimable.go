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

// DelayedClaimableABI is the input ABI used to generate the binding from.
const DelayedClaimableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_start\",\"type\":\"uint256\"},{\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setLimits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"end\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DelayedClaimableBin is the compiled bytecode used for deploying new contracts.
const DelayedClaimableBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556102d0806100306000396000f3006060604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634e71e0c881146100875780638da5cb5b1461009c578063be9a6555146100cb578063c4590d3f146100f0578063e30c397814610109578063efbe1c1c1461011c578063f2fde38b1461012f575b600080fd5b341561009257600080fd5b61009a61014e565b005b34156100a757600080fd5b6100af610200565b604051600160a060020a03909116815260200160405180910390f35b34156100d657600080fd5b6100de61020f565b60405190815260200160405180910390f35b34156100fb57600080fd5b61009a600435602435610215565b341561011457600080fd5b6100af610245565b341561012757600080fd5b6100de610254565b341561013a57600080fd5b61009a600160a060020a036004351661025a565b60015433600160a060020a0390811691161461016957600080fd5b600254431115801561017d57506003544310155b151561018857600080fd5b600154600054600160a060020a0391821691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a038416178255909116909155600255565b600054600160a060020a031681565b60035481565b60005433600160a060020a0390811691161461023057600080fd5b8082111561023d57600080fd5b600255600355565b600154600160a060020a031681565b60025481565b60005433600160a060020a0390811691161461027557600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058200e802ef22548c05c72aa0559f0535718d0a83c2e9282aa69e4e62f21d61bf4cc0029`

// DeployDelayedClaimable deploys a new Ethereum contract, binding an instance of DelayedClaimable to it.
func DeployDelayedClaimable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DelayedClaimable, error) {
	parsed, err := abi.JSON(strings.NewReader(DelayedClaimableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelayedClaimableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelayedClaimable{DelayedClaimableCaller: DelayedClaimableCaller{contract: contract}, DelayedClaimableTransactor: DelayedClaimableTransactor{contract: contract}}, nil
}

// DelayedClaimable is an auto generated Go binding around an Ethereum contract.
type DelayedClaimable struct {
	DelayedClaimableCaller     // Read-only binding to the contract
	DelayedClaimableTransactor // Write-only binding to the contract
}

// DelayedClaimableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedClaimableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedClaimableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedClaimableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedClaimableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedClaimableSession struct {
	Contract     *DelayedClaimable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelayedClaimableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedClaimableCallerSession struct {
	Contract *DelayedClaimableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DelayedClaimableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedClaimableTransactorSession struct {
	Contract     *DelayedClaimableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DelayedClaimableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedClaimableRaw struct {
	Contract *DelayedClaimable // Generic contract binding to access the raw methods on
}

// DelayedClaimableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedClaimableCallerRaw struct {
	Contract *DelayedClaimableCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedClaimableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedClaimableTransactorRaw struct {
	Contract *DelayedClaimableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedClaimable creates a new instance of DelayedClaimable, bound to a specific deployed contract.
func NewDelayedClaimable(address common.Address, backend bind.ContractBackend) (*DelayedClaimable, error) {
	contract, err := bindDelayedClaimable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedClaimable{DelayedClaimableCaller: DelayedClaimableCaller{contract: contract}, DelayedClaimableTransactor: DelayedClaimableTransactor{contract: contract}}, nil
}

// NewDelayedClaimableCaller creates a new read-only instance of DelayedClaimable, bound to a specific deployed contract.
func NewDelayedClaimableCaller(address common.Address, caller bind.ContractCaller) (*DelayedClaimableCaller, error) {
	contract, err := bindDelayedClaimable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedClaimableCaller{contract: contract}, nil
}

// NewDelayedClaimableTransactor creates a new write-only instance of DelayedClaimable, bound to a specific deployed contract.
func NewDelayedClaimableTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedClaimableTransactor, error) {
	contract, err := bindDelayedClaimable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DelayedClaimableTransactor{contract: contract}, nil
}

// bindDelayedClaimable binds a generic wrapper to an already deployed contract.
func bindDelayedClaimable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelayedClaimableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedClaimable *DelayedClaimableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelayedClaimable.Contract.DelayedClaimableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedClaimable *DelayedClaimableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.DelayedClaimableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedClaimable *DelayedClaimableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.DelayedClaimableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedClaimable *DelayedClaimableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelayedClaimable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedClaimable *DelayedClaimableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedClaimable *DelayedClaimableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.contract.Transact(opts, method, params...)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() constant returns(uint256)
func (_DelayedClaimable *DelayedClaimableCaller) End(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelayedClaimable.contract.Call(opts, out, "end")
	return *ret0, err
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() constant returns(uint256)
func (_DelayedClaimable *DelayedClaimableSession) End() (*big.Int, error) {
	return _DelayedClaimable.Contract.End(&_DelayedClaimable.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() constant returns(uint256)
func (_DelayedClaimable *DelayedClaimableCallerSession) End() (*big.Int, error) {
	return _DelayedClaimable.Contract.End(&_DelayedClaimable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DelayedClaimable *DelayedClaimableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DelayedClaimable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DelayedClaimable *DelayedClaimableSession) Owner() (common.Address, error) {
	return _DelayedClaimable.Contract.Owner(&_DelayedClaimable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DelayedClaimable *DelayedClaimableCallerSession) Owner() (common.Address, error) {
	return _DelayedClaimable.Contract.Owner(&_DelayedClaimable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_DelayedClaimable *DelayedClaimableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DelayedClaimable.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_DelayedClaimable *DelayedClaimableSession) PendingOwner() (common.Address, error) {
	return _DelayedClaimable.Contract.PendingOwner(&_DelayedClaimable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_DelayedClaimable *DelayedClaimableCallerSession) PendingOwner() (common.Address, error) {
	return _DelayedClaimable.Contract.PendingOwner(&_DelayedClaimable.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_DelayedClaimable *DelayedClaimableCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelayedClaimable.contract.Call(opts, out, "start")
	return *ret0, err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_DelayedClaimable *DelayedClaimableSession) Start() (*big.Int, error) {
	return _DelayedClaimable.Contract.Start(&_DelayedClaimable.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_DelayedClaimable *DelayedClaimableCallerSession) Start() (*big.Int, error) {
	return _DelayedClaimable.Contract.Start(&_DelayedClaimable.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DelayedClaimable *DelayedClaimableTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedClaimable.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DelayedClaimable *DelayedClaimableSession) ClaimOwnership() (*types.Transaction, error) {
	return _DelayedClaimable.Contract.ClaimOwnership(&_DelayedClaimable.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DelayedClaimable *DelayedClaimableTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _DelayedClaimable.Contract.ClaimOwnership(&_DelayedClaimable.TransactOpts)
}

// SetLimits is a paid mutator transaction binding the contract method 0xc4590d3f.
//
// Solidity: function setLimits(_start uint256, _end uint256) returns()
func (_DelayedClaimable *DelayedClaimableTransactor) SetLimits(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _DelayedClaimable.contract.Transact(opts, "setLimits", _start, _end)
}

// SetLimits is a paid mutator transaction binding the contract method 0xc4590d3f.
//
// Solidity: function setLimits(_start uint256, _end uint256) returns()
func (_DelayedClaimable *DelayedClaimableSession) SetLimits(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.SetLimits(&_DelayedClaimable.TransactOpts, _start, _end)
}

// SetLimits is a paid mutator transaction binding the contract method 0xc4590d3f.
//
// Solidity: function setLimits(_start uint256, _end uint256) returns()
func (_DelayedClaimable *DelayedClaimableTransactorSession) SetLimits(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.SetLimits(&_DelayedClaimable.TransactOpts, _start, _end)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DelayedClaimable *DelayedClaimableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelayedClaimable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DelayedClaimable *DelayedClaimableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.TransferOwnership(&_DelayedClaimable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DelayedClaimable *DelayedClaimableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedClaimable.Contract.TransferOwnership(&_DelayedClaimable.TransactOpts, newOwner)
}
