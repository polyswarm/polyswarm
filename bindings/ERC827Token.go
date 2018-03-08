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

// ERC827TokenABI is the input ABI used to generate the binding from.
const ERC827TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC827TokenBin is the compiled bytecode used for deploying new contracts.
const ERC827TokenBin = `6060604052341561000f57600080fd5b610b228061001e6000396000f3006060604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100c957806316ca3b63146100ff57806318160ddd1461016457806323b872dd146101895780635c17f9f4146101b1578063661884631461021657806370a08231146102385780637272ad4914610257578063a9059cbb146102bc578063ab67aa58146102de578063be45fd621461034a578063d73dd623146103af578063dd62ed3e146103d1575b600080fd5b34156100d457600080fd5b6100eb600160a060020a03600435166024356103f6565b604051901515815260200160405180910390f35b341561010a57600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061046295505050505050565b341561016f57600080fd5b610177610520565b60405190815260200160405180910390f35b341561019457600080fd5b6100eb600160a060020a0360043581169060243516604435610526565b34156101bc57600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506106a895505050505050565b341561022157600080fd5b6100eb600160a060020a03600435166024356106d5565b341561024357600080fd5b610177600160a060020a03600435166107cf565b341561026257600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506107ea95505050505050565b34156102c757600080fd5b6100eb600160a060020a0360043516602435610817565b34156102e957600080fd5b6100eb600160a060020a036004803582169160248035909116916044359160849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061091295505050505050565b341561035557600080fd5b6100eb60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109d295505050505050565b34156103ba57600080fd5b6100eb600160a060020a03600435166024356109ff565b34156103dc57600080fd5b610177600160a060020a0360043581169060243516610aa3565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b600030600160a060020a031684600160a060020a03161415151561048557600080fd5b61048f84846109ff565b5083600160a060020a03168260405180828051906020019080838360005b838110156104c55780820151838201526020016104ad565b50505050905090810190601f1680156104f25780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f1915050151561051657600080fd5b5060019392505050565b60005481565b6000600160a060020a038316151561053d57600080fd5b600160a060020a03841660009081526001602052604090205482111561056257600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561059557600080fd5b600160a060020a0384166000908152600160205260409020546105be908363ffffffff610ace16565b600160a060020a0380861660009081526001602052604080822093909355908516815220546105f3908363ffffffff610ae016565b600160a060020a0380851660009081526001602090815260408083209490945587831682526002815283822033909316825291909152205461063b908363ffffffff610ace16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600030600160a060020a031684600160a060020a0316141515156106cb57600080fd5b61048f84846103f6565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561073257600160a060020a033381166000908152600260209081526040808320938816835292905290812055610769565b610742818463ffffffff610ace16565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600030600160a060020a031684600160a060020a03161415151561080d57600080fd5b61048f84846106d5565b6000600160a060020a038316151561082e57600080fd5b600160a060020a03331660009081526001602052604090205482111561085357600080fd5b600160a060020a03331660009081526001602052604090205461087c908363ffffffff610ace16565b600160a060020a0333811660009081526001602052604080822093909355908516815220546108b1908363ffffffff610ae016565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600030600160a060020a031684600160a060020a03161415151561093557600080fd5b610940858585610526565b5083600160a060020a03168260405180828051906020019080838360005b8381101561097657808201518382015260200161095e565b50505050905090810190601f1680156109a35780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f191505015156109c757600080fd5b506001949350505050565b600030600160a060020a031684600160a060020a0316141515156109f557600080fd5b61048f8484610817565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610a37908363ffffffff610ae016565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600082821115610ada57fe5b50900390565b600082820183811015610aef57fe5b93925050505600a165627a7a723058207ff363a22c6afd56cc1606ec2d3ab7b80d7b89d59dda5b7eca6194cd488a98c80029`

// DeployERC827Token deploys a new Ethereum contract, binding an instance of ERC827Token to it.
func DeployERC827Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC827Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC827TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC827TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC827Token{ERC827TokenCaller: ERC827TokenCaller{contract: contract}, ERC827TokenTransactor: ERC827TokenTransactor{contract: contract}, ERC827TokenFilterer: ERC827TokenFilterer{contract: contract}}, nil
}

// ERC827Token is an auto generated Go binding around an Ethereum contract.
type ERC827Token struct {
	ERC827TokenCaller     // Read-only binding to the contract
	ERC827TokenTransactor // Write-only binding to the contract
	ERC827TokenFilterer   // Log filterer for contract events
}

// ERC827TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC827TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC827TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC827TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC827TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC827TokenSession struct {
	Contract     *ERC827Token      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC827TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC827TokenCallerSession struct {
	Contract *ERC827TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC827TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC827TokenTransactorSession struct {
	Contract     *ERC827TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC827TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC827TokenRaw struct {
	Contract *ERC827Token // Generic contract binding to access the raw methods on
}

// ERC827TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC827TokenCallerRaw struct {
	Contract *ERC827TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC827TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC827TokenTransactorRaw struct {
	Contract *ERC827TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC827Token creates a new instance of ERC827Token, bound to a specific deployed contract.
func NewERC827Token(address common.Address, backend bind.ContractBackend) (*ERC827Token, error) {
	contract, err := bindERC827Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC827Token{ERC827TokenCaller: ERC827TokenCaller{contract: contract}, ERC827TokenTransactor: ERC827TokenTransactor{contract: contract}, ERC827TokenFilterer: ERC827TokenFilterer{contract: contract}}, nil
}

// NewERC827TokenCaller creates a new read-only instance of ERC827Token, bound to a specific deployed contract.
func NewERC827TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC827TokenCaller, error) {
	contract, err := bindERC827Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenCaller{contract: contract}, nil
}

// NewERC827TokenTransactor creates a new write-only instance of ERC827Token, bound to a specific deployed contract.
func NewERC827TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC827TokenTransactor, error) {
	contract, err := bindERC827Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenTransactor{contract: contract}, nil
}

// NewERC827TokenFilterer creates a new log filterer instance of ERC827Token, bound to a specific deployed contract.
func NewERC827TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC827TokenFilterer, error) {
	contract, err := bindERC827Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenFilterer{contract: contract}, nil
}

// bindERC827Token binds a generic wrapper to an already deployed contract.
func bindERC827Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC827TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC827Token *ERC827TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC827Token.Contract.ERC827TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC827Token *ERC827TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC827Token.Contract.ERC827TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC827Token *ERC827TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC827Token.Contract.ERC827TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC827Token *ERC827TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC827Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC827Token *ERC827TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC827Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC827Token *ERC827TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC827Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC827Token *ERC827TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC827Token *ERC827TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC827Token.Contract.Allowance(&_ERC827Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC827Token *ERC827TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC827Token.Contract.Allowance(&_ERC827Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC827Token *ERC827TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC827Token *ERC827TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC827Token.Contract.BalanceOf(&_ERC827Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC827Token *ERC827TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC827Token.Contract.BalanceOf(&_ERC827Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827Token *ERC827TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC827Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827Token *ERC827TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC827Token.Contract.TotalSupply(&_ERC827Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC827Token *ERC827TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC827Token.Contract.TotalSupply(&_ERC827Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.contract.Transact(opts, "approve", _spender, _value, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenSession) Approve(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.Approve(&_ERC827Token.TransactOpts, _spender, _value, _data)
}

// Approve is a paid mutator transaction binding the contract method 0x5c17f9f4.
//
// Solidity: function approve(_spender address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactorSession) Approve(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.Approve(&_ERC827Token.TransactOpts, _spender, _value, _data)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x7272ad49.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue, _data)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x7272ad49.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.DecreaseApproval(&_ERC827Token.TransactOpts, _spender, _subtractedValue, _data)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x7272ad49.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.DecreaseApproval(&_ERC827Token.TransactOpts, _spender, _subtractedValue, _data)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ERC827Token *ERC827TokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC827Token.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ERC827Token *ERC827TokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC827Token.Contract.IncreaseApproval(&_ERC827Token.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_ERC827Token *ERC827TokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC827Token.Contract.IncreaseApproval(&_ERC827Token.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.contract.Transact(opts, "transfer", _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.Transfer(&_ERC827Token.TransactOpts, _to, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.Transfer(&_ERC827Token.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.contract.Transact(opts, "transferFrom", _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.TransferFrom(&_ERC827Token.TransactOpts, _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xab67aa58.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_ERC827Token *ERC827TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC827Token.Contract.TransferFrom(&_ERC827Token.TransactOpts, _from, _to, _value, _data)
}

// ERC827TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC827Token contract.
type ERC827TokenApprovalIterator struct {
	Event *ERC827TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC827TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC827TokenApproval)
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
		it.Event = new(ERC827TokenApproval)
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
func (it *ERC827TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC827TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC827TokenApproval represents a Approval event raised by the ERC827Token contract.
type ERC827TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC827Token *ERC827TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC827TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC827Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenApprovalIterator{contract: _ERC827Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC827Token *ERC827TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC827TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC827Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC827TokenApproval)
				if err := _ERC827Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC827TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC827Token contract.
type ERC827TokenTransferIterator struct {
	Event *ERC827TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC827TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC827TokenTransfer)
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
		it.Event = new(ERC827TokenTransfer)
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
func (it *ERC827TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC827TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC827TokenTransfer represents a Transfer event raised by the ERC827Token contract.
type ERC827TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC827Token *ERC827TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC827TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC827Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC827TokenTransferIterator{contract: _ERC827Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC827Token *ERC827TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC827TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC827Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC827TokenTransfer)
				if err := _ERC827Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
