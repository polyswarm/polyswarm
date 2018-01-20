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

// BountyRegistryABI is the input ABI used to generate the binding from.
const BountyRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"uint128\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"artifactHash\",\"type\":\"bytes32\"},{\"name\":\"artifactURI\",\"type\":\"string\"},{\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"postBounty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bountyGuid\",\"type\":\"uint128\"}],\"name\":\"getNumberOfAssertions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bountyGuids\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfBounties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOUNTY_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOUNTY_AMOUNT_MINIMUM\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"name\":\"verdicts\",\"type\":\"uint256\"},{\"name\":\"bid\",\"type\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"postAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint128\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertionsByGuid\",\"outputs\":[{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"verdicts\",\"type\":\"uint256\"},{\"name\":\"bid\",\"type\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"bountiesByGuid\",\"outputs\":[{\"name\":\"guid\",\"type\":\"uint128\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"artifactHash\",\"type\":\"bytes32\"},{\"name\":\"artifactURI\",\"type\":\"string\"},{\"name\":\"expirationBlock\",\"type\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\"},{\"name\":\"verdicts\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASSERTION_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASSERTION_BID_MINIMUM\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"uint128\"},{\"name\":\"verdicts\",\"type\":\"uint256\"}],\"name\":\"settleBounty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"nectarTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guid\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"artifactHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"artifactURI\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"verdicts\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metdata\",\"type\":\"string\"}],\"name\":\"NewAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"verdicts\",\"type\":\"uint256\"}],\"name\":\"NewVerdict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BountyRegistryBin is the compiled bytecode used for deploying new contracts.
const BountyRegistryBin = `60606040526000805460a060020a60ff0219169055341561001f57600080fd5b6040516020806111658339810160405280805160008054600160a060020a03338116600160a060020a031992831681179093556001805483169093179092556002805492909316911617905550506110e98061007c6000396000f3006060604052600436106100d75763ffffffff60e060020a60003504166324e68f1181146100dc5780633f4ba83a146101145780634ee29ec5146101275780635c975abb14610158578063844b6dfc1461017f5780638456cb59146101b15780638836f3ef146101c45780638da5cb5b146101d7578063b5eabb1e14610206578063b769ccbf14610219578063c638fe101461022c578063d3b922431461025f578063d4c4e7a914610329578063d63b5a9414610219578063eceb122414610219578063f2fde38b14610419578063f57316fb14610438575b600080fd5b34156100e757600080fd5b610112600480356001608060020a03169060248035916044359160643590810191013560843561045a565b005b341561011f57600080fd5b610112610806565b341561013257600080fd5b6101466001608060020a0360043516610885565b60405190815260200160405180910390f35b341561016357600080fd5b61016b6108d1565b604051901515815260200160405180910390f35b341561018a57600080fd5b6101956004356108e1565b6040516001608060020a03909116815260200160405180910390f35b34156101bc57600080fd5b61011261091c565b34156101cf57600080fd5b6101466109a0565b34156101e257600080fd5b6101ea6109a7565b604051600160a060020a03909116815260200160405180910390f35b341561021157600080fd5b6101466109b6565b341561022457600080fd5b6101466109bb565b341561023757600080fd5b610112600480356001608060020a0316906024803591604435916064359081019101356109c0565b341561026a57600080fd5b6102816001608060020a0360043516602435610c67565b604051600160a060020a038516815260208101849052604081018390526080606082018181528354600260001961010060018416150201909116049183018290529060a0830190849080156103175780601f106102ec57610100808354040283529160200191610317565b820191906000526020600020905b8154815290600101906020018083116102fa57829003601f168201915b50509550505050505060405180910390f35b341561033457600080fd5b6103486001608060020a0360043516610cb4565b6040516001608060020a0389168152600160a060020a0388166020820152604081018790526060810186905260a0810184905282151560c082015260e08101829052610100608082018181528654600260018216158402600019019091160491830182905290610120830190879080156104035780601f106103d857610100808354040283529160200191610403565b820191906000526020600020905b8154815290600101906020018083116103e657829003601f168201915b5050995050505050505050505060405180910390f35b341561042457600080fd5b610112600160a060020a0360043516610d11565b341561044357600080fd5b6101126001608060020a0360043516602435610dac565b610462610e93565b6001608060020a038716600090815260046020526040902060010154600160a060020a03161561049157600080fd5b600186101561049f57600080fd5b600083116104ac57600080fd5b600082116104b957600080fd5b600254600160a060020a03166323b872dd33306104dd8a600a63ffffffff610e7d16565b60006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561053557600080fd5b6102c65a03f1151561054657600080fd5b50505060405180519050151561055b57600080fd5b61010060405190810160405280886001608060020a0316815260200133600160a060020a031681526020018781526020018660001916815260200185858080601f01602080910402602001604051908101604052818152929190602084018383808284375050509284525050602090910190506105de844363ffffffff610e7d16565b81526000602080830182905260409283018290526001608060020a038b16825260049052209091508190815181546fffffffffffffffffffffffffffffffff19166001608060020a0391909116178155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556040820151816002015560608201516003820155608082015181600401908051610690929160200190610ee1565b5060a0820151816005015560c082015160068201805460ff191691151591909117905560e08201516007909101555060038054600181016106d18382610f5f565b916000526020600020906002918282040191900660100289909190916101000a8154816001608060020a0302191690836001608060020a03160217905550507f67b5fe7fd81ceeedc37156514979024c4a1f794ae72ac1d3933c6f28a495b169816000015182602001518360400151846060015185608001518660a001516040516001608060020a0387168152600160a060020a0386166020820152604081018590526060810184905260a0810182905260c06080820181815290820184818151815260200191508051906020019080838360005b838110156107be5780820151838201526020016107a6565b50505050905090810190601f1680156107eb5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150505050505050565b60005433600160a060020a0390811691161461082157600080fd5b60005460a060020a900460ff16151561083957600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6001608060020a038116600090815260046020526040812060010154600160a060020a031615156108b557600080fd5b506001608060020a031660009081526005602052604090205490565b60005460a060020a900460ff1681565b60038054829081106108ef57fe5b9060005260206000209060029182820401919006601002915054906101000a90046001608060020a031681565b60005433600160a060020a0390811691161461093757600080fd5b60005460a060020a900460ff161561094e57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6003545b90565b600054600160a060020a031681565b600a81565b600181565b6109c8610f98565b6001608060020a038616600090815260046020526040812060010154600160a060020a031615156109f857600080fd5b6001608060020a038716600090815260046020526040902060050154439011610a2057600080fd5b6001851015610a2e57600080fd5b600254600160a060020a03166323b872dd3330610a5289600163ffffffff610e7d16565b60006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610aaa57600080fd5b6102c65a03f11515610abb57600080fd5b505050604051805190501515610ad057600080fd5b60806040519081016040528033600160a060020a0316815260200187815260200186815260200185858080601f0160208091040260200160405190810160405281815292919060208401838380828437505050929093525050506001608060020a03881660009081526005602052604090208054919350600191808301610b578382610fcf565b600092835260209092208591600402018151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201518160020155606082015181600301908051610bbf929160200190610ee1565b5050500390507fdb43bb1430e831a459231d838fb3b8cd6ea5d8e9f276f06ac2ecbc8a9c490441878351846020015184866040015187606001516040516001608060020a0387168152600160a060020a038616602082015260408101859052606081018490526080810183905260c060a082018181529082018381815181526020019150805190602001908083836000838110156107be5780820151838201526020016107a6565b600560205281600052604060002081815481101515610c8257fe5b60009182526020909120600490910201805460018201546002830154600160a060020a03909216945092509060030184565b600460208190526000918252604090912080546001820154600283015460038401546005850154600686015460078701546001608060020a0390961697600160a060020a0390951696939592949390930192909160ff9091169088565b60005433600160a060020a03908116911614610d2c57600080fd5b600160a060020a0381161515610d4157600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6001608060020a038216600090815260046020526040902060010154600160a060020a03161515610ddc57600080fd5b6001608060020a03821660009081526004602052604090206005015443901115610e0557600080fd5b6001608060020a0382166000908152600460205260409081902060078101839055600601805460ff191660011790557f0da46f8053a8890362a6e1af70e9cbef0e3ca7881d1da3520df4e358727d49c09083908390516001608060020a03909216825260208201526040908101905180910390a15050565b600082820183811015610e8c57fe5b9392505050565b6101006040519081016040908152600080835260208301819052908201819052606082015260808101610ec4610ffb565b815260200160008152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f2257805160ff1916838001178555610f4f565b82800160010185558215610f4f579182015b82811115610f4f578251825591602001919060010190610f34565b50610f5b92915061100d565b5090565b815481835581811511610f93576001016002900481600101600290048360005260206000209182019101610f93919061100d565b505050565b6080604051908101604052806000600160a060020a031681526020016000815260200160008152602001610fca610ffb565b905290565b815481835581811511610f9357600402816004028360005260206000209182019101610f939190611027565b60206040519081016040526000815290565b6109a491905b80821115610f5b5760008155600101611013565b6109a491905b80821115610f5b57805473ffffffffffffffffffffffffffffffffffffffff191681556000600182018190556002820181905561106d6003830182611076565b5060040161102d565b50805460018160011615610100020316600290046000825580601f1061109c57506110ba565b601f0160209004906000526020600020908101906110ba919061100d565b505600a165627a7a72305820ab9da8c25fdfca40061cf3cf3b22776d0b703a0cfa032438e3cf94b6eb076a9f0029`

// DeployBountyRegistry deploys a new Ethereum contract, binding an instance of BountyRegistry to it.
func DeployBountyRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, nectarTokenAddr common.Address) (common.Address, *types.Transaction, *BountyRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(BountyRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BountyRegistryBin), backend, nectarTokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BountyRegistry{BountyRegistryCaller: BountyRegistryCaller{contract: contract}, BountyRegistryTransactor: BountyRegistryTransactor{contract: contract}}, nil
}

// BountyRegistry is an auto generated Go binding around an Ethereum contract.
type BountyRegistry struct {
	BountyRegistryCaller     // Read-only binding to the contract
	BountyRegistryTransactor // Write-only binding to the contract
}

// BountyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BountyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BountyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BountyRegistrySession struct {
	Contract     *BountyRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BountyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BountyRegistryCallerSession struct {
	Contract *BountyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BountyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BountyRegistryTransactorSession struct {
	Contract     *BountyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BountyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BountyRegistryRaw struct {
	Contract *BountyRegistry // Generic contract binding to access the raw methods on
}

// BountyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BountyRegistryCallerRaw struct {
	Contract *BountyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BountyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BountyRegistryTransactorRaw struct {
	Contract *BountyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBountyRegistry creates a new instance of BountyRegistry, bound to a specific deployed contract.
func NewBountyRegistry(address common.Address, backend bind.ContractBackend) (*BountyRegistry, error) {
	contract, err := bindBountyRegistry(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BountyRegistry{BountyRegistryCaller: BountyRegistryCaller{contract: contract}, BountyRegistryTransactor: BountyRegistryTransactor{contract: contract}}, nil
}

// NewBountyRegistryCaller creates a new read-only instance of BountyRegistry, bound to a specific deployed contract.
func NewBountyRegistryCaller(address common.Address, caller bind.ContractCaller) (*BountyRegistryCaller, error) {
	contract, err := bindBountyRegistry(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BountyRegistryCaller{contract: contract}, nil
}

// NewBountyRegistryTransactor creates a new write-only instance of BountyRegistry, bound to a specific deployed contract.
func NewBountyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BountyRegistryTransactor, error) {
	contract, err := bindBountyRegistry(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BountyRegistryTransactor{contract: contract}, nil
}

// bindBountyRegistry binds a generic wrapper to an already deployed contract.
func bindBountyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BountyRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BountyRegistry *BountyRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BountyRegistry.Contract.BountyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BountyRegistry *BountyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyRegistry.Contract.BountyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BountyRegistry *BountyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BountyRegistry.Contract.BountyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BountyRegistry *BountyRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BountyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BountyRegistry *BountyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BountyRegistry *BountyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BountyRegistry.Contract.contract.Transact(opts, method, params...)
}

// ASSERTION_BID_MINIMUM is a free data retrieval call binding the contract method 0xeceb1224.
//
// Solidity: function ASSERTION_BID_MINIMUM() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCaller) ASSERTION_BID_MINIMUM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "ASSERTION_BID_MINIMUM")
	return *ret0, err
}

// ASSERTION_BID_MINIMUM is a free data retrieval call binding the contract method 0xeceb1224.
//
// Solidity: function ASSERTION_BID_MINIMUM() constant returns(uint256)
func (_BountyRegistry *BountyRegistrySession) ASSERTION_BID_MINIMUM() (*big.Int, error) {
	return _BountyRegistry.Contract.ASSERTION_BID_MINIMUM(&_BountyRegistry.CallOpts)
}

// ASSERTION_BID_MINIMUM is a free data retrieval call binding the contract method 0xeceb1224.
//
// Solidity: function ASSERTION_BID_MINIMUM() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCallerSession) ASSERTION_BID_MINIMUM() (*big.Int, error) {
	return _BountyRegistry.Contract.ASSERTION_BID_MINIMUM(&_BountyRegistry.CallOpts)
}

// ASSERTION_FEE is a free data retrieval call binding the contract method 0xd63b5a94.
//
// Solidity: function ASSERTION_FEE() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCaller) ASSERTION_FEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "ASSERTION_FEE")
	return *ret0, err
}

// ASSERTION_FEE is a free data retrieval call binding the contract method 0xd63b5a94.
//
// Solidity: function ASSERTION_FEE() constant returns(uint256)
func (_BountyRegistry *BountyRegistrySession) ASSERTION_FEE() (*big.Int, error) {
	return _BountyRegistry.Contract.ASSERTION_FEE(&_BountyRegistry.CallOpts)
}

// ASSERTION_FEE is a free data retrieval call binding the contract method 0xd63b5a94.
//
// Solidity: function ASSERTION_FEE() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCallerSession) ASSERTION_FEE() (*big.Int, error) {
	return _BountyRegistry.Contract.ASSERTION_FEE(&_BountyRegistry.CallOpts)
}

// BOUNTY_AMOUNT_MINIMUM is a free data retrieval call binding the contract method 0xb769ccbf.
//
// Solidity: function BOUNTY_AMOUNT_MINIMUM() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCaller) BOUNTY_AMOUNT_MINIMUM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "BOUNTY_AMOUNT_MINIMUM")
	return *ret0, err
}

// BOUNTY_AMOUNT_MINIMUM is a free data retrieval call binding the contract method 0xb769ccbf.
//
// Solidity: function BOUNTY_AMOUNT_MINIMUM() constant returns(uint256)
func (_BountyRegistry *BountyRegistrySession) BOUNTY_AMOUNT_MINIMUM() (*big.Int, error) {
	return _BountyRegistry.Contract.BOUNTY_AMOUNT_MINIMUM(&_BountyRegistry.CallOpts)
}

// BOUNTY_AMOUNT_MINIMUM is a free data retrieval call binding the contract method 0xb769ccbf.
//
// Solidity: function BOUNTY_AMOUNT_MINIMUM() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCallerSession) BOUNTY_AMOUNT_MINIMUM() (*big.Int, error) {
	return _BountyRegistry.Contract.BOUNTY_AMOUNT_MINIMUM(&_BountyRegistry.CallOpts)
}

// BOUNTY_FEE is a free data retrieval call binding the contract method 0xb5eabb1e.
//
// Solidity: function BOUNTY_FEE() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCaller) BOUNTY_FEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "BOUNTY_FEE")
	return *ret0, err
}

// BOUNTY_FEE is a free data retrieval call binding the contract method 0xb5eabb1e.
//
// Solidity: function BOUNTY_FEE() constant returns(uint256)
func (_BountyRegistry *BountyRegistrySession) BOUNTY_FEE() (*big.Int, error) {
	return _BountyRegistry.Contract.BOUNTY_FEE(&_BountyRegistry.CallOpts)
}

// BOUNTY_FEE is a free data retrieval call binding the contract method 0xb5eabb1e.
//
// Solidity: function BOUNTY_FEE() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCallerSession) BOUNTY_FEE() (*big.Int, error) {
	return _BountyRegistry.Contract.BOUNTY_FEE(&_BountyRegistry.CallOpts)
}

// AssertionsByGuid is a free data retrieval call binding the contract method 0xd3b92243.
//
// Solidity: function assertionsByGuid( uint128,  uint256) constant returns(author address, verdicts uint256, bid uint256, metadata string)
func (_BountyRegistry *BountyRegistryCaller) AssertionsByGuid(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Author   common.Address
	Verdicts *big.Int
	Bid      *big.Int
	Metadata string
}, error) {
	ret := new(struct {
		Author   common.Address
		Verdicts *big.Int
		Bid      *big.Int
		Metadata string
	})
	out := ret
	err := _BountyRegistry.contract.Call(opts, out, "assertionsByGuid", arg0, arg1)
	return *ret, err
}

// AssertionsByGuid is a free data retrieval call binding the contract method 0xd3b92243.
//
// Solidity: function assertionsByGuid( uint128,  uint256) constant returns(author address, verdicts uint256, bid uint256, metadata string)
func (_BountyRegistry *BountyRegistrySession) AssertionsByGuid(arg0 *big.Int, arg1 *big.Int) (struct {
	Author   common.Address
	Verdicts *big.Int
	Bid      *big.Int
	Metadata string
}, error) {
	return _BountyRegistry.Contract.AssertionsByGuid(&_BountyRegistry.CallOpts, arg0, arg1)
}

// AssertionsByGuid is a free data retrieval call binding the contract method 0xd3b92243.
//
// Solidity: function assertionsByGuid( uint128,  uint256) constant returns(author address, verdicts uint256, bid uint256, metadata string)
func (_BountyRegistry *BountyRegistryCallerSession) AssertionsByGuid(arg0 *big.Int, arg1 *big.Int) (struct {
	Author   common.Address
	Verdicts *big.Int
	Bid      *big.Int
	Metadata string
}, error) {
	return _BountyRegistry.Contract.AssertionsByGuid(&_BountyRegistry.CallOpts, arg0, arg1)
}

// BountiesByGuid is a free data retrieval call binding the contract method 0xd4c4e7a9.
//
// Solidity: function bountiesByGuid( uint128) constant returns(guid uint128, author address, amount uint256, artifactHash bytes32, artifactURI string, expirationBlock uint256, resolved bool, verdicts uint256)
func (_BountyRegistry *BountyRegistryCaller) BountiesByGuid(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Resolved        bool
	Verdicts        *big.Int
}, error) {
	ret := new(struct {
		Guid            *big.Int
		Author          common.Address
		Amount          *big.Int
		ArtifactHash    [32]byte
		ArtifactURI     string
		ExpirationBlock *big.Int
		Resolved        bool
		Verdicts        *big.Int
	})
	out := ret
	err := _BountyRegistry.contract.Call(opts, out, "bountiesByGuid", arg0)
	return *ret, err
}

// BountiesByGuid is a free data retrieval call binding the contract method 0xd4c4e7a9.
//
// Solidity: function bountiesByGuid( uint128) constant returns(guid uint128, author address, amount uint256, artifactHash bytes32, artifactURI string, expirationBlock uint256, resolved bool, verdicts uint256)
func (_BountyRegistry *BountyRegistrySession) BountiesByGuid(arg0 *big.Int) (struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Resolved        bool
	Verdicts        *big.Int
}, error) {
	return _BountyRegistry.Contract.BountiesByGuid(&_BountyRegistry.CallOpts, arg0)
}

// BountiesByGuid is a free data retrieval call binding the contract method 0xd4c4e7a9.
//
// Solidity: function bountiesByGuid( uint128) constant returns(guid uint128, author address, amount uint256, artifactHash bytes32, artifactURI string, expirationBlock uint256, resolved bool, verdicts uint256)
func (_BountyRegistry *BountyRegistryCallerSession) BountiesByGuid(arg0 *big.Int) (struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Resolved        bool
	Verdicts        *big.Int
}, error) {
	return _BountyRegistry.Contract.BountiesByGuid(&_BountyRegistry.CallOpts, arg0)
}

// BountyGuids is a free data retrieval call binding the contract method 0x844b6dfc.
//
// Solidity: function bountyGuids( uint256) constant returns(uint128)
func (_BountyRegistry *BountyRegistryCaller) BountyGuids(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "bountyGuids", arg0)
	return *ret0, err
}

// BountyGuids is a free data retrieval call binding the contract method 0x844b6dfc.
//
// Solidity: function bountyGuids( uint256) constant returns(uint128)
func (_BountyRegistry *BountyRegistrySession) BountyGuids(arg0 *big.Int) (*big.Int, error) {
	return _BountyRegistry.Contract.BountyGuids(&_BountyRegistry.CallOpts, arg0)
}

// BountyGuids is a free data retrieval call binding the contract method 0x844b6dfc.
//
// Solidity: function bountyGuids( uint256) constant returns(uint128)
func (_BountyRegistry *BountyRegistryCallerSession) BountyGuids(arg0 *big.Int) (*big.Int, error) {
	return _BountyRegistry.Contract.BountyGuids(&_BountyRegistry.CallOpts, arg0)
}

// GetNumberOfAssertions is a free data retrieval call binding the contract method 0x4ee29ec5.
//
// Solidity: function getNumberOfAssertions(bountyGuid uint128) constant returns(uint256)
func (_BountyRegistry *BountyRegistryCaller) GetNumberOfAssertions(opts *bind.CallOpts, bountyGuid *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "getNumberOfAssertions", bountyGuid)
	return *ret0, err
}

// GetNumberOfAssertions is a free data retrieval call binding the contract method 0x4ee29ec5.
//
// Solidity: function getNumberOfAssertions(bountyGuid uint128) constant returns(uint256)
func (_BountyRegistry *BountyRegistrySession) GetNumberOfAssertions(bountyGuid *big.Int) (*big.Int, error) {
	return _BountyRegistry.Contract.GetNumberOfAssertions(&_BountyRegistry.CallOpts, bountyGuid)
}

// GetNumberOfAssertions is a free data retrieval call binding the contract method 0x4ee29ec5.
//
// Solidity: function getNumberOfAssertions(bountyGuid uint128) constant returns(uint256)
func (_BountyRegistry *BountyRegistryCallerSession) GetNumberOfAssertions(bountyGuid *big.Int) (*big.Int, error) {
	return _BountyRegistry.Contract.GetNumberOfAssertions(&_BountyRegistry.CallOpts, bountyGuid)
}

// GetNumberOfBounties is a free data retrieval call binding the contract method 0x8836f3ef.
//
// Solidity: function getNumberOfBounties() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCaller) GetNumberOfBounties(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "getNumberOfBounties")
	return *ret0, err
}

// GetNumberOfBounties is a free data retrieval call binding the contract method 0x8836f3ef.
//
// Solidity: function getNumberOfBounties() constant returns(uint256)
func (_BountyRegistry *BountyRegistrySession) GetNumberOfBounties() (*big.Int, error) {
	return _BountyRegistry.Contract.GetNumberOfBounties(&_BountyRegistry.CallOpts)
}

// GetNumberOfBounties is a free data retrieval call binding the contract method 0x8836f3ef.
//
// Solidity: function getNumberOfBounties() constant returns(uint256)
func (_BountyRegistry *BountyRegistryCallerSession) GetNumberOfBounties() (*big.Int, error) {
	return _BountyRegistry.Contract.GetNumberOfBounties(&_BountyRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BountyRegistry *BountyRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BountyRegistry *BountyRegistrySession) Owner() (common.Address, error) {
	return _BountyRegistry.Contract.Owner(&_BountyRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BountyRegistry *BountyRegistryCallerSession) Owner() (common.Address, error) {
	return _BountyRegistry.Contract.Owner(&_BountyRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BountyRegistry *BountyRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BountyRegistry *BountyRegistrySession) Paused() (bool, error) {
	return _BountyRegistry.Contract.Paused(&_BountyRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BountyRegistry *BountyRegistryCallerSession) Paused() (bool, error) {
	return _BountyRegistry.Contract.Paused(&_BountyRegistry.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BountyRegistry *BountyRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BountyRegistry *BountyRegistrySession) Pause() (*types.Transaction, error) {
	return _BountyRegistry.Contract.Pause(&_BountyRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BountyRegistry *BountyRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _BountyRegistry.Contract.Pause(&_BountyRegistry.TransactOpts)
}

// PostAssertion is a paid mutator transaction binding the contract method 0xc638fe10.
//
// Solidity: function postAssertion(bountyGuid uint128, verdicts uint256, bid uint256, metadata string) returns()
func (_BountyRegistry *BountyRegistryTransactor) PostAssertion(opts *bind.TransactOpts, bountyGuid *big.Int, verdicts *big.Int, bid *big.Int, metadata string) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "postAssertion", bountyGuid, verdicts, bid, metadata)
}

// PostAssertion is a paid mutator transaction binding the contract method 0xc638fe10.
//
// Solidity: function postAssertion(bountyGuid uint128, verdicts uint256, bid uint256, metadata string) returns()
func (_BountyRegistry *BountyRegistrySession) PostAssertion(bountyGuid *big.Int, verdicts *big.Int, bid *big.Int, metadata string) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostAssertion(&_BountyRegistry.TransactOpts, bountyGuid, verdicts, bid, metadata)
}

// PostAssertion is a paid mutator transaction binding the contract method 0xc638fe10.
//
// Solidity: function postAssertion(bountyGuid uint128, verdicts uint256, bid uint256, metadata string) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) PostAssertion(bountyGuid *big.Int, verdicts *big.Int, bid *big.Int, metadata string) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostAssertion(&_BountyRegistry.TransactOpts, bountyGuid, verdicts, bid, metadata)
}

// PostBounty is a paid mutator transaction binding the contract method 0x24e68f11.
//
// Solidity: function postBounty(guid uint128, amount uint256, artifactHash bytes32, artifactURI string, durationBlocks uint256) returns()
func (_BountyRegistry *BountyRegistryTransactor) PostBounty(opts *bind.TransactOpts, guid *big.Int, amount *big.Int, artifactHash [32]byte, artifactURI string, durationBlocks *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "postBounty", guid, amount, artifactHash, artifactURI, durationBlocks)
}

// PostBounty is a paid mutator transaction binding the contract method 0x24e68f11.
//
// Solidity: function postBounty(guid uint128, amount uint256, artifactHash bytes32, artifactURI string, durationBlocks uint256) returns()
func (_BountyRegistry *BountyRegistrySession) PostBounty(guid *big.Int, amount *big.Int, artifactHash [32]byte, artifactURI string, durationBlocks *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostBounty(&_BountyRegistry.TransactOpts, guid, amount, artifactHash, artifactURI, durationBlocks)
}

// PostBounty is a paid mutator transaction binding the contract method 0x24e68f11.
//
// Solidity: function postBounty(guid uint128, amount uint256, artifactHash bytes32, artifactURI string, durationBlocks uint256) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) PostBounty(guid *big.Int, amount *big.Int, artifactHash [32]byte, artifactURI string, durationBlocks *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostBounty(&_BountyRegistry.TransactOpts, guid, amount, artifactHash, artifactURI, durationBlocks)
}

// SettleBounty is a paid mutator transaction binding the contract method 0xf57316fb.
//
// Solidity: function settleBounty(guid uint128, verdicts uint256) returns()
func (_BountyRegistry *BountyRegistryTransactor) SettleBounty(opts *bind.TransactOpts, guid *big.Int, verdicts *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "settleBounty", guid, verdicts)
}

// SettleBounty is a paid mutator transaction binding the contract method 0xf57316fb.
//
// Solidity: function settleBounty(guid uint128, verdicts uint256) returns()
func (_BountyRegistry *BountyRegistrySession) SettleBounty(guid *big.Int, verdicts *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.SettleBounty(&_BountyRegistry.TransactOpts, guid, verdicts)
}

// SettleBounty is a paid mutator transaction binding the contract method 0xf57316fb.
//
// Solidity: function settleBounty(guid uint128, verdicts uint256) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) SettleBounty(guid *big.Int, verdicts *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.SettleBounty(&_BountyRegistry.TransactOpts, guid, verdicts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BountyRegistry *BountyRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BountyRegistry *BountyRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BountyRegistry.Contract.TransferOwnership(&_BountyRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BountyRegistry.Contract.TransferOwnership(&_BountyRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BountyRegistry *BountyRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BountyRegistry *BountyRegistrySession) Unpause() (*types.Transaction, error) {
	return _BountyRegistry.Contract.Unpause(&_BountyRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BountyRegistry *BountyRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _BountyRegistry.Contract.Unpause(&_BountyRegistry.TransactOpts)
}
