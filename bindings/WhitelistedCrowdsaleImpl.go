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

// WhitelistedCrowdsaleImplABI is the input ABI used to generate the binding from.
const WhitelistedCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaries\",\"type\":\"address[]\"}],\"name\":\"addManyToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// WhitelistedCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const WhitelistedCrowdsaleImplBin = `6060604052341561000f57600080fd5b6040516060806106fe833981016040528080519190602001805191906020018051915083905082826000831161004457600080fd5b600160a060020a038216151561005957600080fd5b600160a060020a038116151561006e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a0319918216179091556000805493831693821693909317909255600480543390921691909216179055505050610639806100c56000396000f3006060604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e81146100b95780634042b66f146100de578063521eb273146100f15780638ab1d681146101205780638c10671c1461013f5780638da5cb5b1461015d5780639b19251a14610170578063e43252d7146101a3578063ec8ac4d8146101c2578063f2fde38b146101d6578063fc0c546a146101f5575b6100b733610208565b005b34156100c457600080fd5b6100cc6102b0565b60405190815260200160405180910390f35b34156100e957600080fd5b6100cc6102b6565b34156100fc57600080fd5b6101046102bc565b604051600160a060020a03909116815260200160405180910390f35b341561012b57600080fd5b6100b7600160a060020a03600435166102cb565b341561014a57600080fd5b6100b76004803560248101910135610307565b341561016857600080fd5b61010461037c565b341561017b57600080fd5b61018f600160a060020a036004351661038b565b604051901515815260200160405180910390f35b34156101ae57600080fd5b6100b7600160a060020a03600435166103a0565b6100b7600160a060020a0360043516610208565b34156101e157600080fd5b6100b7600160a060020a03600435166103df565b341561020057600080fd5b61010461047a565b3460006102158383610489565b61021e826104bc565b600354909150610234908363ffffffff6104d916565b60035561024183826104f3565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a361029983836104fd565b6102a1610501565b6102ab83836104fd565b505050565b60025481565b60035481565b600154600160a060020a031681565b60045433600160a060020a039081169116146102e657600080fd5b600160a060020a03166000908152600560205260409020805460ff19169055565b60045460009033600160a060020a0390811691161461032557600080fd5b5060005b818110156102ab5760016005600085858581811061034357fe5b60209081029290920135600160a060020a0316835250810191909152604001600020805460ff1916911515919091179055600101610329565b600454600160a060020a031681565b60056020526000908152604090205460ff1681565b60045433600160a060020a039081169116146103bb57600080fd5b600160a060020a03166000908152600560205260409020805460ff19166001179055565b60045433600160a060020a039081169116146103fa57600080fd5b600160a060020a038116151561040f57600080fd5b600454600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600160a060020a038216600090815260056020526040902054829060ff1615156104b257600080fd5b6102ab8383610537565b60006104d36002548361055890919063ffffffff16565b92915050565b6000828201838110156104e857fe5b8091505b5092915050565b6104fd8282610583565b5050565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f19350505050151561053557600080fd5b565b600160a060020a038216151561054c57600080fd5b8015156104fd57600080fd5b60008083151561056b57600091506104ec565b5082820282848281151561057b57fe5b04146104e857fe5b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156105f257600080fd5b5af115156105ff57600080fd5b5050506040518051505050505600a165627a7a7230582072913b7851148099c18905bfba6229dbd2d5bbfbcf7d52e154e79a8e0974bc630029`

// DeployWhitelistedCrowdsaleImpl deploys a new Ethereum contract, binding an instance of WhitelistedCrowdsaleImpl to it.
func DeployWhitelistedCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *WhitelistedCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistedCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistedCrowdsaleImplBin), backend, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistedCrowdsaleImpl{WhitelistedCrowdsaleImplCaller: WhitelistedCrowdsaleImplCaller{contract: contract}, WhitelistedCrowdsaleImplTransactor: WhitelistedCrowdsaleImplTransactor{contract: contract}}, nil
}

// WhitelistedCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type WhitelistedCrowdsaleImpl struct {
	WhitelistedCrowdsaleImplCaller     // Read-only binding to the contract
	WhitelistedCrowdsaleImplTransactor // Write-only binding to the contract
}

// WhitelistedCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistedCrowdsaleImplSession struct {
	Contract     *WhitelistedCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WhitelistedCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistedCrowdsaleImplCallerSession struct {
	Contract *WhitelistedCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// WhitelistedCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistedCrowdsaleImplTransactorSession struct {
	Contract     *WhitelistedCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// WhitelistedCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistedCrowdsaleImplRaw struct {
	Contract *WhitelistedCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// WhitelistedCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleImplCallerRaw struct {
	Contract *WhitelistedCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistedCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistedCrowdsaleImplTransactorRaw struct {
	Contract *WhitelistedCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistedCrowdsaleImpl creates a new instance of WhitelistedCrowdsaleImpl, bound to a specific deployed contract.
func NewWhitelistedCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*WhitelistedCrowdsaleImpl, error) {
	contract, err := bindWhitelistedCrowdsaleImpl(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleImpl{WhitelistedCrowdsaleImplCaller: WhitelistedCrowdsaleImplCaller{contract: contract}, WhitelistedCrowdsaleImplTransactor: WhitelistedCrowdsaleImplTransactor{contract: contract}}, nil
}

// NewWhitelistedCrowdsaleImplCaller creates a new read-only instance of WhitelistedCrowdsaleImpl, bound to a specific deployed contract.
func NewWhitelistedCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*WhitelistedCrowdsaleImplCaller, error) {
	contract, err := bindWhitelistedCrowdsaleImpl(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleImplCaller{contract: contract}, nil
}

// NewWhitelistedCrowdsaleImplTransactor creates a new write-only instance of WhitelistedCrowdsaleImpl, bound to a specific deployed contract.
func NewWhitelistedCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistedCrowdsaleImplTransactor, error) {
	contract, err := bindWhitelistedCrowdsaleImpl(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WhitelistedCrowdsaleImplTransactor{contract: contract}, nil
}

// bindWhitelistedCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindWhitelistedCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistedCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistedCrowdsaleImpl.Contract.WhitelistedCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.WhitelistedCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.WhitelistedCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistedCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistedCrowdsaleImpl.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) Owner() (common.Address, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Owner(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerSession) Owner() (common.Address, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Owner(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WhitelistedCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Rate(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Rate(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistedCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) Token() (common.Address, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Token(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Token(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistedCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Wallet(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Wallet(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WhitelistedCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _WhitelistedCrowdsaleImpl.Contract.WeiRaised(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _WhitelistedCrowdsaleImpl.Contract.WeiRaised(&_WhitelistedCrowdsaleImpl.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistedCrowdsaleImpl.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) Whitelist(arg0 common.Address) (bool, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Whitelist(&_WhitelistedCrowdsaleImpl.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _WhitelistedCrowdsaleImpl.Contract.Whitelist(&_WhitelistedCrowdsaleImpl.CallOpts, arg0)
}

// AddManyToWhitelist is a paid mutator transaction binding the contract method 0x8c10671c.
//
// Solidity: function addManyToWhitelist(_beneficiaries address[]) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactor) AddManyToWhitelist(opts *bind.TransactOpts, _beneficiaries []common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.contract.Transact(opts, "addManyToWhitelist", _beneficiaries)
}

// AddManyToWhitelist is a paid mutator transaction binding the contract method 0x8c10671c.
//
// Solidity: function addManyToWhitelist(_beneficiaries address[]) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) AddManyToWhitelist(_beneficiaries []common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.AddManyToWhitelist(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiaries)
}

// AddManyToWhitelist is a paid mutator transaction binding the contract method 0x8c10671c.
//
// Solidity: function addManyToWhitelist(_beneficiaries address[]) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorSession) AddManyToWhitelist(_beneficiaries []common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.AddManyToWhitelist(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiaries)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactor) AddToWhitelist(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.contract.Transact(opts, "addToWhitelist", _beneficiary)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) AddToWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.AddToWhitelist(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorSession) AddToWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.AddToWhitelist(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.BuyTokens(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.BuyTokens(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.contract.Transact(opts, "removeFromWhitelist", _beneficiary)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) RemoveFromWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.RemoveFromWhitelist(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_beneficiary address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorSession) RemoveFromWhitelist(_beneficiary common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.RemoveFromWhitelist(&_WhitelistedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.TransferOwnership(&_WhitelistedCrowdsaleImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistedCrowdsaleImpl *WhitelistedCrowdsaleImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistedCrowdsaleImpl.Contract.TransferOwnership(&_WhitelistedCrowdsaleImpl.TransactOpts, newOwner)
}
