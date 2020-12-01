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

// RulesABI is the input ABI used to generate the binding from.
const RulesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"exitReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByIndex\",\"outputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"nodeType\",\"type\":\"uint8\"},{\"name\":\"geoHash\",\"type\":\"bytes6\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"organization\",\"type\":\"string\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"group\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_groupSource\",\"type\":\"bytes32\"},{\"name\":\"_groupDestination\",\"type\":\"bytes32\"}],\"name\":\"addConnection\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"nodeType\",\"type\":\"uint8\"},{\"name\":\"geoHash\",\"type\":\"bytes6\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"organization\",\"type\":\"string\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"group\",\"type\":\"bytes32\"}],\"name\":\"addEnode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sourceEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"sourceEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"sourceEnodePort\",\"type\":\"uint16\"},{\"name\":\"destinationEnodeHigh\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeLow\",\"type\":\"bytes32\"},{\"name\":\"destinationEnodeIp\",\"type\":\"bytes16\"},{\"name\":\"destinationEnodePort\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_groupSource\",\"type\":\"bytes32\"},{\"name\":\"_groupDestination\",\"type\":\"bytes32\"}],\"name\":\"removeConnection\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowlist\",\"outputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"nodeType\",\"type\":\"uint8\"},{\"name\":\"geoHash\",\"type\":\"bytes6\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"organization\",\"type\":\"string\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"group\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"name\":\"ip\",\"type\":\"bytes16\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"removeEnode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addsRestrictions\",\"type\":\"bool\"}],\"name\":\"triggerRulesChangeEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enodeHigh\",\"type\":\"bytes32\"},{\"name\":\"_enodeLow\",\"type\":\"bytes32\"},{\"name\":\"_ip\",\"type\":\"bytes16\"},{\"name\":\"_port\",\"type\":\"uint16\"}],\"name\":\"getByEnode\",\"outputs\":[{\"name\":\"nodeType\",\"type\":\"uint8\"},{\"name\":\"geoHash\",\"type\":\"bytes6\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"organization\",\"type\":\"string\"},{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"group\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enterReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nodeIngressAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeIp\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"enodePort\",\"type\":\"uint16\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeRemoved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"enodeHigh\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeLow\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"enodeIp\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"enodePort\",\"type\":\"uint16\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"}]"

// Rules is an auto generated Go binding around an Ethereum contract.
type Rules struct {
	RulesCaller     // Read-only binding to the contract
	RulesTransactor // Write-only binding to the contract
	RulesFilterer   // Log filterer for contract events
}

// RulesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RulesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RulesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RulesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RulesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RulesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RulesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RulesSession struct {
	Contract     *Rules            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RulesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RulesCallerSession struct {
	Contract *RulesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RulesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RulesTransactorSession struct {
	Contract     *RulesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RulesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RulesRaw struct {
	Contract *Rules // Generic contract binding to access the raw methods on
}

// RulesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RulesCallerRaw struct {
	Contract *RulesCaller // Generic read-only contract binding to access the raw methods on
}

// RulesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RulesTransactorRaw struct {
	Contract *RulesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRules creates a new instance of Rules, bound to a specific deployed contract.
func NewRules(address common.Address, backend bind.ContractBackend) (*Rules, error) {
	contract, err := bindRules(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rules{RulesCaller: RulesCaller{contract: contract}, RulesTransactor: RulesTransactor{contract: contract}, RulesFilterer: RulesFilterer{contract: contract}}, nil
}

// NewRulesCaller creates a new read-only instance of Rules, bound to a specific deployed contract.
func NewRulesCaller(address common.Address, caller bind.ContractCaller) (*RulesCaller, error) {
	contract, err := bindRules(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RulesCaller{contract: contract}, nil
}

// NewRulesTransactor creates a new write-only instance of Rules, bound to a specific deployed contract.
func NewRulesTransactor(address common.Address, transactor bind.ContractTransactor) (*RulesTransactor, error) {
	contract, err := bindRules(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RulesTransactor{contract: contract}, nil
}

// NewRulesFilterer creates a new log filterer instance of Rules, bound to a specific deployed contract.
func NewRulesFilterer(address common.Address, filterer bind.ContractFilterer) (*RulesFilterer, error) {
	contract, err := bindRules(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RulesFilterer{contract: contract}, nil
}

// bindRules binds a generic wrapper to an already deployed contract.
func bindRules(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RulesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rules *RulesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rules.Contract.RulesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rules *RulesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rules.Contract.RulesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rules *RulesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rules.Contract.RulesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rules *RulesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rules.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rules *RulesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rules.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rules *RulesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rules.Contract.contract.Transact(opts, method, params...)
}

// Allowlist is a free data retrieval call binding the contract method 0x69c45824.
//
// Solidity: function allowlist(uint256 ) view returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesCaller) Allowlist(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EnodeHigh    [32]byte
	EnodeLow     [32]byte
	Ip           [16]byte
	Port         uint16
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	ret := new(struct {
		EnodeHigh    [32]byte
		EnodeLow     [32]byte
		Ip           [16]byte
		Port         uint16
		NodeType     uint8
		GeoHash      [6]byte
		Name         string
		Organization string
		Did          string
		Group        [32]byte
	})
	out := ret
	err := _Rules.contract.Call(opts, out, "allowlist", arg0)
	return *ret, err
}

// Allowlist is a free data retrieval call binding the contract method 0x69c45824.
//
// Solidity: function allowlist(uint256 ) view returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesSession) Allowlist(arg0 *big.Int) (struct {
	EnodeHigh    [32]byte
	EnodeLow     [32]byte
	Ip           [16]byte
	Port         uint16
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	return _Rules.Contract.Allowlist(&_Rules.CallOpts, arg0)
}

// Allowlist is a free data retrieval call binding the contract method 0x69c45824.
//
// Solidity: function allowlist(uint256 ) view returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesCallerSession) Allowlist(arg0 *big.Int) (struct {
	EnodeHigh    [32]byte
	EnodeLow     [32]byte
	Ip           [16]byte
	Port         uint16
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	return _Rules.Contract.Allowlist(&_Rules.CallOpts, arg0)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) view returns(bytes32)
func (_Rules *RulesCaller) ConnectionAllowed(opts *bind.CallOpts, sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rules.contract.Call(opts, out, "connectionAllowed", sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
	return *ret0, err
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) view returns(bytes32)
func (_Rules *RulesSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _Rules.Contract.ConnectionAllowed(&_Rules.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x3620b1df.
//
// Solidity: function connectionAllowed(bytes32 sourceEnodeHigh, bytes32 sourceEnodeLow, bytes16 sourceEnodeIp, uint16 sourceEnodePort, bytes32 destinationEnodeHigh, bytes32 destinationEnodeLow, bytes16 destinationEnodeIp, uint16 destinationEnodePort) view returns(bytes32)
func (_Rules *RulesCallerSession) ConnectionAllowed(sourceEnodeHigh [32]byte, sourceEnodeLow [32]byte, sourceEnodeIp [16]byte, sourceEnodePort uint16, destinationEnodeHigh [32]byte, destinationEnodeLow [32]byte, destinationEnodeIp [16]byte, destinationEnodePort uint16) ([32]byte, error) {
	return _Rules.Contract.ConnectionAllowed(&_Rules.CallOpts, sourceEnodeHigh, sourceEnodeLow, sourceEnodeIp, sourceEnodePort, destinationEnodeHigh, destinationEnodeLow, destinationEnodeIp, destinationEnodePort)
}

// GetByEnode is a free data retrieval call binding the contract method 0xcefac681.
//
// Solidity: function getByEnode(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port) view returns(uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesCaller) GetByEnode(opts *bind.CallOpts, _enodeHigh [32]byte, _enodeLow [32]byte, _ip [16]byte, _port uint16) (struct {
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	ret := new(struct {
		NodeType     uint8
		GeoHash      [6]byte
		Name         string
		Organization string
		Did          string
		Group        [32]byte
	})
	out := ret
	err := _Rules.contract.Call(opts, out, "getByEnode", _enodeHigh, _enodeLow, _ip, _port)
	return *ret, err
}

// GetByEnode is a free data retrieval call binding the contract method 0xcefac681.
//
// Solidity: function getByEnode(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port) view returns(uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesSession) GetByEnode(_enodeHigh [32]byte, _enodeLow [32]byte, _ip [16]byte, _port uint16) (struct {
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	return _Rules.Contract.GetByEnode(&_Rules.CallOpts, _enodeHigh, _enodeLow, _ip, _port)
}

// GetByEnode is a free data retrieval call binding the contract method 0xcefac681.
//
// Solidity: function getByEnode(bytes32 _enodeHigh, bytes32 _enodeLow, bytes16 _ip, uint16 _port) view returns(uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesCallerSession) GetByEnode(_enodeHigh [32]byte, _enodeLow [32]byte, _ip [16]byte, _port uint16) (struct {
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	return _Rules.Contract.GetByEnode(&_Rules.CallOpts, _enodeHigh, _enodeLow, _ip, _port)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) view returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesCaller) GetByIndex(opts *bind.CallOpts, index *big.Int) (struct {
	EnodeHigh    [32]byte
	EnodeLow     [32]byte
	Ip           [16]byte
	Port         uint16
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	ret := new(struct {
		EnodeHigh    [32]byte
		EnodeLow     [32]byte
		Ip           [16]byte
		Port         uint16
		NodeType     uint8
		GeoHash      [6]byte
		Name         string
		Organization string
		Did          string
		Group        [32]byte
	})
	out := ret
	err := _Rules.contract.Call(opts, out, "getByIndex", index)
	return *ret, err
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) view returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesSession) GetByIndex(index *big.Int) (struct {
	EnodeHigh    [32]byte
	EnodeLow     [32]byte
	Ip           [16]byte
	Port         uint16
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	return _Rules.Contract.GetByIndex(&_Rules.CallOpts, index)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) view returns(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group)
func (_Rules *RulesCallerSession) GetByIndex(index *big.Int) (struct {
	EnodeHigh    [32]byte
	EnodeLow     [32]byte
	Ip           [16]byte
	Port         uint16
	NodeType     uint8
	GeoHash      [6]byte
	Name         string
	Organization string
	Did          string
	Group        [32]byte
}, error) {
	return _Rules.Contract.GetByIndex(&_Rules.CallOpts, index)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Rules *RulesCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rules.contract.Call(opts, out, "getContractVersion")
	return *ret0, err
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Rules *RulesSession) GetContractVersion() (*big.Int, error) {
	return _Rules.Contract.GetContractVersion(&_Rules.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Rules *RulesCallerSession) GetContractVersion() (*big.Int, error) {
	return _Rules.Contract.GetContractVersion(&_Rules.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Rules *RulesCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rules.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Rules *RulesSession) GetSize() (*big.Int, error) {
	return _Rules.Contract.GetSize(&_Rules.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Rules *RulesCallerSession) GetSize() (*big.Int, error) {
	return _Rules.Contract.GetSize(&_Rules.CallOpts)
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() view returns(bool)
func (_Rules *RulesCaller) IsReadOnly(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rules.contract.Call(opts, out, "isReadOnly")
	return *ret0, err
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() view returns(bool)
func (_Rules *RulesSession) IsReadOnly() (bool, error) {
	return _Rules.Contract.IsReadOnly(&_Rules.CallOpts)
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() view returns(bool)
func (_Rules *RulesCallerSession) IsReadOnly() (bool, error) {
	return _Rules.Contract.IsReadOnly(&_Rules.CallOpts)
}

// AddConnection is a paid mutator transaction binding the contract method 0x2da82e23.
//
// Solidity: function addConnection(bytes32 _groupSource, bytes32 _groupDestination) returns(bool)
func (_Rules *RulesTransactor) AddConnection(opts *bind.TransactOpts, _groupSource [32]byte, _groupDestination [32]byte) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "addConnection", _groupSource, _groupDestination)
}

// AddConnection is a paid mutator transaction binding the contract method 0x2da82e23.
//
// Solidity: function addConnection(bytes32 _groupSource, bytes32 _groupDestination) returns(bool)
func (_Rules *RulesSession) AddConnection(_groupSource [32]byte, _groupDestination [32]byte) (*types.Transaction, error) {
	return _Rules.Contract.AddConnection(&_Rules.TransactOpts, _groupSource, _groupDestination)
}

// AddConnection is a paid mutator transaction binding the contract method 0x2da82e23.
//
// Solidity: function addConnection(bytes32 _groupSource, bytes32 _groupDestination) returns(bool)
func (_Rules *RulesTransactorSession) AddConnection(_groupSource [32]byte, _groupDestination [32]byte) (*types.Transaction, error) {
	return _Rules.Contract.AddConnection(&_Rules.TransactOpts, _groupSource, _groupDestination)
}

// AddEnode is a paid mutator transaction binding the contract method 0x307b0d29.
//
// Solidity: function addEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group) returns(bool)
func (_Rules *RulesTransactor) AddEnode(opts *bind.TransactOpts, enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16, nodeType uint8, geoHash [6]byte, name string, organization string, did string, group [32]byte) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "addEnode", enodeHigh, enodeLow, ip, port, nodeType, geoHash, name, organization, did, group)
}

// AddEnode is a paid mutator transaction binding the contract method 0x307b0d29.
//
// Solidity: function addEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group) returns(bool)
func (_Rules *RulesSession) AddEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16, nodeType uint8, geoHash [6]byte, name string, organization string, did string, group [32]byte) (*types.Transaction, error) {
	return _Rules.Contract.AddEnode(&_Rules.TransactOpts, enodeHigh, enodeLow, ip, port, nodeType, geoHash, name, organization, did, group)
}

// AddEnode is a paid mutator transaction binding the contract method 0x307b0d29.
//
// Solidity: function addEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port, uint8 nodeType, bytes6 geoHash, string name, string organization, string did, bytes32 group) returns(bool)
func (_Rules *RulesTransactorSession) AddEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16, nodeType uint8, geoHash [6]byte, name string, organization string, did string, group [32]byte) (*types.Transaction, error) {
	return _Rules.Contract.AddEnode(&_Rules.TransactOpts, enodeHigh, enodeLow, ip, port, nodeType, geoHash, name, organization, did, group)
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_Rules *RulesTransactor) EnterReadOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "enterReadOnly")
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_Rules *RulesSession) EnterReadOnly() (*types.Transaction, error) {
	return _Rules.Contract.EnterReadOnly(&_Rules.TransactOpts)
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_Rules *RulesTransactorSession) EnterReadOnly() (*types.Transaction, error) {
	return _Rules.Contract.EnterReadOnly(&_Rules.TransactOpts)
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_Rules *RulesTransactor) ExitReadOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "exitReadOnly")
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_Rules *RulesSession) ExitReadOnly() (*types.Transaction, error) {
	return _Rules.Contract.ExitReadOnly(&_Rules.TransactOpts)
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_Rules *RulesTransactorSession) ExitReadOnly() (*types.Transaction, error) {
	return _Rules.Contract.ExitReadOnly(&_Rules.TransactOpts)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0x57d3e790.
//
// Solidity: function removeConnection(bytes32 _groupSource, bytes32 _groupDestination) returns(bool)
func (_Rules *RulesTransactor) RemoveConnection(opts *bind.TransactOpts, _groupSource [32]byte, _groupDestination [32]byte) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "removeConnection", _groupSource, _groupDestination)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0x57d3e790.
//
// Solidity: function removeConnection(bytes32 _groupSource, bytes32 _groupDestination) returns(bool)
func (_Rules *RulesSession) RemoveConnection(_groupSource [32]byte, _groupDestination [32]byte) (*types.Transaction, error) {
	return _Rules.Contract.RemoveConnection(&_Rules.TransactOpts, _groupSource, _groupDestination)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0x57d3e790.
//
// Solidity: function removeConnection(bytes32 _groupSource, bytes32 _groupDestination) returns(bool)
func (_Rules *RulesTransactorSession) RemoveConnection(_groupSource [32]byte, _groupDestination [32]byte) (*types.Transaction, error) {
	return _Rules.Contract.RemoveConnection(&_Rules.TransactOpts, _groupSource, _groupDestination)
}

// RemoveEnode is a paid mutator transaction binding the contract method 0xaab2f5eb.
//
// Solidity: function removeEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_Rules *RulesTransactor) RemoveEnode(opts *bind.TransactOpts, enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "removeEnode", enodeHigh, enodeLow, ip, port)
}

// RemoveEnode is a paid mutator transaction binding the contract method 0xaab2f5eb.
//
// Solidity: function removeEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_Rules *RulesSession) RemoveEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _Rules.Contract.RemoveEnode(&_Rules.TransactOpts, enodeHigh, enodeLow, ip, port)
}

// RemoveEnode is a paid mutator transaction binding the contract method 0xaab2f5eb.
//
// Solidity: function removeEnode(bytes32 enodeHigh, bytes32 enodeLow, bytes16 ip, uint16 port) returns(bool)
func (_Rules *RulesTransactorSession) RemoveEnode(enodeHigh [32]byte, enodeLow [32]byte, ip [16]byte, port uint16) (*types.Transaction, error) {
	return _Rules.Contract.RemoveEnode(&_Rules.TransactOpts, enodeHigh, enodeLow, ip, port)
}

// TriggerRulesChangeEvent is a paid mutator transaction binding the contract method 0xcc3a1c41.
//
// Solidity: function triggerRulesChangeEvent(bool addsRestrictions) returns()
func (_Rules *RulesTransactor) TriggerRulesChangeEvent(opts *bind.TransactOpts, addsRestrictions bool) (*types.Transaction, error) {
	return _Rules.contract.Transact(opts, "triggerRulesChangeEvent", addsRestrictions)
}

// TriggerRulesChangeEvent is a paid mutator transaction binding the contract method 0xcc3a1c41.
//
// Solidity: function triggerRulesChangeEvent(bool addsRestrictions) returns()
func (_Rules *RulesSession) TriggerRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _Rules.Contract.TriggerRulesChangeEvent(&_Rules.TransactOpts, addsRestrictions)
}

// TriggerRulesChangeEvent is a paid mutator transaction binding the contract method 0xcc3a1c41.
//
// Solidity: function triggerRulesChangeEvent(bool addsRestrictions) returns()
func (_Rules *RulesTransactorSession) TriggerRulesChangeEvent(addsRestrictions bool) (*types.Transaction, error) {
	return _Rules.Contract.TriggerRulesChangeEvent(&_Rules.TransactOpts, addsRestrictions)
}

// RulesNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the Rules contract.
type RulesNodeAddedIterator struct {
	Event *RulesNodeAdded // Event containing the contract specifics and raw log

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
func (it *RulesNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RulesNodeAdded)
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
		it.Event = new(RulesNodeAdded)
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
func (it *RulesNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RulesNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RulesNodeAdded represents a NodeAdded event raised by the Rules contract.
type RulesNodeAdded struct {
	NodeAdded bool
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	EnodeIp   [16]byte
	EnodePort uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0x983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea.
//
// Solidity: event NodeAdded(bool nodeAdded, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_Rules *RulesFilterer) FilterNodeAdded(opts *bind.FilterOpts) (*RulesNodeAddedIterator, error) {

	logs, sub, err := _Rules.contract.FilterLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return &RulesNodeAddedIterator{contract: _Rules.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0x983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea.
//
// Solidity: event NodeAdded(bool nodeAdded, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_Rules *RulesFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *RulesNodeAdded) (event.Subscription, error) {

	logs, sub, err := _Rules.contract.WatchLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RulesNodeAdded)
				if err := _Rules.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0x983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea.
//
// Solidity: event NodeAdded(bool nodeAdded, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_Rules *RulesFilterer) ParseNodeAdded(log types.Log) (*RulesNodeAdded, error) {
	event := new(RulesNodeAdded)
	if err := _Rules.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RulesNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the Rules contract.
type RulesNodeRemovedIterator struct {
	Event *RulesNodeRemoved // Event containing the contract specifics and raw log

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
func (it *RulesNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RulesNodeRemoved)
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
		it.Event = new(RulesNodeRemoved)
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
func (it *RulesNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RulesNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RulesNodeRemoved represents a NodeRemoved event raised by the Rules contract.
type RulesNodeRemoved struct {
	NodeRemoved bool
	EnodeHigh   [32]byte
	EnodeLow    [32]byte
	EnodeIp     [16]byte
	EnodePort   uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xf05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca45.
//
// Solidity: event NodeRemoved(bool nodeRemoved, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_Rules *RulesFilterer) FilterNodeRemoved(opts *bind.FilterOpts) (*RulesNodeRemovedIterator, error) {

	logs, sub, err := _Rules.contract.FilterLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return &RulesNodeRemovedIterator{contract: _Rules.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xf05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca45.
//
// Solidity: event NodeRemoved(bool nodeRemoved, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_Rules *RulesFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *RulesNodeRemoved) (event.Subscription, error) {

	logs, sub, err := _Rules.contract.WatchLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RulesNodeRemoved)
				if err := _Rules.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// ParseNodeRemoved is a log parse operation binding the contract event 0xf05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca45.
//
// Solidity: event NodeRemoved(bool nodeRemoved, bytes32 enodeHigh, bytes32 enodeLow, bytes16 enodeIp, uint16 enodePort)
func (_Rules *RulesFilterer) ParseNodeRemoved(log types.Log) (*RulesNodeRemoved, error) {
	event := new(RulesNodeRemoved)
	if err := _Rules.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
