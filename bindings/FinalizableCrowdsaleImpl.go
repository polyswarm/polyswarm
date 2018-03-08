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

// FinalizableCrowdsaleImplABI is the input ABI used to generate the binding from.
const FinalizableCrowdsaleImplABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startTime\",\"type\":\"uint256\"},{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// FinalizableCrowdsaleImplBin is the compiled bytecode used for deploying new contracts.
const FinalizableCrowdsaleImplBin = `60606040526006805460a060020a60ff0219169055341561001f57600080fd5b604051608080611bd28339810160405280805191906020018051919060200180519190602001805191508490508383834284101561005c57600080fd5b8383101561006957600080fd5b6000821161007657600080fd5b600160a060020a038116151561008b57600080fd5b6100a06401000000006105cb6100f882021704565b60008054600160a060020a0319908116600160a060020a039384161790915560019590955560029390935560049190915560038054841691831691909117905560068054909216339091161790555061012d92505050565b600061010261011d565b604051809103906000f080151561011857600080fd5b905090565b604051610a358061119d83390190565b6110618061013c6000396000f3006060604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e81146100c45780633197cbb6146100e95780634042b66f146100fc5780634bb278f31461010f578063521eb2731461012257806378e97925146101515780638d4e4083146101645780638da5cb5b1461018b578063ec8ac4d81461019e578063ecb70fb7146101b2578063f2fde38b146101c5578063fc0c546a146101e4575b6100c2336101f7565b005b34156100cf57600080fd5b6100d7610339565b60405190815260200160405180910390f35b34156100f457600080fd5b6100d761033f565b341561010757600080fd5b6100d7610345565b341561011a57600080fd5b6100c261034b565b341561012d57600080fd5b61013561040c565b604051600160a060020a03909116815260200160405180910390f35b341561015c57600080fd5b6100d761041b565b341561016f57600080fd5b610177610421565b604051901515815260200160405180910390f35b341561019657600080fd5b610135610442565b6100c2600160a060020a03600435166101f7565b34156101bd57600080fd5b610177610451565b34156101d057600080fd5b6100c2600160a060020a0360043516610459565b34156101ef57600080fd5b6101356104f4565b600080600160a060020a038316151561020f57600080fd5b610217610503565b151561022257600080fd5b34915061022e82610533565b600554909150610244908363ffffffff61055016565b60055560008054600160a060020a0316906340c10f199085908490604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156102c257600080fd5b6102c65a03f115156102d357600080fd5b505050604051805190505082600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a361033461056a565b505050565b60045481565b60025481565b60055481565b60065433600160a060020a0390811691161461036657600080fd5b60065474010000000000000000000000000000000000000000900460ff161561038e57600080fd5b610396610451565b15156103a157600080fd5b6103a961059e565b7f6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b1768160405160405180910390a16006805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600354600160a060020a031681565b60015481565b60065474010000000000000000000000000000000000000000900460ff1681565b600654600160a060020a031681565b600254421190565b60065433600160a060020a0390811691161461047457600080fd5b600160a060020a038116151561048957600080fd5b600654600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b6000806000600154421015801561051c57506002544211155b91505034151581801561052c5750805b9250505090565b600061054a600454836105a090919063ffffffff16565b92915050565b60008282018381101561055f57fe5b8091505b5092915050565b600354600160a060020a03163480156108fc0290604051600060405180830381858888f19350505050151561059e57600080fd5b565b6000808315156105b35760009150610563565b508282028284828115156105c357fe5b041461055f57fe5b60006105d56105f0565b604051809103906000f08015156105eb57600080fd5b905090565b604051610a3580610601833901905600606060405260038054600160a860020a03191633600160a060020a0316179055610a078061002e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f057806318160ddd1461011257806323b872dd1461013757806340c10f191461015f578063661884631461018157806370a08231146101a35780637d64bcb4146101c25780638da5cb5b146101d5578063a9059cbb14610204578063d73dd62314610226578063dd62ed3e14610248578063f2fde38b1461026d575b600080fd5b34156100d457600080fd5b6100dc61028e565b604051901515815260200160405180910390f35b34156100fb57600080fd5b6100dc600160a060020a036004351660243561029e565b341561011d57600080fd5b61012561030a565b60405190815260200160405180910390f35b341561014257600080fd5b6100dc600160a060020a0360043581169060243516604435610310565b341561016a57600080fd5b6100dc600160a060020a0360043516602435610492565b341561018c57600080fd5b6100dc600160a060020a036004351660243561059f565b34156101ae57600080fd5b610125600160a060020a0360043516610699565b34156101cd57600080fd5b6100dc6106b4565b34156101e057600080fd5b6101e861073f565b604051600160a060020a03909116815260200160405180910390f35b341561020f57600080fd5b6100dc600160a060020a036004351660243561074e565b341561023157600080fd5b6100dc600160a060020a0360043516602435610849565b341561025357600080fd5b610125600160a060020a03600435811690602435166108ed565b341561027857600080fd5b61028c600160a060020a0360043516610918565b005b60035460a060020a900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561032757600080fd5b600160a060020a03841660009081526001602052604090205482111561034c57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561037f57600080fd5b600160a060020a0384166000908152600160205260409020546103a8908363ffffffff6109b316565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103dd908363ffffffff6109c516565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610425908363ffffffff6109b316565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a039081169116146104b057600080fd5b60035460a060020a900460ff16156104c757600080fd5b6000546104da908363ffffffff6109c516565b6000908155600160a060020a038416815260016020526040902054610505908363ffffffff6109c516565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156105fc57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610633565b61060c818463ffffffff6109b316565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146106d257600080fd5b60035460a060020a900460ff16156106e957600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a038316151561076557600080fd5b600160a060020a03331660009081526001602052604090205482111561078a57600080fd5b600160a060020a0333166000908152600160205260409020546107b3908363ffffffff6109b316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107e8908363ffffffff6109c516565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610881908363ffffffff6109c516565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461093357600080fd5b600160a060020a038116151561094857600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828211156109bf57fe5b50900390565b6000828201838110156109d457fe5b93925050505600a165627a7a72305820fa18f4528833494b2fc254750f22bca7ebe05d911e637291d3a0ee6c6a1d86420029a165627a7a72305820ce3879c799d20b7b394b9da4b1a427fadcad652858d465e319c776652a90cd360029606060405260038054600160a860020a03191633600160a060020a0316179055610a078061002e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f057806318160ddd1461011257806323b872dd1461013757806340c10f191461015f578063661884631461018157806370a08231146101a35780637d64bcb4146101c25780638da5cb5b146101d5578063a9059cbb14610204578063d73dd62314610226578063dd62ed3e14610248578063f2fde38b1461026d575b600080fd5b34156100d457600080fd5b6100dc61028e565b604051901515815260200160405180910390f35b34156100fb57600080fd5b6100dc600160a060020a036004351660243561029e565b341561011d57600080fd5b61012561030a565b60405190815260200160405180910390f35b341561014257600080fd5b6100dc600160a060020a0360043581169060243516604435610310565b341561016a57600080fd5b6100dc600160a060020a0360043516602435610492565b341561018c57600080fd5b6100dc600160a060020a036004351660243561059f565b34156101ae57600080fd5b610125600160a060020a0360043516610699565b34156101cd57600080fd5b6100dc6106b4565b34156101e057600080fd5b6101e861073f565b604051600160a060020a03909116815260200160405180910390f35b341561020f57600080fd5b6100dc600160a060020a036004351660243561074e565b341561023157600080fd5b6100dc600160a060020a0360043516602435610849565b341561025357600080fd5b610125600160a060020a03600435811690602435166108ed565b341561027857600080fd5b61028c600160a060020a0360043516610918565b005b60035460a060020a900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561032757600080fd5b600160a060020a03841660009081526001602052604090205482111561034c57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561037f57600080fd5b600160a060020a0384166000908152600160205260409020546103a8908363ffffffff6109b316565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103dd908363ffffffff6109c516565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610425908363ffffffff6109b316565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a039081169116146104b057600080fd5b60035460a060020a900460ff16156104c757600080fd5b6000546104da908363ffffffff6109c516565b6000908155600160a060020a038416815260016020526040902054610505908363ffffffff6109c516565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156105fc57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610633565b61060c818463ffffffff6109b316565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146106d257600080fd5b60035460a060020a900460ff16156106e957600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a038316151561076557600080fd5b600160a060020a03331660009081526001602052604090205482111561078a57600080fd5b600160a060020a0333166000908152600160205260409020546107b3908363ffffffff6109b316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107e8908363ffffffff6109c516565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610881908363ffffffff6109c516565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461093357600080fd5b600160a060020a038116151561094857600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828211156109bf57fe5b50900390565b6000828201838110156109d457fe5b93925050505600a165627a7a72305820fa18f4528833494b2fc254750f22bca7ebe05d911e637291d3a0ee6c6a1d86420029`

// DeployFinalizableCrowdsaleImpl deploys a new Ethereum contract, binding an instance of FinalizableCrowdsaleImpl to it.
func DeployFinalizableCrowdsaleImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _startTime *big.Int, _endTime *big.Int, _rate *big.Int, _wallet common.Address) (common.Address, *types.Transaction, *FinalizableCrowdsaleImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FinalizableCrowdsaleImplBin), backend, _startTime, _endTime, _rate, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalizableCrowdsaleImpl{FinalizableCrowdsaleImplCaller: FinalizableCrowdsaleImplCaller{contract: contract}, FinalizableCrowdsaleImplTransactor: FinalizableCrowdsaleImplTransactor{contract: contract}, FinalizableCrowdsaleImplFilterer: FinalizableCrowdsaleImplFilterer{contract: contract}}, nil
}

// FinalizableCrowdsaleImpl is an auto generated Go binding around an Ethereum contract.
type FinalizableCrowdsaleImpl struct {
	FinalizableCrowdsaleImplCaller     // Read-only binding to the contract
	FinalizableCrowdsaleImplTransactor // Write-only binding to the contract
	FinalizableCrowdsaleImplFilterer   // Log filterer for contract events
}

// FinalizableCrowdsaleImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizableCrowdsaleImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizableCrowdsaleImplSession struct {
	Contract     *FinalizableCrowdsaleImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizableCrowdsaleImplCallerSession struct {
	Contract *FinalizableCrowdsaleImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// FinalizableCrowdsaleImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizableCrowdsaleImplTransactorSession struct {
	Contract     *FinalizableCrowdsaleImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplRaw struct {
	Contract *FinalizableCrowdsaleImpl // Generic contract binding to access the raw methods on
}

// FinalizableCrowdsaleImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplCallerRaw struct {
	Contract *FinalizableCrowdsaleImplCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizableCrowdsaleImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleImplTransactorRaw struct {
	Contract *FinalizableCrowdsaleImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizableCrowdsaleImpl creates a new instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImpl(address common.Address, backend bind.ContractBackend) (*FinalizableCrowdsaleImpl, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImpl{FinalizableCrowdsaleImplCaller: FinalizableCrowdsaleImplCaller{contract: contract}, FinalizableCrowdsaleImplTransactor: FinalizableCrowdsaleImplTransactor{contract: contract}, FinalizableCrowdsaleImplFilterer: FinalizableCrowdsaleImplFilterer{contract: contract}}, nil
}

// NewFinalizableCrowdsaleImplCaller creates a new read-only instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImplCaller(address common.Address, caller bind.ContractCaller) (*FinalizableCrowdsaleImplCaller, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplCaller{contract: contract}, nil
}

// NewFinalizableCrowdsaleImplTransactor creates a new write-only instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImplTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizableCrowdsaleImplTransactor, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplTransactor{contract: contract}, nil
}

// NewFinalizableCrowdsaleImplFilterer creates a new log filterer instance of FinalizableCrowdsaleImpl, bound to a specific deployed contract.
func NewFinalizableCrowdsaleImplFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizableCrowdsaleImplFilterer, error) {
	contract, err := bindFinalizableCrowdsaleImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplFilterer{contract: contract}, nil
}

// bindFinalizableCrowdsaleImpl binds a generic wrapper to an already deployed contract.
func bindFinalizableCrowdsaleImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsaleImpl.Contract.FinalizableCrowdsaleImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.FinalizableCrowdsaleImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.FinalizableCrowdsaleImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsaleImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.contract.Transact(opts, method, params...)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) EndTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.EndTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) EndTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.EndTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) HasEnded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "hasEnded")
	return *ret0, err
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) HasEnded() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.HasEnded(&_FinalizableCrowdsaleImpl.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) HasEnded() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.HasEnded(&_FinalizableCrowdsaleImpl.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) IsFinalized() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.IsFinalized(&_FinalizableCrowdsaleImpl.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) IsFinalized() (bool, error) {
	return _FinalizableCrowdsaleImpl.Contract.IsFinalized(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Owner() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Owner(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Owner() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Owner(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.Rate(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.Rate(&_FinalizableCrowdsaleImpl.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) StartTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.StartTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) StartTime() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.StartTime(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Token() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Token(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Token() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Token(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Wallet(&_FinalizableCrowdsaleImpl.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsaleImpl.Contract.Wallet(&_FinalizableCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FinalizableCrowdsaleImpl.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.WeiRaised(&_FinalizableCrowdsaleImpl.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplCallerSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsaleImpl.Contract.WeiRaised(&_FinalizableCrowdsaleImpl.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.BuyTokens(&_FinalizableCrowdsaleImpl.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.BuyTokens(&_FinalizableCrowdsaleImpl.TransactOpts, beneficiary)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.Finalize(&_FinalizableCrowdsaleImpl.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.Finalize(&_FinalizableCrowdsaleImpl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.TransferOwnership(&_FinalizableCrowdsaleImpl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsaleImpl.Contract.TransferOwnership(&_FinalizableCrowdsaleImpl.TransactOpts, newOwner)
}

// FinalizableCrowdsaleImplFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplFinalizedIterator struct {
	Event *FinalizableCrowdsaleImplFinalized // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleImplFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleImplFinalized)
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
		it.Event = new(FinalizableCrowdsaleImplFinalized)
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
func (it *FinalizableCrowdsaleImplFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleImplFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleImplFinalized represents a Finalized event raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) FilterFinalized(opts *bind.FilterOpts) (*FinalizableCrowdsaleImplFinalizedIterator, error) {

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplFinalizedIterator{contract: _FinalizableCrowdsaleImpl.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleImplFinalized) (event.Subscription, error) {

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleImplFinalized)
				if err := _FinalizableCrowdsaleImpl.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// FinalizableCrowdsaleImplOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplOwnershipTransferredIterator struct {
	Event *FinalizableCrowdsaleImplOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleImplOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleImplOwnershipTransferred)
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
		it.Event = new(FinalizableCrowdsaleImplOwnershipTransferred)
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
func (it *FinalizableCrowdsaleImplOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleImplOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleImplOwnershipTransferred represents a OwnershipTransferred event raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinalizableCrowdsaleImplOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplOwnershipTransferredIterator{contract: _FinalizableCrowdsaleImpl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleImplOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleImplOwnershipTransferred)
				if err := _FinalizableCrowdsaleImpl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// FinalizableCrowdsaleImplTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplTokenPurchaseIterator struct {
	Event *FinalizableCrowdsaleImplTokenPurchase // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleImplTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleImplTokenPurchase)
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
		it.Event = new(FinalizableCrowdsaleImplTokenPurchase)
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
func (it *FinalizableCrowdsaleImplTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleImplTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleImplTokenPurchase represents a TokenPurchase event raised by the FinalizableCrowdsaleImpl contract.
type FinalizableCrowdsaleImplTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*FinalizableCrowdsaleImplTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleImplTokenPurchaseIterator{contract: _FinalizableCrowdsaleImpl.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_FinalizableCrowdsaleImpl *FinalizableCrowdsaleImplFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleImplTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsaleImpl.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleImplTokenPurchase)
				if err := _FinalizableCrowdsaleImpl.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
