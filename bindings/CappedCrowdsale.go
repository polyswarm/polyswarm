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

// CappedCrowdsaleABI is the input ABI used to generate the binding from.
const CappedCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// CappedCrowdsale is an auto generated Go binding around an Ethereum contract.
type CappedCrowdsale struct {
	CappedCrowdsaleCaller     // Read-only binding to the contract
	CappedCrowdsaleTransactor // Write-only binding to the contract
	CappedCrowdsaleFilterer   // Log filterer for contract events
}

// CappedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedCrowdsaleSession struct {
	Contract     *CappedCrowdsale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CappedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedCrowdsaleCallerSession struct {
	Contract *CappedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CappedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedCrowdsaleTransactorSession struct {
	Contract     *CappedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CappedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedCrowdsaleRaw struct {
	Contract *CappedCrowdsale // Generic contract binding to access the raw methods on
}

// CappedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedCrowdsaleCallerRaw struct {
	Contract *CappedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CappedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedCrowdsaleTransactorRaw struct {
	Contract *CappedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedCrowdsale creates a new instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsale(address common.Address, backend bind.ContractBackend) (*CappedCrowdsale, error) {
	contract, err := bindCappedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsale{CappedCrowdsaleCaller: CappedCrowdsaleCaller{contract: contract}, CappedCrowdsaleTransactor: CappedCrowdsaleTransactor{contract: contract}, CappedCrowdsaleFilterer: CappedCrowdsaleFilterer{contract: contract}}, nil
}

// NewCappedCrowdsaleCaller creates a new read-only instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CappedCrowdsaleCaller, error) {
	contract, err := bindCappedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleCaller{contract: contract}, nil
}

// NewCappedCrowdsaleTransactor creates a new write-only instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedCrowdsaleTransactor, error) {
	contract, err := bindCappedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleTransactor{contract: contract}, nil
}

// NewCappedCrowdsaleFilterer creates a new log filterer instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedCrowdsaleFilterer, error) {
	contract, err := bindCappedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleFilterer{contract: contract}, nil
}

// bindCappedCrowdsale binds a generic wrapper to an already deployed contract.
func bindCappedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsale.Contract.CappedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.CappedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.CappedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsale *CappedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsale *CappedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsale *CappedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) Cap() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Cap(&_CappedCrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Cap(&_CappedCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "capReached")
	return *ret0, err
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleSession) CapReached() (bool, error) {
	return _CappedCrowdsale.Contract.CapReached(&_CappedCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) CapReached() (bool, error) {
	return _CappedCrowdsale.Contract.CapReached(&_CappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) Rate() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Rate(&_CappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Rate(&_CappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Token() (common.Address, error) {
	return _CappedCrowdsale.Contract.Token(&_CappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _CappedCrowdsale.Contract.Token(&_CappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Wallet() (common.Address, error) {
	return _CappedCrowdsale.Contract.Wallet(&_CappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _CappedCrowdsale.Contract.Wallet(&_CappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsale.Contract.WeiRaised(&_CappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsale.Contract.WeiRaised(&_CappedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.BuyTokens(&_CappedCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.BuyTokens(&_CappedCrowdsale.TransactOpts, _beneficiary)
}

// CappedCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the CappedCrowdsale contract.
type CappedCrowdsaleTokenPurchaseIterator struct {
	Event *CappedCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

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
func (it *CappedCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCrowdsaleTokenPurchase)
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
		it.Event = new(CappedCrowdsaleTokenPurchase)
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
func (it *CappedCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCrowdsaleTokenPurchase represents a TokenPurchase event raised by the CappedCrowdsale contract.
type CappedCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*CappedCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleTokenPurchaseIterator{contract: _CappedCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *CappedCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCrowdsaleTokenPurchase)
				if err := _CappedCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
