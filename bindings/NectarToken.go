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

// NectarTokenABI is the input ABI used to generate the binding from.
const NectarTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// NectarTokenBin is the compiled bytecode used for deploying new contracts.
const NectarTokenBin = `606060409081526003805460a060020a60ff02191690558051908101604052600e81527f546573744e6574204e65637461720000000000000000000000000000000000006020820152600490805161005b9291602001906100d3565b5060408051908101604052600481527f544e435400000000000000000000000000000000000000000000000000000000602082015260059080516100a39291602001906100d3565b506006805460ff1916601217905560038054600160a060020a033316600160a060020a031990911617905561016e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b61016b91905b8082111561014d5760008155600101610157565b90565b610dcf8061017d6000396000f3006060604052600436106100d75763ffffffff60e060020a60003504166305d2035b81146100dc57806306fdde0314610103578063095ea7b31461018d57806318160ddd146101af57806323b872dd146101d4578063313ce567146101fc57806340c10f1914610225578063661884631461024757806370a08231146102695780637d64bcb4146102885780638da5cb5b1461029b57806395d89b41146102ca578063a9059cbb146102dd578063cae9ca51146102ff578063d73dd62314610364578063dd62ed3e14610386578063f2fde38b146103ab575b600080fd5b34156100e757600080fd5b6100ef6103cc565b604051901515815260200160405180910390f35b341561010e57600080fd5b6101166103dc565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561015257808201518382015260200161013a565b50505050905090810190601f16801561017f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561019857600080fd5b6100ef600160a060020a036004351660243561047a565b34156101ba57600080fd5b6101c26104d4565b60405190815260200160405180910390f35b34156101df57600080fd5b6100ef600160a060020a03600435811690602435166044356104da565b341561020757600080fd5b61020f61065c565b60405160ff909116815260200160405180910390f35b341561023057600080fd5b6100ef600160a060020a0360043516602435610665565b341561025257600080fd5b6100ef600160a060020a0360043516602435610772565b341561027457600080fd5b6101c2600160a060020a036004351661085a565b341561029357600080fd5b6100ef610875565b34156102a657600080fd5b6102ae610900565b604051600160a060020a03909116815260200160405180910390f35b34156102d557600080fd5b61011661090f565b34156102e857600080fd5b6100ef600160a060020a036004351660243561097a565b341561030a57600080fd5b6100ef60048035600160a060020a03169060248035919060649060443590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610a7595505050505050565b341561036f57600080fd5b6100ef600160a060020a0360043516602435610c03565b341561039157600080fd5b6101c2600160a060020a0360043581169060243516610c95565b34156103b657600080fd5b6103ca600160a060020a0360043516610cc0565b005b60035460a060020a900460ff1681565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104725780601f1061044757610100808354040283529160200191610472565b820191906000526020600020905b81548152906001019060200180831161045557829003601f168201915b505050505081565b600160a060020a0333811660008181526002602090815260408083209487168084529490915280822085905590929190600080516020610d848339815191529085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a03831615156104f157600080fd5b600160a060020a03841660009081526001602052604090205482111561051657600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561054957600080fd5b600160a060020a038416600090815260016020526040902054610572908363ffffffff610d5b16565b600160a060020a0380861660009081526001602052604080822093909355908516815220546105a7908363ffffffff610d6d16565b600160a060020a038085166000908152600160209081526040808320949094558783168252600281528382203390931682529190915220546105ef908363ffffffff610d5b16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60065460ff1681565b60035460009033600160a060020a0390811691161461068357600080fd5b60035460a060020a900460ff161561069a57600080fd5b6000546106ad908363ffffffff610d6d16565b6000908155600160a060020a0384168152600160205260409020546106d8908363ffffffff610d6d16565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156107cf57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610806565b6107df818463ffffffff610d5b16565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a033381166000818152600260209081526040808320948916808452949091529081902054600080516020610d84833981519152915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461089357600080fd5b60035460a060020a900460ff16156108aa57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104725780601f1061044757610100808354040283529160200191610472565b6000600160a060020a038316151561099157600080fd5b600160a060020a0333166000908152600160205260409020548211156109b657600080fd5b600160a060020a0333166000908152600160205260409020546109df908363ffffffff610d5b16565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a14908363ffffffff610d6d16565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660008181526002602090815260408083209488168084529490915280822086905590929190600080516020610d848339815191529086905190815260200160405180910390a383600160a060020a03166040517f72656365697665417070726f76616c28616464726573732c75696e743235362c81527f616464726573732c6279746573290000000000000000000000000000000000006020820152602e01604051809103902060e060020a9004338530866040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a03168152602001828051906020019080838360005b83811015610ba4578082015183820152602001610b8c565b50505050905090810190601f168015610bd15780820380516001836020036101000a031916815260200191505b5094505050505060006040518083038160008761646e5a03f1925050501515610bf957600080fd5b5060019392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610c3b908363ffffffff610d6d16565b600160a060020a033381166000818152600260209081526040808320948916808452949091529081902084905591929091600080516020610d8483398151915291905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a03908116911614610cdb57600080fd5b600160a060020a0381161515610cf057600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610d6757fe5b50900390565b600082820183811015610d7c57fe5b939250505056008c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a165627a7a723058203b4149957ff410f573fac2c954a130709e5225523405a8f3ba84d06aa6e5f06d0029`

// DeployNectarToken deploys a new Ethereum contract, binding an instance of NectarToken to it.
func DeployNectarToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NectarToken, error) {
	parsed, err := abi.JSON(strings.NewReader(NectarTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NectarTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NectarToken{NectarTokenCaller: NectarTokenCaller{contract: contract}, NectarTokenTransactor: NectarTokenTransactor{contract: contract}}, nil
}

// NectarToken is an auto generated Go binding around an Ethereum contract.
type NectarToken struct {
	NectarTokenCaller     // Read-only binding to the contract
	NectarTokenTransactor // Write-only binding to the contract
}

// NectarTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NectarTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NectarTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NectarTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NectarTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NectarTokenSession struct {
	Contract     *NectarToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NectarTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NectarTokenCallerSession struct {
	Contract *NectarTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NectarTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NectarTokenTransactorSession struct {
	Contract     *NectarTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NectarTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NectarTokenRaw struct {
	Contract *NectarToken // Generic contract binding to access the raw methods on
}

// NectarTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NectarTokenCallerRaw struct {
	Contract *NectarTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NectarTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NectarTokenTransactorRaw struct {
	Contract *NectarTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNectarToken creates a new instance of NectarToken, bound to a specific deployed contract.
func NewNectarToken(address common.Address, backend bind.ContractBackend) (*NectarToken, error) {
	contract, err := bindNectarToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NectarToken{NectarTokenCaller: NectarTokenCaller{contract: contract}, NectarTokenTransactor: NectarTokenTransactor{contract: contract}}, nil
}

// NewNectarTokenCaller creates a new read-only instance of NectarToken, bound to a specific deployed contract.
func NewNectarTokenCaller(address common.Address, caller bind.ContractCaller) (*NectarTokenCaller, error) {
	contract, err := bindNectarToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NectarTokenCaller{contract: contract}, nil
}

// NewNectarTokenTransactor creates a new write-only instance of NectarToken, bound to a specific deployed contract.
func NewNectarTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NectarTokenTransactor, error) {
	contract, err := bindNectarToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NectarTokenTransactor{contract: contract}, nil
}

// bindNectarToken binds a generic wrapper to an already deployed contract.
func bindNectarToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NectarTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NectarToken *NectarTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NectarToken.Contract.NectarTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NectarToken *NectarTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NectarToken.Contract.NectarTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NectarToken *NectarTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NectarToken.Contract.NectarTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NectarToken *NectarTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NectarToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NectarToken *NectarTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NectarToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NectarToken *NectarTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NectarToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_NectarToken *NectarTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_NectarToken *NectarTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _NectarToken.Contract.Allowance(&_NectarToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_NectarToken *NectarTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _NectarToken.Contract.Allowance(&_NectarToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_NectarToken *NectarTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_NectarToken *NectarTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _NectarToken.Contract.BalanceOf(&_NectarToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_NectarToken *NectarTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _NectarToken.Contract.BalanceOf(&_NectarToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_NectarToken *NectarTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_NectarToken *NectarTokenSession) Decimals() (uint8, error) {
	return _NectarToken.Contract.Decimals(&_NectarToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_NectarToken *NectarTokenCallerSession) Decimals() (uint8, error) {
	return _NectarToken.Contract.Decimals(&_NectarToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_NectarToken *NectarTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_NectarToken *NectarTokenSession) MintingFinished() (bool, error) {
	return _NectarToken.Contract.MintingFinished(&_NectarToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_NectarToken *NectarTokenCallerSession) MintingFinished() (bool, error) {
	return _NectarToken.Contract.MintingFinished(&_NectarToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NectarToken *NectarTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NectarToken *NectarTokenSession) Name() (string, error) {
	return _NectarToken.Contract.Name(&_NectarToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NectarToken *NectarTokenCallerSession) Name() (string, error) {
	return _NectarToken.Contract.Name(&_NectarToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NectarToken *NectarTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NectarToken *NectarTokenSession) Owner() (common.Address, error) {
	return _NectarToken.Contract.Owner(&_NectarToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NectarToken *NectarTokenCallerSession) Owner() (common.Address, error) {
	return _NectarToken.Contract.Owner(&_NectarToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NectarToken *NectarTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NectarToken *NectarTokenSession) Symbol() (string, error) {
	return _NectarToken.Contract.Symbol(&_NectarToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NectarToken *NectarTokenCallerSession) Symbol() (string, error) {
	return _NectarToken.Contract.Symbol(&_NectarToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NectarToken *NectarTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NectarToken *NectarTokenSession) TotalSupply() (*big.Int, error) {
	return _NectarToken.Contract.TotalSupply(&_NectarToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NectarToken *NectarTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NectarToken.Contract.TotalSupply(&_NectarToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Approve(&_NectarToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Approve(&_NectarToken.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_NectarToken *NectarTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_NectarToken *NectarTokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _NectarToken.Contract.ApproveAndCall(&_NectarToken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_NectarToken *NectarTokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _NectarToken.Contract.ApproveAndCall(&_NectarToken.TransactOpts, _spender, _value, _extraData)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_NectarToken *NectarTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.DecreaseApproval(&_NectarToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.DecreaseApproval(&_NectarToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_NectarToken *NectarTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_NectarToken *NectarTokenSession) FinishMinting() (*types.Transaction, error) {
	return _NectarToken.Contract.FinishMinting(&_NectarToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_NectarToken *NectarTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _NectarToken.Contract.FinishMinting(&_NectarToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_NectarToken *NectarTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.IncreaseApproval(&_NectarToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.IncreaseApproval(&_NectarToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_NectarToken *NectarTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Mint(&_NectarToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Mint(&_NectarToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Transfer(&_NectarToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Transfer(&_NectarToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.TransferFrom(&_NectarToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.TransferFrom(&_NectarToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NectarToken *NectarTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NectarToken *NectarTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NectarToken.Contract.TransferOwnership(&_NectarToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NectarToken *NectarTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NectarToken.Contract.TransferOwnership(&_NectarToken.TransactOpts, newOwner)
}
