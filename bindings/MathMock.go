// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MathMockABI is the input ABI used to generate the binding from.
const MathMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint64\"},{\"name\":\"b\",\"type\":\"uint64\"}],\"name\":\"min64\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint64\"},{\"name\":\"b\",\"type\":\"uint64\"}],\"name\":\"max64\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"result\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MathMockBin is the compiled bytecode used for deploying new contracts.
const MathMockBin = `6060604052341561000f57600080fd5b61019f8061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166336b1315c811461005b5780633a4faf7f1461008357806365372147146100a9575b600080fd5b341561006657600080fd5b61008167ffffffffffffffff600435811690602435166100d9565b005b341561008e57600080fd5b61008167ffffffffffffffff60043581169060243516610109565b34156100b457600080fd5b6100bc610113565b60405167ffffffffffffffff909116815260200160405180910390f35b6100e38282610123565b6000805467ffffffffffffffff191667ffffffffffffffff929092169190911790555050565b6100e3828261014f565b60005467ffffffffffffffff1681565b60008167ffffffffffffffff168367ffffffffffffffff16106101465781610148565b825b9392505050565b60008167ffffffffffffffff168367ffffffffffffffff16101561014657816101485600a165627a7a72305820a51bbbcc79eee9f8e8a7a91b209cc8816206bb79bef1587af07237c82c0eb7330029`

// DeployMathMock deploys a new Ethereum contract, binding an instance of MathMock to it.
func DeployMathMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MathMock, error) {
	parsed, err := abi.JSON(strings.NewReader(MathMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MathMock{MathMockCaller: MathMockCaller{contract: contract}, MathMockTransactor: MathMockTransactor{contract: contract}, MathMockFilterer: MathMockFilterer{contract: contract}}, nil
}

// MathMock is an auto generated Go binding around an Ethereum contract.
type MathMock struct {
	MathMockCaller     // Read-only binding to the contract
	MathMockTransactor // Write-only binding to the contract
	MathMockFilterer   // Log filterer for contract events
}

// MathMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathMockSession struct {
	Contract     *MathMock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathMockCallerSession struct {
	Contract *MathMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MathMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathMockTransactorSession struct {
	Contract     *MathMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MathMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathMockRaw struct {
	Contract *MathMock // Generic contract binding to access the raw methods on
}

// MathMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathMockCallerRaw struct {
	Contract *MathMockCaller // Generic read-only contract binding to access the raw methods on
}

// MathMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathMockTransactorRaw struct {
	Contract *MathMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMathMock creates a new instance of MathMock, bound to a specific deployed contract.
func NewMathMock(address common.Address, backend bind.ContractBackend) (*MathMock, error) {
	contract, err := bindMathMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MathMock{MathMockCaller: MathMockCaller{contract: contract}, MathMockTransactor: MathMockTransactor{contract: contract}, MathMockFilterer: MathMockFilterer{contract: contract}}, nil
}

// NewMathMockCaller creates a new read-only instance of MathMock, bound to a specific deployed contract.
func NewMathMockCaller(address common.Address, caller bind.ContractCaller) (*MathMockCaller, error) {
	contract, err := bindMathMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathMockCaller{contract: contract}, nil
}

// NewMathMockTransactor creates a new write-only instance of MathMock, bound to a specific deployed contract.
func NewMathMockTransactor(address common.Address, transactor bind.ContractTransactor) (*MathMockTransactor, error) {
	contract, err := bindMathMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathMockTransactor{contract: contract}, nil
}

// NewMathMockFilterer creates a new log filterer instance of MathMock, bound to a specific deployed contract.
func NewMathMockFilterer(address common.Address, filterer bind.ContractFilterer) (*MathMockFilterer, error) {
	contract, err := bindMathMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathMockFilterer{contract: contract}, nil
}

// bindMathMock binds a generic wrapper to an already deployed contract.
func bindMathMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathMock *MathMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MathMock.Contract.MathMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathMock *MathMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathMock.Contract.MathMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathMock *MathMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathMock.Contract.MathMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathMock *MathMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MathMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathMock *MathMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathMock *MathMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathMock.Contract.contract.Transact(opts, method, params...)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint64)
func (_MathMock *MathMockCaller) Result(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _MathMock.contract.Call(opts, out, "result")
	return *ret0, err
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint64)
func (_MathMock *MathMockSession) Result() (uint64, error) {
	return _MathMock.Contract.Result(&_MathMock.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint64)
func (_MathMock *MathMockCallerSession) Result() (uint64, error) {
	return _MathMock.Contract.Result(&_MathMock.CallOpts)
}

// Max64 is a paid mutator transaction binding the contract method 0x3a4faf7f.
//
// Solidity: function max64(a uint64, b uint64) returns()
func (_MathMock *MathMockTransactor) Max64(opts *bind.TransactOpts, a uint64, b uint64) (*types.Transaction, error) {
	return _MathMock.contract.Transact(opts, "max64", a, b)
}

// Max64 is a paid mutator transaction binding the contract method 0x3a4faf7f.
//
// Solidity: function max64(a uint64, b uint64) returns()
func (_MathMock *MathMockSession) Max64(a uint64, b uint64) (*types.Transaction, error) {
	return _MathMock.Contract.Max64(&_MathMock.TransactOpts, a, b)
}

// Max64 is a paid mutator transaction binding the contract method 0x3a4faf7f.
//
// Solidity: function max64(a uint64, b uint64) returns()
func (_MathMock *MathMockTransactorSession) Max64(a uint64, b uint64) (*types.Transaction, error) {
	return _MathMock.Contract.Max64(&_MathMock.TransactOpts, a, b)
}

// Min64 is a paid mutator transaction binding the contract method 0x36b1315c.
//
// Solidity: function min64(a uint64, b uint64) returns()
func (_MathMock *MathMockTransactor) Min64(opts *bind.TransactOpts, a uint64, b uint64) (*types.Transaction, error) {
	return _MathMock.contract.Transact(opts, "min64", a, b)
}

// Min64 is a paid mutator transaction binding the contract method 0x36b1315c.
//
// Solidity: function min64(a uint64, b uint64) returns()
func (_MathMock *MathMockSession) Min64(a uint64, b uint64) (*types.Transaction, error) {
	return _MathMock.Contract.Min64(&_MathMock.TransactOpts, a, b)
}

// Min64 is a paid mutator transaction binding the contract method 0x36b1315c.
//
// Solidity: function min64(a uint64, b uint64) returns()
func (_MathMock *MathMockTransactorSession) Min64(a uint64, b uint64) (*types.Transaction, error) {
	return _MathMock.Contract.Min64(&_MathMock.TransactOpts, a, b)
}
