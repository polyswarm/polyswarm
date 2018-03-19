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

// WhitelistedCrowdsaleABI is the input ABI used to generate the binding from.
const WhitelistedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaries\",\"type\":\"address[]\"}],\"name\":\"addManyToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// WhitelistedCrowdsale is an auto generated Go binding around an Ethereum contract.
type WhitelistedCrowdsale struct {
	WhitelistedCrowdsaleCaller     // Read-only binding to the contract
	WhitelistedCrowdsaleTransactor // Write-only binding to the contract
	WhitelistedCrowdsaleFilterer   // Log filterer for contract events
}

// WhitelistedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistedCrowdsaleSession struct {
	Contract     *WhitelistedCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WhitelistedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistedCrowdsaleCallerSession struct {
	Contract *WhitelistedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// WhitelistedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistedCrowdsaleTransactorSession struct {
	Contract     *WhitelistedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// WhitelistedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistedCrowdsaleRaw struct {
	Contract *WhitelistedCrowdsale // Generic contract binding to access the raw methods on
}

// WhitelistedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleCallerRaw struct {
	Contract *WhitelistedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleTransactorRaw struct {
	Contract *WhitelistedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistedCrowdsale creates a new instance of WhitelistedCrowdsale, bound to a specific deployed contract.
func NewWhitelistedCrowdsale(address common.Address, backend bind.ContractBackend) (*WhitelistedCrowdsale, error) {
	contract, err := bindWhitelistedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsale{WhitelistedCrowdsaleCaller: WhitelistedCrowdsaleCaller{contract: contract}, WhitelistedCrowdsaleTransactor: WhitelistedCrowdsaleTransactor{contract: contract}, WhitelistedCrowdsaleFilterer: WhitelistedCrowdsaleFilterer{contract: contract}}, nil
}

// NewWhitelistedCrowdsaleCaller creates a new read-only instance of WhitelistedCrowdsale, bound to a specific deployed contract.
func NewWhitelistedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*WhitelistedCrowdsaleCaller, error) {
	contract, err := bindWhitelistedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleCaller{contract: contract}, nil
}

// NewWhitelistedCrowdsaleTransactor creates a new write-only instance of WhitelistedCrowdsale, bound to a specific deployed contract.
func NewWhitelistedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistedCrowdsaleTransactor, error) {
	contract, err := bindWhitelistedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleTransactor{contract: contract}, nil
}

// NewWhitelistedCrowdsaleFilterer creates a new log filterer instance of WhitelistedCrowdsale, bound to a specific deployed contract.
func NewWhitelistedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistedCrowdsaleFilterer, error) {
	contract, err := bindWhitelistedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleFilterer{contract: contract}, nil
}

// bindWhitelistedCrowdsale binds a generic wrapper to an already deployed contract.
func bindWhitelistedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistedCrowdsale.Contract.WhitelistedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.WhitelistedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.WhitelistedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistedCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) Owner() (common.Address, error) {
	return _WhitelistedCrowdsale.Contract.Owner(&_WhitelistedCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _WhitelistedCrowdsale.Contract.Owner(&_WhitelistedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WhitelistedCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) Rate() (*big.Int, error) {
	return _WhitelistedCrowdsale.Contract.Rate(&_WhitelistedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _WhitelistedCrowdsale.Contract.Rate(&_WhitelistedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistedCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) Token() (common.Address, error) {
	return _WhitelistedCrowdsale.Contract.Token(&_WhitelistedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _WhitelistedCrowdsale.Contract.Token(&_WhitelistedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistedCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) Wallet() (common.Address, error) {
	return _WhitelistedCrowdsale.Contract.Wallet(&_WhitelistedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _WhitelistedCrowdsale.Contract.Wallet(&_WhitelistedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WhitelistedCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _WhitelistedCrowdsale.Contract.WeiRaised(&_WhitelistedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _WhitelistedCrowdsale.Contract.WeiRaised(&_WhitelistedCrowdsale.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistedCrowdsale.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) Whitelist(arg0 common.Address) (bool, error) {
	return _WhitelistedCrowdsale.Contract.Whitelist(&_WhitelistedCrowdsale.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _WhitelistedCrowdsale.Contract.Whitelist(&_WhitelistedCrowdsale.CallOpts, arg0)
}

// AddManyToWhitelist is a paid mutator transaction binding the contract method 0x8c10671c.
//
// Solidity: function addManyToWhitelist(_beneficiaries address[]) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactor) AddManyToWhitelist(opts *bind.TransactOpts, _beneficiaries []common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.contract.Transact(opts, "addManyToWhitelist", _beneficiaries)
}

// AddManyToWhitelist is a paid mutator transaction binding the contract method 0x8c10671c.
//
// Solidity: function addManyToWhitelist(_beneficiaries address[]) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) AddManyToWhitelist(_beneficiaries []common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.AddManyToWhitelist(&_WhitelistedCrowdsale.TransactOpts, _beneficiaries)
}

// AddManyToWhitelist is a paid mutator transaction binding the contract method 0x8c10671c.
//
// Solidity: function addManyToWhitelist(_beneficiaries address[]) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorSession) AddManyToWhitelist(_beneficiaries []common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.AddManyToWhitelist(&_WhitelistedCrowdsale.TransactOpts, _beneficiaries)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactor) AddToWhitelist(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.contract.Transact(opts, "addToWhitelist", _beneficiary)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) AddToWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.AddToWhitelist(&_WhitelistedCrowdsale.TransactOpts, _beneficiary)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorSession) AddToWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.AddToWhitelist(&_WhitelistedCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.BuyTokens(&_WhitelistedCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.BuyTokens(&_WhitelistedCrowdsale.TransactOpts, _beneficiary)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.contract.Transact(opts, "removeFromWhitelist", _beneficiary)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) RemoveFromWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.RemoveFromWhitelist(&_WhitelistedCrowdsale.TransactOpts, _beneficiary)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorSession) RemoveFromWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.RemoveFromWhitelist(&_WhitelistedCrowdsale.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.TransferOwnership(&_WhitelistedCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsale.Contract.TransferOwnership(&_WhitelistedCrowdsale.TransactOpts, newOwner)
}

// WhitelistedCrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WhitelistedCrowdsale contract.
type WhitelistedCrowdsaleOwnershipTransferredIterator struct {
	Event *WhitelistedCrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WhitelistedCrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistedCrowdsaleOwnershipTransferred)
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
		it.Event = new(WhitelistedCrowdsaleOwnershipTransferred)
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
func (it *WhitelistedCrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistedCrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistedCrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the WhitelistedCrowdsale contract.
type WhitelistedCrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WhitelistedCrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WhitelistedCrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleOwnershipTransferredIterator{contract: _WhitelistedCrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WhitelistedCrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WhitelistedCrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistedCrowdsaleOwnershipTransferred)
				if err := _WhitelistedCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// WhitelistedCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the WhitelistedCrowdsale contract.
type WhitelistedCrowdsaleTokenPurchaseIterator struct {
	Event *WhitelistedCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

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
func (it *WhitelistedCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistedCrowdsaleTokenPurchase)
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
		it.Event = new(WhitelistedCrowdsaleTokenPurchase)
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
func (it *WhitelistedCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistedCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistedCrowdsaleTokenPurchase represents a TokenPurchase event raised by the WhitelistedCrowdsale contract.
type WhitelistedCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*WhitelistedCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _WhitelistedCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleTokenPurchaseIterator{contract: _WhitelistedCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_WhitelistedCrowdsale *WhitelistedCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *WhitelistedCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _WhitelistedCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistedCrowdsaleTokenPurchase)
				if err := _WhitelistedCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
