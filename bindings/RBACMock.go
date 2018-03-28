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

// RBACMockABI is the input ABI used to generate the binding from.
const RBACMockABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeAdvisor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyAdvisorsCanDoThis\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nobodyCanDoThis\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyAdminsCanDoThis\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eitherAdminOrAdvisorCanDoThis\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"adminRemoveRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"adminAddRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_ADMIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_advisors\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// RBACMockBin is the compiled bytecode used for deploying new contracts.
const RBACMockBin = `6060604052341561000f57600080fd5b604051610b8d380380610b8d83398101604052808051909101905060006100763360408051908101604052600581527f61646d696e000000000000000000000000000000000000000000000000000000602082015264010000000061077e61013b82021704565b6100c03360408051908101604052600781527f61647669736f7200000000000000000000000000000000000000000000000000602082015264010000000061077e61013b82021704565b5060005b81518110156101345761012c8282815181106100dc57fe5b9060200190602002015160408051908101604052600781527f61647669736f7200000000000000000000000000000000000000000000000000602082015264010000000061077e61013b82021704565b6001016100c4565b505061028a565b6101b8826000836040518082805190602001908083835b602083106101715780518252601f199092019160209182019101610152565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209064010000000061088361026582021704565b7fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b7004898282604051600160a060020a038316815260406020820181815290820183818151815260200191508051906020019080838360005b8381101561022657808201518382015260200161020e565b50505050905090810190601f1680156102535780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6108f4806102996000396000f3006060604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c81146100a85780630eeccb841461010957806316aabcbd14610128578063217fe6c61461013b578063420d0ba4146101ae5780634fd3d125146101c157806363a5bc71146101d457806388cee87e146101e7578063b25fa92c14610246578063d391014b146102a5575b600080fd5b34156100b357600080fd5b61010760048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061032f95505050505050565b005b341561011457600080fd5b610107600160a060020a03600435166103a9565b341561013357600080fd5b61010761042a565b341561014657600080fd5b61019a60048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061045695505050505050565b604051901515815260200160405180910390f35b34156101b957600080fd5b6101076104d5565b34156101cc57600080fd5b610107610513565b34156101df57600080fd5b61010761053d565b34156101f257600080fd5b61010760048035600160a060020a03169060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506105a195505050505050565b341561025157600080fd5b61010760048035600160a060020a03169060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506105d395505050505050565b34156102b057600080fd5b6102b8610605565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102f45780820151838201526020016102dc565b50505050905090810190601f1680156103215780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103a5826000836040518082805190602001908083835b602083106103655780518252601f199092019160209182019101610346565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61062716565b5050565b6103d133604080519081016040526005815260d960020a6430b236b4b702602082015261032f565b6103fc8160408051908101604052600781526000805160206108a9833981519152602082015261032f565b6104278160408051908101604052600781526000805160206108a9833981519152602082015261063c565b50565b60408051908101604052600781526000805160206108a98339815191526020820152610427338261032f565b60006104ce836000846040518082805190602001908083835b6020831061048e5780518252601f19909201916020918201910161046f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61075f16565b9392505050565b60408051908101604052600781527f756e6b6e6f776e000000000000000000000000000000000000000000000000006020820152610427338261032f565b61053b33604080519081016040526005815260d960020a6430b236b4b702602082015261032f565b565b61056533604080519081016040526005815260d960020a6430b236b4b7026020820152610456565b8061059657506105963360408051908101604052600781526000805160206108a98339815191526020820152610456565b151561053b57600080fd5b6105c933604080519081016040526005815260d960020a6430b236b4b702602082015261032f565b6103a5828261063c565b6105fb33604080519081016040526005815260d960020a6430b236b4b702602082015261032f565b6103a5828261077e565b604080519081016040526005815260d960020a6430b236b4b702602082015281565b610631828261075f565b15156103a557600080fd5b6106b2826000836040518082805190602001908083835b602083106106725780518252601f199092019160209182019101610653565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61086116565b7fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a8282604051600160a060020a038316815260406020820181815290820183818151815260200191508051906020019080838360005b83811015610720578082015183820152602001610708565b50505050905090810190601f16801561074d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b600160a060020a03166000908152602091909152604090205460ff1690565b6107f4826000836040518082805190602001908083835b602083106107b45780518252601f199092019160209182019101610795565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209063ffffffff61088316565b7fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b7004898282604051600160a060020a0383168152604060208201818152908201838181518152602001915080519060200190808383600083811015610720578082015183820152602001610708565b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0316600090815260209190915260409020805460ff19166001179055560061647669736f7200000000000000000000000000000000000000000000000000a165627a7a723058201562fe4cf8e67c08ba80ac4856be06fbe9e40786738699781b2c3f08e7253ee30029`

// DeployRBACMock deploys a new Ethereum contract, binding an instance of RBACMock to it.
func DeployRBACMock(auth *bind.TransactOpts, backend bind.ContractBackend, _advisors []common.Address) (common.Address, *types.Transaction, *RBACMock, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RBACMockBin), backend, _advisors)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RBACMock{RBACMockCaller: RBACMockCaller{contract: contract}, RBACMockTransactor: RBACMockTransactor{contract: contract}}, nil
}

// RBACMock is an auto generated Go binding around an Ethereum contract.
type RBACMock struct {
	RBACMockCaller     // Read-only binding to the contract
	RBACMockTransactor // Write-only binding to the contract
}

// RBACMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RBACMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RBACMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RBACMockSession struct {
	Contract     *RBACMock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RBACMockCallerSession struct {
	Contract *RBACMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RBACMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RBACMockTransactorSession struct {
	Contract     *RBACMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RBACMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RBACMockRaw struct {
	Contract *RBACMock // Generic contract binding to access the raw methods on
}

// RBACMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RBACMockCallerRaw struct {
	Contract *RBACMockCaller // Generic read-only contract binding to access the raw methods on
}

// RBACMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RBACMockTransactorRaw struct {
	Contract *RBACMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRBACMock creates a new instance of RBACMock, bound to a specific deployed contract.
func NewRBACMock(address common.Address, backend bind.ContractBackend) (*RBACMock, error) {
	contract, err := bindRBACMock(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RBACMock{RBACMockCaller: RBACMockCaller{contract: contract}, RBACMockTransactor: RBACMockTransactor{contract: contract}}, nil
}

// NewRBACMockCaller creates a new read-only instance of RBACMock, bound to a specific deployed contract.
func NewRBACMockCaller(address common.Address, caller bind.ContractCaller) (*RBACMockCaller, error) {
	contract, err := bindRBACMock(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RBACMockCaller{contract: contract}, nil
}

// NewRBACMockTransactor creates a new write-only instance of RBACMock, bound to a specific deployed contract.
func NewRBACMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RBACMockTransactor, error) {
	contract, err := bindRBACMock(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RBACMockTransactor{contract: contract}, nil
}

// bindRBACMock binds a generic wrapper to an already deployed contract.
func bindRBACMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBACMock *RBACMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBACMock.Contract.RBACMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBACMock *RBACMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBACMock.Contract.RBACMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBACMock *RBACMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBACMock.Contract.RBACMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBACMock *RBACMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBACMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBACMock *RBACMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBACMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBACMock *RBACMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBACMock.Contract.contract.Transact(opts, method, params...)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_RBACMock *RBACMockCaller) ROLE_ADMIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RBACMock.contract.Call(opts, out, "ROLE_ADMIN")
	return *ret0, err
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_RBACMock *RBACMockSession) ROLE_ADMIN() (string, error) {
	return _RBACMock.Contract.ROLE_ADMIN(&_RBACMock.CallOpts)
}

// ROLE_ADMIN is a free data retrieval call binding the contract method 0xd391014b.
//
// Solidity: function ROLE_ADMIN() constant returns(string)
func (_RBACMock *RBACMockCallerSession) ROLE_ADMIN() (string, error) {
	return _RBACMock.Contract.ROLE_ADMIN(&_RBACMock.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBACMock *RBACMockCaller) CheckRole(opts *bind.CallOpts, addr common.Address, roleName string) error {
	var ()
	out := &[]interface{}{}
	err := _RBACMock.contract.Call(opts, out, "checkRole", addr, roleName)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBACMock *RBACMockSession) CheckRole(addr common.Address, roleName string) error {
	return _RBACMock.Contract.CheckRole(&_RBACMock.CallOpts, addr, roleName)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBACMock *RBACMockCallerSession) CheckRole(addr common.Address, roleName string) error {
	return _RBACMock.Contract.CheckRole(&_RBACMock.CallOpts, addr, roleName)
}

// EitherAdminOrAdvisorCanDoThis is a free data retrieval call binding the contract method 0x63a5bc71.
//
// Solidity: function eitherAdminOrAdvisorCanDoThis() constant returns()
func (_RBACMock *RBACMockCaller) EitherAdminOrAdvisorCanDoThis(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _RBACMock.contract.Call(opts, out, "eitherAdminOrAdvisorCanDoThis")
	return err
}

// EitherAdminOrAdvisorCanDoThis is a free data retrieval call binding the contract method 0x63a5bc71.
//
// Solidity: function eitherAdminOrAdvisorCanDoThis() constant returns()
func (_RBACMock *RBACMockSession) EitherAdminOrAdvisorCanDoThis() error {
	return _RBACMock.Contract.EitherAdminOrAdvisorCanDoThis(&_RBACMock.CallOpts)
}

// EitherAdminOrAdvisorCanDoThis is a free data retrieval call binding the contract method 0x63a5bc71.
//
// Solidity: function eitherAdminOrAdvisorCanDoThis() constant returns()
func (_RBACMock *RBACMockCallerSession) EitherAdminOrAdvisorCanDoThis() error {
	return _RBACMock.Contract.EitherAdminOrAdvisorCanDoThis(&_RBACMock.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBACMock *RBACMockCaller) HasRole(opts *bind.CallOpts, addr common.Address, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBACMock.contract.Call(opts, out, "hasRole", addr, roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBACMock *RBACMockSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _RBACMock.Contract.HasRole(&_RBACMock.CallOpts, addr, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBACMock *RBACMockCallerSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _RBACMock.Contract.HasRole(&_RBACMock.CallOpts, addr, roleName)
}

// NobodyCanDoThis is a free data retrieval call binding the contract method 0x420d0ba4.
//
// Solidity: function nobodyCanDoThis() constant returns()
func (_RBACMock *RBACMockCaller) NobodyCanDoThis(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _RBACMock.contract.Call(opts, out, "nobodyCanDoThis")
	return err
}

// NobodyCanDoThis is a free data retrieval call binding the contract method 0x420d0ba4.
//
// Solidity: function nobodyCanDoThis() constant returns()
func (_RBACMock *RBACMockSession) NobodyCanDoThis() error {
	return _RBACMock.Contract.NobodyCanDoThis(&_RBACMock.CallOpts)
}

// NobodyCanDoThis is a free data retrieval call binding the contract method 0x420d0ba4.
//
// Solidity: function nobodyCanDoThis() constant returns()
func (_RBACMock *RBACMockCallerSession) NobodyCanDoThis() error {
	return _RBACMock.Contract.NobodyCanDoThis(&_RBACMock.CallOpts)
}

// OnlyAdminsCanDoThis is a free data retrieval call binding the contract method 0x4fd3d125.
//
// Solidity: function onlyAdminsCanDoThis() constant returns()
func (_RBACMock *RBACMockCaller) OnlyAdminsCanDoThis(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _RBACMock.contract.Call(opts, out, "onlyAdminsCanDoThis")
	return err
}

// OnlyAdminsCanDoThis is a free data retrieval call binding the contract method 0x4fd3d125.
//
// Solidity: function onlyAdminsCanDoThis() constant returns()
func (_RBACMock *RBACMockSession) OnlyAdminsCanDoThis() error {
	return _RBACMock.Contract.OnlyAdminsCanDoThis(&_RBACMock.CallOpts)
}

// OnlyAdminsCanDoThis is a free data retrieval call binding the contract method 0x4fd3d125.
//
// Solidity: function onlyAdminsCanDoThis() constant returns()
func (_RBACMock *RBACMockCallerSession) OnlyAdminsCanDoThis() error {
	return _RBACMock.Contract.OnlyAdminsCanDoThis(&_RBACMock.CallOpts)
}

// OnlyAdvisorsCanDoThis is a free data retrieval call binding the contract method 0x16aabcbd.
//
// Solidity: function onlyAdvisorsCanDoThis() constant returns()
func (_RBACMock *RBACMockCaller) OnlyAdvisorsCanDoThis(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _RBACMock.contract.Call(opts, out, "onlyAdvisorsCanDoThis")
	return err
}

// OnlyAdvisorsCanDoThis is a free data retrieval call binding the contract method 0x16aabcbd.
//
// Solidity: function onlyAdvisorsCanDoThis() constant returns()
func (_RBACMock *RBACMockSession) OnlyAdvisorsCanDoThis() error {
	return _RBACMock.Contract.OnlyAdvisorsCanDoThis(&_RBACMock.CallOpts)
}

// OnlyAdvisorsCanDoThis is a free data retrieval call binding the contract method 0x16aabcbd.
//
// Solidity: function onlyAdvisorsCanDoThis() constant returns()
func (_RBACMock *RBACMockCallerSession) OnlyAdvisorsCanDoThis() error {
	return _RBACMock.Contract.OnlyAdvisorsCanDoThis(&_RBACMock.CallOpts)
}

// AdminAddRole is a paid mutator transaction binding the contract method 0xb25fa92c.
//
// Solidity: function adminAddRole(addr address, roleName string) returns()
func (_RBACMock *RBACMockTransactor) AdminAddRole(opts *bind.TransactOpts, addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBACMock.contract.Transact(opts, "adminAddRole", addr, roleName)
}

// AdminAddRole is a paid mutator transaction binding the contract method 0xb25fa92c.
//
// Solidity: function adminAddRole(addr address, roleName string) returns()
func (_RBACMock *RBACMockSession) AdminAddRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBACMock.Contract.AdminAddRole(&_RBACMock.TransactOpts, addr, roleName)
}

// AdminAddRole is a paid mutator transaction binding the contract method 0xb25fa92c.
//
// Solidity: function adminAddRole(addr address, roleName string) returns()
func (_RBACMock *RBACMockTransactorSession) AdminAddRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBACMock.Contract.AdminAddRole(&_RBACMock.TransactOpts, addr, roleName)
}

// AdminRemoveRole is a paid mutator transaction binding the contract method 0x88cee87e.
//
// Solidity: function adminRemoveRole(addr address, roleName string) returns()
func (_RBACMock *RBACMockTransactor) AdminRemoveRole(opts *bind.TransactOpts, addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBACMock.contract.Transact(opts, "adminRemoveRole", addr, roleName)
}

// AdminRemoveRole is a paid mutator transaction binding the contract method 0x88cee87e.
//
// Solidity: function adminRemoveRole(addr address, roleName string) returns()
func (_RBACMock *RBACMockSession) AdminRemoveRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBACMock.Contract.AdminRemoveRole(&_RBACMock.TransactOpts, addr, roleName)
}

// AdminRemoveRole is a paid mutator transaction binding the contract method 0x88cee87e.
//
// Solidity: function adminRemoveRole(addr address, roleName string) returns()
func (_RBACMock *RBACMockTransactorSession) AdminRemoveRole(addr common.Address, roleName string) (*types.Transaction, error) {
	return _RBACMock.Contract.AdminRemoveRole(&_RBACMock.TransactOpts, addr, roleName)
}

// RemoveAdvisor is a paid mutator transaction binding the contract method 0x0eeccb84.
//
// Solidity: function removeAdvisor(_addr address) returns()
func (_RBACMock *RBACMockTransactor) RemoveAdvisor(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RBACMock.contract.Transact(opts, "removeAdvisor", _addr)
}

// RemoveAdvisor is a paid mutator transaction binding the contract method 0x0eeccb84.
//
// Solidity: function removeAdvisor(_addr address) returns()
func (_RBACMock *RBACMockSession) RemoveAdvisor(_addr common.Address) (*types.Transaction, error) {
	return _RBACMock.Contract.RemoveAdvisor(&_RBACMock.TransactOpts, _addr)
}

// RemoveAdvisor is a paid mutator transaction binding the contract method 0x0eeccb84.
//
// Solidity: function removeAdvisor(_addr address) returns()
func (_RBACMock *RBACMockTransactorSession) RemoveAdvisor(_addr common.Address) (*types.Transaction, error) {
	return _RBACMock.Contract.RemoveAdvisor(&_RBACMock.TransactOpts, _addr)
}
