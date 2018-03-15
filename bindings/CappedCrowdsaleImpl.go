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

// CappedCrowdsaleImplABI is the input ABI used to generate the binding from.
const CappedCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startTime\",\"type\":\"uint256\"},{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// CappedCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const CappedCrowdsaleImplBin = `6060604052341561000f57600080fd5b60405160a080611aa983398101604052808051919060200180519190602001805191906020018051919060200180519150819050858585854284101561005457600080fd5b8383101561006157600080fd5b6000821161006e57600080fd5b600160a060020a038116151561008357600080fd5b6100986401000000006104246100ee82021704565b60008054600160a060020a0319908116600160a060020a039384161782556001969096556002949094556004929092556003805490941691161790915581116100e057600080fd5b600655506101239350505050565b60006100f8610113565b604051809103906000f080151561010e57600080fd5b905090565b604051610a798061103083390190565b610efe806101326000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e81146100a35780633197cbb6146100c8578063355274ea146100db5780634042b66f146100ee578063521eb2731461010157806378e9792514610130578063ec8ac4d814610143578063ecb70fb714610157578063fc0c546a1461017e575b6100a133610191565b005b34156100ae57600080fd5b6100b66102c3565b60405190815260200160405180910390f35b34156100d357600080fd5b6100b66102c9565b34156100e657600080fd5b6100b66102cf565b34156100f957600080fd5b6100b66102d5565b341561010c57600080fd5b6101146102db565b604051600160a060020a03909116815260200160405180910390f35b341561013b57600080fd5b6100b66102ea565b6100a1600160a060020a0360043516610191565b341561016257600080fd5b61016a6102f0565b604051901515815260200160405180910390f35b341561018957600080fd5b610114610312565b600080600160a060020a03831615156101a957600080fd5b6101b1610321565b15156101bc57600080fd5b3491506101c882610354565b6005549091506101de908363ffffffff61037116565b600555600054600160a060020a03166340c10f1984836040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561025057600080fd5b5af1151561025d57600080fd5b505050604051805190505082600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36102be61038b565b505050565b60045481565b60025481565b60065481565b60055481565b600354600160a060020a031681565b60015481565b6006546005546000919010156103046103c1565b8061030c5750805b91505090565b600054600160a060020a031681565b60008060065461033c3460055461037190919063ffffffff16565b111590506103486103c9565b801561030c5750919050565b600061036b600454836103f990919063ffffffff16565b92915050565b60008282018381101561038057fe5b8091505b5092915050565b600354600160a060020a03163480156108fc0290604051600060405180830381858888f1935050505015156103bf57600080fd5b565b600254421190565b600080600060015442101580156103e257506002544211155b9150503415158180156103f25750805b9250505090565b60008083151561040c5760009150610384565b5082820282848281151561041c57fe5b041461038057fe5b600061042e610449565b604051809103906000f080151561044457600080fd5b905090565b604051610a798061045a833901905600606060405260038054600160a860020a03191633600160a060020a0316179055610a4b8061002e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f057806318160ddd1461011257806323b872dd1461013757806340c10f191461015f578063661884631461018157806370a08231146101a35780637d64bcb4146101c25780638da5cb5b146101d5578063a9059cbb14610204578063d73dd62314610226578063dd62ed3e14610248578063f2fde38b1461026d575b600080fd5b34156100d457600080fd5b6100dc61028e565b604051901515815260200160405180910390f35b34156100fb57600080fd5b6100dc600160a060020a03600435166024356102af565b341561011d57600080fd5b61012561031b565b60405190815260200160405180910390f35b341561014257600080fd5b6100dc600160a060020a0360043581169060243516604435610321565b341561016a57600080fd5b6100dc600160a060020a03600435166024356104a3565b341561018c57600080fd5b6100dc600160a060020a03600435166024356105c1565b34156101ae57600080fd5b610125600160a060020a03600435166106bb565b34156101cd57600080fd5b6100dc6106d6565b34156101e057600080fd5b6101e8610783565b604051600160a060020a03909116815260200160405180910390f35b341561020f57600080fd5b6100dc600160a060020a0360043516602435610792565b341561023157600080fd5b6100dc600160a060020a036004351660243561088d565b341561025357600080fd5b610125600160a060020a0360043581169060243516610931565b341561027857600080fd5b61028c600160a060020a036004351661095c565b005b60035474010000000000000000000000000000000000000000900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561033857600080fd5b600160a060020a03841660009081526001602052604090205482111561035d57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561039057600080fd5b600160a060020a0384166000908152600160205260409020546103b9908363ffffffff6109f716565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103ee908363ffffffff610a0916565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610436908363ffffffff6109f716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a039081169116146104c157600080fd5b60035474010000000000000000000000000000000000000000900460ff16156104e957600080fd5b6000546104fc908363ffffffff610a0916565b6000908155600160a060020a038416815260016020526040902054610527908363ffffffff610a0916565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561061e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610655565b61062e818463ffffffff6109f716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146106f457600080fd5b60035474010000000000000000000000000000000000000000900460ff161561071c57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156107a957600080fd5b600160a060020a0333166000908152600160205260409020548211156107ce57600080fd5b600160a060020a0333166000908152600160205260409020546107f7908363ffffffff6109f716565b600160a060020a03338116600090815260016020526040808220939093559085168152205461082c908363ffffffff610a0916565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546108c5908363ffffffff610a0916565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461097757600080fd5b600160a060020a038116151561098c57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610a0357fe5b50900390565b600082820183811015610a1857fe5b93925050505600a165627a7a723058208c2b1a98bb39cafb192d0bcfc353def73087dc0a2e807c8e8f102e0d31d4b88c0029a165627a7a723058209b8f76c65db28993678f9ebe6f52929ccd1bf1894fd20c57e3e34d3fef8c81e20029606060405260038054600160a860020a03191633600160a060020a0316179055610a4b8061002e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f057806318160ddd1461011257806323b872dd1461013757806340c10f191461015f578063661884631461018157806370a08231146101a35780637d64bcb4146101c25780638da5cb5b146101d5578063a9059cbb14610204578063d73dd62314610226578063dd62ed3e14610248578063f2fde38b1461026d575b600080fd5b34156100d457600080fd5b6100dc61028e565b604051901515815260200160405180910390f35b34156100fb57600080fd5b6100dc600160a060020a03600435166024356102af565b341561011d57600080fd5b61012561031b565b60405190815260200160405180910390f35b341561014257600080fd5b6100dc600160a060020a0360043581169060243516604435610321565b341561016a57600080fd5b6100dc600160a060020a03600435166024356104a3565b341561018c57600080fd5b6100dc600160a060020a03600435166024356105c1565b34156101ae57600080fd5b610125600160a060020a03600435166106bb565b34156101cd57600080fd5b6100dc6106d6565b34156101e057600080fd5b6101e8610783565b604051600160a060020a03909116815260200160405180910390f35b341561020f57600080fd5b6100dc600160a060020a0360043516602435610792565b341561023157600080fd5b6100dc600160a060020a036004351660243561088d565b341561025357600080fd5b610125600160a060020a0360043581169060243516610931565b341561027857600080fd5b61028c600160a060020a036004351661095c565b005b60035474010000000000000000000000000000000000000000900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561033857600080fd5b600160a060020a03841660009081526001602052604090205482111561035d57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561039057600080fd5b600160a060020a0384166000908152600160205260409020546103b9908363ffffffff6109f716565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103ee908363ffffffff610a0916565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610436908363ffffffff6109f716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a039081169116146104c157600080fd5b60035474010000000000000000000000000000000000000000900460ff16156104e957600080fd5b6000546104fc908363ffffffff610a0916565b6000908155600160a060020a038416815260016020526040902054610527908363ffffffff610a0916565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561061e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610655565b61062e818463ffffffff6109f716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146106f457600080fd5b60035474010000000000000000000000000000000000000000900460ff161561071c57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156107a957600080fd5b600160a060020a0333166000908152600160205260409020548211156107ce57600080fd5b600160a060020a0333166000908152600160205260409020546107f7908363ffffffff6109f716565b600160a060020a03338116600090815260016020526040808220939093559085168152205461082c908363ffffffff610a0916565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546108c5908363ffffffff610a0916565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461097757600080fd5b600160a060020a038116151561098c57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610a0357fe5b50900390565b600082820183811015610a1857fe5b93925050505600a165627a7a723058208c2b1a98bb39cafb192d0bcfc353def73087dc0a2e807c8e8f102e0d31d4b88c0029`

// DeployCappedCrowdsaleImpl deploys a new Ethereum contract, binding an instance of CappedCrowdsaleImpl to it.
func DeployCappedCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _startTime *big.Int, _endTime *big.Int, _rate *big.Int, _wallet common.Address, _cap *big.Int) (common.Address, *types.Transaction, *CappedCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CappedCrowdsaleImplBin), backend, _startTime, _endTime, _rate, _wallet, _cap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CappedCrowdsaleImpl{CappedCrowdsaleImplCaller: CappedCrowdsaleImplCaller{contract: contract}, CappedCrowdsaleImplTransactor: CappedCrowdsaleImplTransactor{contract: contract}, CappedCrowdsaleImplFilterer: CappedCrowdsaleImplFilterer{contract: contract}}, nil
}

// CappedCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type CappedCrowdsaleImpl struct {
	CappedCrowdsaleImplCaller     // Read-only binding to the contract
	CappedCrowdsaleImplTransactor // Write-only binding to the contract
	CappedCrowdsaleImplFilterer   // Log filterer for contract events
}

// CappedCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedCrowdsaleImplSession struct {
	Contract     *CappedCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CappedCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedCrowdsaleImplCallerSession struct {
	Contract *CappedCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CappedCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedCrowdsaleImplTransactorSession struct {
	Contract     *CappedCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CappedCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedCrowdsaleImplRaw struct {
	Contract *CappedCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// CappedCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplCallerRaw struct {
	Contract *CappedCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// CappedCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedCrowdsaleImplTransactorRaw struct {
	Contract *CappedCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedCrowdsaleImpl creates a new instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*CappedCrowdsaleImpl, error) {
	contract, err := bindCappedCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImpl{CappedCrowdsaleImplCaller: CappedCrowdsaleImplCaller{contract: contract}, CappedCrowdsaleImplTransactor: CappedCrowdsaleImplTransactor{contract: contract}, CappedCrowdsaleImplFilterer: CappedCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewCappedCrowdsaleImplCaller creates a new read-only instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*CappedCrowdsaleImplCaller, error) {
	contract, err := bindCappedCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImplCaller{contract: contract}, nil
}

// NewCappedCrowdsaleImplTransactor creates a new write-only instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedCrowdsaleImplTransactor, error) {
	contract, err := bindCappedCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImplTransactor{contract: contract}, nil
}

// NewCappedCrowdsaleImplFilterer creates a new log filterer instance of CappedCrowdsaleImpl, bound to a specific deployed contract.
func NewCappedCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedCrowdsaleImplFilterer, error) {
	contract, err := bindCappedCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImplFilterer{contract: contract}, nil
}

// bindCappedCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindCappedCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsaleImpl.Contract.CappedCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.CappedCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.CappedCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Cap() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Cap(&_CappedCrowdsaleImpl.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Cap() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Cap(&_CappedCrowdsaleImpl.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) EndTime() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.EndTime(&_CappedCrowdsaleImpl.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) EndTime() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.EndTime(&_CappedCrowdsaleImpl.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) HasEnded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "hasEnded")
	return *ret0, err
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) HasEnded() (bool, error) {
	return _CappedCrowdsaleImpl.Contract.HasEnded(&_CappedCrowdsaleImpl.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) HasEnded() (bool, error) {
	return _CappedCrowdsaleImpl.Contract.HasEnded(&_CappedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Rate(&_CappedCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.Rate(&_CappedCrowdsaleImpl.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) StartTime() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.StartTime(&_CappedCrowdsaleImpl.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) StartTime() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.StartTime(&_CappedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Token() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Token(&_CappedCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Token(&_CappedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Wallet(&_CappedCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _CappedCrowdsaleImpl.Contract.Wallet(&_CappedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CappedCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.WeiRaised(&_CappedCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsaleImpl.Contract.WeiRaised(&_CappedCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.BuyTokens(&_CappedCrowdsaleImpl.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CappedCrowdsaleImpl.Contract.BuyTokens(&_CappedCrowdsaleImpl.TransactOpts, beneficiary)
}

// CappedCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the CappedCrowdsaleImpl contract.
type CappedCrowdsaleImplTokenPurchaseIterator struct {
	Event *CappedCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *CappedCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCrowdsaleImplTokenPurchase)
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
		it.Event = new(CappedCrowdsaleImplTokenPurchase)
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
func (it *CappedCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the CappedCrowdsaleImpl contract.
type CappedCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*CappedCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleImplTokenPurchaseIterator{contract: _CappedCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_CappedCrowdsaleImpl *CappedCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *CappedCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCrowdsaleImplTokenPurchase)
				if err := _CappedCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
