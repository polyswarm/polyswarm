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

// FinalizableCrowdsaleImplABI is the input ABI used to generate the binding from.
const FinalizableCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// FinalizableCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const FinalizableCrowdsaleImplBin = `60606040526006805460a060020a60ff0219169055341561001f57600080fd5b60405160a08061072383398101604052808051919060200180519190602001805191906020018051919060200180519150859050848484846000831161006457600080fd5b600160a060020a038216151561007957600080fd5b600160a060020a038116151561008e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a03199182161790915560008054929093169116179055428210156100cd57600080fd5b818110156100da57600080fd5b600491909155600555505060068054600160a060020a033316600160a060020a031990911617905550505061060f806101146000396000f3006060604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631515bc2b81146100c45780632c4e722e146100eb5780634042b66f146101105780634b6753bc146101235780634bb278f314610136578063521eb273146101495780638d4e4083146101785780638da5cb5b1461018b578063b7a8807c1461019e578063ec8ac4d8146101b1578063f2fde38b146101c5578063fc0c546a146101e4575b6100c2336101f7565b005b34156100cf57600080fd5b6100d761029f565b604051901515815260200160405180910390f35b34156100f657600080fd5b6100fe6102a7565b60405190815260200160405180910390f35b341561011b57600080fd5b6100fe6102ad565b341561012e57600080fd5b6100fe6102b3565b341561014157600080fd5b6100c26102b9565b341561015457600080fd5b61015c61037a565b604051600160a060020a03909116815260200160405180910390f35b341561018357600080fd5b6100d7610389565b341561019657600080fd5b61015c6103aa565b34156101a957600080fd5b6100fe6103b9565b6100c2600160a060020a03600435166101f7565b34156101d057600080fd5b6100c2600160a060020a03600435166103bf565b34156101ef57600080fd5b61015c61045a565b3460006102048383610469565b61020d82610496565b600354909150610223908363ffffffff6104b316565b60035561023083826104cd565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36102888383610492565b6102906104d7565b61029a8383610492565b505050565b600554421190565b60025481565b60035481565b60055481565b60065433600160a060020a039081169116146102d457600080fd5b60065474010000000000000000000000000000000000000000900460ff16156102fc57600080fd5b61030461029f565b151561030f57600080fd5b61031761050b565b7f6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b1768160405160405180910390a16006805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600154600160a060020a031681565b60065474010000000000000000000000000000000000000000900460ff1681565b600654600160a060020a031681565b60045481565b60065433600160a060020a039081169116146103da57600080fd5b600160a060020a03811615156103ef57600080fd5b600654600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600454421015801561047d57506005544211155b151561048857600080fd5b610492828261050d565b5050565b60006104ad6002548361052e90919063ffffffff16565b92915050565b6000828201838110156104c257fe5b8091505b5092915050565b6104928282610559565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f19350505050151561050b57600080fd5b565b600160a060020a038216151561052257600080fd5b80151561049257600080fd5b60008083151561054157600091506104c6565b5082820282848281151561055157fe5b04146104c257fe5b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156105c857600080fd5b5af115156105d557600080fd5b5050506040518051505050505600a165627a7a723058203e0e02a63db1fff7d0d9f04e016e56b5a6200bfdba8af414e97a024bf992cef30029`

// DeployFinalizableCrowdsaleImpl deploys a new Ethereum contract, binding an instance of FinalizableCrowdsaleImpl to it.
func DeployFinalizableCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _openingTime *big.Int, _closingTime *big.Int, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *FinalizableCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FinalizableCrowdsaleImplBin), backend, _openingTime, _closingTime, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalizableCrowdsaleImpl{FinalizableCrowdsaleImplCaller: FinalizableCrowdsaleImplCaller{contract: contract}, FinalizableCrowdsaleImplTransactor: FinalizableCrowdsaleImplTransactor{contract: contract}, FinalizableCrowdsaleImplFilterer: FinalizableCrowdsaleImplFilterer{contract: contract}}, nil
}

// FinalizableCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type FinalizableCrowdsaleImpl struct {
	FinalizableCrowdsaleImplCaller     // Read-only binding to the contract
	FinalizableCrowdsaleImplTransactor // Write-only binding to the contract
	FinalizableCrowdsaleImplFilterer   // Log filterer for contract events
}

// FinalizableCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizableCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizableCrowdsaleImplSession struct {
	Contract     *FinalizableCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizableCrowdsaleImplCallerSession struct {
	Contract *FinalizableCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// FinalizableCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizableCrowdsaleImplTransactorSession struct {
	Contract     *FinalizableCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplRaw struct {
	Contract *FinalizableCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// FinalizableCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplCallerRaw struct {
	Contract *FinalizableCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizableCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplTransactorRaw struct {
	Contract *FinalizableCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizableCrowdsaleImpl creates a new instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*FinalizableCrowdsaleImpl, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImpl{FinalizableCrowdsaleImplCaller: FinalizableCrowdsaleImplCaller{contract: contract}, FinalizableCrowdsaleImplTransactor: FinalizableCrowdsaleImplTransactor{contract: contract}, FinalizableCrowdsaleImplFilterer: FinalizableCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewFinalizableCrowdsaleImplCaller creates a new read-only instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*FinalizableCrowdsaleImplCaller, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplCaller{contract: contract}, nil
}

// NewFinalizableCrowdsaleImplTransactor creates a new write-only instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizableCrowdsaleImplTransactor, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplTransactor{contract: contract}, nil
}

// NewFinalizableCrowdsaleImplFilterer creates a new log filterer instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizableCrowdsaleImplFilterer, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplFilterer{contract: contract}, nil
}

// bindFinalizableCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindFinalizableCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsaleImpl.Contract.FinalizableCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.FinalizableCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.FinalizableCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) ClosingTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.ClosingTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) ClosingTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.ClosingTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) HasClosed() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.HasClosed(&_FinalizableCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) HasClosed() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.HasClosed(&_FinalizableCrowdsaleImpl.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) IsFinalized() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.IsFinalized(&_FinalizableCrowdsaleImpl.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) IsFinalized() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.IsFinalized(&_FinalizableCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) OpeningTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.OpeningTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) OpeningTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.OpeningTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Owner() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Owner(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Owner() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Owner(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.Rate(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.Rate(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Token() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Token(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Token(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Wallet(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Wallet(&_FinalizableCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.WeiRaised(&_FinalizableCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.WeiRaised(&_FinalizableCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.BuyTokens(&_FinalizableCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.BuyTokens(&_FinalizableCrowdsaleImpl.TransactOpts, _beneficiary)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.Finalize(&_FinalizableCrowdsaleImpl.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.Finalize(&_FinalizableCrowdsaleImpl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.TransferOwnership(&_FinalizableCrowdsaleImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.TransferOwnership(&_FinalizableCrowdsaleImpl.TransactOpts, newOwner)
}

// FinalizableCrowdsaleImplFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplFinalizedIterator struct {
	Event *FinalizableCrowdsaleImplFinalized // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleImplFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleImplFinalized)
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
		it.Event = new(FinalizableCrowdsaleImplFinalized)
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
func (it *FinalizableCrowdsaleImplFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleImplFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleImplFinalized represents a Finalized event raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) FilterFinalized(opts *bind.FilterOpts) (*FinalizableCrowdsaleImplFinalizedIterator, error) {

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplFinalizedIterator{contract: _FinalizableCrowdsaleImpl.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleImplFinalized) (event.Subscription, error) {

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleImplFinalized)
				if err := _FinalizableCrowdsaleImpl.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// FinalizableCrowdsaleImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplOwnershipTransferredIterator struct {
	Event *FinalizableCrowdsaleImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleImplOwnershipTransferred)
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
		it.Event = new(FinalizableCrowdsaleImplOwnershipTransferred)
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
func (it *FinalizableCrowdsaleImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleImplOwnershipTransferred represents a OwnershipTransferred event raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinalizableCrowdsaleImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplOwnershipTransferredIterator{contract: _FinalizableCrowdsaleImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleImplOwnershipTransferred)
				if err := _FinalizableCrowdsaleImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// FinalizableCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplTokenPurchaseIterator struct {
	Event *FinalizableCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleImplTokenPurchase)
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
		it.Event = new(FinalizableCrowdsaleImplTokenPurchase)
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
func (it *FinalizableCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*FinalizableCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplTokenPurchaseIterator{contract: _FinalizableCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleImplTokenPurchase)
				if err := _FinalizableCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
