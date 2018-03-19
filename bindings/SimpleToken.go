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

// SimpleTokenABI is the input ABI used to generate the binding from.
const SimpleTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// SimpleTokenBin is the compiled bytecode used for deploying new contracts.
const SimpleTokenBin = `6060604052341561000f57600080fd5b69021e19e0c9bab24000006001819055600160a060020a033316600081815260208190526040808220849055919290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91905190815260200160405180910390a3610881806100806000396000f3006060604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b31461014857806318160ddd1461017e57806323b872dd146101a35780632ff2e9dc146101cb578063313ce567146101de578063661884631461020757806370a082311461022957806395d89b4114610248578063a9059cbb1461025b578063d73dd6231461027d578063dd62ed3e1461029f575b600080fd5b34156100c957600080fd5b6100d16102c4565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561010d5780820151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015357600080fd5b61016a600160a060020a03600435166024356102fb565b604051901515815260200160405180910390f35b341561018957600080fd5b610191610367565b60405190815260200160405180910390f35b34156101ae57600080fd5b61016a600160a060020a036004358116906024351660443561036d565b34156101d657600080fd5b6101916104ed565b34156101e957600080fd5b6101f16104fb565b60405160ff909116815260200160405180910390f35b341561021257600080fd5b61016a600160a060020a0360043516602435610500565b341561023457600080fd5b610191600160a060020a03600435166105fa565b341561025357600080fd5b6100d1610615565b341561026657600080fd5b61016a600160a060020a036004351660243561064c565b341561028857600080fd5b61016a600160a060020a036004351660243561075e565b34156102aa57600080fd5b610191600160a060020a0360043581169060243516610802565b60408051908101604052600b81527f53696d706c65546f6b656e000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b6000600160a060020a038316151561038457600080fd5b600160a060020a0384166000908152602081905260409020548211156103a957600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156103dc57600080fd5b600160a060020a038416600090815260208190526040902054610405908363ffffffff61082d16565b600160a060020a03808616600090815260208190526040808220939093559085168152205461043a908363ffffffff61083f16565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610480908363ffffffff61082d16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b69021e19e0c9bab240000081565b601281565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561055d57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610594565b61056d818463ffffffff61082d16565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b60408051908101604052600381527f53494d0000000000000000000000000000000000000000000000000000000000602082015281565b6000600160a060020a038316151561066357600080fd5b600160a060020a03331660009081526020819052604090205482111561068857600080fd5b600160a060020a0333166000908152602081905260409020546106b1908363ffffffff61082d16565b600160a060020a0333811660009081526020819052604080822093909355908516815220546106e6908363ffffffff61083f16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610796908363ffffffff61083f16565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561083957fe5b50900390565b60008282018381101561084e57fe5b93925050505600a165627a7a7230582062b1c7f955b4fc0ac84a3d02c8d5d24c76fb3b5bb5d88a8d2432bc46632814ca0029`

// DeploySimpleToken deploys a new Ethereum contract, binding an instance of SimpleToken to it.
func DeploySimpleToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleToken{SimpleTokenCaller: SimpleTokenCaller{contract: contract}, SimpleTokenTransactor: SimpleTokenTransactor{contract: contract}, SimpleTokenFilterer: SimpleTokenFilterer{contract: contract}}, nil
}

// SimpleToken is an auto generated Go binding around an Ethereum contract.
type SimpleToken struct {
	SimpleTokenCaller     // Read-only binding to the contract
	SimpleTokenTransactor // Write-only binding to the contract
	SimpleTokenFilterer   // Log filterer for contract events
}

// SimpleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleTokenSession struct {
	Contract     *SimpleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleTokenCallerSession struct {
	Contract *SimpleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTokenTransactorSession struct {
	Contract     *SimpleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleTokenRaw struct {
	Contract *SimpleToken // Generic contract binding to access the raw methods on
}

// SimpleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleTokenCallerRaw struct {
	Contract *SimpleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTokenTransactorRaw struct {
	Contract *SimpleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleToken creates a new instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleToken(address common.Address, backend bind.ContractBackend) (*SimpleToken, error) {
	contract, err := bindSimpleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleToken{SimpleTokenCaller: SimpleTokenCaller{contract: contract}, SimpleTokenTransactor: SimpleTokenTransactor{contract: contract}, SimpleTokenFilterer: SimpleTokenFilterer{contract: contract}}, nil
}

// NewSimpleTokenCaller creates a new read-only instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenCaller(address common.Address, caller bind.ContractCaller) (*SimpleTokenCaller, error) {
	contract, err := bindSimpleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenCaller{contract: contract}, nil
}

// NewSimpleTokenTransactor creates a new write-only instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTokenTransactor, error) {
	contract, err := bindSimpleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenTransactor{contract: contract}, nil
}

// NewSimpleTokenFilterer creates a new log filterer instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleTokenFilterer, error) {
	contract, err := bindSimpleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenFilterer{contract: contract}, nil
}

// bindSimpleToken binds a generic wrapper to an already deployed contract.
func bindSimpleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleToken *SimpleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleToken.Contract.SimpleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleToken *SimpleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleToken.Contract.SimpleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleToken *SimpleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleToken.Contract.SimpleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleToken *SimpleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleToken *SimpleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleToken *SimpleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleToken.Contract.contract.Transact(opts, method, params...)
}

// INITIAL_SUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) INITIAL_SUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIAL_SUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) INITIAL_SUPPLY() (*big.Int, error) {
	return _SimpleToken.Contract.INITIAL_SUPPLY(&_SimpleToken.CallOpts)
}

// INITIAL_SUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) INITIAL_SUPPLY() (*big.Int, error) {
	return _SimpleToken.Contract.INITIAL_SUPPLY(&_SimpleToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.Allowance(&_SimpleToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.Allowance(&_SimpleToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SimpleToken *SimpleTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SimpleToken *SimpleTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.BalanceOf(&_SimpleToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SimpleToken *SimpleTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.BalanceOf(&_SimpleToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SimpleToken *SimpleTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SimpleToken *SimpleTokenSession) Decimals() (uint8, error) {
	return _SimpleToken.Contract.Decimals(&_SimpleToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SimpleToken *SimpleTokenCallerSession) Decimals() (uint8, error) {
	return _SimpleToken.Contract.Decimals(&_SimpleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SimpleToken *SimpleTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SimpleToken *SimpleTokenSession) Name() (string, error) {
	return _SimpleToken.Contract.Name(&_SimpleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SimpleToken *SimpleTokenCallerSession) Name() (string, error) {
	return _SimpleToken.Contract.Name(&_SimpleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SimpleToken *SimpleTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SimpleToken *SimpleTokenSession) Symbol() (string, error) {
	return _SimpleToken.Contract.Symbol(&_SimpleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SimpleToken *SimpleTokenCallerSession) Symbol() (string, error) {
	return _SimpleToken.Contract.Symbol(&_SimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) TotalSupply() (*big.Int, error) {
	return _SimpleToken.Contract.TotalSupply(&_SimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SimpleToken.Contract.TotalSupply(&_SimpleToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.Approve(&_SimpleToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.Approve(&_SimpleToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SimpleToken *SimpleTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.DecreaseApproval(&_SimpleToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.DecreaseApproval(&_SimpleToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SimpleToken *SimpleTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.IncreaseApproval(&_SimpleToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.IncreaseApproval(&_SimpleToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.Transfer(&_SimpleToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.Transfer(&_SimpleToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.TransferFrom(&_SimpleToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SimpleToken *SimpleTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.TransferFrom(&_SimpleToken.TransactOpts, _from, _to, _value)
}

// SimpleTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SimpleToken contract.
type SimpleTokenApprovalIterator struct {
	Event *SimpleTokenApproval // Event containing the contract specifics and raw log

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
func (it *SimpleTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenApproval)
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
		it.Event = new(SimpleTokenApproval)
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
func (it *SimpleTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenApproval represents a Approval event raised by the SimpleToken contract.
type SimpleTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SimpleToken *SimpleTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SimpleTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenApprovalIterator{contract: _SimpleToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SimpleToken *SimpleTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SimpleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenApproval)
				if err := _SimpleToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SimpleTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SimpleToken contract.
type SimpleTokenTransferIterator struct {
	Event *SimpleTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SimpleTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenTransfer)
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
		it.Event = new(SimpleTokenTransfer)
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
func (it *SimpleTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenTransfer represents a Transfer event raised by the SimpleToken contract.
type SimpleTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SimpleToken *SimpleTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenTransferIterator{contract: _SimpleToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SimpleToken *SimpleTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenTransfer)
				if err := _SimpleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
