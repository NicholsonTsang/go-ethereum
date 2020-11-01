// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NodeRegistrationImplABI is the input ABI used to generate the binding from.
const NodeRegistrationImplABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"applicantAddress\",\"type\":\"address\"}],\"name\":\"addingPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressNetworkState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressOtphashedValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"deleteMappingElement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashedAdminPW\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedpw\",\"type\":\"bytes32\"}],\"name\":\"setHashedAdminPW\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"applicantAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"otpHashValue\",\"type\":\"bytes32\"}],\"name\":\"settingOtpHashValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"otp\",\"type\":\"string\"}],\"name\":\"verifyOtp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NodeRegistrationImplFuncSigs maps the 4-byte function signature to its string representation.
var NodeRegistrationImplFuncSigs = map[string]string{
	"6d051a1f": "addingPeer(address)",
	"29dc926d": "addressNetworkState(address)",
	"5e04673d": "addressOtphashedValue(address)",
	"f851a440": "admin()",
	"51a1f32b": "deleteMappingElement(address)",
	"5e7bf051": "hashedAdminPW()",
	"131bea52": "setHashedAdminPW(bytes32)",
	"72f9013c": "settingOtpHashValue(address,bytes32)",
	"9d544042": "verifyOtp(string)",
}

// NodeRegistrationImplBin is the compiled bytecode used for deploying new contracts.
var NodeRegistrationImplBin = "0x60806040527f2e6285b5dde400bb8e8306ceffb0f9ade888e91d2b4d3aa8b23981e9c90c5eaa60025534801561003457600080fd5b50600080546001600160a01b0319163317905561047f806100566000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635e7bf051116100665780635e7bf0511461013b5780636d051a1f1461014357806372f9013c146101695780639d54404214610195578063f851a4401461023b57610093565b8063131bea521461009857806329dc926d146100b757806351a1f32b146100ef5780635e04673d14610115575b600080fd5b6100b5600480360360208110156100ae57600080fd5b503561025f565b005b6100dd600480360360208110156100cd57600080fd5b50356001600160a01b031661027b565b60408051918252519081900360200190f35b6100b56004803603602081101561010557600080fd5b50356001600160a01b0316610299565b6100dd6004803603602081101561012b57600080fd5b50356001600160a01b03166102d8565b6100dd6102f3565b6100b56004803603602081101561015957600080fd5b50356001600160a01b03166102f9565b6100b56004803603604081101561017f57600080fd5b506001600160a01b038135169060200135610357565b6100b5600480360360208110156101ab57600080fd5b8101906020810181356401000000008111156101c657600080fd5b8201836020820111156101d857600080fd5b803590602001918460018302840111640100000000831117156101fa57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103c4945050505050565b610243610416565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b0316331461027657600080fd5b600255565b6001600160a01b031660009081526001602052604090206002015490565b6000546001600160a01b031633146102b057600080fd5b6001600160a01b03166000908152600160208190526040822082815590810182905560020155565b6001600160a01b031660009081526001602052604090205490565b60025481565b6000546001600160a01b0316331461031057600080fd5b6001600160a01b0381166000908152600160208190526040909120600201541415610354576001600160a01b03811660009081526001602052604090206002908101555b50565b6000546001600160a01b0316331461036e57600080fd5b610376610425565b50604080516060810182529182526203f4804201602080840191825260008484018181526001600160a01b039096168152600191829052929092209251835551908201559051600290910155565b3360009081526001602052604090206002015461035457336000908152600160209081526040909120548251918301919091201415610354573360009081526001602081905260409091206002015550565b6000546001600160a01b031681565b6040518060600160405280600080191681526020016000815260200160008152509056fea2646970667358221220e19e1855b6c78f216298739bd6844fe69cbb756052f3198d714d11125ab74c0164736f6c63430007040033"

// DeployNodeRegistrationImpl deploys a new Ethereum contract, binding an instance of NodeRegistrationImpl to it.
func DeployNodeRegistrationImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeRegistrationImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRegistrationImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeRegistrationImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeRegistrationImpl{NodeRegistrationImplCaller: NodeRegistrationImplCaller{contract: contract}, NodeRegistrationImplTransactor: NodeRegistrationImplTransactor{contract: contract}, NodeRegistrationImplFilterer: NodeRegistrationImplFilterer{contract: contract}}, nil
}

// NodeRegistrationImpl is an auto generated Go binding around an Ethereum contract.
type NodeRegistrationImpl struct {
	NodeRegistrationImplCaller     // Read-only binding to the contract
	NodeRegistrationImplTransactor // Write-only binding to the contract
	NodeRegistrationImplFilterer   // Log filterer for contract events
}

// NodeRegistrationImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRegistrationImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrationImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRegistrationImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrationImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRegistrationImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrationImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRegistrationImplSession struct {
	Contract     *NodeRegistrationImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NodeRegistrationImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRegistrationImplCallerSession struct {
	Contract *NodeRegistrationImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// NodeRegistrationImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRegistrationImplTransactorSession struct {
	Contract     *NodeRegistrationImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// NodeRegistrationImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRegistrationImplRaw struct {
	Contract *NodeRegistrationImpl // Generic contract binding to access the raw methods on
}

// NodeRegistrationImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRegistrationImplCallerRaw struct {
	Contract *NodeRegistrationImplCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRegistrationImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRegistrationImplTransactorRaw struct {
	Contract *NodeRegistrationImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRegistrationImpl creates a new instance of NodeRegistrationImpl, bound to a specific deployed contract.
func NewNodeRegistrationImpl(address common.Address, backend bind.ContractBackend) (*NodeRegistrationImpl, error) {
	contract, err := bindNodeRegistrationImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationImpl{NodeRegistrationImplCaller: NodeRegistrationImplCaller{contract: contract}, NodeRegistrationImplTransactor: NodeRegistrationImplTransactor{contract: contract}, NodeRegistrationImplFilterer: NodeRegistrationImplFilterer{contract: contract}}, nil
}

// NewNodeRegistrationImplCaller creates a new read-only instance of NodeRegistrationImpl, bound to a specific deployed contract.
func NewNodeRegistrationImplCaller(address common.Address, caller bind.ContractCaller) (*NodeRegistrationImplCaller, error) {
	contract, err := bindNodeRegistrationImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationImplCaller{contract: contract}, nil
}

// NewNodeRegistrationImplTransactor creates a new write-only instance of NodeRegistrationImpl, bound to a specific deployed contract.
func NewNodeRegistrationImplTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRegistrationImplTransactor, error) {
	contract, err := bindNodeRegistrationImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationImplTransactor{contract: contract}, nil
}

// NewNodeRegistrationImplFilterer creates a new log filterer instance of NodeRegistrationImpl, bound to a specific deployed contract.
func NewNodeRegistrationImplFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRegistrationImplFilterer, error) {
	contract, err := bindNodeRegistrationImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationImplFilterer{contract: contract}, nil
}

// bindNodeRegistrationImpl binds a generic wrapper to an already deployed contract.
func bindNodeRegistrationImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRegistrationImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRegistrationImpl *NodeRegistrationImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRegistrationImpl.Contract.NodeRegistrationImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRegistrationImpl *NodeRegistrationImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.NodeRegistrationImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRegistrationImpl *NodeRegistrationImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.NodeRegistrationImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRegistrationImpl *NodeRegistrationImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRegistrationImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.contract.Transact(opts, method, params...)
}

// AddressNetworkState is a free data retrieval call binding the contract method 0x29dc926d.
//
// Solidity: function addressNetworkState(address a) view returns(uint256)
func (_NodeRegistrationImpl *NodeRegistrationImplCaller) AddressNetworkState(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeRegistrationImpl.contract.Call(opts, &out, "addressNetworkState", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressNetworkState is a free data retrieval call binding the contract method 0x29dc926d.
//
// Solidity: function addressNetworkState(address a) view returns(uint256)
func (_NodeRegistrationImpl *NodeRegistrationImplSession) AddressNetworkState(a common.Address) (*big.Int, error) {
	return _NodeRegistrationImpl.Contract.AddressNetworkState(&_NodeRegistrationImpl.CallOpts, a)
}

// AddressNetworkState is a free data retrieval call binding the contract method 0x29dc926d.
//
// Solidity: function addressNetworkState(address a) view returns(uint256)
func (_NodeRegistrationImpl *NodeRegistrationImplCallerSession) AddressNetworkState(a common.Address) (*big.Int, error) {
	return _NodeRegistrationImpl.Contract.AddressNetworkState(&_NodeRegistrationImpl.CallOpts, a)
}

// AddressOtphashedValue is a free data retrieval call binding the contract method 0x5e04673d.
//
// Solidity: function addressOtphashedValue(address a) view returns(bytes32)
func (_NodeRegistrationImpl *NodeRegistrationImplCaller) AddressOtphashedValue(opts *bind.CallOpts, a common.Address) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistrationImpl.contract.Call(opts, &out, "addressOtphashedValue", a)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressOtphashedValue is a free data retrieval call binding the contract method 0x5e04673d.
//
// Solidity: function addressOtphashedValue(address a) view returns(bytes32)
func (_NodeRegistrationImpl *NodeRegistrationImplSession) AddressOtphashedValue(a common.Address) ([32]byte, error) {
	return _NodeRegistrationImpl.Contract.AddressOtphashedValue(&_NodeRegistrationImpl.CallOpts, a)
}

// AddressOtphashedValue is a free data retrieval call binding the contract method 0x5e04673d.
//
// Solidity: function addressOtphashedValue(address a) view returns(bytes32)
func (_NodeRegistrationImpl *NodeRegistrationImplCallerSession) AddressOtphashedValue(a common.Address) ([32]byte, error) {
	return _NodeRegistrationImpl.Contract.AddressOtphashedValue(&_NodeRegistrationImpl.CallOpts, a)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NodeRegistrationImpl *NodeRegistrationImplCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRegistrationImpl.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NodeRegistrationImpl *NodeRegistrationImplSession) Admin() (common.Address, error) {
	return _NodeRegistrationImpl.Contract.Admin(&_NodeRegistrationImpl.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NodeRegistrationImpl *NodeRegistrationImplCallerSession) Admin() (common.Address, error) {
	return _NodeRegistrationImpl.Contract.Admin(&_NodeRegistrationImpl.CallOpts)
}

// HashedAdminPW is a free data retrieval call binding the contract method 0x5e7bf051.
//
// Solidity: function hashedAdminPW() view returns(bytes32)
func (_NodeRegistrationImpl *NodeRegistrationImplCaller) HashedAdminPW(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistrationImpl.contract.Call(opts, &out, "hashedAdminPW")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashedAdminPW is a free data retrieval call binding the contract method 0x5e7bf051.
//
// Solidity: function hashedAdminPW() view returns(bytes32)
func (_NodeRegistrationImpl *NodeRegistrationImplSession) HashedAdminPW() ([32]byte, error) {
	return _NodeRegistrationImpl.Contract.HashedAdminPW(&_NodeRegistrationImpl.CallOpts)
}

// HashedAdminPW is a free data retrieval call binding the contract method 0x5e7bf051.
//
// Solidity: function hashedAdminPW() view returns(bytes32)
func (_NodeRegistrationImpl *NodeRegistrationImplCallerSession) HashedAdminPW() ([32]byte, error) {
	return _NodeRegistrationImpl.Contract.HashedAdminPW(&_NodeRegistrationImpl.CallOpts)
}

// AddingPeer is a paid mutator transaction binding the contract method 0x6d051a1f.
//
// Solidity: function addingPeer(address applicantAddress) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactor) AddingPeer(opts *bind.TransactOpts, applicantAddress common.Address) (*types.Transaction, error) {
	return _NodeRegistrationImpl.contract.Transact(opts, "addingPeer", applicantAddress)
}

// AddingPeer is a paid mutator transaction binding the contract method 0x6d051a1f.
//
// Solidity: function addingPeer(address applicantAddress) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplSession) AddingPeer(applicantAddress common.Address) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.AddingPeer(&_NodeRegistrationImpl.TransactOpts, applicantAddress)
}

// AddingPeer is a paid mutator transaction binding the contract method 0x6d051a1f.
//
// Solidity: function addingPeer(address applicantAddress) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorSession) AddingPeer(applicantAddress common.Address) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.AddingPeer(&_NodeRegistrationImpl.TransactOpts, applicantAddress)
}

// DeleteMappingElement is a paid mutator transaction binding the contract method 0x51a1f32b.
//
// Solidity: function deleteMappingElement(address a) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactor) DeleteMappingElement(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _NodeRegistrationImpl.contract.Transact(opts, "deleteMappingElement", a)
}

// DeleteMappingElement is a paid mutator transaction binding the contract method 0x51a1f32b.
//
// Solidity: function deleteMappingElement(address a) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplSession) DeleteMappingElement(a common.Address) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.DeleteMappingElement(&_NodeRegistrationImpl.TransactOpts, a)
}

// DeleteMappingElement is a paid mutator transaction binding the contract method 0x51a1f32b.
//
// Solidity: function deleteMappingElement(address a) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorSession) DeleteMappingElement(a common.Address) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.DeleteMappingElement(&_NodeRegistrationImpl.TransactOpts, a)
}

// SetHashedAdminPW is a paid mutator transaction binding the contract method 0x131bea52.
//
// Solidity: function setHashedAdminPW(bytes32 hashedpw) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactor) SetHashedAdminPW(opts *bind.TransactOpts, hashedpw [32]byte) (*types.Transaction, error) {
	return _NodeRegistrationImpl.contract.Transact(opts, "setHashedAdminPW", hashedpw)
}

// SetHashedAdminPW is a paid mutator transaction binding the contract method 0x131bea52.
//
// Solidity: function setHashedAdminPW(bytes32 hashedpw) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplSession) SetHashedAdminPW(hashedpw [32]byte) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.SetHashedAdminPW(&_NodeRegistrationImpl.TransactOpts, hashedpw)
}

// SetHashedAdminPW is a paid mutator transaction binding the contract method 0x131bea52.
//
// Solidity: function setHashedAdminPW(bytes32 hashedpw) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorSession) SetHashedAdminPW(hashedpw [32]byte) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.SetHashedAdminPW(&_NodeRegistrationImpl.TransactOpts, hashedpw)
}

// SettingOtpHashValue is a paid mutator transaction binding the contract method 0x72f9013c.
//
// Solidity: function settingOtpHashValue(address applicantAddress, bytes32 otpHashValue) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactor) SettingOtpHashValue(opts *bind.TransactOpts, applicantAddress common.Address, otpHashValue [32]byte) (*types.Transaction, error) {
	return _NodeRegistrationImpl.contract.Transact(opts, "settingOtpHashValue", applicantAddress, otpHashValue)
}

// SettingOtpHashValue is a paid mutator transaction binding the contract method 0x72f9013c.
//
// Solidity: function settingOtpHashValue(address applicantAddress, bytes32 otpHashValue) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplSession) SettingOtpHashValue(applicantAddress common.Address, otpHashValue [32]byte) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.SettingOtpHashValue(&_NodeRegistrationImpl.TransactOpts, applicantAddress, otpHashValue)
}

// SettingOtpHashValue is a paid mutator transaction binding the contract method 0x72f9013c.
//
// Solidity: function settingOtpHashValue(address applicantAddress, bytes32 otpHashValue) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorSession) SettingOtpHashValue(applicantAddress common.Address, otpHashValue [32]byte) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.SettingOtpHashValue(&_NodeRegistrationImpl.TransactOpts, applicantAddress, otpHashValue)
}

// VerifyOtp is a paid mutator transaction binding the contract method 0x9d544042.
//
// Solidity: function verifyOtp(string otp) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactor) VerifyOtp(opts *bind.TransactOpts, otp string) (*types.Transaction, error) {
	return _NodeRegistrationImpl.contract.Transact(opts, "verifyOtp", otp)
}

// VerifyOtp is a paid mutator transaction binding the contract method 0x9d544042.
//
// Solidity: function verifyOtp(string otp) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplSession) VerifyOtp(otp string) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.VerifyOtp(&_NodeRegistrationImpl.TransactOpts, otp)
}

// VerifyOtp is a paid mutator transaction binding the contract method 0x9d544042.
//
// Solidity: function verifyOtp(string otp) returns()
func (_NodeRegistrationImpl *NodeRegistrationImplTransactorSession) VerifyOtp(otp string) (*types.Transaction, error) {
	return _NodeRegistrationImpl.Contract.VerifyOtp(&_NodeRegistrationImpl.TransactOpts, otp)
}

// NodeRegistrationInterfaceABI is the input ABI used to generate the binding from.
const NodeRegistrationInterfaceABI = "[{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeRegistrationInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var NodeRegistrationInterfaceFuncSigs = map[string]string{
	"f851a440": "admin()",
}

// NodeRegistrationInterfaceBin is the compiled bytecode used for deploying new contracts.
var NodeRegistrationInterfaceBin = "0x6080604052348015600f57600080fd5b5060948061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063f851a44014602d575b600080fd5b6033604f565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b03168156fea26469706673582212202667c3f6d8d11221e78dc2f0de81241c68f862b37844273e8836e635194886d664736f6c63430007040033"

// DeployNodeRegistrationInterface deploys a new Ethereum contract, binding an instance of NodeRegistrationInterface to it.
func DeployNodeRegistrationInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeRegistrationInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRegistrationInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeRegistrationInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeRegistrationInterface{NodeRegistrationInterfaceCaller: NodeRegistrationInterfaceCaller{contract: contract}, NodeRegistrationInterfaceTransactor: NodeRegistrationInterfaceTransactor{contract: contract}, NodeRegistrationInterfaceFilterer: NodeRegistrationInterfaceFilterer{contract: contract}}, nil
}

// NodeRegistrationInterface is an auto generated Go binding around an Ethereum contract.
type NodeRegistrationInterface struct {
	NodeRegistrationInterfaceCaller     // Read-only binding to the contract
	NodeRegistrationInterfaceTransactor // Write-only binding to the contract
	NodeRegistrationInterfaceFilterer   // Log filterer for contract events
}

// NodeRegistrationInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRegistrationInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrationInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRegistrationInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrationInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRegistrationInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrationInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRegistrationInterfaceSession struct {
	Contract     *NodeRegistrationInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NodeRegistrationInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRegistrationInterfaceCallerSession struct {
	Contract *NodeRegistrationInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// NodeRegistrationInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRegistrationInterfaceTransactorSession struct {
	Contract     *NodeRegistrationInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// NodeRegistrationInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRegistrationInterfaceRaw struct {
	Contract *NodeRegistrationInterface // Generic contract binding to access the raw methods on
}

// NodeRegistrationInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRegistrationInterfaceCallerRaw struct {
	Contract *NodeRegistrationInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRegistrationInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRegistrationInterfaceTransactorRaw struct {
	Contract *NodeRegistrationInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRegistrationInterface creates a new instance of NodeRegistrationInterface, bound to a specific deployed contract.
func NewNodeRegistrationInterface(address common.Address, backend bind.ContractBackend) (*NodeRegistrationInterface, error) {
	contract, err := bindNodeRegistrationInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationInterface{NodeRegistrationInterfaceCaller: NodeRegistrationInterfaceCaller{contract: contract}, NodeRegistrationInterfaceTransactor: NodeRegistrationInterfaceTransactor{contract: contract}, NodeRegistrationInterfaceFilterer: NodeRegistrationInterfaceFilterer{contract: contract}}, nil
}

// NewNodeRegistrationInterfaceCaller creates a new read-only instance of NodeRegistrationInterface, bound to a specific deployed contract.
func NewNodeRegistrationInterfaceCaller(address common.Address, caller bind.ContractCaller) (*NodeRegistrationInterfaceCaller, error) {
	contract, err := bindNodeRegistrationInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationInterfaceCaller{contract: contract}, nil
}

// NewNodeRegistrationInterfaceTransactor creates a new write-only instance of NodeRegistrationInterface, bound to a specific deployed contract.
func NewNodeRegistrationInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRegistrationInterfaceTransactor, error) {
	contract, err := bindNodeRegistrationInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationInterfaceTransactor{contract: contract}, nil
}

// NewNodeRegistrationInterfaceFilterer creates a new log filterer instance of NodeRegistrationInterface, bound to a specific deployed contract.
func NewNodeRegistrationInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRegistrationInterfaceFilterer, error) {
	contract, err := bindNodeRegistrationInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRegistrationInterfaceFilterer{contract: contract}, nil
}

// bindNodeRegistrationInterface binds a generic wrapper to an already deployed contract.
func bindNodeRegistrationInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRegistrationInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRegistrationInterface *NodeRegistrationInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRegistrationInterface.Contract.NodeRegistrationInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRegistrationInterface *NodeRegistrationInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistrationInterface.Contract.NodeRegistrationInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRegistrationInterface *NodeRegistrationInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRegistrationInterface.Contract.NodeRegistrationInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRegistrationInterface *NodeRegistrationInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRegistrationInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRegistrationInterface *NodeRegistrationInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistrationInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRegistrationInterface *NodeRegistrationInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRegistrationInterface.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NodeRegistrationInterface *NodeRegistrationInterfaceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRegistrationInterface.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NodeRegistrationInterface *NodeRegistrationInterfaceSession) Admin() (common.Address, error) {
	return _NodeRegistrationInterface.Contract.Admin(&_NodeRegistrationInterface.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NodeRegistrationInterface *NodeRegistrationInterfaceCallerSession) Admin() (common.Address, error) {
	return _NodeRegistrationInterface.Contract.Admin(&_NodeRegistrationInterface.CallOpts)
}
