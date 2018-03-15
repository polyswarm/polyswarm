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

// PausableMockABI is the input ABI used to generate the binding from.
const PausableMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drasticMeasureTaken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drasticMeasure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"normalProcess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PausableMockBin is the compiled bytecode used for deploying new contracts.
const PausableMockBin = `60606040526000805460a060020a60ff0219169055341561001f57600080fd5b60008054600160a060020a03191633600160a060020a03161760a860020a60ff0219168155600181905561041b90819061005990396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306661abd811461009d5780633f4ba83a146100c25780635c975abb146100d757806376657b8e146100fe5780638456cb59146101115780638da5cb5b146101245780639958f04514610153578063e7651d7a14610166578063f2fde38b14610179575b600080fd5b34156100a857600080fd5b6100b0610198565b60405190815260200160405180910390f35b34156100cd57600080fd5b6100d561019e565b005b34156100e257600080fd5b6100ea61021d565b604051901515815260200160405180910390f35b341561010957600080fd5b6100ea61022d565b341561011c57600080fd5b6100d561024f565b341561012f57600080fd5b6101376102d3565b604051600160a060020a03909116815260200160405180910390f35b341561015e57600080fd5b6100d56102e2565b341561017157600080fd5b6100d5610333565b341561018457600080fd5b6100d5600160a060020a0360043516610354565b60015481565b60005433600160a060020a039081169116146101b957600080fd5b60005460a060020a900460ff1615156101d157600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60005460a060020a900460ff1681565b6000547501000000000000000000000000000000000000000000900460ff1681565b60005433600160a060020a0390811691161461026a57600080fd5b60005460a060020a900460ff161561028157600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b60005460a060020a900460ff1615156102fa57600080fd5b6000805475ff00000000000000000000000000000000000000000019167501000000000000000000000000000000000000000000179055565b60005460a060020a900460ff161561034a57600080fd5b6001805481019055565b60005433600160a060020a0390811691161461036f57600080fd5b600160a060020a038116151561038457600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058203c58e834786c24036daeb95e058eddc7ac96d806342f79e95907eab9020bced30029`

// DeployPausableMock deploys a new Ethereum contract, binding an instance of PausableMock to it.
func DeployPausableMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausableMock, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableMock{PausableMockCaller: PausableMockCaller{contract: contract}, PausableMockTransactor: PausableMockTransactor{contract: contract}, PausableMockFilterer: PausableMockFilterer{contract: contract}}, nil
}

// PausableMock is an auto generated Go binding around an Ethereum contract.
type PausableMock struct {
	PausableMockCaller     // Read-only binding to the contract
	PausableMockTransactor // Write-only binding to the contract
	PausableMockFilterer   // Log filterer for contract events
}

// PausableMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableMockSession struct {
	Contract     *PausableMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableMockCallerSession struct {
	Contract *PausableMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PausableMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableMockTransactorSession struct {
	Contract     *PausableMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PausableMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableMockRaw struct {
	Contract *PausableMock // Generic contract binding to access the raw methods on
}

// PausableMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableMockCallerRaw struct {
	Contract *PausableMockCaller // Generic read-only contract binding to access the raw methods on
}

// PausableMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableMockTransactorRaw struct {
	Contract *PausableMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableMock creates a new instance of PausableMock, bound to a specific deployed contract.
func NewPausableMock(address common.Address, backend bind.ContractBackend) (*PausableMock, error) {
	contract, err := bindPausableMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableMock{PausableMockCaller: PausableMockCaller{contract: contract}, PausableMockTransactor: PausableMockTransactor{contract: contract}, PausableMockFilterer: PausableMockFilterer{contract: contract}}, nil
}

// NewPausableMockCaller creates a new read-only instance of PausableMock, bound to a specific deployed contract.
func NewPausableMockCaller(address common.Address, caller bind.ContractCaller) (*PausableMockCaller, error) {
	contract, err := bindPausableMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableMockCaller{contract: contract}, nil
}

// NewPausableMockTransactor creates a new write-only instance of PausableMock, bound to a specific deployed contract.
func NewPausableMockTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableMockTransactor, error) {
	contract, err := bindPausableMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableMockTransactor{contract: contract}, nil
}

// NewPausableMockFilterer creates a new log filterer instance of PausableMock, bound to a specific deployed contract.
func NewPausableMockFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableMockFilterer, error) {
	contract, err := bindPausableMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableMockFilterer{contract: contract}, nil
}

// bindPausableMock binds a generic wrapper to an already deployed contract.
func bindPausableMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableMock *PausableMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableMock.Contract.PausableMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableMock *PausableMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableMock.Contract.PausableMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableMock *PausableMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableMock.Contract.PausableMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableMock *PausableMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableMock *PausableMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableMock *PausableMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableMock.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_PausableMock *PausableMockCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableMock.contract.Call(opts, out, "count")
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_PausableMock *PausableMockSession) Count() (*big.Int, error) {
	return _PausableMock.Contract.Count(&_PausableMock.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_PausableMock *PausableMockCallerSession) Count() (*big.Int, error) {
	return _PausableMock.Contract.Count(&_PausableMock.CallOpts)
}

// DrasticMeasureTaken is a free data retrieval call binding the contract method 0x76657b8e.
//
// Solidity: function drasticMeasureTaken() constant returns(bool)
func (_PausableMock *PausableMockCaller) DrasticMeasureTaken(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableMock.contract.Call(opts, out, "drasticMeasureTaken")
	return *ret0, err
}

// DrasticMeasureTaken is a free data retrieval call binding the contract method 0x76657b8e.
//
// Solidity: function drasticMeasureTaken() constant returns(bool)
func (_PausableMock *PausableMockSession) DrasticMeasureTaken() (bool, error) {
	return _PausableMock.Contract.DrasticMeasureTaken(&_PausableMock.CallOpts)
}

// DrasticMeasureTaken is a free data retrieval call binding the contract method 0x76657b8e.
//
// Solidity: function drasticMeasureTaken() constant returns(bool)
func (_PausableMock *PausableMockCallerSession) DrasticMeasureTaken() (bool, error) {
	return _PausableMock.Contract.DrasticMeasureTaken(&_PausableMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableMock *PausableMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableMock.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableMock *PausableMockSession) Owner() (common.Address, error) {
	return _PausableMock.Contract.Owner(&_PausableMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableMock *PausableMockCallerSession) Owner() (common.Address, error) {
	return _PausableMock.Contract.Owner(&_PausableMock.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableMock *PausableMockCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableMock.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableMock *PausableMockSession) Paused() (bool, error) {
	return _PausableMock.Contract.Paused(&_PausableMock.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableMock *PausableMockCallerSession) Paused() (bool, error) {
	return _PausableMock.Contract.Paused(&_PausableMock.CallOpts)
}

// DrasticMeasure is a paid mutator transaction binding the contract method 0x9958f045.
//
// Solidity: function drasticMeasure() returns()
func (_PausableMock *PausableMockTransactor) DrasticMeasure(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableMock.contract.Transact(opts, "drasticMeasure")
}

// DrasticMeasure is a paid mutator transaction binding the contract method 0x9958f045.
//
// Solidity: function drasticMeasure() returns()
func (_PausableMock *PausableMockSession) DrasticMeasure() (*types.Transaction, error) {
	return _PausableMock.Contract.DrasticMeasure(&_PausableMock.TransactOpts)
}

// DrasticMeasure is a paid mutator transaction binding the contract method 0x9958f045.
//
// Solidity: function drasticMeasure() returns()
func (_PausableMock *PausableMockTransactorSession) DrasticMeasure() (*types.Transaction, error) {
	return _PausableMock.Contract.DrasticMeasure(&_PausableMock.TransactOpts)
}

// NormalProcess is a paid mutator transaction binding the contract method 0xe7651d7a.
//
// Solidity: function normalProcess() returns()
func (_PausableMock *PausableMockTransactor) NormalProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableMock.contract.Transact(opts, "normalProcess")
}

// NormalProcess is a paid mutator transaction binding the contract method 0xe7651d7a.
//
// Solidity: function normalProcess() returns()
func (_PausableMock *PausableMockSession) NormalProcess() (*types.Transaction, error) {
	return _PausableMock.Contract.NormalProcess(&_PausableMock.TransactOpts)
}

// NormalProcess is a paid mutator transaction binding the contract method 0xe7651d7a.
//
// Solidity: function normalProcess() returns()
func (_PausableMock *PausableMockTransactorSession) NormalProcess() (*types.Transaction, error) {
	return _PausableMock.Contract.NormalProcess(&_PausableMock.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableMock *PausableMockTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableMock.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableMock *PausableMockSession) Pause() (*types.Transaction, error) {
	return _PausableMock.Contract.Pause(&_PausableMock.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableMock *PausableMockTransactorSession) Pause() (*types.Transaction, error) {
	return _PausableMock.Contract.Pause(&_PausableMock.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableMock *PausableMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PausableMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableMock *PausableMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableMock.Contract.TransferOwnership(&_PausableMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableMock *PausableMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableMock.Contract.TransferOwnership(&_PausableMock.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableMock *PausableMockTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableMock.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableMock *PausableMockSession) Unpause() (*types.Transaction, error) {
	return _PausableMock.Contract.Unpause(&_PausableMock.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableMock *PausableMockTransactorSession) Unpause() (*types.Transaction, error) {
	return _PausableMock.Contract.Unpause(&_PausableMock.TransactOpts)
}

// PausableMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PausableMock contract.
type PausableMockOwnershipTransferredIterator struct {
	Event *PausableMockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableMockOwnershipTransferred)
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
		it.Event = new(PausableMockOwnershipTransferred)
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
func (it *PausableMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableMockOwnershipTransferred represents a OwnershipTransferred event raised by the PausableMock contract.
type PausableMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableMock *PausableMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableMockOwnershipTransferredIterator{contract: _PausableMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableMock *PausableMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableMockOwnershipTransferred)
				if err := _PausableMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableMockPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PausableMock contract.
type PausableMockPauseIterator struct {
	Event *PausableMockPause // Event containing the contract specifics and raw log

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
func (it *PausableMockPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableMockPause)
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
		it.Event = new(PausableMockPause)
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
func (it *PausableMockPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableMockPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableMockPause represents a Pause event raised by the PausableMock contract.
type PausableMockPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PausableMock *PausableMockFilterer) FilterPause(opts *bind.FilterOpts) (*PausableMockPauseIterator, error) {

	logs, sub, err := _PausableMock.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausableMockPauseIterator{contract: _PausableMock.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PausableMock *PausableMockFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausableMockPause) (event.Subscription, error) {

	logs, sub, err := _PausableMock.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableMockPause)
				if err := _PausableMock.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableMockUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PausableMock contract.
type PausableMockUnpauseIterator struct {
	Event *PausableMockUnpause // Event containing the contract specifics and raw log

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
func (it *PausableMockUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableMockUnpause)
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
		it.Event = new(PausableMockUnpause)
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
func (it *PausableMockUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableMockUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableMockUnpause represents a Unpause event raised by the PausableMock contract.
type PausableMockUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PausableMock *PausableMockFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableMockUnpauseIterator, error) {

	logs, sub, err := _PausableMock.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableMockUnpauseIterator{contract: _PausableMock.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PausableMock *PausableMockFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableMockUnpause) (event.Subscription, error) {

	logs, sub, err := _PausableMock.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableMockUnpause)
				if err := _PausableMock.contract.UnpackLog(event, "Unpause", log); err != nil {
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
