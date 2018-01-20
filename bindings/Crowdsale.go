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

// CrowdsaleABI is the input ABI used to generate the binding from.
const CrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startTime\",\"type\":\"uint256\"},{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// CrowdsaleBin is the compiled bytecode used for deploying new contracts.
const CrowdsaleBin = `6060604052341561000f57600080fd5b604051608080611997833981016040528080519190602001805191906020018051919060200180519150504284101561004757600080fd5b8383101561005457600080fd5b6000821161006157600080fd5b600160a060020a038116151561007657600080fd5b61008b6401000000006103bb6100cd82021704565b60008054600160a060020a0319908116600160a060020a0393841617909155600195909555600293909355600491909155600380549093169116179055610102565b60006100d76100f2565b604051809103906000f08015156100ed57600080fd5b905090565b604051610a3580610f6283390190565b610e51806101116000396000f30060606040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632c4e722e81146100985780633197cbb6146100bd5780634042b66f146100d0578063521eb273146100e357806378e9792514610112578063ec8ac4d814610125578063ecb70fb714610139578063fc0c546a14610160575b61009633610173565b005b34156100a357600080fd5b6100ab6102b5565b60405190815260200160405180910390f35b34156100c857600080fd5b6100ab6102bb565b34156100db57600080fd5b6100ab6102c1565b34156100ee57600080fd5b6100f66102c7565b604051600160a060020a03909116815260200160405180910390f35b341561011d57600080fd5b6100ab6102d6565b610096600160a060020a0360043516610173565b341561014457600080fd5b61014c6102dc565b604051901515815260200160405180910390f35b341561016b57600080fd5b6100f66102e4565b600080600160a060020a038316151561018b57600080fd5b6101936102f3565b151561019e57600080fd5b3491506101aa82610323565b6005549091506101c0908363ffffffff61034016565b60055560008054600160a060020a0316906340c10f199085908490604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561023e57600080fd5b6102c65a03f1151561024f57600080fd5b505050604051805190505082600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a36102b061035a565b505050565b60045481565b60025481565b60055481565b600354600160a060020a031681565b60015481565b600254421190565b600054600160a060020a031681565b6000806000600154421015801561030c57506002544211155b91505034151581801561031c5750805b9250505090565b600061033a6004548361039090919063ffffffff16565b92915050565b60008282018381101561034f57fe5b8091505b5092915050565b600354600160a060020a03163480156108fc0290604051600060405180830381858888f19350505050151561038e57600080fd5b565b6000808315156103a35760009150610353565b508282028284828115156103b357fe5b041461034f57fe5b60006103c56103e0565b604051809103906000f08015156103db57600080fd5b905090565b604051610a35806103f1833901905600606060405260038054600160a860020a03191633600160a060020a0316179055610a078061002e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f057806318160ddd1461011257806323b872dd1461013757806340c10f191461015f578063661884631461018157806370a08231146101a35780637d64bcb4146101c25780638da5cb5b146101d5578063a9059cbb14610204578063d73dd62314610226578063dd62ed3e14610248578063f2fde38b1461026d575b600080fd5b34156100d457600080fd5b6100dc61028e565b604051901515815260200160405180910390f35b34156100fb57600080fd5b6100dc600160a060020a036004351660243561029e565b341561011d57600080fd5b61012561030a565b60405190815260200160405180910390f35b341561014257600080fd5b6100dc600160a060020a0360043581169060243516604435610310565b341561016a57600080fd5b6100dc600160a060020a0360043516602435610492565b341561018c57600080fd5b6100dc600160a060020a036004351660243561059f565b34156101ae57600080fd5b610125600160a060020a0360043516610699565b34156101cd57600080fd5b6100dc6106b4565b34156101e057600080fd5b6101e861073f565b604051600160a060020a03909116815260200160405180910390f35b341561020f57600080fd5b6100dc600160a060020a036004351660243561074e565b341561023157600080fd5b6100dc600160a060020a0360043516602435610849565b341561025357600080fd5b610125600160a060020a03600435811690602435166108ed565b341561027857600080fd5b61028c600160a060020a0360043516610918565b005b60035460a060020a900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561032757600080fd5b600160a060020a03841660009081526001602052604090205482111561034c57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561037f57600080fd5b600160a060020a0384166000908152600160205260409020546103a8908363ffffffff6109b316565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103dd908363ffffffff6109c516565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610425908363ffffffff6109b316565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a039081169116146104b057600080fd5b60035460a060020a900460ff16156104c757600080fd5b6000546104da908363ffffffff6109c516565b6000908155600160a060020a038416815260016020526040902054610505908363ffffffff6109c516565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156105fc57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610633565b61060c818463ffffffff6109b316565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146106d257600080fd5b60035460a060020a900460ff16156106e957600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a038316151561076557600080fd5b600160a060020a03331660009081526001602052604090205482111561078a57600080fd5b600160a060020a0333166000908152600160205260409020546107b3908363ffffffff6109b316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107e8908363ffffffff6109c516565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610881908363ffffffff6109c516565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461093357600080fd5b600160a060020a038116151561094857600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828211156109bf57fe5b50900390565b6000828201838110156109d457fe5b93925050505600a165627a7a72305820fa18f4528833494b2fc254750f22bca7ebe05d911e637291d3a0ee6c6a1d86420029a165627a7a723058205766087d829a56fa0942ee507bb94f5574a43467be64d0413e3ffa6285f613030029606060405260038054600160a860020a03191633600160a060020a0316179055610a078061002e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f057806318160ddd1461011257806323b872dd1461013757806340c10f191461015f578063661884631461018157806370a08231146101a35780637d64bcb4146101c25780638da5cb5b146101d5578063a9059cbb14610204578063d73dd62314610226578063dd62ed3e14610248578063f2fde38b1461026d575b600080fd5b34156100d457600080fd5b6100dc61028e565b604051901515815260200160405180910390f35b34156100fb57600080fd5b6100dc600160a060020a036004351660243561029e565b341561011d57600080fd5b61012561030a565b60405190815260200160405180910390f35b341561014257600080fd5b6100dc600160a060020a0360043581169060243516604435610310565b341561016a57600080fd5b6100dc600160a060020a0360043516602435610492565b341561018c57600080fd5b6100dc600160a060020a036004351660243561059f565b34156101ae57600080fd5b610125600160a060020a0360043516610699565b34156101cd57600080fd5b6100dc6106b4565b34156101e057600080fd5b6101e861073f565b604051600160a060020a03909116815260200160405180910390f35b341561020f57600080fd5b6100dc600160a060020a036004351660243561074e565b341561023157600080fd5b6100dc600160a060020a0360043516602435610849565b341561025357600080fd5b610125600160a060020a03600435811690602435166108ed565b341561027857600080fd5b61028c600160a060020a0360043516610918565b005b60035460a060020a900460ff1681565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561032757600080fd5b600160a060020a03841660009081526001602052604090205482111561034c57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561037f57600080fd5b600160a060020a0384166000908152600160205260409020546103a8908363ffffffff6109b316565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103dd908363ffffffff6109c516565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610425908363ffffffff6109b316565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60035460009033600160a060020a039081169116146104b057600080fd5b60035460a060020a900460ff16156104c757600080fd5b6000546104da908363ffffffff6109c516565b6000908155600160a060020a038416815260016020526040902054610505908363ffffffff6109c516565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156105fc57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610633565b61060c818463ffffffff6109b316565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146106d257600080fd5b60035460a060020a900460ff16156106e957600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b6000600160a060020a038316151561076557600080fd5b600160a060020a03331660009081526001602052604090205482111561078a57600080fd5b600160a060020a0333166000908152600160205260409020546107b3908363ffffffff6109b316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107e8908363ffffffff6109c516565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610881908363ffffffff6109c516565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461093357600080fd5b600160a060020a038116151561094857600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828211156109bf57fe5b50900390565b6000828201838110156109d457fe5b93925050505600a165627a7a72305820fa18f4528833494b2fc254750f22bca7ebe05d911e637291d3a0ee6c6a1d86420029`

// DeployCrowdsale deploys a new Ethereum contract, binding an instance of Crowdsale to it.
func DeployCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _startTime *big.Int, _endTime *big.Int, _rate *big.Int, _wallet common.Address) (common.Address, *types.Transaction, *Crowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrowdsaleBin), backend, _startTime, _endTime, _rate, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Crowdsale{CrowdsaleCaller: CrowdsaleCaller{contract: contract}, CrowdsaleTransactor: CrowdsaleTransactor{contract: contract}}, nil
}

// Crowdsale is an auto generated Go binding around an Ethereum contract.
type Crowdsale struct {
	CrowdsaleCaller     // Read-only binding to the contract
	CrowdsaleTransactor // Write-only binding to the contract
}

// CrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdsaleSession struct {
	Contract     *Crowdsale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdsaleCallerSession struct {
	Contract *CrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdsaleTransactorSession struct {
	Contract     *CrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdsaleRaw struct {
	Contract *Crowdsale // Generic contract binding to access the raw methods on
}

// CrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdsaleCallerRaw struct {
	Contract *CrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdsaleTransactorRaw struct {
	Contract *CrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdsale creates a new instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsale(address common.Address, backend bind.ContractBackend) (*Crowdsale, error) {
	contract, err := bindCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crowdsale{CrowdsaleCaller: CrowdsaleCaller{contract: contract}, CrowdsaleTransactor: CrowdsaleTransactor{contract: contract}}, nil
}

// NewCrowdsaleCaller creates a new read-only instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CrowdsaleCaller, error) {
	contract, err := bindCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleCaller{contract: contract}, nil
}

// NewCrowdsaleTransactor creates a new write-only instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdsaleTransactor, error) {
	contract, err := bindCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleTransactor{contract: contract}, nil
}

// bindCrowdsale binds a generic wrapper to an already deployed contract.
func bindCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsale *CrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Crowdsale.Contract.CrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsale *CrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.Contract.CrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsale *CrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsale.Contract.CrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsale *CrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Crowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsale *CrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsale *CrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsale.Contract.contract.Transact(opts, method, params...)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) EndTime() (*big.Int, error) {
	return _Crowdsale.Contract.EndTime(&_Crowdsale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) EndTime() (*big.Int, error) {
	return _Crowdsale.Contract.EndTime(&_Crowdsale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_Crowdsale *CrowdsaleCaller) HasEnded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "hasEnded")
	return *ret0, err
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_Crowdsale *CrowdsaleSession) HasEnded() (bool, error) {
	return _Crowdsale.Contract.HasEnded(&_Crowdsale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) HasEnded() (bool, error) {
	return _Crowdsale.Contract.HasEnded(&_Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) Rate() (*big.Int, error) {
	return _Crowdsale.Contract.Rate(&_Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _Crowdsale.Contract.Rate(&_Crowdsale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) StartTime() (*big.Int, error) {
	return _Crowdsale.Contract.StartTime(&_Crowdsale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) StartTime() (*big.Int, error) {
	return _Crowdsale.Contract.StartTime(&_Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Crowdsale *CrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Crowdsale *CrowdsaleSession) Token() (common.Address, error) {
	return _Crowdsale.Contract.Token(&_Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Token() (common.Address, error) {
	return _Crowdsale.Contract.Token(&_Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Crowdsale *CrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Crowdsale *CrowdsaleSession) Wallet() (common.Address, error) {
	return _Crowdsale.Contract.Wallet(&_Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _Crowdsale.Contract.Wallet(&_Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.WeiRaised(&_Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.WeiRaised(&_Crowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_Crowdsale *CrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_Crowdsale *CrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.BuyTokens(&_Crowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_Crowdsale *CrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.BuyTokens(&_Crowdsale.TransactOpts, beneficiary)
}
