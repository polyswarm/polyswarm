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

// NectarTokenABI is the input ABI used to generate the binding from.
const NectarTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableTransfers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfersEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TransfersEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// NectarTokenBin is the compiled bytecode used for deploying new contracts.
const NectarTokenBin = `606060409081526003805460a060020a60ff02191690558051908101604052600681527f52656b74617200000000000000000000000000000000000000000000000000006020820152600490805161005b9291602001906100d8565b5060408051908101604052600381527f524b540000000000000000000000000000000000000000000000000000000000602082015260059080516100a39291602001906100d8565b5060068054601260ff199091161761ff001916905560038054600160a060020a03191633600160a060020a0316179055610173565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011957805160ff1916838001178555610146565b82800160010185558215610146579182015b8281111561014657825182559160200191906001019061012b565b50610152929150610156565b5090565b61017091905b80821115610152576000815560010161015c565b90565b610f2b806101826000396000f3006060604052600436106100ed5763ffffffff60e060020a60003504166305d2035b81146100f257806306fdde0314610119578063095ea7b3146101a357806318160ddd146101c557806323b872dd146101ea578063313ce5671461021257806340c10f191461023b578063661884631461025d57806370a082311461027f5780637d64bcb41461029e5780638da5cb5b146102b157806395d89b41146102e0578063a9059cbb146102f3578063af35c6c714610315578063bef97c871461032a578063cae9ca511461033d578063d73dd623146103a2578063dd62ed3e146103c4578063f2fde38b146103e9575b600080fd5b34156100fd57600080fd5b610105610408565b604051901515815260200160405180910390f35b341561012457600080fd5b61012c610429565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610168578082015183820152602001610150565b50505050905090810190601f1680156101955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101ae57600080fd5b610105600160a060020a03600435166024356104c7565b34156101d057600080fd5b6101d8610521565b60405190815260200160405180910390f35b34156101f557600080fd5b610105600160a060020a0360043581169060243516604435610527565b341561021d57600080fd5b610225610553565b60405160ff909116815260200160405180910390f35b341561024657600080fd5b610105600160a060020a036004351660243561055c565b341561026857600080fd5b610105600160a060020a036004351660243561067b565b341561028a57600080fd5b6101d8600160a060020a0360043516610763565b34156102a957600080fd5b61010561077e565b34156102bc57600080fd5b6102c461082b565b604051600160a060020a03909116815260200160405180910390f35b34156102eb57600080fd5b61012c61083a565b34156102fe57600080fd5b610105600160a060020a03600435166024356108a5565b341561032057600080fd5b6103286108cf565b005b341561033557600080fd5b61010561093c565b341561034857600080fd5b61010560048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061094a95505050505050565b34156103ad57600080fd5b610105600160a060020a0360043516602435610ad4565b34156103cf57600080fd5b6101d8600160a060020a0360043581169060243516610b66565b34156103f457600080fd5b610328600160a060020a0360043516610b91565b60035474010000000000000000000000000000000000000000900460ff1681565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104bf5780601f10610494576101008083540402835291602001916104bf565b820191906000526020600020905b8154815290600101906020018083116104a257829003601f168201915b505050505081565b600160a060020a0333811660008181526002602090815260408083209487168084529490915280822085905590929190600080516020610ee08339815191529085905190815260200160405180910390a350600192915050565b60015490565b600654600090610100900460ff16151561054057600080fd5b61054b848484610c2c565b949350505050565b60065460ff1681565b60035460009033600160a060020a0390811691161461057a57600080fd5b60035474010000000000000000000000000000000000000000900460ff16156105a257600080fd5b6001546105b5908363ffffffff610dac16565b600155600160a060020a0383166000908152602081905260409020546105e1908363ffffffff610dac16565b600160a060020a0384166000818152602081905260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156106d857600160a060020a03338116600090815260026020908152604080832093881683529290529081205561070f565b6106e8818463ffffffff610dbb16565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a033381166000818152600260209081526040808320948916808452949091529081902054600080516020610ee0833981519152915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b60035460009033600160a060020a0390811691161461079c57600080fd5b60035474010000000000000000000000000000000000000000900460ff16156107c457600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104bf5780601f10610494576101008083540402835291602001916104bf565b600654600090610100900460ff1615156108be57600080fd5b6108c88383610dcd565b9392505050565b60035433600160a060020a039081169116146108ea57600080fd5b600654610100900460ff16156108ff57600080fd5b6006805461ff0019166101001790557feadb24812ab3c9a55c774958184293ebdb6c7f6a2dbab11f397d80c86feb65d360405160405180910390a1565b600654610100900460ff1681565b600160a060020a0333811660008181526002602090815260408083209488168084529490915280822086905590929190600080516020610ee08339815191529086905190815260200160405180910390a383600160a060020a03166040517f72656365697665417070726f76616c28616464726573732c75696e743235362c81527f616464726573732c6279746573290000000000000000000000000000000000006020820152602e01604051809103902060e060020a9004338530866040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a03168152602001828051906020019080838360005b83811015610a79578082015183820152602001610a61565b50505050905090810190601f168015610aa65780820380516001836020036101000a031916815260200191505b509450505050506000604051808303816000875af1925050501515610aca57600080fd5b5060019392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610b0c908363ffffffff610dac16565b600160a060020a033381166000818152600260209081526040808320948916808452949091529081902084905591929091600080516020610ee083398151915291905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a03908116911614610bac57600080fd5b600160a060020a0381161515610bc157600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a0383161515610c4357600080fd5b600160a060020a038416600090815260208190526040902054821115610c6857600080fd5b600160a060020a0380851660009081526002602090815260408083203390941683529290522054821115610c9b57600080fd5b600160a060020a038416600090815260208190526040902054610cc4908363ffffffff610dbb16565b600160a060020a038086166000908152602081905260408082209390935590851681522054610cf9908363ffffffff610dac16565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610d3f908363ffffffff610dbb16565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b6000828201838110156108c857fe5b600082821115610dc757fe5b50900390565b6000600160a060020a0383161515610de457600080fd5b600160a060020a033316600090815260208190526040902054821115610e0957600080fd5b600160a060020a033316600090815260208190526040902054610e32908363ffffffff610dbb16565b600160a060020a033381166000908152602081905260408082209390935590851681522054610e67908363ffffffff610dac16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35060019291505056008c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a165627a7a723058208ef44eaf7916892cf0e6161832f89f51c873f85ab397b07df686957953747ca50029`

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
	return address, tx, &NectarToken{NectarTokenCaller: NectarTokenCaller{contract: contract}, NectarTokenTransactor: NectarTokenTransactor{contract: contract}, NectarTokenFilterer: NectarTokenFilterer{contract: contract}}, nil
}

// NectarToken is an auto generated Go binding around an Ethereum contract.
type NectarToken struct {
	NectarTokenCaller     // Read-only binding to the contract
	NectarTokenTransactor // Write-only binding to the contract
	NectarTokenFilterer   // Log filterer for contract events
}

// NectarTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NectarTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NectarTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NectarTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NectarTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NectarTokenFilterer struct {
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
	contract, err := bindNectarToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NectarToken{NectarTokenCaller: NectarTokenCaller{contract: contract}, NectarTokenTransactor: NectarTokenTransactor{contract: contract}, NectarTokenFilterer: NectarTokenFilterer{contract: contract}}, nil
}

// NewNectarTokenCaller creates a new read-only instance of NectarToken, bound to a specific deployed contract.
func NewNectarTokenCaller(address common.Address, caller bind.ContractCaller) (*NectarTokenCaller, error) {
	contract, err := bindNectarToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NectarTokenCaller{contract: contract}, nil
}

// NewNectarTokenTransactor creates a new write-only instance of NectarToken, bound to a specific deployed contract.
func NewNectarTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NectarTokenTransactor, error) {
	contract, err := bindNectarToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NectarTokenTransactor{contract: contract}, nil
}

// NewNectarTokenFilterer creates a new log filterer instance of NectarToken, bound to a specific deployed contract.
func NewNectarTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NectarTokenFilterer, error) {
	contract, err := bindNectarToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NectarTokenFilterer{contract: contract}, nil
}

// bindNectarToken binds a generic wrapper to an already deployed contract.
func bindNectarToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NectarTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() constant returns(bool)
func (_NectarToken *NectarTokenCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NectarToken.contract.Call(opts, out, "transfersEnabled")
	return *ret0, err
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() constant returns(bool)
func (_NectarToken *NectarTokenSession) TransfersEnabled() (bool, error) {
	return _NectarToken.Contract.TransfersEnabled(&_NectarToken.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() constant returns(bool)
func (_NectarToken *NectarTokenCallerSession) TransfersEnabled() (bool, error) {
	return _NectarToken.Contract.TransfersEnabled(&_NectarToken.CallOpts)
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

// EnableTransfers is a paid mutator transaction binding the contract method 0xaf35c6c7.
//
// Solidity: function enableTransfers() returns()
func (_NectarToken *NectarTokenTransactor) EnableTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "enableTransfers")
}

// EnableTransfers is a paid mutator transaction binding the contract method 0xaf35c6c7.
//
// Solidity: function enableTransfers() returns()
func (_NectarToken *NectarTokenSession) EnableTransfers() (*types.Transaction, error) {
	return _NectarToken.Contract.EnableTransfers(&_NectarToken.TransactOpts)
}

// EnableTransfers is a paid mutator transaction binding the contract method 0xaf35c6c7.
//
// Solidity: function enableTransfers() returns()
func (_NectarToken *NectarTokenTransactorSession) EnableTransfers() (*types.Transaction, error) {
	return _NectarToken.Contract.EnableTransfers(&_NectarToken.TransactOpts)
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
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_NectarToken *NectarTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Transfer(&_NectarToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.Transfer(&_NectarToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NectarToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_NectarToken *NectarTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.TransferFrom(&_NectarToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_NectarToken *NectarTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NectarToken.Contract.TransferFrom(&_NectarToken.TransactOpts, from, to, value)
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

// NectarTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NectarToken contract.
type NectarTokenApprovalIterator struct {
	Event *NectarTokenApproval // Event containing the contract specifics and raw log

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
func (it *NectarTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NectarTokenApproval)
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
		it.Event = new(NectarTokenApproval)
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
func (it *NectarTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NectarTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NectarTokenApproval represents a Approval event raised by the NectarToken contract.
type NectarTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_NectarToken *NectarTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NectarTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NectarToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NectarTokenApprovalIterator{contract: _NectarToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_NectarToken *NectarTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NectarTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NectarToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NectarTokenApproval)
				if err := _NectarToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// NectarTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the NectarToken contract.
type NectarTokenMintIterator struct {
	Event *NectarTokenMint // Event containing the contract specifics and raw log

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
func (it *NectarTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NectarTokenMint)
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
		it.Event = new(NectarTokenMint)
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
func (it *NectarTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NectarTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NectarTokenMint represents a Mint event raised by the NectarToken contract.
type NectarTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_NectarToken *NectarTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*NectarTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NectarToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &NectarTokenMintIterator{contract: _NectarToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_NectarToken *NectarTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *NectarTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NectarToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NectarTokenMint)
				if err := _NectarToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// NectarTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the NectarToken contract.
type NectarTokenMintFinishedIterator struct {
	Event *NectarTokenMintFinished // Event containing the contract specifics and raw log

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
func (it *NectarTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NectarTokenMintFinished)
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
		it.Event = new(NectarTokenMintFinished)
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
func (it *NectarTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NectarTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NectarTokenMintFinished represents a MintFinished event raised by the NectarToken contract.
type NectarTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_NectarToken *NectarTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*NectarTokenMintFinishedIterator, error) {

	logs, sub, err := _NectarToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &NectarTokenMintFinishedIterator{contract: _NectarToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_NectarToken *NectarTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *NectarTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _NectarToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NectarTokenMintFinished)
				if err := _NectarToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// NectarTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NectarToken contract.
type NectarTokenOwnershipTransferredIterator struct {
	Event *NectarTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NectarTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NectarTokenOwnershipTransferred)
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
		it.Event = new(NectarTokenOwnershipTransferred)
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
func (it *NectarTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NectarTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NectarTokenOwnershipTransferred represents a OwnershipTransferred event raised by the NectarToken contract.
type NectarTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NectarToken *NectarTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NectarTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NectarToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NectarTokenOwnershipTransferredIterator{contract: _NectarToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NectarToken *NectarTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NectarTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NectarToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NectarTokenOwnershipTransferred)
				if err := _NectarToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// NectarTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NectarToken contract.
type NectarTokenTransferIterator struct {
	Event *NectarTokenTransfer // Event containing the contract specifics and raw log

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
func (it *NectarTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NectarTokenTransfer)
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
		it.Event = new(NectarTokenTransfer)
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
func (it *NectarTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NectarTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NectarTokenTransfer represents a Transfer event raised by the NectarToken contract.
type NectarTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_NectarToken *NectarTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NectarTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NectarToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NectarTokenTransferIterator{contract: _NectarToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_NectarToken *NectarTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NectarTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NectarToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NectarTokenTransfer)
				if err := _NectarToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// NectarTokenTransfersEnabledIterator is returned from FilterTransfersEnabled and is used to iterate over the raw logs and unpacked data for TransfersEnabled events raised by the NectarToken contract.
type NectarTokenTransfersEnabledIterator struct {
	Event *NectarTokenTransfersEnabled // Event containing the contract specifics and raw log

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
func (it *NectarTokenTransfersEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NectarTokenTransfersEnabled)
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
		it.Event = new(NectarTokenTransfersEnabled)
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
func (it *NectarTokenTransfersEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NectarTokenTransfersEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NectarTokenTransfersEnabled represents a TransfersEnabled event raised by the NectarToken contract.
type NectarTokenTransfersEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfersEnabled is a free log retrieval operation binding the contract event 0xeadb24812ab3c9a55c774958184293ebdb6c7f6a2dbab11f397d80c86feb65d3.
//
// Solidity: event TransfersEnabled()
func (_NectarToken *NectarTokenFilterer) FilterTransfersEnabled(opts *bind.FilterOpts) (*NectarTokenTransfersEnabledIterator, error) {

	logs, sub, err := _NectarToken.contract.FilterLogs(opts, "TransfersEnabled")
	if err != nil {
		return nil, err
	}
	return &NectarTokenTransfersEnabledIterator{contract: _NectarToken.contract, event: "TransfersEnabled", logs: logs, sub: sub}, nil
}

// WatchTransfersEnabled is a free log subscription operation binding the contract event 0xeadb24812ab3c9a55c774958184293ebdb6c7f6a2dbab11f397d80c86feb65d3.
//
// Solidity: event TransfersEnabled()
func (_NectarToken *NectarTokenFilterer) WatchTransfersEnabled(opts *bind.WatchOpts, sink chan<- *NectarTokenTransfersEnabled) (event.Subscription, error) {

	logs, sub, err := _NectarToken.contract.WatchLogs(opts, "TransfersEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NectarTokenTransfersEnabled)
				if err := _NectarToken.contract.UnpackLog(event, "TransfersEnabled", log); err != nil {
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
