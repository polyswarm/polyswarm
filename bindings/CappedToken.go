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

// CappedTokenABI is the input ABI used to generate the binding from.
const CappedTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// CappedTokenBin is the compiled bytecode used for deploying new contracts.
const CappedTokenBin = `60606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610b108339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600455610aa18061006f6000396000f3006060604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100d4578063095ea7b3146100fb57806318160ddd1461011d57806323b872dd14610142578063355274ea1461016a57806340c10f191461017d578063661884631461019f57806370a08231146101c15780637d64bcb4146101e05780638da5cb5b146101f3578063a9059cbb14610222578063d73dd62314610244578063dd62ed3e14610266578063f2fde38b1461028b575b600080fd5b34156100df57600080fd5b6100e76102ac565b604051901515815260200160405180910390f35b341561010657600080fd5b6100e7600160a060020a03600435166024356102bc565b341561012857600080fd5b610130610328565b60405190815260200160405180910390f35b341561014d57600080fd5b6100e7600160a060020a036004358116906024351660443561032e565b341561017557600080fd5b6101306104ae565b341561018857600080fd5b6100e7600160a060020a03600435166024356104b4565b34156101aa57600080fd5b6100e7600160a060020a036004351660243561051b565b34156101cc57600080fd5b610130600160a060020a0360043516610615565b34156101eb57600080fd5b6100e7610630565b34156101fe57600080fd5b6102066106bb565b604051600160a060020a03909116815260200160405180910390f35b341561022d57600080fd5b6100e7600160a060020a03600435166024356106ca565b341561024f57600080fd5b6100e7600160a060020a03600435166024356107dc565b341561027157600080fd5b610130600160a060020a0360043581169060243516610880565b341561029657600080fd5b6102aa600160a060020a03600435166108ab565b005b60035460a060020a900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b6000600160a060020a038316151561034557600080fd5b600160a060020a03841660009081526020819052604090205482111561036a57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561039d57600080fd5b600160a060020a0384166000908152602081905260409020546103c6908363ffffffff61094616565b600160a060020a0380861660009081526020819052604080822093909355908516815220546103fb908363ffffffff61095816565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610441908363ffffffff61094616565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60045481565b60035460009033600160a060020a039081169116146104d257600080fd5b60035460a060020a900460ff16156104e957600080fd5b6004546001546104ff908463ffffffff61095816565b111561050a57600080fd5b6105148383610967565b9392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561057857600160a060020a0333811660009081526002602090815260408083209388168352929052908120556105af565b610588818463ffffffff61094616565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b60035460009033600160a060020a0390811691161461064e57600080fd5b60035460a060020a900460ff161561066557600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156106e157600080fd5b600160a060020a03331660009081526020819052604090205482111561070657600080fd5b600160a060020a03331660009081526020819052604090205461072f908363ffffffff61094616565b600160a060020a033381166000908152602081905260408082209390935590851681522054610764908363ffffffff61095816565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610814908363ffffffff61095816565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146108c657600080fd5b600160a060020a03811615156108db57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561095257fe5b50900390565b60008282018381101561051457fe5b60035460009033600160a060020a0390811691161461098557600080fd5b60035460a060020a900460ff161561099c57600080fd5b6001546109af908363ffffffff61095816565b600155600160a060020a0383166000908152602081905260409020546109db908363ffffffff61095816565b600160a060020a0384166000818152602081905260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a3506001929150505600a165627a7a72305820e729f7726c5c0c5497d5f828cbcae34f18d2c43eb90eb76966e7e7e07c23369e0029`

// DeployCappedToken deploys a new Ethereum contract, binding an instance of CappedToken to it.
func DeployCappedToken(auth *bind.TransactOpts, backend bind.ContractBackend, _cap *big.Int) (common.Address, *types.Transaction, *CappedToken, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CappedTokenBin), backend, _cap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CappedToken{CappedTokenCaller: CappedTokenCaller{contract: contract}, CappedTokenTransactor: CappedTokenTransactor{contract: contract}}, nil
}

// CappedToken is an auto generated Go binding around an Ethereum contract.
type CappedToken struct {
	CappedTokenCaller     // Read-only binding to the contract
	CappedTokenTransactor // Write-only binding to the contract
}

// CappedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedTokenSession struct {
	Contract     *CappedToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CappedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedTokenCallerSession struct {
	Contract *CappedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CappedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedTokenTransactorSession struct {
	Contract     *CappedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CappedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedTokenRaw struct {
	Contract *CappedToken // Generic contract binding to access the raw methods on
}

// CappedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedTokenCallerRaw struct {
	Contract *CappedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CappedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedTokenTransactorRaw struct {
	Contract *CappedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedToken creates a new instance of CappedToken, bound to a specific deployed contract.
func NewCappedToken(address common.Address, backend bind.ContractBackend) (*CappedToken, error) {
	contract, err := bindCappedToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedToken{CappedTokenCaller: CappedTokenCaller{contract: contract}, CappedTokenTransactor: CappedTokenTransactor{contract: contract}}, nil
}

// NewCappedTokenCaller creates a new read-only instance of CappedToken, bound to a specific deployed contract.
func NewCappedTokenCaller(address common.Address, caller bind.ContractCaller) (*CappedTokenCaller, error) {
	contract, err := bindCappedToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CappedTokenCaller{contract: contract}, nil
}

// NewCappedTokenTransactor creates a new write-only instance of CappedToken, bound to a specific deployed contract.
func NewCappedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedTokenTransactor, error) {
	contract, err := bindCappedToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CappedTokenTransactor{contract: contract}, nil
}

// bindCappedToken binds a generic wrapper to an already deployed contract.
func bindCappedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedToken *CappedTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedToken.Contract.CappedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedToken *CappedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedToken.Contract.CappedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedToken *CappedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedToken.Contract.CappedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedToken *CappedTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedToken *CappedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedToken *CappedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_CappedToken *CappedTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_CappedToken *CappedTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CappedToken.Contract.Allowance(&_CappedToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_CappedToken *CappedTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CappedToken.Contract.Allowance(&_CappedToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_CappedToken *CappedTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_CappedToken *CappedTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CappedToken.Contract.BalanceOf(&_CappedToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_CappedToken *CappedTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CappedToken.Contract.BalanceOf(&_CappedToken.CallOpts, _owner)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedToken *CappedTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedToken.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedToken *CappedTokenSession) Cap() (*big.Int, error) {
	return _CappedToken.Contract.Cap(&_CappedToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedToken *CappedTokenCallerSession) Cap() (*big.Int, error) {
	return _CappedToken.Contract.Cap(&_CappedToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_CappedToken *CappedTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CappedToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_CappedToken *CappedTokenSession) MintingFinished() (bool, error) {
	return _CappedToken.Contract.MintingFinished(&_CappedToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_CappedToken *CappedTokenCallerSession) MintingFinished() (bool, error) {
	return _CappedToken.Contract.MintingFinished(&_CappedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CappedToken *CappedTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CappedToken *CappedTokenSession) Owner() (common.Address, error) {
	return _CappedToken.Contract.Owner(&_CappedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CappedToken *CappedTokenCallerSession) Owner() (common.Address, error) {
	return _CappedToken.Contract.Owner(&_CappedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CappedToken *CappedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CappedToken *CappedTokenSession) TotalSupply() (*big.Int, error) {
	return _CappedToken.Contract.TotalSupply(&_CappedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CappedToken *CappedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CappedToken.Contract.TotalSupply(&_CappedToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.Approve(&_CappedToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.Approve(&_CappedToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_CappedToken *CappedTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_CappedToken *CappedTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.DecreaseApproval(&_CappedToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_CappedToken *CappedTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.DecreaseApproval(&_CappedToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_CappedToken *CappedTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_CappedToken *CappedTokenSession) FinishMinting() (*types.Transaction, error) {
	return _CappedToken.Contract.FinishMinting(&_CappedToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_CappedToken *CappedTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _CappedToken.Contract.FinishMinting(&_CappedToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_CappedToken *CappedTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_CappedToken *CappedTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.IncreaseApproval(&_CappedToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_CappedToken *CappedTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.IncreaseApproval(&_CappedToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_CappedToken *CappedTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_CappedToken *CappedTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.Mint(&_CappedToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_CappedToken *CappedTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.Mint(&_CappedToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.Transfer(&_CappedToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.Transfer(&_CappedToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.TransferFrom(&_CappedToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_CappedToken *CappedTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CappedToken.Contract.TransferFrom(&_CappedToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CappedToken *CappedTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CappedToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CappedToken *CappedTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedToken.Contract.TransferOwnership(&_CappedToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CappedToken *CappedTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedToken.Contract.TransferOwnership(&_CappedToken.TransactOpts, newOwner)
}
