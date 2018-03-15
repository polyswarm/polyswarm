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

// HasNoEtherTestABI is the input ABI used to generate the binding from.
const HasNoEtherTestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoEtherTestBin is the compiled bytecode used for deploying new contracts.
const HasNoEtherTestBin = `606060405260008054600160a060020a03191633600160a060020a0316179055341561002a57600080fd5b6101ed806100396000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b81146100635780639f727c2714610092578063f2fde38b146100a5575b341561006157600080fd5b005b341561006e57600080fd5b6100766100c4565b604051600160a060020a03909116815260200160405180910390f35b341561009d57600080fd5b6100616100d3565b34156100b057600080fd5b610061600160a060020a0360043516610126565b600054600160a060020a031681565b60005433600160a060020a039081169116146100ee57600080fd5b600054600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561012457fe5b565b60005433600160a060020a0390811691161461014157600080fd5b600160a060020a038116151561015657600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058207db808b56999de4b9f67bd2b85bd4d5265f774d19d53816baf85c6b87bea2ad30029`

// DeployHasNoEtherTest deploys a new Ethereum contract, binding an instance of HasNoEtherTest to it.
func DeployHasNoEtherTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoEtherTest, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoEtherTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoEtherTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoEtherTest{HasNoEtherTestCaller: HasNoEtherTestCaller{contract: contract}, HasNoEtherTestTransactor: HasNoEtherTestTransactor{contract: contract}, HasNoEtherTestFilterer: HasNoEtherTestFilterer{contract: contract}}, nil
}

// HasNoEtherTest is an auto generated Go binding around an Ethereum contract.
type HasNoEtherTest struct {
	HasNoEtherTestCaller     // Read-only binding to the contract
	HasNoEtherTestTransactor // Write-only binding to the contract
	HasNoEtherTestFilterer   // Log filterer for contract events
}

// HasNoEtherTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoEtherTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoEtherTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasNoEtherTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoEtherTestSession struct {
	Contract     *HasNoEtherTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoEtherTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoEtherTestCallerSession struct {
	Contract *HasNoEtherTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HasNoEtherTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoEtherTestTransactorSession struct {
	Contract     *HasNoEtherTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HasNoEtherTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoEtherTestRaw struct {
	Contract *HasNoEtherTest // Generic contract binding to access the raw methods on
}

// HasNoEtherTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoEtherTestCallerRaw struct {
	Contract *HasNoEtherTestCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoEtherTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoEtherTestTransactorRaw struct {
	Contract *HasNoEtherTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoEtherTest creates a new instance of HasNoEtherTest, bound to a specific deployed contract.
func NewHasNoEtherTest(address common.Address, backend bind.ContractBackend) (*HasNoEtherTest, error) {
	contract, err := bindHasNoEtherTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTest{HasNoEtherTestCaller: HasNoEtherTestCaller{contract: contract}, HasNoEtherTestTransactor: HasNoEtherTestTransactor{contract: contract}, HasNoEtherTestFilterer: HasNoEtherTestFilterer{contract: contract}}, nil
}

// NewHasNoEtherTestCaller creates a new read-only instance of HasNoEtherTest, bound to a specific deployed contract.
func NewHasNoEtherTestCaller(address common.Address, caller bind.ContractCaller) (*HasNoEtherTestCaller, error) {
	contract, err := bindHasNoEtherTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTestCaller{contract: contract}, nil
}

// NewHasNoEtherTestTransactor creates a new write-only instance of HasNoEtherTest, bound to a specific deployed contract.
func NewHasNoEtherTestTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoEtherTestTransactor, error) {
	contract, err := bindHasNoEtherTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTestTransactor{contract: contract}, nil
}

// NewHasNoEtherTestFilterer creates a new log filterer instance of HasNoEtherTest, bound to a specific deployed contract.
func NewHasNoEtherTestFilterer(address common.Address, filterer bind.ContractFilterer) (*HasNoEtherTestFilterer, error) {
	contract, err := bindHasNoEtherTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTestFilterer{contract: contract}, nil
}

// bindHasNoEtherTest binds a generic wrapper to an already deployed contract.
func bindHasNoEtherTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoEtherTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoEtherTest *HasNoEtherTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoEtherTest.Contract.HasNoEtherTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoEtherTest *HasNoEtherTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.HasNoEtherTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoEtherTest *HasNoEtherTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.HasNoEtherTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoEtherTest *HasNoEtherTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoEtherTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoEtherTest *HasNoEtherTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoEtherTest *HasNoEtherTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEtherTest *HasNoEtherTestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoEtherTest.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEtherTest *HasNoEtherTestSession) Owner() (common.Address, error) {
	return _HasNoEtherTest.Contract.Owner(&_HasNoEtherTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEtherTest *HasNoEtherTestCallerSession) Owner() (common.Address, error) {
	return _HasNoEtherTest.Contract.Owner(&_HasNoEtherTest.CallOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEtherTest *HasNoEtherTestTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEtherTest.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEtherTest *HasNoEtherTestSession) ReclaimEther() (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.ReclaimEther(&_HasNoEtherTest.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEtherTest *HasNoEtherTestTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.ReclaimEther(&_HasNoEtherTest.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEtherTest *HasNoEtherTestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEtherTest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEtherTest *HasNoEtherTestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.TransferOwnership(&_HasNoEtherTest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEtherTest *HasNoEtherTestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEtherTest.Contract.TransferOwnership(&_HasNoEtherTest.TransactOpts, newOwner)
}

// HasNoEtherTestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HasNoEtherTest contract.
type HasNoEtherTestOwnershipTransferredIterator struct {
	Event *HasNoEtherTestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HasNoEtherTestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HasNoEtherTestOwnershipTransferred)
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
		it.Event = new(HasNoEtherTestOwnershipTransferred)
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
func (it *HasNoEtherTestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HasNoEtherTestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HasNoEtherTestOwnershipTransferred represents a OwnershipTransferred event raised by the HasNoEtherTest contract.
type HasNoEtherTestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoEtherTest *HasNoEtherTestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HasNoEtherTestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoEtherTest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTestOwnershipTransferredIterator{contract: _HasNoEtherTest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoEtherTest *HasNoEtherTestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HasNoEtherTestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoEtherTest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HasNoEtherTestOwnershipTransferred)
				if err := _HasNoEtherTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
