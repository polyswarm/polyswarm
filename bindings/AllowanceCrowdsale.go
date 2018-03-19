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

// AllowanceCrowdsaleABI is the input ABI used to generate the binding from.
const AllowanceCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remainingTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// AllowanceCrowdsale is an auto generated Go binding around an Ethereum contract.
type AllowanceCrowdsale struct {
	AllowanceCrowdsaleCaller     // Read-only binding to the contract
	AllowanceCrowdsaleTransactor // Write-only binding to the contract
	AllowanceCrowdsaleFilterer   // Log filterer for contract events
}

// AllowanceCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllowanceCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllowanceCrowdsaleSession struct {
	Contract     *AllowanceCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AllowanceCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllowanceCrowdsaleCallerSession struct {
	Contract *AllowanceCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AllowanceCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllowanceCrowdsaleTransactorSession struct {
	Contract     *AllowanceCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AllowanceCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllowanceCrowdsaleRaw struct {
	Contract *AllowanceCrowdsale // Generic contract binding to access the raw methods on
}

// AllowanceCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleCallerRaw struct {
	Contract *AllowanceCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// AllowanceCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllowanceCrowdsaleTransactorRaw struct {
	Contract *AllowanceCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllowanceCrowdsale creates a new instance of AllowanceCrowdsale, bound to a specific deployed contract.
func NewAllowanceCrowdsale(address common.Address, backend bind.ContractBackend) (*AllowanceCrowdsale, error) {
	contract, err := bindAllowanceCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsale{AllowanceCrowdsaleCaller: AllowanceCrowdsaleCaller{contract: contract}, AllowanceCrowdsaleTransactor: AllowanceCrowdsaleTransactor{contract: contract}, AllowanceCrowdsaleFilterer: AllowanceCrowdsaleFilterer{contract: contract}}, nil
}

// NewAllowanceCrowdsaleCaller creates a new read-only instance of AllowanceCrowdsale, bound to a specific deployed contract.
func NewAllowanceCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*AllowanceCrowdsaleCaller, error) {
	contract, err := bindAllowanceCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleCaller{contract: contract}, nil
}

// NewAllowanceCrowdsaleTransactor creates a new write-only instance of AllowanceCrowdsale, bound to a specific deployed contract.
func NewAllowanceCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*AllowanceCrowdsaleTransactor, error) {
	contract, err := bindAllowanceCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleTransactor{contract: contract}, nil
}

// NewAllowanceCrowdsaleFilterer creates a new log filterer instance of AllowanceCrowdsale, bound to a specific deployed contract.
func NewAllowanceCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*AllowanceCrowdsaleFilterer, error) {
	contract, err := bindAllowanceCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleFilterer{contract: contract}, nil
}

// bindAllowanceCrowdsale binds a generic wrapper to an already deployed contract.
func bindAllowanceCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AllowanceCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowanceCrowdsale *AllowanceCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllowanceCrowdsale.Contract.AllowanceCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowanceCrowdsale *AllowanceCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceCrowdsale.Contract.AllowanceCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowanceCrowdsale *AllowanceCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowanceCrowdsale.Contract.AllowanceCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllowanceCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowanceCrowdsale *AllowanceCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowanceCrowdsale *AllowanceCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowanceCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) Rate() (*big.Int, error) {
	return _AllowanceCrowdsale.Contract.Rate(&_AllowanceCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _AllowanceCrowdsale.Contract.Rate(&_AllowanceCrowdsale.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCaller) RemainingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceCrowdsale.contract.Call(opts, out, "remainingTokens")
	return *ret0, err
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) RemainingTokens() (*big.Int, error) {
	return _AllowanceCrowdsale.Contract.RemainingTokens(&_AllowanceCrowdsale.CallOpts)
}

// RemainingTokens is a free data retrieval call binding the contract method 0xbf583903.
//
// Solidity: function remainingTokens() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerSession) RemainingTokens() (*big.Int, error) {
	return _AllowanceCrowdsale.Contract.RemainingTokens(&_AllowanceCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) Token() (common.Address, error) {
	return _AllowanceCrowdsale.Contract.Token(&_AllowanceCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerSession) Token() (common.Address, error) {
	return _AllowanceCrowdsale.Contract.Token(&_AllowanceCrowdsale.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCaller) TokenWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceCrowdsale.contract.Call(opts, out, "tokenWallet")
	return *ret0, err
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) TokenWallet() (common.Address, error) {
	return _AllowanceCrowdsale.Contract.TokenWallet(&_AllowanceCrowdsale.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerSession) TokenWallet() (common.Address, error) {
	return _AllowanceCrowdsale.Contract.TokenWallet(&_AllowanceCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) Wallet() (common.Address, error) {
	return _AllowanceCrowdsale.Contract.Wallet(&_AllowanceCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _AllowanceCrowdsale.Contract.Wallet(&_AllowanceCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _AllowanceCrowdsale.Contract.WeiRaised(&_AllowanceCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _AllowanceCrowdsale.Contract.WeiRaised(&_AllowanceCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_AllowanceCrowdsale *AllowanceCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _AllowanceCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_AllowanceCrowdsale *AllowanceCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _AllowanceCrowdsale.Contract.BuyTokens(&_AllowanceCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_AllowanceCrowdsale *AllowanceCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _AllowanceCrowdsale.Contract.BuyTokens(&_AllowanceCrowdsale.TransactOpts, _beneficiary)
}

// AllowanceCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the AllowanceCrowdsale contract.
type AllowanceCrowdsaleTokenPurchaseIterator struct {
	Event *AllowanceCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

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
func (it *AllowanceCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllowanceCrowdsaleTokenPurchase)
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
		it.Event = new(AllowanceCrowdsaleTokenPurchase)
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
func (it *AllowanceCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllowanceCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllowanceCrowdsaleTokenPurchase represents a TokenPurchase event raised by the AllowanceCrowdsale contract.
type AllowanceCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*AllowanceCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _AllowanceCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &AllowanceCrowdsaleTokenPurchaseIterator{contract: _AllowanceCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_AllowanceCrowdsale *AllowanceCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *AllowanceCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _AllowanceCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllowanceCrowdsaleTokenPurchase)
				if err := _AllowanceCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
