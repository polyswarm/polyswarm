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

// PausableTokenMockABI is the input ABI used to generate the binding from.
const PausableTokenMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PausableTokenMockBin is the compiled bytecode used for deploying new contracts.
const PausableTokenMockBin = `60606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051604080610ab7833981016040528080519190602001805160038054600160a060020a03338116600160a060020a031990921691909117909155939093166000908152602081905260409020929092555050610a35806100826000396000f3006060604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100c957806318160ddd146100ff57806323b872dd146101245780633f4ba83a1461014c5780635c975abb14610161578063661884631461017457806370a08231146101965780638456cb59146101b55780638da5cb5b146101c8578063a9059cbb146101f7578063d73dd62314610219578063dd62ed3e1461023b578063f2fde38b14610260575b600080fd5b34156100d457600080fd5b6100eb600160a060020a036004351660243561027f565b604051901515815260200160405180910390f35b341561010a57600080fd5b6101126102aa565b60405190815260200160405180910390f35b341561012f57600080fd5b6100eb600160a060020a03600435811690602435166044356102b0565b341561015757600080fd5b61015f6102dd565b005b341561016c57600080fd5b6100eb61035c565b341561017f57600080fd5b6100eb600160a060020a036004351660243561036c565b34156101a157600080fd5b610112600160a060020a0360043516610390565b34156101c057600080fd5b61015f6103ab565b34156101d357600080fd5b6101db61042f565b604051600160a060020a03909116815260200160405180910390f35b341561020257600080fd5b6100eb600160a060020a036004351660243561043e565b341561022457600080fd5b6100eb600160a060020a0360043516602435610462565b341561024657600080fd5b610112600160a060020a0360043581169060243516610486565b341561026b57600080fd5b61015f600160a060020a03600435166104b1565b60035460009060a060020a900460ff161561029957600080fd5b6102a3838361054c565b9392505050565b60015490565b60035460009060a060020a900460ff16156102ca57600080fd5b6102d58484846105b8565b949350505050565b60035433600160a060020a039081169116146102f857600080fd5b60035460a060020a900460ff16151561031057600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b60035460009060a060020a900460ff161561038657600080fd5b6102a38383610738565b600160a060020a031660009081526020819052604090205490565b60035433600160a060020a039081169116146103c657600080fd5b60035460a060020a900460ff16156103dd57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60035460009060a060020a900460ff161561045857600080fd5b6102a38383610832565b60035460009060a060020a900460ff161561047c57600080fd5b6102a38383610944565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146104cc57600080fd5b600160a060020a03811615156104e157600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b6000600160a060020a03831615156105cf57600080fd5b600160a060020a0384166000908152602081905260409020548211156105f457600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561062757600080fd5b600160a060020a038416600090815260208190526040902054610650908363ffffffff6109e816565b600160a060020a038086166000908152602081905260408082209390935590851681522054610685908363ffffffff6109fa16565b600160a060020a03808516600090815260208181526040808320949094558783168252600281528382203390931682529190915220546106cb908363ffffffff6109e816565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561079557600160a060020a0333811660009081526002602090815260408083209388168352929052908120556107cc565b6107a5818463ffffffff6109e816565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b6000600160a060020a038316151561084957600080fd5b600160a060020a03331660009081526020819052604090205482111561086e57600080fd5b600160a060020a033316600090815260208190526040902054610897908363ffffffff6109e816565b600160a060020a0333811660009081526020819052604080822093909355908516815220546108cc908363ffffffff6109fa16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a03338116600090815260026020908152604080832093861683529290529081205461097c908363ffffffff6109fa16565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b6000828211156109f457fe5b50900390565b6000828201838110156102a357fe00a165627a7a72305820c9002c2d37ce437fd361ae32bf90b7c2b9dd41d001ce1134b106885ea119d7500029`

// DeployPausableTokenMock deploys a new Ethereum contract, binding an instance of PausableTokenMock to it.
func DeployPausableTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *PausableTokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableTokenMockBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableTokenMock{PausableTokenMockCaller: PausableTokenMockCaller{contract: contract}, PausableTokenMockTransactor: PausableTokenMockTransactor{contract: contract}}, nil
}

// PausableTokenMock is an auto generated Go binding around an Ethereum contract.
type PausableTokenMock struct {
	PausableTokenMockCaller     // Read-only binding to the contract
	PausableTokenMockTransactor // Write-only binding to the contract
}

// PausableTokenMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableTokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableTokenMockSession struct {
	Contract     *PausableTokenMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PausableTokenMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableTokenMockCallerSession struct {
	Contract *PausableTokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PausableTokenMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTokenMockTransactorSession struct {
	Contract     *PausableTokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PausableTokenMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableTokenMockRaw struct {
	Contract *PausableTokenMock // Generic contract binding to access the raw methods on
}

// PausableTokenMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableTokenMockCallerRaw struct {
	Contract *PausableTokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTokenMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTokenMockTransactorRaw struct {
	Contract *PausableTokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableTokenMock creates a new instance of PausableTokenMock, bound to a specific deployed contract.
func NewPausableTokenMock(address common.Address, backend bind.ContractBackend) (*PausableTokenMock, error) {
	contract, err := bindPausableTokenMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableTokenMock{PausableTokenMockCaller: PausableTokenMockCaller{contract: contract}, PausableTokenMockTransactor: PausableTokenMockTransactor{contract: contract}}, nil
}

// NewPausableTokenMockCaller creates a new read-only instance of PausableTokenMock, bound to a specific deployed contract.
func NewPausableTokenMockCaller(address common.Address, caller bind.ContractCaller) (*PausableTokenMockCaller, error) {
	contract, err := bindPausableTokenMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenMockCaller{contract: contract}, nil
}

// NewPausableTokenMockTransactor creates a new write-only instance of PausableTokenMock, bound to a specific deployed contract.
func NewPausableTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTokenMockTransactor, error) {
	contract, err := bindPausableTokenMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PausableTokenMockTransactor{contract: contract}, nil
}

// bindPausableTokenMock binds a generic wrapper to an already deployed contract.
func bindPausableTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableTokenMock *PausableTokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableTokenMock.Contract.PausableTokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableTokenMock *PausableTokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.PausableTokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableTokenMock *PausableTokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.PausableTokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableTokenMock *PausableTokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableTokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableTokenMock *PausableTokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableTokenMock *PausableTokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableTokenMock *PausableTokenMockCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableTokenMock.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableTokenMock *PausableTokenMockSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableTokenMock.Contract.Allowance(&_PausableTokenMock.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableTokenMock *PausableTokenMockCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableTokenMock.Contract.Allowance(&_PausableTokenMock.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableTokenMock *PausableTokenMockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableTokenMock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableTokenMock *PausableTokenMockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableTokenMock.Contract.BalanceOf(&_PausableTokenMock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableTokenMock *PausableTokenMockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableTokenMock.Contract.BalanceOf(&_PausableTokenMock.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableTokenMock *PausableTokenMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableTokenMock.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableTokenMock *PausableTokenMockSession) Owner() (common.Address, error) {
	return _PausableTokenMock.Contract.Owner(&_PausableTokenMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableTokenMock *PausableTokenMockCallerSession) Owner() (common.Address, error) {
	return _PausableTokenMock.Contract.Owner(&_PausableTokenMock.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableTokenMock *PausableTokenMockCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableTokenMock.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableTokenMock *PausableTokenMockSession) Paused() (bool, error) {
	return _PausableTokenMock.Contract.Paused(&_PausableTokenMock.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableTokenMock *PausableTokenMockCallerSession) Paused() (bool, error) {
	return _PausableTokenMock.Contract.Paused(&_PausableTokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableTokenMock *PausableTokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableTokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableTokenMock *PausableTokenMockSession) TotalSupply() (*big.Int, error) {
	return _PausableTokenMock.Contract.TotalSupply(&_PausableTokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableTokenMock *PausableTokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _PausableTokenMock.Contract.TotalSupply(&_PausableTokenMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Approve(&_PausableTokenMock.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Approve(&_PausableTokenMock.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableTokenMock *PausableTokenMockTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableTokenMock *PausableTokenMockSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.DecreaseApproval(&_PausableTokenMock.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableTokenMock *PausableTokenMockTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.DecreaseApproval(&_PausableTokenMock.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableTokenMock *PausableTokenMockTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableTokenMock *PausableTokenMockSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.IncreaseApproval(&_PausableTokenMock.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableTokenMock *PausableTokenMockTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.IncreaseApproval(&_PausableTokenMock.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableTokenMock *PausableTokenMockTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableTokenMock *PausableTokenMockSession) Pause() (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Pause(&_PausableTokenMock.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableTokenMock *PausableTokenMockTransactorSession) Pause() (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Pause(&_PausableTokenMock.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Transfer(&_PausableTokenMock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Transfer(&_PausableTokenMock.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.TransferFrom(&_PausableTokenMock.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableTokenMock *PausableTokenMockTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.TransferFrom(&_PausableTokenMock.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableTokenMock *PausableTokenMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableTokenMock *PausableTokenMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.TransferOwnership(&_PausableTokenMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableTokenMock *PausableTokenMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableTokenMock.Contract.TransferOwnership(&_PausableTokenMock.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableTokenMock *PausableTokenMockTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableTokenMock.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableTokenMock *PausableTokenMockSession) Unpause() (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Unpause(&_PausableTokenMock.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableTokenMock *PausableTokenMockTransactorSession) Unpause() (*types.Transaction, error) {
	return _PausableTokenMock.Contract.Unpause(&_PausableTokenMock.TransactOpts)
}
