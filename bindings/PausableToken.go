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

// PausableTokenABI is the input ABI used to generate the binding from.
const PausableTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PausableTokenBin is the compiled bytecode used for deploying new contracts.
const PausableTokenBin = `606060405260038054600160a860020a03191633600160a060020a0316179055610a358061002e6000396000f3006060604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100c957806318160ddd146100ff57806323b872dd146101245780633f4ba83a1461014c5780635c975abb14610161578063661884631461017457806370a08231146101965780638456cb59146101b55780638da5cb5b146101c8578063a9059cbb146101f7578063d73dd62314610219578063dd62ed3e1461023b578063f2fde38b14610260575b600080fd5b34156100d457600080fd5b6100eb600160a060020a036004351660243561027f565b604051901515815260200160405180910390f35b341561010a57600080fd5b6101126102aa565b60405190815260200160405180910390f35b341561012f57600080fd5b6100eb600160a060020a03600435811690602435166044356102b0565b341561015757600080fd5b61015f6102dd565b005b341561016c57600080fd5b6100eb61035c565b341561017f57600080fd5b6100eb600160a060020a036004351660243561036c565b34156101a157600080fd5b610112600160a060020a0360043516610390565b34156101c057600080fd5b61015f6103ab565b34156101d357600080fd5b6101db61042f565b604051600160a060020a03909116815260200160405180910390f35b341561020257600080fd5b6100eb600160a060020a036004351660243561043e565b341561022457600080fd5b6100eb600160a060020a0360043516602435610462565b341561024657600080fd5b610112600160a060020a0360043581169060243516610486565b341561026b57600080fd5b61015f600160a060020a03600435166104b1565b60035460009060a060020a900460ff161561029957600080fd5b6102a3838361054c565b9392505050565b60015490565b60035460009060a060020a900460ff16156102ca57600080fd5b6102d58484846105b8565b949350505050565b60035433600160a060020a039081169116146102f857600080fd5b60035460a060020a900460ff16151561031057600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b60035460009060a060020a900460ff161561038657600080fd5b6102a38383610738565b600160a060020a031660009081526020819052604090205490565b60035433600160a060020a039081169116146103c657600080fd5b60035460a060020a900460ff16156103dd57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60035460009060a060020a900460ff161561045857600080fd5b6102a38383610832565b60035460009060a060020a900460ff161561047c57600080fd5b6102a38383610944565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146104cc57600080fd5b600160a060020a03811615156104e157600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b6000600160a060020a03831615156105cf57600080fd5b600160a060020a0384166000908152602081905260409020548211156105f457600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561062757600080fd5b600160a060020a038416600090815260208190526040902054610650908363ffffffff6109e816565b600160a060020a038086166000908152602081905260408082209390935590851681522054610685908363ffffffff6109fa16565b600160a060020a03808516600090815260208181526040808320949094558783168252600281528382203390931682529190915220546106cb908363ffffffff6109e816565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561079557600160a060020a0333811660009081526002602090815260408083209388168352929052908120556107cc565b6107a5818463ffffffff6109e816565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b6000600160a060020a038316151561084957600080fd5b600160a060020a03331660009081526020819052604090205482111561086e57600080fd5b600160a060020a033316600090815260208190526040902054610897908363ffffffff6109e816565b600160a060020a0333811660009081526020819052604080822093909355908516815220546108cc908363ffffffff6109fa16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a03338116600090815260026020908152604080832093861683529290529081205461097c908363ffffffff6109fa16565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b6000828211156109f457fe5b50900390565b6000828201838110156102a357fe00a165627a7a723058205bf6fd8159883905fe8df94e7158496ee84242a3c5972bf9791ec6461bc012be0029`

// DeployPausableToken deploys a new Ethereum contract, binding an instance of PausableToken to it.
func DeployPausableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// PausableToken is an auto generated Go binding around an Ethereum contract.
type PausableToken struct {
	PausableTokenCaller     // Read-only binding to the contract
	PausableTokenTransactor // Write-only binding to the contract
	PausableTokenFilterer   // Log filterer for contract events
}

// PausableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableTokenSession struct {
	Contract     *PausableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableTokenCallerSession struct {
	Contract *PausableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PausableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTokenTransactorSession struct {
	Contract     *PausableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PausableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableTokenRaw struct {
	Contract *PausableToken // Generic contract binding to access the raw methods on
}

// PausableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableTokenCallerRaw struct {
	Contract *PausableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTokenTransactorRaw struct {
	Contract *PausableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableToken creates a new instance of PausableToken, bound to a specific deployed contract.
func NewPausableToken(address common.Address, backend bind.ContractBackend) (*PausableToken, error) {
	contract, err := bindPausableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// NewPausableTokenCaller creates a new read-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenCaller(address common.Address, caller bind.ContractCaller) (*PausableTokenCaller, error) {
	contract, err := bindPausableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenCaller{contract: contract}, nil
}

// NewPausableTokenTransactor creates a new write-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTokenTransactor, error) {
	contract, err := bindPausableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransactor{contract: contract}, nil
}

// NewPausableTokenFilterer creates a new log filterer instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableTokenFilterer, error) {
	contract, err := bindPausableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableTokenFilterer{contract: contract}, nil
}

// bindPausableToken binds a generic wrapper to an already deployed contract.
func bindPausableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.PausableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableToken *PausableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableToken *PausableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableToken *PausableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCallerSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableToken *PausableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableToken *PausableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableToken *PausableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// PausableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PausableToken contract.
type PausableTokenApprovalIterator struct {
	Event *PausableTokenApproval // Event containing the contract specifics and raw log

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
func (it *PausableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenApproval)
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
		it.Event = new(PausableTokenApproval)
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
func (it *PausableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenApproval represents a Approval event raised by the PausableToken contract.
type PausableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PausableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenApprovalIterator{contract: _PausableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PausableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenApproval)
				if err := _PausableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PausableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PausableToken contract.
type PausableTokenOwnershipTransferredIterator struct {
	Event *PausableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenOwnershipTransferred)
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
		it.Event = new(PausableTokenOwnershipTransferred)
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
func (it *PausableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PausableToken contract.
type PausableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenOwnershipTransferredIterator{contract: _PausableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenOwnershipTransferred)
				if err := _PausableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PausableToken contract.
type PausableTokenPauseIterator struct {
	Event *PausableTokenPause // Event containing the contract specifics and raw log

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
func (it *PausableTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenPause)
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
		it.Event = new(PausableTokenPause)
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
func (it *PausableTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenPause represents a Pause event raised by the PausableToken contract.
type PausableTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PausableToken *PausableTokenFilterer) FilterPause(opts *bind.FilterOpts) (*PausableTokenPauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenPauseIterator{contract: _PausableToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PausableToken *PausableTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausableTokenPause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenPause)
				if err := _PausableToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PausableToken contract.
type PausableTokenTransferIterator struct {
	Event *PausableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PausableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenTransfer)
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
		it.Event = new(PausableTokenTransfer)
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
func (it *PausableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenTransfer represents a Transfer event raised by the PausableToken contract.
type PausableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PausableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransferIterator{contract: _PausableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PausableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenTransfer)
				if err := _PausableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// PausableTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PausableToken contract.
type PausableTokenUnpauseIterator struct {
	Event *PausableTokenUnpause // Event containing the contract specifics and raw log

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
func (it *PausableTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenUnpause)
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
		it.Event = new(PausableTokenUnpause)
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
func (it *PausableTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenUnpause represents a Unpause event raised by the PausableToken contract.
type PausableTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PausableToken *PausableTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableTokenUnpauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenUnpauseIterator{contract: _PausableToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PausableToken *PausableTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenUnpause)
				if err := _PausableToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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
