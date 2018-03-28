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

// TokenDestructibleABI is the input ABI used to generate the binding from.
const TokenDestructibleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenDestructibleBin is the compiled bytecode used for deploying new contracts.
const TokenDestructibleBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556102ee806100306000396000f30060606040526004361061003d5763ffffffff60e060020a6000350416638da5cb5b8114610042578063c6786e5a14610071578063f2fde38b146100c2575b600080fd5b341561004d57600080fd5b6100556100e1565b604051600160a060020a03909116815260200160405180910390f35b341561007c57600080fd5b6100c060046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437509496506100f095505050505050565b005b34156100cd57600080fd5b6100c0600160a060020a0360043516610227565b600054600160a060020a031681565b600080548190819033600160a060020a0390811691161461011057600080fd5b600092505b83518310156102195783838151811061012a57fe5b90602001906020020151915081600160a060020a03166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561018457600080fd5b5af1151561019157600080fd5b5050506040518051600054909250600160a060020a03808516925063a9059cbb91168360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156101f757600080fd5b5af1151561020457600080fd5b50505060405180515050600190920191610115565b600054600160a060020a0316ff5b60005433600160a060020a0390811691161461024257600080fd5b600160a060020a038116151561025757600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820a95bafa17a08ec1649669f384877f38bae43a1f89d573028f1a0f5d551dd1a820029`

// DeployTokenDestructible deploys a new Ethereum contract, binding an instance of TokenDestructible to it.
func DeployTokenDestructible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenDestructible, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDestructibleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenDestructibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenDestructible{TokenDestructibleCaller: TokenDestructibleCaller{contract: contract}, TokenDestructibleTransactor: TokenDestructibleTransactor{contract: contract}}, nil
}

// TokenDestructible is an auto generated Go binding around an Ethereum contract.
type TokenDestructible struct {
	TokenDestructibleCaller     // Read-only binding to the contract
	TokenDestructibleTransactor // Write-only binding to the contract
}

// TokenDestructibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenDestructibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDestructibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenDestructibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDestructibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenDestructibleSession struct {
	Contract     *TokenDestructible // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TokenDestructibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenDestructibleCallerSession struct {
	Contract *TokenDestructibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TokenDestructibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenDestructibleTransactorSession struct {
	Contract     *TokenDestructibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TokenDestructibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenDestructibleRaw struct {
	Contract *TokenDestructible // Generic contract binding to access the raw methods on
}

// TokenDestructibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenDestructibleCallerRaw struct {
	Contract *TokenDestructibleCaller // Generic read-only contract binding to access the raw methods on
}

// TokenDestructibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenDestructibleTransactorRaw struct {
	Contract *TokenDestructibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenDestructible creates a new instance of TokenDestructible, bound to a specific deployed contract.
func NewTokenDestructible(address common.Address, backend bind.ContractBackend) (*TokenDestructible, error) {
	contract, err := bindTokenDestructible(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenDestructible{TokenDestructibleCaller: TokenDestructibleCaller{contract: contract}, TokenDestructibleTransactor: TokenDestructibleTransactor{contract: contract}}, nil
}

// NewTokenDestructibleCaller creates a new read-only instance of TokenDestructible, bound to a specific deployed contract.
func NewTokenDestructibleCaller(address common.Address, caller bind.ContractCaller) (*TokenDestructibleCaller, error) {
	contract, err := bindTokenDestructible(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDestructibleCaller{contract: contract}, nil
}

// NewTokenDestructibleTransactor creates a new write-only instance of TokenDestructible, bound to a specific deployed contract.
func NewTokenDestructibleTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenDestructibleTransactor, error) {
	contract, err := bindTokenDestructible(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenDestructibleTransactor{contract: contract}, nil
}

// bindTokenDestructible binds a generic wrapper to an already deployed contract.
func bindTokenDestructible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDestructibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDestructible *TokenDestructibleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenDestructible.Contract.TokenDestructibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDestructible *TokenDestructibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDestructible.Contract.TokenDestructibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDestructible *TokenDestructibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDestructible.Contract.TokenDestructibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDestructible *TokenDestructibleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenDestructible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDestructible *TokenDestructibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDestructible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDestructible *TokenDestructibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDestructible.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenDestructible *TokenDestructibleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenDestructible.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenDestructible *TokenDestructibleSession) Owner() (common.Address, error) {
	return _TokenDestructible.Contract.Owner(&_TokenDestructible.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenDestructible *TokenDestructibleCallerSession) Owner() (common.Address, error) {
	return _TokenDestructible.Contract.Owner(&_TokenDestructible.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0xc6786e5a.
//
// Solidity: function destroy(tokens address[]) returns()
func (_TokenDestructible *TokenDestructibleTransactor) Destroy(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _TokenDestructible.contract.Transact(opts, "destroy", tokens)
}

// Destroy is a paid mutator transaction binding the contract method 0xc6786e5a.
//
// Solidity: function destroy(tokens address[]) returns()
func (_TokenDestructible *TokenDestructibleSession) Destroy(tokens []common.Address) (*types.Transaction, error) {
	return _TokenDestructible.Contract.Destroy(&_TokenDestructible.TransactOpts, tokens)
}

// Destroy is a paid mutator transaction binding the contract method 0xc6786e5a.
//
// Solidity: function destroy(tokens address[]) returns()
func (_TokenDestructible *TokenDestructibleTransactorSession) Destroy(tokens []common.Address) (*types.Transaction, error) {
	return _TokenDestructible.Contract.Destroy(&_TokenDestructible.TransactOpts, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenDestructible *TokenDestructibleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenDestructible.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenDestructible *TokenDestructibleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenDestructible.Contract.TransferOwnership(&_TokenDestructible.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenDestructible *TokenDestructibleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenDestructible.Contract.TransferOwnership(&_TokenDestructible.TransactOpts, newOwner)
}
