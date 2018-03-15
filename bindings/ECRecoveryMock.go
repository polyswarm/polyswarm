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

// ECRecoveryMockABI is the input ABI used to generate the binding from.
const ECRecoveryMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addrRecovered\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ECRecoveryMockBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryMockBin = `6060604052341561000f57600080fd5b6102588061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166319045a25811461005057806375f2566e146100cf575b600080fd5b341561005b57600080fd5b6100a6600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506100e295505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100da57600080fd5b6100a6610210565b600073__zeppelin-solidity/contracts/ECRecove__6319045a2584846040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101838152604060248301908152909160440183818151815260200191508051906020019080838360005b83811015610170578082015183820152602001610158565b50505050905090810190601f16801561019d5780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b15156101ba57600080fd5b5af415156101c757600080fd5b50505060405180516000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790555092915050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820bfa35037b89843c9a710206fa232686c347b2b706dfefafbadba78e7e2c53aec0029`

// DeployECRecoveryMock deploys a new Ethereum contract, binding an instance of ECRecoveryMock to it.
func DeployECRecoveryMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecoveryMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecoveryMock{ECRecoveryMockCaller: ECRecoveryMockCaller{contract: contract}, ECRecoveryMockTransactor: ECRecoveryMockTransactor{contract: contract}, ECRecoveryMockFilterer: ECRecoveryMockFilterer{contract: contract}}, nil
}

// ECRecoveryMock is an auto generated Go binding around an Ethereum contract.
type ECRecoveryMock struct {
	ECRecoveryMockCaller     // Read-only binding to the contract
	ECRecoveryMockTransactor // Write-only binding to the contract
	ECRecoveryMockFilterer   // Log filterer for contract events
}

// ECRecoveryMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoveryMockSession struct {
	Contract     *ECRecoveryMock   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryMockCallerSession struct {
	Contract *ECRecoveryMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ECRecoveryMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryMockTransactorSession struct {
	Contract     *ECRecoveryMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ECRecoveryMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryMockRaw struct {
	Contract *ECRecoveryMock // Generic contract binding to access the raw methods on
}

// ECRecoveryMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryMockCallerRaw struct {
	Contract *ECRecoveryMockCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryMockTransactorRaw struct {
	Contract *ECRecoveryMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecoveryMock creates a new instance of ECRecoveryMock, bound to a specific deployed contract.
func NewECRecoveryMock(address common.Address, backend bind.ContractBackend) (*ECRecoveryMock, error) {
	contract, err := bindECRecoveryMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryMock{ECRecoveryMockCaller: ECRecoveryMockCaller{contract: contract}, ECRecoveryMockTransactor: ECRecoveryMockTransactor{contract: contract}, ECRecoveryMockFilterer: ECRecoveryMockFilterer{contract: contract}}, nil
}

// NewECRecoveryMockCaller creates a new read-only instance of ECRecoveryMock, bound to a specific deployed contract.
func NewECRecoveryMockCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryMockCaller, error) {
	contract, err := bindECRecoveryMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryMockCaller{contract: contract}, nil
}

// NewECRecoveryMockTransactor creates a new write-only instance of ECRecoveryMock, bound to a specific deployed contract.
func NewECRecoveryMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryMockTransactor, error) {
	contract, err := bindECRecoveryMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryMockTransactor{contract: contract}, nil
}

// NewECRecoveryMockFilterer creates a new log filterer instance of ECRecoveryMock, bound to a specific deployed contract.
func NewECRecoveryMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryMockFilterer, error) {
	contract, err := bindECRecoveryMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryMockFilterer{contract: contract}, nil
}

// bindECRecoveryMock binds a generic wrapper to an already deployed contract.
func bindECRecoveryMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecoveryMock *ECRecoveryMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecoveryMock.Contract.ECRecoveryMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecoveryMock *ECRecoveryMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecoveryMock.Contract.ECRecoveryMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecoveryMock *ECRecoveryMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecoveryMock.Contract.ECRecoveryMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecoveryMock *ECRecoveryMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecoveryMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecoveryMock *ECRecoveryMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecoveryMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecoveryMock *ECRecoveryMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecoveryMock.Contract.contract.Transact(opts, method, params...)
}

// AddrRecovered is a free data retrieval call binding the contract method 0x75f2566e.
//
// Solidity: function addrRecovered() constant returns(address)
func (_ECRecoveryMock *ECRecoveryMockCaller) AddrRecovered(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ECRecoveryMock.contract.Call(opts, out, "addrRecovered")
	return *ret0, err
}

// AddrRecovered is a free data retrieval call binding the contract method 0x75f2566e.
//
// Solidity: function addrRecovered() constant returns(address)
func (_ECRecoveryMock *ECRecoveryMockSession) AddrRecovered() (common.Address, error) {
	return _ECRecoveryMock.Contract.AddrRecovered(&_ECRecoveryMock.CallOpts)
}

// AddrRecovered is a free data retrieval call binding the contract method 0x75f2566e.
//
// Solidity: function addrRecovered() constant returns(address)
func (_ECRecoveryMock *ECRecoveryMockCallerSession) AddrRecovered() (common.Address, error) {
	return _ECRecoveryMock.Contract.AddrRecovered(&_ECRecoveryMock.CallOpts)
}

// Recover is a paid mutator transaction binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) returns(address)
func (_ECRecoveryMock *ECRecoveryMockTransactor) Recover(opts *bind.TransactOpts, hash [32]byte, sig []byte) (*types.Transaction, error) {
	return _ECRecoveryMock.contract.Transact(opts, "recover", hash, sig)
}

// Recover is a paid mutator transaction binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) returns(address)
func (_ECRecoveryMock *ECRecoveryMockSession) Recover(hash [32]byte, sig []byte) (*types.Transaction, error) {
	return _ECRecoveryMock.Contract.Recover(&_ECRecoveryMock.TransactOpts, hash, sig)
}

// Recover is a paid mutator transaction binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) returns(address)
func (_ECRecoveryMock *ECRecoveryMockTransactorSession) Recover(hash [32]byte, sig []byte) (*types.Transaction, error) {
	return _ECRecoveryMock.Contract.Recover(&_ECRecoveryMock.TransactOpts, hash, sig)
}
