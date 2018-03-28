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

// HasNoTokensABI is the input ABI used to generate the binding from.
const HasNoTokensABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoTokensBin is the compiled bytecode used for deploying new contracts.
const HasNoTokensBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556102e0806100306000396000f3006060604052600436106100485763ffffffff60e060020a60003504166317ffc320811461004d5780638da5cb5b1461006e578063c0ee0b8a1461009d578063f2fde38b146100cc575b600080fd5b341561005857600080fd5b61006c600160a060020a03600435166100eb565b005b341561007957600080fd5b610081610192565b604051600160a060020a03909116815260200160405180910390f35b34156100a857600080fd5b61006c60048035600160a060020a0316906024803591604435918201910135610048565b34156100d757600080fd5b61006c600160a060020a03600435166101a1565b6000805433600160a060020a0390811691161461010757600080fd5b81600160a060020a03166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561015557600080fd5b5af1151561016257600080fd5b505050604051805160005490925061018e9150600160a060020a0384811691168363ffffffff61023c16565b5050565b600054600160a060020a031681565b60005433600160a060020a039081169116146101bc57600080fd5b600160a060020a03811615156101d157600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561029057600080fd5b5af1151561029d57600080fd5b5050506040518051905015156102af57fe5b5050505600a165627a7a723058202bcf87a699a55d83a6d8dc10ec52b688aeb196315a573b1103ac0b7efc0524c90029`

// DeployHasNoTokens deploys a new Ethereum contract, binding an instance of HasNoTokens to it.
func DeployHasNoTokens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoTokensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoTokens{HasNoTokensCaller: HasNoTokensCaller{contract: contract}, HasNoTokensTransactor: HasNoTokensTransactor{contract: contract}}, nil
}

// HasNoTokens is an auto generated Go binding around an Ethereum contract.
type HasNoTokens struct {
	HasNoTokensCaller     // Read-only binding to the contract
	HasNoTokensTransactor // Write-only binding to the contract
}

// HasNoTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoTokensSession struct {
	Contract     *HasNoTokens      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoTokensCallerSession struct {
	Contract *HasNoTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HasNoTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoTokensTransactorSession struct {
	Contract     *HasNoTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HasNoTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoTokensRaw struct {
	Contract *HasNoTokens // Generic contract binding to access the raw methods on
}

// HasNoTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoTokensCallerRaw struct {
	Contract *HasNoTokensCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoTokensTransactorRaw struct {
	Contract *HasNoTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoTokens creates a new instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokens(address common.Address, backend bind.ContractBackend) (*HasNoTokens, error) {
	contract, err := bindHasNoTokens(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoTokens{HasNoTokensCaller: HasNoTokensCaller{contract: contract}, HasNoTokensTransactor: HasNoTokensTransactor{contract: contract}}, nil
}

// NewHasNoTokensCaller creates a new read-only instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensCaller(address common.Address, caller bind.ContractCaller) (*HasNoTokensCaller, error) {
	contract, err := bindHasNoTokens(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensCaller{contract: contract}, nil
}

// NewHasNoTokensTransactor creates a new write-only instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoTokensTransactor, error) {
	contract, err := bindHasNoTokens(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensTransactor{contract: contract}, nil
}

// bindHasNoTokens binds a generic wrapper to an already deployed contract.
func bindHasNoTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoTokens *HasNoTokensRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoTokens.Contract.HasNoTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoTokens *HasNoTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoTokens.Contract.HasNoTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoTokens *HasNoTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoTokens.Contract.HasNoTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoTokens *HasNoTokensCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoTokens *HasNoTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoTokens *HasNoTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoTokens.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoTokens.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensSession) Owner() (common.Address, error) {
	return _HasNoTokens.Contract.Owner(&_HasNoTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensCallerSession) Owner() (common.Address, error) {
	return _HasNoTokens.Contract.Owner(&_HasNoTokens.CallOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.ReclaimToken(&_HasNoTokens.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.ReclaimToken(&_HasNoTokens.TransactOpts, token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TokenFallback(&_HasNoTokens.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TokenFallback(&_HasNoTokens.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TransferOwnership(&_HasNoTokens.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TransferOwnership(&_HasNoTokens.TransactOpts, newOwner)
}
