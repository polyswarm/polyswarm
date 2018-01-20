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

// MessageHelperABI is the input ABI used to generate the binding from.
const MessageHelperABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint256\"},{\"name\":\"text\",\"type\":\"string\"}],\"name\":\"showMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"b32\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"text\",\"type\":\"string\"}],\"name\":\"Show\",\"type\":\"event\"}]"

// MessageHelperBin is the compiled bytecode used for deploying new contracts.
const MessageHelperBin = `6060604052341561000f57600080fd5b6102ce8061001e6000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631b8b921d811461005b5780636e2e2e4d146100db578063a9cc471814610137575b600080fd5b341561006657600080fd5b6100c76004803573ffffffffffffffffffffffffffffffffffffffff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061014c95505050505050565b604051901515815260200160405180910390f35b34156100e657600080fd5b6100c7600480359060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506101ed95505050505050565b341561014257600080fd5b61014a610056565b005b60008273ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b83811015610190578082015183820152602001610178565b50505050905090810190601f1680156101bd5780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f1915050156101e3575060016101e7565b5060005b92915050565b60007f53f52bd5d580b08520e25370d30be415b61c6ff8e378023685ca03eec817b10f8484846040518381526020810183905260606040820181815290820183818151815260200191508051906020019080838360005b8381101561025c578082015183820152602001610244565b50505050905090810190601f1680156102895780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a150600193925050505600a165627a7a7230582079ec090b016f5fa6ccf946affb74c6a5e8aec0119ce2d0126259777a4854a6490029`

// DeployMessageHelper deploys a new Ethereum contract, binding an instance of MessageHelper to it.
func DeployMessageHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessageHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageHelper{MessageHelperCaller: MessageHelperCaller{contract: contract}, MessageHelperTransactor: MessageHelperTransactor{contract: contract}}, nil
}

// MessageHelper is an auto generated Go binding around an Ethereum contract.
type MessageHelper struct {
	MessageHelperCaller     // Read-only binding to the contract
	MessageHelperTransactor // Write-only binding to the contract
}

// MessageHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageHelperSession struct {
	Contract     *MessageHelper    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageHelperCallerSession struct {
	Contract *MessageHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessageHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageHelperTransactorSession struct {
	Contract     *MessageHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessageHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageHelperRaw struct {
	Contract *MessageHelper // Generic contract binding to access the raw methods on
}

// MessageHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageHelperCallerRaw struct {
	Contract *MessageHelperCaller // Generic read-only contract binding to access the raw methods on
}

// MessageHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageHelperTransactorRaw struct {
	Contract *MessageHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageHelper creates a new instance of MessageHelper, bound to a specific deployed contract.
func NewMessageHelper(address common.Address, backend bind.ContractBackend) (*MessageHelper, error) {
	contract, err := bindMessageHelper(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageHelper{MessageHelperCaller: MessageHelperCaller{contract: contract}, MessageHelperTransactor: MessageHelperTransactor{contract: contract}}, nil
}

// NewMessageHelperCaller creates a new read-only instance of MessageHelper, bound to a specific deployed contract.
func NewMessageHelperCaller(address common.Address, caller bind.ContractCaller) (*MessageHelperCaller, error) {
	contract, err := bindMessageHelper(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHelperCaller{contract: contract}, nil
}

// NewMessageHelperTransactor creates a new write-only instance of MessageHelper, bound to a specific deployed contract.
func NewMessageHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageHelperTransactor, error) {
	contract, err := bindMessageHelper(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MessageHelperTransactor{contract: contract}, nil
}

// bindMessageHelper binds a generic wrapper to an already deployed contract.
func bindMessageHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageHelper *MessageHelperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessageHelper.Contract.MessageHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageHelper *MessageHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHelper.Contract.MessageHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageHelper *MessageHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHelper.Contract.MessageHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageHelper *MessageHelperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessageHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageHelper *MessageHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageHelper *MessageHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHelper.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(to address, data bytes) returns(bool)
func (_MessageHelper *MessageHelperTransactor) Call(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _MessageHelper.contract.Transact(opts, "call", to, data)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(to address, data bytes) returns(bool)
func (_MessageHelper *MessageHelperSession) Call(to common.Address, data []byte) (*types.Transaction, error) {
	return _MessageHelper.Contract.Call(&_MessageHelper.TransactOpts, to, data)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(to address, data bytes) returns(bool)
func (_MessageHelper *MessageHelperTransactorSession) Call(to common.Address, data []byte) (*types.Transaction, error) {
	return _MessageHelper.Contract.Call(&_MessageHelper.TransactOpts, to, data)
}

// Fail is a paid mutator transaction binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_MessageHelper *MessageHelperTransactor) Fail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHelper.contract.Transact(opts, "fail")
}

// Fail is a paid mutator transaction binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_MessageHelper *MessageHelperSession) Fail() (*types.Transaction, error) {
	return _MessageHelper.Contract.Fail(&_MessageHelper.TransactOpts)
}

// Fail is a paid mutator transaction binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_MessageHelper *MessageHelperTransactorSession) Fail() (*types.Transaction, error) {
	return _MessageHelper.Contract.Fail(&_MessageHelper.TransactOpts)
}

// ShowMessage is a paid mutator transaction binding the contract method 0x6e2e2e4d.
//
// Solidity: function showMessage(message bytes32, number uint256, text string) returns(bool)
func (_MessageHelper *MessageHelperTransactor) ShowMessage(opts *bind.TransactOpts, message [32]byte, number *big.Int, text string) (*types.Transaction, error) {
	return _MessageHelper.contract.Transact(opts, "showMessage", message, number, text)
}

// ShowMessage is a paid mutator transaction binding the contract method 0x6e2e2e4d.
//
// Solidity: function showMessage(message bytes32, number uint256, text string) returns(bool)
func (_MessageHelper *MessageHelperSession) ShowMessage(message [32]byte, number *big.Int, text string) (*types.Transaction, error) {
	return _MessageHelper.Contract.ShowMessage(&_MessageHelper.TransactOpts, message, number, text)
}

// ShowMessage is a paid mutator transaction binding the contract method 0x6e2e2e4d.
//
// Solidity: function showMessage(message bytes32, number uint256, text string) returns(bool)
func (_MessageHelper *MessageHelperTransactorSession) ShowMessage(message [32]byte, number *big.Int, text string) (*types.Transaction, error) {
	return _MessageHelper.Contract.ShowMessage(&_MessageHelper.TransactOpts, message, number, text)
}
