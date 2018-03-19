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

// SafeERC20HelperABI is the input ABI used to generate the binding from.
const SafeERC20HelperABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"doFailingTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doSucceedingTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doFailingApprove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doFailingTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doSucceedingApprove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doSucceedingTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SafeERC20HelperBin is the compiled bytecode used for deploying new contracts.
const SafeERC20HelperBin = `6060604052341561000f57600080fd5b610017610090565b604051809103906000f080151561002d57600080fd5b60008054600160a060020a031916600160a060020a03929092169190911790556100556100a0565b604051809103906000f080151561006b57600080fd5b60018054600160a060020a031916600160a060020a03929092169190911790556100b0565b6040516101d9806103d283390190565b6040516101e5806105ab83390190565b610313806100bf6000396000f30060606040526004361061005e5763ffffffff60e060020a60003504166350930ca5811461006357806356b2c5bb146100785780636573ee1e1461008b578063afa03ac91461009e578063b19faed8146100b1578063d7228bb5146100c4575b600080fd5b341561006e57600080fd5b6100766100d7565b005b341561008357600080fd5b6100766100fa565b341561009657600080fd5b610076610118565b34156100a957600080fd5b610076610138565b34156100bc57600080fd5b610076610158565b34156100cf57600080fd5b610076610176565b600080546100f891600160a060020a0390911690808063ffffffff61019516565b565b6001546100f890600160a060020a031660008063ffffffff61021b16565b600080546100f891600160a060020a03909116908063ffffffff61029316565b600080546100f891600160a060020a03909116908063ffffffff61021b16565b6001546100f890600160a060020a031660008063ffffffff61029316565b6001546100f890600160a060020a03166000808063ffffffff61019516565b83600160a060020a03166323b872dd84848460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156101f657600080fd5b5af1151561020357600080fd5b50505060405180519050151561021557fe5b50505050565b82600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561026f57600080fd5b5af1151561027c57600080fd5b50505060405180519050151561028e57fe5b505050565b82600160a060020a031663095ea7b3838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561026f57600080fd00a165627a7a72305820c45b898b5e406a9347eb9948f8374649c07377068e002f316fd265540cd2b1c900296060604052341561000f57600080fd5b6101bb8061001e6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100bf57806323b872dd146100e457806370a0823114610119578063a9059cbb1461007c578063dd62ed3e14610145575b600080fd5b341561008757600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043516602435610173565b604051901515815260200160405180910390f35b34156100ca57600080fd5b6100d261017b565b60405190815260200160405180910390f35b34156100ef57600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610180565b341561012457600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff60043516610189565b341561015057600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff600435811690602435165b600092915050565b600090565b60009392505050565b506000905600a165627a7a7230582053cacf9429ed406a357081ee312217ed85c81cfc3abc13f52510a4930c8afddb00296060604052341561000f57600080fd5b6101c78061001e6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100bf57806323b872dd146100e457806370a0823114610119578063a9059cbb1461007c578063dd62ed3e14610145575b600080fd5b341561008757600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043516602435610177565b604051901515815260200160405180910390f35b34156100ca57600080fd5b6100d261017f565b60405190815260200160405180910390f35b34156100ef57600080fd5b6100ab73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610184565b341561012457600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff6004351661018d565b341561015057600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff60043581169060243516610193565b600192915050565b600090565b60019392505050565b50600090565b6000929150505600a165627a7a72305820ca5a3134adce02c5382c754f4840d02a976c0d1e17762e52dbf7ee44ca8bf40b0029`

// DeploySafeERC20Helper deploys a new Ethereum contract, binding an instance of SafeERC20Helper to it.
func DeploySafeERC20Helper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20Helper, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20HelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20HelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20Helper{SafeERC20HelperCaller: SafeERC20HelperCaller{contract: contract}, SafeERC20HelperTransactor: SafeERC20HelperTransactor{contract: contract}, SafeERC20HelperFilterer: SafeERC20HelperFilterer{contract: contract}}, nil
}

// SafeERC20Helper is an auto generated Go binding around an Ethereum contract.
type SafeERC20Helper struct {
	SafeERC20HelperCaller     // Read-only binding to the contract
	SafeERC20HelperTransactor // Write-only binding to the contract
	SafeERC20HelperFilterer   // Log filterer for contract events
}

// SafeERC20HelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20HelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20HelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20HelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20HelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20HelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20HelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20HelperSession struct {
	Contract     *SafeERC20Helper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20HelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20HelperCallerSession struct {
	Contract *SafeERC20HelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SafeERC20HelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20HelperTransactorSession struct {
	Contract     *SafeERC20HelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SafeERC20HelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20HelperRaw struct {
	Contract *SafeERC20Helper // Generic contract binding to access the raw methods on
}

// SafeERC20HelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20HelperCallerRaw struct {
	Contract *SafeERC20HelperCaller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20HelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20HelperTransactorRaw struct {
	Contract *SafeERC20HelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20Helper creates a new instance of SafeERC20Helper, bound to a specific deployed contract.
func NewSafeERC20Helper(address common.Address, backend bind.ContractBackend) (*SafeERC20Helper, error) {
	contract, err := bindSafeERC20Helper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Helper{SafeERC20HelperCaller: SafeERC20HelperCaller{contract: contract}, SafeERC20HelperTransactor: SafeERC20HelperTransactor{contract: contract}, SafeERC20HelperFilterer: SafeERC20HelperFilterer{contract: contract}}, nil
}

// NewSafeERC20HelperCaller creates a new read-only instance of SafeERC20Helper, bound to a specific deployed contract.
func NewSafeERC20HelperCaller(address common.Address, caller bind.ContractCaller) (*SafeERC20HelperCaller, error) {
	contract, err := bindSafeERC20Helper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20HelperCaller{contract: contract}, nil
}

// NewSafeERC20HelperTransactor creates a new write-only instance of SafeERC20Helper, bound to a specific deployed contract.
func NewSafeERC20HelperTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20HelperTransactor, error) {
	contract, err := bindSafeERC20Helper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20HelperTransactor{contract: contract}, nil
}

// NewSafeERC20HelperFilterer creates a new log filterer instance of SafeERC20Helper, bound to a specific deployed contract.
func NewSafeERC20HelperFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20HelperFilterer, error) {
	contract, err := bindSafeERC20Helper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20HelperFilterer{contract: contract}, nil
}

// bindSafeERC20Helper binds a generic wrapper to an already deployed contract.
func bindSafeERC20Helper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20HelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20Helper *SafeERC20HelperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20Helper.Contract.SafeERC20HelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20Helper *SafeERC20HelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.SafeERC20HelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20Helper *SafeERC20HelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.SafeERC20HelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20Helper *SafeERC20HelperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20Helper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20Helper *SafeERC20HelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20Helper *SafeERC20HelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.contract.Transact(opts, method, params...)
}

// DoFailingApprove is a paid mutator transaction binding the contract method 0x6573ee1e.
//
// Solidity: function doFailingApprove() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactor) DoFailingApprove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.contract.Transact(opts, "doFailingApprove")
}

// DoFailingApprove is a paid mutator transaction binding the contract method 0x6573ee1e.
//
// Solidity: function doFailingApprove() returns()
func (_SafeERC20Helper *SafeERC20HelperSession) DoFailingApprove() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoFailingApprove(&_SafeERC20Helper.TransactOpts)
}

// DoFailingApprove is a paid mutator transaction binding the contract method 0x6573ee1e.
//
// Solidity: function doFailingApprove() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactorSession) DoFailingApprove() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoFailingApprove(&_SafeERC20Helper.TransactOpts)
}

// DoFailingTransfer is a paid mutator transaction binding the contract method 0xafa03ac9.
//
// Solidity: function doFailingTransfer() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactor) DoFailingTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.contract.Transact(opts, "doFailingTransfer")
}

// DoFailingTransfer is a paid mutator transaction binding the contract method 0xafa03ac9.
//
// Solidity: function doFailingTransfer() returns()
func (_SafeERC20Helper *SafeERC20HelperSession) DoFailingTransfer() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoFailingTransfer(&_SafeERC20Helper.TransactOpts)
}

// DoFailingTransfer is a paid mutator transaction binding the contract method 0xafa03ac9.
//
// Solidity: function doFailingTransfer() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactorSession) DoFailingTransfer() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoFailingTransfer(&_SafeERC20Helper.TransactOpts)
}

// DoFailingTransferFrom is a paid mutator transaction binding the contract method 0x50930ca5.
//
// Solidity: function doFailingTransferFrom() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactor) DoFailingTransferFrom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.contract.Transact(opts, "doFailingTransferFrom")
}

// DoFailingTransferFrom is a paid mutator transaction binding the contract method 0x50930ca5.
//
// Solidity: function doFailingTransferFrom() returns()
func (_SafeERC20Helper *SafeERC20HelperSession) DoFailingTransferFrom() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoFailingTransferFrom(&_SafeERC20Helper.TransactOpts)
}

// DoFailingTransferFrom is a paid mutator transaction binding the contract method 0x50930ca5.
//
// Solidity: function doFailingTransferFrom() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactorSession) DoFailingTransferFrom() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoFailingTransferFrom(&_SafeERC20Helper.TransactOpts)
}

// DoSucceedingApprove is a paid mutator transaction binding the contract method 0xb19faed8.
//
// Solidity: function doSucceedingApprove() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactor) DoSucceedingApprove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.contract.Transact(opts, "doSucceedingApprove")
}

// DoSucceedingApprove is a paid mutator transaction binding the contract method 0xb19faed8.
//
// Solidity: function doSucceedingApprove() returns()
func (_SafeERC20Helper *SafeERC20HelperSession) DoSucceedingApprove() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoSucceedingApprove(&_SafeERC20Helper.TransactOpts)
}

// DoSucceedingApprove is a paid mutator transaction binding the contract method 0xb19faed8.
//
// Solidity: function doSucceedingApprove() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactorSession) DoSucceedingApprove() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoSucceedingApprove(&_SafeERC20Helper.TransactOpts)
}

// DoSucceedingTransfer is a paid mutator transaction binding the contract method 0x56b2c5bb.
//
// Solidity: function doSucceedingTransfer() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactor) DoSucceedingTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.contract.Transact(opts, "doSucceedingTransfer")
}

// DoSucceedingTransfer is a paid mutator transaction binding the contract method 0x56b2c5bb.
//
// Solidity: function doSucceedingTransfer() returns()
func (_SafeERC20Helper *SafeERC20HelperSession) DoSucceedingTransfer() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoSucceedingTransfer(&_SafeERC20Helper.TransactOpts)
}

// DoSucceedingTransfer is a paid mutator transaction binding the contract method 0x56b2c5bb.
//
// Solidity: function doSucceedingTransfer() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactorSession) DoSucceedingTransfer() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoSucceedingTransfer(&_SafeERC20Helper.TransactOpts)
}

// DoSucceedingTransferFrom is a paid mutator transaction binding the contract method 0xd7228bb5.
//
// Solidity: function doSucceedingTransferFrom() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactor) DoSucceedingTransferFrom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Helper.contract.Transact(opts, "doSucceedingTransferFrom")
}

// DoSucceedingTransferFrom is a paid mutator transaction binding the contract method 0xd7228bb5.
//
// Solidity: function doSucceedingTransferFrom() returns()
func (_SafeERC20Helper *SafeERC20HelperSession) DoSucceedingTransferFrom() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoSucceedingTransferFrom(&_SafeERC20Helper.TransactOpts)
}

// DoSucceedingTransferFrom is a paid mutator transaction binding the contract method 0xd7228bb5.
//
// Solidity: function doSucceedingTransferFrom() returns()
func (_SafeERC20Helper *SafeERC20HelperTransactorSession) DoSucceedingTransferFrom() (*types.Transaction, error) {
	return _SafeERC20Helper.Contract.DoSucceedingTransferFrom(&_SafeERC20Helper.TransactOpts)
}
