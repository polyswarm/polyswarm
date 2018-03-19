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

// RefundableCrowdsaleImplABI is the input ABI used to generate the binding from.
const RefundableCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_goal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// RefundableCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const RefundableCrowdsaleImplBin = `60606040526006805460a060020a60ff0219169055341561001f57600080fd5b60405160c080610f98833981016040528080519190602001805191906020018051919060200180519190602001805191906020018051915081905086868686866000831161006c57600080fd5b600160a060020a038216151561008157600080fd5b600160a060020a038116151561009657600080fd5b60029290925560018054600160a060020a03928316600160a060020a03199182161790915560008054929093169116179055428210156100d557600080fd5b818110156100e257600080fd5b60049190915560055560068054600160a060020a03191633600160a060020a03161790556000811161011357600080fd5b600154600160a060020a031661012761017c565b600160a060020a039091168152602001604051809103906000f080151561014d57600080fd5b60088054600160a060020a031916600160a060020a03929092169190911790556007555061018c945050505050565b60405161060c8061098c83390190565b6107f18061019b6000396000f3006060604052600436106100cc5763ffffffff60e060020a6000350416631515bc2b81146100d75780632c4e722e146100fe57806340193883146101235780634042b66f146101365780634b6753bc146101495780634bb278f31461015c578063521eb2731461016f5780637d3d65221461019e5780638d4e4083146101b15780638da5cb5b146101c4578063b5545a3c146101d7578063b7a8807c146101ea578063ec8ac4d8146101fd578063f2fde38b14610211578063fbfa77cf14610230578063fc0c546a14610243575b6100d533610256565b005b34156100e257600080fd5b6100ea6102fe565b604051901515815260200160405180910390f35b341561010957600080fd5b610111610306565b60405190815260200160405180910390f35b341561012e57600080fd5b61011161030c565b341561014157600080fd5b610111610312565b341561015457600080fd5b610111610318565b341561016757600080fd5b6100d561031e565b341561017a57600080fd5b6101826103df565b604051600160a060020a03909116815260200160405180910390f35b34156101a957600080fd5b6100ea6103ee565b34156101bc57600080fd5b6100ea6103f9565b34156101cf57600080fd5b61018261041a565b34156101e257600080fd5b6100d5610429565b34156101f557600080fd5b6101116104c1565b6100d5600160a060020a0360043516610256565b341561021c57600080fd5b6100d5600160a060020a03600435166104c7565b341561023b57600080fd5b610182610562565b341561024e57600080fd5b610182610571565b3460006102638383610580565b61026c826105ad565b600354909150610282908363ffffffff6105ca16565b60035561028f83826105e4565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36102e783836105a9565b6102ef6105ee565b6102f983836105a9565b505050565b600554421190565b60025481565b60075481565b60035481565b60055481565b60065433600160a060020a0390811691161461033957600080fd5b60065474010000000000000000000000000000000000000000900460ff161561036157600080fd5b6103696102fe565b151561037457600080fd5b61037c610651565b7f6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b1768160405160405180910390a16006805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600154600160a060020a031681565b600754600354101590565b60065474010000000000000000000000000000000000000000900460ff1681565b600654600160a060020a031681565b60065474010000000000000000000000000000000000000000900460ff16151561045257600080fd5b61045a6103ee565b1561046457600080fd5b600854600160a060020a031663fa89401a3360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b15156104b457600080fd5b5af115156102f957600080fd5b60045481565b60065433600160a060020a039081169116146104e257600080fd5b600160a060020a03811615156104f757600080fd5b600654600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600854600160a060020a031681565b600054600160a060020a031681565b600454421015801561059457506005544211155b151561059f57600080fd5b6105a98282610708565b5050565b60006105c46002548361072990919063ffffffff16565b92915050565b6000828201838110156105d957fe5b8091505b5092915050565b6105a98282610754565b600854600160a060020a031663f340fa01343360405160e060020a63ffffffff8516028152600160a060020a0390911660048201526024016000604051808303818588803b151561063e57600080fd5b5af1151561064b57600080fd5b50505050565b6106596103ee565b156106b257600854600160a060020a03166343d726d66040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561069d57600080fd5b5af115156106aa57600080fd5b505050610702565b600854600160a060020a0316638c52dc416040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156106f157600080fd5b5af115156106fe57600080fd5b5050505b6107065b565b600160a060020a038216151561071d57600080fd5b8015156105a957600080fd5b60008083151561073c57600091506105dd565b5082820282848281151561074c57fe5b04146105d957fe5b600054600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156107aa57600080fd5b5af115156107b757600080fd5b5050506040518051505050505600a165627a7a72305820fda035ba6c02e53b8878b52d3ad9a7758794cd32a4ed28502f1c6a51cf8b26ab00296060604052341561000f57600080fd5b60405160208061060c8339810160405280805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b60028054600160a060020a031916600160a060020a03929092169190911760a060020a60ff021916905561057e8061008e6000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166343d726d6811461009d578063521eb273146100b25780638c52dc41146100e15780638da5cb5b146100f4578063c19d93fb14610107578063cb13cddb1461013e578063f2fde38b1461016f578063f340fa011461018e578063fa89401a146101a2575b600080fd5b34156100a857600080fd5b6100b06101c1565b005b34156100bd57600080fd5b6100c561029c565b604051600160a060020a03909116815260200160405180910390f35b34156100ec57600080fd5b6100b06102ab565b34156100ff57600080fd5b6100c561033c565b341561011257600080fd5b61011a61034b565b6040518082600281111561012a57fe5b60ff16815260200191505060405180910390f35b341561014957600080fd5b61015d600160a060020a036004351661035b565b60405190815260200160405180910390f35b341561017a57600080fd5b6100b0600160a060020a036004351661036d565b6100b0600160a060020a0360043516610408565b34156101ad57600080fd5b6100b0600160a060020a036004351661048c565b60005433600160a060020a039081169116146101dc57600080fd5b60006002805460a060020a900460ff16908111156101f657fe5b1461020057600080fd5b6002805474ff00000000000000000000000000000000000000001916740200000000000000000000000000000000000000001790557f1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a60405160405180910390a1600254600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561029a57600080fd5b565b600254600160a060020a031681565b60005433600160a060020a039081169116146102c657600080fd5b60006002805460a060020a900460ff16908111156102e057fe5b146102ea57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a1790557f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8960405160405180910390a1565b600054600160a060020a031681565b60025460a060020a900460ff1681565b60016020526000908152604090205481565b60005433600160a060020a0390811691161461038857600080fd5b600160a060020a038116151561039d57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461042357600080fd5b60006002805460a060020a900460ff169081111561043d57fe5b1461044757600080fd5b600160a060020a038116600090815260016020526040902054610470903463ffffffff61053c16565b600160a060020a03909116600090815260016020526040902055565b600060016002805460a060020a900460ff16908111156104a857fe5b146104b257600080fd5b50600160a060020a038116600081815260016020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156104fb57600080fd5b81600160a060020a03167fd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d06518260405190815260200160405180910390a25050565b60008282018381101561054b57fe5b93925050505600a165627a7a72305820df47823a0e39486806747231e35e14d7a59c983510f8e97ead747f7e312efe420029`

// DeployRefundableCrowdsaleImpl deploys a new Ethereum contract, binding an instance of RefundableCrowdsaleImpl to it.
func DeployRefundableCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _openingTime *big.Int, _closingTime *big.Int, _rate *big.Int, _wallet common.Address, _token common.Address, _goal *big.Int) (common.Address, *types.Transaction, *RefundableCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundableCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RefundableCrowdsaleImplBin), backend, _openingTime, _closingTime, _rate, _wallet, _token, _goal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RefundableCrowdsaleImpl{RefundableCrowdsaleImplCaller: RefundableCrowdsaleImplCaller{contract: contract}, RefundableCrowdsaleImplTransactor: RefundableCrowdsaleImplTransactor{contract: contract}, RefundableCrowdsaleImplFilterer: RefundableCrowdsaleImplFilterer{contract: contract}}, nil
}

// RefundableCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type RefundableCrowdsaleImpl struct {
	RefundableCrowdsaleImplCaller     // Read-only binding to the contract
	RefundableCrowdsaleImplTransactor // Write-only binding to the contract
	RefundableCrowdsaleImplFilterer   // Log filterer for contract events
}

// RefundableCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundableCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundableCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefundableCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundableCrowdsaleImplSession struct {
	Contract     *RefundableCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RefundableCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundableCrowdsaleImplCallerSession struct {
	Contract *RefundableCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// RefundableCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundableCrowdsaleImplTransactorSession struct {
	Contract     *RefundableCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// RefundableCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundableCrowdsaleImplRaw struct {
	Contract *RefundableCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// RefundableCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundableCrowdsaleImplCallerRaw struct {
	Contract *RefundableCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// RefundableCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundableCrowdsaleImplTransactorRaw struct {
	Contract *RefundableCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundableCrowdsaleImpl creates a new instance of RefundableCrowdsaleImpl, bound to a specific deployed contract.
func NewRefundableCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*RefundableCrowdsaleImpl, error) {
	contract, err := bindRefundableCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImpl{RefundableCrowdsaleImplCaller: RefundableCrowdsaleImplCaller{contract: contract}, RefundableCrowdsaleImplTransactor: RefundableCrowdsaleImplTransactor{contract: contract}, RefundableCrowdsaleImplFilterer: RefundableCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewRefundableCrowdsaleImplCaller creates a new read-only instance of RefundableCrowdsaleImpl, bound to a specific deployed contract.
func NewRefundableCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*RefundableCrowdsaleImplCaller, error) {
	contract, err := bindRefundableCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImplCaller{contract: contract}, nil
}

// NewRefundableCrowdsaleImplTransactor creates a new write-only instance of RefundableCrowdsaleImpl, bound to a specific deployed contract.
func NewRefundableCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundableCrowdsaleImplTransactor, error) {
	contract, err := bindRefundableCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImplTransactor{contract: contract}, nil
}

// NewRefundableCrowdsaleImplFilterer creates a new log filterer instance of RefundableCrowdsaleImpl, bound to a specific deployed contract.
func NewRefundableCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*RefundableCrowdsaleImplFilterer, error) {
	contract, err := bindRefundableCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImplFilterer{contract: contract}, nil
}

// bindRefundableCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindRefundableCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundableCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RefundableCrowdsaleImpl.Contract.RefundableCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.RefundableCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.RefundableCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RefundableCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) ClosingTime() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.ClosingTime(&_RefundableCrowdsaleImpl.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) ClosingTime() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.ClosingTime(&_RefundableCrowdsaleImpl.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "goal")
	return *ret0, err
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Goal() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.Goal(&_RefundableCrowdsaleImpl.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) Goal() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.Goal(&_RefundableCrowdsaleImpl.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "goalReached")
	return *ret0, err
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) GoalReached() (bool, error) {
	return _RefundableCrowdsaleImpl.Contract.GoalReached(&_RefundableCrowdsaleImpl.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) GoalReached() (bool, error) {
	return _RefundableCrowdsaleImpl.Contract.GoalReached(&_RefundableCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) HasClosed() (bool, error) {
	return _RefundableCrowdsaleImpl.Contract.HasClosed(&_RefundableCrowdsaleImpl.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) HasClosed() (bool, error) {
	return _RefundableCrowdsaleImpl.Contract.HasClosed(&_RefundableCrowdsaleImpl.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) IsFinalized() (bool, error) {
	return _RefundableCrowdsaleImpl.Contract.IsFinalized(&_RefundableCrowdsaleImpl.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) IsFinalized() (bool, error) {
	return _RefundableCrowdsaleImpl.Contract.IsFinalized(&_RefundableCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) OpeningTime() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.OpeningTime(&_RefundableCrowdsaleImpl.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) OpeningTime() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.OpeningTime(&_RefundableCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Owner() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Owner(&_RefundableCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) Owner() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Owner(&_RefundableCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.Rate(&_RefundableCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.Rate(&_RefundableCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Token() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Token(&_RefundableCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Token(&_RefundableCrowdsaleImpl.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Vault() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Vault(&_RefundableCrowdsaleImpl.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) Vault() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Vault(&_RefundableCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Wallet(&_RefundableCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _RefundableCrowdsaleImpl.Contract.Wallet(&_RefundableCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundableCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.WeiRaised(&_RefundableCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _RefundableCrowdsaleImpl.Contract.WeiRaised(&_RefundableCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.BuyTokens(&_RefundableCrowdsaleImpl.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.BuyTokens(&_RefundableCrowdsaleImpl.TransactOpts, _beneficiary)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactor) ClaimRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.contract.Transact(opts, "claimRefund")
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) ClaimRefund() (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.ClaimRefund(&_RefundableCrowdsaleImpl.TransactOpts)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactorSession) ClaimRefund() (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.ClaimRefund(&_RefundableCrowdsaleImpl.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) Finalize() (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.Finalize(&_RefundableCrowdsaleImpl.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactorSession) Finalize() (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.Finalize(&_RefundableCrowdsaleImpl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.TransferOwnership(&_RefundableCrowdsaleImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsaleImpl.Contract.TransferOwnership(&_RefundableCrowdsaleImpl.TransactOpts, newOwner)
}

// RefundableCrowdsaleImplFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the RefundableCrowdsaleImpl contract.
type RefundableCrowdsaleImplFinalizedIterator struct {
	Event *RefundableCrowdsaleImplFinalized // Event containing the contract specifics and raw log

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
func (it *RefundableCrowdsaleImplFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundableCrowdsaleImplFinalized)
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
		it.Event = new(RefundableCrowdsaleImplFinalized)
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
func (it *RefundableCrowdsaleImplFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundableCrowdsaleImplFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundableCrowdsaleImplFinalized represents a Finalized event raised by the RefundableCrowdsaleImpl contract.
type RefundableCrowdsaleImplFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplFilterer) FilterFinalized(opts *bind.FilterOpts) (*RefundableCrowdsaleImplFinalizedIterator, error) {

	logs, sub, err := _RefundableCrowdsaleImpl.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImplFinalizedIterator{contract: _RefundableCrowdsaleImpl.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *RefundableCrowdsaleImplFinalized) (event.Subscription, error) {

	logs, sub, err := _RefundableCrowdsaleImpl.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundableCrowdsaleImplFinalized)
				if err := _RefundableCrowdsaleImpl.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// RefundableCrowdsaleImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RefundableCrowdsaleImpl contract.
type RefundableCrowdsaleImplOwnershipTransferredIterator struct {
	Event *RefundableCrowdsaleImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RefundableCrowdsaleImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundableCrowdsaleImplOwnershipTransferred)
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
		it.Event = new(RefundableCrowdsaleImplOwnershipTransferred)
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
func (it *RefundableCrowdsaleImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundableCrowdsaleImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundableCrowdsaleImplOwnershipTransferred represents a OwnershipTransferred event raised by the RefundableCrowdsaleImpl contract.
type RefundableCrowdsaleImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RefundableCrowdsaleImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundableCrowdsaleImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImplOwnershipTransferredIterator{contract: _RefundableCrowdsaleImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RefundableCrowdsaleImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundableCrowdsaleImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundableCrowdsaleImplOwnershipTransferred)
				if err := _RefundableCrowdsaleImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RefundableCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the RefundableCrowdsaleImpl contract.
type RefundableCrowdsaleImplTokenPurchaseIterator struct {
	Event *RefundableCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *RefundableCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundableCrowdsaleImplTokenPurchase)
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
		it.Event = new(RefundableCrowdsaleImplTokenPurchase)
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
func (it *RefundableCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundableCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundableCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the RefundableCrowdsaleImpl contract.
type RefundableCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*RefundableCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundableCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleImplTokenPurchaseIterator{contract: _RefundableCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_RefundableCrowdsaleImpl *RefundableCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *RefundableCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundableCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundableCrowdsaleImplTokenPurchase)
				if err := _RefundableCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
