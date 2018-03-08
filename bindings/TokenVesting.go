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

// TokenVestingABI is the input ABI used to generate the binding from.
const TokenVestingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cliff\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"releasableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"vestedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"revocable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"revoked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_start\",\"type\":\"uint256\"},{\"name\":\"_cliff\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_revocable\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenVestingBin is the compiled bytecode used for deploying new contracts.
const TokenVestingBin = `6060604052341561000f57600080fd5b60405160a0806108c8833981016040528080519190602001805191906020018051919060200180519190602001805160008054600160a060020a03191633600160a060020a039081169190911790915590925086161515905061007157600080fd5b8183111561007e57600080fd5b60018054600160a060020a031916600160a060020a0387161790556005805460ff191682151517905560048290556100c384846401000000006100d281026106c41704565b600255505050600355506100e8565b6000828201838110156100e157fe5b9392505050565b6107d1806100f76000396000f3006060604052600436106100ab5763ffffffff60e060020a6000350416630fb5a6b481146100b057806313d033c0146100d55780631726cbc8146100e85780631916558714610107578063384711cc1461012857806338af3eed1461014757806374a8f10314610176578063872a7810146101955780638da5cb5b146101bc5780639852595c146101cf578063be9a6555146101ee578063f2fde38b14610201578063fa01dc0614610220575b600080fd5b34156100bb57600080fd5b6100c361023f565b60405190815260200160405180910390f35b34156100e057600080fd5b6100c3610245565b34156100f357600080fd5b6100c3600160a060020a036004351661024b565b341561011257600080fd5b610126600160a060020a0360043516610283565b005b341561013357600080fd5b6100c3600160a060020a036004351661032f565b341561015257600080fd5b61015a610470565b604051600160a060020a03909116815260200160405180910390f35b341561018157600080fd5b610126600160a060020a036004351661047f565b34156101a057600080fd5b6101a86105d2565b604051901515815260200160405180910390f35b34156101c757600080fd5b61015a6105db565b34156101da57600080fd5b6100c3600160a060020a03600435166105ea565b34156101f957600080fd5b6100c36105fc565b341561020c57600080fd5b610126600160a060020a0360043516610602565b341561022b57600080fd5b6101a8600160a060020a036004351661069d565b60045481565b60025481565b600160a060020a03811660009081526006602052604081205461027d906102718461032f565b9063ffffffff6106b216565b92915050565b600061028e8261024b565b90506000811161029d57600080fd5b600160a060020a0382166000908152600660205260409020546102c6908263ffffffff6106c416565b600160a060020a038084166000818152600660205260409020929092556001546102f89291168363ffffffff6106de16565b7ffb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c5658160405190815260200160405180910390a15050565b600080600083600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561038b57600080fd5b6102c65a03f1151561039c57600080fd5b5050506040518051600160a060020a0386166000908152600660205260409020549093506103d29150839063ffffffff6106c416565b90506002544210156103e75760009250610469565b6004546003546103fc9163ffffffff6106c416565b421015806104225750600160a060020a03841660009081526007602052604090205460ff165b1561042f57809250610469565b61046660045461045a61044d600354426106b290919063ffffffff16565b849063ffffffff61076316565b9063ffffffff61078e16565b92505b5050919050565b600154600160a060020a031681565b600080548190819033600160a060020a0390811691161461049f57600080fd5b60055460ff1615156104b057600080fd5b600160a060020a03841660009081526007602052604090205460ff16156104d657600080fd5b83600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561052d57600080fd5b6102c65a03f1151561053e57600080fd5b5050506040518051905092506105538461024b565b9150610565838363ffffffff6106b216565b600160a060020a038086166000818152600760205260408120805460ff19166001179055549293506105a0929091168363ffffffff6106de16565b7f44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee660405160405180910390a150505050565b60055460ff1681565b600054600160a060020a031681565b60066020526000908152604090205481565b60035481565b60005433600160a060020a0390811691161461061d57600080fd5b600160a060020a038116151561063257600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60076020526000908152604090205460ff1681565b6000828211156106be57fe5b50900390565b6000828201838110156106d357fe5b8091505b5092915050565b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561073b57600080fd5b6102c65a03f1151561074c57600080fd5b50505060405180519050151561075e57fe5b505050565b60008083151561077657600091506106d7565b5082820282848281151561078657fe5b04146106d357fe5b600080828481151561079c57fe5b049493505050505600a165627a7a7230582036eb00e0b0408c1a0499bdb61de165695f05cad3d82aacf23c52a73c16f5276d0029`

// DeployTokenVesting deploys a new Ethereum contract, binding an instance of TokenVesting to it.
func DeployTokenVesting(auth *bind.TransactOpts, backend bind.ContractBackend, _beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _revocable bool) (common.Address, *types.Transaction, *TokenVesting, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenVestingBin), backend, _beneficiary, _start, _cliff, _duration, _revocable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// TokenVesting is an auto generated Go binding around an Ethereum contract.
type TokenVesting struct {
	TokenVestingCaller     // Read-only binding to the contract
	TokenVestingTransactor // Write-only binding to the contract
	TokenVestingFilterer   // Log filterer for contract events
}

// TokenVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenVestingSession struct {
	Contract     *TokenVesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenVestingCallerSession struct {
	Contract *TokenVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenVestingTransactorSession struct {
	Contract     *TokenVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenVestingRaw struct {
	Contract *TokenVesting // Generic contract binding to access the raw methods on
}

// TokenVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenVestingCallerRaw struct {
	Contract *TokenVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokenVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenVestingTransactorRaw struct {
	Contract *TokenVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenVesting creates a new instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVesting(address common.Address, backend bind.ContractBackend) (*TokenVesting, error) {
	contract, err := bindTokenVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// NewTokenVestingCaller creates a new read-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingCaller(address common.Address, caller bind.ContractCaller) (*TokenVestingCaller, error) {
	contract, err := bindTokenVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingCaller{contract: contract}, nil
}

// NewTokenVestingTransactor creates a new write-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenVestingTransactor, error) {
	contract, err := bindTokenVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingTransactor{contract: contract}, nil
}

// NewTokenVestingFilterer creates a new log filterer instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenVestingFilterer, error) {
	contract, err := bindTokenVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenVestingFilterer{contract: contract}, nil
}

// bindTokenVesting binds a generic wrapper to an already deployed contract.
func bindTokenVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.TokenVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingSession) Beneficiary() (common.Address, error) {
	return _TokenVesting.Contract.Beneficiary(&_TokenVesting.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Beneficiary() (common.Address, error) {
	return _TokenVesting.Contract.Beneficiary(&_TokenVesting.CallOpts)
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Cliff(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "cliff")
	return *ret0, err
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Cliff() (*big.Int, error) {
	return _TokenVesting.Contract.Cliff(&_TokenVesting.CallOpts)
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Cliff() (*big.Int, error) {
	return _TokenVesting.Contract.Cliff(&_TokenVesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Duration() (*big.Int, error) {
	return _TokenVesting.Contract.Duration(&_TokenVesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Duration() (*big.Int, error) {
	return _TokenVesting.Contract.Duration(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// ReleasableAmount is a free data retrieval call binding the contract method 0x1726cbc8.
//
// Solidity: function releasableAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) ReleasableAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "releasableAmount", token)
	return *ret0, err
}

// ReleasableAmount is a free data retrieval call binding the contract method 0x1726cbc8.
//
// Solidity: function releasableAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) ReleasableAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleasableAmount(&_TokenVesting.CallOpts, token)
}

// ReleasableAmount is a free data retrieval call binding the contract method 0x1726cbc8.
//
// Solidity: function releasableAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) ReleasableAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleasableAmount(&_TokenVesting.CallOpts, token)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Released(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "released", arg0)
	return *ret0, err
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Released(arg0 common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Released(&_TokenVesting.CallOpts, arg0)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Released(arg0 common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Released(&_TokenVesting.CallOpts, arg0)
}

// Revocable is a free data retrieval call binding the contract method 0x872a7810.
//
// Solidity: function revocable() constant returns(bool)
func (_TokenVesting *TokenVestingCaller) Revocable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "revocable")
	return *ret0, err
}

// Revocable is a free data retrieval call binding the contract method 0x872a7810.
//
// Solidity: function revocable() constant returns(bool)
func (_TokenVesting *TokenVestingSession) Revocable() (bool, error) {
	return _TokenVesting.Contract.Revocable(&_TokenVesting.CallOpts)
}

// Revocable is a free data retrieval call binding the contract method 0x872a7810.
//
// Solidity: function revocable() constant returns(bool)
func (_TokenVesting *TokenVestingCallerSession) Revocable() (bool, error) {
	return _TokenVesting.Contract.Revocable(&_TokenVesting.CallOpts)
}

// Revoked is a free data retrieval call binding the contract method 0xfa01dc06.
//
// Solidity: function revoked( address) constant returns(bool)
func (_TokenVesting *TokenVestingCaller) Revoked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "revoked", arg0)
	return *ret0, err
}

// Revoked is a free data retrieval call binding the contract method 0xfa01dc06.
//
// Solidity: function revoked( address) constant returns(bool)
func (_TokenVesting *TokenVestingSession) Revoked(arg0 common.Address) (bool, error) {
	return _TokenVesting.Contract.Revoked(&_TokenVesting.CallOpts, arg0)
}

// Revoked is a free data retrieval call binding the contract method 0xfa01dc06.
//
// Solidity: function revoked( address) constant returns(bool)
func (_TokenVesting *TokenVestingCallerSession) Revoked(arg0 common.Address) (bool, error) {
	return _TokenVesting.Contract.Revoked(&_TokenVesting.CallOpts, arg0)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "start")
	return *ret0, err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Start() (*big.Int, error) {
	return _TokenVesting.Contract.Start(&_TokenVesting.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Start() (*big.Int, error) {
	return _TokenVesting.Contract.Start(&_TokenVesting.CallOpts)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) VestedAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "vestedAmount", token)
	return *ret0, err
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) VestedAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, token)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) VestedAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(token address) returns()
func (_TokenVesting *TokenVestingTransactor) Release(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "release", token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(token address) returns()
func (_TokenVesting *TokenVestingSession) Release(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(token address) returns()
func (_TokenVesting *TokenVestingTransactorSession) Release(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, token)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(token address) returns()
func (_TokenVesting *TokenVestingTransactor) Revoke(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "revoke", token)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(token address) returns()
func (_TokenVesting *TokenVestingSession) Revoke(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Revoke(&_TokenVesting.TransactOpts, token)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(token address) returns()
func (_TokenVesting *TokenVestingTransactorSession) Revoke(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Revoke(&_TokenVesting.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TokenVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenVesting contract.
type TokenVestingOwnershipTransferredIterator struct {
	Event *TokenVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingOwnershipTransferred)
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
		it.Event = new(TokenVestingOwnershipTransferred)
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
func (it *TokenVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TokenVesting contract.
type TokenVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingOwnershipTransferredIterator{contract: _TokenVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingOwnershipTransferred)
				if err := _TokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenVestingReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the TokenVesting contract.
type TokenVestingReleasedIterator struct {
	Event *TokenVestingReleased // Event containing the contract specifics and raw log

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
func (it *TokenVestingReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingReleased)
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
		it.Event = new(TokenVestingReleased)
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
func (it *TokenVestingReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingReleased represents a Released event raised by the TokenVesting contract.
type TokenVestingReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(amount uint256)
func (_TokenVesting *TokenVestingFilterer) FilterReleased(opts *bind.FilterOpts) (*TokenVestingReleasedIterator, error) {

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return &TokenVestingReleasedIterator{contract: _TokenVesting.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(amount uint256)
func (_TokenVesting *TokenVestingFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *TokenVestingReleased) (event.Subscription, error) {

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingReleased)
				if err := _TokenVesting.contract.UnpackLog(event, "Released", log); err != nil {
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

// TokenVestingRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the TokenVesting contract.
type TokenVestingRevokedIterator struct {
	Event *TokenVestingRevoked // Event containing the contract specifics and raw log

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
func (it *TokenVestingRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingRevoked)
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
		it.Event = new(TokenVestingRevoked)
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
func (it *TokenVestingRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingRevoked represents a Revoked event raised by the TokenVesting contract.
type TokenVestingRevoked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_TokenVesting *TokenVestingFilterer) FilterRevoked(opts *bind.FilterOpts) (*TokenVestingRevokedIterator, error) {

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return &TokenVestingRevokedIterator{contract: _TokenVesting.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_TokenVesting *TokenVestingFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *TokenVestingRevoked) (event.Subscription, error) {

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingRevoked)
				if err := _TokenVesting.contract.UnpackLog(event, "Revoked", log); err != nil {
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
