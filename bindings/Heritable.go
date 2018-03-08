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

// HeritableABI is the input ABI used to generate the binding from.
const HeritableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimHeirOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newHeir\",\"type\":\"address\"}],\"name\":\"setHeir\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"proclaimDeath\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"heartbeat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heartbeatTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heir\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeOfDeath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeHeir\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_heartbeatTimeout\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newHeir\",\"type\":\"address\"}],\"name\":\"HeirChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerHeartbeated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"heir\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeOfDeath\",\"type\":\"uint256\"}],\"name\":\"OwnerProclaimedDead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"HeirOwnershipClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HeritableBin is the compiled bytecode used for deploying new contracts.
const HeritableBin = `6060604052341561000f57600080fd5b6040516020806106438339810160405280805160008054600160a060020a03191633600160a060020a0316179055915061005790508164010000000061005d81026105311704565b506100a4565b60005433600160a060020a0390811691161461007857600080fd5b61008d64010000000061052a61009d82021704565b151561009857600080fd5b600255565b6003541590565b610590806100b36000396000f3006060604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c921e1681146100a857806323defc77146100bd57806324845131146100dc5780633defb962146100ef5780637bca38be146101025780638da5cb5b1461012757806391f2ebb814610156578063b4a8f3e614610169578063ebf88de41461017c578063f2fde38b1461018f575b600080fd5b34156100b357600080fd5b6100bb6101ae565b005b34156100c857600080fd5b6100bb600160a060020a03600435166102a0565b34156100e757600080fd5b6100bb610349565b34156100fa57600080fd5b6100bb6103c9565b341561010d57600080fd5b610115610423565b60405190815260200160405180910390f35b341561013257600080fd5b61013a610429565b604051600160a060020a03909116815260200160405180910390f35b341561016157600080fd5b61013a610438565b341561017457600080fd5b610115610447565b341561018757600080fd5b6100bb61044d565b341561019a57600080fd5b6100bb600160a060020a036004351661048f565b60015433600160a060020a039081169116146101c957600080fd5b6101d161052a565b156101db57600080fd5b600254600354014210156101ee57600080fd5b600154600054600160a060020a0391821691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600154600054600160a060020a0391821691167f1017a357e19071e4408dbd385f24e591aa5bcee52b444dc0c8abddbe6ad29de660405160405180910390a36001546000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091178155600355565b60005433600160a060020a039081169116146102bb57600080fd5b600054600160a060020a03828116911614156102d657600080fd5b6102de6103c9565b600054600160a060020a0380831691167f4e6093f85fa64484abd692810d8a44d508792ff7b7a021d9fbd69fa1c6ff18a060405160405180910390a36001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015433600160a060020a0390811691161461036457600080fd5b61036c61052a565b151561037757600080fd5b600154600054600354600160a060020a0392831692909116907f66389f1f3251c39e71cb18351daefea45bda23e98df0767266b8aed4ff9382779060405190815260200160405180910390a342600355565b60005433600160a060020a039081169116146103e457600080fd5b600054600160a060020a03167fd40b9d9a7572411aff4bbb95ba7bd9a9d4e9b70747ec46f8fe7913c309b6845260405160405180910390a26000600355565b60025481565b600054600160a060020a031681565b600154600160a060020a031681565b60035481565b60005433600160a060020a0390811691161461046857600080fd5b6104706103c9565b6001805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a039081169116146104aa57600080fd5b600160a060020a03811615156104bf57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6003541590565b60005433600160a060020a0390811691161461054c57600080fd5b61055461052a565b151561055f57600080fd5b6002555600a165627a7a72305820c062ecc8a5f82b40b621976fd3f97ce41a022d15631a21bca153d34f38b4235f0029`

// DeployHeritable deploys a new Ethereum contract, binding an instance of Heritable to it.
func DeployHeritable(auth *bind.TransactOpts, backend bind.ContractBackend, _heartbeatTimeout *big.Int) (common.Address, *types.Transaction, *Heritable, error) {
	parsed, err := abi.JSON(strings.NewReader(HeritableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HeritableBin), backend, _heartbeatTimeout)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Heritable{HeritableCaller: HeritableCaller{contract: contract}, HeritableTransactor: HeritableTransactor{contract: contract}, HeritableFilterer: HeritableFilterer{contract: contract}}, nil
}

// Heritable is an auto generated Go binding around an Ethereum contract.
type Heritable struct {
	HeritableCaller     // Read-only binding to the contract
	HeritableTransactor // Write-only binding to the contract
	HeritableFilterer   // Log filterer for contract events
}

// HeritableCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeritableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeritableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeritableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeritableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeritableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeritableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeritableSession struct {
	Contract     *Heritable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeritableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeritableCallerSession struct {
	Contract *HeritableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HeritableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeritableTransactorSession struct {
	Contract     *HeritableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HeritableRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeritableRaw struct {
	Contract *Heritable // Generic contract binding to access the raw methods on
}

// HeritableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeritableCallerRaw struct {
	Contract *HeritableCaller // Generic read-only contract binding to access the raw methods on
}

// HeritableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeritableTransactorRaw struct {
	Contract *HeritableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeritable creates a new instance of Heritable, bound to a specific deployed contract.
func NewHeritable(address common.Address, backend bind.ContractBackend) (*Heritable, error) {
	contract, err := bindHeritable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Heritable{HeritableCaller: HeritableCaller{contract: contract}, HeritableTransactor: HeritableTransactor{contract: contract}, HeritableFilterer: HeritableFilterer{contract: contract}}, nil
}

// NewHeritableCaller creates a new read-only instance of Heritable, bound to a specific deployed contract.
func NewHeritableCaller(address common.Address, caller bind.ContractCaller) (*HeritableCaller, error) {
	contract, err := bindHeritable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeritableCaller{contract: contract}, nil
}

// NewHeritableTransactor creates a new write-only instance of Heritable, bound to a specific deployed contract.
func NewHeritableTransactor(address common.Address, transactor bind.ContractTransactor) (*HeritableTransactor, error) {
	contract, err := bindHeritable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeritableTransactor{contract: contract}, nil
}

// NewHeritableFilterer creates a new log filterer instance of Heritable, bound to a specific deployed contract.
func NewHeritableFilterer(address common.Address, filterer bind.ContractFilterer) (*HeritableFilterer, error) {
	contract, err := bindHeritable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeritableFilterer{contract: contract}, nil
}

// bindHeritable binds a generic wrapper to an already deployed contract.
func bindHeritable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeritableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heritable *HeritableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Heritable.Contract.HeritableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heritable *HeritableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heritable.Contract.HeritableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heritable *HeritableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heritable.Contract.HeritableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heritable *HeritableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Heritable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heritable *HeritableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heritable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heritable *HeritableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heritable.Contract.contract.Transact(opts, method, params...)
}

// HeartbeatTimeout is a free data retrieval call binding the contract method 0x7bca38be.
//
// Solidity: function heartbeatTimeout() constant returns(uint256)
func (_Heritable *HeritableCaller) HeartbeatTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heritable.contract.Call(opts, out, "heartbeatTimeout")
	return *ret0, err
}

// HeartbeatTimeout is a free data retrieval call binding the contract method 0x7bca38be.
//
// Solidity: function heartbeatTimeout() constant returns(uint256)
func (_Heritable *HeritableSession) HeartbeatTimeout() (*big.Int, error) {
	return _Heritable.Contract.HeartbeatTimeout(&_Heritable.CallOpts)
}

// HeartbeatTimeout is a free data retrieval call binding the contract method 0x7bca38be.
//
// Solidity: function heartbeatTimeout() constant returns(uint256)
func (_Heritable *HeritableCallerSession) HeartbeatTimeout() (*big.Int, error) {
	return _Heritable.Contract.HeartbeatTimeout(&_Heritable.CallOpts)
}

// Heir is a free data retrieval call binding the contract method 0x91f2ebb8.
//
// Solidity: function heir() constant returns(address)
func (_Heritable *HeritableCaller) Heir(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heritable.contract.Call(opts, out, "heir")
	return *ret0, err
}

// Heir is a free data retrieval call binding the contract method 0x91f2ebb8.
//
// Solidity: function heir() constant returns(address)
func (_Heritable *HeritableSession) Heir() (common.Address, error) {
	return _Heritable.Contract.Heir(&_Heritable.CallOpts)
}

// Heir is a free data retrieval call binding the contract method 0x91f2ebb8.
//
// Solidity: function heir() constant returns(address)
func (_Heritable *HeritableCallerSession) Heir() (common.Address, error) {
	return _Heritable.Contract.Heir(&_Heritable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Heritable *HeritableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heritable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Heritable *HeritableSession) Owner() (common.Address, error) {
	return _Heritable.Contract.Owner(&_Heritable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Heritable *HeritableCallerSession) Owner() (common.Address, error) {
	return _Heritable.Contract.Owner(&_Heritable.CallOpts)
}

// TimeOfDeath is a free data retrieval call binding the contract method 0xb4a8f3e6.
//
// Solidity: function timeOfDeath() constant returns(uint256)
func (_Heritable *HeritableCaller) TimeOfDeath(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heritable.contract.Call(opts, out, "timeOfDeath")
	return *ret0, err
}

// TimeOfDeath is a free data retrieval call binding the contract method 0xb4a8f3e6.
//
// Solidity: function timeOfDeath() constant returns(uint256)
func (_Heritable *HeritableSession) TimeOfDeath() (*big.Int, error) {
	return _Heritable.Contract.TimeOfDeath(&_Heritable.CallOpts)
}

// TimeOfDeath is a free data retrieval call binding the contract method 0xb4a8f3e6.
//
// Solidity: function timeOfDeath() constant returns(uint256)
func (_Heritable *HeritableCallerSession) TimeOfDeath() (*big.Int, error) {
	return _Heritable.Contract.TimeOfDeath(&_Heritable.CallOpts)
}

// ClaimHeirOwnership is a paid mutator transaction binding the contract method 0x1c921e16.
//
// Solidity: function claimHeirOwnership() returns()
func (_Heritable *HeritableTransactor) ClaimHeirOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heritable.contract.Transact(opts, "claimHeirOwnership")
}

// ClaimHeirOwnership is a paid mutator transaction binding the contract method 0x1c921e16.
//
// Solidity: function claimHeirOwnership() returns()
func (_Heritable *HeritableSession) ClaimHeirOwnership() (*types.Transaction, error) {
	return _Heritable.Contract.ClaimHeirOwnership(&_Heritable.TransactOpts)
}

// ClaimHeirOwnership is a paid mutator transaction binding the contract method 0x1c921e16.
//
// Solidity: function claimHeirOwnership() returns()
func (_Heritable *HeritableTransactorSession) ClaimHeirOwnership() (*types.Transaction, error) {
	return _Heritable.Contract.ClaimHeirOwnership(&_Heritable.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_Heritable *HeritableTransactor) Heartbeat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heritable.contract.Transact(opts, "heartbeat")
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_Heritable *HeritableSession) Heartbeat() (*types.Transaction, error) {
	return _Heritable.Contract.Heartbeat(&_Heritable.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_Heritable *HeritableTransactorSession) Heartbeat() (*types.Transaction, error) {
	return _Heritable.Contract.Heartbeat(&_Heritable.TransactOpts)
}

// ProclaimDeath is a paid mutator transaction binding the contract method 0x24845131.
//
// Solidity: function proclaimDeath() returns()
func (_Heritable *HeritableTransactor) ProclaimDeath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heritable.contract.Transact(opts, "proclaimDeath")
}

// ProclaimDeath is a paid mutator transaction binding the contract method 0x24845131.
//
// Solidity: function proclaimDeath() returns()
func (_Heritable *HeritableSession) ProclaimDeath() (*types.Transaction, error) {
	return _Heritable.Contract.ProclaimDeath(&_Heritable.TransactOpts)
}

// ProclaimDeath is a paid mutator transaction binding the contract method 0x24845131.
//
// Solidity: function proclaimDeath() returns()
func (_Heritable *HeritableTransactorSession) ProclaimDeath() (*types.Transaction, error) {
	return _Heritable.Contract.ProclaimDeath(&_Heritable.TransactOpts)
}

// RemoveHeir is a paid mutator transaction binding the contract method 0xebf88de4.
//
// Solidity: function removeHeir() returns()
func (_Heritable *HeritableTransactor) RemoveHeir(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heritable.contract.Transact(opts, "removeHeir")
}

// RemoveHeir is a paid mutator transaction binding the contract method 0xebf88de4.
//
// Solidity: function removeHeir() returns()
func (_Heritable *HeritableSession) RemoveHeir() (*types.Transaction, error) {
	return _Heritable.Contract.RemoveHeir(&_Heritable.TransactOpts)
}

// RemoveHeir is a paid mutator transaction binding the contract method 0xebf88de4.
//
// Solidity: function removeHeir() returns()
func (_Heritable *HeritableTransactorSession) RemoveHeir() (*types.Transaction, error) {
	return _Heritable.Contract.RemoveHeir(&_Heritable.TransactOpts)
}

// SetHeir is a paid mutator transaction binding the contract method 0x23defc77.
//
// Solidity: function setHeir(newHeir address) returns()
func (_Heritable *HeritableTransactor) SetHeir(opts *bind.TransactOpts, newHeir common.Address) (*types.Transaction, error) {
	return _Heritable.contract.Transact(opts, "setHeir", newHeir)
}

// SetHeir is a paid mutator transaction binding the contract method 0x23defc77.
//
// Solidity: function setHeir(newHeir address) returns()
func (_Heritable *HeritableSession) SetHeir(newHeir common.Address) (*types.Transaction, error) {
	return _Heritable.Contract.SetHeir(&_Heritable.TransactOpts, newHeir)
}

// SetHeir is a paid mutator transaction binding the contract method 0x23defc77.
//
// Solidity: function setHeir(newHeir address) returns()
func (_Heritable *HeritableTransactorSession) SetHeir(newHeir common.Address) (*types.Transaction, error) {
	return _Heritable.Contract.SetHeir(&_Heritable.TransactOpts, newHeir)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Heritable *HeritableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Heritable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Heritable *HeritableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Heritable.Contract.TransferOwnership(&_Heritable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Heritable *HeritableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Heritable.Contract.TransferOwnership(&_Heritable.TransactOpts, newOwner)
}

// HeritableHeirChangedIterator is returned from FilterHeirChanged and is used to iterate over the raw logs and unpacked data for HeirChanged events raised by the Heritable contract.
type HeritableHeirChangedIterator struct {
	Event *HeritableHeirChanged // Event containing the contract specifics and raw log

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
func (it *HeritableHeirChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeritableHeirChanged)
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
		it.Event = new(HeritableHeirChanged)
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
func (it *HeritableHeirChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeritableHeirChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeritableHeirChanged represents a HeirChanged event raised by the Heritable contract.
type HeritableHeirChanged struct {
	Owner   common.Address
	NewHeir common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHeirChanged is a free log retrieval operation binding the contract event 0x4e6093f85fa64484abd692810d8a44d508792ff7b7a021d9fbd69fa1c6ff18a0.
//
// Solidity: event HeirChanged(owner indexed address, newHeir indexed address)
func (_Heritable *HeritableFilterer) FilterHeirChanged(opts *bind.FilterOpts, owner []common.Address, newHeir []common.Address) (*HeritableHeirChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newHeirRule []interface{}
	for _, newHeirItem := range newHeir {
		newHeirRule = append(newHeirRule, newHeirItem)
	}

	logs, sub, err := _Heritable.contract.FilterLogs(opts, "HeirChanged", ownerRule, newHeirRule)
	if err != nil {
		return nil, err
	}
	return &HeritableHeirChangedIterator{contract: _Heritable.contract, event: "HeirChanged", logs: logs, sub: sub}, nil
}

// WatchHeirChanged is a free log subscription operation binding the contract event 0x4e6093f85fa64484abd692810d8a44d508792ff7b7a021d9fbd69fa1c6ff18a0.
//
// Solidity: event HeirChanged(owner indexed address, newHeir indexed address)
func (_Heritable *HeritableFilterer) WatchHeirChanged(opts *bind.WatchOpts, sink chan<- *HeritableHeirChanged, owner []common.Address, newHeir []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var newHeirRule []interface{}
	for _, newHeirItem := range newHeir {
		newHeirRule = append(newHeirRule, newHeirItem)
	}

	logs, sub, err := _Heritable.contract.WatchLogs(opts, "HeirChanged", ownerRule, newHeirRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeritableHeirChanged)
				if err := _Heritable.contract.UnpackLog(event, "HeirChanged", log); err != nil {
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

// HeritableHeirOwnershipClaimedIterator is returned from FilterHeirOwnershipClaimed and is used to iterate over the raw logs and unpacked data for HeirOwnershipClaimed events raised by the Heritable contract.
type HeritableHeirOwnershipClaimedIterator struct {
	Event *HeritableHeirOwnershipClaimed // Event containing the contract specifics and raw log

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
func (it *HeritableHeirOwnershipClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeritableHeirOwnershipClaimed)
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
		it.Event = new(HeritableHeirOwnershipClaimed)
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
func (it *HeritableHeirOwnershipClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeritableHeirOwnershipClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeritableHeirOwnershipClaimed represents a HeirOwnershipClaimed event raised by the Heritable contract.
type HeritableHeirOwnershipClaimed struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHeirOwnershipClaimed is a free log retrieval operation binding the contract event 0x1017a357e19071e4408dbd385f24e591aa5bcee52b444dc0c8abddbe6ad29de6.
//
// Solidity: event HeirOwnershipClaimed(previousOwner indexed address, newOwner indexed address)
func (_Heritable *HeritableFilterer) FilterHeirOwnershipClaimed(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HeritableHeirOwnershipClaimedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heritable.contract.FilterLogs(opts, "HeirOwnershipClaimed", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HeritableHeirOwnershipClaimedIterator{contract: _Heritable.contract, event: "HeirOwnershipClaimed", logs: logs, sub: sub}, nil
}

// WatchHeirOwnershipClaimed is a free log subscription operation binding the contract event 0x1017a357e19071e4408dbd385f24e591aa5bcee52b444dc0c8abddbe6ad29de6.
//
// Solidity: event HeirOwnershipClaimed(previousOwner indexed address, newOwner indexed address)
func (_Heritable *HeritableFilterer) WatchHeirOwnershipClaimed(opts *bind.WatchOpts, sink chan<- *HeritableHeirOwnershipClaimed, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heritable.contract.WatchLogs(opts, "HeirOwnershipClaimed", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeritableHeirOwnershipClaimed)
				if err := _Heritable.contract.UnpackLog(event, "HeirOwnershipClaimed", log); err != nil {
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

// HeritableOwnerHeartbeatedIterator is returned from FilterOwnerHeartbeated and is used to iterate over the raw logs and unpacked data for OwnerHeartbeated events raised by the Heritable contract.
type HeritableOwnerHeartbeatedIterator struct {
	Event *HeritableOwnerHeartbeated // Event containing the contract specifics and raw log

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
func (it *HeritableOwnerHeartbeatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeritableOwnerHeartbeated)
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
		it.Event = new(HeritableOwnerHeartbeated)
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
func (it *HeritableOwnerHeartbeatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeritableOwnerHeartbeatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeritableOwnerHeartbeated represents a OwnerHeartbeated event raised by the Heritable contract.
type HeritableOwnerHeartbeated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerHeartbeated is a free log retrieval operation binding the contract event 0xd40b9d9a7572411aff4bbb95ba7bd9a9d4e9b70747ec46f8fe7913c309b68452.
//
// Solidity: event OwnerHeartbeated(owner indexed address)
func (_Heritable *HeritableFilterer) FilterOwnerHeartbeated(opts *bind.FilterOpts, owner []common.Address) (*HeritableOwnerHeartbeatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Heritable.contract.FilterLogs(opts, "OwnerHeartbeated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &HeritableOwnerHeartbeatedIterator{contract: _Heritable.contract, event: "OwnerHeartbeated", logs: logs, sub: sub}, nil
}

// WatchOwnerHeartbeated is a free log subscription operation binding the contract event 0xd40b9d9a7572411aff4bbb95ba7bd9a9d4e9b70747ec46f8fe7913c309b68452.
//
// Solidity: event OwnerHeartbeated(owner indexed address)
func (_Heritable *HeritableFilterer) WatchOwnerHeartbeated(opts *bind.WatchOpts, sink chan<- *HeritableOwnerHeartbeated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Heritable.contract.WatchLogs(opts, "OwnerHeartbeated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeritableOwnerHeartbeated)
				if err := _Heritable.contract.UnpackLog(event, "OwnerHeartbeated", log); err != nil {
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

// HeritableOwnerProclaimedDeadIterator is returned from FilterOwnerProclaimedDead and is used to iterate over the raw logs and unpacked data for OwnerProclaimedDead events raised by the Heritable contract.
type HeritableOwnerProclaimedDeadIterator struct {
	Event *HeritableOwnerProclaimedDead // Event containing the contract specifics and raw log

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
func (it *HeritableOwnerProclaimedDeadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeritableOwnerProclaimedDead)
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
		it.Event = new(HeritableOwnerProclaimedDead)
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
func (it *HeritableOwnerProclaimedDeadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeritableOwnerProclaimedDeadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeritableOwnerProclaimedDead represents a OwnerProclaimedDead event raised by the Heritable contract.
type HeritableOwnerProclaimedDead struct {
	Owner       common.Address
	Heir        common.Address
	TimeOfDeath *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOwnerProclaimedDead is a free log retrieval operation binding the contract event 0x66389f1f3251c39e71cb18351daefea45bda23e98df0767266b8aed4ff938277.
//
// Solidity: event OwnerProclaimedDead(owner indexed address, heir indexed address, timeOfDeath uint256)
func (_Heritable *HeritableFilterer) FilterOwnerProclaimedDead(opts *bind.FilterOpts, owner []common.Address, heir []common.Address) (*HeritableOwnerProclaimedDeadIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var heirRule []interface{}
	for _, heirItem := range heir {
		heirRule = append(heirRule, heirItem)
	}

	logs, sub, err := _Heritable.contract.FilterLogs(opts, "OwnerProclaimedDead", ownerRule, heirRule)
	if err != nil {
		return nil, err
	}
	return &HeritableOwnerProclaimedDeadIterator{contract: _Heritable.contract, event: "OwnerProclaimedDead", logs: logs, sub: sub}, nil
}

// WatchOwnerProclaimedDead is a free log subscription operation binding the contract event 0x66389f1f3251c39e71cb18351daefea45bda23e98df0767266b8aed4ff938277.
//
// Solidity: event OwnerProclaimedDead(owner indexed address, heir indexed address, timeOfDeath uint256)
func (_Heritable *HeritableFilterer) WatchOwnerProclaimedDead(opts *bind.WatchOpts, sink chan<- *HeritableOwnerProclaimedDead, owner []common.Address, heir []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var heirRule []interface{}
	for _, heirItem := range heir {
		heirRule = append(heirRule, heirItem)
	}

	logs, sub, err := _Heritable.contract.WatchLogs(opts, "OwnerProclaimedDead", ownerRule, heirRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeritableOwnerProclaimedDead)
				if err := _Heritable.contract.UnpackLog(event, "OwnerProclaimedDead", log); err != nil {
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

// HeritableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Heritable contract.
type HeritableOwnershipTransferredIterator struct {
	Event *HeritableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HeritableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeritableOwnershipTransferred)
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
		it.Event = new(HeritableOwnershipTransferred)
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
func (it *HeritableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeritableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeritableOwnershipTransferred represents a OwnershipTransferred event raised by the Heritable contract.
type HeritableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Heritable *HeritableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HeritableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heritable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HeritableOwnershipTransferredIterator{contract: _Heritable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Heritable *HeritableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HeritableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heritable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeritableOwnershipTransferred)
				if err := _Heritable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
