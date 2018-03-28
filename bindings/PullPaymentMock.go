// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PullPaymentMockABI is the input ABI used to generate the binding from.
const PullPaymentMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawPayments\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"callSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// PullPaymentMockBin is the compiled bytecode used for deploying new contracts.
const PullPaymentMockBin = `606060405261024b806100136000396000f3006060604052600436106100605763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416625b448781146100655780636103d70b1461008a578063752c56281461009f578063e2982c21146100c1575b600080fd5b341561007057600080fd5b6100786100e0565b60405190815260200160405180910390f35b341561009557600080fd5b61009d6100e6565b005b34156100aa57600080fd5b61009d600160a060020a036004351660243561017f565b34156100cc57600080fd5b610078600160a060020a0360043516610189565b60015481565b33600160a060020a03811660009081526020819052604090205480151561010c57600080fd5b600160a060020a033016318190101561012457600080fd5b600154610137908263ffffffff61019b16565b600155600160a060020a0382166000818152602081905260408082209190915582156108fc0290839051600060405180830381858888f19350505050151561017b57fe5b5050565b61017b82826101ad565b60006020819052908152604090205481565b6000828211156101a757fe5b50900390565b600160a060020a0382166000908152602081905260409020546101d6908263ffffffff61020916565b600160a060020a038316600090815260208190526040902055600154610202908263ffffffff61020916565b6001555050565b60008282018381101561021857fe5b93925050505600a165627a7a72305820a2b72472a3c616305f21f10a435ecd9267a7ac0e7fdb168e98458b860298030b0029`

// DeployPullPaymentMock deploys a new Ethereum contract, binding an instance of PullPaymentMock to it.
func DeployPullPaymentMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PullPaymentMock, error) {
	parsed, err := abi.JSON(strings.NewReader(PullPaymentMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PullPaymentMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PullPaymentMock{PullPaymentMockCaller: PullPaymentMockCaller{contract: contract}, PullPaymentMockTransactor: PullPaymentMockTransactor{contract: contract}}, nil
}

// PullPaymentMock is an auto generated Go binding around an Ethereum contract.
type PullPaymentMock struct {
	PullPaymentMockCaller     // Read-only binding to the contract
	PullPaymentMockTransactor // Write-only binding to the contract
}

// PullPaymentMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PullPaymentMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullPaymentMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PullPaymentMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullPaymentMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PullPaymentMockSession struct {
	Contract     *PullPaymentMock  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PullPaymentMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PullPaymentMockCallerSession struct {
	Contract *PullPaymentMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PullPaymentMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PullPaymentMockTransactorSession struct {
	Contract     *PullPaymentMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PullPaymentMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PullPaymentMockRaw struct {
	Contract *PullPaymentMock // Generic contract binding to access the raw methods on
}

// PullPaymentMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PullPaymentMockCallerRaw struct {
	Contract *PullPaymentMockCaller // Generic read-only contract binding to access the raw methods on
}

// PullPaymentMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PullPaymentMockTransactorRaw struct {
	Contract *PullPaymentMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPullPaymentMock creates a new instance of PullPaymentMock, bound to a specific deployed contract.
func NewPullPaymentMock(address common.Address, backend bind.ContractBackend) (*PullPaymentMock, error) {
	contract, err := bindPullPaymentMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PullPaymentMock{PullPaymentMockCaller: PullPaymentMockCaller{contract: contract}, PullPaymentMockTransactor: PullPaymentMockTransactor{contract: contract}}, nil
}

// NewPullPaymentMockCaller creates a new read-only instance of PullPaymentMock, bound to a specific deployed contract.
func NewPullPaymentMockCaller(address common.Address, caller bind.ContractCaller) (*PullPaymentMockCaller, error) {
	contract, err := bindPullPaymentMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PullPaymentMockCaller{contract: contract}, nil
}

// NewPullPaymentMockTransactor creates a new write-only instance of PullPaymentMock, bound to a specific deployed contract.
func NewPullPaymentMockTransactor(address common.Address, transactor bind.ContractTransactor) (*PullPaymentMockTransactor, error) {
	contract, err := bindPullPaymentMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PullPaymentMockTransactor{contract: contract}, nil
}

// bindPullPaymentMock binds a generic wrapper to an already deployed contract.
func bindPullPaymentMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PullPaymentMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PullPaymentMock *PullPaymentMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PullPaymentMock.Contract.PullPaymentMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PullPaymentMock *PullPaymentMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullPaymentMock.Contract.PullPaymentMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PullPaymentMock *PullPaymentMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PullPaymentMock.Contract.PullPaymentMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PullPaymentMock *PullPaymentMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PullPaymentMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PullPaymentMock *PullPaymentMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullPaymentMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PullPaymentMock *PullPaymentMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PullPaymentMock.Contract.contract.Transact(opts, method, params...)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_PullPaymentMock *PullPaymentMockCaller) Payments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PullPaymentMock.contract.Call(opts, out, "payments", arg0)
	return *ret0, err
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_PullPaymentMock *PullPaymentMockSession) Payments(arg0 common.Address) (*big.Int, error) {
	return _PullPaymentMock.Contract.Payments(&_PullPaymentMock.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_PullPaymentMock *PullPaymentMockCallerSession) Payments(arg0 common.Address) (*big.Int, error) {
	return _PullPaymentMock.Contract.Payments(&_PullPaymentMock.CallOpts, arg0)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_PullPaymentMock *PullPaymentMockCaller) TotalPayments(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PullPaymentMock.contract.Call(opts, out, "totalPayments")
	return *ret0, err
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_PullPaymentMock *PullPaymentMockSession) TotalPayments() (*big.Int, error) {
	return _PullPaymentMock.Contract.TotalPayments(&_PullPaymentMock.CallOpts)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_PullPaymentMock *PullPaymentMockCallerSession) TotalPayments() (*big.Int, error) {
	return _PullPaymentMock.Contract.TotalPayments(&_PullPaymentMock.CallOpts)
}

// CallSend is a paid mutator transaction binding the contract method 0x752c5628.
//
// Solidity: function callSend(dest address, amount uint256) returns()
func (_PullPaymentMock *PullPaymentMockTransactor) CallSend(opts *bind.TransactOpts, dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PullPaymentMock.contract.Transact(opts, "callSend", dest, amount)
}

// CallSend is a paid mutator transaction binding the contract method 0x752c5628.
//
// Solidity: function callSend(dest address, amount uint256) returns()
func (_PullPaymentMock *PullPaymentMockSession) CallSend(dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PullPaymentMock.Contract.CallSend(&_PullPaymentMock.TransactOpts, dest, amount)
}

// CallSend is a paid mutator transaction binding the contract method 0x752c5628.
//
// Solidity: function callSend(dest address, amount uint256) returns()
func (_PullPaymentMock *PullPaymentMockTransactorSession) CallSend(dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PullPaymentMock.Contract.CallSend(&_PullPaymentMock.TransactOpts, dest, amount)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_PullPaymentMock *PullPaymentMockTransactor) WithdrawPayments(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullPaymentMock.contract.Transact(opts, "withdrawPayments")
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_PullPaymentMock *PullPaymentMockSession) WithdrawPayments() (*types.Transaction, error) {
	return _PullPaymentMock.Contract.WithdrawPayments(&_PullPaymentMock.TransactOpts)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_PullPaymentMock *PullPaymentMockTransactorSession) WithdrawPayments() (*types.Transaction, error) {
	return _PullPaymentMock.Contract.WithdrawPayments(&_PullPaymentMock.TransactOpts)
}
