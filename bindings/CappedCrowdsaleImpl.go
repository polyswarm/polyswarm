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

// CappedCrowdsaleImplABI is the input ABI used to generate the binding from.
const CappedCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// CappedCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const CappedCrowdsaleImplBin = `6060604052341561000f57600080fd5b6040516080806104a98339810160405280805191906020018051919060200180519190602001805191508190508484846000831161004c57600080fd5b600160a060020a038216151561006157600080fd5b600160a060020a038116151561007657600080fd5b60029290925560018054600160a060020a03928316600160a060020a0319918216179091556000805492909316911617815581116100b357600080fd5b600455505050506103e0806100c96000396000f3006060604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e811461008d578063355274ea146100b25780634042b66f146100c55780634f935945146100d8578063521eb273146100ff578063ec8ac4d81461012e578063fc0c546a14610142575b61008b33610155565b005b341561009857600080fd5b6100a06101fd565b60405190815260200160405180910390f35b34156100bd57600080fd5b6100a0610203565b34156100d057600080fd5b6100a0610209565b34156100e357600080fd5b6100eb61020f565b604051901515815260200160405180910390f35b341561010a57600080fd5b61011261021a565b604051600160a060020a03909116815260200160405180910390f35b61008b600160a060020a0360043516610155565b341561014d57600080fd5b610112610229565b3460006101628383610238565b61016b82610267565b600354909150610181908363ffffffff61028416565b60035561018e838261029e565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36101e68383610263565b6101ee6102a8565b6101f88383610263565b505050565b60025481565b60045481565b60035481565b600454600354101590565b600154600160a060020a031681565b600054600160a060020a031681565b61024282826102de565b600454600354610258908363ffffffff61028416565b111561026357600080fd5b5050565b600061027e600254836102ff90919063ffffffff16565b92915050565b60008282018381101561029357fe5b8091505b5092915050565b610263828261032a565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f1935050505015156102dc57600080fd5b565b600160a060020a03821615156102f357600080fd5b80151561026357600080fd5b6000808315156103125760009150610297565b5082820282848281151561032257fe5b041461029357fe5b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561039957600080fd5b5af115156103a657600080fd5b5050506040518051505050505600a165627a7a72305820407490582eb335b8dd504925a7250e5b393ddf1f6fea8c6b0e3ae70e14b6dae60029`

// DeployCappedCrowdsaleImpl deploys a new Ethereum contract, binding an instance of CappedCrowdsaleImpl to it.
func DeployCappedCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, _token common.Address, _cap *big.Int) (common.Address, *types.Transaction, *CappedCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CappedCrowdsaleImplBin), backend, _rate, _wallet, _token, _cap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CappedCrowdsaleImpl{CappedCrowdsaleImplCaller: CappedCrowdsaleImplCaller{contract: contract}, CappedCrowdsaleImplTransactor: CappedCrowdsaleImplTransactor{contract: contract}}, nil
}

// CappedCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type CappedCrowdsaleImpl struct {
	CappedCrowdsaleImplCaller     // Read-only binding to the contract
	CappedCrowdsaleImplTransactor // Write-only binding to the contract
}

// CappedCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedCrowdsaleImplSession struct {
	Contract     *CappedCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CappedCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedCrowdsaleImplCallerSession struct {
	Contract *CappedCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CappedCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedCrowdsaleImplTransactorSession struct {
	Contract     *CappedCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CappedCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedCrowdsaleImplRaw struct {
	Contract *CappedCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// CappedCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplCallerRaw struct {
	Contract *CappedCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// CappedCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplTransactorRaw struct {
	Contract *CappedCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedCrowdsaleImpl creates a new instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*CappedCrowdsaleImpl, error) {
	contract, err := bindCappedCrowdsaleImpl(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImpl{CappedCrowdsaleImplCaller: CappedCrowdsaleImplCaller{contract: contract}, CappedCrowdsaleImplTransactor: CappedCrowdsaleImplTransactor{contract: contract}}, nil
}

// NewCappedCrowdsaleImplCaller creates a new read-only instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*CappedCrowdsaleImplCaller, error) {
	contract, err := bindCappedCrowdsaleImpl(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImplCaller{contract: contract}, nil
}

// NewCappedCrowdsaleImplTransactor creates a new write-only instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedCrowdsaleImplTransactor, error) {
	contract, err := bindCappedCrowdsaleImpl(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImplTransactor{contract: contract}, nil
}

// bindCappedCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindCappedCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsaleImpl.Contract.CappedCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.CappedCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.CappedCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Cap() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Cap(&_CappedCrowdsaleImpl.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Cap() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Cap(&_CappedCrowdsaleImpl.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "capReached")
	return *ret0, err
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) CapReached() (bool, error) {
	return _CappedCrowdsaleImpl.Contract.CapReached(&_CappedCrowdsaleImpl.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) CapReached() (bool, error) {
	return _CappedCrowdsaleImpl.Contract.CapReached(&_CappedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Rate(&_CappedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Rate(&_CappedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Token() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Token(&_CappedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Token(&_CappedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Wallet(&_CappedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Wallet(&_CappedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.WeiRaised(&_CappedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.WeiRaised(&_CappedCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.BuyTokens(&_CappedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.BuyTokens(&_CappedCrowdsaleImpl.TransactOpts, _beneficiary)
}
