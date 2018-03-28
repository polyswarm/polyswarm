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

// SampleCrowdsaleABI is the input ABI used to generate the binding from.
const SampleCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_openingTime\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_goal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// SampleCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const SampleCrowdsaleBin = `60606040526007805460a060020a60ff0219169055341561001f57600080fd5b60405160e08061103a8339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805191508190508787858888876000831161007457600080fd5b600160a060020a038216151561008957600080fd5b600160a060020a038116151561009e57600080fd5b60029290925560018054600160a060020a03928316600160a060020a0319918216179091556000805492909316911617815581116100db57600080fd5b600455428210156100eb57600080fd5b818110156100f857600080fd5b60059190915560065560078054600160a060020a03191633600160a060020a03161790556000811161012957600080fd5b600154600160a060020a031661013d61019f565b600160a060020a039091168152602001604051809103906000f080151561016357600080fd5b60098054600160a060020a031916600160a060020a03929092169190911790556008558281111561019357600080fd5b505050505050506101af565b60405161060c80610a2e83390190565b610870806101be6000396000f3006060604052600436106100e25763ffffffff60e060020a6000350416631515bc2b81146100ed5780632c4e722e14610114578063355274ea14610139578063401938831461014c5780634042b66f1461015f5780634b6753bc146101725780634bb278f3146101855780634f93594514610198578063521eb273146101ab5780637d3d6522146101da5780638d4e4083146101ed5780638da5cb5b14610200578063b5545a3c14610213578063b7a8807c14610226578063ec8ac4d814610239578063f2fde38b1461024d578063fbfa77cf1461026c578063fc0c546a1461027f575b6100eb33610292565b005b34156100f857600080fd5b61010061033a565b604051901515815260200160405180910390f35b341561011f57600080fd5b610127610342565b60405190815260200160405180910390f35b341561014457600080fd5b610127610348565b341561015757600080fd5b61012761034e565b341561016a57600080fd5b610127610354565b341561017d57600080fd5b61012761035a565b341561019057600080fd5b6100eb610360565b34156101a357600080fd5b610100610421565b34156101b657600080fd5b6101be61042c565b604051600160a060020a03909116815260200160405180910390f35b34156101e557600080fd5b61010061043b565b34156101f857600080fd5b610100610446565b341561020b57600080fd5b6101be610467565b341561021e57600080fd5b6100eb610476565b341561023157600080fd5b61012761050e565b6100eb600160a060020a0360043516610292565b341561025857600080fd5b6100eb600160a060020a0360043516610514565b341561027757600080fd5b6101be6105af565b341561028a57600080fd5b6101be6105be565b34600061029f83836105cd565b6102a8826105fa565b6003549091506102be908363ffffffff61061716565b6003556102cb8382610631565b82600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a361032383836105f6565b61032b61063b565b61033583836105f6565b505050565b600654421190565b60025481565b60045481565b60085481565b60035481565b60065481565b60075433600160a060020a0390811691161461037b57600080fd5b60075474010000000000000000000000000000000000000000900460ff16156103a357600080fd5b6103ab61033a565b15156103b657600080fd5b6103be61069e565b7f6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b1768160405160405180910390a16007805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600454600354101590565b600154600160a060020a031681565b600854600354101590565b60075474010000000000000000000000000000000000000000900460ff1681565b600754600160a060020a031681565b60075474010000000000000000000000000000000000000000900460ff16151561049f57600080fd5b6104a761043b565b156104b157600080fd5b600954600160a060020a031663fa89401a3360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561050157600080fd5b5af1151561033557600080fd5b60055481565b60075433600160a060020a0390811691161461052f57600080fd5b600160a060020a038116151561054457600080fd5b600754600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600954600160a060020a031681565b600054600160a060020a031681565b60055442101580156105e157506006544211155b15156105ec57600080fd5b6105f68282610755565b5050565b60006106116002548361078090919063ffffffff16565b92915050565b60008282018381101561062657fe5b8091505b5092915050565b6105f682826107ab565b600954600160a060020a031663f340fa01343360405160e060020a63ffffffff8516028152600160a060020a0390911660048201526024016000604051808303818588803b151561068b57600080fd5b5af1151561069857600080fd5b50505050565b6106a661043b565b156106ff57600954600160a060020a03166343d726d66040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156106ea57600080fd5b5af115156106f757600080fd5b50505061074f565b600954600160a060020a0316638c52dc416040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561073e57600080fd5b5af1151561074b57600080fd5b5050505b6107535b565b61075f8282610823565b600454600354610775908363ffffffff61061716565b11156105f657600080fd5b600080831515610793576000915061062a565b508282028284828115156107a357fe5b041461062657fe5b600054600160a060020a03166340c10f19838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561080157600080fd5b5af1151561080e57600080fd5b5050506040518051905015156105f657600080fd5b600160a060020a038216151561083857600080fd5b8015156105f657600080fd00a165627a7a7230582058dcaf3a9f9cbce5271eac77796cfee35b496817dc3653a918e96cc48424963b00296060604052341561000f57600080fd5b60405160208061060c8339810160405280805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b60028054600160a060020a031916600160a060020a03929092169190911760a060020a60ff021916905561057e8061008e6000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166343d726d6811461009d578063521eb273146100b25780638c52dc41146100e15780638da5cb5b146100f4578063c19d93fb14610107578063cb13cddb1461013e578063f2fde38b1461016f578063f340fa011461018e578063fa89401a146101a2575b600080fd5b34156100a857600080fd5b6100b06101c1565b005b34156100bd57600080fd5b6100c561029c565b604051600160a060020a03909116815260200160405180910390f35b34156100ec57600080fd5b6100b06102ab565b34156100ff57600080fd5b6100c561033c565b341561011257600080fd5b61011a61034b565b6040518082600281111561012a57fe5b60ff16815260200191505060405180910390f35b341561014957600080fd5b61015d600160a060020a036004351661035b565b60405190815260200160405180910390f35b341561017a57600080fd5b6100b0600160a060020a036004351661036d565b6100b0600160a060020a0360043516610408565b34156101ad57600080fd5b6100b0600160a060020a036004351661048c565b60005433600160a060020a039081169116146101dc57600080fd5b60006002805460a060020a900460ff16908111156101f657fe5b1461020057600080fd5b6002805474ff00000000000000000000000000000000000000001916740200000000000000000000000000000000000000001790557f1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a60405160405180910390a1600254600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561029a57600080fd5b565b600254600160a060020a031681565b60005433600160a060020a039081169116146102c657600080fd5b60006002805460a060020a900460ff16908111156102e057fe5b146102ea57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a1790557f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8960405160405180910390a1565b600054600160a060020a031681565b60025460a060020a900460ff1681565b60016020526000908152604090205481565b60005433600160a060020a0390811691161461038857600080fd5b600160a060020a038116151561039d57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461042357600080fd5b60006002805460a060020a900460ff169081111561043d57fe5b1461044757600080fd5b600160a060020a038116600090815260016020526040902054610470903463ffffffff61053c16565b600160a060020a03909116600090815260016020526040902055565b600060016002805460a060020a900460ff16908111156104a857fe5b146104b257600080fd5b50600160a060020a038116600081815260016020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156104fb57600080fd5b81600160a060020a03167fd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d06518260405190815260200160405180910390a25050565b60008282018381101561054b57fe5b93925050505600a165627a7a72305820df47823a0e39486806747231e35e14d7a59c983510f8e97ead747f7e312efe420029`

// DeploySampleCrowdsale deploys a new Ethereum contract, binding an instance of SampleCrowdsale to it.
func DeploySampleCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _openingTime *big.Int, _closingTime *big.Int, _rate *big.Int, _wallet common.Address, _cap *big.Int, _token common.Address, _goal *big.Int) (common.Address, *types.Transaction, *SampleCrowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleCrowdsaleBin), backend, _openingTime, _closingTime, _rate, _wallet, _cap, _token, _goal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleCrowdsale{SampleCrowdsaleCaller: SampleCrowdsaleCaller{contract: contract}, SampleCrowdsaleTransactor: SampleCrowdsaleTransactor{contract: contract}}, nil
}

// SampleCrowdsale is an auto generated Go binding around an Ethereum contract.
type SampleCrowdsale struct {
	SampleCrowdsaleCaller     // Read-only binding to the contract
	SampleCrowdsaleTransactor // Write-only binding to the contract
}

// SampleCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleCrowdsaleSession struct {
	Contract     *SampleCrowdsale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleCrowdsaleCallerSession struct {
	Contract *SampleCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SampleCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleCrowdsaleTransactorSession struct {
	Contract     *SampleCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SampleCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleCrowdsaleRaw struct {
	Contract *SampleCrowdsale // Generic contract binding to access the raw methods on
}

// SampleCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleCrowdsaleCallerRaw struct {
	Contract *SampleCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// SampleCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleCrowdsaleTransactorRaw struct {
	Contract *SampleCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleCrowdsale creates a new instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsale(address common.Address, backend bind.ContractBackend) (*SampleCrowdsale, error) {
	contract, err := bindSampleCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsale{SampleCrowdsaleCaller: SampleCrowdsaleCaller{contract: contract}, SampleCrowdsaleTransactor: SampleCrowdsaleTransactor{contract: contract}}, nil
}

// NewSampleCrowdsaleCaller creates a new read-only instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*SampleCrowdsaleCaller, error) {
	contract, err := bindSampleCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleCaller{contract: contract}, nil
}

// NewSampleCrowdsaleTransactor creates a new write-only instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleCrowdsaleTransactor, error) {
	contract, err := bindSampleCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTransactor{contract: contract}, nil
}

// bindSampleCrowdsale binds a generic wrapper to an already deployed contract.
func bindSampleCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCrowdsale *SampleCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleCrowdsale.Contract.SampleCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCrowdsale *SampleCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.SampleCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCrowdsale *SampleCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.SampleCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCrowdsale *SampleCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCrowdsale *SampleCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCrowdsale *SampleCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) Cap() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Cap(&_SampleCrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Cap(&_SampleCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "capReached")
	return *ret0, err
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) CapReached() (bool, error) {
	return _SampleCrowdsale.Contract.CapReached(&_SampleCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) CapReached() (bool, error) {
	return _SampleCrowdsale.Contract.CapReached(&_SampleCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.ClosingTime(&_SampleCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.ClosingTime(&_SampleCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "goal")
	return *ret0, err
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) Goal() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Goal(&_SampleCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Goal() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Goal(&_SampleCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "goalReached")
	return *ret0, err
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) GoalReached() (bool, error) {
	return _SampleCrowdsale.Contract.GoalReached(&_SampleCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) GoalReached() (bool, error) {
	return _SampleCrowdsale.Contract.GoalReached(&_SampleCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) HasClosed() (bool, error) {
	return _SampleCrowdsale.Contract.HasClosed(&_SampleCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _SampleCrowdsale.Contract.HasClosed(&_SampleCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) IsFinalized() (bool, error) {
	return _SampleCrowdsale.Contract.IsFinalized(&_SampleCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) IsFinalized() (bool, error) {
	return _SampleCrowdsale.Contract.IsFinalized(&_SampleCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "openingTime")
	return *ret0, err
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.OpeningTime(&_SampleCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.OpeningTime(&_SampleCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Owner() (common.Address, error) {
	return _SampleCrowdsale.Contract.Owner(&_SampleCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _SampleCrowdsale.Contract.Owner(&_SampleCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) Rate() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Rate(&_SampleCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Rate(&_SampleCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Token() (common.Address, error) {
	return _SampleCrowdsale.Contract.Token(&_SampleCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Token() (common.Address, error) {
	return _SampleCrowdsale.Contract.Token(&_SampleCrowdsale.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Vault() (common.Address, error) {
	return _SampleCrowdsale.Contract.Vault(&_SampleCrowdsale.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Vault() (common.Address, error) {
	return _SampleCrowdsale.Contract.Vault(&_SampleCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Wallet() (common.Address, error) {
	return _SampleCrowdsale.Contract.Wallet(&_SampleCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _SampleCrowdsale.Contract.Wallet(&_SampleCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _SampleCrowdsale.Contract.WeiRaised(&_SampleCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _SampleCrowdsale.Contract.WeiRaised(&_SampleCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.BuyTokens(&_SampleCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.BuyTokens(&_SampleCrowdsale.TransactOpts, _beneficiary)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) ClaimRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "claimRefund")
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) ClaimRefund() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.ClaimRefund(&_SampleCrowdsale.TransactOpts)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) ClaimRefund() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.ClaimRefund(&_SampleCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.Finalize(&_SampleCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.Finalize(&_SampleCrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.TransferOwnership(&_SampleCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.TransferOwnership(&_SampleCrowdsale.TransactOpts, newOwner)
}
