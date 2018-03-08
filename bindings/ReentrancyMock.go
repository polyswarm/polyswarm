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

// ReentrancyMockABI is the input ABI used to generate the binding from.
const ReentrancyMockABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"countThisRecursive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"countLocalRecursive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"attacker\",\"type\":\"address\"}],\"name\":\"countAndCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyMockBin is the compiled bytecode used for deploying new contracts.
const ReentrancyMockBin = `60606040526000805460ff19169055341561001957600080fd5b60006001556103648061002d6000396000f3006060604052600436106100535763ffffffff60e060020a600035041663083b2732811461005857806361bc221a1461006d5780638c5344fa1461009257806396ffa690146100a8578063b672ad8b146100be575b600080fd5b341561006357600080fd5b61006b6100ea565b005b341561007857600080fd5b61008061011b565b60405190815260200160405180910390f35b341561009d57600080fd5b61006b600435610121565b34156100b357600080fd5b61006b6004356101fa565b34156100c957600080fd5b61006b73ffffffffffffffffffffffffffffffffffffffff6004351661023f565b60005460ff16156100fa57600080fd5b6000805460ff1916600117905561010f61032e565b6000805460ff19169055565b60015481565b60008054819060ff161561013457600080fd5b6000805460ff191660011790556040517f636f756e74546869735265637572736976652875696e743235362900000000008152601b016040518091039020915060008311156101eb5761018561032e565b73ffffffffffffffffffffffffffffffffffffffff301660e060020a8304600019850160405160e060020a63ffffffff8416028152600481019190915260240160006040518083038160008761646e5a03f193505050508015156001146101eb57600080fd5b50506000805460ff1916905550565b60005460ff161561020a57600080fd5b6000805460ff191660011781558111156102325761022661032e565b610232600182036101fa565b506000805460ff19169055565b6000805460ff161561025057600080fd5b6000805460ff1916600117905561026561032e565b6040517f63616c6c6261636b2829000000000000000000000000000000000000000000008152600a01604051809103902090508173ffffffffffffffffffffffffffffffffffffffff16630a2df1ed8260405160e060020a63ffffffff84160281527fffffffff000000000000000000000000000000000000000000000000000000009091166004820152602401600060405180830381600087803b151561030c57600080fd5b6102c65a03f1151561031d57600080fd5b50506000805460ff19169055505050565b60018054810190555600a165627a7a72305820ab5e9d4b7e37c685fbfb640dd80c211dacde23e80531dcca1de07903f69cf96b0029`

// DeployReentrancyMock deploys a new Ethereum contract, binding an instance of ReentrancyMock to it.
func DeployReentrancyMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancyMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyMock{ReentrancyMockCaller: ReentrancyMockCaller{contract: contract}, ReentrancyMockTransactor: ReentrancyMockTransactor{contract: contract}, ReentrancyMockFilterer: ReentrancyMockFilterer{contract: contract}}, nil
}

// ReentrancyMock is an auto generated Go binding around an Ethereum contract.
type ReentrancyMock struct {
	ReentrancyMockCaller     // Read-only binding to the contract
	ReentrancyMockTransactor // Write-only binding to the contract
	ReentrancyMockFilterer   // Log filterer for contract events
}

// ReentrancyMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyMockSession struct {
	Contract     *ReentrancyMock   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyMockCallerSession struct {
	Contract *ReentrancyMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReentrancyMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyMockTransactorSession struct {
	Contract     *ReentrancyMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReentrancyMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyMockRaw struct {
	Contract *ReentrancyMock // Generic contract binding to access the raw methods on
}

// ReentrancyMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyMockCallerRaw struct {
	Contract *ReentrancyMockCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyMockTransactorRaw struct {
	Contract *ReentrancyMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyMock creates a new instance of ReentrancyMock, bound to a specific deployed contract.
func NewReentrancyMock(address common.Address, backend bind.ContractBackend) (*ReentrancyMock, error) {
	contract, err := bindReentrancyMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyMock{ReentrancyMockCaller: ReentrancyMockCaller{contract: contract}, ReentrancyMockTransactor: ReentrancyMockTransactor{contract: contract}, ReentrancyMockFilterer: ReentrancyMockFilterer{contract: contract}}, nil
}

// NewReentrancyMockCaller creates a new read-only instance of ReentrancyMock, bound to a specific deployed contract.
func NewReentrancyMockCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyMockCaller, error) {
	contract, err := bindReentrancyMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyMockCaller{contract: contract}, nil
}

// NewReentrancyMockTransactor creates a new write-only instance of ReentrancyMock, bound to a specific deployed contract.
func NewReentrancyMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyMockTransactor, error) {
	contract, err := bindReentrancyMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyMockTransactor{contract: contract}, nil
}

// NewReentrancyMockFilterer creates a new log filterer instance of ReentrancyMock, bound to a specific deployed contract.
func NewReentrancyMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyMockFilterer, error) {
	contract, err := bindReentrancyMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyMockFilterer{contract: contract}, nil
}

// bindReentrancyMock binds a generic wrapper to an already deployed contract.
func bindReentrancyMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyMock *ReentrancyMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyMock.Contract.ReentrancyMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyMock *ReentrancyMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.ReentrancyMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyMock *ReentrancyMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.ReentrancyMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyMock *ReentrancyMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyMock *ReentrancyMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyMock *ReentrancyMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() constant returns(uint256)
func (_ReentrancyMock *ReentrancyMockCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReentrancyMock.contract.Call(opts, out, "counter")
	return *ret0, err
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() constant returns(uint256)
func (_ReentrancyMock *ReentrancyMockSession) Counter() (*big.Int, error) {
	return _ReentrancyMock.Contract.Counter(&_ReentrancyMock.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() constant returns(uint256)
func (_ReentrancyMock *ReentrancyMockCallerSession) Counter() (*big.Int, error) {
	return _ReentrancyMock.Contract.Counter(&_ReentrancyMock.CallOpts)
}

// Callback is a paid mutator transaction binding the contract method 0x083b2732.
//
// Solidity: function callback() returns()
func (_ReentrancyMock *ReentrancyMockTransactor) Callback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyMock.contract.Transact(opts, "callback")
}

// Callback is a paid mutator transaction binding the contract method 0x083b2732.
//
// Solidity: function callback() returns()
func (_ReentrancyMock *ReentrancyMockSession) Callback() (*types.Transaction, error) {
	return _ReentrancyMock.Contract.Callback(&_ReentrancyMock.TransactOpts)
}

// Callback is a paid mutator transaction binding the contract method 0x083b2732.
//
// Solidity: function callback() returns()
func (_ReentrancyMock *ReentrancyMockTransactorSession) Callback() (*types.Transaction, error) {
	return _ReentrancyMock.Contract.Callback(&_ReentrancyMock.TransactOpts)
}

// CountAndCall is a paid mutator transaction binding the contract method 0xb672ad8b.
//
// Solidity: function countAndCall(attacker address) returns()
func (_ReentrancyMock *ReentrancyMockTransactor) CountAndCall(opts *bind.TransactOpts, attacker common.Address) (*types.Transaction, error) {
	return _ReentrancyMock.contract.Transact(opts, "countAndCall", attacker)
}

// CountAndCall is a paid mutator transaction binding the contract method 0xb672ad8b.
//
// Solidity: function countAndCall(attacker address) returns()
func (_ReentrancyMock *ReentrancyMockSession) CountAndCall(attacker common.Address) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.CountAndCall(&_ReentrancyMock.TransactOpts, attacker)
}

// CountAndCall is a paid mutator transaction binding the contract method 0xb672ad8b.
//
// Solidity: function countAndCall(attacker address) returns()
func (_ReentrancyMock *ReentrancyMockTransactorSession) CountAndCall(attacker common.Address) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.CountAndCall(&_ReentrancyMock.TransactOpts, attacker)
}

// CountLocalRecursive is a paid mutator transaction binding the contract method 0x96ffa690.
//
// Solidity: function countLocalRecursive(n uint256) returns()
func (_ReentrancyMock *ReentrancyMockTransactor) CountLocalRecursive(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _ReentrancyMock.contract.Transact(opts, "countLocalRecursive", n)
}

// CountLocalRecursive is a paid mutator transaction binding the contract method 0x96ffa690.
//
// Solidity: function countLocalRecursive(n uint256) returns()
func (_ReentrancyMock *ReentrancyMockSession) CountLocalRecursive(n *big.Int) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.CountLocalRecursive(&_ReentrancyMock.TransactOpts, n)
}

// CountLocalRecursive is a paid mutator transaction binding the contract method 0x96ffa690.
//
// Solidity: function countLocalRecursive(n uint256) returns()
func (_ReentrancyMock *ReentrancyMockTransactorSession) CountLocalRecursive(n *big.Int) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.CountLocalRecursive(&_ReentrancyMock.TransactOpts, n)
}

// CountThisRecursive is a paid mutator transaction binding the contract method 0x8c5344fa.
//
// Solidity: function countThisRecursive(n uint256) returns()
func (_ReentrancyMock *ReentrancyMockTransactor) CountThisRecursive(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _ReentrancyMock.contract.Transact(opts, "countThisRecursive", n)
}

// CountThisRecursive is a paid mutator transaction binding the contract method 0x8c5344fa.
//
// Solidity: function countThisRecursive(n uint256) returns()
func (_ReentrancyMock *ReentrancyMockSession) CountThisRecursive(n *big.Int) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.CountThisRecursive(&_ReentrancyMock.TransactOpts, n)
}

// CountThisRecursive is a paid mutator transaction binding the contract method 0x8c5344fa.
//
// Solidity: function countThisRecursive(n uint256) returns()
func (_ReentrancyMock *ReentrancyMockTransactorSession) CountThisRecursive(n *big.Int) (*types.Transaction, error) {
	return _ReentrancyMock.Contract.CountThisRecursive(&_ReentrancyMock.TransactOpts, n)
}
