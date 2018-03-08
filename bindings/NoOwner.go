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

// NoOwnerABI is the input ABI used to generate the binding from.
const NoOwnerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"reclaimContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// NoOwnerBin is the compiled bytecode used for deploying new contracts.
const NoOwnerBin = `606060405260008054600160a060020a03191633600160a060020a0316179055341561002a57600080fd5b61042b806100396000396000f30060606040526004361061005e5763ffffffff60e060020a60003504166317ffc320811461006b5780632aed7f3f1461008a5780638da5cb5b146100a95780639f727c27146100d8578063c0ee0b8a146100eb578063f2fde38b1461011a575b341561006957600080fd5b005b341561007657600080fd5b610069600160a060020a0360043516610139565b341561009557600080fd5b610069600160a060020a03600435166101ed565b34156100b457600080fd5b6100bc610278565b604051600160a060020a03909116815260200160405180910390f35b34156100e357600080fd5b610069610287565b34156100f657600080fd5b61006960048035600160a060020a03169060248035916044359182019101356102da565b341561012557600080fd5b610069600160a060020a03600435166102df565b6000805433600160a060020a0390811691161461015557600080fd5b81600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156101ac57600080fd5b6102c65a03f115156101bd57600080fd5b50505060405180516000549092506101e99150600160a060020a0384811691168363ffffffff61037a16565b5050565b6000805433600160a060020a0390811691161461020957600080fd5b506000548190600160a060020a038083169163f2fde38b911660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561026057600080fd5b6102c65a03f1151561027157600080fd5b5050505050565b600054600160a060020a031681565b60005433600160a060020a039081169116146102a257600080fd5b600054600160a060020a039081169030163180156108fc0290604051600060405180830381858888f1935050505015156102d857fe5b565b600080fd5b60005433600160a060020a039081169116146102fa57600080fd5b600160a060020a038116151561030f57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156103d757600080fd5b6102c65a03f115156103e857600080fd5b5050506040518051905015156103fa57fe5b5050505600a165627a7a72305820b9816ca5c5827a7ecd0147acc72398440cfd78e22f046565524d41fbd70b224b0029`

// DeployNoOwner deploys a new Ethereum contract, binding an instance of NoOwner to it.
func DeployNoOwner(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOwner, error) {
	parsed, err := abi.JSON(strings.NewReader(NoOwnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NoOwnerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOwner{NoOwnerCaller: NoOwnerCaller{contract: contract}, NoOwnerTransactor: NoOwnerTransactor{contract: contract}, NoOwnerFilterer: NoOwnerFilterer{contract: contract}}, nil
}

// NoOwner is an auto generated Go binding around an Ethereum contract.
type NoOwner struct {
	NoOwnerCaller     // Read-only binding to the contract
	NoOwnerTransactor // Write-only binding to the contract
	NoOwnerFilterer   // Log filterer for contract events
}

// NoOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoOwnerSession struct {
	Contract     *NoOwner          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoOwnerCallerSession struct {
	Contract *NoOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NoOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoOwnerTransactorSession struct {
	Contract     *NoOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NoOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoOwnerRaw struct {
	Contract *NoOwner // Generic contract binding to access the raw methods on
}

// NoOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoOwnerCallerRaw struct {
	Contract *NoOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// NoOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoOwnerTransactorRaw struct {
	Contract *NoOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoOwner creates a new instance of NoOwner, bound to a specific deployed contract.
func NewNoOwner(address common.Address, backend bind.ContractBackend) (*NoOwner, error) {
	contract, err := bindNoOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOwner{NoOwnerCaller: NoOwnerCaller{contract: contract}, NoOwnerTransactor: NoOwnerTransactor{contract: contract}, NoOwnerFilterer: NoOwnerFilterer{contract: contract}}, nil
}

// NewNoOwnerCaller creates a new read-only instance of NoOwner, bound to a specific deployed contract.
func NewNoOwnerCaller(address common.Address, caller bind.ContractCaller) (*NoOwnerCaller, error) {
	contract, err := bindNoOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOwnerCaller{contract: contract}, nil
}

// NewNoOwnerTransactor creates a new write-only instance of NoOwner, bound to a specific deployed contract.
func NewNoOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOwnerTransactor, error) {
	contract, err := bindNoOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOwnerTransactor{contract: contract}, nil
}

// NewNoOwnerFilterer creates a new log filterer instance of NoOwner, bound to a specific deployed contract.
func NewNoOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOwnerFilterer, error) {
	contract, err := bindNoOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOwnerFilterer{contract: contract}, nil
}

// bindNoOwner binds a generic wrapper to an already deployed contract.
func bindNoOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOwner *NoOwnerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NoOwner.Contract.NoOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOwner *NoOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOwner.Contract.NoOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOwner *NoOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOwner.Contract.NoOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOwner *NoOwnerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NoOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOwner *NoOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOwner *NoOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOwner.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NoOwner *NoOwnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NoOwner.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NoOwner *NoOwnerSession) Owner() (common.Address, error) {
	return _NoOwner.Contract.Owner(&_NoOwner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NoOwner *NoOwnerCallerSession) Owner() (common.Address, error) {
	return _NoOwner.Contract.Owner(&_NoOwner.CallOpts)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_NoOwner *NoOwnerTransactor) ReclaimContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "reclaimContract", contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_NoOwner *NoOwnerSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimContract(&_NoOwner.TransactOpts, contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_NoOwner *NoOwnerTransactorSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimContract(&_NoOwner.TransactOpts, contractAddr)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_NoOwner *NoOwnerTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_NoOwner *NoOwnerSession) ReclaimEther() (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimEther(&_NoOwner.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_NoOwner *NoOwnerTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimEther(&_NoOwner.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_NoOwner *NoOwnerTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_NoOwner *NoOwnerSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimToken(&_NoOwner.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_NoOwner *NoOwnerTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimToken(&_NoOwner.TransactOpts, token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_NoOwner *NoOwnerTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_NoOwner *NoOwnerSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _NoOwner.Contract.TokenFallback(&_NoOwner.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_NoOwner *NoOwnerTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _NoOwner.Contract.TokenFallback(&_NoOwner.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NoOwner *NoOwnerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NoOwner *NoOwnerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.TransferOwnership(&_NoOwner.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NoOwner *NoOwnerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.TransferOwnership(&_NoOwner.TransactOpts, newOwner)
}

// NoOwnerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NoOwner contract.
type NoOwnerOwnershipTransferredIterator struct {
	Event *NoOwnerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NoOwnerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoOwnerOwnershipTransferred)
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
		it.Event = new(NoOwnerOwnershipTransferred)
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
func (it *NoOwnerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoOwnerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoOwnerOwnershipTransferred represents a OwnershipTransferred event raised by the NoOwner contract.
type NoOwnerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NoOwner *NoOwnerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NoOwnerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NoOwner.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NoOwnerOwnershipTransferredIterator{contract: _NoOwner.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NoOwner *NoOwnerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NoOwnerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NoOwner.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoOwnerOwnershipTransferred)
				if err := _NoOwner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
