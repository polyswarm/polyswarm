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

// ContactableABI is the input ABI used to generate the binding from.
const ContactableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"contactInformation\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"info\",\"type\":\"string\"}],\"name\":\"setContactInformation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ContactableBin is the compiled bytecode used for deploying new contracts.
const ContactableBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556103d2806100306000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166336f7ab5e81146100665780638da5cb5b146100f0578063b967a52e1461011f578063f2fde38b14610172575b600080fd5b341561007157600080fd5b610079610191565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100b557808201518382015260200161009d565b50505050905090810190601f1680156100e25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100fb57600080fd5b61010361022f565b604051600160a060020a03909116815260200160405180910390f35b341561012a57600080fd5b61017060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061023e95505050505050565b005b341561017d57600080fd5b610170600160a060020a0360043516610270565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102275780601f106101fc57610100808354040283529160200191610227565b820191906000526020600020905b81548152906001019060200180831161020a57829003601f168201915b505050505081565b600054600160a060020a031681565b60005433600160a060020a0390811691161461025957600080fd5b600181805161026c92916020019061030b565b5050565b60005433600160a060020a0390811691161461028b57600080fd5b600160a060020a03811615156102a057600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061034c57805160ff1916838001178555610379565b82800160010185558215610379579182015b8281111561037957825182559160200191906001019061035e565b50610385929150610389565b5090565b6103a391905b80821115610385576000815560010161038f565b905600a165627a7a7230582028e8bca0dbaffcef20ba0325b836ea524cdcad392990028be59316b82f24da570029`

// DeployContactable deploys a new Ethereum contract, binding an instance of Contactable to it.
func DeployContactable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contactable, error) {
	parsed, err := abi.JSON(strings.NewReader(ContactableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContactableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contactable{ContactableCaller: ContactableCaller{contract: contract}, ContactableTransactor: ContactableTransactor{contract: contract}, ContactableFilterer: ContactableFilterer{contract: contract}}, nil
}

// Contactable is an auto generated Go binding around an Ethereum contract.
type Contactable struct {
	ContactableCaller     // Read-only binding to the contract
	ContactableTransactor // Write-only binding to the contract
	ContactableFilterer   // Log filterer for contract events
}

// ContactableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContactableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContactableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContactableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContactableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContactableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContactableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContactableSession struct {
	Contract     *Contactable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContactableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContactableCallerSession struct {
	Contract *ContactableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContactableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContactableTransactorSession struct {
	Contract     *ContactableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContactableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContactableRaw struct {
	Contract *Contactable // Generic contract binding to access the raw methods on
}

// ContactableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContactableCallerRaw struct {
	Contract *ContactableCaller // Generic read-only contract binding to access the raw methods on
}

// ContactableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContactableTransactorRaw struct {
	Contract *ContactableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContactable creates a new instance of Contactable, bound to a specific deployed contract.
func NewContactable(address common.Address, backend bind.ContractBackend) (*Contactable, error) {
	contract, err := bindContactable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contactable{ContactableCaller: ContactableCaller{contract: contract}, ContactableTransactor: ContactableTransactor{contract: contract}, ContactableFilterer: ContactableFilterer{contract: contract}}, nil
}

// NewContactableCaller creates a new read-only instance of Contactable, bound to a specific deployed contract.
func NewContactableCaller(address common.Address, caller bind.ContractCaller) (*ContactableCaller, error) {
	contract, err := bindContactable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContactableCaller{contract: contract}, nil
}

// NewContactableTransactor creates a new write-only instance of Contactable, bound to a specific deployed contract.
func NewContactableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContactableTransactor, error) {
	contract, err := bindContactable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContactableTransactor{contract: contract}, nil
}

// NewContactableFilterer creates a new log filterer instance of Contactable, bound to a specific deployed contract.
func NewContactableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContactableFilterer, error) {
	contract, err := bindContactable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContactableFilterer{contract: contract}, nil
}

// bindContactable binds a generic wrapper to an already deployed contract.
func bindContactable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContactableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contactable *ContactableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contactable.Contract.ContactableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contactable *ContactableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contactable.Contract.ContactableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contactable *ContactableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contactable.Contract.ContactableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contactable *ContactableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contactable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contactable *ContactableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contactable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contactable *ContactableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contactable.Contract.contract.Transact(opts, method, params...)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_Contactable *ContactableCaller) ContactInformation(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contactable.contract.Call(opts, out, "contactInformation")
	return *ret0, err
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_Contactable *ContactableSession) ContactInformation() (string, error) {
	return _Contactable.Contract.ContactInformation(&_Contactable.CallOpts)
}

// ContactInformation is a free data retrieval call binding the contract method 0x36f7ab5e.
//
// Solidity: function contactInformation() constant returns(string)
func (_Contactable *ContactableCallerSession) ContactInformation() (string, error) {
	return _Contactable.Contract.ContactInformation(&_Contactable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contactable *ContactableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contactable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contactable *ContactableSession) Owner() (common.Address, error) {
	return _Contactable.Contract.Owner(&_Contactable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contactable *ContactableCallerSession) Owner() (common.Address, error) {
	return _Contactable.Contract.Owner(&_Contactable.CallOpts)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(info string) returns()
func (_Contactable *ContactableTransactor) SetContactInformation(opts *bind.TransactOpts, info string) (*types.Transaction, error) {
	return _Contactable.contract.Transact(opts, "setContactInformation", info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(info string) returns()
func (_Contactable *ContactableSession) SetContactInformation(info string) (*types.Transaction, error) {
	return _Contactable.Contract.SetContactInformation(&_Contactable.TransactOpts, info)
}

// SetContactInformation is a paid mutator transaction binding the contract method 0xb967a52e.
//
// Solidity: function setContactInformation(info string) returns()
func (_Contactable *ContactableTransactorSession) SetContactInformation(info string) (*types.Transaction, error) {
	return _Contactable.Contract.SetContactInformation(&_Contactable.TransactOpts, info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Contactable *ContactableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contactable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Contactable *ContactableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contactable.Contract.TransferOwnership(&_Contactable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Contactable *ContactableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contactable.Contract.TransferOwnership(&_Contactable.TransactOpts, newOwner)
}

// ContactableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contactable contract.
type ContactableOwnershipTransferredIterator struct {
	Event *ContactableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContactableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContactableOwnershipTransferred)
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
		it.Event = new(ContactableOwnershipTransferred)
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
func (it *ContactableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContactableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContactableOwnershipTransferred represents a OwnershipTransferred event raised by the Contactable contract.
type ContactableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Contactable *ContactableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContactableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contactable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContactableOwnershipTransferredIterator{contract: _Contactable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Contactable *ContactableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContactableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contactable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContactableOwnershipTransferred)
				if err := _Contactable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
