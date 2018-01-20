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

// DetailedERC20MockABI is the input ABI used to generate the binding from.
const DetailedERC20MockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// DetailedERC20MockBin is the compiled bytecode used for deploying new contracts.
const DetailedERC20MockBin = `6060604052341561000f57600080fd5b604051610a0e380380610a0e83398101604052808051820191906020018051820191906020018051915083905082826003838051610051929160200190610085565b506004828051610065929160200190610085565b506005805460ff191660ff92909216919091179055506101209350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c657805160ff19168380011785556100f3565b828001600101855582156100f3579182015b828111156100f35782518255916020019190600101906100d8565b506100ff929150610103565b5090565b61011d91905b808211156100ff5760008155600101610109565b90565b6108df8061012f6000396000f3006060604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100b3578063095ea7b31461013d57806318160ddd1461017357806323b872dd14610198578063313ce567146101c057806366188463146101e957806370a082311461020b57806395d89b411461022a578063a9059cbb1461023d578063d73dd6231461025f578063dd62ed3e14610281575b600080fd5b34156100be57600080fd5b6100c66102a6565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101025780820151838201526020016100ea565b50505050905090810190601f16801561012f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561014857600080fd5b61015f600160a060020a0360043516602435610344565b604051901515815260200160405180910390f35b341561017e57600080fd5b6101866103b0565b60405190815260200160405180910390f35b34156101a357600080fd5b61015f600160a060020a03600435811690602435166044356103b6565b34156101cb57600080fd5b6101d3610538565b60405160ff909116815260200160405180910390f35b34156101f457600080fd5b61015f600160a060020a0360043516602435610541565b341561021657600080fd5b610186600160a060020a036004351661063b565b341561023557600080fd5b6100c6610656565b341561024857600080fd5b61015f600160a060020a03600435166024356106c1565b341561026a57600080fd5b61015f600160a060020a03600435166024356107bc565b341561028c57600080fd5b610186600160a060020a0360043581169060243516610860565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561033c5780601f106103115761010080835404028352916020019161033c565b820191906000526020600020905b81548152906001019060200180831161031f57829003601f168201915b505050505081565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a03831615156103cd57600080fd5b600160a060020a0384166000908152600160205260409020548211156103f257600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561042557600080fd5b600160a060020a03841660009081526001602052604090205461044e908363ffffffff61088b16565b600160a060020a038086166000908152600160205260408082209390935590851681522054610483908363ffffffff61089d16565b600160a060020a038085166000908152600160209081526040808320949094558783168252600281528382203390931682529190915220546104cb908363ffffffff61088b16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60055460ff1681565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561059e57600160a060020a0333811660009081526002602090815260408083209388168352929052908120556105d5565b6105ae818463ffffffff61088b16565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60048054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561033c5780601f106103115761010080835404028352916020019161033c565b6000600160a060020a03831615156106d857600080fd5b600160a060020a0333166000908152600160205260409020548211156106fd57600080fd5b600160a060020a033316600090815260016020526040902054610726908363ffffffff61088b16565b600160a060020a03338116600090815260016020526040808220939093559085168152205461075b908363ffffffff61089d16565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546107f4908363ffffffff61089d16565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561089757fe5b50900390565b6000828201838110156108ac57fe5b93925050505600a165627a7a72305820079b40df76d5ed0585f52fd4a8ab35979061effb94ed6bda8c3de9a7c47d661c0029`

// DeployDetailedERC20Mock deploys a new Ethereum contract, binding an instance of DetailedERC20Mock to it.
func DeployDetailedERC20Mock(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *DetailedERC20Mock, error) {
	parsed, err := abi.JSON(strings.NewReader(DetailedERC20MockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DetailedERC20MockBin), backend, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DetailedERC20Mock{DetailedERC20MockCaller: DetailedERC20MockCaller{contract: contract}, DetailedERC20MockTransactor: DetailedERC20MockTransactor{contract: contract}}, nil
}

// DetailedERC20Mock is an auto generated Go binding around an Ethereum contract.
type DetailedERC20Mock struct {
	DetailedERC20MockCaller     // Read-only binding to the contract
	DetailedERC20MockTransactor // Write-only binding to the contract
}

// DetailedERC20MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DetailedERC20MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedERC20MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DetailedERC20MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedERC20MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DetailedERC20MockSession struct {
	Contract     *DetailedERC20Mock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DetailedERC20MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DetailedERC20MockCallerSession struct {
	Contract *DetailedERC20MockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DetailedERC20MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DetailedERC20MockTransactorSession struct {
	Contract     *DetailedERC20MockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DetailedERC20MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DetailedERC20MockRaw struct {
	Contract *DetailedERC20Mock // Generic contract binding to access the raw methods on
}

// DetailedERC20MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DetailedERC20MockCallerRaw struct {
	Contract *DetailedERC20MockCaller // Generic read-only contract binding to access the raw methods on
}

// DetailedERC20MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DetailedERC20MockTransactorRaw struct {
	Contract *DetailedERC20MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDetailedERC20Mock creates a new instance of DetailedERC20Mock, bound to a specific deployed contract.
func NewDetailedERC20Mock(address common.Address, backend bind.ContractBackend) (*DetailedERC20Mock, error) {
	contract, err := bindDetailedERC20Mock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DetailedERC20Mock{DetailedERC20MockCaller: DetailedERC20MockCaller{contract: contract}, DetailedERC20MockTransactor: DetailedERC20MockTransactor{contract: contract}}, nil
}

// NewDetailedERC20MockCaller creates a new read-only instance of DetailedERC20Mock, bound to a specific deployed contract.
func NewDetailedERC20MockCaller(address common.Address, caller bind.ContractCaller) (*DetailedERC20MockCaller, error) {
	contract, err := bindDetailedERC20Mock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DetailedERC20MockCaller{contract: contract}, nil
}

// NewDetailedERC20MockTransactor creates a new write-only instance of DetailedERC20Mock, bound to a specific deployed contract.
func NewDetailedERC20MockTransactor(address common.Address, transactor bind.ContractTransactor) (*DetailedERC20MockTransactor, error) {
	contract, err := bindDetailedERC20Mock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DetailedERC20MockTransactor{contract: contract}, nil
}

// bindDetailedERC20Mock binds a generic wrapper to an already deployed contract.
func bindDetailedERC20Mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DetailedERC20MockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DetailedERC20Mock *DetailedERC20MockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DetailedERC20Mock.Contract.DetailedERC20MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DetailedERC20Mock *DetailedERC20MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.DetailedERC20MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DetailedERC20Mock *DetailedERC20MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.DetailedERC20MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DetailedERC20Mock *DetailedERC20MockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DetailedERC20Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DetailedERC20Mock *DetailedERC20MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DetailedERC20Mock *DetailedERC20MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_DetailedERC20Mock *DetailedERC20MockCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedERC20Mock.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_DetailedERC20Mock *DetailedERC20MockSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DetailedERC20Mock.Contract.Allowance(&_DetailedERC20Mock.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_DetailedERC20Mock *DetailedERC20MockCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DetailedERC20Mock.Contract.Allowance(&_DetailedERC20Mock.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_DetailedERC20Mock *DetailedERC20MockCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedERC20Mock.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_DetailedERC20Mock *DetailedERC20MockSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DetailedERC20Mock.Contract.BalanceOf(&_DetailedERC20Mock.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_DetailedERC20Mock *DetailedERC20MockCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DetailedERC20Mock.Contract.BalanceOf(&_DetailedERC20Mock.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedERC20Mock *DetailedERC20MockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DetailedERC20Mock.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedERC20Mock *DetailedERC20MockSession) Decimals() (uint8, error) {
	return _DetailedERC20Mock.Contract.Decimals(&_DetailedERC20Mock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedERC20Mock *DetailedERC20MockCallerSession) Decimals() (uint8, error) {
	return _DetailedERC20Mock.Contract.Decimals(&_DetailedERC20Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedERC20Mock *DetailedERC20MockCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DetailedERC20Mock.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedERC20Mock *DetailedERC20MockSession) Name() (string, error) {
	return _DetailedERC20Mock.Contract.Name(&_DetailedERC20Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedERC20Mock *DetailedERC20MockCallerSession) Name() (string, error) {
	return _DetailedERC20Mock.Contract.Name(&_DetailedERC20Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedERC20Mock *DetailedERC20MockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DetailedERC20Mock.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedERC20Mock *DetailedERC20MockSession) Symbol() (string, error) {
	return _DetailedERC20Mock.Contract.Symbol(&_DetailedERC20Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedERC20Mock *DetailedERC20MockCallerSession) Symbol() (string, error) {
	return _DetailedERC20Mock.Contract.Symbol(&_DetailedERC20Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedERC20Mock *DetailedERC20MockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedERC20Mock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedERC20Mock *DetailedERC20MockSession) TotalSupply() (*big.Int, error) {
	return _DetailedERC20Mock.Contract.TotalSupply(&_DetailedERC20Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedERC20Mock *DetailedERC20MockCallerSession) TotalSupply() (*big.Int, error) {
	return _DetailedERC20Mock.Contract.TotalSupply(&_DetailedERC20Mock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.Approve(&_DetailedERC20Mock.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.Approve(&_DetailedERC20Mock.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.DecreaseApproval(&_DetailedERC20Mock.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.DecreaseApproval(&_DetailedERC20Mock.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.IncreaseApproval(&_DetailedERC20Mock.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.IncreaseApproval(&_DetailedERC20Mock.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.Transfer(&_DetailedERC20Mock.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.Transfer(&_DetailedERC20Mock.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.TransferFrom(&_DetailedERC20Mock.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_DetailedERC20Mock *DetailedERC20MockTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DetailedERC20Mock.Contract.TransferFrom(&_DetailedERC20Mock.TransactOpts, _from, _to, _value)
}
