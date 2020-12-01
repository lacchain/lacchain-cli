// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissioning

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

// PermissioningABI is the input ABI used to generate the binding from.
const PermissioningABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllContractKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RULES_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ADMIN_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sourceEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"sourceEnodePort\",\"type\":\"uint16\"},{\"name\":\"destinationEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"destinationEnodePort\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addsRestrictions\",\"type\":\"bool\"}],\"name\":\"emitRulesChangeEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"removeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addsRestrictions\",\"type\":\"bool\"}],\"name\":\"NodePermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractName\",\"type\":\"bytes32\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"}]"

// Permissioning is an auto generated Go binding around an Ethereum contract.
type Permissioning struct {
	PermissioningCaller     // Read-only binding to the contract
	PermissioningTransactor // Write-only binding to the contract
	PermissioningFilterer   // Log filterer for contract events
}

// PermissioningCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissioningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissioningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissioningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissioningFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissioningFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissioningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissioningSession struct {
	Contract     *Permissioning    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissioningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissioningCallerSession struct {
	Contract *PermissioningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PermissioningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissioningTransactorSession struct {
	Contract     *PermissioningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PermissioningRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissioningRaw struct {
	Contract *Permissioning // Generic contract binding to access the raw methods on
}

// PermissioningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissioningCallerRaw struct {
	Contract *PermissioningCaller // Generic read-only contract binding to access the raw methods on
}

// PermissioningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissioningTransactorRaw struct {
	Contract *PermissioningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissioning creates a new instance of Permissioning, bound to a specific deployed contract.
func NewPermissioning(address common.Address, backend bind.ContractBackend) (*Permissioning, error) {
	contract, err := bindPermissioning(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Permissioning{PermissioningCaller: PermissioningCaller{contract: contract}, PermissioningTransactor: PermissioningTransactor{contract: contract}, PermissioningFilterer: PermissioningFilterer{contract: contract}}, nil
}

// NewPermissioningCaller creates a new read-only instance of Permissioning, bound to a specific deployed contract.
func NewPermissioningCaller(address common.Address, caller bind.ContractCaller) (*PermissioningCaller, error) {
	contract, err := bindPermissioning(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissioningCaller{contract: contract}, nil
}

// NewPermissioningTransactor creates a new write-only instance of Permissioning, bound to a specific deployed contract.
func NewPermissioningTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissioningTransactor, error) {
	contract, err := bindPermissioning(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissioningTransactor{contract: contract}, nil
}

// NewPermissioningFilterer creates a new log filterer instance of Permissioning, bound to a specific deployed contract.
func NewPermissioningFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissioningFilterer, error) {
	contract, err := bindPermissioning(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissioningFilterer{contract: contract}, nil
}

// bindPermissioning binds a generic wrapper to an already deployed contract.
func bindPermissioning(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissioningABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissioning *PermissioningRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissioning.Contract.PermissioningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissioning *PermissioningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissioning.Contract.PermissioningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissioning *PermissioningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissioning.Contract.PermissioningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissioning *PermissioningCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissioning.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissioning *PermissioningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissioning.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissioning *PermissioningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissioning.Contract.contract.Transact(opts, method, params...)
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() view returns(bytes32)
func (_Permissioning *PermissioningCaller) ADMINCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "ADMIN_CONTRACT")
	return *ret0, err
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() view returns(bytes32)
func (_Permissioning *PermissioningSession) ADMINCONTRACT() ([32]byte, error) {
	return _Permissioning.Contract.ADMINCONTRACT(&_Permissioning.CallOpts)
}

// ADMINCONTRACT is a free data retrieval call binding the contract method 0x1e7c27cb.
//
// Solidity: function ADMIN_CONTRACT() view returns(bytes32)
func (_Permissioning *PermissioningCallerSession) ADMINCONTRACT() ([32]byte, error) {
	return _Permissioning.Contract.ADMINCONTRACT(&_Permissioning.CallOpts)
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() view returns(bytes32)
func (_Permissioning *PermissioningCaller) RULESCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "RULES_CONTRACT")
	return *ret0, err
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() view returns(bytes32)
func (_Permissioning *PermissioningSession) RULESCONTRACT() ([32]byte, error) {
	return _Permissioning.Contract.RULESCONTRACT(&_Permissioning.CallOpts)
}

// RULESCONTRACT is a free data retrieval call binding the contract method 0x11601306.
//
// Solidity: function RULES_CONTRACT() view returns(bytes32)
func (_Permissioning *PermissioningCallerSession) RULESCONTRACT() ([32]byte, error) {
	return _Permissioning.Contract.RULESCONTRACT(&_Permissioning.CallOpts)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) view returns(bytes32)
func (_Permissioning *PermissioningCaller) ConnectionAllowed(opts *bind.CallOpts, sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "connectionAllowed", sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
	return *ret0, err
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) view returns(bytes32)
func (_Permissioning *PermissioningSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _Permissioning.Contract.ConnectionAllowed(&_Permissioning.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) view returns(bytes32)
func (_Permissioning *PermissioningCallerSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _Permissioning.Contract.ConnectionAllowed(&_Permissioning.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() view returns(bytes32[])
func (_Permissioning *PermissioningCaller) GetAllContractKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "getAllContractKeys")
	return *ret0, err
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() view returns(bytes32[])
func (_Permissioning *PermissioningSession) GetAllContractKeys() ([][32]byte, error) {
	return _Permissioning.Contract.GetAllContractKeys(&_Permissioning.CallOpts)
}

// GetAllContractKeys is a free data retrieval call binding the contract method 0x10d9042e.
//
// Solidity: function getAllContractKeys() view returns(bytes32[])
func (_Permissioning *PermissioningCallerSession) GetAllContractKeys() ([][32]byte, error) {
	return _Permissioning.Contract.GetAllContractKeys(&_Permissioning.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) view returns(address)
func (_Permissioning *PermissioningCaller) GetContractAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "getContractAddress", name)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) view returns(address)
func (_Permissioning *PermissioningSession) GetContractAddress(name [32]byte) (common.Address, error) {
	return _Permissioning.Contract.GetContractAddress(&_Permissioning.CallOpts, name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(bytes32 name) view returns(address)
func (_Permissioning *PermissioningCallerSession) GetContractAddress(name [32]byte) (common.Address, error) {
	return _Permissioning.Contract.GetContractAddress(&_Permissioning.CallOpts, name)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Permissioning *PermissioningCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "getContractVersion")
	return *ret0, err
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Permissioning *PermissioningSession) GetContractVersion() (*big.Int, error) {
	return _Permissioning.Contract.GetContractVersion(&_Permissioning.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Permissioning *PermissioningCallerSession) GetContractVersion() (*big.Int, error) {
	return _Permissioning.Contract.GetContractVersion(&_Permissioning.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Permissioning *PermissioningCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Permissioning *PermissioningSession) GetSize() (*big.Int, error) {
	return _Permissioning.Contract.GetSize(&_Permissioning.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Permissioning *PermissioningCallerSession) GetSize() (*big.Int, error) {
	return _Permissioning.Contract.GetSize(&_Permissioning.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) view returns(bool)
func (_Permissioning *PermissioningCaller) IsAuthorized(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "isAuthorized", account)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) view returns(bool)
func (_Permissioning *PermissioningSession) IsAuthorized(account common.Address) (bool, error) {
	return _Permissioning.Contract.IsAuthorized(&_Permissioning.CallOpts, account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) view returns(bool)
func (_Permissioning *PermissioningCallerSession) IsAuthorized(account common.Address) (bool, error) {
	return _Permissioning.Contract.IsAuthorized(&_Permissioning.CallOpts, account)
}

// EmitRulesChangeEvent is a paid mutator transaction binding the contract method 0x4dc3fefc.
//
// Solidity: function emitRulesChangeEvent(bool addsRestrictions) returns()
func (_Permissioning *PermissioningTransactor) EmitRulesChangeEvent(opts *bind.TransactOpts, addsRestrictions bool) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "emitRulesChangeEvent", addsRestrictions)
}

// EmitRulesChangeEvent is a paid mutator transaction binding the contract method 0x4dc3fefc.
//
// Solidity: function emitRulesChangeEvent(bool addsRestrictions) returns()
func (_Permissioning *PermissioningSession) EmitRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _Permissioning.Contract.EmitRulesChangeEvent(&_Permissioning.TransactOpts, addsRestrictions)
}

// EmitRulesChangeEvent is a paid mutator transaction binding the contract method 0x4dc3fefc.
//
// Solidity: function emitRulesChangeEvent(bool addsRestrictions) returns()
func (_Permissioning *PermissioningTransactorSession) EmitRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _Permissioning.Contract.EmitRulesChangeEvent(&_Permissioning.TransactOpts, addsRestrictions)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_Permissioning *PermissioningTransactor) RemoveContract(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "removeContract", _name)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_Permissioning *PermissioningSession) RemoveContract(_name [32]byte) (*types.Transaction, error) {
	return _Permissioning.Contract.RemoveContract(&_Permissioning.TransactOpts, _name)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xa43e04d8.
//
// Solidity: function removeContract(bytes32 _name) returns(bool)
func (_Permissioning *PermissioningTransactorSession) RemoveContract(_name [32]byte) (*types.Transaction, error) {
	return _Permissioning.Contract.RemoveContract(&_Permissioning.TransactOpts, _name)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_Permissioning *PermissioningTransactor) SetContractAddress(opts *bind.TransactOpts, name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "setContractAddress", name, addr)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_Permissioning *PermissioningSession) SetContractAddress(name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Permissioning.Contract.SetContractAddress(&_Permissioning.TransactOpts, name, addr)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xe001f841.
//
// Solidity: function setContractAddress(bytes32 name, address addr) returns(bool)
func (_Permissioning *PermissioningTransactorSession) SetContractAddress(name [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Permissioning.Contract.SetContractAddress(&_Permissioning.TransactOpts, name, addr)
}

// PermissioningNodePermissionsUpdatedIterator is returned from FilterNodePermissionsUpdated and is used to iterate over the raw logs and unpacked data for NodePermissionsUpdated events raised by the Permissioning contract.
type PermissioningNodePermissionsUpdatedIterator struct {
	Event *PermissioningNodePermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *PermissioningNodePermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissioningNodePermissionsUpdated)
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
		it.Event = new(PermissioningNodePermissionsUpdated)
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
func (it *PermissioningNodePermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissioningNodePermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissioningNodePermissionsUpdated represents a NodePermissionsUpdated event raised by the Permissioning contract.
type PermissioningNodePermissionsUpdated struct {
	AddsRestrictions bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNodePermissionsUpdated is a free log retrieval operation binding the contract event 0x66120f934b66d52127e448f8e94c2460ea62821335e0dd18e89ed38a4a09b413.
//
// Solidity: event NodePermissionsUpdated(bool addsRestrictions)
func (_Permissioning *PermissioningFilterer) FilterNodePermissionsUpdated(opts *bind.FilterOpts) (*PermissioningNodePermissionsUpdatedIterator, error) {

	logs, sub, err := _Permissioning.contract.FilterLogs(opts, "NodePermissionsUpdated")
	if err != nil {
		return nil, err
	}
	return &PermissioningNodePermissionsUpdatedIterator{contract: _Permissioning.contract, event: "NodePermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchNodePermissionsUpdated is a free log subscription operation binding the contract event 0x66120f934b66d52127e448f8e94c2460ea62821335e0dd18e89ed38a4a09b413.
//
// Solidity: event NodePermissionsUpdated(bool addsRestrictions)
func (_Permissioning *PermissioningFilterer) WatchNodePermissionsUpdated(opts *bind.WatchOpts, sink chan<- *PermissioningNodePermissionsUpdated) (event.Subscription, error) {

	logs, sub, err := _Permissioning.contract.WatchLogs(opts, "NodePermissionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissioningNodePermissionsUpdated)
				if err := _Permissioning.contract.UnpackLog(event, "NodePermissionsUpdated", log); err != nil {
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

// ParseNodePermissionsUpdated is a log parse operation binding the contract event 0x66120f934b66d52127e448f8e94c2460ea62821335e0dd18e89ed38a4a09b413.
//
// Solidity: event NodePermissionsUpdated(bool addsRestrictions)
func (_Permissioning *PermissioningFilterer) ParseNodePermissionsUpdated(log types.Log) (*PermissioningNodePermissionsUpdated, error) {
	event := new(PermissioningNodePermissionsUpdated)
	if err := _Permissioning.contract.UnpackLog(event, "NodePermissionsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PermissioningRegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the Permissioning contract.
type PermissioningRegistryUpdatedIterator struct {
	Event *PermissioningRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *PermissioningRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissioningRegistryUpdated)
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
		it.Event = new(PermissioningRegistryUpdated)
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
func (it *PermissioningRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissioningRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissioningRegistryUpdated represents a RegistryUpdated event raised by the Permissioning contract.
type PermissioningRegistryUpdated struct {
	ContractAddress common.Address
	ContractName    [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_Permissioning *PermissioningFilterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*PermissioningRegistryUpdatedIterator, error) {

	logs, sub, err := _Permissioning.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &PermissioningRegistryUpdatedIterator{contract: _Permissioning.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_Permissioning *PermissioningFilterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *PermissioningRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Permissioning.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissioningRegistryUpdated)
				if err := _Permissioning.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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

// ParseRegistryUpdated is a log parse operation binding the contract event 0xe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af.
//
// Solidity: event RegistryUpdated(address contractAddress, bytes32 contractName)
func (_Permissioning *PermissioningFilterer) ParseRegistryUpdated(log types.Log) (*PermissioningRegistryUpdated, error) {
	event := new(PermissioningRegistryUpdated)
	if err := _Permissioning.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
