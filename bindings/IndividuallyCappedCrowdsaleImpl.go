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

// IndividuallyCappedCrowdsaleImplABI is the input ABI used to generate the binding from.
const IndividuallyCappedCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"caps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getUserCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaries\",\"type\":\"address[]\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setGroupCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getUserContribution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setUserCap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// IndividuallyCappedCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const IndividuallyCappedCrowdsaleImplBin = `6060604052341561000f57600080fd5b6040516060806107a7833981016040528080519190602001805191906020018051915083905082826000831161004457600080fd5b600160a060020a038216151561005957600080fd5b600160a060020a038116151561006e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a03199182161790915560008054938316938216939093179092556004805433909216919092161790555050506106e2806100c56000396000f3006060604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e81146100cf5780634042b66f146100f457806342e94c9014610107578063521eb2731461012657806366d97b21146101555780638b58c64c146101745780638da5cb5b14610193578063a31f61fc146101a6578063bb8b2b47146101c8578063c3143fe5146101e7578063ec8ac4d814610209578063f2fde38b1461021d578063fc0c546a1461023c575b6100cd3361024f565b005b34156100da57600080fd5b6100e26102f7565b60405190815260200160405180910390f35b34156100ff57600080fd5b6100e26102fd565b341561011257600080fd5b6100e2600160a060020a0360043516610303565b341561013157600080fd5b610139610315565b604051600160a060020a03909116815260200160405180910390f35b341561016057600080fd5b6100e2600160a060020a0360043516610324565b341561017f57600080fd5b6100e2600160a060020a0360043516610336565b341561019e57600080fd5b610139610351565b34156101b157600080fd5b6100cd602460048035828101929101359035610360565b34156101d357600080fd5b6100e2600160a060020a03600435166103cc565b34156101f257600080fd5b6100cd600160a060020a03600435166024356103e7565b6100cd600160a060020a036004351661024f565b341561022857600080fd5b6100cd600160a060020a036004351661041e565b341561024757600080fd5b6101396104b9565b34600061025c83836104c8565b61026582610516565b60035490915061027b908363ffffffff61053316565b600355610288838261054d565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36102e08383610557565b6102e86105aa565b6102f28383610512565b505050565b60025481565b60035481565b60056020526000908152604090205481565b600154600160a060020a031681565b60066020526000908152604090205481565b600160a060020a031660009081526006602052604090205490565b600454600160a060020a031681565b60045460009033600160a060020a0390811691161461037e57600080fd5b5060005b828110156103c657816006600086868581811061039b57fe5b60209081029290920135600160a060020a031683525081019190915260400160002055600101610382565b50505050565b600160a060020a031660009081526005602052604090205490565b60045433600160a060020a0390811691161461040257600080fd5b600160a060020a03909116600090815260066020526040902055565b60045433600160a060020a0390811691161461043957600080fd5b600160a060020a038116151561044e57600080fd5b600454600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b6104d282826105e0565b600160a060020a038216600090815260066020908152604080832054600590925290912054610507908363ffffffff61053316565b111561051257600080fd5b5050565b600061052d6002548361060190919063ffffffff16565b92915050565b60008282018381101561054257fe5b8091505b5092915050565b610512828261062c565b6105618282610512565b600160a060020a03821660009081526005602052604090205461058a908263ffffffff61053316565b600160a060020a0390921660009081526005602052604090209190915550565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f1935050505015156105de57600080fd5b565b600160a060020a03821615156105f557600080fd5b80151561051257600080fd5b6000808315156106145760009150610546565b5082820282848281151561062457fe5b041461054257fe5b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561069b57600080fd5b5af115156106a857600080fd5b5050506040518051505050505600a165627a7a723058208b97606e4fabfeb67662d3a3bf1ac48857c411992697116bcdcb0dfaab01da010029`

// DeployIndividuallyCappedCrowdsaleImpl deploys a new Ethereum contract, binding an instance of IndividuallyCappedCrowdsaleImpl to it.
func DeployIndividuallyCappedCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *IndividuallyCappedCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(IndividuallyCappedCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IndividuallyCappedCrowdsaleImplBin), backend, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IndividuallyCappedCrowdsaleImpl{IndividuallyCappedCrowdsaleImplCaller: IndividuallyCappedCrowdsaleImplCaller{contract: contract}, IndividuallyCappedCrowdsaleImplTransactor: IndividuallyCappedCrowdsaleImplTransactor{contract: contract}}, nil
}

// IndividuallyCappedCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleImpl struct {
	IndividuallyCappedCrowdsaleImplCaller     // Read-only binding to the contract
	IndividuallyCappedCrowdsaleImplTransactor // Write-only binding to the contract
}

// IndividuallyCappedCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndividuallyCappedCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndividuallyCappedCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndividuallyCappedCrowdsaleImplSession struct {
	Contract     *IndividuallyCappedCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IndividuallyCappedCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndividuallyCappedCrowdsaleImplCallerSession struct {
	Contract *IndividuallyCappedCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// IndividuallyCappedCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndividuallyCappedCrowdsaleImplTransactorSession struct {
	Contract     *IndividuallyCappedCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// IndividuallyCappedCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleImplRaw struct {
	Contract *IndividuallyCappedCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// IndividuallyCappedCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleImplCallerRaw struct {
	Contract *IndividuallyCappedCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// IndividuallyCappedCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndividuallyCappedCrowdsaleImplTransactorRaw struct {
	Contract *IndividuallyCappedCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndividuallyCappedCrowdsaleImpl creates a new instance of IndividuallyCappedCrowdsaleImpl, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*IndividuallyCappedCrowdsaleImpl, error) {
	contract, err := bindIndividuallyCappedCrowdsaleImpl(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleImpl{IndividuallyCappedCrowdsaleImplCaller: IndividuallyCappedCrowdsaleImplCaller{contract: contract}, IndividuallyCappedCrowdsaleImplTransactor: IndividuallyCappedCrowdsaleImplTransactor{contract: contract}}, nil
}

// NewIndividuallyCappedCrowdsaleImplCaller creates a new read-only instance of IndividuallyCappedCrowdsaleImpl, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*IndividuallyCappedCrowdsaleImplCaller, error) {
	contract, err := bindIndividuallyCappedCrowdsaleImpl(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleImplCaller{contract: contract}, nil
}

// NewIndividuallyCappedCrowdsaleImplTransactor creates a new write-only instance of IndividuallyCappedCrowdsaleImpl, bound to a specific deployed contract.
func NewIndividuallyCappedCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*IndividuallyCappedCrowdsaleImplTransactor, error) {
	contract, err := bindIndividuallyCappedCrowdsaleImpl(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &IndividuallyCappedCrowdsaleImplTransactor{contract: contract}, nil
}

// bindIndividuallyCappedCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindIndividuallyCappedCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IndividuallyCappedCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IndividuallyCappedCrowdsaleImpl.Contract.IndividuallyCappedCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.IndividuallyCappedCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.IndividuallyCappedCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IndividuallyCappedCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) Caps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "caps", arg0)
	return *ret0, err
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) Caps(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Caps(&_IndividuallyCappedCrowdsaleImpl.CallOpts, arg0)
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) Caps(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Caps(&_IndividuallyCappedCrowdsaleImpl.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "contributions", arg0)
	return *ret0, err
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Contributions(&_IndividuallyCappedCrowdsaleImpl.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions( address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Contributions(&_IndividuallyCappedCrowdsaleImpl.CallOpts, arg0)
}

// GetUserCap is a free data retrieval call binding the contract method 0x8b58c64c.
//
// Solidity: function getUserCap(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) GetUserCap(opts *bind.CallOpts, _beneficiary common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "getUserCap", _beneficiary)
	return *ret0, err
}

// GetUserCap is a free data retrieval call binding the contract method 0x8b58c64c.
//
// Solidity: function getUserCap(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) GetUserCap(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.GetUserCap(&_IndividuallyCappedCrowdsaleImpl.CallOpts, _beneficiary)
}

// GetUserCap is a free data retrieval call binding the contract method 0x8b58c64c.
//
// Solidity: function getUserCap(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) GetUserCap(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.GetUserCap(&_IndividuallyCappedCrowdsaleImpl.CallOpts, _beneficiary)
}

// GetUserContribution is a free data retrieval call binding the contract method 0xbb8b2b47.
//
// Solidity: function getUserContribution(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) GetUserContribution(opts *bind.CallOpts, _beneficiary common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "getUserContribution", _beneficiary)
	return *ret0, err
}

// GetUserContribution is a free data retrieval call binding the contract method 0xbb8b2b47.
//
// Solidity: function getUserContribution(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) GetUserContribution(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.GetUserContribution(&_IndividuallyCappedCrowdsaleImpl.CallOpts, _beneficiary)
}

// GetUserContribution is a free data retrieval call binding the contract method 0xbb8b2b47.
//
// Solidity: function getUserContribution(_beneficiary address) constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) GetUserContribution(_beneficiary common.Address) (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.GetUserContribution(&_IndividuallyCappedCrowdsaleImpl.CallOpts, _beneficiary)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) Owner() (common.Address, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Owner(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) Owner() (common.Address, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Owner(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Rate(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Rate(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) Token() (common.Address, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Token(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Token(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Wallet(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.Wallet(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IndividuallyCappedCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.WeiRaised(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.WeiRaised(&_IndividuallyCappedCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.BuyTokens(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.BuyTokens(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// SetGroupCap is a paid mutator transaction binding the contract method 0xa31f61fc.
//
// Solidity: function setGroupCap(_beneficiaries address[], _cap uint256) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactor) SetGroupCap(opts *bind.TransactOpts, _beneficiaries []common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.contract.Transact(opts, "setGroupCap", _beneficiaries, _cap)
}

// SetGroupCap is a paid mutator transaction binding the contract method 0xa31f61fc.
//
// Solidity: function setGroupCap(_beneficiaries address[], _cap uint256) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) SetGroupCap(_beneficiaries []common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.SetGroupCap(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, _beneficiaries, _cap)
}

// SetGroupCap is a paid mutator transaction binding the contract method 0xa31f61fc.
//
// Solidity: function setGroupCap(_beneficiaries address[], _cap uint256) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactorSession) SetGroupCap(_beneficiaries []common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.SetGroupCap(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, _beneficiaries, _cap)
}

// SetUserCap is a paid mutator transaction binding the contract method 0xc3143fe5.
//
// Solidity: function setUserCap(_beneficiary address, _cap uint256) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactor) SetUserCap(opts *bind.TransactOpts, _beneficiary common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.contract.Transact(opts, "setUserCap", _beneficiary, _cap)
}

// SetUserCap is a paid mutator transaction binding the contract method 0xc3143fe5.
//
// Solidity: function setUserCap(_beneficiary address, _cap uint256) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) SetUserCap(_beneficiary common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.SetUserCap(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, _beneficiary, _cap)
}

// SetUserCap is a paid mutator transaction binding the contract method 0xc3143fe5.
//
// Solidity: function setUserCap(_beneficiary address, _cap uint256) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactorSession) SetUserCap(_beneficiary common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.SetUserCap(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, _beneficiary, _cap)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.TransferOwnership(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_IndividuallyCappedCrowdsaleImpl *IndividuallyCappedCrowdsaleImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IndividuallyCappedCrowdsaleImpl.Contract.TransferOwnership(&_IndividuallyCappedCrowdsaleImpl.TransactOpts, newOwner)
}
