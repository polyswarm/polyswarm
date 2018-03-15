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

// CappedTokenABI is the input ABI used to generate the binding from.
const CappedTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// CappedTokenBin is the compiled bytecode used for deploying new contracts.
const CappedTokenBin = `60606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610afa8339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600455610a8b8061006f6000396000f3006060604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100d4578063095ea7b3146100fb57806318160ddd1461011d57806323b872dd14610142578063355274ea1461016a57806340c10f191461017d578063661884631461019f57806370a08231146101c15780637d64bcb4146101e05780638da5cb5b146101f3578063a9059cbb14610222578063d73dd62314610244578063dd62ed3e14610266578063f2fde38b1461028b575b600080fd5b34156100df57600080fd5b6100e76102ac565b604051901515815260200160405180910390f35b341561010657600080fd5b6100e7600160a060020a03600435166024356102bc565b341561012857600080fd5b610130610328565b60405190815260200160405180910390f35b341561014d57600080fd5b6100e7600160a060020a036004358116906024351660443561032e565b341561017557600080fd5b6101306104b0565b341561018857600080fd5b6100e7600160a060020a03600435166024356104b6565b34156101aa57600080fd5b6100e7600160a060020a036004351660243561051d565b34156101cc57600080fd5b610130600160a060020a0360043516610617565b34156101eb57600080fd5b6100e7610632565b34156101fe57600080fd5b6102066106bd565b604051600160a060020a03909116815260200160405180910390f35b341561022d57600080fd5b6100e7600160a060020a03600435166024356106cc565b341561024f57600080fd5b6100e7600160a060020a03600435166024356107c7565b341561027157600080fd5b610130600160a060020a036004358116906024351661086b565b341561029657600080fd5b6102aa600160a060020a0360043516610896565b005b60035460a060020a900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561034557600080fd5b600160a060020a03841660009081526001602052604090205482111561036a57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561039d57600080fd5b600160a060020a0384166000908152600160205260409020546103c6908363ffffffff61093116565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103fb908363ffffffff61094316565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610443908363ffffffff61093116565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60045481565b60035460009033600160a060020a039081169116146104d457600080fd5b60035460a060020a900460ff16156104eb57600080fd5b600454600054610501908463ffffffff61094316565b111561050c57600080fd5b6105168383610952565b9392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561057a57600160a060020a0333811660009081526002602090815260408083209388168352929052908120556105b1565b61058a818463ffffffff61093116565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461065057600080fd5b60035460a060020a900460ff161561066757600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156106e357600080fd5b600160a060020a03331660009081526001602052604090205482111561070857600080fd5b600160a060020a033316600090815260016020526040902054610731908363ffffffff61093116565b600160a060020a033381166000908152600160205260408082209390935590851681522054610766908363ffffffff61094316565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546107ff908363ffffffff61094316565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146108b157600080fd5b600160a060020a03811615156108c657600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561093d57fe5b50900390565b60008282018381101561051657fe5b60035460009033600160a060020a0390811691161461097057600080fd5b60035460a060020a900460ff161561098757600080fd5b60005461099a908363ffffffff61094316565b6000908155600160a060020a0384168152600160205260409020546109c5908363ffffffff61094316565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a3506001929150505600a165627a7a7230582029a644b3b69c9af338a46462057e095d4a2ab99eefd2e9068e93f436707d40170029`

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
	return address, tx, &CappedToken{CappedTokenCaller: CappedTokenCaller{contract: contract}, CappedTokenTransactor: CappedTokenTransactor{contract: contract}, CappedTokenFilterer: CappedTokenFilterer{contract: contract}}, nil
}

// CappedToken is an auto generated Go binding around an Ethereum contract.
type CappedToken struct {
	CappedTokenCaller     // Read-only binding to the contract
	CappedTokenTransactor // Write-only binding to the contract
	CappedTokenFilterer   // Log filterer for contract events
}

// CappedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedTokenFilterer struct {
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
	contract, err := bindCappedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedToken{CappedTokenCaller: CappedTokenCaller{contract: contract}, CappedTokenTransactor: CappedTokenTransactor{contract: contract}, CappedTokenFilterer: CappedTokenFilterer{contract: contract}}, nil
}

// NewCappedTokenCaller creates a new read-only instance of CappedToken, bound to a specific deployed contract.
func NewCappedTokenCaller(address common.Address, caller bind.ContractCaller) (*CappedTokenCaller, error) {
	contract, err := bindCappedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedTokenCaller{contract: contract}, nil
}

// NewCappedTokenTransactor creates a new write-only instance of CappedToken, bound to a specific deployed contract.
func NewCappedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedTokenTransactor, error) {
	contract, err := bindCappedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedTokenTransactor{contract: contract}, nil
}

// NewCappedTokenFilterer creates a new log filterer instance of CappedToken, bound to a specific deployed contract.
func NewCappedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedTokenFilterer, error) {
	contract, err := bindCappedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedTokenFilterer{contract: contract}, nil
}

// bindCappedToken binds a generic wrapper to an already deployed contract.
func bindCappedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// CappedTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CappedToken contract.
type CappedTokenApprovalIterator struct {
	Event *CappedTokenApproval // Event containing the contract specifics and raw log

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
func (it *CappedTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTokenApproval)
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
		it.Event = new(CappedTokenApproval)
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
func (it *CappedTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTokenApproval represents a Approval event raised by the CappedToken contract.
type CappedTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_CappedToken *CappedTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CappedTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CappedToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CappedTokenApprovalIterator{contract: _CappedToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_CappedToken *CappedTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CappedTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CappedToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTokenApproval)
				if err := _CappedToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// CappedTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CappedToken contract.
type CappedTokenMintIterator struct {
	Event *CappedTokenMint // Event containing the contract specifics and raw log

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
func (it *CappedTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTokenMint)
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
		it.Event = new(CappedTokenMint)
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
func (it *CappedTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTokenMint represents a Mint event raised by the CappedToken contract.
type CappedTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_CappedToken *CappedTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*CappedTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CappedToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &CappedTokenMintIterator{contract: _CappedToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_CappedToken *CappedTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CappedTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CappedToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTokenMint)
				if err := _CappedToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// CappedTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the CappedToken contract.
type CappedTokenMintFinishedIterator struct {
	Event *CappedTokenMintFinished // Event containing the contract specifics and raw log

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
func (it *CappedTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTokenMintFinished)
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
		it.Event = new(CappedTokenMintFinished)
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
func (it *CappedTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTokenMintFinished represents a MintFinished event raised by the CappedToken contract.
type CappedTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_CappedToken *CappedTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*CappedTokenMintFinishedIterator, error) {

	logs, sub, err := _CappedToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &CappedTokenMintFinishedIterator{contract: _CappedToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_CappedToken *CappedTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *CappedTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _CappedToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTokenMintFinished)
				if err := _CappedToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// CappedTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CappedToken contract.
type CappedTokenOwnershipTransferredIterator struct {
	Event *CappedTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CappedTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTokenOwnershipTransferred)
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
		it.Event = new(CappedTokenOwnershipTransferred)
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
func (it *CappedTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CappedToken contract.
type CappedTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CappedToken *CappedTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CappedTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CappedTokenOwnershipTransferredIterator{contract: _CappedToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CappedToken *CappedTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CappedTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTokenOwnershipTransferred)
				if err := _CappedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// CappedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CappedToken contract.
type CappedTokenTransferIterator struct {
	Event *CappedTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CappedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTokenTransfer)
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
		it.Event = new(CappedTokenTransfer)
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
func (it *CappedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTokenTransfer represents a Transfer event raised by the CappedToken contract.
type CappedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_CappedToken *CappedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CappedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CappedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CappedTokenTransferIterator{contract: _CappedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_CappedToken *CappedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CappedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CappedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTokenTransfer)
				if err := _CappedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
