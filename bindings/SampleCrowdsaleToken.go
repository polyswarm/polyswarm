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

// SampleCrowdsaleTokenABI is the input ABI used to generate the binding from.
const SampleCrowdsaleTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// SampleCrowdsaleTokenBin is the compiled bytecode used for deploying new contracts.
const SampleCrowdsaleTokenBin = `606060405260038054600160a860020a03191633600160a060020a0316179055610bbb8061002e6000396000f3006060604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100ea57806306fdde0314610111578063095ea7b31461019b57806318160ddd146101bd57806323b872dd146101e2578063313ce5671461020a57806340c10f1914610233578063661884631461025557806370a08231146102775780637d64bcb4146102965780638da5cb5b146102a957806395d89b41146102d8578063a9059cbb146102eb578063d73dd6231461030d578063dd62ed3e1461032f578063f2fde38b14610354575b600080fd5b34156100f557600080fd5b6100fd610375565b604051901515815260200160405180910390f35b341561011c57600080fd5b610124610396565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610160578082015183820152602001610148565b50505050905090810190601f16801561018d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101a657600080fd5b6100fd600160a060020a03600435166024356103cd565b34156101c857600080fd5b6101d0610439565b60405190815260200160405180910390f35b34156101ed57600080fd5b6100fd600160a060020a036004358116906024351660443561043f565b341561021557600080fd5b61021d6105bf565b60405160ff909116815260200160405180910390f35b341561023e57600080fd5b6100fd600160a060020a03600435166024356105c4565b341561026057600080fd5b6100fd600160a060020a03600435166024356106e3565b341561028257600080fd5b6101d0600160a060020a03600435166107dd565b34156102a157600080fd5b6100fd6107f8565b34156102b457600080fd5b6102bc6108a5565b604051600160a060020a03909116815260200160405180910390f35b34156102e357600080fd5b6101246108b4565b34156102f657600080fd5b6100fd600160a060020a03600435166024356108eb565b341561031857600080fd5b6100fd600160a060020a03600435166024356109fd565b341561033a57600080fd5b6101d0600160a060020a0360043581169060243516610aa1565b341561035f57600080fd5b610373600160a060020a0360043516610acc565b005b60035474010000000000000000000000000000000000000000900460ff1681565b60408051908101604052601681527f53616d706c652043726f776473616c6520546f6b656e00000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b6000600160a060020a038316151561045657600080fd5b600160a060020a03841660009081526020819052604090205482111561047b57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156104ae57600080fd5b600160a060020a0384166000908152602081905260409020546104d7908363ffffffff610b6716565b600160a060020a03808616600090815260208190526040808220939093559085168152205461050c908363ffffffff610b7916565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610552908363ffffffff610b6716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b601281565b60035460009033600160a060020a039081169116146105e257600080fd5b60035474010000000000000000000000000000000000000000900460ff161561060a57600080fd5b60015461061d908363ffffffff610b7916565b600155600160a060020a038316600090815260208190526040902054610649908363ffffffff610b7916565b600160a060020a0384166000818152602081905260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561074057600160a060020a033381166000908152600260209081526040808320938816835292905290812055610777565b610750818463ffffffff610b6716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b60035460009033600160a060020a0390811691161461081657600080fd5b60035474010000000000000000000000000000000000000000900460ff161561083e57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60408051908101604052600381527f5343540000000000000000000000000000000000000000000000000000000000602082015281565b6000600160a060020a038316151561090257600080fd5b600160a060020a03331660009081526020819052604090205482111561092757600080fd5b600160a060020a033316600090815260208190526040902054610950908363ffffffff610b6716565b600160a060020a033381166000908152602081905260408082209390935590851681522054610985908363ffffffff610b7916565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610a35908363ffffffff610b7916565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a03908116911614610ae757600080fd5b600160a060020a0381161515610afc57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610b7357fe5b50900390565b600082820183811015610b8857fe5b93925050505600a165627a7a7230582005732c882f8fe2900e2ec1adcbb1f0581ff004ace8bb11f1b89bbe1211910a0e0029`

// DeploySampleCrowdsaleToken deploys a new Ethereum contract, binding an instance of SampleCrowdsaleToken to it.
func DeploySampleCrowdsaleToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SampleCrowdsaleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCrowdsaleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleCrowdsaleTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleCrowdsaleToken{SampleCrowdsaleTokenCaller: SampleCrowdsaleTokenCaller{contract: contract}, SampleCrowdsaleTokenTransactor: SampleCrowdsaleTokenTransactor{contract: contract}, SampleCrowdsaleTokenFilterer: SampleCrowdsaleTokenFilterer{contract: contract}}, nil
}

// SampleCrowdsaleToken is an auto generated Go binding around an Ethereum contract.
type SampleCrowdsaleToken struct {
	SampleCrowdsaleTokenCaller     // Read-only binding to the contract
	SampleCrowdsaleTokenTransactor // Write-only binding to the contract
	SampleCrowdsaleTokenFilterer   // Log filterer for contract events
}

// SampleCrowdsaleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleCrowdsaleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleCrowdsaleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleCrowdsaleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleCrowdsaleTokenSession struct {
	Contract     *SampleCrowdsaleToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SampleCrowdsaleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleCrowdsaleTokenCallerSession struct {
	Contract *SampleCrowdsaleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SampleCrowdsaleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleCrowdsaleTokenTransactorSession struct {
	Contract     *SampleCrowdsaleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SampleCrowdsaleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleCrowdsaleTokenRaw struct {
	Contract *SampleCrowdsaleToken // Generic contract binding to access the raw methods on
}

// SampleCrowdsaleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleCrowdsaleTokenCallerRaw struct {
	Contract *SampleCrowdsaleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SampleCrowdsaleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleCrowdsaleTokenTransactorRaw struct {
	Contract *SampleCrowdsaleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleCrowdsaleToken creates a new instance of SampleCrowdsaleToken, bound to a specific deployed contract.
func NewSampleCrowdsaleToken(address common.Address, backend bind.ContractBackend) (*SampleCrowdsaleToken, error) {
	contract, err := bindSampleCrowdsaleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleToken{SampleCrowdsaleTokenCaller: SampleCrowdsaleTokenCaller{contract: contract}, SampleCrowdsaleTokenTransactor: SampleCrowdsaleTokenTransactor{contract: contract}, SampleCrowdsaleTokenFilterer: SampleCrowdsaleTokenFilterer{contract: contract}}, nil
}

// NewSampleCrowdsaleTokenCaller creates a new read-only instance of SampleCrowdsaleToken, bound to a specific deployed contract.
func NewSampleCrowdsaleTokenCaller(address common.Address, caller bind.ContractCaller) (*SampleCrowdsaleTokenCaller, error) {
	contract, err := bindSampleCrowdsaleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenCaller{contract: contract}, nil
}

// NewSampleCrowdsaleTokenTransactor creates a new write-only instance of SampleCrowdsaleToken, bound to a specific deployed contract.
func NewSampleCrowdsaleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleCrowdsaleTokenTransactor, error) {
	contract, err := bindSampleCrowdsaleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenTransactor{contract: contract}, nil
}

// NewSampleCrowdsaleTokenFilterer creates a new log filterer instance of SampleCrowdsaleToken, bound to a specific deployed contract.
func NewSampleCrowdsaleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleCrowdsaleTokenFilterer, error) {
	contract, err := bindSampleCrowdsaleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenFilterer{contract: contract}, nil
}

// bindSampleCrowdsaleToken binds a generic wrapper to an already deployed contract.
func bindSampleCrowdsaleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCrowdsaleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleCrowdsaleToken.Contract.SampleCrowdsaleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.SampleCrowdsaleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.SampleCrowdsaleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleCrowdsaleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SampleCrowdsaleToken.Contract.Allowance(&_SampleCrowdsaleToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SampleCrowdsaleToken.Contract.Allowance(&_SampleCrowdsaleToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SampleCrowdsaleToken.Contract.BalanceOf(&_SampleCrowdsaleToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SampleCrowdsaleToken.Contract.BalanceOf(&_SampleCrowdsaleToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Decimals() (uint8, error) {
	return _SampleCrowdsaleToken.Contract.Decimals(&_SampleCrowdsaleToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) Decimals() (uint8, error) {
	return _SampleCrowdsaleToken.Contract.Decimals(&_SampleCrowdsaleToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) MintingFinished() (bool, error) {
	return _SampleCrowdsaleToken.Contract.MintingFinished(&_SampleCrowdsaleToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) MintingFinished() (bool, error) {
	return _SampleCrowdsaleToken.Contract.MintingFinished(&_SampleCrowdsaleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Name() (string, error) {
	return _SampleCrowdsaleToken.Contract.Name(&_SampleCrowdsaleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) Name() (string, error) {
	return _SampleCrowdsaleToken.Contract.Name(&_SampleCrowdsaleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Owner() (common.Address, error) {
	return _SampleCrowdsaleToken.Contract.Owner(&_SampleCrowdsaleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) Owner() (common.Address, error) {
	return _SampleCrowdsaleToken.Contract.Owner(&_SampleCrowdsaleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Symbol() (string, error) {
	return _SampleCrowdsaleToken.Contract.Symbol(&_SampleCrowdsaleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) Symbol() (string, error) {
	return _SampleCrowdsaleToken.Contract.Symbol(&_SampleCrowdsaleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsaleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) TotalSupply() (*big.Int, error) {
	return _SampleCrowdsaleToken.Contract.TotalSupply(&_SampleCrowdsaleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SampleCrowdsaleToken.Contract.TotalSupply(&_SampleCrowdsaleToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.Approve(&_SampleCrowdsaleToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.Approve(&_SampleCrowdsaleToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.DecreaseApproval(&_SampleCrowdsaleToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.DecreaseApproval(&_SampleCrowdsaleToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) FinishMinting() (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.FinishMinting(&_SampleCrowdsaleToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.FinishMinting(&_SampleCrowdsaleToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.IncreaseApproval(&_SampleCrowdsaleToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.IncreaseApproval(&_SampleCrowdsaleToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.Mint(&_SampleCrowdsaleToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.Mint(&_SampleCrowdsaleToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.Transfer(&_SampleCrowdsaleToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.Transfer(&_SampleCrowdsaleToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.TransferFrom(&_SampleCrowdsaleToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.TransferFrom(&_SampleCrowdsaleToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.TransferOwnership(&_SampleCrowdsaleToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsaleToken.Contract.TransferOwnership(&_SampleCrowdsaleToken.TransactOpts, newOwner)
}

// SampleCrowdsaleTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenApprovalIterator struct {
	Event *SampleCrowdsaleTokenApproval // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleTokenApproval)
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
		it.Event = new(SampleCrowdsaleTokenApproval)
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
func (it *SampleCrowdsaleTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleTokenApproval represents a Approval event raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SampleCrowdsaleTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenApprovalIterator{contract: _SampleCrowdsaleToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleTokenApproval)
				if err := _SampleCrowdsaleToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SampleCrowdsaleTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenMintIterator struct {
	Event *SampleCrowdsaleTokenMint // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleTokenMint)
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
		it.Event = new(SampleCrowdsaleTokenMint)
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
func (it *SampleCrowdsaleTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleTokenMint represents a Mint event raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*SampleCrowdsaleTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenMintIterator{contract: _SampleCrowdsaleToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleTokenMint)
				if err := _SampleCrowdsaleToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// SampleCrowdsaleTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenMintFinishedIterator struct {
	Event *SampleCrowdsaleTokenMintFinished // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleTokenMintFinished)
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
		it.Event = new(SampleCrowdsaleTokenMintFinished)
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
func (it *SampleCrowdsaleTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleTokenMintFinished represents a MintFinished event raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*SampleCrowdsaleTokenMintFinishedIterator, error) {

	logs, sub, err := _SampleCrowdsaleToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenMintFinishedIterator{contract: _SampleCrowdsaleToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _SampleCrowdsaleToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleTokenMintFinished)
				if err := _SampleCrowdsaleToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// SampleCrowdsaleTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenOwnershipTransferredIterator struct {
	Event *SampleCrowdsaleTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleTokenOwnershipTransferred)
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
		it.Event = new(SampleCrowdsaleTokenOwnershipTransferred)
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
func (it *SampleCrowdsaleTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleTokenOwnershipTransferred represents a OwnershipTransferred event raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SampleCrowdsaleTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenOwnershipTransferredIterator{contract: _SampleCrowdsaleToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleTokenOwnershipTransferred)
				if err := _SampleCrowdsaleToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SampleCrowdsaleTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenTransferIterator struct {
	Event *SampleCrowdsaleTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleTokenTransfer)
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
		it.Event = new(SampleCrowdsaleTokenTransfer)
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
func (it *SampleCrowdsaleTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleTokenTransfer represents a Transfer event raised by the SampleCrowdsaleToken contract.
type SampleCrowdsaleTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SampleCrowdsaleTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenTransferIterator{contract: _SampleCrowdsaleToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SampleCrowdsaleToken *SampleCrowdsaleTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SampleCrowdsaleToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleTokenTransfer)
				if err := _SampleCrowdsaleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
