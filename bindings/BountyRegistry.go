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

// BountyRegistryABI is the input ABI used to generate the binding from.
const BountyRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"uint128\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"artifactURI\",\"type\":\"string\"},{\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"postBounty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"arbiter\",\"type\":\"address\"}],\"name\":\"removeArbiter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"name\":\"bid\",\"type\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\"},{\"name\":\"verdicts\",\"type\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"postAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bountyGuid\",\"type\":\"uint128\"}],\"name\":\"getNumberOfAssertions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"arbiters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bountyGuids\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfBounties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newArbiter\",\"type\":\"address\"}],\"name\":\"addArbiter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOUNTY_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOUNTY_AMOUNT_MINIMUM\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint128\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertionsByGuid\",\"outputs\":[{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"bid\",\"type\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\"},{\"name\":\"verdicts\",\"type\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"bountiesByGuid\",\"outputs\":[{\"name\":\"guid\",\"type\":\"uint128\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"artifactURI\",\"type\":\"string\"},{\"name\":\"expirationBlock\",\"type\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\"},{\"name\":\"verdicts\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASSERTION_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASSERTION_BID_MINIMUM\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"name\":\"verdicts\",\"type\":\"uint256\"}],\"name\":\"settleBounty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"nectarTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guid\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"artifactURI\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"mask\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"verdicts\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metdata\",\"type\":\"string\"}],\"name\":\"NewAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bountyGuid\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"verdicts\",\"type\":\"uint256\"}],\"name\":\"NewVerdict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BountyRegistryBin is the compiled bytecode used for deploying new contracts.
const BountyRegistryBin = `60606040526000805460a060020a60ff0219169055341561001f57600080fd5b6040516020806119f88339810160405280805160008054600160a060020a03338116600160a060020a0319928316811790935560018054831690931790925560028054929093169116179055505061197c8061007c6000396000f3006060604052600436106100f75763ffffffff60e060020a6000350416623b654881146100fc5780633487e08c146101305780633f4ba83a1461014f57806348aa2e93146101625780634ee29ec5146101995780635c975abb146101ca5780637bf2bb10146101f1578063844b6dfc146102105780638456cb59146102425780638836f3ef146102555780638da5cb5b14610268578063b538d3bc14610297578063b5eabb1e146102b6578063b769ccbf146102b6578063d3b92243146102c9578063d4c4e7a914610391578063d63b5a94146102b6578063eceb1224146102b6578063f2fde38b14610468578063f57316fb14610487575b600080fd5b341561010757600080fd5b61012e600480356001608060020a03169060248035916044359182019101356064356104a9565b005b341561013b57600080fd5b61012e600160a060020a036004351661083d565b341561015a57600080fd5b61012e610890565b341561016d57600080fd5b61012e600480356001608060020a0316906024803591604435916064359160843591820191013561090f565b34156101a457600080fd5b6101b86001608060020a0360043516610c35565b60405190815260200160405180910390f35b34156101d557600080fd5b6101dd610c81565b604051901515815260200160405180910390f35b34156101fc57600080fd5b6101dd600160a060020a0360043516610c91565b341561021b57600080fd5b610226600435610ca6565b6040516001608060020a03909116815260200160405180910390f35b341561024d57600080fd5b61012e610ce1565b341561026057600080fd5b6101b8610d65565b341561027357600080fd5b61027b610d6c565b604051600160a060020a03909116815260200160405180910390f35b34156102a257600080fd5b61012e600160a060020a0360043516610d7b565b34156102c157600080fd5b6101b8610e0c565b34156102d457600080fd5b6102eb6001608060020a0360043516602435610e17565b6040518086600160a060020a0316600160a060020a0316815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561035257808201518382015260200161033a565b50505050905090810190601f16801561037f5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561039c57600080fd5b6103b06001608060020a0360043516610f14565b6040516001608060020a0388168152600160a060020a0387166020820152604081018690526080810184905282151560a082015260c0810182905260e06060820181815290820186818151815260200191508051906020019080838360005b8381101561042757808201518382015260200161040f565b50505050905090810190601f1680156104545780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b341561047357600080fd5b61012e600160a060020a036004351661100a565b341561049257600080fd5b61012e6001608060020a03600435166024356110a5565b6104b1611720565b60005460a060020a900460ff16156104c857600080fd5b6001608060020a038616600090815260046020526040902060010154600160a060020a0316156104f757600080fd5b66de0b6b3a76400085101561050b57600080fd5b6000831161051857600080fd5b6000821161052557600080fd5b600254600160a060020a03166323b872dd333061054f8966de0b6b3a76400063ffffffff6116b216565b60405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561059e57600080fd5b5af115156105ab57600080fd5b5050506040518051905015156105c057600080fd5b60e060405190810160405280876001608060020a0316815260200133600160a060020a0316815260200186815260200185858080601f0160208091040260200160405190810160405281815292919060208401838380828437505050928452505060209091019050610638844363ffffffff6116b216565b81526000602080830182905260409283018290526001608060020a038a16825260049052209091508190815181546fffffffffffffffffffffffffffffffff19166001608060020a0391909116178155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015181600201556060820151816003019080516106e0929160200190611766565b506080820151816004015560a082015160058201805460ff191691151591909117905560c082015160069091015550600380546001810161072183826117e4565b50600091825260209091206002820401805460019092166010026101000a6001608060020a0381810219909316928916029190911790557f1836d68e5bbde3ba859b17bd89128c486fd2c707b6791706b6efa78219d85062815182602001518360400151846060015185608001516040516001608060020a0386168152600160a060020a0385166020820152604081018490526080810182905260a06060820181815290820184818151815260200191508051906020019080838360005b838110156107f75780820151838201526020016107df565b50505050905090810190601f1680156108245780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1505050505050565b60005460a060020a900460ff161561085457600080fd5b60005433600160a060020a0390811691161461086f57600080fd5b600160a060020a03166000908152600660205260409020805460ff19169055565b60005433600160a060020a039081169116146108ab57600080fd5b60005460a060020a900460ff1615156108c357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b61091761181d565b6000805460a060020a900460ff161561092f57600080fd5b6001608060020a038816600090815260046020526040902060010154600160a060020a0316151561095f57600080fd5b66de0b6b3a76400087101561097357600080fd5b6001608060020a0388166000908152600460208190526040909120015443901161099c57600080fd5b600254600160a060020a03166323b872dd33306109c68b66de0b6b3a76400063ffffffff6116b216565b60405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610a1557600080fd5b5af11515610a2257600080fd5b505050604051805190501515610a3757600080fd5b60a06040519081016040528033600160a060020a0316815260200188815260200187815260200186815260200185858080601f0160208091040260200160405190810160405281815292919060208401838380828437505050929093525050506001608060020a03891660009081526005602052604090208054919350600191808301610ac4838261185b565b600092835260209092208591600502018151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101556040820151816002015560608201518160030155608082015181600401908051610b36929160200190611766565b5050500390507f6485719c794bff0febee47b18180d293415dc05661685490274f4287bc25d32c8883518385602001518660400151876060015188608001516040516001608060020a0388168152600160a060020a038716602082015260408101869052606081018590526080810184905260a0810183905260e060c0820181815290820183818151815260200191508051906020019080838360005b83811015610beb578082015183820152602001610bd3565b50505050905090810190601f168015610c185780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15050505050505050565b6001608060020a038116600090815260046020526040812060010154600160a060020a03161515610c6557600080fd5b506001608060020a031660009081526005602052604090205490565b60005460a060020a900460ff1681565b60066020526000908152604090205460ff1681565b6003805482908110610cb457fe5b9060005260206000209060029182820401919006601002915054906101000a90046001608060020a031681565b60005433600160a060020a03908116911614610cfc57600080fd5b60005460a060020a900460ff1615610d1357600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6003545b90565b600054600160a060020a031681565b60005460a060020a900460ff1615610d9257600080fd5b60005433600160a060020a03908116911614610dad57600080fd5b600160a060020a0381161515610dc257600080fd5b600160a060020a03811660009081526006602052604090205460ff1615610de857600080fd5b600160a060020a03166000908152600660205260409020805460ff19166001179055565b66de0b6b3a76400081565b600560205281600052604060002081815481101515610e3257fe5b9060005260206000209060050201600091509150508060000160009054906101000a9004600160a060020a031690806001015490806002015490806003015490806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f0a5780601f10610edf57610100808354040283529160200191610f0a565b820191906000526020600020905b815481529060010190602001808311610eed57829003601f168201915b5050505050905085565b60046020908152600091825260409182902080546001808301546002808501546003860180546001608060020a0390961698600160a060020a03909416979196959094600019908216156101000201169190910491601f8301829004820290910190519081016040528092919081815260200182805460018160011615610100020316600290048015610fe85780601f10610fbd57610100808354040283529160200191610fe8565b820191906000526020600020905b815481529060010190602001808311610fcb57829003601f168201915b50505050600483015460058401546006909401549293909260ff909116915087565b60005433600160a060020a0390811691161461102557600080fd5b600160a060020a038116151561103a57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6110ad611720565b6110b5611887565b600160a060020a0333166000908152600660205260408120548190819081908190819081908190819060ff1615156110ec57600080fd5b60005460a060020a900460ff161561110357600080fd5b6001608060020a038d16600090815260046020526040908190209060e09051908101604090815282546001608060020a03168252600180840154600160a060020a03166020808501919091526002808601548486015260038601805495969560608801959194600019908216156101000201169190910491601f8301819004810201905190810160405280929190818152602001828054600181600116156101000203166002900480156111f85780601f106111cd576101008083540402835291602001916111f8565b820191906000526020600020905b8154815290600101906020018083116111db57829003601f168201915b50505050508152602001600482015481526020016005820160009054906101000a900460ff161515151581526020016006820154815250509a50600560008e6001608060020a03166001608060020a0316815260200190815260200160002080548060200260200160405190810160405281815291906000602084015b8282101561139557600084815260209020600583020160a060405190810160405290816000820160009054906101000a9004600160a060020a0316600160a060020a0316600160a060020a03168152602001600182015481526020016002820154815260200160038201548152602001600482018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561137d5780601f106113525761010080835404028352916020019161137d565b820191906000526020600020905b81548152906001019060200180831161136057829003601f168201915b50505050508152505081526020019060010190611275565b5050505099506000600160a060020a03168b60200151600160a060020a031614156113bf57600080fd5b438b6080015111156113d057600080fd5b60c08b018c9052600160a08c015260009850895197508a60400151965061141b61140766de0b6b3a7640008a63ffffffff6116cc16565b66de0b6b3a7640009063ffffffff6116b216565b955060009450600098505b87891015611499578b8a8a8151811061143b57fe5b90602001906020020151606001511461148e576114788a8a8151811061145d57fe5b9060200190602002015160200151889063ffffffff6116b216565b965061148b85600163ffffffff6116b216565b94505b600190980197611426565b6114ba60016114ae8a8863ffffffff6116f716565b9063ffffffff6116b216565b93506114cc878563ffffffff61170916565b925083878115156114d957fe5b06915060009050600098505b878910156115d4578b8a8a815181106114fa57fe5b906020019060200201516060015114156115c957611538838b8b8151811061151e57fe5b90602001906020020151602001519063ffffffff6116b216565b600254909150600160a060020a031663a9059cbb8b8b8151811061155857fe5b90602001906020020151518360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156115a757600080fd5b5af115156115b457600080fd5b5050506040518051905015156115c957600080fd5b6001909801976114e5565b600254600160a060020a031663a9059cbb336115fa856114ae888c63ffffffff6116b216565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561163d57600080fd5b5af1151561164a57600080fd5b50505060405180519050151561165f57600080fd5b7f0da46f8053a8890362a6e1af70e9cbef0e3ca7881d1da3520df4e358727d49c08d8d6040516001608060020a03909216825260208201526040908101905180910390a150505050505050505050505050565b6000828201838110156116c157fe5b8091505b5092915050565b6000808315156116df57600091506116c5565b508282028284828115156116ef57fe5b04146116c157fe5b60008282111561170357fe5b50900390565b600080828481151561171757fe5b04949350505050565b60e060405190810160409081526000808352602083018190529082015260608101611749611887565b815260200160008152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117a757805160ff19168380011785556117d4565b828001600101855582156117d4579182015b828111156117d45782518255916020019190600101906117b9565b506117e0929150611899565b5090565b8154818355818115116118185760010160029004816001016002900483600052602060002091820191016118189190611899565b505050565b60a0604051908101604052806000600160a060020a03168152602001600081526020016000815260200160008152602001611856611887565b905290565b8154818355818115116118185760050281600502836000526020600020918201910161181891906118b3565b60206040519081016040526000815290565b610d6991905b808211156117e0576000815560010161189f565b610d6991905b808211156117e057805473ffffffffffffffffffffffffffffffffffffffff1916815560006001820181905560028201819055600382018190556119006004830182611909565b506005016118b9565b50805460018160011615610100020316600290046000825580601f1061192f575061194d565b601f01602090049060005260206000209081019061194d9190611899565b505600a165627a7a723058204f8c0d865286caae0e4329b39912826c488322f16cbf018daf63e280411574420029`

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
	return address, tx, &BountyRegistry{BountyRegistryCaller: BountyRegistryCaller{contract: contract}, BountyRegistryTransactor: BountyRegistryTransactor{contract: contract}, BountyRegistryFilterer: BountyRegistryFilterer{contract: contract}}, nil
}

// BountyRegistry is an auto generated Go binding around an Ethereum contract.
type BountyRegistry struct {
	BountyRegistryCaller     // Read-only binding to the contract
	BountyRegistryTransactor // Write-only binding to the contract
	BountyRegistryFilterer   // Log filterer for contract events
}

// BountyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BountyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BountyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BountyRegistryFilterer struct {
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
	contract, err := bindBountyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BountyRegistry{BountyRegistryCaller: BountyRegistryCaller{contract: contract}, BountyRegistryTransactor: BountyRegistryTransactor{contract: contract}, BountyRegistryFilterer: BountyRegistryFilterer{contract: contract}}, nil
}

// NewBountyRegistryCaller creates a new read-only instance of BountyRegistry, bound to a specific deployed contract.
func NewBountyRegistryCaller(address common.Address, caller bind.ContractCaller) (*BountyRegistryCaller, error) {
	contract, err := bindBountyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BountyRegistryCaller{contract: contract}, nil
}

// NewBountyRegistryTransactor creates a new write-only instance of BountyRegistry, bound to a specific deployed contract.
func NewBountyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BountyRegistryTransactor, error) {
	contract, err := bindBountyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BountyRegistryTransactor{contract: contract}, nil
}

// NewBountyRegistryFilterer creates a new log filterer instance of BountyRegistry, bound to a specific deployed contract.
func NewBountyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BountyRegistryFilterer, error) {
	contract, err := bindBountyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BountyRegistryFilterer{contract: contract}, nil
}

// bindBountyRegistry binds a generic wrapper to an already deployed contract.
func bindBountyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BountyRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// Arbiters is a free data retrieval call binding the contract method 0x7bf2bb10.
//
// Solidity: function arbiters( address) constant returns(bool)
func (_BountyRegistry *BountyRegistryCaller) Arbiters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BountyRegistry.contract.Call(opts, out, "arbiters", arg0)
	return *ret0, err
}

// Arbiters is a free data retrieval call binding the contract method 0x7bf2bb10.
//
// Solidity: function arbiters( address) constant returns(bool)
func (_BountyRegistry *BountyRegistrySession) Arbiters(arg0 common.Address) (bool, error) {
	return _BountyRegistry.Contract.Arbiters(&_BountyRegistry.CallOpts, arg0)
}

// Arbiters is a free data retrieval call binding the contract method 0x7bf2bb10.
//
// Solidity: function arbiters( address) constant returns(bool)
func (_BountyRegistry *BountyRegistryCallerSession) Arbiters(arg0 common.Address) (bool, error) {
	return _BountyRegistry.Contract.Arbiters(&_BountyRegistry.CallOpts, arg0)
}

// AssertionsByGuid is a free data retrieval call binding the contract method 0xd3b92243.
//
// Solidity: function assertionsByGuid( uint128,  uint256) constant returns(author address, bid uint256, mask uint256, verdicts uint256, metadata string)
func (_BountyRegistry *BountyRegistryCaller) AssertionsByGuid(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Author   common.Address
	Bid      *big.Int
	Mask     *big.Int
	Verdicts *big.Int
	Metadata string
}, error) {
	ret := new(struct {
		Author   common.Address
		Bid      *big.Int
		Mask     *big.Int
		Verdicts *big.Int
		Metadata string
	})
	out := ret
	err := _BountyRegistry.contract.Call(opts, out, "assertionsByGuid", arg0, arg1)
	return *ret, err
}

// AssertionsByGuid is a free data retrieval call binding the contract method 0xd3b92243.
//
// Solidity: function assertionsByGuid( uint128,  uint256) constant returns(author address, bid uint256, mask uint256, verdicts uint256, metadata string)
func (_BountyRegistry *BountyRegistrySession) AssertionsByGuid(arg0 *big.Int, arg1 *big.Int) (struct {
	Author   common.Address
	Bid      *big.Int
	Mask     *big.Int
	Verdicts *big.Int
	Metadata string
}, error) {
	return _BountyRegistry.Contract.AssertionsByGuid(&_BountyRegistry.CallOpts, arg0, arg1)
}

// AssertionsByGuid is a free data retrieval call binding the contract method 0xd3b92243.
//
// Solidity: function assertionsByGuid( uint128,  uint256) constant returns(author address, bid uint256, mask uint256, verdicts uint256, metadata string)
func (_BountyRegistry *BountyRegistryCallerSession) AssertionsByGuid(arg0 *big.Int, arg1 *big.Int) (struct {
	Author   common.Address
	Bid      *big.Int
	Mask     *big.Int
	Verdicts *big.Int
	Metadata string
}, error) {
	return _BountyRegistry.Contract.AssertionsByGuid(&_BountyRegistry.CallOpts, arg0, arg1)
}

// BountiesByGuid is a free data retrieval call binding the contract method 0xd4c4e7a9.
//
// Solidity: function bountiesByGuid( uint128) constant returns(guid uint128, author address, amount uint256, artifactURI string, expirationBlock uint256, resolved bool, verdicts uint256)
func (_BountyRegistry *BountyRegistryCaller) BountiesByGuid(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactURI     string
	ExpirationBlock *big.Int
	Resolved        bool
	Verdicts        *big.Int
}, error) {
	ret := new(struct {
		Guid            *big.Int
		Author          common.Address
		Amount          *big.Int
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
// Solidity: function bountiesByGuid( uint128) constant returns(guid uint128, author address, amount uint256, artifactURI string, expirationBlock uint256, resolved bool, verdicts uint256)
func (_BountyRegistry *BountyRegistrySession) BountiesByGuid(arg0 *big.Int) (struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactURI     string
	ExpirationBlock *big.Int
	Resolved        bool
	Verdicts        *big.Int
}, error) {
	return _BountyRegistry.Contract.BountiesByGuid(&_BountyRegistry.CallOpts, arg0)
}

// BountiesByGuid is a free data retrieval call binding the contract method 0xd4c4e7a9.
//
// Solidity: function bountiesByGuid( uint128) constant returns(guid uint128, author address, amount uint256, artifactURI string, expirationBlock uint256, resolved bool, verdicts uint256)
func (_BountyRegistry *BountyRegistryCallerSession) BountiesByGuid(arg0 *big.Int) (struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
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

// AddArbiter is a paid mutator transaction binding the contract method 0xb538d3bc.
//
// Solidity: function addArbiter(newArbiter address) returns()
func (_BountyRegistry *BountyRegistryTransactor) AddArbiter(opts *bind.TransactOpts, newArbiter common.Address) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "addArbiter", newArbiter)
}

// AddArbiter is a paid mutator transaction binding the contract method 0xb538d3bc.
//
// Solidity: function addArbiter(newArbiter address) returns()
func (_BountyRegistry *BountyRegistrySession) AddArbiter(newArbiter common.Address) (*types.Transaction, error) {
	return _BountyRegistry.Contract.AddArbiter(&_BountyRegistry.TransactOpts, newArbiter)
}

// AddArbiter is a paid mutator transaction binding the contract method 0xb538d3bc.
//
// Solidity: function addArbiter(newArbiter address) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) AddArbiter(newArbiter common.Address) (*types.Transaction, error) {
	return _BountyRegistry.Contract.AddArbiter(&_BountyRegistry.TransactOpts, newArbiter)
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

// PostAssertion is a paid mutator transaction binding the contract method 0x48aa2e93.
//
// Solidity: function postAssertion(bountyGuid uint128, bid uint256, mask uint256, verdicts uint256, metadata string) returns()
func (_BountyRegistry *BountyRegistryTransactor) PostAssertion(opts *bind.TransactOpts, bountyGuid *big.Int, bid *big.Int, mask *big.Int, verdicts *big.Int, metadata string) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "postAssertion", bountyGuid, bid, mask, verdicts, metadata)
}

// PostAssertion is a paid mutator transaction binding the contract method 0x48aa2e93.
//
// Solidity: function postAssertion(bountyGuid uint128, bid uint256, mask uint256, verdicts uint256, metadata string) returns()
func (_BountyRegistry *BountyRegistrySession) PostAssertion(bountyGuid *big.Int, bid *big.Int, mask *big.Int, verdicts *big.Int, metadata string) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostAssertion(&_BountyRegistry.TransactOpts, bountyGuid, bid, mask, verdicts, metadata)
}

// PostAssertion is a paid mutator transaction binding the contract method 0x48aa2e93.
//
// Solidity: function postAssertion(bountyGuid uint128, bid uint256, mask uint256, verdicts uint256, metadata string) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) PostAssertion(bountyGuid *big.Int, bid *big.Int, mask *big.Int, verdicts *big.Int, metadata string) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostAssertion(&_BountyRegistry.TransactOpts, bountyGuid, bid, mask, verdicts, metadata)
}

// PostBounty is a paid mutator transaction binding the contract method 0x003b6548.
//
// Solidity: function postBounty(guid uint128, amount uint256, artifactURI string, durationBlocks uint256) returns()
func (_BountyRegistry *BountyRegistryTransactor) PostBounty(opts *bind.TransactOpts, guid *big.Int, amount *big.Int, artifactURI string, durationBlocks *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "postBounty", guid, amount, artifactURI, durationBlocks)
}

// PostBounty is a paid mutator transaction binding the contract method 0x003b6548.
//
// Solidity: function postBounty(guid uint128, amount uint256, artifactURI string, durationBlocks uint256) returns()
func (_BountyRegistry *BountyRegistrySession) PostBounty(guid *big.Int, amount *big.Int, artifactURI string, durationBlocks *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostBounty(&_BountyRegistry.TransactOpts, guid, amount, artifactURI, durationBlocks)
}

// PostBounty is a paid mutator transaction binding the contract method 0x003b6548.
//
// Solidity: function postBounty(guid uint128, amount uint256, artifactURI string, durationBlocks uint256) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) PostBounty(guid *big.Int, amount *big.Int, artifactURI string, durationBlocks *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.PostBounty(&_BountyRegistry.TransactOpts, guid, amount, artifactURI, durationBlocks)
}

// RemoveArbiter is a paid mutator transaction binding the contract method 0x3487e08c.
//
// Solidity: function removeArbiter(arbiter address) returns()
func (_BountyRegistry *BountyRegistryTransactor) RemoveArbiter(opts *bind.TransactOpts, arbiter common.Address) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "removeArbiter", arbiter)
}

// RemoveArbiter is a paid mutator transaction binding the contract method 0x3487e08c.
//
// Solidity: function removeArbiter(arbiter address) returns()
func (_BountyRegistry *BountyRegistrySession) RemoveArbiter(arbiter common.Address) (*types.Transaction, error) {
	return _BountyRegistry.Contract.RemoveArbiter(&_BountyRegistry.TransactOpts, arbiter)
}

// RemoveArbiter is a paid mutator transaction binding the contract method 0x3487e08c.
//
// Solidity: function removeArbiter(arbiter address) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) RemoveArbiter(arbiter common.Address) (*types.Transaction, error) {
	return _BountyRegistry.Contract.RemoveArbiter(&_BountyRegistry.TransactOpts, arbiter)
}

// SettleBounty is a paid mutator transaction binding the contract method 0xf57316fb.
//
// Solidity: function settleBounty(bountyGuid uint128, verdicts uint256) returns()
func (_BountyRegistry *BountyRegistryTransactor) SettleBounty(opts *bind.TransactOpts, bountyGuid *big.Int, verdicts *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.contract.Transact(opts, "settleBounty", bountyGuid, verdicts)
}

// SettleBounty is a paid mutator transaction binding the contract method 0xf57316fb.
//
// Solidity: function settleBounty(bountyGuid uint128, verdicts uint256) returns()
func (_BountyRegistry *BountyRegistrySession) SettleBounty(bountyGuid *big.Int, verdicts *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.SettleBounty(&_BountyRegistry.TransactOpts, bountyGuid, verdicts)
}

// SettleBounty is a paid mutator transaction binding the contract method 0xf57316fb.
//
// Solidity: function settleBounty(bountyGuid uint128, verdicts uint256) returns()
func (_BountyRegistry *BountyRegistryTransactorSession) SettleBounty(bountyGuid *big.Int, verdicts *big.Int) (*types.Transaction, error) {
	return _BountyRegistry.Contract.SettleBounty(&_BountyRegistry.TransactOpts, bountyGuid, verdicts)
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

// BountyRegistryNewAssertionIterator is returned from FilterNewAssertion and is used to iterate over the raw logs and unpacked data for NewAssertion events raised by the BountyRegistry contract.
type BountyRegistryNewAssertionIterator struct {
	Event *BountyRegistryNewAssertion // Event containing the contract specifics and raw log

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
func (it *BountyRegistryNewAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyRegistryNewAssertion)
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
		it.Event = new(BountyRegistryNewAssertion)
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
func (it *BountyRegistryNewAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyRegistryNewAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyRegistryNewAssertion represents a NewAssertion event raised by the BountyRegistry contract.
type BountyRegistryNewAssertion struct {
	BountyGuid *big.Int
	Author     common.Address
	Index      *big.Int
	Bid        *big.Int
	Mask       *big.Int
	Verdicts   *big.Int
	Metdata    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewAssertion is a free log retrieval operation binding the contract event 0x6485719c794bff0febee47b18180d293415dc05661685490274f4287bc25d32c.
//
// Solidity: event NewAssertion(bountyGuid uint128, author address, index uint256, bid uint256, mask uint256, verdicts uint256, metdata string)
func (_BountyRegistry *BountyRegistryFilterer) FilterNewAssertion(opts *bind.FilterOpts) (*BountyRegistryNewAssertionIterator, error) {

	logs, sub, err := _BountyRegistry.contract.FilterLogs(opts, "NewAssertion")
	if err != nil {
		return nil, err
	}
	return &BountyRegistryNewAssertionIterator{contract: _BountyRegistry.contract, event: "NewAssertion", logs: logs, sub: sub}, nil
}

// WatchNewAssertion is a free log subscription operation binding the contract event 0x6485719c794bff0febee47b18180d293415dc05661685490274f4287bc25d32c.
//
// Solidity: event NewAssertion(bountyGuid uint128, author address, index uint256, bid uint256, mask uint256, verdicts uint256, metdata string)
func (_BountyRegistry *BountyRegistryFilterer) WatchNewAssertion(opts *bind.WatchOpts, sink chan<- *BountyRegistryNewAssertion) (event.Subscription, error) {

	logs, sub, err := _BountyRegistry.contract.WatchLogs(opts, "NewAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyRegistryNewAssertion)
				if err := _BountyRegistry.contract.UnpackLog(event, "NewAssertion", log); err != nil {
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

// BountyRegistryNewBountyIterator is returned from FilterNewBounty and is used to iterate over the raw logs and unpacked data for NewBounty events raised by the BountyRegistry contract.
type BountyRegistryNewBountyIterator struct {
	Event *BountyRegistryNewBounty // Event containing the contract specifics and raw log

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
func (it *BountyRegistryNewBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyRegistryNewBounty)
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
		it.Event = new(BountyRegistryNewBounty)
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
func (it *BountyRegistryNewBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyRegistryNewBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyRegistryNewBounty represents a NewBounty event raised by the BountyRegistry contract.
type BountyRegistryNewBounty struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactURI     string
	ExpirationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewBounty is a free log retrieval operation binding the contract event 0x1836d68e5bbde3ba859b17bd89128c486fd2c707b6791706b6efa78219d85062.
//
// Solidity: event NewBounty(guid uint128, author address, amount uint256, artifactURI string, expirationBlock uint256)
func (_BountyRegistry *BountyRegistryFilterer) FilterNewBounty(opts *bind.FilterOpts) (*BountyRegistryNewBountyIterator, error) {

	logs, sub, err := _BountyRegistry.contract.FilterLogs(opts, "NewBounty")
	if err != nil {
		return nil, err
	}
	return &BountyRegistryNewBountyIterator{contract: _BountyRegistry.contract, event: "NewBounty", logs: logs, sub: sub}, nil
}

// WatchNewBounty is a free log subscription operation binding the contract event 0x1836d68e5bbde3ba859b17bd89128c486fd2c707b6791706b6efa78219d85062.
//
// Solidity: event NewBounty(guid uint128, author address, amount uint256, artifactURI string, expirationBlock uint256)
func (_BountyRegistry *BountyRegistryFilterer) WatchNewBounty(opts *bind.WatchOpts, sink chan<- *BountyRegistryNewBounty) (event.Subscription, error) {

	logs, sub, err := _BountyRegistry.contract.WatchLogs(opts, "NewBounty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyRegistryNewBounty)
				if err := _BountyRegistry.contract.UnpackLog(event, "NewBounty", log); err != nil {
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

// BountyRegistryNewVerdictIterator is returned from FilterNewVerdict and is used to iterate over the raw logs and unpacked data for NewVerdict events raised by the BountyRegistry contract.
type BountyRegistryNewVerdictIterator struct {
	Event *BountyRegistryNewVerdict // Event containing the contract specifics and raw log

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
func (it *BountyRegistryNewVerdictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyRegistryNewVerdict)
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
		it.Event = new(BountyRegistryNewVerdict)
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
func (it *BountyRegistryNewVerdictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyRegistryNewVerdictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyRegistryNewVerdict represents a NewVerdict event raised by the BountyRegistry contract.
type BountyRegistryNewVerdict struct {
	BountyGuid *big.Int
	Verdicts   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewVerdict is a free log retrieval operation binding the contract event 0x0da46f8053a8890362a6e1af70e9cbef0e3ca7881d1da3520df4e358727d49c0.
//
// Solidity: event NewVerdict(bountyGuid uint128, verdicts uint256)
func (_BountyRegistry *BountyRegistryFilterer) FilterNewVerdict(opts *bind.FilterOpts) (*BountyRegistryNewVerdictIterator, error) {

	logs, sub, err := _BountyRegistry.contract.FilterLogs(opts, "NewVerdict")
	if err != nil {
		return nil, err
	}
	return &BountyRegistryNewVerdictIterator{contract: _BountyRegistry.contract, event: "NewVerdict", logs: logs, sub: sub}, nil
}

// WatchNewVerdict is a free log subscription operation binding the contract event 0x0da46f8053a8890362a6e1af70e9cbef0e3ca7881d1da3520df4e358727d49c0.
//
// Solidity: event NewVerdict(bountyGuid uint128, verdicts uint256)
func (_BountyRegistry *BountyRegistryFilterer) WatchNewVerdict(opts *bind.WatchOpts, sink chan<- *BountyRegistryNewVerdict) (event.Subscription, error) {

	logs, sub, err := _BountyRegistry.contract.WatchLogs(opts, "NewVerdict")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyRegistryNewVerdict)
				if err := _BountyRegistry.contract.UnpackLog(event, "NewVerdict", log); err != nil {
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

// BountyRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BountyRegistry contract.
type BountyRegistryOwnershipTransferredIterator struct {
	Event *BountyRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BountyRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyRegistryOwnershipTransferred)
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
		it.Event = new(BountyRegistryOwnershipTransferred)
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
func (it *BountyRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BountyRegistry contract.
type BountyRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BountyRegistry *BountyRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BountyRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BountyRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BountyRegistryOwnershipTransferredIterator{contract: _BountyRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BountyRegistry *BountyRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BountyRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BountyRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyRegistryOwnershipTransferred)
				if err := _BountyRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BountyRegistryPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the BountyRegistry contract.
type BountyRegistryPauseIterator struct {
	Event *BountyRegistryPause // Event containing the contract specifics and raw log

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
func (it *BountyRegistryPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyRegistryPause)
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
		it.Event = new(BountyRegistryPause)
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
func (it *BountyRegistryPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyRegistryPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyRegistryPause represents a Pause event raised by the BountyRegistry contract.
type BountyRegistryPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_BountyRegistry *BountyRegistryFilterer) FilterPause(opts *bind.FilterOpts) (*BountyRegistryPauseIterator, error) {

	logs, sub, err := _BountyRegistry.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &BountyRegistryPauseIterator{contract: _BountyRegistry.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_BountyRegistry *BountyRegistryFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *BountyRegistryPause) (event.Subscription, error) {

	logs, sub, err := _BountyRegistry.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyRegistryPause)
				if err := _BountyRegistry.contract.UnpackLog(event, "Pause", log); err != nil {
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

// BountyRegistryUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the BountyRegistry contract.
type BountyRegistryUnpauseIterator struct {
	Event *BountyRegistryUnpause // Event containing the contract specifics and raw log

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
func (it *BountyRegistryUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyRegistryUnpause)
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
		it.Event = new(BountyRegistryUnpause)
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
func (it *BountyRegistryUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyRegistryUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyRegistryUnpause represents a Unpause event raised by the BountyRegistry contract.
type BountyRegistryUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_BountyRegistry *BountyRegistryFilterer) FilterUnpause(opts *bind.FilterOpts) (*BountyRegistryUnpauseIterator, error) {

	logs, sub, err := _BountyRegistry.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &BountyRegistryUnpauseIterator{contract: _BountyRegistry.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_BountyRegistry *BountyRegistryFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *BountyRegistryUnpause) (event.Subscription, error) {

	logs, sub, err := _BountyRegistry.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyRegistryUnpause)
				if err := _BountyRegistry.contract.UnpackLog(event, "Unpause", log); err != nil {
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
