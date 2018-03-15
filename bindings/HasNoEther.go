// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// HasNoEtherABI is the input ABI used to generate the binding from.
const HasNoEtherABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoEtherBin is the compiled bytecode used for deploying new contracts.
const HasNoEtherBin = `606060405260008054600160a060020a03191633600160a060020a0316179055341561002a57600080fd5b6101ed806100396000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b81146100635780639f727c2714610092578063f2fde38b146100a5575b341561006157600080fd5b005b341561006e57600080fd5b6100766100c4565b604051600160a060020a03909116815260200160405180910390f35b341561009d57600080fd5b6100616100d3565b34156100b057600080fd5b610061600160a060020a0360043516610126565b600054600160a060020a031681565b60005433600160a060020a039081169116146100ee57600080fd5b600054600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561012457fe5b565b60005433600160a060020a0390811691161461014157600080fd5b600160a060020a038116151561015657600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d460fbf84915e5f71b16a3bf65e58a8a43aeb98781f41b2cb50c0f10a33321640029`

// DeployHasNoEther deploys a new Ethereum contract, binding an instance of HasNoEther to it.
func DeployHasNoEther(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoEther, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoEtherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoEtherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoEther{HasNoEtherCaller: HasNoEtherCaller{contract: contract}, HasNoEtherTransactor: HasNoEtherTransactor{contract: contract}, HasNoEtherFilterer: HasNoEtherFilterer{contract: contract}}, nil
}

// HasNoEther is an auto generated Go binding around an Ethereum contract.
type HasNoEther struct {
	HasNoEtherCaller     // Read-only binding to the contract
	HasNoEtherTransactor // Write-only binding to the contract
	HasNoEtherFilterer   // Log filterer for contract events
}

// HasNoEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasNoEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoEtherSession struct {
	Contract     *HasNoEther       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoEtherCallerSession struct {
	Contract *HasNoEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HasNoEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoEtherTransactorSession struct {
	Contract     *HasNoEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HasNoEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoEtherRaw struct {
	Contract *HasNoEther // Generic contract binding to access the raw methods on
}

// HasNoEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoEtherCallerRaw struct {
	Contract *HasNoEtherCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoEtherTransactorRaw struct {
	Contract *HasNoEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoEther creates a new instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEther(address common.Address, backend bind.ContractBackend) (*HasNoEther, error) {
	contract, err := bindHasNoEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoEther{HasNoEtherCaller: HasNoEtherCaller{contract: contract}, HasNoEtherTransactor: HasNoEtherTransactor{contract: contract}, HasNoEtherFilterer: HasNoEtherFilterer{contract: contract}}, nil
}

// NewHasNoEtherCaller creates a new read-only instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEtherCaller(address common.Address, caller bind.ContractCaller) (*HasNoEtherCaller, error) {
	contract, err := bindHasNoEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherCaller{contract: contract}, nil
}

// NewHasNoEtherTransactor creates a new write-only instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoEtherTransactor, error) {
	contract, err := bindHasNoEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTransactor{contract: contract}, nil
}

// NewHasNoEtherFilterer creates a new log filterer instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*HasNoEtherFilterer, error) {
	contract, err := bindHasNoEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherFilterer{contract: contract}, nil
}

// bindHasNoEther binds a generic wrapper to an already deployed contract.
func bindHasNoEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoEther *HasNoEtherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoEther.Contract.HasNoEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoEther *HasNoEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEther.Contract.HasNoEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoEther *HasNoEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoEther.Contract.HasNoEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoEther *HasNoEtherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoEther *HasNoEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoEther *HasNoEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoEther.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEther *HasNoEtherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoEther.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEther *HasNoEtherSession) Owner() (common.Address, error) {
	return _HasNoEther.Contract.Owner(&_HasNoEther.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEther *HasNoEtherCallerSession) Owner() (common.Address, error) {
	return _HasNoEther.Contract.Owner(&_HasNoEther.CallOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEther *HasNoEtherTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEther.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEther *HasNoEtherSession) ReclaimEther() (*types.Transaction, error) {
	return _HasNoEther.Contract.ReclaimEther(&_HasNoEther.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEther *HasNoEtherTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _HasNoEther.Contract.ReclaimEther(&_HasNoEther.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEther *HasNoEtherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEther.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEther *HasNoEtherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEther.Contract.TransferOwnership(&_HasNoEther.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEther *HasNoEtherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEther.Contract.TransferOwnership(&_HasNoEther.TransactOpts, newOwner)
}

// HasNoEtherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HasNoEther contract.
type HasNoEtherOwnershipTransferredIterator struct {
	Event *HasNoEtherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HasNoEtherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HasNoEtherOwnershipTransferred)
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
		it.Event = new(HasNoEtherOwnershipTransferred)
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
func (it *HasNoEtherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HasNoEtherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HasNoEtherOwnershipTransferred represents a OwnershipTransferred event raised by the HasNoEther contract.
type HasNoEtherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoEther *HasNoEtherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HasNoEtherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoEther.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherOwnershipTransferredIterator{contract: _HasNoEther.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoEther *HasNoEtherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HasNoEtherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoEther.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HasNoEtherOwnershipTransferred)
				if err := _HasNoEther.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
