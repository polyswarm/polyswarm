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

// ERC20FailingMockABI is the input ABI used to generate the binding from.
const ERC20FailingMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20FailingMockBin is the compiled bytecode used for deploying new contracts.
const ERC20FailingMockBin = `6060604052341561000f57600080fd5b6101bc8061001e6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100bf57806323b872dd146100e457806370a0823114610119578063a9059cbb1461007c578063dd62ed3e14610145575b600080fd5b341561008757600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043516602435610173565b604051901515815260200160405180910390f35b34156100ca57600080fd5b6100d261017b565b60405190815260200160405180910390f35b34156100ef57600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610181565b341561012457600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff6004351661018a565b341561015057600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff600435811690602435165b600092915050565b60005481565b60009392505050565b506000905600a165627a7a72305820ff513700d105bf612803b32e105889683c8c56d5b85d90d52b7b4a62e40776310029`

// DeployERC20FailingMock deploys a new Ethereum contract, binding an instance of ERC20FailingMock to it.
func DeployERC20FailingMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20FailingMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20FailingMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20FailingMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20FailingMock{ERC20FailingMockCaller: ERC20FailingMockCaller{contract: contract}, ERC20FailingMockTransactor: ERC20FailingMockTransactor{contract: contract}, ERC20FailingMockFilterer: ERC20FailingMockFilterer{contract: contract}}, nil
}

// ERC20FailingMock is an auto generated Go binding around an Ethereum contract.
type ERC20FailingMock struct {
	ERC20FailingMockCaller     // Read-only binding to the contract
	ERC20FailingMockTransactor // Write-only binding to the contract
	ERC20FailingMockFilterer   // Log filterer for contract events
}

// ERC20FailingMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20FailingMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20FailingMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20FailingMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20FailingMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20FailingMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20FailingMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20FailingMockSession struct {
	Contract     *ERC20FailingMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20FailingMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20FailingMockCallerSession struct {
	Contract *ERC20FailingMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20FailingMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20FailingMockTransactorSession struct {
	Contract     *ERC20FailingMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20FailingMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20FailingMockRaw struct {
	Contract *ERC20FailingMock // Generic contract binding to access the raw methods on
}

// ERC20FailingMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20FailingMockCallerRaw struct {
	Contract *ERC20FailingMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20FailingMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20FailingMockTransactorRaw struct {
	Contract *ERC20FailingMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20FailingMock creates a new instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMock(address common.Address, backend bind.ContractBackend) (*ERC20FailingMock, error) {
	contract, err := bindERC20FailingMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMock{ERC20FailingMockCaller: ERC20FailingMockCaller{contract: contract}, ERC20FailingMockTransactor: ERC20FailingMockTransactor{contract: contract}, ERC20FailingMockFilterer: ERC20FailingMockFilterer{contract: contract}}, nil
}

// NewERC20FailingMockCaller creates a new read-only instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMockCaller(address common.Address, caller bind.ContractCaller) (*ERC20FailingMockCaller, error) {
	contract, err := bindERC20FailingMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockCaller{contract: contract}, nil
}

// NewERC20FailingMockTransactor creates a new write-only instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20FailingMockTransactor, error) {
	contract, err := bindERC20FailingMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockTransactor{contract: contract}, nil
}

// NewERC20FailingMockFilterer creates a new log filterer instance of ERC20FailingMock, bound to a specific deployed contract.
func NewERC20FailingMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20FailingMockFilterer, error) {
	contract, err := bindERC20FailingMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockFilterer{contract: contract}, nil
}

// bindERC20FailingMock binds a generic wrapper to an already deployed contract.
func bindERC20FailingMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20FailingMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20FailingMock *ERC20FailingMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20FailingMock.Contract.ERC20FailingMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20FailingMock *ERC20FailingMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.ERC20FailingMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20FailingMock *ERC20FailingMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.ERC20FailingMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20FailingMock *ERC20FailingMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20FailingMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20FailingMock *ERC20FailingMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20FailingMock *ERC20FailingMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20FailingMock.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.Allowance(&_ERC20FailingMock.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.Allowance(&_ERC20FailingMock.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20FailingMock.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.BalanceOf(&_ERC20FailingMock.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20FailingMock.Contract.BalanceOf(&_ERC20FailingMock.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20FailingMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockSession) TotalSupply() (*big.Int, error) {
	return _ERC20FailingMock.Contract.TotalSupply(&_ERC20FailingMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20FailingMock *ERC20FailingMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20FailingMock.Contract.TotalSupply(&_ERC20FailingMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Approve(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Approve(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Transfer(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer( address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.Transfer(&_ERC20FailingMock.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.TransferFrom(&_ERC20FailingMock.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom( address,  address,  uint256) returns(bool)
func (_ERC20FailingMock *ERC20FailingMockTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _ERC20FailingMock.Contract.TransferFrom(&_ERC20FailingMock.TransactOpts, arg0, arg1, arg2)
}

// ERC20FailingMockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20FailingMock contract.
type ERC20FailingMockApprovalIterator struct {
	Event *ERC20FailingMockApproval // Event containing the contract specifics and raw log

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
func (it *ERC20FailingMockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20FailingMockApproval)
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
		it.Event = new(ERC20FailingMockApproval)
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
func (it *ERC20FailingMockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20FailingMockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20FailingMockApproval represents a Approval event raised by the ERC20FailingMock contract.
type ERC20FailingMockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20FailingMock *ERC20FailingMockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20FailingMockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20FailingMock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockApprovalIterator{contract: _ERC20FailingMock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20FailingMock *ERC20FailingMockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20FailingMockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20FailingMock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20FailingMockApproval)
				if err := _ERC20FailingMock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20FailingMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20FailingMock contract.
type ERC20FailingMockTransferIterator struct {
	Event *ERC20FailingMockTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20FailingMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20FailingMockTransfer)
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
		it.Event = new(ERC20FailingMockTransfer)
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
func (it *ERC20FailingMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20FailingMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20FailingMockTransfer represents a Transfer event raised by the ERC20FailingMock contract.
type ERC20FailingMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20FailingMock *ERC20FailingMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20FailingMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20FailingMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20FailingMockTransferIterator{contract: _ERC20FailingMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20FailingMock *ERC20FailingMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20FailingMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20FailingMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20FailingMockTransfer)
				if err := _ERC20FailingMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
