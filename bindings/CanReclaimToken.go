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

// CanReclaimTokenABI is the input ABI used to generate the binding from.
const CanReclaimTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CanReclaimTokenBin is the compiled bytecode used for deploying new contracts.
const CanReclaimTokenBin = `606060405260008054600160a060020a033316600160a060020a03199091161790556102a6806100306000396000f30060606040526004361061003d5763ffffffff60e060020a60003504166317ffc32081146100425780638da5cb5b14610063578063f2fde38b14610092575b600080fd5b341561004d57600080fd5b610061600160a060020a03600435166100b1565b005b341561006e57600080fd5b610076610158565b604051600160a060020a03909116815260200160405180910390f35b341561009d57600080fd5b610061600160a060020a0360043516610167565b6000805433600160a060020a039081169116146100cd57600080fd5b81600160a060020a03166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561011b57600080fd5b5af1151561012857600080fd5b50505060405180516000549092506101549150600160a060020a0384811691168363ffffffff61020216565b5050565b600054600160a060020a031681565b60005433600160a060020a0390811691161461018257600080fd5b600160a060020a038116151561019757600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561025657600080fd5b5af1151561026357600080fd5b50505060405180519050151561027557fe5b5050505600a165627a7a72305820ca323eef6e9fc5d4a4dd8f094c96110e6a34362d9e68432fb2ce74890e5401f40029`

// DeployCanReclaimToken deploys a new Ethereum contract, binding an instance of CanReclaimToken to it.
func DeployCanReclaimToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanReclaimToken, error) {
	parsed, err := abi.JSON(strings.NewReader(CanReclaimTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CanReclaimTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanReclaimToken{CanReclaimTokenCaller: CanReclaimTokenCaller{contract: contract}, CanReclaimTokenTransactor: CanReclaimTokenTransactor{contract: contract}, CanReclaimTokenFilterer: CanReclaimTokenFilterer{contract: contract}}, nil
}

// CanReclaimToken is an auto generated Go binding around an Ethereum contract.
type CanReclaimToken struct {
	CanReclaimTokenCaller     // Read-only binding to the contract
	CanReclaimTokenTransactor // Write-only binding to the contract
	CanReclaimTokenFilterer   // Log filterer for contract events
}

// CanReclaimTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanReclaimTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanReclaimTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanReclaimTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanReclaimTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanReclaimTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanReclaimTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanReclaimTokenSession struct {
	Contract     *CanReclaimToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanReclaimTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanReclaimTokenCallerSession struct {
	Contract *CanReclaimTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CanReclaimTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanReclaimTokenTransactorSession struct {
	Contract     *CanReclaimTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CanReclaimTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanReclaimTokenRaw struct {
	Contract *CanReclaimToken // Generic contract binding to access the raw methods on
}

// CanReclaimTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanReclaimTokenCallerRaw struct {
	Contract *CanReclaimTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CanReclaimTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanReclaimTokenTransactorRaw struct {
	Contract *CanReclaimTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanReclaimToken creates a new instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimToken(address common.Address, backend bind.ContractBackend) (*CanReclaimToken, error) {
	contract, err := bindCanReclaimToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanReclaimToken{CanReclaimTokenCaller: CanReclaimTokenCaller{contract: contract}, CanReclaimTokenTransactor: CanReclaimTokenTransactor{contract: contract}, CanReclaimTokenFilterer: CanReclaimTokenFilterer{contract: contract}}, nil
}

// NewCanReclaimTokenCaller creates a new read-only instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimTokenCaller(address common.Address, caller bind.ContractCaller) (*CanReclaimTokenCaller, error) {
	contract, err := bindCanReclaimToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenCaller{contract: contract}, nil
}

// NewCanReclaimTokenTransactor creates a new write-only instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CanReclaimTokenTransactor, error) {
	contract, err := bindCanReclaimToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenTransactor{contract: contract}, nil
}

// NewCanReclaimTokenFilterer creates a new log filterer instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CanReclaimTokenFilterer, error) {
	contract, err := bindCanReclaimToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenFilterer{contract: contract}, nil
}

// bindCanReclaimToken binds a generic wrapper to an already deployed contract.
func bindCanReclaimToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CanReclaimTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanReclaimToken *CanReclaimTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanReclaimToken.Contract.CanReclaimTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanReclaimToken *CanReclaimTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.CanReclaimTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanReclaimToken *CanReclaimTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.CanReclaimTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanReclaimToken *CanReclaimTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanReclaimToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanReclaimToken *CanReclaimTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanReclaimToken *CanReclaimTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanReclaimToken *CanReclaimTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanReclaimToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanReclaimToken *CanReclaimTokenSession) Owner() (common.Address, error) {
	return _CanReclaimToken.Contract.Owner(&_CanReclaimToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanReclaimToken *CanReclaimTokenCallerSession) Owner() (common.Address, error) {
	return _CanReclaimToken.Contract.Owner(&_CanReclaimToken.CallOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_CanReclaimToken *CanReclaimTokenSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.ReclaimToken(&_CanReclaimToken.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.ReclaimToken(&_CanReclaimToken.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanReclaimToken *CanReclaimTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.TransferOwnership(&_CanReclaimToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.TransferOwnership(&_CanReclaimToken.TransactOpts, newOwner)
}

// CanReclaimTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CanReclaimToken contract.
type CanReclaimTokenOwnershipTransferredIterator struct {
	Event *CanReclaimTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CanReclaimTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanReclaimTokenOwnershipTransferred)
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
		it.Event = new(CanReclaimTokenOwnershipTransferred)
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
func (it *CanReclaimTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanReclaimTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanReclaimTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CanReclaimToken contract.
type CanReclaimTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CanReclaimToken *CanReclaimTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CanReclaimTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanReclaimToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenOwnershipTransferredIterator{contract: _CanReclaimToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CanReclaimToken *CanReclaimTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CanReclaimTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanReclaimToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanReclaimTokenOwnershipTransferred)
				if err := _CanReclaimToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
