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

// ERC23TokenMockABI is the input ABI used to generate the binding from.
const ERC23TokenMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferERC23\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC23TokenMockBin is the compiled bytecode used for deploying new contracts.
const ERC23TokenMockBin = `6060604052341561000f57600080fd5b6040516040806104338339810160405280805191906020018051600160a060020a0390931660009081526001602052604081208490559290925550506103d98061005a6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd8114610066578063685c60de1461008b57806370a0823114610104578063a9059cbb14610123575b600080fd5b341561007157600080fd5b610079610145565b60405190815260200160405180910390f35b341561009657600080fd5b6100f060048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061014b95505050505050565b604051901515815260200160405180910390f35b341561010f57600080fd5b610079600160a060020a036004351661026f565b341561012e57600080fd5b6100f0600160a060020a036004351660243561028a565b60005481565b600080600061015a868661028a565b50853b151991508115610263575084600160a060020a03811663c0ee0b8a3387876040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102015780820151838201526020016101e9565b50505050905090810190601f16801561022e5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b151561024e57600080fd5b6102c65a03f1151561025f57600080fd5b5050505b50600195945050505050565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a03831615156102a157600080fd5b600160a060020a0333166000908152600160205260409020548211156102c657600080fd5b600160a060020a0333166000908152600160205260409020546102ef908363ffffffff61038516565b600160a060020a033381166000908152600160205260408082209390935590851681522054610324908363ffffffff61039716565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b60008282111561039157fe5b50900390565b6000828201838110156103a657fe5b93925050505600a165627a7a7230582005ed958abc5ad984a39104e8dc6f349d02ce147b0800ce9c47304f9a707509790029`

// DeployERC23TokenMock deploys a new Ethereum contract, binding an instance of ERC23TokenMock to it.
func DeployERC23TokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *ERC23TokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC23TokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC23TokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC23TokenMock{ERC23TokenMockCaller: ERC23TokenMockCaller{contract: contract}, ERC23TokenMockTransactor: ERC23TokenMockTransactor{contract: contract}, ERC23TokenMockFilterer: ERC23TokenMockFilterer{contract: contract}}, nil
}

// ERC23TokenMock is an auto generated Go binding around an Ethereum contract.
type ERC23TokenMock struct {
	ERC23TokenMockCaller     // Read-only binding to the contract
	ERC23TokenMockTransactor // Write-only binding to the contract
	ERC23TokenMockFilterer   // Log filterer for contract events
}

// ERC23TokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC23TokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC23TokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC23TokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC23TokenMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC23TokenMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC23TokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC23TokenMockSession struct {
	Contract     *ERC23TokenMock   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC23TokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC23TokenMockCallerSession struct {
	Contract *ERC23TokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC23TokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC23TokenMockTransactorSession struct {
	Contract     *ERC23TokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC23TokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC23TokenMockRaw struct {
	Contract *ERC23TokenMock // Generic contract binding to access the raw methods on
}

// ERC23TokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC23TokenMockCallerRaw struct {
	Contract *ERC23TokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC23TokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC23TokenMockTransactorRaw struct {
	Contract *ERC23TokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC23TokenMock creates a new instance of ERC23TokenMock, bound to a specific deployed contract.
func NewERC23TokenMock(address common.Address, backend bind.ContractBackend) (*ERC23TokenMock, error) {
	contract, err := bindERC23TokenMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC23TokenMock{ERC23TokenMockCaller: ERC23TokenMockCaller{contract: contract}, ERC23TokenMockTransactor: ERC23TokenMockTransactor{contract: contract}, ERC23TokenMockFilterer: ERC23TokenMockFilterer{contract: contract}}, nil
}

// NewERC23TokenMockCaller creates a new read-only instance of ERC23TokenMock, bound to a specific deployed contract.
func NewERC23TokenMockCaller(address common.Address, caller bind.ContractCaller) (*ERC23TokenMockCaller, error) {
	contract, err := bindERC23TokenMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC23TokenMockCaller{contract: contract}, nil
}

// NewERC23TokenMockTransactor creates a new write-only instance of ERC23TokenMock, bound to a specific deployed contract.
func NewERC23TokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC23TokenMockTransactor, error) {
	contract, err := bindERC23TokenMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC23TokenMockTransactor{contract: contract}, nil
}

// NewERC23TokenMockFilterer creates a new log filterer instance of ERC23TokenMock, bound to a specific deployed contract.
func NewERC23TokenMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC23TokenMockFilterer, error) {
	contract, err := bindERC23TokenMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC23TokenMockFilterer{contract: contract}, nil
}

// bindERC23TokenMock binds a generic wrapper to an already deployed contract.
func bindERC23TokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC23TokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC23TokenMock *ERC23TokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC23TokenMock.Contract.ERC23TokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC23TokenMock *ERC23TokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.ERC23TokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC23TokenMock *ERC23TokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.ERC23TokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC23TokenMock *ERC23TokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC23TokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC23TokenMock *ERC23TokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC23TokenMock *ERC23TokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC23TokenMock *ERC23TokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC23TokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC23TokenMock *ERC23TokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC23TokenMock.Contract.BalanceOf(&_ERC23TokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC23TokenMock *ERC23TokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC23TokenMock.Contract.BalanceOf(&_ERC23TokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC23TokenMock *ERC23TokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC23TokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC23TokenMock *ERC23TokenMockSession) TotalSupply() (*big.Int, error) {
	return _ERC23TokenMock.Contract.TotalSupply(&_ERC23TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC23TokenMock *ERC23TokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC23TokenMock.Contract.TotalSupply(&_ERC23TokenMock.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC23TokenMock *ERC23TokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC23TokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC23TokenMock *ERC23TokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.Transfer(&_ERC23TokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC23TokenMock *ERC23TokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.Transfer(&_ERC23TokenMock.TransactOpts, _to, _value)
}

// TransferERC23 is a paid mutator transaction binding the contract method 0x685c60de.
//
// Solidity: function transferERC23(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC23TokenMock *ERC23TokenMockTransactor) TransferERC23(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC23TokenMock.contract.Transact(opts, "transferERC23", _to, _value, _data)
}

// TransferERC23 is a paid mutator transaction binding the contract method 0x685c60de.
//
// Solidity: function transferERC23(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC23TokenMock *ERC23TokenMockSession) TransferERC23(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.TransferERC23(&_ERC23TokenMock.TransactOpts, _to, _value, _data)
}

// TransferERC23 is a paid mutator transaction binding the contract method 0x685c60de.
//
// Solidity: function transferERC23(_to address, _value uint256, _data bytes) returns(success bool)
func (_ERC23TokenMock *ERC23TokenMockTransactorSession) TransferERC23(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC23TokenMock.Contract.TransferERC23(&_ERC23TokenMock.TransactOpts, _to, _value, _data)
}

// ERC23TokenMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC23TokenMock contract.
type ERC23TokenMockTransferIterator struct {
	Event *ERC23TokenMockTransfer // Event containing the contract specifics and raw log

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
func (it *ERC23TokenMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC23TokenMockTransfer)
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
		it.Event = new(ERC23TokenMockTransfer)
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
func (it *ERC23TokenMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC23TokenMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC23TokenMockTransfer represents a Transfer event raised by the ERC23TokenMock contract.
type ERC23TokenMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC23TokenMock *ERC23TokenMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC23TokenMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC23TokenMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC23TokenMockTransferIterator{contract: _ERC23TokenMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC23TokenMock *ERC23TokenMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC23TokenMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC23TokenMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC23TokenMockTransfer)
				if err := _ERC23TokenMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
