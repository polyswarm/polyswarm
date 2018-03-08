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

// SimpleSavingsWalletABI is the input ABI used to generate the binding from.
const SimpleSavingsWalletABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimHeirOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newHeir\",\"type\":\"address\"}],\"name\":\"setHeir\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"proclaimDeath\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"heartbeat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heartbeatTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heir\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"payee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeOfDeath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeHeir\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_heartbeatTimeout\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newHeir\",\"type\":\"address\"}],\"name\":\"HeirChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerHeartbeated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"heir\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeOfDeath\",\"type\":\"uint256\"}],\"name\":\"OwnerProclaimedDead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"HeirOwnershipClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SimpleSavingsWalletBin is the compiled bytecode used for deploying new contracts.
const SimpleSavingsWalletBin = `6060604052341561000f57600080fd5b60405160208061079d8339810160405280805160008054600160a060020a03191633600160a060020a031617905591508190506100588164010000000061005f81026106891704565b50506100a6565b60005433600160a060020a0390811691161461007a57600080fd5b61008f64010000000061068261009f82021704565b151561009a57600080fd5b600255565b6003541590565b6106e8806100b56000396000f3006060604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c921e1681146100fe57806323defc771461011357806324845131146101325780633defb962146101455780637bca38be146101585780638da5cb5b1461017d57806391f2ebb8146101ac5780639e1a00aa146101bf578063b4a8f3e6146101e1578063ebf88de4146101f4578063f2fde38b14610207575b33600160a060020a03167f74cf3d18d0ddca79038197ad0dd2c7fa5005ef61a5d1ed190e8a8a437e2fcf103430600160a060020a03163160405191825260208201526040908101905180910390a2005b341561010957600080fd5b610111610226565b005b341561011e57600080fd5b610111600160a060020a0360043516610318565b341561013d57600080fd5b6101116103c1565b341561015057600080fd5b610111610441565b341561016357600080fd5b61016b61049b565b60405190815260200160405180910390f35b341561018857600080fd5b6101906104a1565b604051600160a060020a03909116815260200160405180910390f35b34156101b757600080fd5b6101906104b0565b34156101ca57600080fd5b610111600160a060020a03600435166024356104bf565b34156101ec57600080fd5b61016b61059f565b34156101ff57600080fd5b6101116105a5565b341561021257600080fd5b610111600160a060020a03600435166105e7565b60015433600160a060020a0390811691161461024157600080fd5b610249610682565b1561025357600080fd5b6002546003540142101561026657600080fd5b600154600054600160a060020a0391821691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600154600054600160a060020a0391821691167f1017a357e19071e4408dbd385f24e591aa5bcee52b444dc0c8abddbe6ad29de660405160405180910390a36001546000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091178155600355565b60005433600160a060020a0390811691161461033357600080fd5b600054600160a060020a038281169116141561034e57600080fd5b610356610441565b600054600160a060020a0380831691167f4e6093f85fa64484abd692810d8a44d508792ff7b7a021d9fbd69fa1c6ff18a060405160405180910390a36001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015433600160a060020a039081169116146103dc57600080fd5b6103e4610682565b15156103ef57600080fd5b600154600054600354600160a060020a0392831692909116907f66389f1f3251c39e71cb18351daefea45bda23e98df0767266b8aed4ff9382779060405190815260200160405180910390a342600355565b60005433600160a060020a0390811691161461045c57600080fd5b600054600160a060020a03167fd40b9d9a7572411aff4bbb95ba7bd9a9d4e9b70747ec46f8fe7913c309b6845260405160405180910390a26000600355565b60025481565b600054600160a060020a031681565b600154600160a060020a031681565b60005433600160a060020a039081169116146104da57600080fd5b600160a060020a03821615801590610504575030600160a060020a031682600160a060020a031614155b151561050f57600080fd5b6000811161051c57600080fd5b600160a060020a03821681156108fc0282604051600060405180830381858888f19350505050151561054d57600080fd5b81600160a060020a03167f6356739d963da01dc3533acba7203430fcc14f2175d48a8dd0973d7db49c785e8230600160a060020a03163160405191825260208201526040908101905180910390a25050565b60035481565b60005433600160a060020a039081169116146105c057600080fd5b6105c8610441565b6001805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a0390811691161461060257600080fd5b600160a060020a038116151561061757600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6003541590565b60005433600160a060020a039081169116146106a457600080fd5b6106ac610682565b15156106b757600080fd5b6002555600a165627a7a72305820eca33a0bec202eaa29cac279d62dba235e580b7703fc05651b717e9e8fa3f1790029`

// DeploySimpleSavingsWallet deploys a new Ethereum contract, binding an instance of SimpleSavingsWallet to it.
func DeploySimpleSavingsWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _heartbeatTimeout *big.Int) (common.Address, *types.Transaction, *SimpleSavingsWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleSavingsWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleSavingsWalletBin), backend, _heartbeatTimeout)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleSavingsWallet{SimpleSavingsWalletCaller: SimpleSavingsWalletCaller{contract: contract}, SimpleSavingsWalletTransactor: SimpleSavingsWalletTransactor{contract: contract}, SimpleSavingsWalletFilterer: SimpleSavingsWalletFilterer{contract: contract}}, nil
}

// SimpleSavingsWallet is an auto generated Go binding around an Ethereum contract.
type SimpleSavingsWallet struct {
	SimpleSavingsWalletCaller     // Read-only binding to the contract
	SimpleSavingsWalletTransactor // Write-only binding to the contract
	SimpleSavingsWalletFilterer   // Log filterer for contract events
}

// SimpleSavingsWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleSavingsWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSavingsWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleSavingsWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSavingsWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleSavingsWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSavingsWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSavingsWalletSession struct {
	Contract     *SimpleSavingsWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SimpleSavingsWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleSavingsWalletCallerSession struct {
	Contract *SimpleSavingsWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SimpleSavingsWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleSavingsWalletTransactorSession struct {
	Contract     *SimpleSavingsWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SimpleSavingsWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleSavingsWalletRaw struct {
	Contract *SimpleSavingsWallet // Generic contract binding to access the raw methods on
}

// SimpleSavingsWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleSavingsWalletCallerRaw struct {
	Contract *SimpleSavingsWalletCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleSavingsWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleSavingsWalletTransactorRaw struct {
	Contract *SimpleSavingsWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleSavingsWallet creates a new instance of SimpleSavingsWallet, bound to a specific deployed contract.
func NewSimpleSavingsWallet(address common.Address, backend bind.ContractBackend) (*SimpleSavingsWallet, error) {
	contract, err := bindSimpleSavingsWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWallet{SimpleSavingsWalletCaller: SimpleSavingsWalletCaller{contract: contract}, SimpleSavingsWalletTransactor: SimpleSavingsWalletTransactor{contract: contract}, SimpleSavingsWalletFilterer: SimpleSavingsWalletFilterer{contract: contract}}, nil
}

// NewSimpleSavingsWalletCaller creates a new read-only instance of SimpleSavingsWallet, bound to a specific deployed contract.
func NewSimpleSavingsWalletCaller(address common.Address, caller bind.ContractCaller) (*SimpleSavingsWalletCaller, error) {
	contract, err := bindSimpleSavingsWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletCaller{contract: contract}, nil
}

// NewSimpleSavingsWalletTransactor creates a new write-only instance of SimpleSavingsWallet, bound to a specific deployed contract.
func NewSimpleSavingsWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleSavingsWalletTransactor, error) {
	contract, err := bindSimpleSavingsWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletTransactor{contract: contract}, nil
}

// NewSimpleSavingsWalletFilterer creates a new log filterer instance of SimpleSavingsWallet, bound to a specific deployed contract.
func NewSimpleSavingsWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleSavingsWalletFilterer, error) {
	contract, err := bindSimpleSavingsWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletFilterer{contract: contract}, nil
}

// bindSimpleSavingsWallet binds a generic wrapper to an already deployed contract.
func bindSimpleSavingsWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleSavingsWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleSavingsWallet *SimpleSavingsWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleSavingsWallet.Contract.SimpleSavingsWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleSavingsWallet *SimpleSavingsWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.SimpleSavingsWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleSavingsWallet *SimpleSavingsWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.SimpleSavingsWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleSavingsWallet *SimpleSavingsWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleSavingsWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.contract.Transact(opts, method, params...)
}

// HeartbeatTimeout is a free data retrieval call binding the contract method 0x7bca38be.
//
// Solidity: function heartbeatTimeout() constant returns(uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletCaller) HeartbeatTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleSavingsWallet.contract.Call(opts, out, "heartbeatTimeout")
	return *ret0, err
}

// HeartbeatTimeout is a free data retrieval call binding the contract method 0x7bca38be.
//
// Solidity: function heartbeatTimeout() constant returns(uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) HeartbeatTimeout() (*big.Int, error) {
	return _SimpleSavingsWallet.Contract.HeartbeatTimeout(&_SimpleSavingsWallet.CallOpts)
}

// HeartbeatTimeout is a free data retrieval call binding the contract method 0x7bca38be.
//
// Solidity: function heartbeatTimeout() constant returns(uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletCallerSession) HeartbeatTimeout() (*big.Int, error) {
	return _SimpleSavingsWallet.Contract.HeartbeatTimeout(&_SimpleSavingsWallet.CallOpts)
}

// Heir is a free data retrieval call binding the contract method 0x91f2ebb8.
//
// Solidity: function heir() constant returns(address)
func (_SimpleSavingsWallet *SimpleSavingsWalletCaller) Heir(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimpleSavingsWallet.contract.Call(opts, out, "heir")
	return *ret0, err
}

// Heir is a free data retrieval call binding the contract method 0x91f2ebb8.
//
// Solidity: function heir() constant returns(address)
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) Heir() (common.Address, error) {
	return _SimpleSavingsWallet.Contract.Heir(&_SimpleSavingsWallet.CallOpts)
}

// Heir is a free data retrieval call binding the contract method 0x91f2ebb8.
//
// Solidity: function heir() constant returns(address)
func (_SimpleSavingsWallet *SimpleSavingsWalletCallerSession) Heir() (common.Address, error) {
	return _SimpleSavingsWallet.Contract.Heir(&_SimpleSavingsWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleSavingsWallet *SimpleSavingsWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimpleSavingsWallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) Owner() (common.Address, error) {
	return _SimpleSavingsWallet.Contract.Owner(&_SimpleSavingsWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleSavingsWallet *SimpleSavingsWalletCallerSession) Owner() (common.Address, error) {
	return _SimpleSavingsWallet.Contract.Owner(&_SimpleSavingsWallet.CallOpts)
}

// TimeOfDeath is a free data retrieval call binding the contract method 0xb4a8f3e6.
//
// Solidity: function timeOfDeath() constant returns(uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletCaller) TimeOfDeath(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleSavingsWallet.contract.Call(opts, out, "timeOfDeath")
	return *ret0, err
}

// TimeOfDeath is a free data retrieval call binding the contract method 0xb4a8f3e6.
//
// Solidity: function timeOfDeath() constant returns(uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) TimeOfDeath() (*big.Int, error) {
	return _SimpleSavingsWallet.Contract.TimeOfDeath(&_SimpleSavingsWallet.CallOpts)
}

// TimeOfDeath is a free data retrieval call binding the contract method 0xb4a8f3e6.
//
// Solidity: function timeOfDeath() constant returns(uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletCallerSession) TimeOfDeath() (*big.Int, error) {
	return _SimpleSavingsWallet.Contract.TimeOfDeath(&_SimpleSavingsWallet.CallOpts)
}

// ClaimHeirOwnership is a paid mutator transaction binding the contract method 0x1c921e16.
//
// Solidity: function claimHeirOwnership() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) ClaimHeirOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "claimHeirOwnership")
}

// ClaimHeirOwnership is a paid mutator transaction binding the contract method 0x1c921e16.
//
// Solidity: function claimHeirOwnership() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) ClaimHeirOwnership() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.ClaimHeirOwnership(&_SimpleSavingsWallet.TransactOpts)
}

// ClaimHeirOwnership is a paid mutator transaction binding the contract method 0x1c921e16.
//
// Solidity: function claimHeirOwnership() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) ClaimHeirOwnership() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.ClaimHeirOwnership(&_SimpleSavingsWallet.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) Heartbeat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "heartbeat")
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) Heartbeat() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.Heartbeat(&_SimpleSavingsWallet.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) Heartbeat() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.Heartbeat(&_SimpleSavingsWallet.TransactOpts)
}

// ProclaimDeath is a paid mutator transaction binding the contract method 0x24845131.
//
// Solidity: function proclaimDeath() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) ProclaimDeath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "proclaimDeath")
}

// ProclaimDeath is a paid mutator transaction binding the contract method 0x24845131.
//
// Solidity: function proclaimDeath() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) ProclaimDeath() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.ProclaimDeath(&_SimpleSavingsWallet.TransactOpts)
}

// ProclaimDeath is a paid mutator transaction binding the contract method 0x24845131.
//
// Solidity: function proclaimDeath() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) ProclaimDeath() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.ProclaimDeath(&_SimpleSavingsWallet.TransactOpts)
}

// RemoveHeir is a paid mutator transaction binding the contract method 0xebf88de4.
//
// Solidity: function removeHeir() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) RemoveHeir(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "removeHeir")
}

// RemoveHeir is a paid mutator transaction binding the contract method 0xebf88de4.
//
// Solidity: function removeHeir() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) RemoveHeir() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.RemoveHeir(&_SimpleSavingsWallet.TransactOpts)
}

// RemoveHeir is a paid mutator transaction binding the contract method 0xebf88de4.
//
// Solidity: function removeHeir() returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) RemoveHeir() (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.RemoveHeir(&_SimpleSavingsWallet.TransactOpts)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(payee address, amount uint256) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) SendTo(opts *bind.TransactOpts, payee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "sendTo", payee, amount)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(payee address, amount uint256) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) SendTo(payee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.SendTo(&_SimpleSavingsWallet.TransactOpts, payee, amount)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(payee address, amount uint256) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) SendTo(payee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.SendTo(&_SimpleSavingsWallet.TransactOpts, payee, amount)
}

// SetHeir is a paid mutator transaction binding the contract method 0x23defc77.
//
// Solidity: function setHeir(newHeir address) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) SetHeir(opts *bind.TransactOpts, newHeir common.Address) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "setHeir", newHeir)
}

// SetHeir is a paid mutator transaction binding the contract method 0x23defc77.
//
// Solidity: function setHeir(newHeir address) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) SetHeir(newHeir common.Address) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.SetHeir(&_SimpleSavingsWallet.TransactOpts, newHeir)
}

// SetHeir is a paid mutator transaction binding the contract method 0x23defc77.
//
// Solidity: function setHeir(newHeir address) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) SetHeir(newHeir common.Address) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.SetHeir(&_SimpleSavingsWallet.TransactOpts, newHeir)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SimpleSavingsWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.TransferOwnership(&_SimpleSavingsWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SimpleSavingsWallet *SimpleSavingsWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleSavingsWallet.Contract.TransferOwnership(&_SimpleSavingsWallet.TransactOpts, newOwner)
}

// SimpleSavingsWalletHeirChangedIterator is returned from FilterHeirChanged and is used to iterate over the raw logs and unpacked data for HeirChanged events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletHeirChangedIterator struct {
	Event *SimpleSavingsWalletHeirChanged // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletHeirChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletHeirChanged)
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
		it.Event = new(SimpleSavingsWalletHeirChanged)
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
func (it *SimpleSavingsWalletHeirChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletHeirChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletHeirChanged represents a HeirChanged event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletHeirChanged struct {
	Owner   common.Address
	NewHeir common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHeirChanged is a free log retrieval operation binding the contract event 0x4e6093f85fa64484abd692810d8a44d508792ff7b7a021d9fbd69fa1c6ff18a0.
//
// Solidity: event HeirChanged(owner indexed address, newHeir indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterHeirChanged(opts *bind.FilterOpts, owner []common.Address, newHeir []common.Address) (*SimpleSavingsWalletHeirChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newHeirRule []interface{}
	for _, newHeirItem := range newHeir {
		newHeirRule = append(newHeirRule, newHeirItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "HeirChanged", ownerRule, newHeirRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletHeirChangedIterator{contract: _SimpleSavingsWallet.contract, event: "HeirChanged", logs: logs, sub: sub}, nil
}

// WatchHeirChanged is a free log subscription operation binding the contract event 0x4e6093f85fa64484abd692810d8a44d508792ff7b7a021d9fbd69fa1c6ff18a0.
//
// Solidity: event HeirChanged(owner indexed address, newHeir indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchHeirChanged(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletHeirChanged, owner []common.Address, newHeir []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newHeirRule []interface{}
	for _, newHeirItem := range newHeir {
		newHeirRule = append(newHeirRule, newHeirItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "HeirChanged", ownerRule, newHeirRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletHeirChanged)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "HeirChanged", log); err != nil {
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

// SimpleSavingsWalletHeirOwnershipClaimedIterator is returned from FilterHeirOwnershipClaimed and is used to iterate over the raw logs and unpacked data for HeirOwnershipClaimed events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletHeirOwnershipClaimedIterator struct {
	Event *SimpleSavingsWalletHeirOwnershipClaimed // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletHeirOwnershipClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletHeirOwnershipClaimed)
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
		it.Event = new(SimpleSavingsWalletHeirOwnershipClaimed)
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
func (it *SimpleSavingsWalletHeirOwnershipClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletHeirOwnershipClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletHeirOwnershipClaimed represents a HeirOwnershipClaimed event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletHeirOwnershipClaimed struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHeirOwnershipClaimed is a free log retrieval operation binding the contract event 0x1017a357e19071e4408dbd385f24e591aa5bcee52b444dc0c8abddbe6ad29de6.
//
// Solidity: event HeirOwnershipClaimed(previousOwner indexed address, newOwner indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterHeirOwnershipClaimed(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SimpleSavingsWalletHeirOwnershipClaimedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "HeirOwnershipClaimed", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletHeirOwnershipClaimedIterator{contract: _SimpleSavingsWallet.contract, event: "HeirOwnershipClaimed", logs: logs, sub: sub}, nil
}

// WatchHeirOwnershipClaimed is a free log subscription operation binding the contract event 0x1017a357e19071e4408dbd385f24e591aa5bcee52b444dc0c8abddbe6ad29de6.
//
// Solidity: event HeirOwnershipClaimed(previousOwner indexed address, newOwner indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchHeirOwnershipClaimed(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletHeirOwnershipClaimed, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "HeirOwnershipClaimed", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletHeirOwnershipClaimed)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "HeirOwnershipClaimed", log); err != nil {
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

// SimpleSavingsWalletOwnerHeartbeatedIterator is returned from FilterOwnerHeartbeated and is used to iterate over the raw logs and unpacked data for OwnerHeartbeated events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletOwnerHeartbeatedIterator struct {
	Event *SimpleSavingsWalletOwnerHeartbeated // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletOwnerHeartbeatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletOwnerHeartbeated)
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
		it.Event = new(SimpleSavingsWalletOwnerHeartbeated)
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
func (it *SimpleSavingsWalletOwnerHeartbeatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletOwnerHeartbeatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletOwnerHeartbeated represents a OwnerHeartbeated event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletOwnerHeartbeated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerHeartbeated is a free log retrieval operation binding the contract event 0xd40b9d9a7572411aff4bbb95ba7bd9a9d4e9b70747ec46f8fe7913c309b68452.
//
// Solidity: event OwnerHeartbeated(owner indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterOwnerHeartbeated(opts *bind.FilterOpts, owner []common.Address) (*SimpleSavingsWalletOwnerHeartbeatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "OwnerHeartbeated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletOwnerHeartbeatedIterator{contract: _SimpleSavingsWallet.contract, event: "OwnerHeartbeated", logs: logs, sub: sub}, nil
}

// WatchOwnerHeartbeated is a free log subscription operation binding the contract event 0xd40b9d9a7572411aff4bbb95ba7bd9a9d4e9b70747ec46f8fe7913c309b68452.
//
// Solidity: event OwnerHeartbeated(owner indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchOwnerHeartbeated(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletOwnerHeartbeated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "OwnerHeartbeated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletOwnerHeartbeated)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "OwnerHeartbeated", log); err != nil {
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

// SimpleSavingsWalletOwnerProclaimedDeadIterator is returned from FilterOwnerProclaimedDead and is used to iterate over the raw logs and unpacked data for OwnerProclaimedDead events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletOwnerProclaimedDeadIterator struct {
	Event *SimpleSavingsWalletOwnerProclaimedDead // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletOwnerProclaimedDeadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletOwnerProclaimedDead)
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
		it.Event = new(SimpleSavingsWalletOwnerProclaimedDead)
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
func (it *SimpleSavingsWalletOwnerProclaimedDeadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletOwnerProclaimedDeadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletOwnerProclaimedDead represents a OwnerProclaimedDead event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletOwnerProclaimedDead struct {
	Owner       common.Address
	Heir        common.Address
	TimeOfDeath *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOwnerProclaimedDead is a free log retrieval operation binding the contract event 0x66389f1f3251c39e71cb18351daefea45bda23e98df0767266b8aed4ff938277.
//
// Solidity: event OwnerProclaimedDead(owner indexed address, heir indexed address, timeOfDeath uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterOwnerProclaimedDead(opts *bind.FilterOpts, owner []common.Address, heir []common.Address) (*SimpleSavingsWalletOwnerProclaimedDeadIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var heirRule []interface{}
	for _, heirItem := range heir {
		heirRule = append(heirRule, heirItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "OwnerProclaimedDead", ownerRule, heirRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletOwnerProclaimedDeadIterator{contract: _SimpleSavingsWallet.contract, event: "OwnerProclaimedDead", logs: logs, sub: sub}, nil
}

// WatchOwnerProclaimedDead is a free log subscription operation binding the contract event 0x66389f1f3251c39e71cb18351daefea45bda23e98df0767266b8aed4ff938277.
//
// Solidity: event OwnerProclaimedDead(owner indexed address, heir indexed address, timeOfDeath uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchOwnerProclaimedDead(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletOwnerProclaimedDead, owner []common.Address, heir []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var heirRule []interface{}
	for _, heirItem := range heir {
		heirRule = append(heirRule, heirItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "OwnerProclaimedDead", ownerRule, heirRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletOwnerProclaimedDead)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "OwnerProclaimedDead", log); err != nil {
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

// SimpleSavingsWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletOwnershipTransferredIterator struct {
	Event *SimpleSavingsWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletOwnershipTransferred)
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
		it.Event = new(SimpleSavingsWalletOwnershipTransferred)
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
func (it *SimpleSavingsWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletOwnershipTransferred represents a OwnershipTransferred event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SimpleSavingsWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletOwnershipTransferredIterator{contract: _SimpleSavingsWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletOwnershipTransferred)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SimpleSavingsWalletReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletReceivedIterator struct {
	Event *SimpleSavingsWalletReceived // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletReceived)
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
		it.Event = new(SimpleSavingsWalletReceived)
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
func (it *SimpleSavingsWalletReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletReceived represents a Received event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletReceived struct {
	Payer   common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x74cf3d18d0ddca79038197ad0dd2c7fa5005ef61a5d1ed190e8a8a437e2fcf10.
//
// Solidity: event Received(payer indexed address, amount uint256, balance uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterReceived(opts *bind.FilterOpts, payer []common.Address) (*SimpleSavingsWalletReceivedIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "Received", payerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletReceivedIterator{contract: _SimpleSavingsWallet.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x74cf3d18d0ddca79038197ad0dd2c7fa5005ef61a5d1ed190e8a8a437e2fcf10.
//
// Solidity: event Received(payer indexed address, amount uint256, balance uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletReceived, payer []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "Received", payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletReceived)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "Received", log); err != nil {
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

// SimpleSavingsWalletSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletSentIterator struct {
	Event *SimpleSavingsWalletSent // Event containing the contract specifics and raw log

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
func (it *SimpleSavingsWalletSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleSavingsWalletSent)
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
		it.Event = new(SimpleSavingsWalletSent)
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
func (it *SimpleSavingsWalletSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleSavingsWalletSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleSavingsWalletSent represents a Sent event raised by the SimpleSavingsWallet contract.
type SimpleSavingsWalletSent struct {
	Payee   common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x6356739d963da01dc3533acba7203430fcc14f2175d48a8dd0973d7db49c785e.
//
// Solidity: event Sent(payee indexed address, amount uint256, balance uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) FilterSent(opts *bind.FilterOpts, payee []common.Address) (*SimpleSavingsWalletSentIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.FilterLogs(opts, "Sent", payeeRule)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletSentIterator{contract: _SimpleSavingsWallet.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x6356739d963da01dc3533acba7203430fcc14f2175d48a8dd0973d7db49c785e.
//
// Solidity: event Sent(payee indexed address, amount uint256, balance uint256)
func (_SimpleSavingsWallet *SimpleSavingsWalletFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *SimpleSavingsWalletSent, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _SimpleSavingsWallet.contract.WatchLogs(opts, "Sent", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleSavingsWalletSent)
				if err := _SimpleSavingsWallet.contract.UnpackLog(event, "Sent", log); err != nil {
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
