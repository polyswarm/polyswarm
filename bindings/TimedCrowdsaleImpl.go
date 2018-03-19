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

// TimedCrowdsaleImplABI is the input ABI used to generate the binding from.
const TimedCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// TimedCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const TimedCrowdsaleImplBin = `6060604052341561000f57600080fd5b60405160a0806104e683398101604052808051919060200180519190602001805191906020018051919060200180519150859050848484846000831161005457600080fd5b600160a060020a038216151561006957600080fd5b600160a060020a038116151561007e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a03199182161790915560008054929093169116179055428210156100bd57600080fd5b818110156100ca57600080fd5b60049190915560055550505050506103ff806100e76000396000f30060606040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631515bc2b81146100985780632c4e722e146100bf5780634042b66f146100e45780634b6753bc146100f7578063521eb2731461010a578063b7a8807c14610139578063ec8ac4d81461014c578063fc0c546a14610160575b61009633610173565b005b34156100a357600080fd5b6100ab61021b565b604051901515815260200160405180910390f35b34156100ca57600080fd5b6100d2610223565b60405190815260200160405180910390f35b34156100ef57600080fd5b6100d2610229565b341561010257600080fd5b6100d261022f565b341561011557600080fd5b61011d610235565b604051600160a060020a03909116815260200160405180910390f35b341561014457600080fd5b6100d2610244565b610096600160a060020a0360043516610173565b341561016b57600080fd5b61011d61024a565b3460006101808383610259565b61018982610286565b60035490915061019f908363ffffffff6102a316565b6003556101ac83826102bd565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36102048383610282565b61020c6102c7565b6102168383610282565b505050565b600554421190565b60025481565b60035481565b60055481565b600154600160a060020a031681565b60045481565b600054600160a060020a031681565b600454421015801561026d57506005544211155b151561027857600080fd5b61028282826102fd565b5050565b600061029d6002548361031e90919063ffffffff16565b92915050565b6000828201838110156102b257fe5b8091505b5092915050565b6102828282610349565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f1935050505015156102fb57600080fd5b565b600160a060020a038216151561031257600080fd5b80151561028257600080fd5b60008083151561033157600091506102b6565b5082820282848281151561034157fe5b04146102b257fe5b600054600160a060020a031663a9059cbb83836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156103b857600080fd5b5af115156103c557600080fd5b5050506040518051505050505600a165627a7a723058209069c54eae74bc82a3427942299bdbd83235f1dae84b870827d5b4ad9ed0d7ff0029`

// DeployTimedCrowdsaleImpl deploys a new Ethereum contract, binding an instance of TimedCrowdsaleImpl to it.
func DeployTimedCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _openingTime *big.Int, _closingTime *big.Int, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *TimedCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TimedCrowdsaleImplBin), backend, _openingTime, _closingTime, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimedCrowdsaleImpl{TimedCrowdsaleImplCaller: TimedCrowdsaleImplCaller{contract: contract}, TimedCrowdsaleImplTransactor: TimedCrowdsaleImplTransactor{contract: contract}, TimedCrowdsaleImplFilterer: TimedCrowdsaleImplFilterer{contract: contract}}, nil
}

// TimedCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type TimedCrowdsaleImpl struct {
	TimedCrowdsaleImplCaller     // Read-only binding to the contract
	TimedCrowdsaleImplTransactor // Write-only binding to the contract
	TimedCrowdsaleImplFilterer   // Log filterer for contract events
}

// TimedCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimedCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedCrowdsaleImplSession struct {
	Contract     *TimedCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimedCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedCrowdsaleImplCallerSession struct {
	Contract *TimedCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TimedCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedCrowdsaleImplTransactorSession struct {
	Contract     *TimedCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TimedCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedCrowdsaleImplRaw struct {
	Contract *TimedCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// TimedCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedCrowdsaleImplCallerRaw struct {
	Contract *TimedCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// TimedCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedCrowdsaleImplTransactorRaw struct {
	Contract *TimedCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimedCrowdsaleImpl creates a new instance of TimedCrowdsaleImpl, bound to a specific deployed contract.
func NewTimedCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*TimedCrowdsaleImpl, error) {
	contract, err := bindTimedCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleImpl{TimedCrowdsaleImplCaller: TimedCrowdsaleImplCaller{contract: contract}, TimedCrowdsaleImplTransactor: TimedCrowdsaleImplTransactor{contract: contract}, TimedCrowdsaleImplFilterer: TimedCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewTimedCrowdsaleImplCaller creates a new read-only instance of TimedCrowdsaleImpl, bound to a specific deployed contract.
func NewTimedCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*TimedCrowdsaleImplCaller, error) {
	contract, err := bindTimedCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleImplCaller{contract: contract}, nil
}

// NewTimedCrowdsaleImplTransactor creates a new write-only instance of TimedCrowdsaleImpl, bound to a specific deployed contract.
func NewTimedCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedCrowdsaleImplTransactor, error) {
	contract, err := bindTimedCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleImplTransactor{contract: contract}, nil
}

// NewTimedCrowdsaleImplFilterer creates a new log filterer instance of TimedCrowdsaleImpl, bound to a specific deployed contract.
func NewTimedCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*TimedCrowdsaleImplFilterer, error) {
	contract, err := bindTimedCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleImplFilterer{contract: contract}, nil
}

// bindTimedCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindTimedCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsaleImpl.Contract.TimedCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.Contract.TimedCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.Contract.TimedCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.ClosingTime(&_TimedCrowdsaleImpl.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.ClosingTime(&_TimedCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) HasClosed() (bool, error) {
	return _TimedCrowdsaleImpl.Contract.HasClosed(&_TimedCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) HasClosed() (bool, error) {
	return _TimedCrowdsaleImpl.Contract.HasClosed(&_TimedCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.OpeningTime(&_TimedCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.OpeningTime(&_TimedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.Rate(&_TimedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.Rate(&_TimedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) Token() (common.Address, error) {
	return _TimedCrowdsaleImpl.Contract.Token(&_TimedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _TimedCrowdsaleImpl.Contract.Token(&_TimedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _TimedCrowdsaleImpl.Contract.Wallet(&_TimedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _TimedCrowdsaleImpl.Contract.Wallet(&_TimedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimedCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.WeiRaised(&_TimedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsaleImpl.Contract.WeiRaised(&_TimedCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.Contract.BuyTokens(&_TimedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsaleImpl.Contract.BuyTokens(&_TimedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// TimedCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the TimedCrowdsaleImpl contract.
type TimedCrowdsaleImplTokenPurchaseIterator struct {
	Event *TimedCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleImplTokenPurchase)
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
		it.Event = new(TimedCrowdsaleImplTokenPurchase)
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
func (it *TimedCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the TimedCrowdsaleImpl contract.
type TimedCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*TimedCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _TimedCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleImplTokenPurchaseIterator{contract: _TimedCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_TimedCrowdsaleImpl *TimedCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _TimedCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleImplTokenPurchase)
				if err := _TimedCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
