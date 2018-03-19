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

// MintedCrowdsaleImplABI is the input ABI used to generate the binding from.
const MintedCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// MintedCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const MintedCrowdsaleImplBin = `6060604052341561000f57600080fd5b604051606080610410833981016040528080519190602001805191906020018051915083905082826000831161004457600080fd5b600160a060020a038216151561005957600080fd5b600160a060020a038116151561006e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a0319918216179091556000805493909216921691909117905550505061035b806100b56000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e81146100775780634042b66f1461009c578063521eb273146100af578063ec8ac4d8146100de578063fc0c546a146100f2575b61007533610105565b005b341561008257600080fd5b61008a6101ad565b60405190815260200160405180910390f35b34156100a757600080fd5b61008a6101b3565b34156100ba57600080fd5b6100c26101b9565b604051600160a060020a03909116815260200160405180910390f35b610075600160a060020a0360043516610105565b34156100fd57600080fd5b6100c26101c8565b34600061011283836101d7565b61011b826101fc565b600354909150610131908363ffffffff61021916565b60035561013e8382610233565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a361019683836101f8565b61019e61023d565b6101a883836101f8565b505050565b60025481565b60035481565b600154600160a060020a031681565b600054600160a060020a031681565b600160a060020a03821615156101ec57600080fd5b8015156101f857600080fd5b5050565b60006102136002548361027390919063ffffffff16565b92915050565b60008282018381101561022857fe5b8091505b5092915050565b6101f8828261029e565b600154600160a060020a03163480156108fc0290604051600060405180830381858888f19350505050151561027157600080fd5b565b600080831515610286576000915061022c565b5082820282848281151561029657fe5b041461022857fe5b600054600160a060020a03166340c10f1983836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561030d57600080fd5b5af1151561031a57600080fd5b5050506040518051905015156101f857600080fd00a165627a7a72305820fd8a73a10bf11ad19ddedc2748c2247664a57d4fb2641d6b1ece3bf7546a123c0029`

// DeployMintedCrowdsaleImpl deploys a new Ethereum contract, binding an instance of MintedCrowdsaleImpl to it.
func DeployMintedCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *MintedCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(MintedCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MintedCrowdsaleImplBin), backend, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintedCrowdsaleImpl{MintedCrowdsaleImplCaller: MintedCrowdsaleImplCaller{contract: contract}, MintedCrowdsaleImplTransactor: MintedCrowdsaleImplTransactor{contract: contract}, MintedCrowdsaleImplFilterer: MintedCrowdsaleImplFilterer{contract: contract}}, nil
}

// MintedCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type MintedCrowdsaleImpl struct {
	MintedCrowdsaleImplCaller     // Read-only binding to the contract
	MintedCrowdsaleImplTransactor // Write-only binding to the contract
	MintedCrowdsaleImplFilterer   // Log filterer for contract events
}

// MintedCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintedCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintedCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintedCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintedCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintedCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintedCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintedCrowdsaleImplSession struct {
	Contract     *MintedCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MintedCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintedCrowdsaleImplCallerSession struct {
	Contract *MintedCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MintedCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintedCrowdsaleImplTransactorSession struct {
	Contract     *MintedCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MintedCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintedCrowdsaleImplRaw struct {
	Contract *MintedCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// MintedCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintedCrowdsaleImplCallerRaw struct {
	Contract *MintedCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// MintedCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintedCrowdsaleImplTransactorRaw struct {
	Contract *MintedCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintedCrowdsaleImpl creates a new instance of MintedCrowdsaleImpl, bound to a specific deployed contract.
func NewMintedCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*MintedCrowdsaleImpl, error) {
	contract, err := bindMintedCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleImpl{MintedCrowdsaleImplCaller: MintedCrowdsaleImplCaller{contract: contract}, MintedCrowdsaleImplTransactor: MintedCrowdsaleImplTransactor{contract: contract}, MintedCrowdsaleImplFilterer: MintedCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewMintedCrowdsaleImplCaller creates a new read-only instance of MintedCrowdsaleImpl, bound to a specific deployed contract.
func NewMintedCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*MintedCrowdsaleImplCaller, error) {
	contract, err := bindMintedCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleImplCaller{contract: contract}, nil
}

// NewMintedCrowdsaleImplTransactor creates a new write-only instance of MintedCrowdsaleImpl, bound to a specific deployed contract.
func NewMintedCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*MintedCrowdsaleImplTransactor, error) {
	contract, err := bindMintedCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleImplTransactor{contract: contract}, nil
}

// NewMintedCrowdsaleImplFilterer creates a new log filterer instance of MintedCrowdsaleImpl, bound to a specific deployed contract.
func NewMintedCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*MintedCrowdsaleImplFilterer, error) {
	contract, err := bindMintedCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleImplFilterer{contract: contract}, nil
}

// bindMintedCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindMintedCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintedCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintedCrowdsaleImpl.Contract.MintedCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.Contract.MintedCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.Contract.MintedCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintedCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintedCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _MintedCrowdsaleImpl.Contract.Rate(&_MintedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _MintedCrowdsaleImpl.Contract.Rate(&_MintedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintedCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplSession) Token() (common.Address, error) {
	return _MintedCrowdsaleImpl.Contract.Token(&_MintedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _MintedCrowdsaleImpl.Contract.Token(&_MintedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintedCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _MintedCrowdsaleImpl.Contract.Wallet(&_MintedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _MintedCrowdsaleImpl.Contract.Wallet(&_MintedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintedCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _MintedCrowdsaleImpl.Contract.WeiRaised(&_MintedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _MintedCrowdsaleImpl.Contract.WeiRaised(&_MintedCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.Contract.BuyTokens(&_MintedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _MintedCrowdsaleImpl.Contract.BuyTokens(&_MintedCrowdsaleImpl.TransactOpts, _beneficiary)
}

// MintedCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the MintedCrowdsaleImpl contract.
type MintedCrowdsaleImplTokenPurchaseIterator struct {
	Event *MintedCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *MintedCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintedCrowdsaleImplTokenPurchase)
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
		it.Event = new(MintedCrowdsaleImplTokenPurchase)
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
func (it *MintedCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintedCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintedCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the MintedCrowdsaleImpl contract.
type MintedCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*MintedCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _MintedCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &MintedCrowdsaleImplTokenPurchaseIterator{contract: _MintedCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_MintedCrowdsaleImpl *MintedCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *MintedCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _MintedCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintedCrowdsaleImplTokenPurchase)
				if err := _MintedCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
