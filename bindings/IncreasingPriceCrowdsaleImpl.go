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

// IncreasingPriceCrowdsaleImplABI is the input ABI used to generate the binding from.
const IncreasingPriceCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_initialRate\",\"type\":\"uint256\"},{\"name\":\"_finalRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// IncreasingPriceCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const IncreasingPriceCrowdsaleImplBin = `6060604052341561000f57600080fd5b60405160c08061063283398101604052808051919060200180519190602001805191906020018051919060200180519190602001805191508290508187878388886000831161005d57600080fd5b600160a060020a038216151561007257600080fd5b600160a060020a038116151561008757600080fd5b60029290925560018054600160a060020a03928316600160a060020a03199182161790915560008054929093169116179055428210156100c657600080fd5b818110156100d357600080fd5b600491909155600555808210156100e957600080fd5b600081116100f657600080fd5b60069190915560075550505050505061051e806101146000396000f3006060604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631515bc2b81146100b957806321106109146100e05780632c4e722e146101055780634042b66f146101185780634b6753bc1461012b578063521eb2731461013e5780639e51051f1461016d578063b7a8807c14610180578063ec8ac4d814610193578063f7fb07b0146101a7578063fc0c546a146101ba575b6100b7336101cd565b005b34156100c457600080fd5b6100cc610275565b604051901515815260200160405180910390f35b34156100eb57600080fd5b6100f361027d565b60405190815260200160405180910390f35b341561011057600080fd5b6100f3610283565b341561012357600080fd5b6100f3610289565b341561013657600080fd5b6100f361028f565b341561014957600080fd5b610151610295565b604051600160a060020a03909116815260200160405180910390f35b341561017857600080fd5b6100f36102a4565b341561018b57600080fd5b6100f36102aa565b6100b7600160a060020a03600435166101cd565b34156101b257600080fd5b6100f36102b0565b34156101c557600080fd5b610151610339565b3460006101da8383610348565b6101e382610375565b6003549091506101f9908363ffffffff61039916565b60035561020683826103b3565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a361025e8383610371565b6102666103bd565b6102708383610371565b505050565b600554421190565b60075481565b60025481565b60035481565b60055481565b600154600160a060020a031681565b60065481565b60045481565b6000806000806102cb600454426103f390919063ffffffff16565b92506102e46004546005546103f390919063ffffffff16565b91506102fd6007546006546103f390919063ffffffff16565b905061033161032283610316868563ffffffff61040516565b9063ffffffff61043016565b6006549063ffffffff6103f316565b935050505090565b600054600160a060020a031681565b600454421015801561035c57506005544211155b151561036757600080fd5b6103718282610447565b5050565b6000806103806102b0565b9050610392818463ffffffff61040516565b9392505050565b6000828201838110156103a857fe5b8091505b5092915050565b6103718282610468565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f1935050505015156103f157600080fd5b565b6000828211156103ff57fe5b50900390565b60008083151561041857600091506103ac565b5082820282848281151561042857fe5b04146103a857fe5b600080828481151561043e57fe5b04949350505050565b600160a060020a038216151561045c57600080fd5b80151561037157600080fd5b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156104d757600080fd5b5af115156104e457600080fd5b5050506040518051505050505600a165627a7a72305820b5eee734f82cee6cfdc782ad55ccfddeebd133fd3d36ba97a4dd533991937db70029`

// DeployIncreasingPriceCrowdsaleImpl deploys a new Ethereum contract, binding an instance of IncreasingPriceCrowdsaleImpl to it.
func DeployIncreasingPriceCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _openingTime *big.Int, _closingTime *big.Int, _wallet common.Address, _token common.Address, _initialRate *big.Int, _finalRate *big.Int) (common.Address, *types.Transaction, *IncreasingPriceCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(IncreasingPriceCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncreasingPriceCrowdsaleImplBin), backend, _openingTime, _closingTime, _wallet, _token, _initialRate, _finalRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncreasingPriceCrowdsaleImpl{IncreasingPriceCrowdsaleImplCaller: IncreasingPriceCrowdsaleImplCaller{contract: contract}, IncreasingPriceCrowdsaleImplTransactor: IncreasingPriceCrowdsaleImplTransactor{contract: contract}, IncreasingPriceCrowdsaleImplFilterer: IncreasingPriceCrowdsaleImplFilterer{contract: contract}}, nil
}

// IncreasingPriceCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleImpl struct {
	IncreasingPriceCrowdsaleImplCaller     // Read-only binding to the contract
	IncreasingPriceCrowdsaleImplTransactor // Write-only binding to the contract
	IncreasingPriceCrowdsaleImplFilterer   // Log filterer for contract events
}

// IncreasingPriceCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncreasingPriceCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncreasingPriceCrowdsaleImplSession struct {
	Contract     *IncreasingPriceCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IncreasingPriceCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncreasingPriceCrowdsaleImplCallerSession struct {
	Contract *IncreasingPriceCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// IncreasingPriceCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncreasingPriceCrowdsaleImplTransactorSession struct {
	Contract     *IncreasingPriceCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// IncreasingPriceCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleImplRaw struct {
	Contract *IncreasingPriceCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// IncreasingPriceCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleImplCallerRaw struct {
	Contract *IncreasingPriceCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// IncreasingPriceCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleImplTransactorRaw struct {
	Contract *IncreasingPriceCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncreasingPriceCrowdsaleImpl creates a new instance of IncreasingPriceCrowdsaleImpl, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*IncreasingPriceCrowdsaleImpl, error) {
	contract, err := bindIncreasingPriceCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleImpl{IncreasingPriceCrowdsaleImplCaller: IncreasingPriceCrowdsaleImplCaller{contract: contract}, IncreasingPriceCrowdsaleImplTransactor: IncreasingPriceCrowdsaleImplTransactor{contract: contract}, IncreasingPriceCrowdsaleImplFilterer: IncreasingPriceCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewIncreasingPriceCrowdsaleImplCaller creates a new read-only instance of IncreasingPriceCrowdsaleImpl, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*IncreasingPriceCrowdsaleImplCaller, error) {
	contract, err := bindIncreasingPriceCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleImplCaller{contract: contract}, nil
}

// NewIncreasingPriceCrowdsaleImplTransactor creates a new write-only instance of IncreasingPriceCrowdsaleImpl, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*IncreasingPriceCrowdsaleImplTransactor, error) {
	contract, err := bindIncreasingPriceCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleImplTransactor{contract: contract}, nil
}

// NewIncreasingPriceCrowdsaleImplFilterer creates a new log filterer instance of IncreasingPriceCrowdsaleImpl, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*IncreasingPriceCrowdsaleImplFilterer, error) {
	contract, err := bindIncreasingPriceCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleImplFilterer{contract: contract}, nil
}

// bindIncreasingPriceCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindIncreasingPriceCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncreasingPriceCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncreasingPriceCrowdsaleImpl.Contract.IncreasingPriceCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.IncreasingPriceCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.IncreasingPriceCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncreasingPriceCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) ClosingTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.ClosingTime(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) ClosingTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.ClosingTime(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) FinalRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "finalRate")
	return *ret0, err
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) FinalRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.FinalRate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) FinalRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.FinalRate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) GetCurrentRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "getCurrentRate")
	return *ret0, err
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) GetCurrentRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.GetCurrentRate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) GetCurrentRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.GetCurrentRate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) HasClosed() (bool, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.HasClosed(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) HasClosed() (bool, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.HasClosed(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) InitialRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "initialRate")
	return *ret0, err
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) InitialRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.InitialRate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) InitialRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.InitialRate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) OpeningTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.OpeningTime(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) OpeningTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.OpeningTime(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.Rate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.Rate(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) Token() (common.Address, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.Token(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.Token(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.Wallet(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.Wallet(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncreasingPriceCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.WeiRaised(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.WeiRaised(&_IncreasingPriceCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.BuyTokens(&_IncreasingPriceCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsaleImpl.Contract.BuyTokens(&_IncreasingPriceCrowdsaleImpl.TransactOpts, _beneficiary)
}

// IncreasingPriceCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the IncreasingPriceCrowdsaleImpl contract.
type IncreasingPriceCrowdsaleImplTokenPurchaseIterator struct {
	Event *IncreasingPriceCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *IncreasingPriceCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncreasingPriceCrowdsaleImplTokenPurchase)
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
		it.Event = new(IncreasingPriceCrowdsaleImplTokenPurchase)
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
func (it *IncreasingPriceCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncreasingPriceCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncreasingPriceCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the IncreasingPriceCrowdsaleImpl contract.
type IncreasingPriceCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*IncreasingPriceCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IncreasingPriceCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleImplTokenPurchaseIterator{contract: _IncreasingPriceCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_IncreasingPriceCrowdsaleImpl *IncreasingPriceCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *IncreasingPriceCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IncreasingPriceCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncreasingPriceCrowdsaleImplTokenPurchase)
				if err := _IncreasingPriceCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
