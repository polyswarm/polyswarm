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

// PullPaymentABI is the input ABI used to generate the binding from.
const PullPaymentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawPayments\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PullPaymentBin is the compiled bytecode used for deploying new contracts.
const PullPaymentBin = `6060604052341561000f57600080fd5b6101d68061001e6000396000f3006060604052600436106100555763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416625b4487811461005a5780636103d70b1461007f578063e2982c2114610094575b600080fd5b341561006557600080fd5b61006d6100c0565b60405190815260200160405180910390f35b341561008a57600080fd5b6100926100c6565b005b341561009f57600080fd5b61006d73ffffffffffffffffffffffffffffffffffffffff60043516610186565b60015481565b3373ffffffffffffffffffffffffffffffffffffffff81166000908152602081905260409020548015156100f957600080fd5b73ffffffffffffffffffffffffffffffffffffffff3016318190101561011e57600080fd5b600154610131908263ffffffff61019816565b60015573ffffffffffffffffffffffffffffffffffffffff82166000818152602081905260408082209190915582156108fc0290839051600060405180830381858888f19350505050151561018257fe5b5050565b60006020819052908152604090205481565b6000828211156101a457fe5b509003905600a165627a7a723058208ecbbbcfbca382580fe57cd74702b8cb85a35374065c5d266faf10fb915547400029`

// DeployPullPayment deploys a new Ethereum contract, binding an instance of PullPayment to it.
func DeployPullPayment(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PullPayment, error) {
	parsed, err := abi.JSON(strings.NewReader(PullPaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PullPaymentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PullPayment{PullPaymentCaller: PullPaymentCaller{contract: contract}, PullPaymentTransactor: PullPaymentTransactor{contract: contract}, PullPaymentFilterer: PullPaymentFilterer{contract: contract}}, nil
}

// PullPayment is an auto generated Go binding around an Ethereum contract.
type PullPayment struct {
	PullPaymentCaller     // Read-only binding to the contract
	PullPaymentTransactor // Write-only binding to the contract
	PullPaymentFilterer   // Log filterer for contract events
}

// PullPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PullPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PullPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PullPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PullPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PullPaymentSession struct {
	Contract     *PullPayment      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PullPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PullPaymentCallerSession struct {
	Contract *PullPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PullPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PullPaymentTransactorSession struct {
	Contract     *PullPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PullPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PullPaymentRaw struct {
	Contract *PullPayment // Generic contract binding to access the raw methods on
}

// PullPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PullPaymentCallerRaw struct {
	Contract *PullPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PullPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PullPaymentTransactorRaw struct {
	Contract *PullPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPullPayment creates a new instance of PullPayment, bound to a specific deployed contract.
func NewPullPayment(address common.Address, backend bind.ContractBackend) (*PullPayment, error) {
	contract, err := bindPullPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PullPayment{PullPaymentCaller: PullPaymentCaller{contract: contract}, PullPaymentTransactor: PullPaymentTransactor{contract: contract}, PullPaymentFilterer: PullPaymentFilterer{contract: contract}}, nil
}

// NewPullPaymentCaller creates a new read-only instance of PullPayment, bound to a specific deployed contract.
func NewPullPaymentCaller(address common.Address, caller bind.ContractCaller) (*PullPaymentCaller, error) {
	contract, err := bindPullPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PullPaymentCaller{contract: contract}, nil
}

// NewPullPaymentTransactor creates a new write-only instance of PullPayment, bound to a specific deployed contract.
func NewPullPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PullPaymentTransactor, error) {
	contract, err := bindPullPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PullPaymentTransactor{contract: contract}, nil
}

// NewPullPaymentFilterer creates a new log filterer instance of PullPayment, bound to a specific deployed contract.
func NewPullPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PullPaymentFilterer, error) {
	contract, err := bindPullPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PullPaymentFilterer{contract: contract}, nil
}

// bindPullPayment binds a generic wrapper to an already deployed contract.
func bindPullPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PullPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PullPayment *PullPaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PullPayment.Contract.PullPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PullPayment *PullPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullPayment.Contract.PullPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PullPayment *PullPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PullPayment.Contract.PullPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PullPayment *PullPaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PullPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PullPayment *PullPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PullPayment *PullPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PullPayment.Contract.contract.Transact(opts, method, params...)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_PullPayment *PullPaymentCaller) Payments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PullPayment.contract.Call(opts, out, "payments", arg0)
	return *ret0, err
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_PullPayment *PullPaymentSession) Payments(arg0 common.Address) (*big.Int, error) {
	return _PullPayment.Contract.Payments(&_PullPayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments( address) constant returns(uint256)
func (_PullPayment *PullPaymentCallerSession) Payments(arg0 common.Address) (*big.Int, error) {
	return _PullPayment.Contract.Payments(&_PullPayment.CallOpts, arg0)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_PullPayment *PullPaymentCaller) TotalPayments(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PullPayment.contract.Call(opts, out, "totalPayments")
	return *ret0, err
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_PullPayment *PullPaymentSession) TotalPayments() (*big.Int, error) {
	return _PullPayment.Contract.TotalPayments(&_PullPayment.CallOpts)
}

// TotalPayments is a free data retrieval call binding the contract method 0x005b4487.
//
// Solidity: function totalPayments() constant returns(uint256)
func (_PullPayment *PullPaymentCallerSession) TotalPayments() (*big.Int, error) {
	return _PullPayment.Contract.TotalPayments(&_PullPayment.CallOpts)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_PullPayment *PullPaymentTransactor) WithdrawPayments(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PullPayment.contract.Transact(opts, "withdrawPayments")
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_PullPayment *PullPaymentSession) WithdrawPayments() (*types.Transaction, error) {
	return _PullPayment.Contract.WithdrawPayments(&_PullPayment.TransactOpts)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x6103d70b.
//
// Solidity: function withdrawPayments() returns()
func (_PullPayment *PullPaymentTransactorSession) WithdrawPayments() (*types.Transaction, error) {
	return _PullPayment.Contract.WithdrawPayments(&_PullPayment.TransactOpts)
}
