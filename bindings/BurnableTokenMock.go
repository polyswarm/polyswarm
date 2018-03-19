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

// BurnableTokenMockABI is the input ABI used to generate the binding from.
const BurnableTokenMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BurnableTokenMockBin is the compiled bytecode used for deploying new contracts.
const BurnableTokenMockBin = `6060604052341561000f57600080fd5b6040516040806103d18339810160405280805191906020018051600160a060020a0390931660009081526020819052604090208390555050600155610378806100596000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461006657806342966c681461008b57806370a08231146100a3578063a9059cbb146100c2575b600080fd5b341561007157600080fd5b6100796100f8565b60405190815260200160405180910390f35b341561009657600080fd5b6100a16004356100fe565b005b34156100ae57600080fd5b610079600160a060020a03600435166101f7565b34156100cd57600080fd5b6100e4600160a060020a0360043516602435610212565b604051901515815260200160405180910390f35b60015490565b600160a060020a03331660009081526020819052604081205482111561012357600080fd5b5033600160a060020a0381166000908152602081905260409020546101489083610324565b600160a060020a038216600090815260208190526040902055600154610174908363ffffffff61032416565b600155600160a060020a0381167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58360405190815260200160405180910390a26000600160a060020a0382167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35050565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a038316151561022957600080fd5b600160a060020a03331660009081526020819052604090205482111561024e57600080fd5b600160a060020a033316600090815260208190526040902054610277908363ffffffff61032416565b600160a060020a0333811660009081526020819052604080822093909355908516815220546102ac908363ffffffff61033616565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b60008282111561033057fe5b50900390565b60008282018381101561034557fe5b93925050505600a165627a7a723058204b6aa9bc866604f24feea7fe8f7890b5bdfefbca262881ca30ae9ea43782c2960029`

// DeployBurnableTokenMock deploys a new Ethereum contract, binding an instance of BurnableTokenMock to it.
func DeployBurnableTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *BurnableTokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnableTokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableTokenMock{BurnableTokenMockCaller: BurnableTokenMockCaller{contract: contract}, BurnableTokenMockTransactor: BurnableTokenMockTransactor{contract: contract}, BurnableTokenMockFilterer: BurnableTokenMockFilterer{contract: contract}}, nil
}

// BurnableTokenMock is an auto generated Go binding around an Ethereum contract.
type BurnableTokenMock struct {
	BurnableTokenMockCaller     // Read-only binding to the contract
	BurnableTokenMockTransactor // Write-only binding to the contract
	BurnableTokenMockFilterer   // Log filterer for contract events
}

// BurnableTokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnableTokenMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenMockSession struct {
	Contract     *BurnableTokenMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BurnableTokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenMockCallerSession struct {
	Contract *BurnableTokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BurnableTokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenMockTransactorSession struct {
	Contract     *BurnableTokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BurnableTokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenMockRaw struct {
	Contract *BurnableTokenMock // Generic contract binding to access the raw methods on
}

// BurnableTokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenMockCallerRaw struct {
	Contract *BurnableTokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenMockTransactorRaw struct {
	Contract *BurnableTokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableTokenMock creates a new instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMock(address common.Address, backend bind.ContractBackend) (*BurnableTokenMock, error) {
	contract, err := bindBurnableTokenMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMock{BurnableTokenMockCaller: BurnableTokenMockCaller{contract: contract}, BurnableTokenMockTransactor: BurnableTokenMockTransactor{contract: contract}, BurnableTokenMockFilterer: BurnableTokenMockFilterer{contract: contract}}, nil
}

// NewBurnableTokenMockCaller creates a new read-only instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMockCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenMockCaller, error) {
	contract, err := bindBurnableTokenMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockCaller{contract: contract}, nil
}

// NewBurnableTokenMockTransactor creates a new write-only instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenMockTransactor, error) {
	contract, err := bindBurnableTokenMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockTransactor{contract: contract}, nil
}

// NewBurnableTokenMockFilterer creates a new log filterer instance of BurnableTokenMock, bound to a specific deployed contract.
func NewBurnableTokenMockFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnableTokenMockFilterer, error) {
	contract, err := bindBurnableTokenMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockFilterer{contract: contract}, nil
}

// bindBurnableTokenMock binds a generic wrapper to an already deployed contract.
func bindBurnableTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableTokenMock *BurnableTokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableTokenMock.Contract.BurnableTokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableTokenMock *BurnableTokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.BurnableTokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableTokenMock *BurnableTokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.BurnableTokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableTokenMock *BurnableTokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableTokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableTokenMock *BurnableTokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableTokenMock *BurnableTokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableTokenMock *BurnableTokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableTokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableTokenMock *BurnableTokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableTokenMock.Contract.BalanceOf(&_BurnableTokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableTokenMock *BurnableTokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableTokenMock.Contract.BalanceOf(&_BurnableTokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableTokenMock *BurnableTokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableTokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableTokenMock *BurnableTokenMockSession) TotalSupply() (*big.Int, error) {
	return _BurnableTokenMock.Contract.TotalSupply(&_BurnableTokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableTokenMock *BurnableTokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableTokenMock.Contract.TotalSupply(&_BurnableTokenMock.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableTokenMock *BurnableTokenMockTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableTokenMock *BurnableTokenMockSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Burn(&_BurnableTokenMock.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableTokenMock *BurnableTokenMockTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Burn(&_BurnableTokenMock.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableTokenMock *BurnableTokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableTokenMock *BurnableTokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Transfer(&_BurnableTokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableTokenMock *BurnableTokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableTokenMock.Contract.Transfer(&_BurnableTokenMock.TransactOpts, _to, _value)
}

// BurnableTokenMockBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BurnableTokenMock contract.
type BurnableTokenMockBurnIterator struct {
	Event *BurnableTokenMockBurn // Event containing the contract specifics and raw log

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
func (it *BurnableTokenMockBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenMockBurn)
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
		it.Event = new(BurnableTokenMockBurn)
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
func (it *BurnableTokenMockBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenMockBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenMockBurn represents a Burn event raised by the BurnableTokenMock contract.
type BurnableTokenMockBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(burner indexed address, value uint256)
func (_BurnableTokenMock *BurnableTokenMockFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*BurnableTokenMockBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableTokenMock.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockBurnIterator{contract: _BurnableTokenMock.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(burner indexed address, value uint256)
func (_BurnableTokenMock *BurnableTokenMockFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BurnableTokenMockBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableTokenMock.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenMockBurn)
				if err := _BurnableTokenMock.contract.UnpackLog(event, "Burn", log); err != nil {
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

// BurnableTokenMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnableTokenMock contract.
type BurnableTokenMockTransferIterator struct {
	Event *BurnableTokenMockTransfer // Event containing the contract specifics and raw log

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
func (it *BurnableTokenMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenMockTransfer)
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
		it.Event = new(BurnableTokenMockTransfer)
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
func (it *BurnableTokenMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenMockTransfer represents a Transfer event raised by the BurnableTokenMock contract.
type BurnableTokenMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableTokenMock *BurnableTokenMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnableTokenMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableTokenMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMockTransferIterator{contract: _BurnableTokenMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableTokenMock *BurnableTokenMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnableTokenMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableTokenMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenMockTransfer)
				if err := _BurnableTokenMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
