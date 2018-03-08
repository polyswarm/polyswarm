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

// HasNoTokensABI is the input ABI used to generate the binding from.
const HasNoTokensABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoTokensBin is the compiled bytecode used for deploying new contracts.
const HasNoTokensBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556102fa806100306000396000f3006060604052600436106100485763ffffffff60e060020a60003504166317ffc320811461004d5780638da5cb5b1461006e578063c0ee0b8a1461009d578063f2fde38b146100cc575b600080fd5b341561005857600080fd5b61006c600160a060020a03600435166100eb565b005b341561007957600080fd5b61008161019f565b604051600160a060020a03909116815260200160405180910390f35b34156100a857600080fd5b61006c60048035600160a060020a0316906024803591604435918201910135610048565b34156100d757600080fd5b61006c600160a060020a03600435166101ae565b6000805433600160a060020a0390811691161461010757600080fd5b81600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561015e57600080fd5b6102c65a03f1151561016f57600080fd5b505050604051805160005490925061019b9150600160a060020a0384811691168363ffffffff61024916565b5050565b600054600160a060020a031681565b60005433600160a060020a039081169116146101c957600080fd5b600160a060020a03811615156101de57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156102a657600080fd5b6102c65a03f115156102b757600080fd5b5050506040518051905015156102c957fe5b5050505600a165627a7a72305820a00ebb3b6614ec53bc2a4b94856e82b373263b12f9ef1a415fbb2ce412f2491e0029`

// DeployHasNoTokens deploys a new Ethereum contract, binding an instance of HasNoTokens to it.
func DeployHasNoTokens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoTokensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoTokens{HasNoTokensCaller: HasNoTokensCaller{contract: contract}, HasNoTokensTransactor: HasNoTokensTransactor{contract: contract}, HasNoTokensFilterer: HasNoTokensFilterer{contract: contract}}, nil
}

// HasNoTokens is an auto generated Go binding around an Ethereum contract.
type HasNoTokens struct {
	HasNoTokensCaller     // Read-only binding to the contract
	HasNoTokensTransactor // Write-only binding to the contract
	HasNoTokensFilterer   // Log filterer for contract events
}

// HasNoTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasNoTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoTokensSession struct {
	Contract     *HasNoTokens      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoTokensCallerSession struct {
	Contract *HasNoTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HasNoTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoTokensTransactorSession struct {
	Contract     *HasNoTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HasNoTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoTokensRaw struct {
	Contract *HasNoTokens // Generic contract binding to access the raw methods on
}

// HasNoTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoTokensCallerRaw struct {
	Contract *HasNoTokensCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoTokensTransactorRaw struct {
	Contract *HasNoTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoTokens creates a new instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokens(address common.Address, backend bind.ContractBackend) (*HasNoTokens, error) {
	contract, err := bindHasNoTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoTokens{HasNoTokensCaller: HasNoTokensCaller{contract: contract}, HasNoTokensTransactor: HasNoTokensTransactor{contract: contract}, HasNoTokensFilterer: HasNoTokensFilterer{contract: contract}}, nil
}

// NewHasNoTokensCaller creates a new read-only instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensCaller(address common.Address, caller bind.ContractCaller) (*HasNoTokensCaller, error) {
	contract, err := bindHasNoTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensCaller{contract: contract}, nil
}

// NewHasNoTokensTransactor creates a new write-only instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoTokensTransactor, error) {
	contract, err := bindHasNoTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensTransactor{contract: contract}, nil
}

// NewHasNoTokensFilterer creates a new log filterer instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*HasNoTokensFilterer, error) {
	contract, err := bindHasNoTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensFilterer{contract: contract}, nil
}

// bindHasNoTokens binds a generic wrapper to an already deployed contract.
func bindHasNoTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoTokens *HasNoTokensRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoTokens.Contract.HasNoTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoTokens *HasNoTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoTokens.Contract.HasNoTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoTokens *HasNoTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoTokens.Contract.HasNoTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoTokens *HasNoTokensCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoTokens *HasNoTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoTokens *HasNoTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoTokens.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoTokens.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensSession) Owner() (common.Address, error) {
	return _HasNoTokens.Contract.Owner(&_HasNoTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensCallerSession) Owner() (common.Address, error) {
	return _HasNoTokens.Contract.Owner(&_HasNoTokens.CallOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.ReclaimToken(&_HasNoTokens.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.ReclaimToken(&_HasNoTokens.TransactOpts, token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TokenFallback(&_HasNoTokens.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TokenFallback(&_HasNoTokens.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TransferOwnership(&_HasNoTokens.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TransferOwnership(&_HasNoTokens.TransactOpts, newOwner)
}

// HasNoTokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HasNoTokens contract.
type HasNoTokensOwnershipTransferredIterator struct {
	Event *HasNoTokensOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HasNoTokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HasNoTokensOwnershipTransferred)
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
		it.Event = new(HasNoTokensOwnershipTransferred)
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
func (it *HasNoTokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HasNoTokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HasNoTokensOwnershipTransferred represents a OwnershipTransferred event raised by the HasNoTokens contract.
type HasNoTokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoTokens *HasNoTokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HasNoTokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoTokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensOwnershipTransferredIterator{contract: _HasNoTokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoTokens *HasNoTokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HasNoTokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoTokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HasNoTokensOwnershipTransferred)
				if err := _HasNoTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
