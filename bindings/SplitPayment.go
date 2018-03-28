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

// SplitPaymentABI is the input ABI used to generate the binding from.
const SplitPaymentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payees\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_payees\",\"type\":\"address[]\"},{\"name\":\"_shares\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// SplitPaymentBin is the compiled bytecode used for deploying new contracts.
const SplitPaymentBin = `60606040526000805560006001556040516106253803806106258339810160405280805182019190602001805190910190506000815183511461004157600080fd5b5060005b82518110156100975761008f83828151811061005d57fe5b9060200190602002015183838151811061007357fe5b9060200190602002015164010000000061034f61009f82021704565b600101610045565b5050506101ad565b600160a060020a03821615156100b457600080fd5b600081116100c157600080fd5b600160a060020a038216600090815260026020526040902054156100e457600080fd5b60048054600181016100f68382610163565b5060009182526020808320919091018054600160a060020a031916600160a060020a0386169081179091558252600290526040812082905554610146908264010000000061014d81026102e11704565b6000555050565b60008282018381101561015c57fe5b9392505050565b8154818355818115116101875760008381526020902061018791810190830161018c565b505050565b6101aa91905b808211156101a65760008155600101610192565b5090565b90565b610469806101bc6000396000f3006060604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633a98ef3981146100795780634e71d92d1461009e57806363037b0c146100b15780639852595c146100e3578063ce7c2ac214610102578063e33b7de314610121575b005b341561008457600080fd5b61008c610134565b60405190815260200160405180910390f35b34156100a957600080fd5b61007761013a565b34156100bc57600080fd5b6100c760043561028f565b604051600160a060020a03909116815260200160405180910390f35b34156100ee57600080fd5b61008c600160a060020a03600435166102b7565b341561010d57600080fd5b61008c600160a060020a03600435166102c9565b341561012c57600080fd5b61008c6102db565b60005481565b33600160a060020a038116600090815260026020526040812054819081901161016257600080fd5b60015461018090600160a060020a033016319063ffffffff6102e116565b600160a060020a038416600090815260036020908152604080832054835460029093529220549294506101db926101cf91906101c390879063ffffffff6102fb16565b9063ffffffff61032616565b9063ffffffff61033d16565b90508015156101e957600080fd5b600160a060020a033016318190101561020157600080fd5b600160a060020a03831660009081526003602052604090205461022a908263ffffffff6102e116565b600160a060020a038416600090815260036020526040902055600154610256908263ffffffff6102e116565b600155600160a060020a03831681156108fc0282604051600060405180830381858888f19350505050151561028a57600080fd5b505050565b600480548290811061029d57fe5b600091825260209091200154600160a060020a0316905081565b60036020526000908152604090205481565b60026020526000908152604090205481565b60015481565b6000828201838110156102f057fe5b8091505b5092915050565b60008083151561030e57600091506102f4565b5082820282848281151561031e57fe5b04146102f057fe5b600080828481151561033457fe5b04949350505050565b60008282111561034957fe5b50900390565b600160a060020a038216151561036457600080fd5b6000811161037157600080fd5b600160a060020a0382166000908152600260205260409020541561039457600080fd5b60048054600181016103a683826103fd565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915582526002905260408120829055546103f690826102e1565b6000555050565b81548183558181151161028a5760008381526020902061028a91810190830161043a91905b808211156104365760008155600101610422565b5090565b905600a165627a7a7230582016dbf138f2f9b9530f1ffcf333adc5d02f7c88abd8d10ebd5cf103e82538f7c20029`

// DeploySplitPayment deploys a new Ethereum contract, binding an instance of SplitPayment to it.
func DeploySplitPayment(auth *bind.TransactOpts, backend bind.ContractBackend, _payees []common.Address, _shares []*big.Int) (common.Address, *types.Transaction, *SplitPayment, error) {
	parsed, err := abi.JSON(strings.NewReader(SplitPaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SplitPaymentBin), backend, _payees, _shares)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SplitPayment{SplitPaymentCaller: SplitPaymentCaller{contract: contract}, SplitPaymentTransactor: SplitPaymentTransactor{contract: contract}}, nil
}

// SplitPayment is an auto generated Go binding around an Ethereum contract.
type SplitPayment struct {
	SplitPaymentCaller     // Read-only binding to the contract
	SplitPaymentTransactor // Write-only binding to the contract
}

// SplitPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type SplitPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SplitPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SplitPaymentSession struct {
	Contract     *SplitPayment     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SplitPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SplitPaymentCallerSession struct {
	Contract *SplitPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SplitPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SplitPaymentTransactorSession struct {
	Contract     *SplitPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SplitPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type SplitPaymentRaw struct {
	Contract *SplitPayment // Generic contract binding to access the raw methods on
}

// SplitPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SplitPaymentCallerRaw struct {
	Contract *SplitPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// SplitPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SplitPaymentTransactorRaw struct {
	Contract *SplitPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSplitPayment creates a new instance of SplitPayment, bound to a specific deployed contract.
func NewSplitPayment(address common.Address, backend bind.ContractBackend) (*SplitPayment, error) {
	contract, err := bindSplitPayment(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SplitPayment{SplitPaymentCaller: SplitPaymentCaller{contract: contract}, SplitPaymentTransactor: SplitPaymentTransactor{contract: contract}}, nil
}

// NewSplitPaymentCaller creates a new read-only instance of SplitPayment, bound to a specific deployed contract.
func NewSplitPaymentCaller(address common.Address, caller bind.ContractCaller) (*SplitPaymentCaller, error) {
	contract, err := bindSplitPayment(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SplitPaymentCaller{contract: contract}, nil
}

// NewSplitPaymentTransactor creates a new write-only instance of SplitPayment, bound to a specific deployed contract.
func NewSplitPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*SplitPaymentTransactor, error) {
	contract, err := bindSplitPayment(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SplitPaymentTransactor{contract: contract}, nil
}

// bindSplitPayment binds a generic wrapper to an already deployed contract.
func bindSplitPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SplitPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SplitPayment *SplitPaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SplitPayment.Contract.SplitPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SplitPayment *SplitPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitPayment.Contract.SplitPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SplitPayment *SplitPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SplitPayment.Contract.SplitPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SplitPayment *SplitPaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SplitPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SplitPayment *SplitPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SplitPayment *SplitPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SplitPayment.Contract.contract.Transact(opts, method, params...)
}

// Payees is a free data retrieval call binding the contract method 0x63037b0c.
//
// Solidity: function payees( uint256) constant returns(address)
func (_SplitPayment *SplitPaymentCaller) Payees(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SplitPayment.contract.Call(opts, out, "payees", arg0)
	return *ret0, err
}

// Payees is a free data retrieval call binding the contract method 0x63037b0c.
//
// Solidity: function payees( uint256) constant returns(address)
func (_SplitPayment *SplitPaymentSession) Payees(arg0 *big.Int) (common.Address, error) {
	return _SplitPayment.Contract.Payees(&_SplitPayment.CallOpts, arg0)
}

// Payees is a free data retrieval call binding the contract method 0x63037b0c.
//
// Solidity: function payees( uint256) constant returns(address)
func (_SplitPayment *SplitPaymentCallerSession) Payees(arg0 *big.Int) (common.Address, error) {
	return _SplitPayment.Contract.Payees(&_SplitPayment.CallOpts, arg0)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_SplitPayment *SplitPaymentCaller) Released(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SplitPayment.contract.Call(opts, out, "released", arg0)
	return *ret0, err
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_SplitPayment *SplitPaymentSession) Released(arg0 common.Address) (*big.Int, error) {
	return _SplitPayment.Contract.Released(&_SplitPayment.CallOpts, arg0)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_SplitPayment *SplitPaymentCallerSession) Released(arg0 common.Address) (*big.Int, error) {
	return _SplitPayment.Contract.Released(&_SplitPayment.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares( address) constant returns(uint256)
func (_SplitPayment *SplitPaymentCaller) Shares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SplitPayment.contract.Call(opts, out, "shares", arg0)
	return *ret0, err
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares( address) constant returns(uint256)
func (_SplitPayment *SplitPaymentSession) Shares(arg0 common.Address) (*big.Int, error) {
	return _SplitPayment.Contract.Shares(&_SplitPayment.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares( address) constant returns(uint256)
func (_SplitPayment *SplitPaymentCallerSession) Shares(arg0 common.Address) (*big.Int, error) {
	return _SplitPayment.Contract.Shares(&_SplitPayment.CallOpts, arg0)
}

// TotalReleased is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() constant returns(uint256)
func (_SplitPayment *SplitPaymentCaller) TotalReleased(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SplitPayment.contract.Call(opts, out, "totalReleased")
	return *ret0, err
}

// TotalReleased is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() constant returns(uint256)
func (_SplitPayment *SplitPaymentSession) TotalReleased() (*big.Int, error) {
	return _SplitPayment.Contract.TotalReleased(&_SplitPayment.CallOpts)
}

// TotalReleased is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() constant returns(uint256)
func (_SplitPayment *SplitPaymentCallerSession) TotalReleased() (*big.Int, error) {
	return _SplitPayment.Contract.TotalReleased(&_SplitPayment.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() constant returns(uint256)
func (_SplitPayment *SplitPaymentCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SplitPayment.contract.Call(opts, out, "totalShares")
	return *ret0, err
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() constant returns(uint256)
func (_SplitPayment *SplitPaymentSession) TotalShares() (*big.Int, error) {
	return _SplitPayment.Contract.TotalShares(&_SplitPayment.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() constant returns(uint256)
func (_SplitPayment *SplitPaymentCallerSession) TotalShares() (*big.Int, error) {
	return _SplitPayment.Contract.TotalShares(&_SplitPayment.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_SplitPayment *SplitPaymentTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitPayment.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_SplitPayment *SplitPaymentSession) Claim() (*types.Transaction, error) {
	return _SplitPayment.Contract.Claim(&_SplitPayment.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_SplitPayment *SplitPaymentTransactorSession) Claim() (*types.Transaction, error) {
	return _SplitPayment.Contract.Claim(&_SplitPayment.TransactOpts)
}
