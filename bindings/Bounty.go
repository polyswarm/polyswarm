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

// BountyABI is the input ABI used to generate the binding from.
const BountyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"researchers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawPayments\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createTarget\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"claimed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"createdAddress\",\"type\":\"address\"}],\"name\":\"TargetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Bounty is an auto generated Go binding around an Ethereum contract.
type Bounty struct {
	BountyCaller     // Read-only binding to the contract
	BountyTransactor // Write-only binding to the contract
}

// BountyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BountyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BountyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BountySession struct {
	Contract     *Bounty           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BountyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BountyCallerSession struct {
	Contract *BountyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BountyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BountyTransactorSession struct {
	Contract     *BountyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BountyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BountyRaw struct {
	Contract *Bounty // Generic contract binding to access the raw methods on
}

// BountyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BountyCallerRaw struct {
	Contract *BountyCaller // Generic read-only contract binding to access the raw methods on
}

// BountyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BountyTransactorRaw struct {
	Contract *BountyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBounty creates a new instance of Bounty, bound to a specific deployed contract.
func NewBounty(address common.Address, backend bind.ContractBackend) (*Bounty, error) {
	contract, err := bindBounty(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bounty{BountyCaller: BountyCaller{contract: contract}, BountyTransactor: BountyTransactor{contract: contract}}, nil
}

// NewBountyCaller creates a new read-only instance of Bounty, bound to a specific deployed contract.
func NewBountyCaller(address common.Address, caller bind.ContractCaller) (*BountyCaller, error) {
	contract, err := bindBounty(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BountyCaller{contract: contract}, nil
}

// NewBountyTransactor creates a new write-only instance of Bounty, bound to a specific deployed contract.
func NewBountyTransactor(address common.Address, transactor bind.ContractTransactor) (*BountyTransactor, error) {
	contract, err := bindBounty(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BountyTransactor{contract: contract}, nil
}

// bindBounty binds a generic wrapper to an already deployed contract.
func bindBounty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BountyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bounty *BountyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bounty.Contract.BountyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bounty *BountyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bounty.Contract.BountyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bounty *BountyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bounty.Contract.BountyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bounty *BountyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bounty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bounty *BountyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bounty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bounty *BountyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bounty.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(bool)
func (_Bounty *BountyCaller) Claimed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bounty.contract.Call(opts, out, "claimed")
	return *ret0, err
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(bool)
func (_Bounty *BountySession) Claimed() (bool, error) {
	return _Bounty.Contract.Claimed(&_Bounty.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(bool)
func (_Bounty *BountyCallerSession) Claimed() (bool, error) {
	return _Bounty.Contract.Claimed(&_Bounty.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bounty *BountyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bounty.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bounty *BountySession) Owner() (common.Address, error) {
	return _Bounty.Contract.Owner(&_Bounty.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bounty *BountyCallerSession) Owner() (common.Address, error) {
	return _Bounty.Contract.Owner(&_Bounty.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_Bounty *BountyCaller) Payments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bounty.contract.Call(opts, out, "payments", arg0)
	return *ret0, err
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_Bounty *BountySession) Payments(arg0 common.Address) (*big.Int, error) {
	return _Bounty.Contract.Payments(&_Bounty.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_Bounty *BountyCallerSession) Payments(arg0 common.Address) (*big.Int, error) {
	return _Bounty.Contract.Payments(&_Bounty.CallOpts, arg0)
}

// Researchers is a free data retrieval call binding the contract method 0x01bc237d.
//
// Solidity: function researchers( address) constant returns(address)
func (_Bounty *BountyCaller) Researchers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bounty.contract.Call(opts, out, "researchers", arg0)
	return *ret0, err
}

// Researchers is a free data retrieval call binding the contract method 0x01bc237d.
//
// Solidity: function researchers( address) constant returns(address)
func (_Bounty *BountySession) Researchers(arg0 common.Address) (common.Address, error) {
	return _Bounty.Contract.Researchers(&_Bounty.CallOpts, arg0)
}

// Researchers is a free data retrieval call binding the contract method 0x01bc237d.
//
// Solidity: function researchers( address) constant returns(address)
func (_Bounty *BountyCallerSession) Researchers(arg0 common.Address) (common.Address, error) {
	return _Bounty.Contract.Researchers(&_Bounty.CallOpts, arg0)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_Bounty *BountyCaller) TotalPayments(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bounty.contract.Call(opts, out, "totalPayments")
	return *ret0, err
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_Bounty *BountySession) TotalPayments() (*big.Int, error) {
	return _Bounty.Contract.TotalPayments(&_Bounty.CallOpts)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_Bounty *BountyCallerSession) TotalPayments() (*big.Int, error) {
	return _Bounty.Contract.TotalPayments(&_Bounty.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(target address) returns()
func (_Bounty *BountyTransactor) Claim(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Bounty.contract.Transact(opts, "claim", target)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(target address) returns()
func (_Bounty *BountySession) Claim(target common.Address) (*types.Transaction, error) {
	return _Bounty.Contract.Claim(&_Bounty.TransactOpts, target)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(target address) returns()
func (_Bounty *BountyTransactorSession) Claim(target common.Address) (*types.Transaction, error) {
	return _Bounty.Contract.Claim(&_Bounty.TransactOpts, target)
}

// CreateTarget is a paid mutator transaction binding the contract method 0xc98165b6.
//
// Solidity: function createTarget() returns(address)
func (_Bounty *BountyTransactor) CreateTarget(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bounty.contract.Transact(opts, "createTarget")
}

// CreateTarget is a paid mutator transaction binding the contract method 0xc98165b6.
//
// Solidity: function createTarget() returns(address)
func (_Bounty *BountySession) CreateTarget() (*types.Transaction, error) {
	return _Bounty.Contract.CreateTarget(&_Bounty.TransactOpts)
}

// CreateTarget is a paid mutator transaction binding the contract method 0xc98165b6.
//
// Solidity: function createTarget() returns(address)
func (_Bounty *BountyTransactorSession) CreateTarget() (*types.Transaction, error) {
	return _Bounty.Contract.CreateTarget(&_Bounty.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Bounty *BountyTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bounty.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Bounty *BountySession) Destroy() (*types.Transaction, error) {
	return _Bounty.Contract.Destroy(&_Bounty.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Bounty *BountyTransactorSession) Destroy() (*types.Transaction, error) {
	return _Bounty.Contract.Destroy(&_Bounty.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Bounty *BountyTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Bounty.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Bounty *BountySession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Bounty.Contract.DestroyAndSend(&_Bounty.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Bounty *BountyTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Bounty.Contract.DestroyAndSend(&_Bounty.TransactOpts, _recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bounty *BountyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bounty.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bounty *BountySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bounty.Contract.TransferOwnership(&_Bounty.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bounty *BountyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bounty.Contract.TransferOwnership(&_Bounty.TransactOpts, newOwner)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_Bounty *BountyTransactor) WithdrawPayments(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bounty.contract.Transact(opts, "withdrawPayments")
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_Bounty *BountySession) WithdrawPayments() (*types.Transaction, error) {
	return _Bounty.Contract.WithdrawPayments(&_Bounty.TransactOpts)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_Bounty *BountyTransactorSession) WithdrawPayments() (*types.Transaction, error) {
	return _Bounty.Contract.WithdrawPayments(&_Bounty.TransactOpts)
}
