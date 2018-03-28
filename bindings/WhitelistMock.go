// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WhitelistMockABI is the input ABI used to generate the binding from.
const WhitelistMockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyWhitelistedCanDoThis\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// WhitelistMockBin is the compiled bytecode used for deploying new contracts.
const WhitelistMockBin = `606060405260008054600160a060020a033316600160a060020a0319909116179055610523806100306000396000f30060606040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324953eaa8114610092578063286dd3f5146100f55780637b9417c8146101145780638da5cb5b146101335780639b19251a14610162578063e2ec6ec314610181578063e9cac389146101d0578063f2fde38b146101e5575b600080fd5b341561009d57600080fd5b6100e1600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061020495505050505050565b604051901515815260200160405180910390f35b341561010057600080fd5b6100e1600160a060020a0360043516610265565b341561011f57600080fd5b6100e1600160a060020a036004351661030a565b341561013e57600080fd5b6101466103b4565b604051600160a060020a03909116815260200160405180910390f35b341561016d57600080fd5b6100e1600160a060020a03600435166103c3565b341561018c57600080fd5b6100e160046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437509496506103d895505050505050565b34156101db57600080fd5b6101e3610433565b005b34156101f057600080fd5b6101e3600160a060020a036004351661045c565b60008054819033600160a060020a0390811691161461022257600080fd5b5060005b825181101561025f5761024d83828151811061023e57fe5b90602001906020020151610265565b1561025757600191505b600101610226565b50919050565b6000805433600160a060020a0390811691161461028157600080fd5b600160a060020a03821660009081526001602052604090205460ff161561030557600160a060020a03821660009081526001602052604090819020805460ff191690557ff1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a90839051600160a060020a03909116815260200160405180910390a15060015b919050565b6000805433600160a060020a0390811691161461032657600080fd5b600160a060020a03821660009081526001602052604090205460ff16151561030557600160a060020a038216600090815260016020819052604091829020805460ff191690911790557fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f90839051600160a060020a03909116815260200160405180910390a1506001919050565b600054600160a060020a031681565b60016020526000908152604090205460ff1681565b60008054819033600160a060020a039081169116146103f657600080fd5b5060005b825181101561025f5761042183828151811061041257fe5b9060200190602002015161030a565b1561042b57600191505b6001016103fa565b600160a060020a03331660009081526001602052604090205460ff16151561045a57600080fd5b565b60005433600160a060020a0390811691161461047757600080fd5b600160a060020a038116151561048c57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582012cacbf30a11eab124ceeaebb9a6b2ba4db6a866911e9599b9e76e3dbb28bdfe0029`

// DeployWhitelistMock deploys a new Ethereum contract, binding an instance of WhitelistMock to it.
func DeployWhitelistMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitelistMock, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistMock{WhitelistMockCaller: WhitelistMockCaller{contract: contract}, WhitelistMockTransactor: WhitelistMockTransactor{contract: contract}}, nil
}

// WhitelistMock is an auto generated Go binding around an Ethereum contract.
type WhitelistMock struct {
	WhitelistMockCaller     // Read-only binding to the contract
	WhitelistMockTransactor // Write-only binding to the contract
}

// WhitelistMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistMockSession struct {
	Contract     *WhitelistMock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistMockCallerSession struct {
	Contract *WhitelistMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WhitelistMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistMockTransactorSession struct {
	Contract     *WhitelistMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WhitelistMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistMockRaw struct {
	Contract *WhitelistMock // Generic contract binding to access the raw methods on
}

// WhitelistMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistMockCallerRaw struct {
	Contract *WhitelistMockCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistMockTransactorRaw struct {
	Contract *WhitelistMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistMock creates a new instance of WhitelistMock, bound to a specific deployed contract.
func NewWhitelistMock(address common.Address, backend bind.ContractBackend) (*WhitelistMock, error) {
	contract, err := bindWhitelistMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistMock{WhitelistMockCaller: WhitelistMockCaller{contract: contract}, WhitelistMockTransactor: WhitelistMockTransactor{contract: contract}}, nil
}

// NewWhitelistMockCaller creates a new read-only instance of WhitelistMock, bound to a specific deployed contract.
func NewWhitelistMockCaller(address common.Address, caller bind.ContractCaller) (*WhitelistMockCaller, error) {
	contract, err := bindWhitelistMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistMockCaller{contract: contract}, nil
}

// NewWhitelistMockTransactor creates a new write-only instance of WhitelistMock, bound to a specific deployed contract.
func NewWhitelistMockTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistMockTransactor, error) {
	contract, err := bindWhitelistMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WhitelistMockTransactor{contract: contract}, nil
}

// bindWhitelistMock binds a generic wrapper to an already deployed contract.
func bindWhitelistMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistMock *WhitelistMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistMock.Contract.WhitelistMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistMock *WhitelistMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistMock.Contract.WhitelistMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistMock *WhitelistMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistMock.Contract.WhitelistMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistMock *WhitelistMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistMock *WhitelistMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistMock *WhitelistMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistMock.Contract.contract.Transact(opts, method, params...)
}

// OnlyWhitelistedCanDoThis is a free data retrieval call binding the contract method 0xe9cac389.
//
// Solidity: function onlyWhitelistedCanDoThis() constant returns()
func (_WhitelistMock *WhitelistMockCaller) OnlyWhitelistedCanDoThis(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _WhitelistMock.contract.Call(opts, out, "onlyWhitelistedCanDoThis")
	return err
}

// OnlyWhitelistedCanDoThis is a free data retrieval call binding the contract method 0xe9cac389.
//
// Solidity: function onlyWhitelistedCanDoThis() constant returns()
func (_WhitelistMock *WhitelistMockSession) OnlyWhitelistedCanDoThis() error {
	return _WhitelistMock.Contract.OnlyWhitelistedCanDoThis(&_WhitelistMock.CallOpts)
}

// OnlyWhitelistedCanDoThis is a free data retrieval call binding the contract method 0xe9cac389.
//
// Solidity: function onlyWhitelistedCanDoThis() constant returns()
func (_WhitelistMock *WhitelistMockCallerSession) OnlyWhitelistedCanDoThis() error {
	return _WhitelistMock.Contract.OnlyWhitelistedCanDoThis(&_WhitelistMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistMock *WhitelistMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistMock.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistMock *WhitelistMockSession) Owner() (common.Address, error) {
	return _WhitelistMock.Contract.Owner(&_WhitelistMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WhitelistMock *WhitelistMockCallerSession) Owner() (common.Address, error) {
	return _WhitelistMock.Contract.Owner(&_WhitelistMock.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistMock *WhitelistMockCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistMock.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistMock *WhitelistMockSession) Whitelist(arg0 common.Address) (bool, error) {
	return _WhitelistMock.Contract.Whitelist(&_WhitelistMock.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_WhitelistMock *WhitelistMockCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _WhitelistMock.Contract.Whitelist(&_WhitelistMock.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _WhitelistMock.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_WhitelistMock *WhitelistMockSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.AddAddressToWhitelist(&_WhitelistMock.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.AddAddressToWhitelist(&_WhitelistMock.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _WhitelistMock.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_WhitelistMock *WhitelistMockSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.AddAddressesToWhitelist(&_WhitelistMock.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.AddAddressesToWhitelist(&_WhitelistMock.TransactOpts, addrs)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _WhitelistMock.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_WhitelistMock *WhitelistMockSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.RemoveAddressFromWhitelist(&_WhitelistMock.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.RemoveAddressFromWhitelist(&_WhitelistMock.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _WhitelistMock.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_WhitelistMock *WhitelistMockSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.RemoveAddressesFromWhitelist(&_WhitelistMock.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_WhitelistMock *WhitelistMockTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.RemoveAddressesFromWhitelist(&_WhitelistMock.TransactOpts, addrs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistMock *WhitelistMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistMock *WhitelistMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.TransferOwnership(&_WhitelistMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WhitelistMock *WhitelistMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistMock.Contract.TransferOwnership(&_WhitelistMock.TransactOpts, newOwner)
}
