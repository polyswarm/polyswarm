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

// AllowanceCrowdsaleImplABI is the input ABI used to generate the binding from.
const AllowanceCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_tokenWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// AllowanceCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const AllowanceCrowdsaleImplBin = `6060604052341561000f57600080fd5b6040516080806104e98339810160405280805191906020018051919060200180519190602001805191508190508484846000831161004c57600080fd5b600160a060020a038216151561006157600080fd5b600160a060020a038116151561007657600080fd5b60029290925560018054600160a060020a0319908116600160a060020a03938416179091556000805490911692821692909217909155811615156100b957600080fd5b60048054600160a060020a03909216600160a060020a0319909216919091179055505050506103fc806100ed6000396000f3006060604052600436106100695763ffffffff60e060020a6000350416632c4e722e81146100745780634042b66f14610099578063521eb273146100ac578063bf583903146100db578063bff99c6c146100ee578063ec8ac4d814610101578063fc0c546a14610115575b61007233610128565b005b341561007f57600080fd5b6100876101d0565b60405190815260200160405180910390f35b34156100a457600080fd5b6100876101d6565b34156100b757600080fd5b6100bf6101dc565b604051600160a060020a03909116815260200160405180910390f35b34156100e657600080fd5b6100876101eb565b34156100f957600080fd5b6100bf610266565b610072600160a060020a0360043516610128565b341561012057600080fd5b6100bf610275565b3460006101358383610284565b61013e826102a9565b600354909150610154908363ffffffff6102c616565b60035561016183826102e0565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36101b983836102a5565b6101c16102ea565b6101cb83836102a5565b505050565b60025481565b60035481565b600154600160a060020a031681565b60008054600454600160a060020a039182169163dd62ed3e91163060405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561024b57600080fd5b5af1151561025857600080fd5b505050604051805191505090565b600454600160a060020a031681565b600054600160a060020a031681565b600160a060020a038216151561029957600080fd5b8015156102a557600080fd5b5050565b60006102c06002548361032090919063ffffffff16565b92915050565b6000828201838110156102d557fe5b8091505b5092915050565b6102a5828261034b565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f19350505050151561031e57600080fd5b565b60008083151561033357600091506102d9565b5082820282848281151561034357fe5b04146102d557fe5b600054600454600160a060020a03918216916323b872dd9116848460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156103b557600080fd5b5af115156103c257600080fd5b5050506040518051505050505600a165627a7a72305820f08edcb6bcaf6937f466c7b8a3371ebfd8c83b5b725e1dbd74b63b99df4363f20029`

// DeployAllowanceCrowdsaleImpl deploys a new Ethereum contract, binding an instance of AllowanceCrowdsaleImpl to it.
func DeployAllowanceCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, _token common.Address, _tokenWallet common.Address) (common.Address, *types.Transaction, *AllowanceCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(AllowanceCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AllowanceCrowdsaleImplBin), backend, _rate, _wallet, _token, _tokenWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllowanceCrowdsaleImpl{AllowanceCrowdsaleImplCaller: AllowanceCrowdsaleImplCaller{contract: contract}, AllowanceCrowdsaleImplTransactor: AllowanceCrowdsaleImplTransactor{contract: contract}, AllowanceCrowdsaleImplFilterer: AllowanceCrowdsaleImplFilterer{contract: contract}}, nil
}

// AllowanceCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type AllowanceCrowdsaleImpl struct {
	AllowanceCrowdsaleImplCaller     // Read-only binding to the contract
	AllowanceCrowdsaleImplTransactor // Write-only binding to the contract
	AllowanceCrowdsaleImplFilterer   // Log filterer for contract events
}

// AllowanceCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllowanceCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllowanceCrowdsaleImplSession struct {
	Contract     *AllowanceCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AllowanceCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllowanceCrowdsaleImplCallerSession struct {
	Contract *AllowanceCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AllowanceCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllowanceCrowdsaleImplTransactorSession struct {
	Contract     *AllowanceCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AllowanceCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllowanceCrowdsaleImplRaw struct {
	Contract *AllowanceCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// AllowanceCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleImplCallerRaw struct {
	Contract *AllowanceCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// AllowanceCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleImplTransactorRaw struct {
	Contract *AllowanceCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllowanceCrowdsaleImpl creates a new instance of AllowanceCrowdsaleImpl, bound to a specific deployed contract.
func NewAllowanceCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*AllowanceCrowdsaleImpl, error) {
	contract, err := bindAllowanceCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleImpl{AllowanceCrowdsaleImplCaller: AllowanceCrowdsaleImplCaller{contract: contract}, AllowanceCrowdsaleImplTransactor: AllowanceCrowdsaleImplTransactor{contract: contract}, AllowanceCrowdsaleImplFilterer: AllowanceCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewAllowanceCrowdsaleImplCaller creates a new read-only instance of AllowanceCrowdsaleImpl, bound to a specific deployed contract.
func NewAllowanceCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*AllowanceCrowdsaleImplCaller, error) {
	contract, err := bindAllowanceCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleImplCaller{contract: contract}, nil
}

// NewAllowanceCrowdsaleImplTransactor creates a new write-only instance of AllowanceCrowdsaleImpl, bound to a specific deployed contract.
func NewAllowanceCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*AllowanceCrowdsaleImplTransactor, error) {
	contract, err := bindAllowanceCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleImplTransactor{contract: contract}, nil
}

// NewAllowanceCrowdsaleImplFilterer creates a new log filterer instance of AllowanceCrowdsaleImpl, bound to a specific deployed contract.
func NewAllowanceCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*AllowanceCrowdsaleImplFilterer, error) {
	contract, err := bindAllowanceCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleImplFilterer{contract: contract}, nil
}

// bindAllowanceCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindAllowanceCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AllowanceCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllowanceCrowdsaleImpl.Contract.AllowanceCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.Contract.AllowanceCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.Contract.AllowanceCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllowanceCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _AllowanceCrowdsaleImpl.Contract.Rate(&_AllowanceCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _AllowanceCrowdsaleImpl.Contract.Rate(&_AllowanceCrowdsaleImpl.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCaller) RemainingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceCrowdsaleImpl.contract.Call(opts, out, "remainingTokens")
	return *ret0, err
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) RemainingTokens() (*big.Int, error) {
	return _AllowanceCrowdsaleImpl.Contract.RemainingTokens(&_AllowanceCrowdsaleImpl.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerSession) RemainingTokens() (*big.Int, error) {
	return _AllowanceCrowdsaleImpl.Contract.RemainingTokens(&_AllowanceCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) Token() (common.Address, error) {
	return _AllowanceCrowdsaleImpl.Contract.Token(&_AllowanceCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _AllowanceCrowdsaleImpl.Contract.Token(&_AllowanceCrowdsaleImpl.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCaller) TokenWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceCrowdsaleImpl.contract.Call(opts, out, "tokenWallet")
	return *ret0, err
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) TokenWallet() (common.Address, error) {
	return _AllowanceCrowdsaleImpl.Contract.TokenWallet(&_AllowanceCrowdsaleImpl.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerSession) TokenWallet() (common.Address, error) {
	return _AllowanceCrowdsaleImpl.Contract.TokenWallet(&_AllowanceCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _AllowanceCrowdsaleImpl.Contract.Wallet(&_AllowanceCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _AllowanceCrowdsaleImpl.Contract.Wallet(&_AllowanceCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _AllowanceCrowdsaleImpl.Contract.WeiRaised(&_AllowanceCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _AllowanceCrowdsaleImpl.Contract.WeiRaised(&_AllowanceCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.Contract.BuyTokens(&_AllowanceCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _AllowanceCrowdsaleImpl.Contract.BuyTokens(&_AllowanceCrowdsaleImpl.TransactOpts, _beneficiary)
}

// AllowanceCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the AllowanceCrowdsaleImpl contract.
type AllowanceCrowdsaleImplTokenPurchaseIterator struct {
	Event *AllowanceCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *AllowanceCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllowanceCrowdsaleImplTokenPurchase)
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
		it.Event = new(AllowanceCrowdsaleImplTokenPurchase)
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
func (it *AllowanceCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllowanceCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllowanceCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the AllowanceCrowdsaleImpl contract.
type AllowanceCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*AllowanceCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _AllowanceCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleImplTokenPurchaseIterator{contract: _AllowanceCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_AllowanceCrowdsaleImpl *AllowanceCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *AllowanceCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _AllowanceCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllowanceCrowdsaleImplTokenPurchase)
				if err := _AllowanceCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
