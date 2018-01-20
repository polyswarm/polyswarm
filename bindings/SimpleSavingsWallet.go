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
	return address, tx, &SimpleSavingsWallet{SimpleSavingsWalletCaller: SimpleSavingsWalletCaller{contract: contract}, SimpleSavingsWalletTransactor: SimpleSavingsWalletTransactor{contract: contract}}, nil
}

// SimpleSavingsWallet is an auto generated Go binding around an Ethereum contract.
type SimpleSavingsWallet struct {
	SimpleSavingsWalletCaller     // Read-only binding to the contract
	SimpleSavingsWalletTransactor // Write-only binding to the contract
}

// SimpleSavingsWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleSavingsWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSavingsWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleSavingsWalletTransactor struct {
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
	contract, err := bindSimpleSavingsWallet(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWallet{SimpleSavingsWalletCaller: SimpleSavingsWalletCaller{contract: contract}, SimpleSavingsWalletTransactor: SimpleSavingsWalletTransactor{contract: contract}}, nil
}

// NewSimpleSavingsWalletCaller creates a new read-only instance of SimpleSavingsWallet, bound to a specific deployed contract.
func NewSimpleSavingsWalletCaller(address common.Address, caller bind.ContractCaller) (*SimpleSavingsWalletCaller, error) {
	contract, err := bindSimpleSavingsWallet(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletCaller{contract: contract}, nil
}

// NewSimpleSavingsWalletTransactor creates a new write-only instance of SimpleSavingsWallet, bound to a specific deployed contract.
func NewSimpleSavingsWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleSavingsWalletTransactor, error) {
	contract, err := bindSimpleSavingsWallet(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SimpleSavingsWalletTransactor{contract: contract}, nil
}

// bindSimpleSavingsWallet binds a generic wrapper to an already deployed contract.
func bindSimpleSavingsWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleSavingsWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
