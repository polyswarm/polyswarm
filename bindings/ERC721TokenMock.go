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

// ERC721TokenMockABI is the input ABI used to generate the binding from.
const ERC721TokenMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approvedFor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"takeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC721TokenMockBin is the compiled bytecode used for deploying new contracts.
const ERC721TokenMockBin = `6060604052341561000f57600080fd5b6109e38061001e6000396000f3006060604052600436106100a35763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100a857806318160ddd146100cc5780632a6dd48f146100f157806340c10f191461012357806342966c68146101455780635a3f26721461015b5780636352211e146101cd57806370a08231146101e3578063a9059cbb14610202578063b2e6ceeb14610224575b600080fd5b34156100b357600080fd5b6100ca600160a060020a036004351660243561023a565b005b34156100d757600080fd5b6100df61032c565b60405190815260200160405180910390f35b34156100fc57600080fd5b610107600435610333565b604051600160a060020a03909116815260200160405180910390f35b341561012e57600080fd5b6100ca600160a060020a036004351660243561034e565b341561015057600080fd5b6100ca60043561035c565b341561016657600080fd5b61017a600160a060020a0360043516610368565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101b95780820151838201526020016101a1565b505050509050019250505060405180910390f35b34156101d857600080fd5b6101076004356103eb565b34156101ee57600080fd5b6100df600160a060020a0360043516610415565b341561020d57600080fd5b6100ca600160a060020a0360043516602435610430565b341561022f57600080fd5b6100ca600435610467565b60008133600160a060020a0316610250826103eb565b600160a060020a03161461026357600080fd5b61026c836103eb565b9150600160a060020a03848116908316141561028757600080fd5b61029083610333565b600160a060020a03161515806102ae5750600160a060020a03841615155b156103265760008381526002602052604090819020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591908416907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259086905190815260200160405180910390a35b50505050565b6000545b90565b600090815260026020526040902054600160a060020a031690565b610358828261048f565b5050565b610365816104f1565b50565b610370610968565b6003600083600160a060020a0316600160a060020a031681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156103df57602002820191906000526020600020905b8154815260200190600101908083116103cb575b50505050509050919050565b600081815260016020526040812054600160a060020a031680151561040f57600080fd5b92915050565b600160a060020a031660009081526003602052604090205490565b8033600160a060020a0316610444826103eb565b600160a060020a03161461045757600080fd5b610462338484610586565b505050565b610471338261064c565b151561047c57600080fd5b610365610488826103eb565b3383610586565b600160a060020a03821615156104a457600080fd5b6104ae8282610672565b81600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b8033600160a060020a0316610505826103eb565b600160a060020a03161461051857600080fd5b61052182610333565b600160a060020a031615610539576105393383610737565b61054333836107c9565b600033600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35050565b600160a060020a038216151561059b57600080fd5b6105a4816103eb565b600160a060020a03838116911614156105bc57600080fd5b82600160a060020a03166105cf826103eb565b600160a060020a0316146105e257600080fd5b6105ec8382610737565b6105f683826107c9565b6106008282610672565b81600160a060020a031683600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a3505050565b600082600160a060020a031661066183610333565b600160a060020a0316149392505050565b600081815260016020526040812054600160a060020a03161561069457600080fd5b6000828152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790556106d183610415565b600160a060020a0384166000908152600360205260409020805491925090600181016106fd838261097a565b50600091825260208083209190910184905583825260049052604081208290555461072f90600163ffffffff61094016565b600055505050565b81600160a060020a031661074a826103eb565b600160a060020a03161461075d57600080fd5b600081815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff19169055600160a060020a038416907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259084905190815260200160405180910390a35050565b600080600084600160a060020a03166107e1856103eb565b600160a060020a0316146107f457600080fd5b600084815260046020526040902054925061081f600161081387610415565b9063ffffffff61095616565b600160a060020a03861660009081526003602052604090208054919350908390811061084757fe5b6000918252602080832090910154868352600182526040808420805473ffffffffffffffffffffffffffffffffffffffff19169055600160a060020a03891684526003909252912080549192508291859081106108a057fe5b6000918252602080832090910192909255600160a060020a03871681526003909152604081208054849081106108d257fe5b6000918252602080832090910192909255600160a060020a038716815260039091526040902080549061090990600019830161097a565b5060008481526004602052604080822082905582825281208490555461093690600163ffffffff61095616565b6000555050505050565b60008282018381101561094f57fe5b9392505050565b60008282111561096257fe5b50900390565b60206040519081016040526000815290565b8154818355818115116104625760008381526020902061046291810190830161033091905b808211156109b3576000815560010161099f565b50905600a165627a7a72305820e326c9c0a5febebbac79f7d43b8ddb7a99e395b30b6911e1beb5f9ea59bc263e0029`

// DeployERC721TokenMock deploys a new Ethereum contract, binding an instance of ERC721TokenMock to it.
func DeployERC721TokenMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721TokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721TokenMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721TokenMock{ERC721TokenMockCaller: ERC721TokenMockCaller{contract: contract}, ERC721TokenMockTransactor: ERC721TokenMockTransactor{contract: contract}}, nil
}

// ERC721TokenMock is an auto generated Go binding around an Ethereum contract.
type ERC721TokenMock struct {
	ERC721TokenMockCaller     // Read-only binding to the contract
	ERC721TokenMockTransactor // Write-only binding to the contract
}

// ERC721TokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721TokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721TokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721TokenMockSession struct {
	Contract     *ERC721TokenMock  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721TokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721TokenMockCallerSession struct {
	Contract *ERC721TokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC721TokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TokenMockTransactorSession struct {
	Contract     *ERC721TokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC721TokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721TokenMockRaw struct {
	Contract *ERC721TokenMock // Generic contract binding to access the raw methods on
}

// ERC721TokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721TokenMockCallerRaw struct {
	Contract *ERC721TokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721TokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TokenMockTransactorRaw struct {
	Contract *ERC721TokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721TokenMock creates a new instance of ERC721TokenMock, bound to a specific deployed contract.
func NewERC721TokenMock(address common.Address, backend bind.ContractBackend) (*ERC721TokenMock, error) {
	contract, err := bindERC721TokenMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenMock{ERC721TokenMockCaller: ERC721TokenMockCaller{contract: contract}, ERC721TokenMockTransactor: ERC721TokenMockTransactor{contract: contract}}, nil
}

// NewERC721TokenMockCaller creates a new read-only instance of ERC721TokenMock, bound to a specific deployed contract.
func NewERC721TokenMockCaller(address common.Address, caller bind.ContractCaller) (*ERC721TokenMockCaller, error) {
	contract, err := bindERC721TokenMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenMockCaller{contract: contract}, nil
}

// NewERC721TokenMockTransactor creates a new write-only instance of ERC721TokenMock, bound to a specific deployed contract.
func NewERC721TokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721TokenMockTransactor, error) {
	contract, err := bindERC721TokenMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenMockTransactor{contract: contract}, nil
}

// bindERC721TokenMock binds a generic wrapper to an already deployed contract.
func bindERC721TokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721TokenMock *ERC721TokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721TokenMock.Contract.ERC721TokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721TokenMock *ERC721TokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.ERC721TokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721TokenMock *ERC721TokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.ERC721TokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721TokenMock *ERC721TokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721TokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721TokenMock *ERC721TokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721TokenMock *ERC721TokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.contract.Transact(opts, method, params...)
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_ERC721TokenMock *ERC721TokenMockCaller) ApprovedFor(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721TokenMock.contract.Call(opts, out, "approvedFor", _tokenId)
	return *ret0, err
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_ERC721TokenMock *ERC721TokenMockSession) ApprovedFor(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenMock.Contract.ApprovedFor(&_ERC721TokenMock.CallOpts, _tokenId)
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_ERC721TokenMock *ERC721TokenMockCallerSession) ApprovedFor(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenMock.Contract.ApprovedFor(&_ERC721TokenMock.CallOpts, _tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721TokenMock *ERC721TokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721TokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721TokenMock *ERC721TokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721TokenMock.Contract.BalanceOf(&_ERC721TokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721TokenMock *ERC721TokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721TokenMock.Contract.BalanceOf(&_ERC721TokenMock.CallOpts, _owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721TokenMock *ERC721TokenMockCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721TokenMock.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721TokenMock *ERC721TokenMockSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenMock.Contract.OwnerOf(&_ERC721TokenMock.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721TokenMock *ERC721TokenMockCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721TokenMock.Contract.OwnerOf(&_ERC721TokenMock.CallOpts, _tokenId)
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_ERC721TokenMock *ERC721TokenMockCaller) TokensOf(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ERC721TokenMock.contract.Call(opts, out, "tokensOf", _owner)
	return *ret0, err
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_ERC721TokenMock *ERC721TokenMockSession) TokensOf(_owner common.Address) ([]*big.Int, error) {
	return _ERC721TokenMock.Contract.TokensOf(&_ERC721TokenMock.CallOpts, _owner)
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_ERC721TokenMock *ERC721TokenMockCallerSession) TokensOf(_owner common.Address) ([]*big.Int, error) {
	return _ERC721TokenMock.Contract.TokensOf(&_ERC721TokenMock.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721TokenMock *ERC721TokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721TokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721TokenMock *ERC721TokenMockSession) TotalSupply() (*big.Int, error) {
	return _ERC721TokenMock.Contract.TotalSupply(&_ERC721TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721TokenMock *ERC721TokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721TokenMock.Contract.TotalSupply(&_ERC721TokenMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Approve(&_ERC721TokenMock.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Approve(&_ERC721TokenMock.TransactOpts, _to, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Burn(&_ERC721TokenMock.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Burn(&_ERC721TokenMock.TransactOpts, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.contract.Transact(opts, "mint", _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Mint(&_ERC721TokenMock.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactorSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Mint(&_ERC721TokenMock.TransactOpts, _to, _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactor) TakeOwnership(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.contract.Transact(opts, "takeOwnership", _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockSession) TakeOwnership(_tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.TakeOwnership(&_ERC721TokenMock.TransactOpts, _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactorSession) TakeOwnership(_tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.TakeOwnership(&_ERC721TokenMock.TransactOpts, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Transfer(&_ERC721TokenMock.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_ERC721TokenMock *ERC721TokenMockTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721TokenMock.Contract.Transfer(&_ERC721TokenMock.TransactOpts, _to, _tokenId)
}
