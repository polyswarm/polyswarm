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

// PostDeliveryCrowdsaleImplABI is the input ABI used to generate the binding from.
const PostDeliveryCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// PostDeliveryCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const PostDeliveryCrowdsaleImplBin = `6060604052341561000f57600080fd5b60405160a0806105e183398101604052808051919060200180519190602001805191906020018051919060200180519150859050848484846000831161005457600080fd5b600160a060020a038216151561006957600080fd5b600160a060020a038116151561007e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a03199182161790915560008054929093169116179055428210156100bd57600080fd5b818110156100ca57600080fd5b60049190915560055550505050506104fa806100e76000396000f3006060604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631515bc2b81146100ae57806327e235e3146100d55780632c4e722e146101065780634042b66f146101195780634b6753bc1461012c578063521eb2731461013f5780638d8f2adb1461016e578063b7a8807c14610181578063ec8ac4d814610194578063fc0c546a146101a8575b6100ac336101bb565b005b34156100b957600080fd5b6100c1610263565b604051901515815260200160405180910390f35b34156100e057600080fd5b6100f4600160a060020a036004351661026b565b60405190815260200160405180910390f35b341561011157600080fd5b6100f461027d565b341561012457600080fd5b6100f4610283565b341561013757600080fd5b6100f4610289565b341561014a57600080fd5b61015261028f565b604051600160a060020a03909116815260200160405180910390f35b341561017957600080fd5b6100ac61029e565b341561018c57600080fd5b6100f4610300565b6100ac600160a060020a03600435166101bb565b34156101b357600080fd5b610152610306565b3460006101c88383610315565b6101d182610342565b6003549091506101e7908363ffffffff61035f16565b6003556101f48382610379565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a361024c838361033e565b6102546103c2565b61025e838361033e565b505050565b600554421190565b60066020526000908152604090205481565b60025481565b60035481565b60055481565b600154600160a060020a031681565b60006102a8610263565b15156102b357600080fd5b50600160a060020a0333166000908152600660205260408120549081116102d957600080fd5b33600160a060020a0381166000908152600660205260408120556102fd90826103f8565b50565b60045481565b600054600160a060020a031681565b600454421015801561032957506005544211155b151561033457600080fd5b61033e8282610482565b5050565b6000610359600254836104a390919063ffffffff16565b92915050565b60008282018381101561036e57fe5b8091505b5092915050565b600160a060020a0382166000908152600660205260409020546103a2908263ffffffff61035f16565b600160a060020a0390921660009081526006602052604090209190915550565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f1935050505015156103f657600080fd5b565b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561046757600080fd5b5af1151561047457600080fd5b505050604051805150505050565b600160a060020a038216151561049757600080fd5b80151561033e57600080fd5b6000808315156104b65760009150610372565b508282028284828115156104c657fe5b041461036e57fe00a165627a7a7230582039ae2a9ae593cb89207ba6482ef878294003f73a6c885118acfddbca67e9d70d0029`

// DeployPostDeliveryCrowdsaleImpl deploys a new Ethereum contract, binding an instance of PostDeliveryCrowdsaleImpl to it.
func DeployPostDeliveryCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _openingTime *big.Int, _closingTime *big.Int, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *PostDeliveryCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(PostDeliveryCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PostDeliveryCrowdsaleImplBin), backend, _openingTime, _closingTime, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PostDeliveryCrowdsaleImpl{PostDeliveryCrowdsaleImplCaller: PostDeliveryCrowdsaleImplCaller{contract: contract}, PostDeliveryCrowdsaleImplTransactor: PostDeliveryCrowdsaleImplTransactor{contract: contract}}, nil
}

// PostDeliveryCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleImpl struct {
	PostDeliveryCrowdsaleImplCaller     // Read-only binding to the contract
	PostDeliveryCrowdsaleImplTransactor // Write-only binding to the contract
}

// PostDeliveryCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostDeliveryCrowdsaleImplSession struct {
	Contract     *PostDeliveryCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PostDeliveryCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostDeliveryCrowdsaleImplCallerSession struct {
	Contract *PostDeliveryCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PostDeliveryCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostDeliveryCrowdsaleImplTransactorSession struct {
	Contract     *PostDeliveryCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PostDeliveryCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleImplRaw struct {
	Contract *PostDeliveryCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// PostDeliveryCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleImplCallerRaw struct {
	Contract *PostDeliveryCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// PostDeliveryCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleImplTransactorRaw struct {
	Contract *PostDeliveryCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPostDeliveryCrowdsaleImpl creates a new instance of PostDeliveryCrowdsaleImpl, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*PostDeliveryCrowdsaleImpl, error) {
	contract, err := bindPostDeliveryCrowdsaleImpl(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleImpl{PostDeliveryCrowdsaleImplCaller: PostDeliveryCrowdsaleImplCaller{contract: contract}, PostDeliveryCrowdsaleImplTransactor: PostDeliveryCrowdsaleImplTransactor{contract: contract}}, nil
}

// NewPostDeliveryCrowdsaleImplCaller creates a new read-only instance of PostDeliveryCrowdsaleImpl, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*PostDeliveryCrowdsaleImplCaller, error) {
	contract, err := bindPostDeliveryCrowdsaleImpl(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleImplCaller{contract: contract}, nil
}

// NewPostDeliveryCrowdsaleImplTransactor creates a new write-only instance of PostDeliveryCrowdsaleImpl, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*PostDeliveryCrowdsaleImplTransactor, error) {
	contract, err := bindPostDeliveryCrowdsaleImpl(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleImplTransactor{contract: contract}, nil
}

// bindPostDeliveryCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindPostDeliveryCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PostDeliveryCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PostDeliveryCrowdsaleImpl.Contract.PostDeliveryCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.PostDeliveryCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.PostDeliveryCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PostDeliveryCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Balances(&_PostDeliveryCrowdsaleImpl.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Balances(&_PostDeliveryCrowdsaleImpl.CallOpts, arg0)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) ClosingTime() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.ClosingTime(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) ClosingTime() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.ClosingTime(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) HasClosed() (bool, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.HasClosed(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) HasClosed() (bool, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.HasClosed(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) OpeningTime() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.OpeningTime(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) OpeningTime() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.OpeningTime(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Rate(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Rate(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) Token() (common.Address, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Token(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Token(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Wallet(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.Wallet(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PostDeliveryCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.WeiRaised(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.WeiRaised(&_PostDeliveryCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.BuyTokens(&_PostDeliveryCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.BuyTokens(&_PostDeliveryCrowdsaleImpl.TransactOpts, _beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplTransactor) WithdrawTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.contract.Transact(opts, "withdrawTokens")
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplSession) WithdrawTokens() (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.WithdrawTokens(&_PostDeliveryCrowdsaleImpl.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_PostDeliveryCrowdsaleImpl *PostDeliveryCrowdsaleImplTransactorSession) WithdrawTokens() (*types.Transaction, error) {
	return _PostDeliveryCrowdsaleImpl.Contract.WithdrawTokens(&_PostDeliveryCrowdsaleImpl.TransactOpts)
}
