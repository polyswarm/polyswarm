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

// IndividuallyCappedCrowdsaleABI is the input ABI used to generate the binding from.
const IndividuallyCappedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"caps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getUserCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaries\",\"type\":\"address[]\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setGroupCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getUserContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setUserCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// IndividuallyCappedCrowdsale is an auto generated Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsale struct {
	IndividuallyCappedCrowdsaleCaller     // Read-only binding to the contract
	IndividuallyCappedCrowdsaleTransactor // Write-only binding to the contract
}

// IndividuallyCappedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndividuallyCappedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleTransactor struct {
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
	contract, err := bindIndividuallyCappedCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsale{IndividuallyCappedCrowdsaleCaller: IndividuallyCappedCrowdsaleCaller{contract: contract}, IndividuallyCappedCrowdsaleTransactor: IndividuallyCappedCrowdsaleTransactor{contract: contract}}, nil
}

// NewIndividuallyCappedCrowdsaleCaller creates a new read-only instance of IndividuallyCappedCrowdsale, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*IndividuallyCappedCrowdsaleCaller, error) {
	contract, err := bindIndividuallyCappedCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleCaller{contract: contract}, nil
}

// NewIndividuallyCappedCrowdsaleTransactor creates a new write-only instance of IndividuallyCappedCrowdsale, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*IndividuallyCappedCrowdsaleTransactor, error) {
	contract, err := bindIndividuallyCappedCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleTransactor{contract: contract}, nil
}

// bindIndividuallyCappedCrowdsale binds a generic wrapper to an already deployed contract.
func bindIndividuallyCappedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IndividuallyCappedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
