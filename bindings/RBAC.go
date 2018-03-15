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

// RBACABI is the input ABI used to generate the binding from.
const RBACABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"adminRemoveRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"adminAddRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// RBACBin is the compiled bytecode used for deploying new contracts.
const RBACBin = `6060604052341561000f57600080fd5b6100593360408051908101604052600581527f61646d696e00000000000000000000000000000000000000000000000000000060208201526401000000006105a261005e82021704565b6101ad565b6100db826000836040518082805190602001908083835b602083106100945780518252601f199092019160209182019101610075565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020906401000000006106a761018882021704565b7fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b7004898282604051600160a060020a038316815260406020820181815290820183818151815260200191508051906020019080838360005b83811015610149578082015183820152602001610131565b50505050905090810190601f1680156101765780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6106f8806101bc6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c8114610071578063217fe6c6146100d257806388cee87e14610145578063b25fa92c146101a4578063d391014b14610203575b600080fd5b341561007c57600080fd5b6100d060048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061028d95505050505050565b005b34156100dd57600080fd5b61013160048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061030795505050505050565b604051901515815260200160405180910390f35b341561015057600080fd5b6100d060048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061038695505050505050565b34156101af57600080fd5b6100d060048035600160a060020a03169060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506103cd95505050505050565b341561020e57600080fd5b610216610414565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561025257808201518382015260200161023a565b50505050905090810190601f16801561027f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610303826000836040518082805190602001908083835b602083106102c35780518252601f1990920191602091820191016102a4565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61044b16565b5050565b600061037f836000846040518082805190602001908083835b6020831061033f5780518252601f199092019160209182019101610320565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61046016565b9392505050565b6103c33360408051908101604052600581527f61646d696e000000000000000000000000000000000000000000000000000000602082015261028d565b610303828261047f565b61040a3360408051908101604052600581527f61646d696e000000000000000000000000000000000000000000000000000000602082015261028d565b61030382826105a2565b60408051908101604052600581527f61646d696e000000000000000000000000000000000000000000000000000000602082015281565b6104558282610460565b151561030357600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b6104f5826000836040518082805190602001908083835b602083106104b55780518252601f199092019160209182019101610496565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61068516565b7fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a8282604051600160a060020a038316815260406020820181815290820183818151815260200191508051906020019080838360005b8381101561056357808201518382015260200161054b565b50505050905090810190601f1680156105905780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b610618826000836040518082805190602001908083835b602083106105d85780518252601f1990920191602091820191016105b9565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff6106a716565b7fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b7004898282604051600160a060020a038316815260406020820181815290820183818151815260200191508051906020019080838360008381101561056357808201518382015260200161054b565b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0316600090815260209190915260409020805460ff191660011790555600a165627a7a723058203abf707a2582145d8c0a9d7a171054e634b2198ab25b625727e8be2058edfeb70029`

// DeployRBAC deploys a new Ethereum contract, binding an instance of RBAC to it.
func DeployRBAC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RBAC, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RBACBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// RBAC is an auto generated Go binding around an Ethereum contract.
type RBAC struct {
	RBACCaller     // Read-only binding to the contract
	RBACTransactor // Write-only binding to the contract
	RBACFilterer   // Log filterer for contract events
}

// RBACCaller is an auto generated read-only Go binding around an Ethereum contract.
type RBACCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RBACTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RBACFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RBACSession struct {
	Contract     *RBAC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RBACCallerSession struct {
	Contract *RBACCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RBACTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RBACTransactorSession struct {
	Contract     *RBACTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACRaw is an auto generated low-level Go binding around an Ethereum contract.
type RBACRaw struct {
	Contract *RBAC // Generic contract binding to access the raw methods on
}

// RBACCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RBACCallerRaw struct {
	Contract *RBACCaller // Generic read-only contract binding to access the raw methods on
}

// RBACTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RBACTransactorRaw struct {
	Contract *RBACTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRBAC creates a new instance of RBAC, bound to a specific deployed contract.
func NewRBAC(address common.Address, backend bind.ContractBackend) (*RBAC, error) {
	contract, err := bindRBAC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// NewRBACCaller creates a new read-only instance of RBAC, bound to a specific deployed contract.
func NewRBACCaller(address common.Address, caller bind.ContractCaller) (*RBACCaller, error) {
	contract, err := bindRBAC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RBACCaller{contract: contract}, nil
}

// NewRBACTransactor creates a new write-only instance of RBAC, bound to a specific deployed contract.
func NewRBACTransactor(address common.Address, transactor bind.ContractTransactor) (*RBACTransactor, error) {
	contract, err := bindRBAC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RBACTransactor{contract: contract}, nil
}

// NewRBACFilterer creates a new log filterer instance of RBAC, bound to a specific deployed contract.
func NewRBACFilterer(address common.Address, filterer bind.ContractFilterer) (*RBACFilterer, error) {
	contract, err := bindRBAC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RBACFilterer{contract: contract}, nil
}

// bindRBAC binds a generic wrapper to an already deployed contract.
func bindRBAC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.RBACCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transact(opts, method, params...)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_RBAC *RBACCaller) ROLE_ADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RBAC.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_RBAC *RBACSession) ROLE_ADMIN() (string, error) {
	return _RBAC.Contract.ROLE_ADMIN(&_RBAC.CallOpts)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_RBAC *RBACCallerSession) ROLE_ADMIN() (string, error) {
	return _RBAC.Contract.ROLE_ADMIN(&_RBAC.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBAC *RBACCaller) CheckRole(opts *bind.CallOpts, addr common.Address, roleName string) error {
	var ()
	out := &[]interface{}{}
	err := _RBAC.contract.Call(opts, out, "checkRole", addr, roleName)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBAC *RBACSession) CheckRole(addr common.Address, roleName string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, addr, roleName)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBAC *RBACCallerSession) CheckRole(addr common.Address, roleName string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, addr, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBAC *RBACCaller) HasRole(opts *bind.CallOpts, addr common.Address, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBAC.contract.Call(opts, out, "hasRole", addr, roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBAC *RBACSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, addr, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBAC *RBACCallerSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, addr, roleName)
}

// AdminAddRole is a paid mutator transaction binding the contract method 0xb25fa92c.
//
// Solidity: function adminAddRole(addr address, roleName string) returns()
func (_RBAC *RBACTransactor) AdminAddRole(opts *bind.TransactOpts, addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBAC.contract.Transact(opts, "adminAddRole", addr, roleName)
}

// AdminAddRole is a paid mutator transaction binding the contract method 0xb25fa92c.
//
// Solidity: function adminAddRole(addr address, roleName string) returns()
func (_RBAC *RBACSession) AdminAddRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBAC.Contract.AdminAddRole(&_RBAC.TransactOpts, addr, roleName)
}

// AdminAddRole is a paid mutator transaction binding the contract method 0xb25fa92c.
//
// Solidity: function adminAddRole(addr address, roleName string) returns()
func (_RBAC *RBACTransactorSession) AdminAddRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBAC.Contract.AdminAddRole(&_RBAC.TransactOpts, addr, roleName)
}

// AdminRemoveRole is a paid mutator transaction binding the contract method 0x88cee87e.
//
// Solidity: function adminRemoveRole(addr address, roleName string) returns()
func (_RBAC *RBACTransactor) AdminRemoveRole(opts *bind.TransactOpts, addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBAC.contract.Transact(opts, "adminRemoveRole", addr, roleName)
}

// AdminRemoveRole is a paid mutator transaction binding the contract method 0x88cee87e.
//
// Solidity: function adminRemoveRole(addr address, roleName string) returns()
func (_RBAC *RBACSession) AdminRemoveRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBAC.Contract.AdminRemoveRole(&_RBAC.TransactOpts, addr, roleName)
}

// AdminRemoveRole is a paid mutator transaction binding the contract method 0x88cee87e.
//
// Solidity: function adminRemoveRole(addr address, roleName string) returns()
func (_RBAC *RBACTransactorSession) AdminRemoveRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBAC.Contract.AdminRemoveRole(&_RBAC.TransactOpts, addr, roleName)
}

// RBACRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the RBAC contract.
type RBACRoleAddedIterator struct {
	Event *RBACRoleAdded // Event containing the contract specifics and raw log

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
func (it *RBACRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleAdded)
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
		it.Event = new(RBACRoleAdded)
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
func (it *RBACRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleAdded represents a RoleAdded event raised by the RBAC contract.
type RBACRoleAdded struct {
	Addr     common.Address
	RoleName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: event RoleAdded(addr address, roleName string)
func (_RBAC *RBACFilterer) FilterRoleAdded(opts *bind.FilterOpts) (*RBACRoleAddedIterator, error) {

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleAdded")
	if err != nil {
		return nil, err
	}
	return &RBACRoleAddedIterator{contract: _RBAC.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: event RoleAdded(addr address, roleName string)
func (_RBAC *RBACFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *RBACRoleAdded) (event.Subscription, error) {

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleAdded)
				if err := _RBAC.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// RBACRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the RBAC contract.
type RBACRoleRemovedIterator struct {
	Event *RBACRoleRemoved // Event containing the contract specifics and raw log

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
func (it *RBACRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleRemoved)
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
		it.Event = new(RBACRoleRemoved)
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
func (it *RBACRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleRemoved represents a RoleRemoved event raised by the RBAC contract.
type RBACRoleRemoved struct {
	Addr     common.Address
	RoleName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: event RoleRemoved(addr address, roleName string)
func (_RBAC *RBACFilterer) FilterRoleRemoved(opts *bind.FilterOpts) (*RBACRoleRemovedIterator, error) {

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleRemoved")
	if err != nil {
		return nil, err
	}
	return &RBACRoleRemovedIterator{contract: _RBAC.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: event RoleRemoved(addr address, roleName string)
func (_RBAC *RBACFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *RBACRoleRemoved) (event.Subscription, error) {

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleRemoved)
				if err := _RBAC.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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
