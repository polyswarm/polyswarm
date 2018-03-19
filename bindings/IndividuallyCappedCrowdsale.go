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

// IndividuallyCappedCrowdsaleABI is the input ABI used to generate the binding from.
const IndividuallyCappedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"caps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getUserCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaries\",\"type\":\"address[]\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setGroupCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getUserContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setUserCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// IndividuallyCappedCrowdsale is an auto generated Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsale struct {
	IndividuallyCappedCrowdsaleCaller     // Read-only binding to the contract
	IndividuallyCappedCrowdsaleTransactor // Write-only binding to the contract
	IndividuallyCappedCrowdsaleFilterer   // Log filterer for contract events
}

// IndividuallyCappedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndividuallyCappedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndividuallyCappedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndividuallyCappedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndividuallyCappedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndividuallyCappedCrowdsaleSession struct {
	Contract     *IndividuallyCappedCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IndividuallyCappedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndividuallyCappedCrowdsaleCallerSession struct {
	Contract *IndividuallyCappedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IndividuallyCappedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndividuallyCappedCrowdsaleTransactorSession struct {
	Contract     *IndividuallyCappedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IndividuallyCappedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleRaw struct {
	Contract *IndividuallyCappedCrowdsale // Generic contract binding to access the raw methods on
}

// IndividuallyCappedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleCallerRaw struct {
	Contract *IndividuallyCappedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// IndividuallyCappedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleTransactorRaw struct {
	Contract *IndividuallyCappedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndividuallyCappedCrowdsale creates a new instance of IndividuallyCappedCrowdsale, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsale(address common.Address, backend bind.ContractBackend) (*IndividuallyCappedCrowdsale, error) {
	contract, err := bindIndividuallyCappedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsale{IndividuallyCappedCrowdsaleCaller: IndividuallyCappedCrowdsaleCaller{contract: contract}, IndividuallyCappedCrowdsaleTransactor: IndividuallyCappedCrowdsaleTransactor{contract: contract}, IndividuallyCappedCrowdsaleFilterer: IndividuallyCappedCrowdsaleFilterer{contract: contract}}, nil
}

// NewIndividuallyCappedCrowdsaleCaller creates a new read-only instance of IndividuallyCappedCrowdsale, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*IndividuallyCappedCrowdsaleCaller, error) {
	contract, err := bindIndividuallyCappedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleCaller{contract: contract}, nil
}

// NewIndividuallyCappedCrowdsaleTransactor creates a new write-only instance of IndividuallyCappedCrowdsale, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*IndividuallyCappedCrowdsaleTransactor, error) {
	contract, err := bindIndividuallyCappedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleTransactor{contract: contract}, nil
}

// NewIndividuallyCappedCrowdsaleFilterer creates a new log filterer instance of IndividuallyCappedCrowdsale, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*IndividuallyCappedCrowdsaleFilterer, error) {
	contract, err := bindIndividuallyCappedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleFilterer{contract: contract}, nil
}

// bindIndividuallyCappedCrowdsale binds a generic wrapper to an already deployed contract.
func bindIndividuallyCappedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IndividuallyCappedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IndividuallyCappedCrowdsale.Contract.IndividuallyCappedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.IndividuallyCappedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.IndividuallyCappedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IndividuallyCappedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) Caps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "caps", arg0)
	return *ret0, err
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) Caps(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.Caps(&_IndividuallyCappedCrowdsale.CallOpts, arg0)
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) Caps(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.Caps(&_IndividuallyCappedCrowdsale.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "contributions", arg0)
	return *ret0, err
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.Contributions(&_IndividuallyCappedCrowdsale.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.Contributions(&_IndividuallyCappedCrowdsale.CallOpts, arg0)
}

// GetUserCap is a free data retrieval call binding the contract method 0x8b58c64c.
//
// Solidity: function getUserCap(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) GetUserCap(opts *bind.CallOpts, _beneficiary common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "getUserCap", _beneficiary)
	return *ret0, err
}

// GetUserCap is a free data retrieval call binding the contract method 0x8b58c64c.
//
// Solidity: function getUserCap(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) GetUserCap(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.GetUserCap(&_IndividuallyCappedCrowdsale.CallOpts, _beneficiary)
}

// GetUserCap is a free data retrieval call binding the contract method 0x8b58c64c.
//
// Solidity: function getUserCap(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) GetUserCap(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.GetUserCap(&_IndividuallyCappedCrowdsale.CallOpts, _beneficiary)
}

// GetUserContribution is a free data retrieval call binding the contract method 0xbb8b2b47.
//
// Solidity: function getUserContribution(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) GetUserContribution(opts *bind.CallOpts, _beneficiary common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "getUserContribution", _beneficiary)
	return *ret0, err
}

// GetUserContribution is a free data retrieval call binding the contract method 0xbb8b2b47.
//
// Solidity: function getUserContribution(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) GetUserContribution(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.GetUserContribution(&_IndividuallyCappedCrowdsale.CallOpts, _beneficiary)
}

// GetUserContribution is a free data retrieval call binding the contract method 0xbb8b2b47.
//
// Solidity: function getUserContribution(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) GetUserContribution(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.GetUserContribution(&_IndividuallyCappedCrowdsale.CallOpts, _beneficiary)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) Owner() (common.Address, error) {
	return _IndividuallyCappedCrowdsale.Contract.Owner(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _IndividuallyCappedCrowdsale.Contract.Owner(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) Rate() (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.Rate(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.Rate(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) Token() (common.Address, error) {
	return _IndividuallyCappedCrowdsale.Contract.Token(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _IndividuallyCappedCrowdsale.Contract.Token(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) Wallet() (common.Address, error) {
	return _IndividuallyCappedCrowdsale.Contract.Wallet(&_IndividuallyCappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _IndividuallyCappedCrowdsale.Contract.Wallet(&_IndividuallyCappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.WeiRaised(&_IndividuallyCappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _IndividuallyCappedCrowdsale.Contract.WeiRaised(&_IndividuallyCappedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.BuyTokens(&_IndividuallyCappedCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.BuyTokens(&_IndividuallyCappedCrowdsale.TransactOpts, _beneficiary)
}

// SetGroupCap is a paid mutator transaction binding the contract method 0xa31f61fc.
//
// Solidity: function setGroupCap(_beneficiaries address[], _cap uint256) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactor) SetGroupCap(opts *bind.TransactOpts, _beneficiaries []common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.contract.Transact(opts, "setGroupCap", _beneficiaries, _cap)
}

// SetGroupCap is a paid mutator transaction binding the contract method 0xa31f61fc.
//
// Solidity: function setGroupCap(_beneficiaries address[], _cap uint256) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) SetGroupCap(_beneficiaries []common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.SetGroupCap(&_IndividuallyCappedCrowdsale.TransactOpts, _beneficiaries, _cap)
}

// SetGroupCap is a paid mutator transaction binding the contract method 0xa31f61fc.
//
// Solidity: function setGroupCap(_beneficiaries address[], _cap uint256) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactorSession) SetGroupCap(_beneficiaries []common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.SetGroupCap(&_IndividuallyCappedCrowdsale.TransactOpts, _beneficiaries, _cap)
}

// SetUserCap is a paid mutator transaction binding the contract method 0xc3143fe5.
//
// Solidity: function setUserCap(_beneficiary address, _cap uint256) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactor) SetUserCap(opts *bind.TransactOpts, _beneficiary common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.contract.Transact(opts, "setUserCap", _beneficiary, _cap)
}

// SetUserCap is a paid mutator transaction binding the contract method 0xc3143fe5.
//
// Solidity: function setUserCap(_beneficiary address, _cap uint256) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) SetUserCap(_beneficiary common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.SetUserCap(&_IndividuallyCappedCrowdsale.TransactOpts, _beneficiary, _cap)
}

// SetUserCap is a paid mutator transaction binding the contract method 0xc3143fe5.
//
// Solidity: function setUserCap(_beneficiary address, _cap uint256) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactorSession) SetUserCap(_beneficiary common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.SetUserCap(&_IndividuallyCappedCrowdsale.TransactOpts, _beneficiary, _cap)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.TransferOwnership(&_IndividuallyCappedCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsale.Contract.TransferOwnership(&_IndividuallyCappedCrowdsale.TransactOpts, newOwner)
}

// IndividuallyCappedCrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IndividuallyCappedCrowdsale contract.
type IndividuallyCappedCrowdsaleOwnershipTransferredIterator struct {
	Event *IndividuallyCappedCrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IndividuallyCappedCrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndividuallyCappedCrowdsaleOwnershipTransferred)
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
		it.Event = new(IndividuallyCappedCrowdsaleOwnershipTransferred)
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
func (it *IndividuallyCappedCrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndividuallyCappedCrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndividuallyCappedCrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the IndividuallyCappedCrowdsale contract.
type IndividuallyCappedCrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IndividuallyCappedCrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IndividuallyCappedCrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleOwnershipTransferredIterator{contract: _IndividuallyCappedCrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IndividuallyCappedCrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IndividuallyCappedCrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndividuallyCappedCrowdsaleOwnershipTransferred)
				if err := _IndividuallyCappedCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// IndividuallyCappedCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the IndividuallyCappedCrowdsale contract.
type IndividuallyCappedCrowdsaleTokenPurchaseIterator struct {
	Event *IndividuallyCappedCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

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
func (it *IndividuallyCappedCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndividuallyCappedCrowdsaleTokenPurchase)
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
		it.Event = new(IndividuallyCappedCrowdsaleTokenPurchase)
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
func (it *IndividuallyCappedCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndividuallyCappedCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndividuallyCappedCrowdsaleTokenPurchase represents a TokenPurchase event raised by the IndividuallyCappedCrowdsale contract.
type IndividuallyCappedCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*IndividuallyCappedCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IndividuallyCappedCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleTokenPurchaseIterator{contract: _IndividuallyCappedCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_IndividuallyCappedCrowdsale *IndividuallyCappedCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *IndividuallyCappedCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IndividuallyCappedCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndividuallyCappedCrowdsaleTokenPurchase)
				if err := _IndividuallyCappedCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
