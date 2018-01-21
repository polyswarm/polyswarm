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

// InsecureTargetBountyABI is the input ABI used to generate the binding from.
const InsecureTargetBountyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"researchers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawPayments\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createTarget\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"claimed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"createdAddress\",\"type\":\"address\"}],\"name\":\"TargetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// InsecureTargetBountyBin is the compiled bytecode used for deploying new contracts.
const InsecureTargetBountyBin = `606060405260028054600160a060020a03191633600160a060020a031617905561072c8061002e6000396000f3006060604052600436106100ad5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416625b448781146100d757806301bc237d146100fc5780631e83409a146101375780636103d70b1461015657806383197ef0146101695780638da5cb5b1461017c578063c98165b61461018f578063e2982c21146101a2578063e834a834146101c1578063f2fde38b146101e8578063f5074f4114610207575b60025474010000000000000000000000000000000000000000900460ff16156100d557600080fd5b005b34156100e257600080fd5b6100ea610226565b60405190815260200160405180910390f35b341561010757600080fd5b61011b600160a060020a036004351661022c565b604051600160a060020a03909116815260200160405180910390f35b341561014257600080fd5b6100d5600160a060020a0360043516610247565b341561016157600080fd5b6100d561033f565b341561017457600080fd5b6100d56103d8565b341561018757600080fd5b61011b610401565b341561019a57600080fd5b61011b610410565b34156101ad57600080fd5b6100ea600160a060020a036004351661049f565b34156101cc57600080fd5b6101d46104b1565b604051901515815260200160405180910390f35b34156101f357600080fd5b6100d5600160a060020a03600435166104d2565b341561021257600080fd5b6100d5600160a060020a036004351661056d565b60015481565b600360205260009081526040902054600160a060020a031681565b600160a060020a038082166000908152600360205260409020541680151561026e57600080fd5b81600160a060020a031663e79487da6000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156102cd57600080fd5b6102c65a03f115156102de57600080fd5b50505060405180511590506102f257600080fd5b6103068130600160a060020a031631610594565b50506002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b33600160a060020a03811660009081526020819052604090205480151561036557600080fd5b600160a060020a033016318190101561037d57600080fd5b600154610390908263ffffffff6105f016565b600155600160a060020a0382166000818152602081905260408082209190915582156108fc0290839051600060405180830381858888f1935050505015156103d457fe5b5050565b60025433600160a060020a039081169116146103f357600080fd5b600254600160a060020a0316ff5b600254600160a060020a031681565b60008061041b610602565b600160a060020a0381811660009081526003602052604090819020805473ffffffffffffffffffffffffffffffffffffffff191633909316929092179091559091507fe62d909feaad4aecbcea8fef32a9b41552373e45f5acb7ce7fc0a997180e7ae590829051600160a060020a03909116815260200160405180910390a1919050565b60006020819052908152604090205481565b60025474010000000000000000000000000000000000000000900460ff1681565b60025433600160a060020a039081169116146104ed57600080fd5b600160a060020a038116151561050257600080fd5b600254600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025433600160a060020a0390811691161461058857600080fd5b80600160a060020a0316ff5b600160a060020a0382166000908152602081905260409020546105bd908263ffffffff61062716565b600160a060020a0383166000908152602081905260409020556001546105e9908263ffffffff61062716565b6001555050565b6000828211156105fc57fe5b50900390565b600061060c61063d565b604051809103906000f080151561062257600080fd5b905090565b60008282018381101561063657fe5b9392505050565b60405160b48061064d83390190560060606040523415600e57600080fd5b60988061001c6000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e79487da81146043575b600080fd5b3415604d57600080fd5b60536067565b604051901515815260200160405180910390f35b6000905600a165627a7a723058200f63b2b3de783fc2b5ea47ecd95e0d60fa23f99a52955d9920e4235197f068dc0029a165627a7a7230582074f2fd9e4deea2c277d13ef0e967225a9219a8d0c2215f4703d166d282aa2fc00029`

// DeployInsecureTargetBounty deploys a new Ethereum contract, binding an instance of InsecureTargetBounty to it.
func DeployInsecureTargetBounty(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InsecureTargetBounty, error) {
	parsed, err := abi.JSON(strings.NewReader(InsecureTargetBountyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InsecureTargetBountyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InsecureTargetBounty{InsecureTargetBountyCaller: InsecureTargetBountyCaller{contract: contract}, InsecureTargetBountyTransactor: InsecureTargetBountyTransactor{contract: contract}}, nil
}

// InsecureTargetBounty is an auto generated Go binding around an Ethereum contract.
type InsecureTargetBounty struct {
	InsecureTargetBountyCaller     // Read-only binding to the contract
	InsecureTargetBountyTransactor // Write-only binding to the contract
}

// InsecureTargetBountyCaller is an auto generated read-only Go binding around an Ethereum contract.
type InsecureTargetBountyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsecureTargetBountyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InsecureTargetBountyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsecureTargetBountySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InsecureTargetBountySession struct {
	Contract     *InsecureTargetBounty // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InsecureTargetBountyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InsecureTargetBountyCallerSession struct {
	Contract *InsecureTargetBountyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InsecureTargetBountyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InsecureTargetBountyTransactorSession struct {
	Contract     *InsecureTargetBountyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InsecureTargetBountyRaw is an auto generated low-level Go binding around an Ethereum contract.
type InsecureTargetBountyRaw struct {
	Contract *InsecureTargetBounty // Generic contract binding to access the raw methods on
}

// InsecureTargetBountyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InsecureTargetBountyCallerRaw struct {
	Contract *InsecureTargetBountyCaller // Generic read-only contract binding to access the raw methods on
}

// InsecureTargetBountyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InsecureTargetBountyTransactorRaw struct {
	Contract *InsecureTargetBountyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInsecureTargetBounty creates a new instance of InsecureTargetBounty, bound to a specific deployed contract.
func NewInsecureTargetBounty(address common.Address, backend bind.ContractBackend) (*InsecureTargetBounty, error) {
	contract, err := bindInsecureTargetBounty(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetBounty{InsecureTargetBountyCaller: InsecureTargetBountyCaller{contract: contract}, InsecureTargetBountyTransactor: InsecureTargetBountyTransactor{contract: contract}}, nil
}

// NewInsecureTargetBountyCaller creates a new read-only instance of InsecureTargetBounty, bound to a specific deployed contract.
func NewInsecureTargetBountyCaller(address common.Address, caller bind.ContractCaller) (*InsecureTargetBountyCaller, error) {
	contract, err := bindInsecureTargetBounty(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetBountyCaller{contract: contract}, nil
}

// NewInsecureTargetBountyTransactor creates a new write-only instance of InsecureTargetBounty, bound to a specific deployed contract.
func NewInsecureTargetBountyTransactor(address common.Address, transactor bind.ContractTransactor) (*InsecureTargetBountyTransactor, error) {
	contract, err := bindInsecureTargetBounty(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &InsecureTargetBountyTransactor{contract: contract}, nil
}

// bindInsecureTargetBounty binds a generic wrapper to an already deployed contract.
func bindInsecureTargetBounty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InsecureTargetBountyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsecureTargetBounty *InsecureTargetBountyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InsecureTargetBounty.Contract.InsecureTargetBountyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsecureTargetBounty *InsecureTargetBountyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.InsecureTargetBountyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsecureTargetBounty *InsecureTargetBountyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.InsecureTargetBountyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsecureTargetBounty *InsecureTargetBountyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InsecureTargetBounty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsecureTargetBounty *InsecureTargetBountyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsecureTargetBounty *InsecureTargetBountyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(bool)
func (_InsecureTargetBounty *InsecureTargetBountyCaller) Claimed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InsecureTargetBounty.contract.Call(opts, out, "claimed")
	return *ret0, err
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(bool)
func (_InsecureTargetBounty *InsecureTargetBountySession) Claimed() (bool, error) {
	return _InsecureTargetBounty.Contract.Claimed(&_InsecureTargetBounty.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(bool)
func (_InsecureTargetBounty *InsecureTargetBountyCallerSession) Claimed() (bool, error) {
	return _InsecureTargetBounty.Contract.Claimed(&_InsecureTargetBounty.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InsecureTargetBounty *InsecureTargetBountyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _InsecureTargetBounty.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InsecureTargetBounty *InsecureTargetBountySession) Owner() (common.Address, error) {
	return _InsecureTargetBounty.Contract.Owner(&_InsecureTargetBounty.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InsecureTargetBounty *InsecureTargetBountyCallerSession) Owner() (common.Address, error) {
	return _InsecureTargetBounty.Contract.Owner(&_InsecureTargetBounty.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_InsecureTargetBounty *InsecureTargetBountyCaller) Payments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InsecureTargetBounty.contract.Call(opts, out, "payments", arg0)
	return *ret0, err
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_InsecureTargetBounty *InsecureTargetBountySession) Payments(arg0 common.Address) (*big.Int, error) {
	return _InsecureTargetBounty.Contract.Payments(&_InsecureTargetBounty.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_InsecureTargetBounty *InsecureTargetBountyCallerSession) Payments(arg0 common.Address) (*big.Int, error) {
	return _InsecureTargetBounty.Contract.Payments(&_InsecureTargetBounty.CallOpts, arg0)
}

// Researchers is a free data retrieval call binding the contract method 0x01bc237d.
//
// Solidity: function researchers( address) constant returns(address)
func (_InsecureTargetBounty *InsecureTargetBountyCaller) Researchers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _InsecureTargetBounty.contract.Call(opts, out, "researchers", arg0)
	return *ret0, err
}

// Researchers is a free data retrieval call binding the contract method 0x01bc237d.
//
// Solidity: function researchers( address) constant returns(address)
func (_InsecureTargetBounty *InsecureTargetBountySession) Researchers(arg0 common.Address) (common.Address, error) {
	return _InsecureTargetBounty.Contract.Researchers(&_InsecureTargetBounty.CallOpts, arg0)
}

// Researchers is a free data retrieval call binding the contract method 0x01bc237d.
//
// Solidity: function researchers( address) constant returns(address)
func (_InsecureTargetBounty *InsecureTargetBountyCallerSession) Researchers(arg0 common.Address) (common.Address, error) {
	return _InsecureTargetBounty.Contract.Researchers(&_InsecureTargetBounty.CallOpts, arg0)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_InsecureTargetBounty *InsecureTargetBountyCaller) TotalPayments(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InsecureTargetBounty.contract.Call(opts, out, "totalPayments")
	return *ret0, err
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_InsecureTargetBounty *InsecureTargetBountySession) TotalPayments() (*big.Int, error) {
	return _InsecureTargetBounty.Contract.TotalPayments(&_InsecureTargetBounty.CallOpts)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_InsecureTargetBounty *InsecureTargetBountyCallerSession) TotalPayments() (*big.Int, error) {
	return _InsecureTargetBounty.Contract.TotalPayments(&_InsecureTargetBounty.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(target address) returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactor) Claim(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.contract.Transact(opts, "claim", target)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(target address) returns()
func (_InsecureTargetBounty *InsecureTargetBountySession) Claim(target common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.Claim(&_InsecureTargetBounty.TransactOpts, target)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(target address) returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactorSession) Claim(target common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.Claim(&_InsecureTargetBounty.TransactOpts, target)
}

// CreateTarget is a paid mutator transaction binding the contract method 0xc98165b6.
//
// Solidity: function createTarget() returns(address)
func (_InsecureTargetBounty *InsecureTargetBountyTransactor) CreateTarget(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetBounty.contract.Transact(opts, "createTarget")
}

// CreateTarget is a paid mutator transaction binding the contract method 0xc98165b6.
//
// Solidity: function createTarget() returns(address)
func (_InsecureTargetBounty *InsecureTargetBountySession) CreateTarget() (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.CreateTarget(&_InsecureTargetBounty.TransactOpts)
}

// CreateTarget is a paid mutator transaction binding the contract method 0xc98165b6.
//
// Solidity: function createTarget() returns(address)
func (_InsecureTargetBounty *InsecureTargetBountyTransactorSession) CreateTarget() (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.CreateTarget(&_InsecureTargetBounty.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetBounty.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_InsecureTargetBounty *InsecureTargetBountySession) Destroy() (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.Destroy(&_InsecureTargetBounty.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactorSession) Destroy() (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.Destroy(&_InsecureTargetBounty.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_InsecureTargetBounty *InsecureTargetBountySession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.DestroyAndSend(&_InsecureTargetBounty.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.DestroyAndSend(&_InsecureTargetBounty.TransactOpts, _recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_InsecureTargetBounty *InsecureTargetBountySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.TransferOwnership(&_InsecureTargetBounty.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.TransferOwnership(&_InsecureTargetBounty.TransactOpts, newOwner)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactor) WithdrawPayments(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsecureTargetBounty.contract.Transact(opts, "withdrawPayments")
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_InsecureTargetBounty *InsecureTargetBountySession) WithdrawPayments() (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.WithdrawPayments(&_InsecureTargetBounty.TransactOpts)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_InsecureTargetBounty *InsecureTargetBountyTransactorSession) WithdrawPayments() (*types.Transaction, error) {
	return _InsecureTargetBounty.Contract.WithdrawPayments(&_InsecureTargetBounty.TransactOpts)
}