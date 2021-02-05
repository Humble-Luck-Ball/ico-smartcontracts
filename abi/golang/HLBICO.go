// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HLBICO

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

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c0d7e0b46268c2a522edd4ee652cd04f62d3d8a7cfc8fa9d658e5d30619608a764736f6c63430007060033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// CappedTimedCrowdsaleABI is the input ABI used to generate the binding from.
const CappedTimedCrowdsaleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CappedTimedCrowdsaleFuncSigs maps the 4-byte function signature to its string representation.
var CappedTimedCrowdsaleFuncSigs = map[string]string{
	"ec8ac4d8": "buyTokens(address)",
	"355274ea": "cap()",
	"4f935945": "capReached()",
	"4b6753bc": "closingTime()",
	"1515bc2b": "hasClosed()",
	"47535d7b": "isOpen()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
}

// CappedTimedCrowdsale is an auto generated Go binding around an Ethereum contract.
type CappedTimedCrowdsale struct {
	CappedTimedCrowdsaleCaller     // Read-only binding to the contract
	CappedTimedCrowdsaleTransactor // Write-only binding to the contract
	CappedTimedCrowdsaleFilterer   // Log filterer for contract events
}

// CappedTimedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedTimedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTimedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedTimedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTimedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedTimedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedTimedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedTimedCrowdsaleSession struct {
	Contract     *CappedTimedCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CappedTimedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedTimedCrowdsaleCallerSession struct {
	Contract *CappedTimedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CappedTimedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedTimedCrowdsaleTransactorSession struct {
	Contract     *CappedTimedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CappedTimedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedTimedCrowdsaleRaw struct {
	Contract *CappedTimedCrowdsale // Generic contract binding to access the raw methods on
}

// CappedTimedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedTimedCrowdsaleCallerRaw struct {
	Contract *CappedTimedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CappedTimedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedTimedCrowdsaleTransactorRaw struct {
	Contract *CappedTimedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedTimedCrowdsale creates a new instance of CappedTimedCrowdsale, bound to a specific deployed contract.
func NewCappedTimedCrowdsale(address common.Address, backend bind.ContractBackend) (*CappedTimedCrowdsale, error) {
	contract, err := bindCappedTimedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedTimedCrowdsale{CappedTimedCrowdsaleCaller: CappedTimedCrowdsaleCaller{contract: contract}, CappedTimedCrowdsaleTransactor: CappedTimedCrowdsaleTransactor{contract: contract}, CappedTimedCrowdsaleFilterer: CappedTimedCrowdsaleFilterer{contract: contract}}, nil
}

// NewCappedTimedCrowdsaleCaller creates a new read-only instance of CappedTimedCrowdsale, bound to a specific deployed contract.
func NewCappedTimedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CappedTimedCrowdsaleCaller, error) {
	contract, err := bindCappedTimedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedTimedCrowdsaleCaller{contract: contract}, nil
}

// NewCappedTimedCrowdsaleTransactor creates a new write-only instance of CappedTimedCrowdsale, bound to a specific deployed contract.
func NewCappedTimedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedTimedCrowdsaleTransactor, error) {
	contract, err := bindCappedTimedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedTimedCrowdsaleTransactor{contract: contract}, nil
}

// NewCappedTimedCrowdsaleFilterer creates a new log filterer instance of CappedTimedCrowdsale, bound to a specific deployed contract.
func NewCappedTimedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedTimedCrowdsaleFilterer, error) {
	contract, err := bindCappedTimedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedTimedCrowdsaleFilterer{contract: contract}, nil
}

// bindCappedTimedCrowdsale binds a generic wrapper to an already deployed contract.
func bindCappedTimedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedTimedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedTimedCrowdsale.Contract.CappedTimedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.CappedTimedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.CappedTimedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedTimedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) Cap() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.Cap(&_CappedTimedCrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.Cap(&_CappedTimedCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) CapReached() (bool, error) {
	return _CappedTimedCrowdsale.Contract.CapReached(&_CappedTimedCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) CapReached() (bool, error) {
	return _CappedTimedCrowdsale.Contract.CapReached(&_CappedTimedCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.ClosingTime(&_CappedTimedCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.ClosingTime(&_CappedTimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) HasClosed() (bool, error) {
	return _CappedTimedCrowdsale.Contract.HasClosed(&_CappedTimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _CappedTimedCrowdsale.Contract.HasClosed(&_CappedTimedCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) IsOpen() (bool, error) {
	return _CappedTimedCrowdsale.Contract.IsOpen(&_CappedTimedCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _CappedTimedCrowdsale.Contract.IsOpen(&_CappedTimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.OpeningTime(&_CappedTimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.OpeningTime(&_CappedTimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) Rate() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.Rate(&_CappedTimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.Rate(&_CappedTimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) Token() (common.Address, error) {
	return _CappedTimedCrowdsale.Contract.Token(&_CappedTimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _CappedTimedCrowdsale.Contract.Token(&_CappedTimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) Wallet() (common.Address, error) {
	return _CappedTimedCrowdsale.Contract.Wallet(&_CappedTimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _CappedTimedCrowdsale.Contract.Wallet(&_CappedTimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedTimedCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.WeiRaised(&_CappedTimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _CappedTimedCrowdsale.Contract.WeiRaised(&_CappedTimedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.BuyTokens(&_CappedTimedCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.BuyTokens(&_CappedTimedCrowdsale.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.Fallback(&_CappedTimedCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.Fallback(&_CappedTimedCrowdsale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedTimedCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.Receive(&_CappedTimedCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _CappedTimedCrowdsale.Contract.Receive(&_CappedTimedCrowdsale.TransactOpts)
}

// CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the CappedTimedCrowdsale contract.
type CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *CappedTimedCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTimedCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(CappedTimedCrowdsaleTimedCrowdsaleExtended)
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
func (it *CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTimedCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the CappedTimedCrowdsale contract.
type CappedTimedCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _CappedTimedCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &CappedTimedCrowdsaleTimedCrowdsaleExtendedIterator{contract: _CappedTimedCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *CappedTimedCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _CappedTimedCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTimedCrowdsaleTimedCrowdsaleExtended)
				if err := _CappedTimedCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*CappedTimedCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(CappedTimedCrowdsaleTimedCrowdsaleExtended)
	if err := _CappedTimedCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedTimedCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the CappedTimedCrowdsale contract.
type CappedTimedCrowdsaleTokensPurchasedIterator struct {
	Event *CappedTimedCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *CappedTimedCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedTimedCrowdsaleTokensPurchased)
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
		it.Event = new(CappedTimedCrowdsaleTokensPurchased)
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
func (it *CappedTimedCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedTimedCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedTimedCrowdsaleTokensPurchased represents a TokensPurchased event raised by the CappedTimedCrowdsale contract.
type CappedTimedCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*CappedTimedCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedTimedCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CappedTimedCrowdsaleTokensPurchasedIterator{contract: _CappedTimedCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *CappedTimedCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedTimedCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedTimedCrowdsaleTokensPurchased)
				if err := _CappedTimedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CappedTimedCrowdsale *CappedTimedCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*CappedTimedCrowdsaleTokensPurchased, error) {
	event := new(CappedTimedCrowdsaleTokensPurchased)
	if err := _CappedTimedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalEscrowABI is the input ABI used to generate the binding from.
const ConditionalEscrowABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"depositsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawalAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ConditionalEscrowFuncSigs maps the 4-byte function signature to its string representation.
var ConditionalEscrowFuncSigs = map[string]string{
	"f340fa01": "deposit(address)",
	"e3a9db1a": "depositsOf(address)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"51cff8d9": "withdraw(address)",
	"685ca194": "withdrawalAllowed(address)",
}

// ConditionalEscrow is an auto generated Go binding around an Ethereum contract.
type ConditionalEscrow struct {
	ConditionalEscrowCaller     // Read-only binding to the contract
	ConditionalEscrowTransactor // Write-only binding to the contract
	ConditionalEscrowFilterer   // Log filterer for contract events
}

// ConditionalEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConditionalEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionalEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConditionalEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionalEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConditionalEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionalEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConditionalEscrowSession struct {
	Contract     *ConditionalEscrow // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConditionalEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConditionalEscrowCallerSession struct {
	Contract *ConditionalEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConditionalEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConditionalEscrowTransactorSession struct {
	Contract     *ConditionalEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConditionalEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConditionalEscrowRaw struct {
	Contract *ConditionalEscrow // Generic contract binding to access the raw methods on
}

// ConditionalEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConditionalEscrowCallerRaw struct {
	Contract *ConditionalEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// ConditionalEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConditionalEscrowTransactorRaw struct {
	Contract *ConditionalEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConditionalEscrow creates a new instance of ConditionalEscrow, bound to a specific deployed contract.
func NewConditionalEscrow(address common.Address, backend bind.ContractBackend) (*ConditionalEscrow, error) {
	contract, err := bindConditionalEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrow{ConditionalEscrowCaller: ConditionalEscrowCaller{contract: contract}, ConditionalEscrowTransactor: ConditionalEscrowTransactor{contract: contract}, ConditionalEscrowFilterer: ConditionalEscrowFilterer{contract: contract}}, nil
}

// NewConditionalEscrowCaller creates a new read-only instance of ConditionalEscrow, bound to a specific deployed contract.
func NewConditionalEscrowCaller(address common.Address, caller bind.ContractCaller) (*ConditionalEscrowCaller, error) {
	contract, err := bindConditionalEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrowCaller{contract: contract}, nil
}

// NewConditionalEscrowTransactor creates a new write-only instance of ConditionalEscrow, bound to a specific deployed contract.
func NewConditionalEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ConditionalEscrowTransactor, error) {
	contract, err := bindConditionalEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrowTransactor{contract: contract}, nil
}

// NewConditionalEscrowFilterer creates a new log filterer instance of ConditionalEscrow, bound to a specific deployed contract.
func NewConditionalEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ConditionalEscrowFilterer, error) {
	contract, err := bindConditionalEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrowFilterer{contract: contract}, nil
}

// bindConditionalEscrow binds a generic wrapper to an already deployed contract.
func bindConditionalEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConditionalEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConditionalEscrow *ConditionalEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConditionalEscrow.Contract.ConditionalEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConditionalEscrow *ConditionalEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.ConditionalEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConditionalEscrow *ConditionalEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.ConditionalEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConditionalEscrow *ConditionalEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConditionalEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConditionalEscrow *ConditionalEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConditionalEscrow *ConditionalEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.contract.Transact(opts, method, params...)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_ConditionalEscrow *ConditionalEscrowCaller) DepositsOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConditionalEscrow.contract.Call(opts, &out, "depositsOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_ConditionalEscrow *ConditionalEscrowSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _ConditionalEscrow.Contract.DepositsOf(&_ConditionalEscrow.CallOpts, payee)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_ConditionalEscrow *ConditionalEscrowCallerSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _ConditionalEscrow.Contract.DepositsOf(&_ConditionalEscrow.CallOpts, payee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConditionalEscrow *ConditionalEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConditionalEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConditionalEscrow *ConditionalEscrowSession) Owner() (common.Address, error) {
	return _ConditionalEscrow.Contract.Owner(&_ConditionalEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConditionalEscrow *ConditionalEscrowCallerSession) Owner() (common.Address, error) {
	return _ConditionalEscrow.Contract.Owner(&_ConditionalEscrow.CallOpts)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address payee) view returns(bool)
func (_ConditionalEscrow *ConditionalEscrowCaller) WithdrawalAllowed(opts *bind.CallOpts, payee common.Address) (bool, error) {
	var out []interface{}
	err := _ConditionalEscrow.contract.Call(opts, &out, "withdrawalAllowed", payee)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address payee) view returns(bool)
func (_ConditionalEscrow *ConditionalEscrowSession) WithdrawalAllowed(payee common.Address) (bool, error) {
	return _ConditionalEscrow.Contract.WithdrawalAllowed(&_ConditionalEscrow.CallOpts, payee)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address payee) view returns(bool)
func (_ConditionalEscrow *ConditionalEscrowCallerSession) WithdrawalAllowed(payee common.Address) (bool, error) {
	return _ConditionalEscrow.Contract.WithdrawalAllowed(&_ConditionalEscrow.CallOpts, payee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address payee) payable returns()
func (_ConditionalEscrow *ConditionalEscrowTransactor) Deposit(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.contract.Transact(opts, "deposit", payee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address payee) payable returns()
func (_ConditionalEscrow *ConditionalEscrowSession) Deposit(payee common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.Deposit(&_ConditionalEscrow.TransactOpts, payee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address payee) payable returns()
func (_ConditionalEscrow *ConditionalEscrowTransactorSession) Deposit(payee common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.Deposit(&_ConditionalEscrow.TransactOpts, payee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConditionalEscrow *ConditionalEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionalEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConditionalEscrow *ConditionalEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.RenounceOwnership(&_ConditionalEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConditionalEscrow *ConditionalEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.RenounceOwnership(&_ConditionalEscrow.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConditionalEscrow *ConditionalEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConditionalEscrow *ConditionalEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.TransferOwnership(&_ConditionalEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConditionalEscrow *ConditionalEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.TransferOwnership(&_ConditionalEscrow.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_ConditionalEscrow *ConditionalEscrowTransactor) Withdraw(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.contract.Transact(opts, "withdraw", payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_ConditionalEscrow *ConditionalEscrowSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.Withdraw(&_ConditionalEscrow.TransactOpts, payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_ConditionalEscrow *ConditionalEscrowTransactorSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _ConditionalEscrow.Contract.Withdraw(&_ConditionalEscrow.TransactOpts, payee)
}

// ConditionalEscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ConditionalEscrow contract.
type ConditionalEscrowDepositedIterator struct {
	Event *ConditionalEscrowDeposited // Event containing the contract specifics and raw log

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
func (it *ConditionalEscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalEscrowDeposited)
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
		it.Event = new(ConditionalEscrowDeposited)
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
func (it *ConditionalEscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalEscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalEscrowDeposited represents a Deposited event raised by the ConditionalEscrow contract.
type ConditionalEscrowDeposited struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_ConditionalEscrow *ConditionalEscrowFilterer) FilterDeposited(opts *bind.FilterOpts, payee []common.Address) (*ConditionalEscrowDepositedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _ConditionalEscrow.contract.FilterLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrowDepositedIterator{contract: _ConditionalEscrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_ConditionalEscrow *ConditionalEscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ConditionalEscrowDeposited, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _ConditionalEscrow.contract.WatchLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalEscrowDeposited)
				if err := _ConditionalEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_ConditionalEscrow *ConditionalEscrowFilterer) ParseDeposited(log types.Log) (*ConditionalEscrowDeposited, error) {
	event := new(ConditionalEscrowDeposited)
	if err := _ConditionalEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConditionalEscrow contract.
type ConditionalEscrowOwnershipTransferredIterator struct {
	Event *ConditionalEscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConditionalEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalEscrowOwnershipTransferred)
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
		it.Event = new(ConditionalEscrowOwnershipTransferred)
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
func (it *ConditionalEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the ConditionalEscrow contract.
type ConditionalEscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConditionalEscrow *ConditionalEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConditionalEscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConditionalEscrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrowOwnershipTransferredIterator{contract: _ConditionalEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConditionalEscrow *ConditionalEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConditionalEscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConditionalEscrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalEscrowOwnershipTransferred)
				if err := _ConditionalEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConditionalEscrow *ConditionalEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*ConditionalEscrowOwnershipTransferred, error) {
	event := new(ConditionalEscrowOwnershipTransferred)
	if err := _ConditionalEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalEscrowWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ConditionalEscrow contract.
type ConditionalEscrowWithdrawnIterator struct {
	Event *ConditionalEscrowWithdrawn // Event containing the contract specifics and raw log

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
func (it *ConditionalEscrowWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalEscrowWithdrawn)
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
		it.Event = new(ConditionalEscrowWithdrawn)
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
func (it *ConditionalEscrowWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalEscrowWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalEscrowWithdrawn represents a Withdrawn event raised by the ConditionalEscrow contract.
type ConditionalEscrowWithdrawn struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_ConditionalEscrow *ConditionalEscrowFilterer) FilterWithdrawn(opts *bind.FilterOpts, payee []common.Address) (*ConditionalEscrowWithdrawnIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _ConditionalEscrow.contract.FilterLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalEscrowWithdrawnIterator{contract: _ConditionalEscrow.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_ConditionalEscrow *ConditionalEscrowFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ConditionalEscrowWithdrawn, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _ConditionalEscrow.contract.WatchLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalEscrowWithdrawn)
				if err := _ConditionalEscrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_ConditionalEscrow *ConditionalEscrowFilterer) ParseWithdrawn(log types.Log) (*ConditionalEscrowWithdrawn, error) {
	event := new(ConditionalEscrowWithdrawn)
	if err := _ConditionalEscrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// CrowdsaleMintABI is the input ABI used to generate the binding from.
const CrowdsaleMintABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rateReceived\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"walletReceived\",\"type\":\"address\"},{\"internalType\":\"contractLBCToken\",\"name\":\"tokenReceived\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CrowdsaleMintFuncSigs maps the 4-byte function signature to its string representation.
var CrowdsaleMintFuncSigs = map[string]string{
	"ec8ac4d8": "buyTokens(address)",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
}

// CrowdsaleMintBin is the compiled bytecode used for deploying new contracts.
var CrowdsaleMintBin = "0x608060405234801561001057600080fd5b506040516107163803806107168339818101604052606081101561003357600080fd5b508051602082015160409092015160016000559091908261009b576040805162461bcd60e51b815260206004820152601460248201527f43726f776473616c653a20726174652069732030000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166100e05760405162461bcd60e51b81526004018080602001828103825260258152602001806106f16025913960400191505060405180910390fd5b6001600160a01b0381166101255760405162461bcd60e51b81526004018080602001828103825260248152602001806106cd6024913960400191505060405180910390fd5b600392909255600280546001600160a01b039283166001600160a01b03199182161790915560018054929093169116179055610567806101666000396000f3fe60806040526004361061004e5760003560e01c80632c4e722e146100705780634042b66f14610097578063521eb273146100ac578063ec8ac4d8146100dd578063fc0c546a1461010357610065565b366100655761006361005e610118565b61011c565b005b61006361005e610118565b34801561007c57600080fd5b5061008561022c565b60408051918252519081900360200190f35b3480156100a357600080fd5b50610085610232565b3480156100b857600080fd5b506100c1610238565b604080516001600160a01b039092168252519081900360200190f35b610063600480360360208110156100f357600080fd5b50356001600160a01b031661011c565b34801561010f57600080fd5b506100c1610247565b3390565b60026000541415610174576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002600055346101848282610256565b600061018f826102f1565b60045490915061019f908361030e565b6004556101ac838261036f565b826001600160a01b03166101be610118565b6001600160a01b03167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b8484604051808381526020018281526020019250505060405180910390a361021083836102ed565b610218610379565b61022283836102ed565b5050600160005550565b60035490565b60045490565b6002546001600160a01b031690565b6001546001600160a01b031690565b6001600160a01b03821661029b5760405162461bcd60e51b815260040180806020018281038252602a815260200180610508602a913960400191505060405180910390fd5b806102ed576040805162461bcd60e51b815260206004820152601960248201527f43726f776473616c653a20776569416d6f756e74206973203000000000000000604482015290519081900360640190fd5b5050565b6000610308600354836103b590919063ffffffff16565b92915050565b600082820183811015610368576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6102ed828261040e565b6002546040516001600160a01b03909116903480156108fc02916000818181858888f193505050501580156103b2573d6000803e3d6000fd5b50565b6000826103c457506000610308565b828202828482816103d157fe5b04146103685760405162461bcd60e51b81526004018080602001828103825260218152602001806104e76021913960400191505060405180910390fd5b600154604080516340c10f1960e01b81526001600160a01b03858116600483015260248201859052915191909216916340c10f199160448083019260209291908290030181600087803b15801561046457600080fd5b505af1158015610478573d6000803e3d6000fd5b505050506040513d602081101561048e57600080fd5b505115156001146102ed576040805162461bcd60e51b815260206004820152601960248201527f43726f776473616c653a206d696e74696e67206661696c656400000000000000604482015290519081900360640190fdfe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7743726f776473616c653a2062656e656669636961727920697320746865207a65726f2061646472657373a2646970667358221220e588f7a4978bc22f8d71a57ab6694e9eb8bae3d94ae5185a75f2f79b9cfe003464736f6c6343000706003343726f776473616c653a20746f6b656e20697320746865207a65726f206164647265737343726f776473616c653a2077616c6c657420697320746865207a65726f2061646472657373"

// DeployCrowdsaleMint deploys a new Ethereum contract, binding an instance of CrowdsaleMint to it.
func DeployCrowdsaleMint(auth *bind.TransactOpts, backend bind.ContractBackend, rateReceived *big.Int, walletReceived common.Address, tokenReceived common.Address) (common.Address, *types.Transaction, *CrowdsaleMint, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleMintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrowdsaleMintBin), backend, rateReceived, walletReceived, tokenReceived)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrowdsaleMint{CrowdsaleMintCaller: CrowdsaleMintCaller{contract: contract}, CrowdsaleMintTransactor: CrowdsaleMintTransactor{contract: contract}, CrowdsaleMintFilterer: CrowdsaleMintFilterer{contract: contract}}, nil
}

// CrowdsaleMint is an auto generated Go binding around an Ethereum contract.
type CrowdsaleMint struct {
	CrowdsaleMintCaller     // Read-only binding to the contract
	CrowdsaleMintTransactor // Write-only binding to the contract
	CrowdsaleMintFilterer   // Log filterer for contract events
}

// CrowdsaleMintCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdsaleMintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleMintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdsaleMintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleMintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdsaleMintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleMintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdsaleMintSession struct {
	Contract     *CrowdsaleMint    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdsaleMintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdsaleMintCallerSession struct {
	Contract *CrowdsaleMintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CrowdsaleMintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdsaleMintTransactorSession struct {
	Contract     *CrowdsaleMintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CrowdsaleMintRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdsaleMintRaw struct {
	Contract *CrowdsaleMint // Generic contract binding to access the raw methods on
}

// CrowdsaleMintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdsaleMintCallerRaw struct {
	Contract *CrowdsaleMintCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdsaleMintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdsaleMintTransactorRaw struct {
	Contract *CrowdsaleMintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdsaleMint creates a new instance of CrowdsaleMint, bound to a specific deployed contract.
func NewCrowdsaleMint(address common.Address, backend bind.ContractBackend) (*CrowdsaleMint, error) {
	contract, err := bindCrowdsaleMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleMint{CrowdsaleMintCaller: CrowdsaleMintCaller{contract: contract}, CrowdsaleMintTransactor: CrowdsaleMintTransactor{contract: contract}, CrowdsaleMintFilterer: CrowdsaleMintFilterer{contract: contract}}, nil
}

// NewCrowdsaleMintCaller creates a new read-only instance of CrowdsaleMint, bound to a specific deployed contract.
func NewCrowdsaleMintCaller(address common.Address, caller bind.ContractCaller) (*CrowdsaleMintCaller, error) {
	contract, err := bindCrowdsaleMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleMintCaller{contract: contract}, nil
}

// NewCrowdsaleMintTransactor creates a new write-only instance of CrowdsaleMint, bound to a specific deployed contract.
func NewCrowdsaleMintTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdsaleMintTransactor, error) {
	contract, err := bindCrowdsaleMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleMintTransactor{contract: contract}, nil
}

// NewCrowdsaleMintFilterer creates a new log filterer instance of CrowdsaleMint, bound to a specific deployed contract.
func NewCrowdsaleMintFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdsaleMintFilterer, error) {
	contract, err := bindCrowdsaleMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleMintFilterer{contract: contract}, nil
}

// bindCrowdsaleMint binds a generic wrapper to an already deployed contract.
func bindCrowdsaleMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleMintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdsaleMint *CrowdsaleMintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdsaleMint.Contract.CrowdsaleMintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdsaleMint *CrowdsaleMintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.CrowdsaleMintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdsaleMint *CrowdsaleMintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.CrowdsaleMintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdsaleMint *CrowdsaleMintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdsaleMint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdsaleMint *CrowdsaleMintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdsaleMint *CrowdsaleMintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrowdsaleMint *CrowdsaleMintCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdsaleMint.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrowdsaleMint *CrowdsaleMintSession) Rate() (*big.Int, error) {
	return _CrowdsaleMint.Contract.Rate(&_CrowdsaleMint.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrowdsaleMint *CrowdsaleMintCallerSession) Rate() (*big.Int, error) {
	return _CrowdsaleMint.Contract.Rate(&_CrowdsaleMint.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdsaleMint *CrowdsaleMintCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdsaleMint.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdsaleMint *CrowdsaleMintSession) Token() (common.Address, error) {
	return _CrowdsaleMint.Contract.Token(&_CrowdsaleMint.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdsaleMint *CrowdsaleMintCallerSession) Token() (common.Address, error) {
	return _CrowdsaleMint.Contract.Token(&_CrowdsaleMint.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CrowdsaleMint *CrowdsaleMintCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdsaleMint.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CrowdsaleMint *CrowdsaleMintSession) Wallet() (common.Address, error) {
	return _CrowdsaleMint.Contract.Wallet(&_CrowdsaleMint.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CrowdsaleMint *CrowdsaleMintCallerSession) Wallet() (common.Address, error) {
	return _CrowdsaleMint.Contract.Wallet(&_CrowdsaleMint.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CrowdsaleMint *CrowdsaleMintCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdsaleMint.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CrowdsaleMint *CrowdsaleMintSession) WeiRaised() (*big.Int, error) {
	return _CrowdsaleMint.Contract.WeiRaised(&_CrowdsaleMint.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CrowdsaleMint *CrowdsaleMintCallerSession) WeiRaised() (*big.Int, error) {
	return _CrowdsaleMint.Contract.WeiRaised(&_CrowdsaleMint.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_CrowdsaleMint *CrowdsaleMintTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _CrowdsaleMint.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_CrowdsaleMint *CrowdsaleMintSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.BuyTokens(&_CrowdsaleMint.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_CrowdsaleMint *CrowdsaleMintTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.BuyTokens(&_CrowdsaleMint.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrowdsaleMint *CrowdsaleMintTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CrowdsaleMint.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrowdsaleMint *CrowdsaleMintSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.Fallback(&_CrowdsaleMint.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrowdsaleMint *CrowdsaleMintTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.Fallback(&_CrowdsaleMint.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdsaleMint *CrowdsaleMintTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdsaleMint.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdsaleMint *CrowdsaleMintSession) Receive() (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.Receive(&_CrowdsaleMint.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdsaleMint *CrowdsaleMintTransactorSession) Receive() (*types.Transaction, error) {
	return _CrowdsaleMint.Contract.Receive(&_CrowdsaleMint.TransactOpts)
}

// CrowdsaleMintTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the CrowdsaleMint contract.
type CrowdsaleMintTokensPurchasedIterator struct {
	Event *CrowdsaleMintTokensPurchased // Event containing the contract specifics and raw log

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
func (it *CrowdsaleMintTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleMintTokensPurchased)
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
		it.Event = new(CrowdsaleMintTokensPurchased)
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
func (it *CrowdsaleMintTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleMintTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleMintTokensPurchased represents a TokensPurchased event raised by the CrowdsaleMint contract.
type CrowdsaleMintTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CrowdsaleMint *CrowdsaleMintFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*CrowdsaleMintTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CrowdsaleMint.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleMintTokensPurchasedIterator{contract: _CrowdsaleMint.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CrowdsaleMint *CrowdsaleMintFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *CrowdsaleMintTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CrowdsaleMint.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleMintTokensPurchased)
				if err := _CrowdsaleMint.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CrowdsaleMint *CrowdsaleMintFilterer) ParseTokensPurchased(log types.Log) (*CrowdsaleMintTokensPurchased, error) {
	event := new(CrowdsaleMintTokensPurchased)
	if err := _CrowdsaleMint.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CappedUnburnableABI is the input ABI used to generate the binding from.
const ERC20CappedUnburnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20CappedUnburnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20CappedUnburnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"355274ea": "cap()",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"5c975abb": "paused()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20CappedUnburnable is an auto generated Go binding around an Ethereum contract.
type ERC20CappedUnburnable struct {
	ERC20CappedUnburnableCaller     // Read-only binding to the contract
	ERC20CappedUnburnableTransactor // Write-only binding to the contract
	ERC20CappedUnburnableFilterer   // Log filterer for contract events
}

// ERC20CappedUnburnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CappedUnburnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CappedUnburnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CappedUnburnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CappedUnburnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CappedUnburnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CappedUnburnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CappedUnburnableSession struct {
	Contract     *ERC20CappedUnburnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20CappedUnburnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CappedUnburnableCallerSession struct {
	Contract *ERC20CappedUnburnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ERC20CappedUnburnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CappedUnburnableTransactorSession struct {
	Contract     *ERC20CappedUnburnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ERC20CappedUnburnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CappedUnburnableRaw struct {
	Contract *ERC20CappedUnburnable // Generic contract binding to access the raw methods on
}

// ERC20CappedUnburnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CappedUnburnableCallerRaw struct {
	Contract *ERC20CappedUnburnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CappedUnburnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CappedUnburnableTransactorRaw struct {
	Contract *ERC20CappedUnburnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20CappedUnburnable creates a new instance of ERC20CappedUnburnable, bound to a specific deployed contract.
func NewERC20CappedUnburnable(address common.Address, backend bind.ContractBackend) (*ERC20CappedUnburnable, error) {
	contract, err := bindERC20CappedUnburnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnable{ERC20CappedUnburnableCaller: ERC20CappedUnburnableCaller{contract: contract}, ERC20CappedUnburnableTransactor: ERC20CappedUnburnableTransactor{contract: contract}, ERC20CappedUnburnableFilterer: ERC20CappedUnburnableFilterer{contract: contract}}, nil
}

// NewERC20CappedUnburnableCaller creates a new read-only instance of ERC20CappedUnburnable, bound to a specific deployed contract.
func NewERC20CappedUnburnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20CappedUnburnableCaller, error) {
	contract, err := bindERC20CappedUnburnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnableCaller{contract: contract}, nil
}

// NewERC20CappedUnburnableTransactor creates a new write-only instance of ERC20CappedUnburnable, bound to a specific deployed contract.
func NewERC20CappedUnburnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CappedUnburnableTransactor, error) {
	contract, err := bindERC20CappedUnburnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnableTransactor{contract: contract}, nil
}

// NewERC20CappedUnburnableFilterer creates a new log filterer instance of ERC20CappedUnburnable, bound to a specific deployed contract.
func NewERC20CappedUnburnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CappedUnburnableFilterer, error) {
	contract, err := bindERC20CappedUnburnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnableFilterer{contract: contract}, nil
}

// bindERC20CappedUnburnable binds a generic wrapper to an already deployed contract.
func bindERC20CappedUnburnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20CappedUnburnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CappedUnburnable *ERC20CappedUnburnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CappedUnburnable.Contract.ERC20CappedUnburnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CappedUnburnable *ERC20CappedUnburnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.ERC20CappedUnburnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CappedUnburnable *ERC20CappedUnburnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.ERC20CappedUnburnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CappedUnburnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.Allowance(&_ERC20CappedUnburnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.Allowance(&_ERC20CappedUnburnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.BalanceOf(&_ERC20CappedUnburnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.BalanceOf(&_ERC20CappedUnburnable.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Cap() (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.Cap(&_ERC20CappedUnburnable.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) Cap() (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.Cap(&_ERC20CappedUnburnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Decimals() (uint8, error) {
	return _ERC20CappedUnburnable.Contract.Decimals(&_ERC20CappedUnburnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) Decimals() (uint8, error) {
	return _ERC20CappedUnburnable.Contract.Decimals(&_ERC20CappedUnburnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Name() (string, error) {
	return _ERC20CappedUnburnable.Contract.Name(&_ERC20CappedUnburnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) Name() (string, error) {
	return _ERC20CappedUnburnable.Contract.Name(&_ERC20CappedUnburnable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Paused() (bool, error) {
	return _ERC20CappedUnburnable.Contract.Paused(&_ERC20CappedUnburnable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) Paused() (bool, error) {
	return _ERC20CappedUnburnable.Contract.Paused(&_ERC20CappedUnburnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Symbol() (string, error) {
	return _ERC20CappedUnburnable.Contract.Symbol(&_ERC20CappedUnburnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) Symbol() (string, error) {
	return _ERC20CappedUnburnable.Contract.Symbol(&_ERC20CappedUnburnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20CappedUnburnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.TotalSupply(&_ERC20CappedUnburnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20CappedUnburnable.Contract.TotalSupply(&_ERC20CappedUnburnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.Approve(&_ERC20CappedUnburnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.Approve(&_ERC20CappedUnburnable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.DecreaseAllowance(&_ERC20CappedUnburnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.DecreaseAllowance(&_ERC20CappedUnburnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.IncreaseAllowance(&_ERC20CappedUnburnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.IncreaseAllowance(&_ERC20CappedUnburnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.Transfer(&_ERC20CappedUnburnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.Transfer(&_ERC20CappedUnburnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.TransferFrom(&_ERC20CappedUnburnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CappedUnburnable.Contract.TransferFrom(&_ERC20CappedUnburnable.TransactOpts, sender, recipient, amount)
}

// ERC20CappedUnburnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnableApprovalIterator struct {
	Event *ERC20CappedUnburnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20CappedUnburnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CappedUnburnableApproval)
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
		it.Event = new(ERC20CappedUnburnableApproval)
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
func (it *ERC20CappedUnburnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CappedUnburnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CappedUnburnableApproval represents a Approval event raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20CappedUnburnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20CappedUnburnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnableApprovalIterator{contract: _ERC20CappedUnburnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20CappedUnburnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20CappedUnburnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CappedUnburnableApproval)
				if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) ParseApproval(log types.Log) (*ERC20CappedUnburnableApproval, error) {
	event := new(ERC20CappedUnburnableApproval)
	if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CappedUnburnablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnablePausedIterator struct {
	Event *ERC20CappedUnburnablePaused // Event containing the contract specifics and raw log

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
func (it *ERC20CappedUnburnablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CappedUnburnablePaused)
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
		it.Event = new(ERC20CappedUnburnablePaused)
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
func (it *ERC20CappedUnburnablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CappedUnburnablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CappedUnburnablePaused represents a Paused event raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20CappedUnburnablePausedIterator, error) {

	logs, sub, err := _ERC20CappedUnburnable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnablePausedIterator{contract: _ERC20CappedUnburnable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20CappedUnburnablePaused) (event.Subscription, error) {

	logs, sub, err := _ERC20CappedUnburnable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CappedUnburnablePaused)
				if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) ParsePaused(log types.Log) (*ERC20CappedUnburnablePaused, error) {
	event := new(ERC20CappedUnburnablePaused)
	if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CappedUnburnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnableTransferIterator struct {
	Event *ERC20CappedUnburnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20CappedUnburnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CappedUnburnableTransfer)
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
		it.Event = new(ERC20CappedUnburnableTransfer)
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
func (it *ERC20CappedUnburnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CappedUnburnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CappedUnburnableTransfer represents a Transfer event raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20CappedUnburnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CappedUnburnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnableTransferIterator{contract: _ERC20CappedUnburnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20CappedUnburnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CappedUnburnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CappedUnburnableTransfer)
				if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) ParseTransfer(log types.Log) (*ERC20CappedUnburnableTransfer, error) {
	event := new(ERC20CappedUnburnableTransfer)
	if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CappedUnburnableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnableUnpausedIterator struct {
	Event *ERC20CappedUnburnableUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20CappedUnburnableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CappedUnburnableUnpaused)
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
		it.Event = new(ERC20CappedUnburnableUnpaused)
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
func (it *ERC20CappedUnburnableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CappedUnburnableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CappedUnburnableUnpaused represents a Unpaused event raised by the ERC20CappedUnburnable contract.
type ERC20CappedUnburnableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20CappedUnburnableUnpausedIterator, error) {

	logs, sub, err := _ERC20CappedUnburnable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20CappedUnburnableUnpausedIterator{contract: _ERC20CappedUnburnable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20CappedUnburnableUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20CappedUnburnable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CappedUnburnableUnpaused)
				if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20CappedUnburnable *ERC20CappedUnburnableFilterer) ParseUnpaused(log types.Log) (*ERC20CappedUnburnableUnpaused, error) {
	event := new(ERC20CappedUnburnableUnpaused)
	if err := _ERC20CappedUnburnable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUnburnableABI is the input ABI used to generate the binding from.
const ERC20PausableUnburnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20PausableUnburnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20PausableUnburnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"5c975abb": "paused()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20PausableUnburnable is an auto generated Go binding around an Ethereum contract.
type ERC20PausableUnburnable struct {
	ERC20PausableUnburnableCaller     // Read-only binding to the contract
	ERC20PausableUnburnableTransactor // Write-only binding to the contract
	ERC20PausableUnburnableFilterer   // Log filterer for contract events
}

// ERC20PausableUnburnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PausableUnburnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PausableUnburnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PausableUnburnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PausableUnburnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PausableUnburnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PausableUnburnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PausableUnburnableSession struct {
	Contract     *ERC20PausableUnburnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20PausableUnburnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PausableUnburnableCallerSession struct {
	Contract *ERC20PausableUnburnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC20PausableUnburnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PausableUnburnableTransactorSession struct {
	Contract     *ERC20PausableUnburnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC20PausableUnburnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PausableUnburnableRaw struct {
	Contract *ERC20PausableUnburnable // Generic contract binding to access the raw methods on
}

// ERC20PausableUnburnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PausableUnburnableCallerRaw struct {
	Contract *ERC20PausableUnburnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PausableUnburnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PausableUnburnableTransactorRaw struct {
	Contract *ERC20PausableUnburnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20PausableUnburnable creates a new instance of ERC20PausableUnburnable, bound to a specific deployed contract.
func NewERC20PausableUnburnable(address common.Address, backend bind.ContractBackend) (*ERC20PausableUnburnable, error) {
	contract, err := bindERC20PausableUnburnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnable{ERC20PausableUnburnableCaller: ERC20PausableUnburnableCaller{contract: contract}, ERC20PausableUnburnableTransactor: ERC20PausableUnburnableTransactor{contract: contract}, ERC20PausableUnburnableFilterer: ERC20PausableUnburnableFilterer{contract: contract}}, nil
}

// NewERC20PausableUnburnableCaller creates a new read-only instance of ERC20PausableUnburnable, bound to a specific deployed contract.
func NewERC20PausableUnburnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20PausableUnburnableCaller, error) {
	contract, err := bindERC20PausableUnburnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnableCaller{contract: contract}, nil
}

// NewERC20PausableUnburnableTransactor creates a new write-only instance of ERC20PausableUnburnable, bound to a specific deployed contract.
func NewERC20PausableUnburnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PausableUnburnableTransactor, error) {
	contract, err := bindERC20PausableUnburnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnableTransactor{contract: contract}, nil
}

// NewERC20PausableUnburnableFilterer creates a new log filterer instance of ERC20PausableUnburnable, bound to a specific deployed contract.
func NewERC20PausableUnburnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PausableUnburnableFilterer, error) {
	contract, err := bindERC20PausableUnburnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnableFilterer{contract: contract}, nil
}

// bindERC20PausableUnburnable binds a generic wrapper to an already deployed contract.
func bindERC20PausableUnburnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PausableUnburnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PausableUnburnable *ERC20PausableUnburnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PausableUnburnable.Contract.ERC20PausableUnburnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PausableUnburnable *ERC20PausableUnburnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.ERC20PausableUnburnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PausableUnburnable *ERC20PausableUnburnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.ERC20PausableUnburnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PausableUnburnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PausableUnburnable.Contract.Allowance(&_ERC20PausableUnburnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PausableUnburnable.Contract.Allowance(&_ERC20PausableUnburnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PausableUnburnable.Contract.BalanceOf(&_ERC20PausableUnburnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PausableUnburnable.Contract.BalanceOf(&_ERC20PausableUnburnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Decimals() (uint8, error) {
	return _ERC20PausableUnburnable.Contract.Decimals(&_ERC20PausableUnburnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) Decimals() (uint8, error) {
	return _ERC20PausableUnburnable.Contract.Decimals(&_ERC20PausableUnburnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Name() (string, error) {
	return _ERC20PausableUnburnable.Contract.Name(&_ERC20PausableUnburnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) Name() (string, error) {
	return _ERC20PausableUnburnable.Contract.Name(&_ERC20PausableUnburnable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Paused() (bool, error) {
	return _ERC20PausableUnburnable.Contract.Paused(&_ERC20PausableUnburnable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) Paused() (bool, error) {
	return _ERC20PausableUnburnable.Contract.Paused(&_ERC20PausableUnburnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Symbol() (string, error) {
	return _ERC20PausableUnburnable.Contract.Symbol(&_ERC20PausableUnburnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) Symbol() (string, error) {
	return _ERC20PausableUnburnable.Contract.Symbol(&_ERC20PausableUnburnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PausableUnburnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20PausableUnburnable.Contract.TotalSupply(&_ERC20PausableUnburnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20PausableUnburnable.Contract.TotalSupply(&_ERC20PausableUnburnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.Approve(&_ERC20PausableUnburnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.Approve(&_ERC20PausableUnburnable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.DecreaseAllowance(&_ERC20PausableUnburnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.DecreaseAllowance(&_ERC20PausableUnburnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.IncreaseAllowance(&_ERC20PausableUnburnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.IncreaseAllowance(&_ERC20PausableUnburnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.Transfer(&_ERC20PausableUnburnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.Transfer(&_ERC20PausableUnburnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.TransferFrom(&_ERC20PausableUnburnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUnburnable.Contract.TransferFrom(&_ERC20PausableUnburnable.TransactOpts, sender, recipient, amount)
}

// ERC20PausableUnburnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnableApprovalIterator struct {
	Event *ERC20PausableUnburnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUnburnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUnburnableApproval)
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
		it.Event = new(ERC20PausableUnburnableApproval)
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
func (it *ERC20PausableUnburnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUnburnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUnburnableApproval represents a Approval event raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PausableUnburnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PausableUnburnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnableApprovalIterator{contract: _ERC20PausableUnburnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PausableUnburnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PausableUnburnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUnburnableApproval)
				if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) ParseApproval(log types.Log) (*ERC20PausableUnburnableApproval, error) {
	event := new(ERC20PausableUnburnableApproval)
	if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUnburnablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnablePausedIterator struct {
	Event *ERC20PausableUnburnablePaused // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUnburnablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUnburnablePaused)
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
		it.Event = new(ERC20PausableUnburnablePaused)
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
func (it *ERC20PausableUnburnablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUnburnablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUnburnablePaused represents a Paused event raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20PausableUnburnablePausedIterator, error) {

	logs, sub, err := _ERC20PausableUnburnable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnablePausedIterator{contract: _ERC20PausableUnburnable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20PausableUnburnablePaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PausableUnburnable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUnburnablePaused)
				if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) ParsePaused(log types.Log) (*ERC20PausableUnburnablePaused, error) {
	event := new(ERC20PausableUnburnablePaused)
	if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUnburnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnableTransferIterator struct {
	Event *ERC20PausableUnburnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUnburnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUnburnableTransfer)
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
		it.Event = new(ERC20PausableUnburnableTransfer)
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
func (it *ERC20PausableUnburnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUnburnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUnburnableTransfer represents a Transfer event raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PausableUnburnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PausableUnburnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnableTransferIterator{contract: _ERC20PausableUnburnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PausableUnburnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PausableUnburnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUnburnableTransfer)
				if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) ParseTransfer(log types.Log) (*ERC20PausableUnburnableTransfer, error) {
	event := new(ERC20PausableUnburnableTransfer)
	if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUnburnableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnableUnpausedIterator struct {
	Event *ERC20PausableUnburnableUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUnburnableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUnburnableUnpaused)
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
		it.Event = new(ERC20PausableUnburnableUnpaused)
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
func (it *ERC20PausableUnburnableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUnburnableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUnburnableUnpaused represents a Unpaused event raised by the ERC20PausableUnburnable contract.
type ERC20PausableUnburnableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20PausableUnburnableUnpausedIterator, error) {

	logs, sub, err := _ERC20PausableUnburnable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUnburnableUnpausedIterator{contract: _ERC20PausableUnburnable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20PausableUnburnableUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PausableUnburnable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUnburnableUnpaused)
				if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PausableUnburnable *ERC20PausableUnburnableFilterer) ParseUnpaused(log types.Log) (*ERC20PausableUnburnableUnpaused, error) {
	event := new(ERC20PausableUnburnableUnpaused)
	if err := _ERC20PausableUnburnable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UnburnableABI is the input ABI used to generate the binding from.
const ERC20UnburnableABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameGiven\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbolGiven\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20UnburnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20UnburnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20UnburnableBin is the compiled bytecode used for deploying new contracts.
var ERC20UnburnableBin = "0x60806040523480156200001157600080fd5b5060405162000cd438038062000cd4833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250508251620001b491506003906020850190620001e0565b508051620001ca906004906020840190620001e0565b50506005805460ff19166012179055506200028c565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000218576000855562000263565b82601f106200023357805160ff191683800117855562000263565b8280016001018555821562000263579182015b828111156200026357825182559160200191906001019062000246565b506200027192915062000275565b5090565b5b8082111562000271576000815560010162000276565b610a38806200029c6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063395093511161007157806339509351146101d957806370a082311461020557806395d89b411461022b578063a457c2d714610233578063a9059cbb1461025f578063dd62ed3e1461028b576100a9565b806306fdde03146100ae578063095ea7b31461012b57806318160ddd1461016b57806323b872dd14610185578063313ce567146101bb575b600080fd5b6100b66102b9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f05781810151838201526020016100d8565b50505050905090810190601f16801561011d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101576004803603604081101561014157600080fd5b506001600160a01b03813516906020013561034f565b604080519115158252519081900360200190f35b61017361036c565b60408051918252519081900360200190f35b6101576004803603606081101561019b57600080fd5b506001600160a01b03813581169160208101359091169060400135610372565b6101c36103f9565b6040805160ff9092168252519081900360200190f35b610157600480360360408110156101ef57600080fd5b506001600160a01b038135169060200135610402565b6101736004803603602081101561021b57600080fd5b50356001600160a01b0316610450565b6100b661046b565b6101576004803603604081101561024957600080fd5b506001600160a01b0381351690602001356104cc565b6101576004803603604081101561027557600080fd5b506001600160a01b038135169060200135610534565b610173600480360360408110156102a157600080fd5b506001600160a01b0381358116916020013516610548565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b820191906000526020600020905b81548152906001019060200180831161032857829003601f168201915b5050505050905090565b600061036361035c610573565b8484610577565b50600192915050565b60025490565b600061037f848484610663565b6103ef8461038b610573565b6103ea856040518060600160405280603281526020016109d1603291396001600160a01b038a166000908152600160205260408120906103c9610573565b6001600160a01b0316815260208101919091526040016000205491906107be565b610577565b5060019392505050565b60055460ff1690565b600061036361040f610573565b846103ea8560016000610420610573565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610855565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b60006103636104d9610573565b846103ea856040518060600160405280602f81526020016109a2602f913960016000610503610573565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906107be565b6000610363610541610573565b8484610663565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166105bc5760405162461bcd60e51b815260040180806020018281038252602e8152602001806108e8602e913960400191505060405180910390fd5b6001600160a01b0382166106015760405162461bcd60e51b815260040180806020018281038252602c8152602001806108bc602c913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106a85760405162461bcd60e51b815260040180806020018281038252602f815260200180610916602f913960400191505060405180910390fd5b6001600160a01b0382166106ed5760405162461bcd60e51b815260040180806020018281038252602d815260200180610945602d913960400191505060405180910390fd5b6106f88383836108b6565b61073581604051806060016040528060308152602001610972603091396001600160a01b03861660009081526020819052604090205491906107be565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546107649082610855565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561084d5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108125781810151838201526020016107fa565b50505050905090810190601f16801561083f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156108af576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe4552433230556e6275726e61626c653a20617070726f766520746f20746865207a65726f20616464726573734552433230556e6275726e61626c653a20617070726f76652066726f6d20746865207a65726f20616464726573734552433230556e6275726e61626c653a207472616e736665722066726f6d20746865207a65726f20616464726573734552433230556e6275726e61626c653a207472616e7366657220746f20746865207a65726f20616464726573734552433230556e6275726e61626c653a207472616e7366657220616d6f756e7420657863656564732062616c616e63654552433230556e6275726e61626c653a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f4552433230556e6275726e61626c653a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365a2646970667358221220695918832d93ec89b70932ed71f72ca8f575c85b9a9d1407c1f1087c5b1e9d8464736f6c63430007060033"

// DeployERC20Unburnable deploys a new Ethereum contract, binding an instance of ERC20Unburnable to it.
func DeployERC20Unburnable(auth *bind.TransactOpts, backend bind.ContractBackend, nameGiven string, symbolGiven string) (common.Address, *types.Transaction, *ERC20Unburnable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UnburnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20UnburnableBin), backend, nameGiven, symbolGiven)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Unburnable{ERC20UnburnableCaller: ERC20UnburnableCaller{contract: contract}, ERC20UnburnableTransactor: ERC20UnburnableTransactor{contract: contract}, ERC20UnburnableFilterer: ERC20UnburnableFilterer{contract: contract}}, nil
}

// ERC20Unburnable is an auto generated Go binding around an Ethereum contract.
type ERC20Unburnable struct {
	ERC20UnburnableCaller     // Read-only binding to the contract
	ERC20UnburnableTransactor // Write-only binding to the contract
	ERC20UnburnableFilterer   // Log filterer for contract events
}

// ERC20UnburnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20UnburnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UnburnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20UnburnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UnburnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20UnburnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UnburnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20UnburnableSession struct {
	Contract     *ERC20Unburnable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20UnburnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20UnburnableCallerSession struct {
	Contract *ERC20UnburnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC20UnburnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20UnburnableTransactorSession struct {
	Contract     *ERC20UnburnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC20UnburnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20UnburnableRaw struct {
	Contract *ERC20Unburnable // Generic contract binding to access the raw methods on
}

// ERC20UnburnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20UnburnableCallerRaw struct {
	Contract *ERC20UnburnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20UnburnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20UnburnableTransactorRaw struct {
	Contract *ERC20UnburnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Unburnable creates a new instance of ERC20Unburnable, bound to a specific deployed contract.
func NewERC20Unburnable(address common.Address, backend bind.ContractBackend) (*ERC20Unburnable, error) {
	contract, err := bindERC20Unburnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Unburnable{ERC20UnburnableCaller: ERC20UnburnableCaller{contract: contract}, ERC20UnburnableTransactor: ERC20UnburnableTransactor{contract: contract}, ERC20UnburnableFilterer: ERC20UnburnableFilterer{contract: contract}}, nil
}

// NewERC20UnburnableCaller creates a new read-only instance of ERC20Unburnable, bound to a specific deployed contract.
func NewERC20UnburnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20UnburnableCaller, error) {
	contract, err := bindERC20Unburnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UnburnableCaller{contract: contract}, nil
}

// NewERC20UnburnableTransactor creates a new write-only instance of ERC20Unburnable, bound to a specific deployed contract.
func NewERC20UnburnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20UnburnableTransactor, error) {
	contract, err := bindERC20Unburnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UnburnableTransactor{contract: contract}, nil
}

// NewERC20UnburnableFilterer creates a new log filterer instance of ERC20Unburnable, bound to a specific deployed contract.
func NewERC20UnburnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20UnburnableFilterer, error) {
	contract, err := bindERC20Unburnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20UnburnableFilterer{contract: contract}, nil
}

// bindERC20Unburnable binds a generic wrapper to an already deployed contract.
func bindERC20Unburnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UnburnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Unburnable *ERC20UnburnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Unburnable.Contract.ERC20UnburnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Unburnable *ERC20UnburnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.ERC20UnburnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Unburnable *ERC20UnburnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.ERC20UnburnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Unburnable *ERC20UnburnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Unburnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Unburnable *ERC20UnburnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Unburnable *ERC20UnburnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Unburnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Unburnable.Contract.Allowance(&_ERC20Unburnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Unburnable.Contract.Allowance(&_ERC20Unburnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Unburnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Unburnable.Contract.BalanceOf(&_ERC20Unburnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Unburnable.Contract.BalanceOf(&_ERC20Unburnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Unburnable *ERC20UnburnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Unburnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Unburnable *ERC20UnburnableSession) Decimals() (uint8, error) {
	return _ERC20Unburnable.Contract.Decimals(&_ERC20Unburnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Unburnable *ERC20UnburnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Unburnable.Contract.Decimals(&_ERC20Unburnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Unburnable *ERC20UnburnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Unburnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Unburnable *ERC20UnburnableSession) Name() (string, error) {
	return _ERC20Unburnable.Contract.Name(&_ERC20Unburnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Unburnable *ERC20UnburnableCallerSession) Name() (string, error) {
	return _ERC20Unburnable.Contract.Name(&_ERC20Unburnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Unburnable *ERC20UnburnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Unburnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Unburnable *ERC20UnburnableSession) Symbol() (string, error) {
	return _ERC20Unburnable.Contract.Symbol(&_ERC20Unburnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Unburnable *ERC20UnburnableCallerSession) Symbol() (string, error) {
	return _ERC20Unburnable.Contract.Symbol(&_ERC20Unburnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Unburnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Unburnable.Contract.TotalSupply(&_ERC20Unburnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Unburnable *ERC20UnburnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Unburnable.Contract.TotalSupply(&_ERC20Unburnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.Approve(&_ERC20Unburnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.Approve(&_ERC20Unburnable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.DecreaseAllowance(&_ERC20Unburnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.DecreaseAllowance(&_ERC20Unburnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.IncreaseAllowance(&_ERC20Unburnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.IncreaseAllowance(&_ERC20Unburnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.Transfer(&_ERC20Unburnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.Transfer(&_ERC20Unburnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.TransferFrom(&_ERC20Unburnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Unburnable *ERC20UnburnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Unburnable.Contract.TransferFrom(&_ERC20Unburnable.TransactOpts, sender, recipient, amount)
}

// ERC20UnburnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Unburnable contract.
type ERC20UnburnableApprovalIterator struct {
	Event *ERC20UnburnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20UnburnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UnburnableApproval)
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
		it.Event = new(ERC20UnburnableApproval)
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
func (it *ERC20UnburnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UnburnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UnburnableApproval represents a Approval event raised by the ERC20Unburnable contract.
type ERC20UnburnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Unburnable *ERC20UnburnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20UnburnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Unburnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UnburnableApprovalIterator{contract: _ERC20Unburnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Unburnable *ERC20UnburnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20UnburnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Unburnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UnburnableApproval)
				if err := _ERC20Unburnable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Unburnable *ERC20UnburnableFilterer) ParseApproval(log types.Log) (*ERC20UnburnableApproval, error) {
	event := new(ERC20UnburnableApproval)
	if err := _ERC20Unburnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UnburnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Unburnable contract.
type ERC20UnburnableTransferIterator struct {
	Event *ERC20UnburnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20UnburnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UnburnableTransfer)
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
		it.Event = new(ERC20UnburnableTransfer)
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
func (it *ERC20UnburnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UnburnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UnburnableTransfer represents a Transfer event raised by the ERC20Unburnable contract.
type ERC20UnburnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Unburnable *ERC20UnburnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20UnburnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Unburnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UnburnableTransferIterator{contract: _ERC20Unburnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Unburnable *ERC20UnburnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20UnburnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Unburnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UnburnableTransfer)
				if err := _ERC20Unburnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Unburnable *ERC20UnburnableFilterer) ParseTransfer(log types.Log) (*ERC20UnburnableTransfer, error) {
	event := new(ERC20UnburnableTransfer)
	if err := _ERC20Unburnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowABI is the input ABI used to generate the binding from.
const EscrowABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"depositsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EscrowFuncSigs maps the 4-byte function signature to its string representation.
var EscrowFuncSigs = map[string]string{
	"f340fa01": "deposit(address)",
	"e3a9db1a": "depositsOf(address)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"51cff8d9": "withdraw(address)",
}

// EscrowBin is the compiled bytecode used for deploying new contracts.
var EscrowBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6106d28061007d6000396000f3fe6080604052600436106100555760003560e01c806351cff8d91461005a578063715018a61461008f5780638da5cb5b146100a4578063e3a9db1a146100d5578063f2fde38b1461011a578063f340fa011461014d575b600080fd5b34801561006657600080fd5b5061008d6004803603602081101561007d57600080fd5b50356001600160a01b0316610173565b005b34801561009b57600080fd5b5061008d610236565b3480156100b057600080fd5b506100b96102d8565b604080516001600160a01b039092168252519081900360200190f35b3480156100e157600080fd5b50610108600480360360208110156100f857600080fd5b50356001600160a01b03166102e7565b60408051918252519081900360200190f35b34801561012657600080fd5b5061008d6004803603602081101561013d57600080fd5b50356001600160a01b0316610302565b61008d6004803603602081101561016357600080fd5b50356001600160a01b03166103fa565b61017b6104cd565b6000546001600160a01b039081169116146101cb576040805162461bcd60e51b8152602060048201819052602482015260008051602061067d833981519152604482015290519081900360640190fd5b6001600160a01b03811660008181526001602052604081208054919055906101f390826104d1565b6040805182815290516001600160a01b038416917f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5919081900360200190a25050565b61023e6104cd565b6000546001600160a01b0390811691161461028e576040805162461bcd60e51b8152602060048201819052602482015260008051602061067d833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b6001600160a01b031660009081526001602052604090205490565b61030a6104cd565b6000546001600160a01b0390811691161461035a576040805162461bcd60e51b8152602060048201819052602482015260008051602061067d833981519152604482015290519081900360640190fd5b6001600160a01b03811661039f5760405162461bcd60e51b815260040180806020018281038252602681526020018061061d6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6104026104cd565b6000546001600160a01b03908116911614610452576040805162461bcd60e51b8152602060048201819052602482015260008051602061067d833981519152604482015290519081900360640190fd5b6001600160a01b038116600090815260016020526040902054349061047790826105bb565b6001600160a01b038316600081815260016020908152604091829020939093558051848152905191927f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c492918290030190a25050565b3390565b80471015610526576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d8060008114610571576040519150601f19603f3d011682016040523d82523d6000602084013e610576565b606091505b50509050806105b65760405162461bcd60e51b815260040180806020018281038252603a815260200180610643603a913960400191505060405180910390fd5b505050565b600082820183811015610615576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d617920686176652072657665727465644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212201b7f747e6e5ad40da3869236dec2c4362000323a9e19f56f8dc117cb910a38c964736f6c63430007060033"

// DeployEscrow deploys a new Ethereum contract, binding an instance of Escrow to it.
func DeployEscrow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Escrow, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EscrowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_Escrow *EscrowCaller) DepositsOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "depositsOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_Escrow *EscrowSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _Escrow.Contract.DepositsOf(&_Escrow.CallOpts, payee)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_Escrow *EscrowCallerSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _Escrow.Contract.DepositsOf(&_Escrow.CallOpts, payee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Escrow *EscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Escrow *EscrowSession) Owner() (common.Address, error) {
	return _Escrow.Contract.Owner(&_Escrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Escrow *EscrowCallerSession) Owner() (common.Address, error) {
	return _Escrow.Contract.Owner(&_Escrow.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address payee) payable returns()
func (_Escrow *EscrowTransactor) Deposit(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "deposit", payee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address payee) payable returns()
func (_Escrow *EscrowSession) Deposit(payee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts, payee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address payee) payable returns()
func (_Escrow *EscrowTransactorSession) Deposit(payee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts, payee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Escrow *EscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Escrow *EscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _Escrow.Contract.RenounceOwnership(&_Escrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Escrow *EscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Escrow.Contract.RenounceOwnership(&_Escrow.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Escrow *EscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Escrow *EscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.TransferOwnership(&_Escrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Escrow *EscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.TransferOwnership(&_Escrow.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_Escrow *EscrowTransactor) Withdraw(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "withdraw", payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_Escrow *EscrowSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts, payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_Escrow *EscrowTransactorSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts, payee)
}

// EscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Escrow contract.
type EscrowDepositedIterator struct {
	Event *EscrowDeposited // Event containing the contract specifics and raw log

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
func (it *EscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowDeposited)
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
		it.Event = new(EscrowDeposited)
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
func (it *EscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowDeposited represents a Deposited event raised by the Escrow contract.
type EscrowDeposited struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_Escrow *EscrowFilterer) FilterDeposited(opts *bind.FilterOpts, payee []common.Address) (*EscrowDepositedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return &EscrowDepositedIterator{contract: _Escrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_Escrow *EscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EscrowDeposited, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowDeposited)
				if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_Escrow *EscrowFilterer) ParseDeposited(log types.Log) (*EscrowDeposited, error) {
	event := new(EscrowDeposited)
	if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Escrow contract.
type EscrowOwnershipTransferredIterator struct {
	Event *EscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOwnershipTransferred)
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
		it.Event = new(EscrowOwnershipTransferred)
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
func (it *EscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOwnershipTransferred represents a OwnershipTransferred event raised by the Escrow contract.
type EscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Escrow *EscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EscrowOwnershipTransferredIterator{contract: _Escrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Escrow *EscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOwnershipTransferred)
				if err := _Escrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Escrow *EscrowFilterer) ParseOwnershipTransferred(log types.Log) (*EscrowOwnershipTransferred, error) {
	event := new(EscrowOwnershipTransferred)
	if err := _Escrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Escrow contract.
type EscrowWithdrawnIterator struct {
	Event *EscrowWithdrawn // Event containing the contract specifics and raw log

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
func (it *EscrowWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowWithdrawn)
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
		it.Event = new(EscrowWithdrawn)
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
func (it *EscrowWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowWithdrawn represents a Withdrawn event raised by the Escrow contract.
type EscrowWithdrawn struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_Escrow *EscrowFilterer) FilterWithdrawn(opts *bind.FilterOpts, payee []common.Address) (*EscrowWithdrawnIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return &EscrowWithdrawnIterator{contract: _Escrow.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_Escrow *EscrowFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EscrowWithdrawn, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowWithdrawn)
				if err := _Escrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_Escrow *EscrowFilterer) ParseWithdrawn(log types.Log) (*EscrowWithdrawn, error) {
	event := new(EscrowWithdrawn)
	if err := _Escrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalizableCrowdsaleABI is the input ABI used to generate the binding from.
const FinalizableCrowdsaleABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"CrowdsaleFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// FinalizableCrowdsaleFuncSigs maps the 4-byte function signature to its string representation.
var FinalizableCrowdsaleFuncSigs = map[string]string{
	"ec8ac4d8": "buyTokens(address)",
	"4b6753bc": "closingTime()",
	"4bb278f3": "finalize()",
	"b3f05b97": "finalized()",
	"1515bc2b": "hasClosed()",
	"47535d7b": "isOpen()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
}

// FinalizableCrowdsale is an auto generated Go binding around an Ethereum contract.
type FinalizableCrowdsale struct {
	FinalizableCrowdsaleCaller     // Read-only binding to the contract
	FinalizableCrowdsaleTransactor // Write-only binding to the contract
	FinalizableCrowdsaleFilterer   // Log filterer for contract events
}

// FinalizableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizableCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizableCrowdsaleSession struct {
	Contract     *FinalizableCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizableCrowdsaleCallerSession struct {
	Contract *FinalizableCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FinalizableCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizableCrowdsaleTransactorSession struct {
	Contract     *FinalizableCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FinalizableCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizableCrowdsaleRaw struct {
	Contract *FinalizableCrowdsale // Generic contract binding to access the raw methods on
}

// FinalizableCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleCallerRaw struct {
	Contract *FinalizableCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizableCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizableCrowdsaleTransactorRaw struct {
	Contract *FinalizableCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizableCrowdsale creates a new instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsale(address common.Address, backend bind.ContractBackend) (*FinalizableCrowdsale, error) {
	contract, err := bindFinalizableCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsale{FinalizableCrowdsaleCaller: FinalizableCrowdsaleCaller{contract: contract}, FinalizableCrowdsaleTransactor: FinalizableCrowdsaleTransactor{contract: contract}, FinalizableCrowdsaleFilterer: FinalizableCrowdsaleFilterer{contract: contract}}, nil
}

// NewFinalizableCrowdsaleCaller creates a new read-only instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*FinalizableCrowdsaleCaller, error) {
	contract, err := bindFinalizableCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleCaller{contract: contract}, nil
}

// NewFinalizableCrowdsaleTransactor creates a new write-only instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizableCrowdsaleTransactor, error) {
	contract, err := bindFinalizableCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleTransactor{contract: contract}, nil
}

// NewFinalizableCrowdsaleFilterer creates a new log filterer instance of FinalizableCrowdsale, bound to a specific deployed contract.
func NewFinalizableCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizableCrowdsaleFilterer, error) {
	contract, err := bindFinalizableCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleFilterer{contract: contract}, nil
}

// bindFinalizableCrowdsale binds a generic wrapper to an already deployed contract.
func bindFinalizableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsale *FinalizableCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsale.Contract.FinalizableCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsale *FinalizableCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.FinalizableCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsale *FinalizableCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.FinalizableCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalizableCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.ClosingTime(&_FinalizableCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.ClosingTime(&_FinalizableCrowdsale.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Finalized() (bool, error) {
	return _FinalizableCrowdsale.Contract.Finalized(&_FinalizableCrowdsale.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Finalized() (bool, error) {
	return _FinalizableCrowdsale.Contract.Finalized(&_FinalizableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) HasClosed() (bool, error) {
	return _FinalizableCrowdsale.Contract.HasClosed(&_FinalizableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _FinalizableCrowdsale.Contract.HasClosed(&_FinalizableCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) IsOpen() (bool, error) {
	return _FinalizableCrowdsale.Contract.IsOpen(&_FinalizableCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _FinalizableCrowdsale.Contract.IsOpen(&_FinalizableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.OpeningTime(&_FinalizableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.OpeningTime(&_FinalizableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.Rate(&_FinalizableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.Rate(&_FinalizableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Token() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Token(&_FinalizableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Token() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Token(&_FinalizableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Wallet(&_FinalizableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _FinalizableCrowdsale.Contract.Wallet(&_FinalizableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FinalizableCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.WeiRaised(&_FinalizableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_FinalizableCrowdsale *FinalizableCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _FinalizableCrowdsale.Contract.WeiRaised(&_FinalizableCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.BuyTokens(&_FinalizableCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.BuyTokens(&_FinalizableCrowdsale.TransactOpts, beneficiary)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Finalize(&_FinalizableCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Finalize(&_FinalizableCrowdsale.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Fallback(&_FinalizableCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Fallback(&_FinalizableCrowdsale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalizableCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Receive(&_FinalizableCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FinalizableCrowdsale *FinalizableCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _FinalizableCrowdsale.Contract.Receive(&_FinalizableCrowdsale.TransactOpts)
}

// FinalizableCrowdsaleCrowdsaleFinalizedIterator is returned from FilterCrowdsaleFinalized and is used to iterate over the raw logs and unpacked data for CrowdsaleFinalized events raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleCrowdsaleFinalizedIterator struct {
	Event *FinalizableCrowdsaleCrowdsaleFinalized // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleCrowdsaleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleCrowdsaleFinalized)
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
		it.Event = new(FinalizableCrowdsaleCrowdsaleFinalized)
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
func (it *FinalizableCrowdsaleCrowdsaleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleCrowdsaleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleCrowdsaleFinalized represents a CrowdsaleFinalized event raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleCrowdsaleFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrowdsaleFinalized is a free log retrieval operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) FilterCrowdsaleFinalized(opts *bind.FilterOpts) (*FinalizableCrowdsaleCrowdsaleFinalizedIterator, error) {

	logs, sub, err := _FinalizableCrowdsale.contract.FilterLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleCrowdsaleFinalizedIterator{contract: _FinalizableCrowdsale.contract, event: "CrowdsaleFinalized", logs: logs, sub: sub}, nil
}

// WatchCrowdsaleFinalized is a free log subscription operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) WatchCrowdsaleFinalized(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleCrowdsaleFinalized) (event.Subscription, error) {

	logs, sub, err := _FinalizableCrowdsale.contract.WatchLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleCrowdsaleFinalized)
				if err := _FinalizableCrowdsale.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
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

// ParseCrowdsaleFinalized is a log parse operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) ParseCrowdsaleFinalized(log types.Log) (*FinalizableCrowdsaleCrowdsaleFinalized, error) {
	event := new(FinalizableCrowdsaleCrowdsaleFinalized)
	if err := _FinalizableCrowdsale.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalizableCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *FinalizableCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(FinalizableCrowdsaleTimedCrowdsaleExtended)
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
func (it *FinalizableCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*FinalizableCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _FinalizableCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleTimedCrowdsaleExtendedIterator{contract: _FinalizableCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _FinalizableCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleTimedCrowdsaleExtended)
				if err := _FinalizableCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*FinalizableCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(FinalizableCrowdsaleTimedCrowdsaleExtended)
	if err := _FinalizableCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalizableCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleTokensPurchasedIterator struct {
	Event *FinalizableCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *FinalizableCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalizableCrowdsaleTokensPurchased)
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
		it.Event = new(FinalizableCrowdsaleTokensPurchased)
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
func (it *FinalizableCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalizableCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalizableCrowdsaleTokensPurchased represents a TokensPurchased event raised by the FinalizableCrowdsale contract.
type FinalizableCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*FinalizableCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &FinalizableCrowdsaleTokensPurchasedIterator{contract: _FinalizableCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *FinalizableCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _FinalizableCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalizableCrowdsaleTokensPurchased)
				if err := _FinalizableCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_FinalizableCrowdsale *FinalizableCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*FinalizableCrowdsaleTokensPurchased, error) {
	event := new(FinalizableCrowdsaleTokensPurchased)
	if err := _FinalizableCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOABI is the input ABI used to generate the binding from.
const HLBICOABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRateReceived\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"walletReceived\",\"type\":\"address\"},{\"internalType\":\"contractLBCToken\",\"name\":\"tokenReceived\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"openingTimeReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTimeReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goalReceived\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changerAddress\",\"type\":\"address\"}],\"name\":\"ChangedReserveAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelisterAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changerAddress\",\"type\":\"address\"}],\"name\":\"ChangedWhitelisterAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CrowdsaleFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelistingAddress\",\"type\":\"address\"}],\"name\":\"InitializedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGoal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTranche\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxInvest\",\"type\":\"uint256\"}],\"name\":\"UpdatedCaps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"_deployingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_reserveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_whitelistingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistedKYC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistedRegistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"coef\",\"type\":\"uint256\"}],\"name\":\"adjustEtherValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReserveAddress\",\"type\":\"address\"}],\"name\":\"changeReserveAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"changeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWhitelisterAddress\",\"type\":\"address\"}],\"name\":\"changeWhitelister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"}],\"name\":\"claimRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"etherTranche\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whitelistingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEtherToInvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKYCInvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRegisteredInvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// HLBICOFuncSigs maps the 4-byte function signature to its string representation.
var HLBICOFuncSigs = map[string]string{
	"1cd273cf": "_deployingAddress()",
	"ae5b5b11": "_reserveAddress()",
	"f5e343bb": "_whitelistingAddress()",
	"7c8b38c8": "addWhitelistedKYC(address)",
	"fc400611": "addWhitelistedRegistered(address)",
	"329425c5": "adjustEtherValue(uint256)",
	"70a08231": "balanceOf(address)",
	"ec8ac4d8": "buyTokens(address)",
	"355274ea": "cap()",
	"4f935945": "capReached()",
	"477fea02": "changeReserveAddress(address)",
	"66829b16": "changeToken(address)",
	"966aeece": "changeWhitelister(address)",
	"bffa55d5": "claimRefund(address)",
	"4b6753bc": "closingTime()",
	"3f701544": "etherTranche()",
	"4bb278f3": "finalize()",
	"b3f05b97": "finalized()",
	"40193883": "goal()",
	"7d3d6522": "goalReached()",
	"1515bc2b": "hasClosed()",
	"f09a4016": "init(address,address)",
	"158ef93e": "initialized()",
	"47535d7b": "isOpen()",
	"3af32abf": "isWhitelisted(address)",
	"cba8b279": "maxEtherToInvest()",
	"6ab84491": "maxKYCInvest()",
	"c559ea47": "maxRegisteredInvest()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"291d9549": "removeWhitelisted(address)",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
	"49df728c": "withdrawTokens(address)",
}

// HLBICOBin is the compiled bytecode used for deploying new contracts.
var HLBICOBin = "0x60806040523480156200001157600080fd5b50604051620036f5380380620036f5833981810160405260e08110156200003757600080fd5b508051602082015160408301516060840151608085015160a086015160c09096015160016000559495939492939192909190808285858a8a8a82620000c3576040805162461bcd60e51b815260206004820152601460248201527f43726f776473616c653a20726174652069732030000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166200010a5760405162461bcd60e51b8152600401808060200182810382526025815260200180620036d06025913960400191505060405180910390fd5b6001600160a01b038116620001515760405162461bcd60e51b8152600401808060200182810382526024815260200180620036ac6024913960400191505060405180910390fd5b600392909255600280546001600160a01b039283166001600160a01b0319918216179091556001805492909316911617905542821015620001c45760405162461bcd60e51b8152600401808060200182810382526033815260200180620036426033913960400191505060405180910390fd5b818111620002045760405162461bcd60e51b8152600401808060200182810382526037815260200180620036756037913960400191505060405180910390fd5b6005919091556006558062000260576040805162461bcd60e51b815260206004820152601960248201527f43617070656443726f776473616c653a20636170206973203000000000000000604482015290519081900360640190fd5b6007556008805460ff1916905580620002c0576040805162461bcd60e51b815260206004820152601e60248201527f526566756e6461626c6543726f776473616c653a20676f616c20697320300000604482015290519081900360640190fd5b620002ca620003b9565b604051620002d890620003c8565b6001600160a01b03909116815260405190819003602001906000f08015801562000306573d6000803e3d6000fd5b50600a80546001600160a01b0319166001600160a01b03929092169190911790556009556040516200033890620003d6565b604051809103906000f08015801562000355573d6000803e3d6000fd5b50600c80546001600160a01b03929092166001600160a01b0319928316179055600d8054909116331790555050681043561a88293000006014555050678ac7230489e800006012555050670de0b6b3a76400006013556010556000601555620003e4565b6002546001600160a01b031690565b610d4f806200253083390190565b6103c3806200327f83390190565b61213c80620003f46000396000f3fe6080604052600436106102085760003560e01c8063521eb27311610118578063b7a8807c116100a0578063ec8ac4d81161006f578063ec8ac4d814610606578063f09a40161461062c578063f5e343bb14610667578063fc0c546a1461067c578063fc400611146106915761021f565b8063b7a8807c146105a9578063bffa55d5146105be578063c559ea47146105f1578063cba8b279146104bc5761021f565b80637c8b38c8116100e75780637c8b38c8146105045780637d3d652214610537578063966aeece1461054c578063ae5b5b111461057f578063b3f05b97146105945761021f565b8063521eb2731461047457806366829b16146104895780636ab84491146104bc57806370a08231146104d15761021f565b80633f7015441161019b578063477fea021161016a578063477fea02146103cf57806349df728c146104025780634b6753bc146104355780634bb278f31461044a5780634f9359451461045f5761021f565b80633f7015441461037b57806340193883146103905780634042b66f146103a557806347535d7b146103ba5761021f565b80632c4e722e116101d75780632c4e722e146102cc578063329425c5146102f3578063355274ea1461031d5780633af32abf146103325761021f565b80631515bc2b1461022a578063158ef93e146102535780631cd273cf14610268578063291d9549146102995761021f565b3661021f5761021d6102186106c4565b6106c8565b005b61021d6102186106c4565b34801561023657600080fd5b5061023f6107d8565b604080519115158252519081900360200190f35b34801561025f57600080fd5b5061023f6107e7565b34801561027457600080fd5b5061027d6107f7565b604080516001600160a01b039092168252519081900360200190f35b3480156102a557600080fd5b5061021d600480360360208110156102bc57600080fd5b50356001600160a01b0316610806565b3480156102d857600080fd5b506102e161086b565b60408051918252519081900360200190f35b3480156102ff57600080fd5b5061021d6004803603602081101561031657600080fd5b5035610871565b34801561032957600080fd5b506102e16109f3565b34801561033e57600080fd5b506103656004803603602081101561035557600080fd5b50356001600160a01b03166109f9565b6040805160ff9092168252519081900360200190f35b34801561038757600080fd5b506102e1610a79565b34801561039c57600080fd5b506102e1610a7f565b3480156103b157600080fd5b506102e1610a85565b3480156103c657600080fd5b5061023f610a8b565b3480156103db57600080fd5b5061021d600480360360208110156103f257600080fd5b50356001600160a01b0316610aa4565b34801561040e57600080fd5b5061021d6004803603602081101561042557600080fd5b50356001600160a01b0316610b52565b34801561044157600080fd5b506102e1610be1565b34801561045657600080fd5b5061021d610be7565b34801561046b57600080fd5b5061023f610cc2565b34801561048057600080fd5b5061027d610cd6565b34801561049557600080fd5b5061021d600480360360208110156104ac57600080fd5b50356001600160a01b0316610ce5565b3480156104c857600080fd5b506102e1610d37565b3480156104dd57600080fd5b506102e1600480360360208110156104f457600080fd5b50356001600160a01b0316610d3d565b34801561051057600080fd5b5061021d6004803603602081101561052757600080fd5b50356001600160a01b0316610d58565b34801561054357600080fd5b5061023f610dbc565b34801561055857600080fd5b5061021d6004803603602081101561056f57600080fd5b50356001600160a01b0316610dc9565b34801561058b57600080fd5b5061027d610e77565b3480156105a057600080fd5b5061023f610e86565b3480156105b557600080fd5b506102e1610e8f565b3480156105ca57600080fd5b5061021d600480360360208110156105e157600080fd5b50356001600160a01b0316610e95565b3480156105fd57600080fd5b506102e1610f85565b61021d6004803603602081101561061c57600080fd5b50356001600160a01b03166106c8565b34801561063857600080fd5b5061021d6004803603604081101561064f57600080fd5b506001600160a01b0381358116916020013516610f8b565b34801561067357600080fd5b5061027d611129565b34801561068857600080fd5b5061027d611138565b34801561069d57600080fd5b5061021d600480360360208110156106b457600080fd5b50356001600160a01b0316611147565b3390565b60026000541415610720576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026000553461073082826111ab565b600061073b82611215565b60045490915061074b9083611226565b6004556107588382611287565b826001600160a01b031661076a6106c4565b6001600160a01b03167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b8484604051808381526020018281526020019250505060405180910390a36107bc8383611211565b6107c4611291565b6107ce8383611211565b5050600160005550565b60006107e261129b565b905090565b600c54600160a01b900460ff1681565b600d546001600160a01b031681565b600e546001600160a01b031661081a6106c4565b6001600160a01b03161461085f5760405162461bcd60e51b8152600401808060200182810382526031815260200180611e3d6031913960400191505060405180910390fd5b610868816112b3565b50565b60105490565b600d546001600160a01b031633146108ba5760405162461bcd60e51b81526004018080602001828103825260388152602001806120ad6038913960400191505060405180910390fd5b6000811180156108cb575061271081105b6109065760405162461bcd60e51b8152600401808060200182810382526034815260200180611d826034913960400191505060405180910390fd5b61092c6109276103e86109218461091b610a7f565b90611346565b9061139f565b6113e1565b6109466109416103e86109218461091b6109f3565b6113e6565b6109616103e86109218360145461134690919063ffffffff16565b601455601254610979906103e8906109219084611346565b601255601354610991906103e8906109219084611346565b6013557f0c8ac593d921e6388dbbe810ac16aa0fa330ca2fae1ff2b027e6b8c2bb0f92426109bd610a7f565b6109c56109f3565b601454601254604080519485526020850193909352838301919091526060830152519081900360800190a150565b60075490565b60006001600160a01b038216610a56576040805162461bcd60e51b815260206004820152601f60248201527f484c4349434f3a206163636f756e74206973207a65726f206164647265737300604482015290519081900360640190fd5b506001600160a01b03811660009081526011602052604090205460ff165b919050565b60145490565b60095490565b60045490565b600060055442101580156107e257505060065442111590565b600d546001600160a01b03163314610aed5760405162461bcd60e51b81526004018080602001828103825260388152602001806120ad6038913960400191505060405180910390fd5b600f80546001600160a01b0319166001600160a01b038316179055610b106106c4565b6001600160a01b0316816001600160a01b03167f7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b5460405160405180910390a350565b610b5a610e86565b610b955760405162461bcd60e51b815260040180806020018281038252602e815260200180611dde602e913960400191505060405180910390fd5b610b9d610dbc565b610bd85760405162461bcd60e51b8152600401808060200182810382526031815260200180611e0c6031913960400191505060405180910390fd5b610868816113eb565b60065490565b60085460ff1615610c295760405162461bcd60e51b8152600401808060200182810382526027815260200180611f896027913960400191505060405180910390fd5b610c316107d8565b610c82576040805162461bcd60e51b815260206004820181905260248201527f46696e616c697a61626c6543726f776473616c653a206e6f7420636c6f736564604482015290519081900360640190fd5b6008805460ff19166001179055610c97611522565b6040517f9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf73183690600090a1565b6000600754610ccf610a85565b1015905090565b6002546001600160a01b031690565b600d546001600160a01b03163314610d2e5760405162461bcd60e51b81526004018080602001828103825260388152602001806120ad6038913960400191505060405180910390fd5b6108688161155e565b60125490565b6001600160a01b03166000908152600b602052604090205490565b600e546001600160a01b0316610d6c6106c4565b6001600160a01b031614610db15760405162461bcd60e51b8152600401808060200182810382526031815260200180611e3d6031913960400191505060405180910390fd5b610868816002611580565b6000600954610ccf610a85565b600d546001600160a01b03163314610e125760405162461bcd60e51b81526004018080602001828103825260388152602001806120ad6038913960400191505060405180910390fd5b600e80546001600160a01b0319166001600160a01b038316179055610e356106c4565b6001600160a01b0316816001600160a01b03167f63c70fa8d3aed6b9cda2842d4152a4ccf78aab6db236b9b22ad898ca7541e11360405160405180910390a350565b600f546001600160a01b031681565b60085460ff1690565b60055490565b610e9d610e86565b610ed85760405162461bcd60e51b81526004018080602001828103825260228152602001806120696022913960400191505060405180910390fd5b610ee0610dbc565b15610f1c5760405162461bcd60e51b8152600401808060200182810382526021815260200180611fd56021913960400191505060405180910390fd5b600a54604080516351cff8d960e01b81526001600160a01b038481166004830152915191909216916351cff8d991602480830192600092919082900301818387803b158015610f6a57600080fd5b505af1158015610f7e573d6000803e3d6000fd5b5050505050565b60135490565b600c54600160a01b900460ff1615610fd45760405162461bcd60e51b8152600401808060200182810382526028815260200180611e6e6028913960400191505060405180910390fd5b600d546001600160a01b0316331461101d5760405162461bcd60e51b81526004018080602001828103825260388152602001806120ad6038913960400191505060405180910390fd5b6001600160a01b0382166110625760405162461bcd60e51b81526004018080602001828103825260288152602001806120416028913960400191505060405180910390fd5b6001600160a01b0381166110a75760405162461bcd60e51b8152600401808060200182810382526023815260200180611f0d6023913960400191505060405180910390fd5b600e80546001600160a01b038085166001600160a01b03199283168117909355600f805491851691909216179055600c805460ff60a01b1916600160a01b1790556110f06106c4565b6001600160a01b03167f1136ccf874c0c85bed8c1488e35920195c98fda6c939c473a46510d35826d5a160405160405180910390a35050565b600e546001600160a01b031681565b6001546001600160a01b031690565b600e546001600160a01b031661115b6106c4565b6001600160a01b0316146111a05760405162461bcd60e51b8152600401808060200182810382526031815260200180611e3d6031913960400191505060405180910390fd5b610868816001611580565b60006111b6836109f9565b60ff16116111f55760405162461bcd60e51b815260040180806020018281038252602281526020018061208b6022913960400191505060405180910390fd5b61120781611202846109f9565b61166b565b611211828261170d565b5050565b60006112208261177f565b92915050565b600082820183811015611280576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b611211828261185e565b611299611868565b565b60006112a5610cc2565b806107e257506107e2611870565b60006112be826109f9565b60ff16116112fd5760405162461bcd60e51b81526004018080602001828103825260228152602001806120e56022913960400191505060405180910390fd5b6001600160a01b038116600081815260116020526040808220805460ff19169055517f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b69190a250565b60008261135557506000611220565b8282028284828161136257fe5b04146112805760405162461bcd60e51b8152600401808060200182810382526021815260200180611f686021913960400191505060405180910390fd5b600061128083836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250611878565b600955565b600755565b6113f36107d8565b61142e5760405162461bcd60e51b81526004018080602001828103825260218152602001806120206021913960400191505060405180910390fd5b6001600160a01b0381166000908152600b6020526040902054806114835760405162461bcd60e51b8152600401808060200182810382526038815260200180611f306038913960400191505060405180910390fd5b6001600160a01b038083166000908152600b6020526040812055600c541663beabacc86114ae611138565b84846040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b031681526020018281526020019350505050600060405180830381600087803b15801561150657600080fd5b505af115801561151a573d6000803e3d6000fd5b505050505050565b61152a610dbc565b1561155657600f54611556906001600160a01b03166115516064610921600561091b610a85565b61191a565b6112996119f2565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b8060ff166001148061159557508060ff166002145b6115d05760405162461bcd60e51b8152600401808060200182810382526028815260200180611db66028913960400191505060405180910390fd5b8060ff166115dd836109f9565b60ff161061161c5760405162461bcd60e51b8152600401808060200182810382526023815260200180611e966023913960400191505060405180910390fd5b6001600160a01b038216600081815260116020526040808220805460ff191660ff8616179055517fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f9190a25050565b611673610d37565b8211156116b15760405162461bcd60e51b8152600401808060200182810382526025815260200180611fb06025913960400191505060405180910390fd5b601354821180156116c557508060ff166002145b806116d257506013548211155b6112115760405162461bcd60e51b8152600401808060200182810382526054815260200180611eb96054913960600191505060405180910390fd5b6117178282611b45565b60075461172c82611726610a85565b90611226565b1115611211576040805162461bcd60e51b815260206004820152601d60248201527f43617070656443726f776473616c653a20636170206578636565646564000000604482015290519081900360640190fd5b6000611789610a8b565b61179557506000610a74565b6015546000906117a59084611226565b60158190556014541015611825576014546015546117c291611ba8565b60158190556000906117d5908590611ba8565b90506117fa6117f36117e561086b565b60155461091b908890611ba8565b8390611226565b6010805460ae19019055915061181d6117f361181461086b565b60155490611346565b915050611844565b61184161183a61183361086b565b8590611346565b8290611226565b90505b6112806118576064610921846005611346565b8290611ba8565b6112118282611bea565b611299611c38565b600654421190565b600081836119045760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156118c95781810151838201526020016118b1565b50505050905090810190601f1680156118f65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161191057fe5b0495945050505050565b600154604080516340c10f1960e01b81526001600160a01b03858116600483015260248201859052915191909216916340c10f199160448083019260209291908290030181600087803b15801561197057600080fd5b505af1158015611984573d6000803e3d6000fd5b505050506040513d602081101561199a57600080fd5b50511515600114611211576040805162461bcd60e51b815260206004820152601960248201527f43726f776473616c653a206d696e74696e67206661696c656400000000000000604482015290519081900360640190fd5b6119fa610dbc565b15611ad457600a60009054906101000a90046001600160a01b03166001600160a01b03166343d726d66040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611a4f57600080fd5b505af1158015611a63573d6000803e3d6000fd5b50505050600a60009054906101000a90046001600160a01b03166001600160a01b0316639af6549a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611ab757600080fd5b505af1158015611acb573d6000803e3d6000fd5b50505050611b3d565b600a60009054906101000a90046001600160a01b03166001600160a01b0316638c52dc416040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611b2457600080fd5b505af1158015611b38573d6000803e3d6000fd5b505050505b611299611299565b611b4d610a8b565b611b9e576040805162461bcd60e51b815260206004820152601860248201527f54696d656443726f776473616c653a206e6f74206f70656e0000000000000000604482015290519081900360640190fd5b6112118282611c90565b600061128083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611d27565b6001600160a01b0382166000908152600b6020526040902054611c0d9082611226565b6001600160a01b038084166000908152600b6020526040902091909155600c5461121191168261191a565b600a546001600160a01b031663f340fa0134611c526106c4565b6040518363ffffffff1660e01b815260040180826001600160a01b031681526020019150506000604051808303818588803b158015610f6a57600080fd5b6001600160a01b038216611cd55760405162461bcd60e51b815260040180806020018281038252602a815260200180611ff6602a913960400191505060405180910390fd5b80611211576040805162461bcd60e51b815260206004820152601960248201527f43726f776473616c653a20776569416d6f756e74206973203000000000000000604482015290519081900360640190fd5b60008184841115611d795760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156118c95781810151838201526020016118b1565b50505090039056fe484c4249434f3a20636f65662069736e27742077697468696e2072616e6765206f6620617574686f72697a65642076616c756573484c4249434f3a2077686974656c697374696e6720666c6167206d7573742062652031206f722032526566756e6461626c65506f737444656c697665727943726f776473616c653a206e6f742066696e616c697a6564526566756e6461626c65506f737444656c697665727943726f776473616c653a20676f616c206e6f742072656163686564484c4249434f3a2063616c6c657220646f6573206e6f742068617665207468652057686974656c697374656420726f6c65484c4249434f3a20636f6e747261637420697320616c726561647920696e697469616c697a65642e484c4349434f3a206163636f756e7420616c72656164792077686974656c6973746564484c4249434f3a2043616e6e6f7420696e76657374206d6f7265207468616e204e6f4b5943206c696d69742e2055736572206e6565647320746f2070617373204b59432073637265656e696e672066697273742e484c4249434f3a2072657365727665416464726573732063616e6e6f74206265203078506f737444656c697665727943726f776473616c653a2062656e6566696369617279206973206e6f742064756520616e7920746f6b656e73536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7746696e616c697a61626c6543726f776473616c653a20616c72656164792066696e616c697a6564484c4249434f3a2043616e6e6f7420696e76657374206d6f7265207468616e206c696d6974526566756e6461626c6543726f776473616c653a20676f616c207265616368656443726f776473616c653a2062656e656669636961727920697320746865207a65726f2061646472657373506f737444656c697665727943726f776473616c653a206e6f7420636c6f736564484c4249434f3a2077686974656c697374696e67416464726573732063616e6e6f74206265203078526566756e6461626c6543726f776473616c653a206e6f742066696e616c697a6564484c4249434f3a204163636f756e74206973206e6f742077686974656c6973746564484c4249434f3a206f6e6c7920746865206465706c6f79696e6720616464726573732063616e2063616c6c2074686973206d6574686f642e484c4349434f3a206163636f756e74206973206e6f742077686974656c6973746564a2646970667358221220e902dc50ab5709537e7530a6bec9528b763cced88ba7c67e7d0202939db6f65e64736f6c63430007060033608060405234801561001057600080fd5b50604051610d4f380380610d4f8339818101604052602081101561003357600080fd5b5051600061003f6100fe565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001600160a01b0381166100ce5760405162461bcd60e51b815260040180806020018281038252602d815260200180610d22602d913960400191505060405180910390fd5b6002805460ff196001600160a01b039390931661010002610100600160a81b031990911617919091169055610102565b3390565b610c11806101116000396000f3fe6080604052600436106100a75760003560e01c80638da5cb5b116100645780638da5cb5b146101985780639af6549a146101ad578063c19d93fb146101c2578063e3a9db1a146101f8578063f2fde38b1461023d578063f340fa0114610270576100a7565b806338af3eed146100ac57806343d726d6146100dd57806351cff8d9146100f4578063685ca19414610127578063715018a61461016e5780638c52dc4114610183575b600080fd5b3480156100b857600080fd5b506100c1610296565b604080516001600160a01b039092168252519081900360200190f35b3480156100e957600080fd5b506100f26102aa565b005b34801561010057600080fd5b506100f26004803603602081101561011757600080fd5b50356001600160a01b0316610388565b34801561013357600080fd5b5061015a6004803603602081101561014a57600080fd5b50356001600160a01b03166103d8565b604080519115158252519081900360200190f35b34801561017a57600080fd5b506100f26103f4565b34801561018f57600080fd5b506100f2610496565b3480156101a457600080fd5b506100c1610575565b3480156101b957600080fd5b506100f2610584565b3480156101ce57600080fd5b506101d7610611565b604051808260028111156101e757fe5b815260200191505060405180910390f35b34801561020457600080fd5b5061022b6004803603602081101561021b57600080fd5b50356001600160a01b031661061a565b60408051918252519081900360200190f35b34801561024957600080fd5b506100f26004803603602081101561026057600080fd5b50356001600160a01b0316610635565b6100f26004803603602081101561028657600080fd5b50356001600160a01b031661072d565b60025461010090046001600160a01b031690565b6102b2610785565b6000546001600160a01b03908116911614610302576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b60006002805460ff169081111561031557fe5b146103515760405162461bcd60e51b8152600401808060200182810382526029815260200180610b616029913960400191505060405180910390fd5b6002805460ff1916811790556040517f088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f90600090a1565b610391816103d8565b6103cc5760405162461bcd60e51b8152600401808060200182810382526033815260200180610b2e6033913960400191505060405180910390fd5b6103d581610789565b50565b600060016002805460ff16908111156103ed57fe5b1492915050565b6103fc610785565b6000546001600160a01b0390811691161461044c576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b61049e610785565b6000546001600160a01b039081169116146104ee576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b60006002805460ff169081111561050157fe5b1461053d5760405162461bcd60e51b8152600401808060200182810382526032815260200180610baa6032913960400191505060405180910390fd5b6002805460ff191660011790556040517f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8990600090a1565b6000546001600160a01b031690565b6002805460ff168181111561059557fe5b146105d15760405162461bcd60e51b8152600401808060200182810382526038815260200180610a6b6038913960400191505060405180910390fd5b6002546040516001600160a01b0361010090920491909116904780156108fc02916000818181858888f193505050501580156103d5573d6000803e3d6000fd5b60025460ff1690565b6001600160a01b031660009081526001602052604090205490565b61063d610785565b6000546001600160a01b0390811691161461068d576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b6001600160a01b0381166106d25760405162461bcd60e51b8152600401808060200182810382526026815260200180610aa36026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006002805460ff169081111561074057fe5b1461077c5760405162461bcd60e51b815260040180806020018281038252602b815260200180610b03602b913960400191505060405180910390fd5b6103d58161084c565b3390565b610791610785565b6000546001600160a01b039081169116146107e1576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b6001600160a01b0381166000818152600160205260408120805491905590610809908261091f565b6040805182815290516001600160a01b038416917f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5919081900360200190a25050565b610854610785565b6000546001600160a01b039081169116146108a4576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526001602052604090205434906108c99082610a09565b6001600160a01b038316600081815260016020908152604091829020939093558051848152905191927f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c492918290030190a25050565b80471015610974576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d80600081146109bf576040519150601f19603f3d011682016040523d82523d6000602084013e6109c4565b606091505b5050905080610a045760405162461bcd60e51b815260040180806020018281038252603a815260200180610ac9603a913960400191505060405180910390fd5b505050565b600082820183811015610a63576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe526566756e64457363726f773a2062656e65666963696172792063616e206f6e6c79207769746864726177207768696c6520636c6f7365644f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d61792068617665207265766572746564526566756e64457363726f773a2063616e206f6e6c79206465706f736974207768696c6520616374697665436f6e646974696f6e616c457363726f773a207061796565206973206e6f7420616c6c6f77656420746f207769746864726177526566756e64457363726f773a2063616e206f6e6c7920636c6f7365207768696c65206163746976654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572526566756e64457363726f773a2063616e206f6e6c7920656e61626c6520726566756e6473207768696c6520616374697665a2646970667358221220415d58c8a6368e7076ac118e308228241ab855c49dceef8931507f3d6665159d64736f6c63430007060033526566756e64457363726f773a2062656e656669636961727920697320746865207a65726f2061646472657373608060405234801561001057600080fd5b50600061001b610076565b600080546001600160a01b0383166001600160a01b0319909116811790915560408051918252519192507f4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9919081900360200190a15061007a565b3390565b61033a806100896000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632348238c14610046578063beabacc81461006e578063c6dbdf61146100a4575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100c8565b005b61006c6004803603606081101561008457600080fd5b506001600160a01b038135811691602081013590911690604001356101ba565b6100ac61029b565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b03166100dc6102aa565b6001600160a01b0316146101215760405162461bcd60e51b815260040180806020018281038252602c8152602001806102d9602c913960400191505060405180910390fd5b6001600160a01b0381166101665760405162461bcd60e51b815260040180806020018281038252602a8152602001806102af602a913960400191505060405180910390fd5b600080546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d99181900360200190a150565b6000546001600160a01b03166101ce6102aa565b6001600160a01b0316146102135760405162461bcd60e51b815260040180806020018281038252602c8152602001806102d9602c913960400191505060405180910390fd5b826001600160a01b031663a9059cbb83836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561026a57600080fd5b505af115801561027e573d6000803e3d6000fd5b505050506040513d602081101561029457600080fd5b5050505050565b6000546001600160a01b031690565b339056fe5365636f6e646172793a206e6577207072696d61727920697320746865207a65726f20616464726573735365636f6e646172793a2063616c6c6572206973206e6f7420746865207072696d617279206163636f756e74a2646970667358221220cdd24c05d08289961d37f26f404cdfc4688c42bbea0ae52bbe09cd50a126f02464736f6c6343000706003354696d656443726f776473616c653a206f70656e696e672074696d65206973206265666f72652063757272656e742074696d6554696d656443726f776473616c653a206f70656e696e672074696d65206973206e6f74206265666f726520636c6f73696e672074696d6543726f776473616c653a20746f6b656e20697320746865207a65726f206164647265737343726f776473616c653a2077616c6c657420697320746865207a65726f2061646472657373"

// DeployHLBICO deploys a new Ethereum contract, binding an instance of HLBICO to it.
func DeployHLBICO(auth *bind.TransactOpts, backend bind.ContractBackend, initialRateReceived *big.Int, walletReceived common.Address, tokenReceived common.Address, openingTimeReceived *big.Int, closingTimeReceived *big.Int, capReceived *big.Int, goalReceived *big.Int) (common.Address, *types.Transaction, *HLBICO, error) {
	parsed, err := abi.JSON(strings.NewReader(HLBICOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HLBICOBin), backend, initialRateReceived, walletReceived, tokenReceived, openingTimeReceived, closingTimeReceived, capReceived, goalReceived)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HLBICO{HLBICOCaller: HLBICOCaller{contract: contract}, HLBICOTransactor: HLBICOTransactor{contract: contract}, HLBICOFilterer: HLBICOFilterer{contract: contract}}, nil
}

// HLBICO is an auto generated Go binding around an Ethereum contract.
type HLBICO struct {
	HLBICOCaller     // Read-only binding to the contract
	HLBICOTransactor // Write-only binding to the contract
	HLBICOFilterer   // Log filterer for contract events
}

// HLBICOCaller is an auto generated read-only Go binding around an Ethereum contract.
type HLBICOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HLBICOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HLBICOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HLBICOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HLBICOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HLBICOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HLBICOSession struct {
	Contract     *HLBICO           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HLBICOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HLBICOCallerSession struct {
	Contract *HLBICOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HLBICOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HLBICOTransactorSession struct {
	Contract     *HLBICOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HLBICORaw is an auto generated low-level Go binding around an Ethereum contract.
type HLBICORaw struct {
	Contract *HLBICO // Generic contract binding to access the raw methods on
}

// HLBICOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HLBICOCallerRaw struct {
	Contract *HLBICOCaller // Generic read-only contract binding to access the raw methods on
}

// HLBICOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HLBICOTransactorRaw struct {
	Contract *HLBICOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHLBICO creates a new instance of HLBICO, bound to a specific deployed contract.
func NewHLBICO(address common.Address, backend bind.ContractBackend) (*HLBICO, error) {
	contract, err := bindHLBICO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HLBICO{HLBICOCaller: HLBICOCaller{contract: contract}, HLBICOTransactor: HLBICOTransactor{contract: contract}, HLBICOFilterer: HLBICOFilterer{contract: contract}}, nil
}

// NewHLBICOCaller creates a new read-only instance of HLBICO, bound to a specific deployed contract.
func NewHLBICOCaller(address common.Address, caller bind.ContractCaller) (*HLBICOCaller, error) {
	contract, err := bindHLBICO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HLBICOCaller{contract: contract}, nil
}

// NewHLBICOTransactor creates a new write-only instance of HLBICO, bound to a specific deployed contract.
func NewHLBICOTransactor(address common.Address, transactor bind.ContractTransactor) (*HLBICOTransactor, error) {
	contract, err := bindHLBICO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HLBICOTransactor{contract: contract}, nil
}

// NewHLBICOFilterer creates a new log filterer instance of HLBICO, bound to a specific deployed contract.
func NewHLBICOFilterer(address common.Address, filterer bind.ContractFilterer) (*HLBICOFilterer, error) {
	contract, err := bindHLBICO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HLBICOFilterer{contract: contract}, nil
}

// bindHLBICO binds a generic wrapper to an already deployed contract.
func bindHLBICO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HLBICOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HLBICO *HLBICORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HLBICO.Contract.HLBICOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HLBICO *HLBICORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HLBICO.Contract.HLBICOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HLBICO *HLBICORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HLBICO.Contract.HLBICOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HLBICO *HLBICOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HLBICO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HLBICO *HLBICOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HLBICO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HLBICO *HLBICOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HLBICO.Contract.contract.Transact(opts, method, params...)
}

// DeployingAddress is a free data retrieval call binding the contract method 0x1cd273cf.
//
// Solidity: function _deployingAddress() view returns(address)
func (_HLBICO *HLBICOCaller) DeployingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "_deployingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployingAddress is a free data retrieval call binding the contract method 0x1cd273cf.
//
// Solidity: function _deployingAddress() view returns(address)
func (_HLBICO *HLBICOSession) DeployingAddress() (common.Address, error) {
	return _HLBICO.Contract.DeployingAddress(&_HLBICO.CallOpts)
}

// DeployingAddress is a free data retrieval call binding the contract method 0x1cd273cf.
//
// Solidity: function _deployingAddress() view returns(address)
func (_HLBICO *HLBICOCallerSession) DeployingAddress() (common.Address, error) {
	return _HLBICO.Contract.DeployingAddress(&_HLBICO.CallOpts)
}

// ReserveAddress is a free data retrieval call binding the contract method 0xae5b5b11.
//
// Solidity: function _reserveAddress() view returns(address)
func (_HLBICO *HLBICOCaller) ReserveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "_reserveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveAddress is a free data retrieval call binding the contract method 0xae5b5b11.
//
// Solidity: function _reserveAddress() view returns(address)
func (_HLBICO *HLBICOSession) ReserveAddress() (common.Address, error) {
	return _HLBICO.Contract.ReserveAddress(&_HLBICO.CallOpts)
}

// ReserveAddress is a free data retrieval call binding the contract method 0xae5b5b11.
//
// Solidity: function _reserveAddress() view returns(address)
func (_HLBICO *HLBICOCallerSession) ReserveAddress() (common.Address, error) {
	return _HLBICO.Contract.ReserveAddress(&_HLBICO.CallOpts)
}

// WhitelistingAddress is a free data retrieval call binding the contract method 0xf5e343bb.
//
// Solidity: function _whitelistingAddress() view returns(address)
func (_HLBICO *HLBICOCaller) WhitelistingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "_whitelistingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistingAddress is a free data retrieval call binding the contract method 0xf5e343bb.
//
// Solidity: function _whitelistingAddress() view returns(address)
func (_HLBICO *HLBICOSession) WhitelistingAddress() (common.Address, error) {
	return _HLBICO.Contract.WhitelistingAddress(&_HLBICO.CallOpts)
}

// WhitelistingAddress is a free data retrieval call binding the contract method 0xf5e343bb.
//
// Solidity: function _whitelistingAddress() view returns(address)
func (_HLBICO *HLBICOCallerSession) WhitelistingAddress() (common.Address, error) {
	return _HLBICO.Contract.WhitelistingAddress(&_HLBICO.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HLBICO *HLBICOCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HLBICO *HLBICOSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HLBICO.Contract.BalanceOf(&_HLBICO.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HLBICO *HLBICOCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HLBICO.Contract.BalanceOf(&_HLBICO.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_HLBICO *HLBICOCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_HLBICO *HLBICOSession) Cap() (*big.Int, error) {
	return _HLBICO.Contract.Cap(&_HLBICO.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) Cap() (*big.Int, error) {
	return _HLBICO.Contract.Cap(&_HLBICO.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_HLBICO *HLBICOCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_HLBICO *HLBICOSession) CapReached() (bool, error) {
	return _HLBICO.Contract.CapReached(&_HLBICO.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_HLBICO *HLBICOCallerSession) CapReached() (bool, error) {
	return _HLBICO.Contract.CapReached(&_HLBICO.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_HLBICO *HLBICOCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_HLBICO *HLBICOSession) ClosingTime() (*big.Int, error) {
	return _HLBICO.Contract.ClosingTime(&_HLBICO.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) ClosingTime() (*big.Int, error) {
	return _HLBICO.Contract.ClosingTime(&_HLBICO.CallOpts)
}

// EtherTranche is a free data retrieval call binding the contract method 0x3f701544.
//
// Solidity: function etherTranche() view returns(uint256)
func (_HLBICO *HLBICOCaller) EtherTranche(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "etherTranche")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EtherTranche is a free data retrieval call binding the contract method 0x3f701544.
//
// Solidity: function etherTranche() view returns(uint256)
func (_HLBICO *HLBICOSession) EtherTranche() (*big.Int, error) {
	return _HLBICO.Contract.EtherTranche(&_HLBICO.CallOpts)
}

// EtherTranche is a free data retrieval call binding the contract method 0x3f701544.
//
// Solidity: function etherTranche() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) EtherTranche() (*big.Int, error) {
	return _HLBICO.Contract.EtherTranche(&_HLBICO.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_HLBICO *HLBICOCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_HLBICO *HLBICOSession) Finalized() (bool, error) {
	return _HLBICO.Contract.Finalized(&_HLBICO.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_HLBICO *HLBICOCallerSession) Finalized() (bool, error) {
	return _HLBICO.Contract.Finalized(&_HLBICO.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_HLBICO *HLBICOCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "goal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_HLBICO *HLBICOSession) Goal() (*big.Int, error) {
	return _HLBICO.Contract.Goal(&_HLBICO.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) Goal() (*big.Int, error) {
	return _HLBICO.Contract.Goal(&_HLBICO.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_HLBICO *HLBICOCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "goalReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_HLBICO *HLBICOSession) GoalReached() (bool, error) {
	return _HLBICO.Contract.GoalReached(&_HLBICO.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_HLBICO *HLBICOCallerSession) GoalReached() (bool, error) {
	return _HLBICO.Contract.GoalReached(&_HLBICO.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_HLBICO *HLBICOCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_HLBICO *HLBICOSession) HasClosed() (bool, error) {
	return _HLBICO.Contract.HasClosed(&_HLBICO.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_HLBICO *HLBICOCallerSession) HasClosed() (bool, error) {
	return _HLBICO.Contract.HasClosed(&_HLBICO.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_HLBICO *HLBICOCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_HLBICO *HLBICOSession) Initialized() (bool, error) {
	return _HLBICO.Contract.Initialized(&_HLBICO.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_HLBICO *HLBICOCallerSession) Initialized() (bool, error) {
	return _HLBICO.Contract.Initialized(&_HLBICO.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_HLBICO *HLBICOCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_HLBICO *HLBICOSession) IsOpen() (bool, error) {
	return _HLBICO.Contract.IsOpen(&_HLBICO.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_HLBICO *HLBICOCallerSession) IsOpen() (bool, error) {
	return _HLBICO.Contract.IsOpen(&_HLBICO.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(uint8)
func (_HLBICO *HLBICOCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (uint8, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(uint8)
func (_HLBICO *HLBICOSession) IsWhitelisted(account common.Address) (uint8, error) {
	return _HLBICO.Contract.IsWhitelisted(&_HLBICO.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(uint8)
func (_HLBICO *HLBICOCallerSession) IsWhitelisted(account common.Address) (uint8, error) {
	return _HLBICO.Contract.IsWhitelisted(&_HLBICO.CallOpts, account)
}

// MaxEtherToInvest is a free data retrieval call binding the contract method 0xcba8b279.
//
// Solidity: function maxEtherToInvest() view returns(uint256)
func (_HLBICO *HLBICOCaller) MaxEtherToInvest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "maxEtherToInvest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEtherToInvest is a free data retrieval call binding the contract method 0xcba8b279.
//
// Solidity: function maxEtherToInvest() view returns(uint256)
func (_HLBICO *HLBICOSession) MaxEtherToInvest() (*big.Int, error) {
	return _HLBICO.Contract.MaxEtherToInvest(&_HLBICO.CallOpts)
}

// MaxEtherToInvest is a free data retrieval call binding the contract method 0xcba8b279.
//
// Solidity: function maxEtherToInvest() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) MaxEtherToInvest() (*big.Int, error) {
	return _HLBICO.Contract.MaxEtherToInvest(&_HLBICO.CallOpts)
}

// MaxKYCInvest is a free data retrieval call binding the contract method 0x6ab84491.
//
// Solidity: function maxKYCInvest() view returns(uint256)
func (_HLBICO *HLBICOCaller) MaxKYCInvest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "maxKYCInvest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxKYCInvest is a free data retrieval call binding the contract method 0x6ab84491.
//
// Solidity: function maxKYCInvest() view returns(uint256)
func (_HLBICO *HLBICOSession) MaxKYCInvest() (*big.Int, error) {
	return _HLBICO.Contract.MaxKYCInvest(&_HLBICO.CallOpts)
}

// MaxKYCInvest is a free data retrieval call binding the contract method 0x6ab84491.
//
// Solidity: function maxKYCInvest() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) MaxKYCInvest() (*big.Int, error) {
	return _HLBICO.Contract.MaxKYCInvest(&_HLBICO.CallOpts)
}

// MaxRegisteredInvest is a free data retrieval call binding the contract method 0xc559ea47.
//
// Solidity: function maxRegisteredInvest() view returns(uint256)
func (_HLBICO *HLBICOCaller) MaxRegisteredInvest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "maxRegisteredInvest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRegisteredInvest is a free data retrieval call binding the contract method 0xc559ea47.
//
// Solidity: function maxRegisteredInvest() view returns(uint256)
func (_HLBICO *HLBICOSession) MaxRegisteredInvest() (*big.Int, error) {
	return _HLBICO.Contract.MaxRegisteredInvest(&_HLBICO.CallOpts)
}

// MaxRegisteredInvest is a free data retrieval call binding the contract method 0xc559ea47.
//
// Solidity: function maxRegisteredInvest() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) MaxRegisteredInvest() (*big.Int, error) {
	return _HLBICO.Contract.MaxRegisteredInvest(&_HLBICO.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_HLBICO *HLBICOCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_HLBICO *HLBICOSession) OpeningTime() (*big.Int, error) {
	return _HLBICO.Contract.OpeningTime(&_HLBICO.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) OpeningTime() (*big.Int, error) {
	return _HLBICO.Contract.OpeningTime(&_HLBICO.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_HLBICO *HLBICOCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_HLBICO *HLBICOSession) Rate() (*big.Int, error) {
	return _HLBICO.Contract.Rate(&_HLBICO.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) Rate() (*big.Int, error) {
	return _HLBICO.Contract.Rate(&_HLBICO.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_HLBICO *HLBICOCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_HLBICO *HLBICOSession) Token() (common.Address, error) {
	return _HLBICO.Contract.Token(&_HLBICO.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_HLBICO *HLBICOCallerSession) Token() (common.Address, error) {
	return _HLBICO.Contract.Token(&_HLBICO.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_HLBICO *HLBICOCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_HLBICO *HLBICOSession) Wallet() (common.Address, error) {
	return _HLBICO.Contract.Wallet(&_HLBICO.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_HLBICO *HLBICOCallerSession) Wallet() (common.Address, error) {
	return _HLBICO.Contract.Wallet(&_HLBICO.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_HLBICO *HLBICOCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HLBICO.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_HLBICO *HLBICOSession) WeiRaised() (*big.Int, error) {
	return _HLBICO.Contract.WeiRaised(&_HLBICO.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_HLBICO *HLBICOCallerSession) WeiRaised() (*big.Int, error) {
	return _HLBICO.Contract.WeiRaised(&_HLBICO.CallOpts)
}

// AddWhitelistedKYC is a paid mutator transaction binding the contract method 0x7c8b38c8.
//
// Solidity: function addWhitelistedKYC(address account) returns()
func (_HLBICO *HLBICOTransactor) AddWhitelistedKYC(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "addWhitelistedKYC", account)
}

// AddWhitelistedKYC is a paid mutator transaction binding the contract method 0x7c8b38c8.
//
// Solidity: function addWhitelistedKYC(address account) returns()
func (_HLBICO *HLBICOSession) AddWhitelistedKYC(account common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.AddWhitelistedKYC(&_HLBICO.TransactOpts, account)
}

// AddWhitelistedKYC is a paid mutator transaction binding the contract method 0x7c8b38c8.
//
// Solidity: function addWhitelistedKYC(address account) returns()
func (_HLBICO *HLBICOTransactorSession) AddWhitelistedKYC(account common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.AddWhitelistedKYC(&_HLBICO.TransactOpts, account)
}

// AddWhitelistedRegistered is a paid mutator transaction binding the contract method 0xfc400611.
//
// Solidity: function addWhitelistedRegistered(address account) returns()
func (_HLBICO *HLBICOTransactor) AddWhitelistedRegistered(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "addWhitelistedRegistered", account)
}

// AddWhitelistedRegistered is a paid mutator transaction binding the contract method 0xfc400611.
//
// Solidity: function addWhitelistedRegistered(address account) returns()
func (_HLBICO *HLBICOSession) AddWhitelistedRegistered(account common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.AddWhitelistedRegistered(&_HLBICO.TransactOpts, account)
}

// AddWhitelistedRegistered is a paid mutator transaction binding the contract method 0xfc400611.
//
// Solidity: function addWhitelistedRegistered(address account) returns()
func (_HLBICO *HLBICOTransactorSession) AddWhitelistedRegistered(account common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.AddWhitelistedRegistered(&_HLBICO.TransactOpts, account)
}

// AdjustEtherValue is a paid mutator transaction binding the contract method 0x329425c5.
//
// Solidity: function adjustEtherValue(uint256 coef) returns()
func (_HLBICO *HLBICOTransactor) AdjustEtherValue(opts *bind.TransactOpts, coef *big.Int) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "adjustEtherValue", coef)
}

// AdjustEtherValue is a paid mutator transaction binding the contract method 0x329425c5.
//
// Solidity: function adjustEtherValue(uint256 coef) returns()
func (_HLBICO *HLBICOSession) AdjustEtherValue(coef *big.Int) (*types.Transaction, error) {
	return _HLBICO.Contract.AdjustEtherValue(&_HLBICO.TransactOpts, coef)
}

// AdjustEtherValue is a paid mutator transaction binding the contract method 0x329425c5.
//
// Solidity: function adjustEtherValue(uint256 coef) returns()
func (_HLBICO *HLBICOTransactorSession) AdjustEtherValue(coef *big.Int) (*types.Transaction, error) {
	return _HLBICO.Contract.AdjustEtherValue(&_HLBICO.TransactOpts, coef)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_HLBICO *HLBICOTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_HLBICO *HLBICOSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.BuyTokens(&_HLBICO.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_HLBICO *HLBICOTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.BuyTokens(&_HLBICO.TransactOpts, beneficiary)
}

// ChangeReserveAddress is a paid mutator transaction binding the contract method 0x477fea02.
//
// Solidity: function changeReserveAddress(address newReserveAddress) returns()
func (_HLBICO *HLBICOTransactor) ChangeReserveAddress(opts *bind.TransactOpts, newReserveAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "changeReserveAddress", newReserveAddress)
}

// ChangeReserveAddress is a paid mutator transaction binding the contract method 0x477fea02.
//
// Solidity: function changeReserveAddress(address newReserveAddress) returns()
func (_HLBICO *HLBICOSession) ChangeReserveAddress(newReserveAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ChangeReserveAddress(&_HLBICO.TransactOpts, newReserveAddress)
}

// ChangeReserveAddress is a paid mutator transaction binding the contract method 0x477fea02.
//
// Solidity: function changeReserveAddress(address newReserveAddress) returns()
func (_HLBICO *HLBICOTransactorSession) ChangeReserveAddress(newReserveAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ChangeReserveAddress(&_HLBICO.TransactOpts, newReserveAddress)
}

// ChangeToken is a paid mutator transaction binding the contract method 0x66829b16.
//
// Solidity: function changeToken(address newToken) returns()
func (_HLBICO *HLBICOTransactor) ChangeToken(opts *bind.TransactOpts, newToken common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "changeToken", newToken)
}

// ChangeToken is a paid mutator transaction binding the contract method 0x66829b16.
//
// Solidity: function changeToken(address newToken) returns()
func (_HLBICO *HLBICOSession) ChangeToken(newToken common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ChangeToken(&_HLBICO.TransactOpts, newToken)
}

// ChangeToken is a paid mutator transaction binding the contract method 0x66829b16.
//
// Solidity: function changeToken(address newToken) returns()
func (_HLBICO *HLBICOTransactorSession) ChangeToken(newToken common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ChangeToken(&_HLBICO.TransactOpts, newToken)
}

// ChangeWhitelister is a paid mutator transaction binding the contract method 0x966aeece.
//
// Solidity: function changeWhitelister(address newWhitelisterAddress) returns()
func (_HLBICO *HLBICOTransactor) ChangeWhitelister(opts *bind.TransactOpts, newWhitelisterAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "changeWhitelister", newWhitelisterAddress)
}

// ChangeWhitelister is a paid mutator transaction binding the contract method 0x966aeece.
//
// Solidity: function changeWhitelister(address newWhitelisterAddress) returns()
func (_HLBICO *HLBICOSession) ChangeWhitelister(newWhitelisterAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ChangeWhitelister(&_HLBICO.TransactOpts, newWhitelisterAddress)
}

// ChangeWhitelister is a paid mutator transaction binding the contract method 0x966aeece.
//
// Solidity: function changeWhitelister(address newWhitelisterAddress) returns()
func (_HLBICO *HLBICOTransactorSession) ChangeWhitelister(newWhitelisterAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ChangeWhitelister(&_HLBICO.TransactOpts, newWhitelisterAddress)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_HLBICO *HLBICOTransactor) ClaimRefund(opts *bind.TransactOpts, refundee common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "claimRefund", refundee)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_HLBICO *HLBICOSession) ClaimRefund(refundee common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ClaimRefund(&_HLBICO.TransactOpts, refundee)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_HLBICO *HLBICOTransactorSession) ClaimRefund(refundee common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.ClaimRefund(&_HLBICO.TransactOpts, refundee)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_HLBICO *HLBICOTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_HLBICO *HLBICOSession) Finalize() (*types.Transaction, error) {
	return _HLBICO.Contract.Finalize(&_HLBICO.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_HLBICO *HLBICOTransactorSession) Finalize() (*types.Transaction, error) {
	return _HLBICO.Contract.Finalize(&_HLBICO.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address whitelistingAddress, address reserveAddress) returns()
func (_HLBICO *HLBICOTransactor) Init(opts *bind.TransactOpts, whitelistingAddress common.Address, reserveAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "init", whitelistingAddress, reserveAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address whitelistingAddress, address reserveAddress) returns()
func (_HLBICO *HLBICOSession) Init(whitelistingAddress common.Address, reserveAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.Init(&_HLBICO.TransactOpts, whitelistingAddress, reserveAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address whitelistingAddress, address reserveAddress) returns()
func (_HLBICO *HLBICOTransactorSession) Init(whitelistingAddress common.Address, reserveAddress common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.Init(&_HLBICO.TransactOpts, whitelistingAddress, reserveAddress)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_HLBICO *HLBICOTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_HLBICO *HLBICOSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.RemoveWhitelisted(&_HLBICO.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_HLBICO *HLBICOTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.RemoveWhitelisted(&_HLBICO.TransactOpts, account)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_HLBICO *HLBICOTransactor) WithdrawTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _HLBICO.contract.Transact(opts, "withdrawTokens", beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_HLBICO *HLBICOSession) WithdrawTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.WithdrawTokens(&_HLBICO.TransactOpts, beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_HLBICO *HLBICOTransactorSession) WithdrawTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _HLBICO.Contract.WithdrawTokens(&_HLBICO.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HLBICO *HLBICOTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _HLBICO.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HLBICO *HLBICOSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HLBICO.Contract.Fallback(&_HLBICO.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HLBICO *HLBICOTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HLBICO.Contract.Fallback(&_HLBICO.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HLBICO *HLBICOTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HLBICO.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HLBICO *HLBICOSession) Receive() (*types.Transaction, error) {
	return _HLBICO.Contract.Receive(&_HLBICO.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HLBICO *HLBICOTransactorSession) Receive() (*types.Transaction, error) {
	return _HLBICO.Contract.Receive(&_HLBICO.TransactOpts)
}

// HLBICOChangedReserveAddressIterator is returned from FilterChangedReserveAddress and is used to iterate over the raw logs and unpacked data for ChangedReserveAddress events raised by the HLBICO contract.
type HLBICOChangedReserveAddressIterator struct {
	Event *HLBICOChangedReserveAddress // Event containing the contract specifics and raw log

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
func (it *HLBICOChangedReserveAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOChangedReserveAddress)
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
		it.Event = new(HLBICOChangedReserveAddress)
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
func (it *HLBICOChangedReserveAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOChangedReserveAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOChangedReserveAddress represents a ChangedReserveAddress event raised by the HLBICO contract.
type HLBICOChangedReserveAddress struct {
	ReserveAddress common.Address
	ChangerAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedReserveAddress is a free log retrieval operation binding the contract event 0x7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b54.
//
// Solidity: event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress)
func (_HLBICO *HLBICOFilterer) FilterChangedReserveAddress(opts *bind.FilterOpts, reserveAddress []common.Address, changerAddress []common.Address) (*HLBICOChangedReserveAddressIterator, error) {

	var reserveAddressRule []interface{}
	for _, reserveAddressItem := range reserveAddress {
		reserveAddressRule = append(reserveAddressRule, reserveAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "ChangedReserveAddress", reserveAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HLBICOChangedReserveAddressIterator{contract: _HLBICO.contract, event: "ChangedReserveAddress", logs: logs, sub: sub}, nil
}

// WatchChangedReserveAddress is a free log subscription operation binding the contract event 0x7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b54.
//
// Solidity: event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress)
func (_HLBICO *HLBICOFilterer) WatchChangedReserveAddress(opts *bind.WatchOpts, sink chan<- *HLBICOChangedReserveAddress, reserveAddress []common.Address, changerAddress []common.Address) (event.Subscription, error) {

	var reserveAddressRule []interface{}
	for _, reserveAddressItem := range reserveAddress {
		reserveAddressRule = append(reserveAddressRule, reserveAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "ChangedReserveAddress", reserveAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOChangedReserveAddress)
				if err := _HLBICO.contract.UnpackLog(event, "ChangedReserveAddress", log); err != nil {
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

// ParseChangedReserveAddress is a log parse operation binding the contract event 0x7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b54.
//
// Solidity: event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress)
func (_HLBICO *HLBICOFilterer) ParseChangedReserveAddress(log types.Log) (*HLBICOChangedReserveAddress, error) {
	event := new(HLBICOChangedReserveAddress)
	if err := _HLBICO.contract.UnpackLog(event, "ChangedReserveAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOChangedWhitelisterAddressIterator is returned from FilterChangedWhitelisterAddress and is used to iterate over the raw logs and unpacked data for ChangedWhitelisterAddress events raised by the HLBICO contract.
type HLBICOChangedWhitelisterAddressIterator struct {
	Event *HLBICOChangedWhitelisterAddress // Event containing the contract specifics and raw log

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
func (it *HLBICOChangedWhitelisterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOChangedWhitelisterAddress)
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
		it.Event = new(HLBICOChangedWhitelisterAddress)
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
func (it *HLBICOChangedWhitelisterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOChangedWhitelisterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOChangedWhitelisterAddress represents a ChangedWhitelisterAddress event raised by the HLBICO contract.
type HLBICOChangedWhitelisterAddress struct {
	WhitelisterAddress common.Address
	ChangerAddress     common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangedWhitelisterAddress is a free log retrieval operation binding the contract event 0x63c70fa8d3aed6b9cda2842d4152a4ccf78aab6db236b9b22ad898ca7541e113.
//
// Solidity: event ChangedWhitelisterAddress(address indexed whitelisterAddress, address indexed changerAddress)
func (_HLBICO *HLBICOFilterer) FilterChangedWhitelisterAddress(opts *bind.FilterOpts, whitelisterAddress []common.Address, changerAddress []common.Address) (*HLBICOChangedWhitelisterAddressIterator, error) {

	var whitelisterAddressRule []interface{}
	for _, whitelisterAddressItem := range whitelisterAddress {
		whitelisterAddressRule = append(whitelisterAddressRule, whitelisterAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "ChangedWhitelisterAddress", whitelisterAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HLBICOChangedWhitelisterAddressIterator{contract: _HLBICO.contract, event: "ChangedWhitelisterAddress", logs: logs, sub: sub}, nil
}

// WatchChangedWhitelisterAddress is a free log subscription operation binding the contract event 0x63c70fa8d3aed6b9cda2842d4152a4ccf78aab6db236b9b22ad898ca7541e113.
//
// Solidity: event ChangedWhitelisterAddress(address indexed whitelisterAddress, address indexed changerAddress)
func (_HLBICO *HLBICOFilterer) WatchChangedWhitelisterAddress(opts *bind.WatchOpts, sink chan<- *HLBICOChangedWhitelisterAddress, whitelisterAddress []common.Address, changerAddress []common.Address) (event.Subscription, error) {

	var whitelisterAddressRule []interface{}
	for _, whitelisterAddressItem := range whitelisterAddress {
		whitelisterAddressRule = append(whitelisterAddressRule, whitelisterAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "ChangedWhitelisterAddress", whitelisterAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOChangedWhitelisterAddress)
				if err := _HLBICO.contract.UnpackLog(event, "ChangedWhitelisterAddress", log); err != nil {
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

// ParseChangedWhitelisterAddress is a log parse operation binding the contract event 0x63c70fa8d3aed6b9cda2842d4152a4ccf78aab6db236b9b22ad898ca7541e113.
//
// Solidity: event ChangedWhitelisterAddress(address indexed whitelisterAddress, address indexed changerAddress)
func (_HLBICO *HLBICOFilterer) ParseChangedWhitelisterAddress(log types.Log) (*HLBICOChangedWhitelisterAddress, error) {
	event := new(HLBICOChangedWhitelisterAddress)
	if err := _HLBICO.contract.UnpackLog(event, "ChangedWhitelisterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOCrowdsaleFinalizedIterator is returned from FilterCrowdsaleFinalized and is used to iterate over the raw logs and unpacked data for CrowdsaleFinalized events raised by the HLBICO contract.
type HLBICOCrowdsaleFinalizedIterator struct {
	Event *HLBICOCrowdsaleFinalized // Event containing the contract specifics and raw log

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
func (it *HLBICOCrowdsaleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOCrowdsaleFinalized)
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
		it.Event = new(HLBICOCrowdsaleFinalized)
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
func (it *HLBICOCrowdsaleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOCrowdsaleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOCrowdsaleFinalized represents a CrowdsaleFinalized event raised by the HLBICO contract.
type HLBICOCrowdsaleFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrowdsaleFinalized is a free log retrieval operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_HLBICO *HLBICOFilterer) FilterCrowdsaleFinalized(opts *bind.FilterOpts) (*HLBICOCrowdsaleFinalizedIterator, error) {

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return &HLBICOCrowdsaleFinalizedIterator{contract: _HLBICO.contract, event: "CrowdsaleFinalized", logs: logs, sub: sub}, nil
}

// WatchCrowdsaleFinalized is a free log subscription operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_HLBICO *HLBICOFilterer) WatchCrowdsaleFinalized(opts *bind.WatchOpts, sink chan<- *HLBICOCrowdsaleFinalized) (event.Subscription, error) {

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOCrowdsaleFinalized)
				if err := _HLBICO.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
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

// ParseCrowdsaleFinalized is a log parse operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_HLBICO *HLBICOFilterer) ParseCrowdsaleFinalized(log types.Log) (*HLBICOCrowdsaleFinalized, error) {
	event := new(HLBICOCrowdsaleFinalized)
	if err := _HLBICO.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOInitializedContractIterator is returned from FilterInitializedContract and is used to iterate over the raw logs and unpacked data for InitializedContract events raised by the HLBICO contract.
type HLBICOInitializedContractIterator struct {
	Event *HLBICOInitializedContract // Event containing the contract specifics and raw log

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
func (it *HLBICOInitializedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOInitializedContract)
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
		it.Event = new(HLBICOInitializedContract)
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
func (it *HLBICOInitializedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOInitializedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOInitializedContract represents a InitializedContract event raised by the HLBICO contract.
type HLBICOInitializedContract struct {
	ChangerAddress      common.Address
	WhitelistingAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInitializedContract is a free log retrieval operation binding the contract event 0x1136ccf874c0c85bed8c1488e35920195c98fda6c939c473a46510d35826d5a1.
//
// Solidity: event InitializedContract(address indexed changerAddress, address indexed whitelistingAddress)
func (_HLBICO *HLBICOFilterer) FilterInitializedContract(opts *bind.FilterOpts, changerAddress []common.Address, whitelistingAddress []common.Address) (*HLBICOInitializedContractIterator, error) {

	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}
	var whitelistingAddressRule []interface{}
	for _, whitelistingAddressItem := range whitelistingAddress {
		whitelistingAddressRule = append(whitelistingAddressRule, whitelistingAddressItem)
	}

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "InitializedContract", changerAddressRule, whitelistingAddressRule)
	if err != nil {
		return nil, err
	}
	return &HLBICOInitializedContractIterator{contract: _HLBICO.contract, event: "InitializedContract", logs: logs, sub: sub}, nil
}

// WatchInitializedContract is a free log subscription operation binding the contract event 0x1136ccf874c0c85bed8c1488e35920195c98fda6c939c473a46510d35826d5a1.
//
// Solidity: event InitializedContract(address indexed changerAddress, address indexed whitelistingAddress)
func (_HLBICO *HLBICOFilterer) WatchInitializedContract(opts *bind.WatchOpts, sink chan<- *HLBICOInitializedContract, changerAddress []common.Address, whitelistingAddress []common.Address) (event.Subscription, error) {

	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}
	var whitelistingAddressRule []interface{}
	for _, whitelistingAddressItem := range whitelistingAddress {
		whitelistingAddressRule = append(whitelistingAddressRule, whitelistingAddressItem)
	}

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "InitializedContract", changerAddressRule, whitelistingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOInitializedContract)
				if err := _HLBICO.contract.UnpackLog(event, "InitializedContract", log); err != nil {
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

// ParseInitializedContract is a log parse operation binding the contract event 0x1136ccf874c0c85bed8c1488e35920195c98fda6c939c473a46510d35826d5a1.
//
// Solidity: event InitializedContract(address indexed changerAddress, address indexed whitelistingAddress)
func (_HLBICO *HLBICOFilterer) ParseInitializedContract(log types.Log) (*HLBICOInitializedContract, error) {
	event := new(HLBICOInitializedContract)
	if err := _HLBICO.contract.UnpackLog(event, "InitializedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the HLBICO contract.
type HLBICOTimedCrowdsaleExtendedIterator struct {
	Event *HLBICOTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *HLBICOTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOTimedCrowdsaleExtended)
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
		it.Event = new(HLBICOTimedCrowdsaleExtended)
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
func (it *HLBICOTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the HLBICO contract.
type HLBICOTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_HLBICO *HLBICOFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*HLBICOTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &HLBICOTimedCrowdsaleExtendedIterator{contract: _HLBICO.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_HLBICO *HLBICOFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *HLBICOTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOTimedCrowdsaleExtended)
				if err := _HLBICO.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_HLBICO *HLBICOFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*HLBICOTimedCrowdsaleExtended, error) {
	event := new(HLBICOTimedCrowdsaleExtended)
	if err := _HLBICO.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the HLBICO contract.
type HLBICOTokensPurchasedIterator struct {
	Event *HLBICOTokensPurchased // Event containing the contract specifics and raw log

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
func (it *HLBICOTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOTokensPurchased)
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
		it.Event = new(HLBICOTokensPurchased)
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
func (it *HLBICOTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOTokensPurchased represents a TokensPurchased event raised by the HLBICO contract.
type HLBICOTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_HLBICO *HLBICOFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*HLBICOTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &HLBICOTokensPurchasedIterator{contract: _HLBICO.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_HLBICO *HLBICOFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *HLBICOTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOTokensPurchased)
				if err := _HLBICO.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_HLBICO *HLBICOFilterer) ParseTokensPurchased(log types.Log) (*HLBICOTokensPurchased, error) {
	event := new(HLBICOTokensPurchased)
	if err := _HLBICO.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOUpdatedCapsIterator is returned from FilterUpdatedCaps and is used to iterate over the raw logs and unpacked data for UpdatedCaps events raised by the HLBICO contract.
type HLBICOUpdatedCapsIterator struct {
	Event *HLBICOUpdatedCaps // Event containing the contract specifics and raw log

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
func (it *HLBICOUpdatedCapsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOUpdatedCaps)
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
		it.Event = new(HLBICOUpdatedCaps)
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
func (it *HLBICOUpdatedCapsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOUpdatedCapsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOUpdatedCaps represents a UpdatedCaps event raised by the HLBICO contract.
type HLBICOUpdatedCaps struct {
	NewGoal      *big.Int
	NewCap       *big.Int
	NewTranche   *big.Int
	NewMaxInvest *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCaps is a free log retrieval operation binding the contract event 0x0c8ac593d921e6388dbbe810ac16aa0fa330ca2fae1ff2b027e6b8c2bb0f9242.
//
// Solidity: event UpdatedCaps(uint256 newGoal, uint256 newCap, uint256 newTranche, uint256 newMaxInvest)
func (_HLBICO *HLBICOFilterer) FilterUpdatedCaps(opts *bind.FilterOpts) (*HLBICOUpdatedCapsIterator, error) {

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "UpdatedCaps")
	if err != nil {
		return nil, err
	}
	return &HLBICOUpdatedCapsIterator{contract: _HLBICO.contract, event: "UpdatedCaps", logs: logs, sub: sub}, nil
}

// WatchUpdatedCaps is a free log subscription operation binding the contract event 0x0c8ac593d921e6388dbbe810ac16aa0fa330ca2fae1ff2b027e6b8c2bb0f9242.
//
// Solidity: event UpdatedCaps(uint256 newGoal, uint256 newCap, uint256 newTranche, uint256 newMaxInvest)
func (_HLBICO *HLBICOFilterer) WatchUpdatedCaps(opts *bind.WatchOpts, sink chan<- *HLBICOUpdatedCaps) (event.Subscription, error) {

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "UpdatedCaps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOUpdatedCaps)
				if err := _HLBICO.contract.UnpackLog(event, "UpdatedCaps", log); err != nil {
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

// ParseUpdatedCaps is a log parse operation binding the contract event 0x0c8ac593d921e6388dbbe810ac16aa0fa330ca2fae1ff2b027e6b8c2bb0f9242.
//
// Solidity: event UpdatedCaps(uint256 newGoal, uint256 newCap, uint256 newTranche, uint256 newMaxInvest)
func (_HLBICO *HLBICOFilterer) ParseUpdatedCaps(log types.Log) (*HLBICOUpdatedCaps, error) {
	event := new(HLBICOUpdatedCaps)
	if err := _HLBICO.contract.UnpackLog(event, "UpdatedCaps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the HLBICO contract.
type HLBICOWhitelistedAddedIterator struct {
	Event *HLBICOWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *HLBICOWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOWhitelistedAdded)
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
		it.Event = new(HLBICOWhitelistedAdded)
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
func (it *HLBICOWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOWhitelistedAdded represents a WhitelistedAdded event raised by the HLBICO contract.
type HLBICOWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address indexed account)
func (_HLBICO *HLBICOFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts, account []common.Address) (*HLBICOWhitelistedAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &HLBICOWhitelistedAddedIterator{contract: _HLBICO.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address indexed account)
func (_HLBICO *HLBICOFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *HLBICOWhitelistedAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOWhitelistedAdded)
				if err := _HLBICO.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address indexed account)
func (_HLBICO *HLBICOFilterer) ParseWhitelistedAdded(log types.Log) (*HLBICOWhitelistedAdded, error) {
	event := new(HLBICOWhitelistedAdded)
	if err := _HLBICO.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HLBICOWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the HLBICO contract.
type HLBICOWhitelistedRemovedIterator struct {
	Event *HLBICOWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *HLBICOWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HLBICOWhitelistedRemoved)
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
		it.Event = new(HLBICOWhitelistedRemoved)
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
func (it *HLBICOWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HLBICOWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HLBICOWhitelistedRemoved represents a WhitelistedRemoved event raised by the HLBICO contract.
type HLBICOWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed account)
func (_HLBICO *HLBICOFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts, account []common.Address) (*HLBICOWhitelistedRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HLBICO.contract.FilterLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &HLBICOWhitelistedRemovedIterator{contract: _HLBICO.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed account)
func (_HLBICO *HLBICOFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *HLBICOWhitelistedRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HLBICO.contract.WatchLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HLBICOWhitelistedRemoved)
				if err := _HLBICO.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed account)
func (_HLBICO *HLBICOFilterer) ParseWhitelistedRemoved(log types.Log) (*HLBICOWhitelistedRemoved, error) {
	event := new(HLBICOWhitelistedRemoved)
	if err := _HLBICO.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenABI is the input ABI used to generate the binding from.
const LBCTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minterAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changerAddress\",\"type\":\"address\"}],\"name\":\"ChangedMinterAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauserAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changerAddress\",\"type\":\"address\"}],\"name\":\"ChangedPauserAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changerAddress\",\"type\":\"address\"}],\"name\":\"ChangedReserveAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"}],\"name\":\"InitializedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_deployingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_pauserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_reserveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinterAddress\",\"type\":\"address\"}],\"name\":\"changeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauserAddress\",\"type\":\"address\"}],\"name\":\"changePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauserAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LBCTokenFuncSigs maps the 4-byte function signature to its string representation.
var LBCTokenFuncSigs = map[string]string{
	"1cd273cf": "_deployingAddress()",
	"9c495671": "_minterAddress()",
	"715df33b": "_pauserAddress()",
	"ae5b5b11": "_reserveAddress()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"355274ea": "cap()",
	"2c4d4d18": "changeMinter(address)",
	"2cd271e7": "changePauser(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"184b9559": "init(address,address,address)",
	"158ef93e": "initialized()",
	"40c10f19": "mint(address,uint256)",
	"06fdde03": "name()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"3f4ba83a": "unpause()",
}

// LBCTokenBin is the compiled bytecode used for deploying new contracts.
var LBCTokenBin = "0x60806040523480156200001157600080fd5b506040516200186938038062001869833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b506040525050506af8277896582678ac00000082828160039080519060200190620001c592919062000267565b508051620001db90600490602084019062000267565b50506005805461ff001960ff19909116601217169055508062000245576040805162461bcd60e51b815260206004820152601560248201527f45524332304361707065643a2063617020697320300000000000000000000000604482015290519081900360640190fd5b600655505060078054610100600160a81b031916336101000217905562000313565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200029f5760008555620002ea565b82601f10620002ba57805160ff1916838001178555620002ea565b82800160010185558215620002ea579182015b82811115620002ea578251825591602001919060010190620002cd565b50620002f8929150620002fc565b5090565b5b80821115620002f85760008155600101620002fd565b61154680620003236000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80633f4ba83a116100c357806395d89b411161007c57806395d89b41146103e05780639c495671146103e8578063a457c2d7146103f0578063a9059cbb1461041c578063ae5b5b1114610448578063dd62ed3e1461045057610158565b80633f4ba83a1461036e57806340c10f19146103765780635c975abb146103a257806370a08231146103aa578063715df33b146103d05780638456cb59146103d857610158565b806323b872dd1161011557806323b872dd1461029a5780632c4d4d18146102d05780632cd271e7146102f6578063313ce5671461031c578063355274ea1461033a578063395093511461034257610158565b806306fdde031461015d578063095ea7b3146101da578063158ef93e1461021a57806318160ddd14610222578063184b95591461023c5780631cd273cf14610276575b600080fd5b61016561047e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561019f578181015183820152602001610187565b50505050905090810190601f1680156101cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610206600480360360408110156101f057600080fd5b506001600160a01b038135169060200135610514565b604080519115158252519081900360200190f35b610206610531565b61022a61053a565b60408051918252519081900360200190f35b6102746004803603606081101561025257600080fd5b506001600160a01b038135811691602081013582169160409091013516610540565b005b61027e61076d565b604080516001600160a01b039092168252519081900360200190f35b610206600480360360608110156102b057600080fd5b506001600160a01b03813581169160208101359091169060400135610781565b610274600480360360208110156102e657600080fd5b50356001600160a01b0316610808565b6102746004803603602081101561030c57600080fd5b50356001600160a01b0316610906565b610324610a04565b6040805160ff9092168252519081900360200190f35b61022a610a0d565b6102066004803603604081101561035857600080fd5b506001600160a01b038135169060200135610a13565b610274610a61565b6102066004803603604081101561038c57600080fd5b506001600160a01b038135169060200135610ab4565b610206610b0a565b61022a600480360360208110156103c057600080fd5b50356001600160a01b0316610b18565b61027e610b33565b610274610b42565b610165610b93565b61027e610bf4565b6102066004803603604081101561040657600080fd5b506001600160a01b038135169060200135610c03565b6102066004803603604081101561043257600080fd5b506001600160a01b038135169060200135610c6b565b61027e610c7f565b61022a6004803603604081101561046657600080fd5b506001600160a01b0381358116916020013516610c8e565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561050a5780601f106104df5761010080835404028352916020019161050a565b820191906000526020600020905b8154815290600101906020018083116104ed57829003601f168201915b5050505050905090565b6000610528610521610cb9565b8484610cbd565b50600192915050565b60075460ff1681565b60025490565b60075460ff1615610598576040805162461bcd60e51b815260206004820181905260248201527f436f6e747261637420697320616c726561647920696e697469616c697a65642e604482015290519081900360640190fd5b60075461010090046001600160a01b031633146105e65760405162461bcd60e51b815260040180806020018281038252603081526020018061130f6030913960400191505060405180910390fd5b6001600160a01b038316610641576040805162461bcd60e51b815260206004820152601b60248201527f5f6d696e746572416464726573732063616e6e6f742062652030780000000000604482015290519081900360640190fd5b6001600160a01b03821661069c576040805162461bcd60e51b815260206004820152601b60248201527f5f706175736572416464726573732063616e6e6f742062652030780000000000604482015290519081900360640190fd5b6001600160a01b0381166106f7576040805162461bcd60e51b815260206004820152601c60248201527f5f72657365727665416464726573732063616e6e6f7420626520307800000000604482015290519081900360640190fd5b600980546001600160a01b038086166001600160a01b03199283161790925560088054858416908316179055600a805492841692909116821790556007805460ff191660011790556040517faf827135deb94e19593430f16f577c1d98befd728375ca86115192ab05848fcb90600090a2505050565b60075461010090046001600160a01b031681565b600061078e848484610da9565b6107fe8461079a610cb9565b6107f9856040518060600160405280603281526020016114df603291396001600160a01b038a166000908152600160205260408120906107d8610cb9565b6001600160a01b031681526020810191909152604001600020549190610f04565b610cbd565b5060019392505050565b6009546001600160a01b031633146108515760405162461bcd60e51b815260040180806020018281038252602d81526020018061139c602d913960400191505060405180910390fd5b600554610100900460ff16156108a1576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b600980546001600160a01b0319166001600160a01b0383161790556108c4610cb9565b6001600160a01b0316816001600160a01b03167fc6cc4544f0cd8f11ff892e5ea4c8609793a185ba6f5c7ddd2a1f5399bee6774c60405160405180910390a350565b6008546001600160a01b0316331461094f5760405162461bcd60e51b815260040180806020018281038252602d8152602001806113c9602d913960400191505060405180910390fd5b600554610100900460ff161561099f576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0383161790556109c2610cb9565b6001600160a01b0316816001600160a01b03167f2f9de1084416a5a11e235865f4bbfd5c7c9244354207963ce0df47d2ec6193b860405160405180910390a350565b60055460ff1690565b60065490565b6000610528610a20610cb9565b846107f98560016000610a31610cb9565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610f9b565b6008546001600160a01b03163314610aaa5760405162461bcd60e51b815260040180806020018281038252602d8152602001806113c9602d913960400191505060405180910390fd5b610ab2610ffc565b565b6009546000906001600160a01b03163314610b005760405162461bcd60e51b815260040180806020018281038252602d81526020018061139c602d913960400191505060405180910390fd5b61052883836110a0565b600554610100900460ff1690565b6001600160a01b031660009081526020819052604090205490565b6008546001600160a01b031681565b6008546001600160a01b03163314610b8b5760405162461bcd60e51b815260040180806020018281038252602d8152602001806113c9602d913960400191505060405180910390fd5b610ab261117a565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561050a5780601f106104df5761010080835404028352916020019161050a565b6009546001600160a01b031681565b6000610528610c10610cb9565b846107f9856040518060600160405280602f81526020016114b0602f913960016000610c3a610cb9565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190610f04565b6000610528610c78610cb9565b8484610da9565b600a546001600160a01b031681565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b038316610d025760405162461bcd60e51b815260040180806020018281038252602e81526020018061133f602e913960400191505060405180910390fd5b6001600160a01b038216610d475760405162461bcd60e51b815260040180806020018281038252602c8152602001806112e3602c913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610dee5760405162461bcd60e51b815260040180806020018281038252602f81526020018061136d602f913960400191505060405180910390fd5b6001600160a01b038216610e335760405162461bcd60e51b815260040180806020018281038252602d81526020018061142a602d913960400191505060405180910390fd5b610e3e838383611202565b610e7b81604051806060016040528060308152602001611480603091396001600160a01b0386166000908152602081905260409020549190610f04565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610eaa9082610f9b565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610f935760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f58578181015183820152602001610f40565b50505050905090810190601f168015610f855780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610ff5576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600554610100900460ff1661104f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6005805461ff00191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611083610cb9565b604080516001600160a01b039092168252519081900360200190a1565b6001600160a01b0382166110e55760405162461bcd60e51b81526004018080602001828103825260298152602001806114576029913960400191505060405180910390fd5b6110f160008383611202565b6002546110fe9082610f9b565b6002556001600160a01b0382166000908152602081905260409020546111249082610f9b565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600554610100900460ff16156111ca576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6005805461ff0019166101001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611083610cb9565b61120d838383611212565b505050565b61121d838383611293565b6001600160a01b03831661120d576006546112408261123a61053a565b90610f9b565b111561120d576040805162461bcd60e51b815260206004820152601960248201527f45524332304361707065643a2063617020657863656564656400000000000000604482015290519081900360640190fd5b61129e83838361120d565b6112a6610b0a565b1561120d5760405162461bcd60e51b81526004018080602001828103825260348152602001806113f66034913960400191505060405180910390fdfe4552433230556e6275726e61626c653a20617070726f766520746f20746865207a65726f20616464726573734f6e6c7920746865206465706c6f79696e6720616464726573732063616e2063616c6c2074686973206d6574686f642e4552433230556e6275726e61626c653a20617070726f76652066726f6d20746865207a65726f20616464726573734552433230556e6275726e61626c653a207472616e736665722066726f6d20746865207a65726f20616464726573734f6e6c7920746865206d696e74657220616464726573732063616e2063616c6c2074686973206d6574686f642e4f6e6c79207468652070617573657220616464726573732063616e2063616c6c2074686973206d6574686f642e45524332305061757361626c65556e6275726e61626c653a20746f6b656e207472616e73666572207768696c65207061757365644552433230556e6275726e61626c653a207472616e7366657220746f20746865207a65726f20616464726573734552433230556e6275726e61626c653a206d696e7420746f20746865207a65726f20616464726573734552433230556e6275726e61626c653a207472616e7366657220616d6f756e7420657863656564732062616c616e63654552433230556e6275726e61626c653a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f4552433230556e6275726e61626c653a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365a264697066735822122063c16f839510aa6e12bc2353259160ac5a72aad3582dfb0b5a2d861abbbb2f5364736f6c63430007060033"

// DeployLBCToken deploys a new Ethereum contract, binding an instance of LBCToken to it.
func DeployLBCToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *LBCToken, error) {
	parsed, err := abi.JSON(strings.NewReader(LBCTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LBCTokenBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LBCToken{LBCTokenCaller: LBCTokenCaller{contract: contract}, LBCTokenTransactor: LBCTokenTransactor{contract: contract}, LBCTokenFilterer: LBCTokenFilterer{contract: contract}}, nil
}

// LBCToken is an auto generated Go binding around an Ethereum contract.
type LBCToken struct {
	LBCTokenCaller     // Read-only binding to the contract
	LBCTokenTransactor // Write-only binding to the contract
	LBCTokenFilterer   // Log filterer for contract events
}

// LBCTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LBCTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBCTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LBCTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBCTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LBCTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LBCTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LBCTokenSession struct {
	Contract     *LBCToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LBCTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LBCTokenCallerSession struct {
	Contract *LBCTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LBCTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LBCTokenTransactorSession struct {
	Contract     *LBCTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LBCTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LBCTokenRaw struct {
	Contract *LBCToken // Generic contract binding to access the raw methods on
}

// LBCTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LBCTokenCallerRaw struct {
	Contract *LBCTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LBCTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LBCTokenTransactorRaw struct {
	Contract *LBCTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLBCToken creates a new instance of LBCToken, bound to a specific deployed contract.
func NewLBCToken(address common.Address, backend bind.ContractBackend) (*LBCToken, error) {
	contract, err := bindLBCToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LBCToken{LBCTokenCaller: LBCTokenCaller{contract: contract}, LBCTokenTransactor: LBCTokenTransactor{contract: contract}, LBCTokenFilterer: LBCTokenFilterer{contract: contract}}, nil
}

// NewLBCTokenCaller creates a new read-only instance of LBCToken, bound to a specific deployed contract.
func NewLBCTokenCaller(address common.Address, caller bind.ContractCaller) (*LBCTokenCaller, error) {
	contract, err := bindLBCToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LBCTokenCaller{contract: contract}, nil
}

// NewLBCTokenTransactor creates a new write-only instance of LBCToken, bound to a specific deployed contract.
func NewLBCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LBCTokenTransactor, error) {
	contract, err := bindLBCToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LBCTokenTransactor{contract: contract}, nil
}

// NewLBCTokenFilterer creates a new log filterer instance of LBCToken, bound to a specific deployed contract.
func NewLBCTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LBCTokenFilterer, error) {
	contract, err := bindLBCToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LBCTokenFilterer{contract: contract}, nil
}

// bindLBCToken binds a generic wrapper to an already deployed contract.
func bindLBCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LBCTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LBCToken *LBCTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LBCToken.Contract.LBCTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LBCToken *LBCTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBCToken.Contract.LBCTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LBCToken *LBCTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LBCToken.Contract.LBCTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LBCToken *LBCTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LBCToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LBCToken *LBCTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBCToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LBCToken *LBCTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LBCToken.Contract.contract.Transact(opts, method, params...)
}

// DeployingAddress is a free data retrieval call binding the contract method 0x1cd273cf.
//
// Solidity: function _deployingAddress() view returns(address)
func (_LBCToken *LBCTokenCaller) DeployingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "_deployingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployingAddress is a free data retrieval call binding the contract method 0x1cd273cf.
//
// Solidity: function _deployingAddress() view returns(address)
func (_LBCToken *LBCTokenSession) DeployingAddress() (common.Address, error) {
	return _LBCToken.Contract.DeployingAddress(&_LBCToken.CallOpts)
}

// DeployingAddress is a free data retrieval call binding the contract method 0x1cd273cf.
//
// Solidity: function _deployingAddress() view returns(address)
func (_LBCToken *LBCTokenCallerSession) DeployingAddress() (common.Address, error) {
	return _LBCToken.Contract.DeployingAddress(&_LBCToken.CallOpts)
}

// MinterAddress is a free data retrieval call binding the contract method 0x9c495671.
//
// Solidity: function _minterAddress() view returns(address)
func (_LBCToken *LBCTokenCaller) MinterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "_minterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MinterAddress is a free data retrieval call binding the contract method 0x9c495671.
//
// Solidity: function _minterAddress() view returns(address)
func (_LBCToken *LBCTokenSession) MinterAddress() (common.Address, error) {
	return _LBCToken.Contract.MinterAddress(&_LBCToken.CallOpts)
}

// MinterAddress is a free data retrieval call binding the contract method 0x9c495671.
//
// Solidity: function _minterAddress() view returns(address)
func (_LBCToken *LBCTokenCallerSession) MinterAddress() (common.Address, error) {
	return _LBCToken.Contract.MinterAddress(&_LBCToken.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0x715df33b.
//
// Solidity: function _pauserAddress() view returns(address)
func (_LBCToken *LBCTokenCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "_pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0x715df33b.
//
// Solidity: function _pauserAddress() view returns(address)
func (_LBCToken *LBCTokenSession) PauserAddress() (common.Address, error) {
	return _LBCToken.Contract.PauserAddress(&_LBCToken.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0x715df33b.
//
// Solidity: function _pauserAddress() view returns(address)
func (_LBCToken *LBCTokenCallerSession) PauserAddress() (common.Address, error) {
	return _LBCToken.Contract.PauserAddress(&_LBCToken.CallOpts)
}

// ReserveAddress is a free data retrieval call binding the contract method 0xae5b5b11.
//
// Solidity: function _reserveAddress() view returns(address)
func (_LBCToken *LBCTokenCaller) ReserveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "_reserveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveAddress is a free data retrieval call binding the contract method 0xae5b5b11.
//
// Solidity: function _reserveAddress() view returns(address)
func (_LBCToken *LBCTokenSession) ReserveAddress() (common.Address, error) {
	return _LBCToken.Contract.ReserveAddress(&_LBCToken.CallOpts)
}

// ReserveAddress is a free data retrieval call binding the contract method 0xae5b5b11.
//
// Solidity: function _reserveAddress() view returns(address)
func (_LBCToken *LBCTokenCallerSession) ReserveAddress() (common.Address, error) {
	return _LBCToken.Contract.ReserveAddress(&_LBCToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LBCToken *LBCTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LBCToken *LBCTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LBCToken.Contract.Allowance(&_LBCToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LBCToken *LBCTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LBCToken.Contract.Allowance(&_LBCToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LBCToken *LBCTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LBCToken *LBCTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LBCToken.Contract.BalanceOf(&_LBCToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LBCToken *LBCTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LBCToken.Contract.BalanceOf(&_LBCToken.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_LBCToken *LBCTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_LBCToken *LBCTokenSession) Cap() (*big.Int, error) {
	return _LBCToken.Contract.Cap(&_LBCToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_LBCToken *LBCTokenCallerSession) Cap() (*big.Int, error) {
	return _LBCToken.Contract.Cap(&_LBCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LBCToken *LBCTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LBCToken *LBCTokenSession) Decimals() (uint8, error) {
	return _LBCToken.Contract.Decimals(&_LBCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LBCToken *LBCTokenCallerSession) Decimals() (uint8, error) {
	return _LBCToken.Contract.Decimals(&_LBCToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_LBCToken *LBCTokenCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_LBCToken *LBCTokenSession) Initialized() (bool, error) {
	return _LBCToken.Contract.Initialized(&_LBCToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_LBCToken *LBCTokenCallerSession) Initialized() (bool, error) {
	return _LBCToken.Contract.Initialized(&_LBCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LBCToken *LBCTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LBCToken *LBCTokenSession) Name() (string, error) {
	return _LBCToken.Contract.Name(&_LBCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LBCToken *LBCTokenCallerSession) Name() (string, error) {
	return _LBCToken.Contract.Name(&_LBCToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LBCToken *LBCTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LBCToken *LBCTokenSession) Paused() (bool, error) {
	return _LBCToken.Contract.Paused(&_LBCToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LBCToken *LBCTokenCallerSession) Paused() (bool, error) {
	return _LBCToken.Contract.Paused(&_LBCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LBCToken *LBCTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LBCToken *LBCTokenSession) Symbol() (string, error) {
	return _LBCToken.Contract.Symbol(&_LBCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LBCToken *LBCTokenCallerSession) Symbol() (string, error) {
	return _LBCToken.Contract.Symbol(&_LBCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LBCToken *LBCTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LBCToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LBCToken *LBCTokenSession) TotalSupply() (*big.Int, error) {
	return _LBCToken.Contract.TotalSupply(&_LBCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LBCToken *LBCTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LBCToken.Contract.TotalSupply(&_LBCToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.Approve(&_LBCToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.Approve(&_LBCToken.TransactOpts, spender, amount)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address newMinterAddress) returns()
func (_LBCToken *LBCTokenTransactor) ChangeMinter(opts *bind.TransactOpts, newMinterAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "changeMinter", newMinterAddress)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address newMinterAddress) returns()
func (_LBCToken *LBCTokenSession) ChangeMinter(newMinterAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.Contract.ChangeMinter(&_LBCToken.TransactOpts, newMinterAddress)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address newMinterAddress) returns()
func (_LBCToken *LBCTokenTransactorSession) ChangeMinter(newMinterAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.Contract.ChangeMinter(&_LBCToken.TransactOpts, newMinterAddress)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauserAddress) returns()
func (_LBCToken *LBCTokenTransactor) ChangePauser(opts *bind.TransactOpts, newPauserAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "changePauser", newPauserAddress)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauserAddress) returns()
func (_LBCToken *LBCTokenSession) ChangePauser(newPauserAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.Contract.ChangePauser(&_LBCToken.TransactOpts, newPauserAddress)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauserAddress) returns()
func (_LBCToken *LBCTokenTransactorSession) ChangePauser(newPauserAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.Contract.ChangePauser(&_LBCToken.TransactOpts, newPauserAddress)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LBCToken *LBCTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LBCToken *LBCTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.DecreaseAllowance(&_LBCToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LBCToken *LBCTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.DecreaseAllowance(&_LBCToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LBCToken *LBCTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LBCToken *LBCTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.IncreaseAllowance(&_LBCToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LBCToken *LBCTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.IncreaseAllowance(&_LBCToken.TransactOpts, spender, addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address minterAddress, address pauserAddress, address reserveAddress) returns()
func (_LBCToken *LBCTokenTransactor) Init(opts *bind.TransactOpts, minterAddress common.Address, pauserAddress common.Address, reserveAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "init", minterAddress, pauserAddress, reserveAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address minterAddress, address pauserAddress, address reserveAddress) returns()
func (_LBCToken *LBCTokenSession) Init(minterAddress common.Address, pauserAddress common.Address, reserveAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.Contract.Init(&_LBCToken.TransactOpts, minterAddress, pauserAddress, reserveAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address minterAddress, address pauserAddress, address reserveAddress) returns()
func (_LBCToken *LBCTokenTransactorSession) Init(minterAddress common.Address, pauserAddress common.Address, reserveAddress common.Address) (*types.Transaction, error) {
	return _LBCToken.Contract.Init(&_LBCToken.TransactOpts, minterAddress, pauserAddress, reserveAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.Mint(&_LBCToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.Mint(&_LBCToken.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LBCToken *LBCTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LBCToken *LBCTokenSession) Pause() (*types.Transaction, error) {
	return _LBCToken.Contract.Pause(&_LBCToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LBCToken *LBCTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _LBCToken.Contract.Pause(&_LBCToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.Transfer(&_LBCToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.Transfer(&_LBCToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.TransferFrom(&_LBCToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LBCToken *LBCTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LBCToken.Contract.TransferFrom(&_LBCToken.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LBCToken *LBCTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LBCToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LBCToken *LBCTokenSession) Unpause() (*types.Transaction, error) {
	return _LBCToken.Contract.Unpause(&_LBCToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LBCToken *LBCTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _LBCToken.Contract.Unpause(&_LBCToken.TransactOpts)
}

// LBCTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LBCToken contract.
type LBCTokenApprovalIterator struct {
	Event *LBCTokenApproval // Event containing the contract specifics and raw log

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
func (it *LBCTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenApproval)
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
		it.Event = new(LBCTokenApproval)
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
func (it *LBCTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenApproval represents a Approval event raised by the LBCToken contract.
type LBCTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LBCToken *LBCTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LBCTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LBCTokenApprovalIterator{contract: _LBCToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LBCToken *LBCTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LBCTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenApproval)
				if err := _LBCToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LBCToken *LBCTokenFilterer) ParseApproval(log types.Log) (*LBCTokenApproval, error) {
	event := new(LBCTokenApproval)
	if err := _LBCToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenChangedMinterAddressIterator is returned from FilterChangedMinterAddress and is used to iterate over the raw logs and unpacked data for ChangedMinterAddress events raised by the LBCToken contract.
type LBCTokenChangedMinterAddressIterator struct {
	Event *LBCTokenChangedMinterAddress // Event containing the contract specifics and raw log

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
func (it *LBCTokenChangedMinterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenChangedMinterAddress)
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
		it.Event = new(LBCTokenChangedMinterAddress)
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
func (it *LBCTokenChangedMinterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenChangedMinterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenChangedMinterAddress represents a ChangedMinterAddress event raised by the LBCToken contract.
type LBCTokenChangedMinterAddress struct {
	MinterAddress  common.Address
	ChangerAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedMinterAddress is a free log retrieval operation binding the contract event 0xc6cc4544f0cd8f11ff892e5ea4c8609793a185ba6f5c7ddd2a1f5399bee6774c.
//
// Solidity: event ChangedMinterAddress(address indexed minterAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) FilterChangedMinterAddress(opts *bind.FilterOpts, minterAddress []common.Address, changerAddress []common.Address) (*LBCTokenChangedMinterAddressIterator, error) {

	var minterAddressRule []interface{}
	for _, minterAddressItem := range minterAddress {
		minterAddressRule = append(minterAddressRule, minterAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "ChangedMinterAddress", minterAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return &LBCTokenChangedMinterAddressIterator{contract: _LBCToken.contract, event: "ChangedMinterAddress", logs: logs, sub: sub}, nil
}

// WatchChangedMinterAddress is a free log subscription operation binding the contract event 0xc6cc4544f0cd8f11ff892e5ea4c8609793a185ba6f5c7ddd2a1f5399bee6774c.
//
// Solidity: event ChangedMinterAddress(address indexed minterAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) WatchChangedMinterAddress(opts *bind.WatchOpts, sink chan<- *LBCTokenChangedMinterAddress, minterAddress []common.Address, changerAddress []common.Address) (event.Subscription, error) {

	var minterAddressRule []interface{}
	for _, minterAddressItem := range minterAddress {
		minterAddressRule = append(minterAddressRule, minterAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "ChangedMinterAddress", minterAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenChangedMinterAddress)
				if err := _LBCToken.contract.UnpackLog(event, "ChangedMinterAddress", log); err != nil {
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

// ParseChangedMinterAddress is a log parse operation binding the contract event 0xc6cc4544f0cd8f11ff892e5ea4c8609793a185ba6f5c7ddd2a1f5399bee6774c.
//
// Solidity: event ChangedMinterAddress(address indexed minterAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) ParseChangedMinterAddress(log types.Log) (*LBCTokenChangedMinterAddress, error) {
	event := new(LBCTokenChangedMinterAddress)
	if err := _LBCToken.contract.UnpackLog(event, "ChangedMinterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenChangedPauserAddressIterator is returned from FilterChangedPauserAddress and is used to iterate over the raw logs and unpacked data for ChangedPauserAddress events raised by the LBCToken contract.
type LBCTokenChangedPauserAddressIterator struct {
	Event *LBCTokenChangedPauserAddress // Event containing the contract specifics and raw log

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
func (it *LBCTokenChangedPauserAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenChangedPauserAddress)
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
		it.Event = new(LBCTokenChangedPauserAddress)
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
func (it *LBCTokenChangedPauserAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenChangedPauserAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenChangedPauserAddress represents a ChangedPauserAddress event raised by the LBCToken contract.
type LBCTokenChangedPauserAddress struct {
	PauserAddress  common.Address
	ChangerAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedPauserAddress is a free log retrieval operation binding the contract event 0x2f9de1084416a5a11e235865f4bbfd5c7c9244354207963ce0df47d2ec6193b8.
//
// Solidity: event ChangedPauserAddress(address indexed pauserAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) FilterChangedPauserAddress(opts *bind.FilterOpts, pauserAddress []common.Address, changerAddress []common.Address) (*LBCTokenChangedPauserAddressIterator, error) {

	var pauserAddressRule []interface{}
	for _, pauserAddressItem := range pauserAddress {
		pauserAddressRule = append(pauserAddressRule, pauserAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "ChangedPauserAddress", pauserAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return &LBCTokenChangedPauserAddressIterator{contract: _LBCToken.contract, event: "ChangedPauserAddress", logs: logs, sub: sub}, nil
}

// WatchChangedPauserAddress is a free log subscription operation binding the contract event 0x2f9de1084416a5a11e235865f4bbfd5c7c9244354207963ce0df47d2ec6193b8.
//
// Solidity: event ChangedPauserAddress(address indexed pauserAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) WatchChangedPauserAddress(opts *bind.WatchOpts, sink chan<- *LBCTokenChangedPauserAddress, pauserAddress []common.Address, changerAddress []common.Address) (event.Subscription, error) {

	var pauserAddressRule []interface{}
	for _, pauserAddressItem := range pauserAddress {
		pauserAddressRule = append(pauserAddressRule, pauserAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "ChangedPauserAddress", pauserAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenChangedPauserAddress)
				if err := _LBCToken.contract.UnpackLog(event, "ChangedPauserAddress", log); err != nil {
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

// ParseChangedPauserAddress is a log parse operation binding the contract event 0x2f9de1084416a5a11e235865f4bbfd5c7c9244354207963ce0df47d2ec6193b8.
//
// Solidity: event ChangedPauserAddress(address indexed pauserAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) ParseChangedPauserAddress(log types.Log) (*LBCTokenChangedPauserAddress, error) {
	event := new(LBCTokenChangedPauserAddress)
	if err := _LBCToken.contract.UnpackLog(event, "ChangedPauserAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenChangedReserveAddressIterator is returned from FilterChangedReserveAddress and is used to iterate over the raw logs and unpacked data for ChangedReserveAddress events raised by the LBCToken contract.
type LBCTokenChangedReserveAddressIterator struct {
	Event *LBCTokenChangedReserveAddress // Event containing the contract specifics and raw log

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
func (it *LBCTokenChangedReserveAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenChangedReserveAddress)
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
		it.Event = new(LBCTokenChangedReserveAddress)
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
func (it *LBCTokenChangedReserveAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenChangedReserveAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenChangedReserveAddress represents a ChangedReserveAddress event raised by the LBCToken contract.
type LBCTokenChangedReserveAddress struct {
	ReserveAddress common.Address
	ChangerAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedReserveAddress is a free log retrieval operation binding the contract event 0x7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b54.
//
// Solidity: event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) FilterChangedReserveAddress(opts *bind.FilterOpts, reserveAddress []common.Address, changerAddress []common.Address) (*LBCTokenChangedReserveAddressIterator, error) {

	var reserveAddressRule []interface{}
	for _, reserveAddressItem := range reserveAddress {
		reserveAddressRule = append(reserveAddressRule, reserveAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "ChangedReserveAddress", reserveAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return &LBCTokenChangedReserveAddressIterator{contract: _LBCToken.contract, event: "ChangedReserveAddress", logs: logs, sub: sub}, nil
}

// WatchChangedReserveAddress is a free log subscription operation binding the contract event 0x7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b54.
//
// Solidity: event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) WatchChangedReserveAddress(opts *bind.WatchOpts, sink chan<- *LBCTokenChangedReserveAddress, reserveAddress []common.Address, changerAddress []common.Address) (event.Subscription, error) {

	var reserveAddressRule []interface{}
	for _, reserveAddressItem := range reserveAddress {
		reserveAddressRule = append(reserveAddressRule, reserveAddressItem)
	}
	var changerAddressRule []interface{}
	for _, changerAddressItem := range changerAddress {
		changerAddressRule = append(changerAddressRule, changerAddressItem)
	}

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "ChangedReserveAddress", reserveAddressRule, changerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenChangedReserveAddress)
				if err := _LBCToken.contract.UnpackLog(event, "ChangedReserveAddress", log); err != nil {
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

// ParseChangedReserveAddress is a log parse operation binding the contract event 0x7c07f5a6c7576e884f7d9e01c457774a9059fb680caf5cdd02c86700d3ce1b54.
//
// Solidity: event ChangedReserveAddress(address indexed reserveAddress, address indexed changerAddress)
func (_LBCToken *LBCTokenFilterer) ParseChangedReserveAddress(log types.Log) (*LBCTokenChangedReserveAddress, error) {
	event := new(LBCTokenChangedReserveAddress)
	if err := _LBCToken.contract.UnpackLog(event, "ChangedReserveAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenInitializedContractIterator is returned from FilterInitializedContract and is used to iterate over the raw logs and unpacked data for InitializedContract events raised by the LBCToken contract.
type LBCTokenInitializedContractIterator struct {
	Event *LBCTokenInitializedContract // Event containing the contract specifics and raw log

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
func (it *LBCTokenInitializedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenInitializedContract)
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
		it.Event = new(LBCTokenInitializedContract)
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
func (it *LBCTokenInitializedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenInitializedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenInitializedContract represents a InitializedContract event raised by the LBCToken contract.
type LBCTokenInitializedContract struct {
	ReserveAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitializedContract is a free log retrieval operation binding the contract event 0xaf827135deb94e19593430f16f577c1d98befd728375ca86115192ab05848fcb.
//
// Solidity: event InitializedContract(address indexed reserveAddress)
func (_LBCToken *LBCTokenFilterer) FilterInitializedContract(opts *bind.FilterOpts, reserveAddress []common.Address) (*LBCTokenInitializedContractIterator, error) {

	var reserveAddressRule []interface{}
	for _, reserveAddressItem := range reserveAddress {
		reserveAddressRule = append(reserveAddressRule, reserveAddressItem)
	}

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "InitializedContract", reserveAddressRule)
	if err != nil {
		return nil, err
	}
	return &LBCTokenInitializedContractIterator{contract: _LBCToken.contract, event: "InitializedContract", logs: logs, sub: sub}, nil
}

// WatchInitializedContract is a free log subscription operation binding the contract event 0xaf827135deb94e19593430f16f577c1d98befd728375ca86115192ab05848fcb.
//
// Solidity: event InitializedContract(address indexed reserveAddress)
func (_LBCToken *LBCTokenFilterer) WatchInitializedContract(opts *bind.WatchOpts, sink chan<- *LBCTokenInitializedContract, reserveAddress []common.Address) (event.Subscription, error) {

	var reserveAddressRule []interface{}
	for _, reserveAddressItem := range reserveAddress {
		reserveAddressRule = append(reserveAddressRule, reserveAddressItem)
	}

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "InitializedContract", reserveAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenInitializedContract)
				if err := _LBCToken.contract.UnpackLog(event, "InitializedContract", log); err != nil {
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

// ParseInitializedContract is a log parse operation binding the contract event 0xaf827135deb94e19593430f16f577c1d98befd728375ca86115192ab05848fcb.
//
// Solidity: event InitializedContract(address indexed reserveAddress)
func (_LBCToken *LBCTokenFilterer) ParseInitializedContract(log types.Log) (*LBCTokenInitializedContract, error) {
	event := new(LBCTokenInitializedContract)
	if err := _LBCToken.contract.UnpackLog(event, "InitializedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LBCToken contract.
type LBCTokenPausedIterator struct {
	Event *LBCTokenPaused // Event containing the contract specifics and raw log

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
func (it *LBCTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenPaused)
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
		it.Event = new(LBCTokenPaused)
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
func (it *LBCTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenPaused represents a Paused event raised by the LBCToken contract.
type LBCTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LBCToken *LBCTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*LBCTokenPausedIterator, error) {

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LBCTokenPausedIterator{contract: _LBCToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LBCToken *LBCTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LBCTokenPaused) (event.Subscription, error) {

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenPaused)
				if err := _LBCToken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LBCToken *LBCTokenFilterer) ParsePaused(log types.Log) (*LBCTokenPaused, error) {
	event := new(LBCTokenPaused)
	if err := _LBCToken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LBCToken contract.
type LBCTokenTransferIterator struct {
	Event *LBCTokenTransfer // Event containing the contract specifics and raw log

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
func (it *LBCTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenTransfer)
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
		it.Event = new(LBCTokenTransfer)
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
func (it *LBCTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenTransfer represents a Transfer event raised by the LBCToken contract.
type LBCTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LBCToken *LBCTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LBCTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LBCTokenTransferIterator{contract: _LBCToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LBCToken *LBCTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LBCTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenTransfer)
				if err := _LBCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LBCToken *LBCTokenFilterer) ParseTransfer(log types.Log) (*LBCTokenTransfer, error) {
	event := new(LBCTokenTransfer)
	if err := _LBCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LBCTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LBCToken contract.
type LBCTokenUnpausedIterator struct {
	Event *LBCTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *LBCTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LBCTokenUnpaused)
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
		it.Event = new(LBCTokenUnpaused)
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
func (it *LBCTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LBCTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LBCTokenUnpaused represents a Unpaused event raised by the LBCToken contract.
type LBCTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LBCToken *LBCTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LBCTokenUnpausedIterator, error) {

	logs, sub, err := _LBCToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LBCTokenUnpausedIterator{contract: _LBCToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LBCToken *LBCTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LBCTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _LBCToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LBCTokenUnpaused)
				if err := _LBCToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LBCToken *LBCTokenFilterer) ParseUnpaused(log types.Log) (*LBCTokenUnpaused, error) {
	event := new(LBCTokenUnpaused)
	if err := _LBCToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6102c78061007d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063715018a6146100465780638da5cb5b14610050578063f2fde38b14610074575b600080fd5b61004e61009a565b005b61005861014e565b604080516001600160a01b039092168252519081900360200190f35b61004e6004803603602081101561008a57600080fd5b50356001600160a01b031661015d565b6100a2610267565b6000546001600160a01b03908116911614610104576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b610165610267565b6000546001600160a01b039081169116146101c7576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b03811661020c5760405162461bcd60e51b815260040180806020018281038252602681526020018061026c6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a26469706673582212209816523e0eb3cfe60fbd117cc778e891342fc323a769c72865702a684fdf20f064736f6c63430007060033"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// PausableBin is the compiled bytecode used for deploying new contracts.
var PausableBin = "0x6080604052348015600f57600080fd5b506000805460ff191690556086806100286000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80635c975abb14602d575b600080fd5b60336047565b604080519115158252519081900360200190f35b60005460ff169056fea2646970667358221220c121d7f28a6d8265c26de77e9b8f26472fc05719efa508b2c56e4c802dd7a9c464736f6c63430007060033"

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostDeliveryCrowdsaleABI is the input ABI used to generate the binding from.
const PostDeliveryCrowdsaleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// PostDeliveryCrowdsaleFuncSigs maps the 4-byte function signature to its string representation.
var PostDeliveryCrowdsaleFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"ec8ac4d8": "buyTokens(address)",
	"4b6753bc": "closingTime()",
	"1515bc2b": "hasClosed()",
	"47535d7b": "isOpen()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
	"49df728c": "withdrawTokens(address)",
}

// PostDeliveryCrowdsale is an auto generated Go binding around an Ethereum contract.
type PostDeliveryCrowdsale struct {
	PostDeliveryCrowdsaleCaller     // Read-only binding to the contract
	PostDeliveryCrowdsaleTransactor // Write-only binding to the contract
	PostDeliveryCrowdsaleFilterer   // Log filterer for contract events
}

// PostDeliveryCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PostDeliveryCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostDeliveryCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostDeliveryCrowdsaleSession struct {
	Contract     *PostDeliveryCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PostDeliveryCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostDeliveryCrowdsaleCallerSession struct {
	Contract *PostDeliveryCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PostDeliveryCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostDeliveryCrowdsaleTransactorSession struct {
	Contract     *PostDeliveryCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PostDeliveryCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleRaw struct {
	Contract *PostDeliveryCrowdsale // Generic contract binding to access the raw methods on
}

// PostDeliveryCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleCallerRaw struct {
	Contract *PostDeliveryCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// PostDeliveryCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostDeliveryCrowdsaleTransactorRaw struct {
	Contract *PostDeliveryCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPostDeliveryCrowdsale creates a new instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsale(address common.Address, backend bind.ContractBackend) (*PostDeliveryCrowdsale, error) {
	contract, err := bindPostDeliveryCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsale{PostDeliveryCrowdsaleCaller: PostDeliveryCrowdsaleCaller{contract: contract}, PostDeliveryCrowdsaleTransactor: PostDeliveryCrowdsaleTransactor{contract: contract}, PostDeliveryCrowdsaleFilterer: PostDeliveryCrowdsaleFilterer{contract: contract}}, nil
}

// NewPostDeliveryCrowdsaleCaller creates a new read-only instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*PostDeliveryCrowdsaleCaller, error) {
	contract, err := bindPostDeliveryCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleCaller{contract: contract}, nil
}

// NewPostDeliveryCrowdsaleTransactor creates a new write-only instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*PostDeliveryCrowdsaleTransactor, error) {
	contract, err := bindPostDeliveryCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleTransactor{contract: contract}, nil
}

// NewPostDeliveryCrowdsaleFilterer creates a new log filterer instance of PostDeliveryCrowdsale, bound to a specific deployed contract.
func NewPostDeliveryCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*PostDeliveryCrowdsaleFilterer, error) {
	contract, err := bindPostDeliveryCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleFilterer{contract: contract}, nil
}

// bindPostDeliveryCrowdsale binds a generic wrapper to an already deployed contract.
func bindPostDeliveryCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PostDeliveryCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PostDeliveryCrowdsale.Contract.PostDeliveryCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.PostDeliveryCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.PostDeliveryCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PostDeliveryCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.BalanceOf(&_PostDeliveryCrowdsale.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.BalanceOf(&_PostDeliveryCrowdsale.CallOpts, account)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.ClosingTime(&_PostDeliveryCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.ClosingTime(&_PostDeliveryCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) HasClosed() (bool, error) {
	return _PostDeliveryCrowdsale.Contract.HasClosed(&_PostDeliveryCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _PostDeliveryCrowdsale.Contract.HasClosed(&_PostDeliveryCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) IsOpen() (bool, error) {
	return _PostDeliveryCrowdsale.Contract.IsOpen(&_PostDeliveryCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _PostDeliveryCrowdsale.Contract.IsOpen(&_PostDeliveryCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.OpeningTime(&_PostDeliveryCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.OpeningTime(&_PostDeliveryCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Rate() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.Rate(&_PostDeliveryCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.Rate(&_PostDeliveryCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Token() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Token(&_PostDeliveryCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Token() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Token(&_PostDeliveryCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Wallet() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Wallet(&_PostDeliveryCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _PostDeliveryCrowdsale.Contract.Wallet(&_PostDeliveryCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PostDeliveryCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.WeiRaised(&_PostDeliveryCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _PostDeliveryCrowdsale.Contract.WeiRaised(&_PostDeliveryCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.BuyTokens(&_PostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.BuyTokens(&_PostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactor) WithdrawTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.contract.Transact(opts, "withdrawTokens", beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) WithdrawTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.WithdrawTokens(&_PostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorSession) WithdrawTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.WithdrawTokens(&_PostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.Fallback(&_PostDeliveryCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.Fallback(&_PostDeliveryCrowdsale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.Receive(&_PostDeliveryCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _PostDeliveryCrowdsale.Contract.Receive(&_PostDeliveryCrowdsale.TransactOpts)
}

// PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the PostDeliveryCrowdsale contract.
type PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *PostDeliveryCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostDeliveryCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(PostDeliveryCrowdsaleTimedCrowdsaleExtended)
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
func (it *PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostDeliveryCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the PostDeliveryCrowdsale contract.
type PostDeliveryCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _PostDeliveryCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator{contract: _PostDeliveryCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *PostDeliveryCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _PostDeliveryCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostDeliveryCrowdsaleTimedCrowdsaleExtended)
				if err := _PostDeliveryCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*PostDeliveryCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(PostDeliveryCrowdsaleTimedCrowdsaleExtended)
	if err := _PostDeliveryCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostDeliveryCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the PostDeliveryCrowdsale contract.
type PostDeliveryCrowdsaleTokensPurchasedIterator struct {
	Event *PostDeliveryCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *PostDeliveryCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostDeliveryCrowdsaleTokensPurchased)
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
		it.Event = new(PostDeliveryCrowdsaleTokensPurchased)
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
func (it *PostDeliveryCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostDeliveryCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostDeliveryCrowdsaleTokensPurchased represents a TokensPurchased event raised by the PostDeliveryCrowdsale contract.
type PostDeliveryCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*PostDeliveryCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _PostDeliveryCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &PostDeliveryCrowdsaleTokensPurchasedIterator{contract: _PostDeliveryCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *PostDeliveryCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _PostDeliveryCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostDeliveryCrowdsaleTokensPurchased)
				if err := _PostDeliveryCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_PostDeliveryCrowdsale *PostDeliveryCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*PostDeliveryCrowdsaleTokensPurchased, error) {
	event := new(PostDeliveryCrowdsaleTokensPurchased)
	if err := _PostDeliveryCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyGuardBin is the compiled bytecode used for deploying new contracts.
var ReentrancyGuardBin = "0x6080604052348015600f57600080fd5b506001600055603f8060226000396000f3fe6080604052600080fdfea2646970667358221220981a5a57401bbb0dcaf8587a63f23b178df4b2cb6f9b293fbd68f9af8c31daf964736f6c63430007060033"

// DeployReentrancyGuard deploys a new Ethereum contract, binding an instance of ReentrancyGuard to it.
func DeployReentrancyGuard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyGuard, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancyGuardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// RefundEscrowABI is the input ABI used to generate the binding from.
const RefundEscrowABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"beneficiaryReceived\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RefundsClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RefundsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiaryWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"depositsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRefunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumRefundEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawalAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RefundEscrowFuncSigs maps the 4-byte function signature to its string representation.
var RefundEscrowFuncSigs = map[string]string{
	"38af3eed": "beneficiary()",
	"9af6549a": "beneficiaryWithdraw()",
	"43d726d6": "close()",
	"f340fa01": "deposit(address)",
	"e3a9db1a": "depositsOf(address)",
	"8c52dc41": "enableRefunds()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"c19d93fb": "state()",
	"f2fde38b": "transferOwnership(address)",
	"51cff8d9": "withdraw(address)",
	"685ca194": "withdrawalAllowed(address)",
}

// RefundEscrowBin is the compiled bytecode used for deploying new contracts.
var RefundEscrowBin = "0x608060405234801561001057600080fd5b50604051610d4f380380610d4f8339818101604052602081101561003357600080fd5b5051600061003f6100fe565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001600160a01b0381166100ce5760405162461bcd60e51b815260040180806020018281038252602d815260200180610d22602d913960400191505060405180910390fd5b6002805460ff196001600160a01b039390931661010002610100600160a81b031990911617919091169055610102565b3390565b610c11806101116000396000f3fe6080604052600436106100a75760003560e01c80638da5cb5b116100645780638da5cb5b146101985780639af6549a146101ad578063c19d93fb146101c2578063e3a9db1a146101f8578063f2fde38b1461023d578063f340fa0114610270576100a7565b806338af3eed146100ac57806343d726d6146100dd57806351cff8d9146100f4578063685ca19414610127578063715018a61461016e5780638c52dc4114610183575b600080fd5b3480156100b857600080fd5b506100c1610296565b604080516001600160a01b039092168252519081900360200190f35b3480156100e957600080fd5b506100f26102aa565b005b34801561010057600080fd5b506100f26004803603602081101561011757600080fd5b50356001600160a01b0316610388565b34801561013357600080fd5b5061015a6004803603602081101561014a57600080fd5b50356001600160a01b03166103d8565b604080519115158252519081900360200190f35b34801561017a57600080fd5b506100f26103f4565b34801561018f57600080fd5b506100f2610496565b3480156101a457600080fd5b506100c1610575565b3480156101b957600080fd5b506100f2610584565b3480156101ce57600080fd5b506101d7610611565b604051808260028111156101e757fe5b815260200191505060405180910390f35b34801561020457600080fd5b5061022b6004803603602081101561021b57600080fd5b50356001600160a01b031661061a565b60408051918252519081900360200190f35b34801561024957600080fd5b506100f26004803603602081101561026057600080fd5b50356001600160a01b0316610635565b6100f26004803603602081101561028657600080fd5b50356001600160a01b031661072d565b60025461010090046001600160a01b031690565b6102b2610785565b6000546001600160a01b03908116911614610302576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b60006002805460ff169081111561031557fe5b146103515760405162461bcd60e51b8152600401808060200182810382526029815260200180610b616029913960400191505060405180910390fd5b6002805460ff1916811790556040517f088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f90600090a1565b610391816103d8565b6103cc5760405162461bcd60e51b8152600401808060200182810382526033815260200180610b2e6033913960400191505060405180910390fd5b6103d581610789565b50565b600060016002805460ff16908111156103ed57fe5b1492915050565b6103fc610785565b6000546001600160a01b0390811691161461044c576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b61049e610785565b6000546001600160a01b039081169116146104ee576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b60006002805460ff169081111561050157fe5b1461053d5760405162461bcd60e51b8152600401808060200182810382526032815260200180610baa6032913960400191505060405180910390fd5b6002805460ff191660011790556040517f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8990600090a1565b6000546001600160a01b031690565b6002805460ff168181111561059557fe5b146105d15760405162461bcd60e51b8152600401808060200182810382526038815260200180610a6b6038913960400191505060405180910390fd5b6002546040516001600160a01b0361010090920491909116904780156108fc02916000818181858888f193505050501580156103d5573d6000803e3d6000fd5b60025460ff1690565b6001600160a01b031660009081526001602052604090205490565b61063d610785565b6000546001600160a01b0390811691161461068d576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b6001600160a01b0381166106d25760405162461bcd60e51b8152600401808060200182810382526026815260200180610aa36026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006002805460ff169081111561074057fe5b1461077c5760405162461bcd60e51b815260040180806020018281038252602b815260200180610b03602b913960400191505060405180910390fd5b6103d58161084c565b3390565b610791610785565b6000546001600160a01b039081169116146107e1576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b6001600160a01b0381166000818152600160205260408120805491905590610809908261091f565b6040805182815290516001600160a01b038416917f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5919081900360200190a25050565b610854610785565b6000546001600160a01b039081169116146108a4576040805162461bcd60e51b81526020600482018190526024820152600080516020610b8a833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526001602052604090205434906108c99082610a09565b6001600160a01b038316600081815260016020908152604091829020939093558051848152905191927f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c492918290030190a25050565b80471015610974576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d80600081146109bf576040519150601f19603f3d011682016040523d82523d6000602084013e6109c4565b606091505b5050905080610a045760405162461bcd60e51b815260040180806020018281038252603a815260200180610ac9603a913960400191505060405180910390fd5b505050565b600082820183811015610a63576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe526566756e64457363726f773a2062656e65666963696172792063616e206f6e6c79207769746864726177207768696c6520636c6f7365644f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d61792068617665207265766572746564526566756e64457363726f773a2063616e206f6e6c79206465706f736974207768696c6520616374697665436f6e646974696f6e616c457363726f773a207061796565206973206e6f7420616c6c6f77656420746f207769746864726177526566756e64457363726f773a2063616e206f6e6c7920636c6f7365207768696c65206163746976654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572526566756e64457363726f773a2063616e206f6e6c7920656e61626c6520726566756e6473207768696c6520616374697665a2646970667358221220415d58c8a6368e7076ac118e308228241ab855c49dceef8931507f3d6665159d64736f6c63430007060033526566756e64457363726f773a2062656e656669636961727920697320746865207a65726f2061646472657373"

// DeployRefundEscrow deploys a new Ethereum contract, binding an instance of RefundEscrow to it.
func DeployRefundEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, beneficiaryReceived common.Address) (common.Address, *types.Transaction, *RefundEscrow, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundEscrowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RefundEscrowBin), backend, beneficiaryReceived)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RefundEscrow{RefundEscrowCaller: RefundEscrowCaller{contract: contract}, RefundEscrowTransactor: RefundEscrowTransactor{contract: contract}, RefundEscrowFilterer: RefundEscrowFilterer{contract: contract}}, nil
}

// RefundEscrow is an auto generated Go binding around an Ethereum contract.
type RefundEscrow struct {
	RefundEscrowCaller     // Read-only binding to the contract
	RefundEscrowTransactor // Write-only binding to the contract
	RefundEscrowFilterer   // Log filterer for contract events
}

// RefundEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefundEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundEscrowSession struct {
	Contract     *RefundEscrow     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefundEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundEscrowCallerSession struct {
	Contract *RefundEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RefundEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundEscrowTransactorSession struct {
	Contract     *RefundEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RefundEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundEscrowRaw struct {
	Contract *RefundEscrow // Generic contract binding to access the raw methods on
}

// RefundEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundEscrowCallerRaw struct {
	Contract *RefundEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// RefundEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundEscrowTransactorRaw struct {
	Contract *RefundEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundEscrow creates a new instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrow(address common.Address, backend bind.ContractBackend) (*RefundEscrow, error) {
	contract, err := bindRefundEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundEscrow{RefundEscrowCaller: RefundEscrowCaller{contract: contract}, RefundEscrowTransactor: RefundEscrowTransactor{contract: contract}, RefundEscrowFilterer: RefundEscrowFilterer{contract: contract}}, nil
}

// NewRefundEscrowCaller creates a new read-only instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrowCaller(address common.Address, caller bind.ContractCaller) (*RefundEscrowCaller, error) {
	contract, err := bindRefundEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowCaller{contract: contract}, nil
}

// NewRefundEscrowTransactor creates a new write-only instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundEscrowTransactor, error) {
	contract, err := bindRefundEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowTransactor{contract: contract}, nil
}

// NewRefundEscrowFilterer creates a new log filterer instance of RefundEscrow, bound to a specific deployed contract.
func NewRefundEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*RefundEscrowFilterer, error) {
	contract, err := bindRefundEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowFilterer{contract: contract}, nil
}

// bindRefundEscrow binds a generic wrapper to an already deployed contract.
func bindRefundEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundEscrow *RefundEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundEscrow.Contract.RefundEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundEscrow *RefundEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.Contract.RefundEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundEscrow *RefundEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundEscrow.Contract.RefundEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundEscrow *RefundEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundEscrow *RefundEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundEscrow *RefundEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundEscrow.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RefundEscrow *RefundEscrowCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RefundEscrow *RefundEscrowSession) Beneficiary() (common.Address, error) {
	return _RefundEscrow.Contract.Beneficiary(&_RefundEscrow.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RefundEscrow *RefundEscrowCallerSession) Beneficiary() (common.Address, error) {
	return _RefundEscrow.Contract.Beneficiary(&_RefundEscrow.CallOpts)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_RefundEscrow *RefundEscrowCaller) DepositsOf(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "depositsOf", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_RefundEscrow *RefundEscrowSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _RefundEscrow.Contract.DepositsOf(&_RefundEscrow.CallOpts, payee)
}

// DepositsOf is a free data retrieval call binding the contract method 0xe3a9db1a.
//
// Solidity: function depositsOf(address payee) view returns(uint256)
func (_RefundEscrow *RefundEscrowCallerSession) DepositsOf(payee common.Address) (*big.Int, error) {
	return _RefundEscrow.Contract.DepositsOf(&_RefundEscrow.CallOpts, payee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefundEscrow *RefundEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefundEscrow *RefundEscrowSession) Owner() (common.Address, error) {
	return _RefundEscrow.Contract.Owner(&_RefundEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefundEscrow *RefundEscrowCallerSession) Owner() (common.Address, error) {
	return _RefundEscrow.Contract.Owner(&_RefundEscrow.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_RefundEscrow *RefundEscrowCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_RefundEscrow *RefundEscrowSession) State() (uint8, error) {
	return _RefundEscrow.Contract.State(&_RefundEscrow.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_RefundEscrow *RefundEscrowCallerSession) State() (uint8, error) {
	return _RefundEscrow.Contract.State(&_RefundEscrow.CallOpts)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address ) view returns(bool)
func (_RefundEscrow *RefundEscrowCaller) WithdrawalAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RefundEscrow.contract.Call(opts, &out, "withdrawalAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address ) view returns(bool)
func (_RefundEscrow *RefundEscrowSession) WithdrawalAllowed(arg0 common.Address) (bool, error) {
	return _RefundEscrow.Contract.WithdrawalAllowed(&_RefundEscrow.CallOpts, arg0)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0x685ca194.
//
// Solidity: function withdrawalAllowed(address ) view returns(bool)
func (_RefundEscrow *RefundEscrowCallerSession) WithdrawalAllowed(arg0 common.Address) (bool, error) {
	return _RefundEscrow.Contract.WithdrawalAllowed(&_RefundEscrow.CallOpts, arg0)
}

// BeneficiaryWithdraw is a paid mutator transaction binding the contract method 0x9af6549a.
//
// Solidity: function beneficiaryWithdraw() returns()
func (_RefundEscrow *RefundEscrowTransactor) BeneficiaryWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "beneficiaryWithdraw")
}

// BeneficiaryWithdraw is a paid mutator transaction binding the contract method 0x9af6549a.
//
// Solidity: function beneficiaryWithdraw() returns()
func (_RefundEscrow *RefundEscrowSession) BeneficiaryWithdraw() (*types.Transaction, error) {
	return _RefundEscrow.Contract.BeneficiaryWithdraw(&_RefundEscrow.TransactOpts)
}

// BeneficiaryWithdraw is a paid mutator transaction binding the contract method 0x9af6549a.
//
// Solidity: function beneficiaryWithdraw() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) BeneficiaryWithdraw() (*types.Transaction, error) {
	return _RefundEscrow.Contract.BeneficiaryWithdraw(&_RefundEscrow.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundEscrow *RefundEscrowTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundEscrow *RefundEscrowSession) Close() (*types.Transaction, error) {
	return _RefundEscrow.Contract.Close(&_RefundEscrow.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) Close() (*types.Transaction, error) {
	return _RefundEscrow.Contract.Close(&_RefundEscrow.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address refundee) payable returns()
func (_RefundEscrow *RefundEscrowTransactor) Deposit(opts *bind.TransactOpts, refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "deposit", refundee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address refundee) payable returns()
func (_RefundEscrow *RefundEscrowSession) Deposit(refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Deposit(&_RefundEscrow.TransactOpts, refundee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address refundee) payable returns()
func (_RefundEscrow *RefundEscrowTransactorSession) Deposit(refundee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Deposit(&_RefundEscrow.TransactOpts, refundee)
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundEscrow *RefundEscrowTransactor) EnableRefunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "enableRefunds")
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundEscrow *RefundEscrowSession) EnableRefunds() (*types.Transaction, error) {
	return _RefundEscrow.Contract.EnableRefunds(&_RefundEscrow.TransactOpts)
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) EnableRefunds() (*types.Transaction, error) {
	return _RefundEscrow.Contract.EnableRefunds(&_RefundEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefundEscrow *RefundEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefundEscrow *RefundEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _RefundEscrow.Contract.RenounceOwnership(&_RefundEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefundEscrow *RefundEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RefundEscrow.Contract.RenounceOwnership(&_RefundEscrow.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefundEscrow *RefundEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefundEscrow *RefundEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.TransferOwnership(&_RefundEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.TransferOwnership(&_RefundEscrow.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_RefundEscrow *RefundEscrowTransactor) Withdraw(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.contract.Transact(opts, "withdraw", payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_RefundEscrow *RefundEscrowSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Withdraw(&_RefundEscrow.TransactOpts, payee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address payee) returns()
func (_RefundEscrow *RefundEscrowTransactorSession) Withdraw(payee common.Address) (*types.Transaction, error) {
	return _RefundEscrow.Contract.Withdraw(&_RefundEscrow.TransactOpts, payee)
}

// RefundEscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the RefundEscrow contract.
type RefundEscrowDepositedIterator struct {
	Event *RefundEscrowDeposited // Event containing the contract specifics and raw log

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
func (it *RefundEscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowDeposited)
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
		it.Event = new(RefundEscrowDeposited)
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
func (it *RefundEscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowDeposited represents a Deposited event raised by the RefundEscrow contract.
type RefundEscrowDeposited struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) FilterDeposited(opts *bind.FilterOpts, payee []common.Address) (*RefundEscrowDepositedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowDepositedIterator{contract: _RefundEscrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RefundEscrowDeposited, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "Deposited", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowDeposited)
				if err := _RefundEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) ParseDeposited(log types.Log) (*RefundEscrowDeposited, error) {
	event := new(RefundEscrowDeposited)
	if err := _RefundEscrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RefundEscrow contract.
type RefundEscrowOwnershipTransferredIterator struct {
	Event *RefundEscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RefundEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowOwnershipTransferred)
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
		it.Event = new(RefundEscrowOwnershipTransferred)
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
func (it *RefundEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the RefundEscrow contract.
type RefundEscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefundEscrow *RefundEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RefundEscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowOwnershipTransferredIterator{contract: _RefundEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefundEscrow *RefundEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RefundEscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowOwnershipTransferred)
				if err := _RefundEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefundEscrow *RefundEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*RefundEscrowOwnershipTransferred, error) {
	event := new(RefundEscrowOwnershipTransferred)
	if err := _RefundEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowRefundsClosedIterator is returned from FilterRefundsClosed and is used to iterate over the raw logs and unpacked data for RefundsClosed events raised by the RefundEscrow contract.
type RefundEscrowRefundsClosedIterator struct {
	Event *RefundEscrowRefundsClosed // Event containing the contract specifics and raw log

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
func (it *RefundEscrowRefundsClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowRefundsClosed)
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
		it.Event = new(RefundEscrowRefundsClosed)
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
func (it *RefundEscrowRefundsClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowRefundsClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowRefundsClosed represents a RefundsClosed event raised by the RefundEscrow contract.
type RefundEscrowRefundsClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefundsClosed is a free log retrieval operation binding the contract event 0x088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f.
//
// Solidity: event RefundsClosed()
func (_RefundEscrow *RefundEscrowFilterer) FilterRefundsClosed(opts *bind.FilterOpts) (*RefundEscrowRefundsClosedIterator, error) {

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "RefundsClosed")
	if err != nil {
		return nil, err
	}
	return &RefundEscrowRefundsClosedIterator{contract: _RefundEscrow.contract, event: "RefundsClosed", logs: logs, sub: sub}, nil
}

// WatchRefundsClosed is a free log subscription operation binding the contract event 0x088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f.
//
// Solidity: event RefundsClosed()
func (_RefundEscrow *RefundEscrowFilterer) WatchRefundsClosed(opts *bind.WatchOpts, sink chan<- *RefundEscrowRefundsClosed) (event.Subscription, error) {

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "RefundsClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowRefundsClosed)
				if err := _RefundEscrow.contract.UnpackLog(event, "RefundsClosed", log); err != nil {
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

// ParseRefundsClosed is a log parse operation binding the contract event 0x088672c3a6e342f7cd94a65ba63b79df24a8973927b4d05d803c44bbf787d12f.
//
// Solidity: event RefundsClosed()
func (_RefundEscrow *RefundEscrowFilterer) ParseRefundsClosed(log types.Log) (*RefundEscrowRefundsClosed, error) {
	event := new(RefundEscrowRefundsClosed)
	if err := _RefundEscrow.contract.UnpackLog(event, "RefundsClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowRefundsEnabledIterator is returned from FilterRefundsEnabled and is used to iterate over the raw logs and unpacked data for RefundsEnabled events raised by the RefundEscrow contract.
type RefundEscrowRefundsEnabledIterator struct {
	Event *RefundEscrowRefundsEnabled // Event containing the contract specifics and raw log

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
func (it *RefundEscrowRefundsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowRefundsEnabled)
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
		it.Event = new(RefundEscrowRefundsEnabled)
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
func (it *RefundEscrowRefundsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowRefundsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowRefundsEnabled represents a RefundsEnabled event raised by the RefundEscrow contract.
type RefundEscrowRefundsEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefundsEnabled is a free log retrieval operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundEscrow *RefundEscrowFilterer) FilterRefundsEnabled(opts *bind.FilterOpts) (*RefundEscrowRefundsEnabledIterator, error) {

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "RefundsEnabled")
	if err != nil {
		return nil, err
	}
	return &RefundEscrowRefundsEnabledIterator{contract: _RefundEscrow.contract, event: "RefundsEnabled", logs: logs, sub: sub}, nil
}

// WatchRefundsEnabled is a free log subscription operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundEscrow *RefundEscrowFilterer) WatchRefundsEnabled(opts *bind.WatchOpts, sink chan<- *RefundEscrowRefundsEnabled) (event.Subscription, error) {

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "RefundsEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowRefundsEnabled)
				if err := _RefundEscrow.contract.UnpackLog(event, "RefundsEnabled", log); err != nil {
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

// ParseRefundsEnabled is a log parse operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundEscrow *RefundEscrowFilterer) ParseRefundsEnabled(log types.Log) (*RefundEscrowRefundsEnabled, error) {
	event := new(RefundEscrowRefundsEnabled)
	if err := _RefundEscrow.contract.UnpackLog(event, "RefundsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundEscrowWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RefundEscrow contract.
type RefundEscrowWithdrawnIterator struct {
	Event *RefundEscrowWithdrawn // Event containing the contract specifics and raw log

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
func (it *RefundEscrowWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundEscrowWithdrawn)
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
		it.Event = new(RefundEscrowWithdrawn)
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
func (it *RefundEscrowWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundEscrowWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundEscrowWithdrawn represents a Withdrawn event raised by the RefundEscrow contract.
type RefundEscrowWithdrawn struct {
	Payee     common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) FilterWithdrawn(opts *bind.FilterOpts, payee []common.Address) (*RefundEscrowWithdrawnIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.FilterLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return &RefundEscrowWithdrawnIterator{contract: _RefundEscrow.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RefundEscrowWithdrawn, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _RefundEscrow.contract.WatchLogs(opts, "Withdrawn", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundEscrowWithdrawn)
				if err := _RefundEscrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed payee, uint256 weiAmount)
func (_RefundEscrow *RefundEscrowFilterer) ParseWithdrawn(log types.Log) (*RefundEscrowWithdrawn, error) {
	event := new(RefundEscrowWithdrawn)
	if err := _RefundEscrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundableCrowdsaleABI is the input ABI used to generate the binding from.
const RefundableCrowdsaleABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"CrowdsaleFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"}],\"name\":\"claimRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// RefundableCrowdsaleFuncSigs maps the 4-byte function signature to its string representation.
var RefundableCrowdsaleFuncSigs = map[string]string{
	"ec8ac4d8": "buyTokens(address)",
	"bffa55d5": "claimRefund(address)",
	"4b6753bc": "closingTime()",
	"4bb278f3": "finalize()",
	"b3f05b97": "finalized()",
	"40193883": "goal()",
	"7d3d6522": "goalReached()",
	"1515bc2b": "hasClosed()",
	"47535d7b": "isOpen()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
}

// RefundableCrowdsale is an auto generated Go binding around an Ethereum contract.
type RefundableCrowdsale struct {
	RefundableCrowdsaleCaller     // Read-only binding to the contract
	RefundableCrowdsaleTransactor // Write-only binding to the contract
	RefundableCrowdsaleFilterer   // Log filterer for contract events
}

// RefundableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefundableCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundableCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundableCrowdsaleSession struct {
	Contract     *RefundableCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RefundableCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundableCrowdsaleCallerSession struct {
	Contract *RefundableCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RefundableCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundableCrowdsaleTransactorSession struct {
	Contract     *RefundableCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RefundableCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundableCrowdsaleRaw struct {
	Contract *RefundableCrowdsale // Generic contract binding to access the raw methods on
}

// RefundableCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundableCrowdsaleCallerRaw struct {
	Contract *RefundableCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// RefundableCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundableCrowdsaleTransactorRaw struct {
	Contract *RefundableCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundableCrowdsale creates a new instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsale(address common.Address, backend bind.ContractBackend) (*RefundableCrowdsale, error) {
	contract, err := bindRefundableCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsale{RefundableCrowdsaleCaller: RefundableCrowdsaleCaller{contract: contract}, RefundableCrowdsaleTransactor: RefundableCrowdsaleTransactor{contract: contract}, RefundableCrowdsaleFilterer: RefundableCrowdsaleFilterer{contract: contract}}, nil
}

// NewRefundableCrowdsaleCaller creates a new read-only instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*RefundableCrowdsaleCaller, error) {
	contract, err := bindRefundableCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleCaller{contract: contract}, nil
}

// NewRefundableCrowdsaleTransactor creates a new write-only instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundableCrowdsaleTransactor, error) {
	contract, err := bindRefundableCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleTransactor{contract: contract}, nil
}

// NewRefundableCrowdsaleFilterer creates a new log filterer instance of RefundableCrowdsale, bound to a specific deployed contract.
func NewRefundableCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*RefundableCrowdsaleFilterer, error) {
	contract, err := bindRefundableCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleFilterer{contract: contract}, nil
}

// bindRefundableCrowdsale binds a generic wrapper to an already deployed contract.
func bindRefundableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundableCrowdsale *RefundableCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundableCrowdsale.Contract.RefundableCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundableCrowdsale *RefundableCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.RefundableCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundableCrowdsale *RefundableCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.RefundableCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundableCrowdsale *RefundableCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundableCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.ClosingTime(&_RefundableCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.ClosingTime(&_RefundableCrowdsale.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Finalized() (bool, error) {
	return _RefundableCrowdsale.Contract.Finalized(&_RefundableCrowdsale.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Finalized() (bool, error) {
	return _RefundableCrowdsale.Contract.Finalized(&_RefundableCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "goal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Goal() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Goal(&_RefundableCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Goal() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Goal(&_RefundableCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "goalReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) GoalReached() (bool, error) {
	return _RefundableCrowdsale.Contract.GoalReached(&_RefundableCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) GoalReached() (bool, error) {
	return _RefundableCrowdsale.Contract.GoalReached(&_RefundableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) HasClosed() (bool, error) {
	return _RefundableCrowdsale.Contract.HasClosed(&_RefundableCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _RefundableCrowdsale.Contract.HasClosed(&_RefundableCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) IsOpen() (bool, error) {
	return _RefundableCrowdsale.Contract.IsOpen(&_RefundableCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _RefundableCrowdsale.Contract.IsOpen(&_RefundableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.OpeningTime(&_RefundableCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.OpeningTime(&_RefundableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Rate() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Rate(&_RefundableCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.Rate(&_RefundableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Token() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Token(&_RefundableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Token() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Token(&_RefundableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Wallet() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Wallet(&_RefundableCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _RefundableCrowdsale.Contract.Wallet(&_RefundableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundableCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.WeiRaised(&_RefundableCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_RefundableCrowdsale *RefundableCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _RefundableCrowdsale.Contract.WeiRaised(&_RefundableCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.BuyTokens(&_RefundableCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.BuyTokens(&_RefundableCrowdsale.TransactOpts, beneficiary)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) ClaimRefund(opts *bind.TransactOpts, refundee common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "claimRefund", refundee)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) ClaimRefund(refundee common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.ClaimRefund(&_RefundableCrowdsale.TransactOpts, refundee)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) ClaimRefund(refundee common.Address) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.ClaimRefund(&_RefundableCrowdsale.TransactOpts, refundee)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Finalize(&_RefundableCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Finalize(&_RefundableCrowdsale.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Fallback(&_RefundableCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Fallback(&_RefundableCrowdsale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundableCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Receive(&_RefundableCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefundableCrowdsale *RefundableCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _RefundableCrowdsale.Contract.Receive(&_RefundableCrowdsale.TransactOpts)
}

// RefundableCrowdsaleCrowdsaleFinalizedIterator is returned from FilterCrowdsaleFinalized and is used to iterate over the raw logs and unpacked data for CrowdsaleFinalized events raised by the RefundableCrowdsale contract.
type RefundableCrowdsaleCrowdsaleFinalizedIterator struct {
	Event *RefundableCrowdsaleCrowdsaleFinalized // Event containing the contract specifics and raw log

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
func (it *RefundableCrowdsaleCrowdsaleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundableCrowdsaleCrowdsaleFinalized)
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
		it.Event = new(RefundableCrowdsaleCrowdsaleFinalized)
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
func (it *RefundableCrowdsaleCrowdsaleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundableCrowdsaleCrowdsaleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundableCrowdsaleCrowdsaleFinalized represents a CrowdsaleFinalized event raised by the RefundableCrowdsale contract.
type RefundableCrowdsaleCrowdsaleFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrowdsaleFinalized is a free log retrieval operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) FilterCrowdsaleFinalized(opts *bind.FilterOpts) (*RefundableCrowdsaleCrowdsaleFinalizedIterator, error) {

	logs, sub, err := _RefundableCrowdsale.contract.FilterLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleCrowdsaleFinalizedIterator{contract: _RefundableCrowdsale.contract, event: "CrowdsaleFinalized", logs: logs, sub: sub}, nil
}

// WatchCrowdsaleFinalized is a free log subscription operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) WatchCrowdsaleFinalized(opts *bind.WatchOpts, sink chan<- *RefundableCrowdsaleCrowdsaleFinalized) (event.Subscription, error) {

	logs, sub, err := _RefundableCrowdsale.contract.WatchLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundableCrowdsaleCrowdsaleFinalized)
				if err := _RefundableCrowdsale.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
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

// ParseCrowdsaleFinalized is a log parse operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) ParseCrowdsaleFinalized(log types.Log) (*RefundableCrowdsaleCrowdsaleFinalized, error) {
	event := new(RefundableCrowdsaleCrowdsaleFinalized)
	if err := _RefundableCrowdsale.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundableCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the RefundableCrowdsale contract.
type RefundableCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *RefundableCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *RefundableCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundableCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(RefundableCrowdsaleTimedCrowdsaleExtended)
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
func (it *RefundableCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundableCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundableCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the RefundableCrowdsale contract.
type RefundableCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*RefundableCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _RefundableCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleTimedCrowdsaleExtendedIterator{contract: _RefundableCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *RefundableCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _RefundableCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundableCrowdsaleTimedCrowdsaleExtended)
				if err := _RefundableCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*RefundableCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(RefundableCrowdsaleTimedCrowdsaleExtended)
	if err := _RefundableCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundableCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the RefundableCrowdsale contract.
type RefundableCrowdsaleTokensPurchasedIterator struct {
	Event *RefundableCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *RefundableCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundableCrowdsaleTokensPurchased)
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
		it.Event = new(RefundableCrowdsaleTokensPurchased)
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
func (it *RefundableCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundableCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundableCrowdsaleTokensPurchased represents a TokensPurchased event raised by the RefundableCrowdsale contract.
type RefundableCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*RefundableCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundableCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &RefundableCrowdsaleTokensPurchasedIterator{contract: _RefundableCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *RefundableCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundableCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundableCrowdsaleTokensPurchased)
				if err := _RefundableCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_RefundableCrowdsale *RefundableCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*RefundableCrowdsaleTokensPurchased, error) {
	event := new(RefundableCrowdsaleTokensPurchased)
	if err := _RefundableCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundablePostDeliveryCrowdsaleABI is the input ABI used to generate the binding from.
const RefundablePostDeliveryCrowdsaleABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"CrowdsaleFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"}],\"name\":\"claimRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// RefundablePostDeliveryCrowdsaleFuncSigs maps the 4-byte function signature to its string representation.
var RefundablePostDeliveryCrowdsaleFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"ec8ac4d8": "buyTokens(address)",
	"bffa55d5": "claimRefund(address)",
	"4b6753bc": "closingTime()",
	"4bb278f3": "finalize()",
	"b3f05b97": "finalized()",
	"40193883": "goal()",
	"7d3d6522": "goalReached()",
	"1515bc2b": "hasClosed()",
	"47535d7b": "isOpen()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
	"49df728c": "withdrawTokens(address)",
}

// RefundablePostDeliveryCrowdsale is an auto generated Go binding around an Ethereum contract.
type RefundablePostDeliveryCrowdsale struct {
	RefundablePostDeliveryCrowdsaleCaller     // Read-only binding to the contract
	RefundablePostDeliveryCrowdsaleTransactor // Write-only binding to the contract
	RefundablePostDeliveryCrowdsaleFilterer   // Log filterer for contract events
}

// RefundablePostDeliveryCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundablePostDeliveryCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundablePostDeliveryCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundablePostDeliveryCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundablePostDeliveryCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefundablePostDeliveryCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundablePostDeliveryCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundablePostDeliveryCrowdsaleSession struct {
	Contract     *RefundablePostDeliveryCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// RefundablePostDeliveryCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundablePostDeliveryCrowdsaleCallerSession struct {
	Contract *RefundablePostDeliveryCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// RefundablePostDeliveryCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundablePostDeliveryCrowdsaleTransactorSession struct {
	Contract     *RefundablePostDeliveryCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// RefundablePostDeliveryCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundablePostDeliveryCrowdsaleRaw struct {
	Contract *RefundablePostDeliveryCrowdsale // Generic contract binding to access the raw methods on
}

// RefundablePostDeliveryCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundablePostDeliveryCrowdsaleCallerRaw struct {
	Contract *RefundablePostDeliveryCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// RefundablePostDeliveryCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundablePostDeliveryCrowdsaleTransactorRaw struct {
	Contract *RefundablePostDeliveryCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundablePostDeliveryCrowdsale creates a new instance of RefundablePostDeliveryCrowdsale, bound to a specific deployed contract.
func NewRefundablePostDeliveryCrowdsale(address common.Address, backend bind.ContractBackend) (*RefundablePostDeliveryCrowdsale, error) {
	contract, err := bindRefundablePostDeliveryCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsale{RefundablePostDeliveryCrowdsaleCaller: RefundablePostDeliveryCrowdsaleCaller{contract: contract}, RefundablePostDeliveryCrowdsaleTransactor: RefundablePostDeliveryCrowdsaleTransactor{contract: contract}, RefundablePostDeliveryCrowdsaleFilterer: RefundablePostDeliveryCrowdsaleFilterer{contract: contract}}, nil
}

// NewRefundablePostDeliveryCrowdsaleCaller creates a new read-only instance of RefundablePostDeliveryCrowdsale, bound to a specific deployed contract.
func NewRefundablePostDeliveryCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*RefundablePostDeliveryCrowdsaleCaller, error) {
	contract, err := bindRefundablePostDeliveryCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsaleCaller{contract: contract}, nil
}

// NewRefundablePostDeliveryCrowdsaleTransactor creates a new write-only instance of RefundablePostDeliveryCrowdsale, bound to a specific deployed contract.
func NewRefundablePostDeliveryCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundablePostDeliveryCrowdsaleTransactor, error) {
	contract, err := bindRefundablePostDeliveryCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsaleTransactor{contract: contract}, nil
}

// NewRefundablePostDeliveryCrowdsaleFilterer creates a new log filterer instance of RefundablePostDeliveryCrowdsale, bound to a specific deployed contract.
func NewRefundablePostDeliveryCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*RefundablePostDeliveryCrowdsaleFilterer, error) {
	contract, err := bindRefundablePostDeliveryCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsaleFilterer{contract: contract}, nil
}

// bindRefundablePostDeliveryCrowdsale binds a generic wrapper to an already deployed contract.
func bindRefundablePostDeliveryCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundablePostDeliveryCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundablePostDeliveryCrowdsale.Contract.RefundablePostDeliveryCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.RefundablePostDeliveryCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.RefundablePostDeliveryCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefundablePostDeliveryCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.BalanceOf(&_RefundablePostDeliveryCrowdsale.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.BalanceOf(&_RefundablePostDeliveryCrowdsale.CallOpts, account)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.ClosingTime(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.ClosingTime(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Finalized() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Finalized(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) Finalized() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Finalized(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "goal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Goal() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Goal(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) Goal() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Goal(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "goalReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) GoalReached() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.GoalReached(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) GoalReached() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.GoalReached(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) HasClosed() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.HasClosed(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.HasClosed(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) IsOpen() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.IsOpen(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.IsOpen(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.OpeningTime(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.OpeningTime(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Rate() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Rate(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Rate(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Token() (common.Address, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Token(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) Token() (common.Address, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Token(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Wallet() (common.Address, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Wallet(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Wallet(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RefundablePostDeliveryCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.WeiRaised(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.WeiRaised(&_RefundablePostDeliveryCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.BuyTokens(&_RefundablePostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.BuyTokens(&_RefundablePostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactor) ClaimRefund(opts *bind.TransactOpts, refundee common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.contract.Transact(opts, "claimRefund", refundee)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) ClaimRefund(refundee common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.ClaimRefund(&_RefundablePostDeliveryCrowdsale.TransactOpts, refundee)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address refundee) returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorSession) ClaimRefund(refundee common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.ClaimRefund(&_RefundablePostDeliveryCrowdsale.TransactOpts, refundee)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Finalize(&_RefundablePostDeliveryCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Finalize(&_RefundablePostDeliveryCrowdsale.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactor) WithdrawTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.contract.Transact(opts, "withdrawTokens", beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) WithdrawTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.WithdrawTokens(&_RefundablePostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address beneficiary) returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorSession) WithdrawTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.WithdrawTokens(&_RefundablePostDeliveryCrowdsale.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Fallback(&_RefundablePostDeliveryCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Fallback(&_RefundablePostDeliveryCrowdsale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Receive(&_RefundablePostDeliveryCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _RefundablePostDeliveryCrowdsale.Contract.Receive(&_RefundablePostDeliveryCrowdsale.TransactOpts)
}

// RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator is returned from FilterCrowdsaleFinalized and is used to iterate over the raw logs and unpacked data for CrowdsaleFinalized events raised by the RefundablePostDeliveryCrowdsale contract.
type RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator struct {
	Event *RefundablePostDeliveryCrowdsaleCrowdsaleFinalized // Event containing the contract specifics and raw log

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
func (it *RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundablePostDeliveryCrowdsaleCrowdsaleFinalized)
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
		it.Event = new(RefundablePostDeliveryCrowdsaleCrowdsaleFinalized)
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
func (it *RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundablePostDeliveryCrowdsaleCrowdsaleFinalized represents a CrowdsaleFinalized event raised by the RefundablePostDeliveryCrowdsale contract.
type RefundablePostDeliveryCrowdsaleCrowdsaleFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrowdsaleFinalized is a free log retrieval operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) FilterCrowdsaleFinalized(opts *bind.FilterOpts) (*RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator, error) {

	logs, sub, err := _RefundablePostDeliveryCrowdsale.contract.FilterLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsaleCrowdsaleFinalizedIterator{contract: _RefundablePostDeliveryCrowdsale.contract, event: "CrowdsaleFinalized", logs: logs, sub: sub}, nil
}

// WatchCrowdsaleFinalized is a free log subscription operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) WatchCrowdsaleFinalized(opts *bind.WatchOpts, sink chan<- *RefundablePostDeliveryCrowdsaleCrowdsaleFinalized) (event.Subscription, error) {

	logs, sub, err := _RefundablePostDeliveryCrowdsale.contract.WatchLogs(opts, "CrowdsaleFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundablePostDeliveryCrowdsaleCrowdsaleFinalized)
				if err := _RefundablePostDeliveryCrowdsale.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
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

// ParseCrowdsaleFinalized is a log parse operation binding the contract event 0x9270cc390c096600a1c17c44345a1ba689fafd99d97487b10cfccf86cf731836.
//
// Solidity: event CrowdsaleFinalized()
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) ParseCrowdsaleFinalized(log types.Log) (*RefundablePostDeliveryCrowdsaleCrowdsaleFinalized, error) {
	event := new(RefundablePostDeliveryCrowdsaleCrowdsaleFinalized)
	if err := _RefundablePostDeliveryCrowdsale.contract.UnpackLog(event, "CrowdsaleFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the RefundablePostDeliveryCrowdsale contract.
type RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended)
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
func (it *RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the RefundablePostDeliveryCrowdsale contract.
type RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _RefundablePostDeliveryCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtendedIterator{contract: _RefundablePostDeliveryCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _RefundablePostDeliveryCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended)
				if err := _RefundablePostDeliveryCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(RefundablePostDeliveryCrowdsaleTimedCrowdsaleExtended)
	if err := _RefundablePostDeliveryCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefundablePostDeliveryCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the RefundablePostDeliveryCrowdsale contract.
type RefundablePostDeliveryCrowdsaleTokensPurchasedIterator struct {
	Event *RefundablePostDeliveryCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *RefundablePostDeliveryCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundablePostDeliveryCrowdsaleTokensPurchased)
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
		it.Event = new(RefundablePostDeliveryCrowdsaleTokensPurchased)
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
func (it *RefundablePostDeliveryCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundablePostDeliveryCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundablePostDeliveryCrowdsaleTokensPurchased represents a TokensPurchased event raised by the RefundablePostDeliveryCrowdsale contract.
type RefundablePostDeliveryCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*RefundablePostDeliveryCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundablePostDeliveryCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &RefundablePostDeliveryCrowdsaleTokensPurchasedIterator{contract: _RefundablePostDeliveryCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *RefundablePostDeliveryCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundablePostDeliveryCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundablePostDeliveryCrowdsaleTokensPurchased)
				if err := _RefundablePostDeliveryCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_RefundablePostDeliveryCrowdsale *RefundablePostDeliveryCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*RefundablePostDeliveryCrowdsaleTokensPurchased, error) {
	event := new(RefundablePostDeliveryCrowdsaleTokensPurchased)
	if err := _RefundablePostDeliveryCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208a68b6dcf2bf162ea8b6ae699adc982a5b3a6869f4f1cb0e8ce1dfd11f500b3064736f6c63430007060033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204d96e73385ece43f5b505689162b5ccdcc2051f546d3b9daeedc826592624ea864736f6c63430007060033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SecondaryABI is the input ABI used to generate the binding from.
const SecondaryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"PrimaryTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"primary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"transferPrimary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SecondaryFuncSigs maps the 4-byte function signature to its string representation.
var SecondaryFuncSigs = map[string]string{
	"c6dbdf61": "primary()",
	"2348238c": "transferPrimary(address)",
}

// Secondary is an auto generated Go binding around an Ethereum contract.
type Secondary struct {
	SecondaryCaller     // Read-only binding to the contract
	SecondaryTransactor // Write-only binding to the contract
	SecondaryFilterer   // Log filterer for contract events
}

// SecondaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecondaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecondaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecondaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecondaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecondaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecondarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecondarySession struct {
	Contract     *Secondary        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecondaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecondaryCallerSession struct {
	Contract *SecondaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SecondaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecondaryTransactorSession struct {
	Contract     *SecondaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SecondaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecondaryRaw struct {
	Contract *Secondary // Generic contract binding to access the raw methods on
}

// SecondaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecondaryCallerRaw struct {
	Contract *SecondaryCaller // Generic read-only contract binding to access the raw methods on
}

// SecondaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecondaryTransactorRaw struct {
	Contract *SecondaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecondary creates a new instance of Secondary, bound to a specific deployed contract.
func NewSecondary(address common.Address, backend bind.ContractBackend) (*Secondary, error) {
	contract, err := bindSecondary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Secondary{SecondaryCaller: SecondaryCaller{contract: contract}, SecondaryTransactor: SecondaryTransactor{contract: contract}, SecondaryFilterer: SecondaryFilterer{contract: contract}}, nil
}

// NewSecondaryCaller creates a new read-only instance of Secondary, bound to a specific deployed contract.
func NewSecondaryCaller(address common.Address, caller bind.ContractCaller) (*SecondaryCaller, error) {
	contract, err := bindSecondary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecondaryCaller{contract: contract}, nil
}

// NewSecondaryTransactor creates a new write-only instance of Secondary, bound to a specific deployed contract.
func NewSecondaryTransactor(address common.Address, transactor bind.ContractTransactor) (*SecondaryTransactor, error) {
	contract, err := bindSecondary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecondaryTransactor{contract: contract}, nil
}

// NewSecondaryFilterer creates a new log filterer instance of Secondary, bound to a specific deployed contract.
func NewSecondaryFilterer(address common.Address, filterer bind.ContractFilterer) (*SecondaryFilterer, error) {
	contract, err := bindSecondary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecondaryFilterer{contract: contract}, nil
}

// bindSecondary binds a generic wrapper to an already deployed contract.
func bindSecondary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SecondaryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Secondary *SecondaryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Secondary.Contract.SecondaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Secondary *SecondaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Secondary.Contract.SecondaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Secondary *SecondaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Secondary.Contract.SecondaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Secondary *SecondaryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Secondary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Secondary *SecondaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Secondary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Secondary *SecondaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Secondary.Contract.contract.Transact(opts, method, params...)
}

// Primary is a free data retrieval call binding the contract method 0xc6dbdf61.
//
// Solidity: function primary() view returns(address)
func (_Secondary *SecondaryCaller) Primary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Secondary.contract.Call(opts, &out, "primary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Primary is a free data retrieval call binding the contract method 0xc6dbdf61.
//
// Solidity: function primary() view returns(address)
func (_Secondary *SecondarySession) Primary() (common.Address, error) {
	return _Secondary.Contract.Primary(&_Secondary.CallOpts)
}

// Primary is a free data retrieval call binding the contract method 0xc6dbdf61.
//
// Solidity: function primary() view returns(address)
func (_Secondary *SecondaryCallerSession) Primary() (common.Address, error) {
	return _Secondary.Contract.Primary(&_Secondary.CallOpts)
}

// TransferPrimary is a paid mutator transaction binding the contract method 0x2348238c.
//
// Solidity: function transferPrimary(address recipient) returns()
func (_Secondary *SecondaryTransactor) TransferPrimary(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Secondary.contract.Transact(opts, "transferPrimary", recipient)
}

// TransferPrimary is a paid mutator transaction binding the contract method 0x2348238c.
//
// Solidity: function transferPrimary(address recipient) returns()
func (_Secondary *SecondarySession) TransferPrimary(recipient common.Address) (*types.Transaction, error) {
	return _Secondary.Contract.TransferPrimary(&_Secondary.TransactOpts, recipient)
}

// TransferPrimary is a paid mutator transaction binding the contract method 0x2348238c.
//
// Solidity: function transferPrimary(address recipient) returns()
func (_Secondary *SecondaryTransactorSession) TransferPrimary(recipient common.Address) (*types.Transaction, error) {
	return _Secondary.Contract.TransferPrimary(&_Secondary.TransactOpts, recipient)
}

// SecondaryPrimaryTransferredIterator is returned from FilterPrimaryTransferred and is used to iterate over the raw logs and unpacked data for PrimaryTransferred events raised by the Secondary contract.
type SecondaryPrimaryTransferredIterator struct {
	Event *SecondaryPrimaryTransferred // Event containing the contract specifics and raw log

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
func (it *SecondaryPrimaryTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecondaryPrimaryTransferred)
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
		it.Event = new(SecondaryPrimaryTransferred)
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
func (it *SecondaryPrimaryTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecondaryPrimaryTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecondaryPrimaryTransferred represents a PrimaryTransferred event raised by the Secondary contract.
type SecondaryPrimaryTransferred struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimaryTransferred is a free log retrieval operation binding the contract event 0x4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9.
//
// Solidity: event PrimaryTransferred(address recipient)
func (_Secondary *SecondaryFilterer) FilterPrimaryTransferred(opts *bind.FilterOpts) (*SecondaryPrimaryTransferredIterator, error) {

	logs, sub, err := _Secondary.contract.FilterLogs(opts, "PrimaryTransferred")
	if err != nil {
		return nil, err
	}
	return &SecondaryPrimaryTransferredIterator{contract: _Secondary.contract, event: "PrimaryTransferred", logs: logs, sub: sub}, nil
}

// WatchPrimaryTransferred is a free log subscription operation binding the contract event 0x4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9.
//
// Solidity: event PrimaryTransferred(address recipient)
func (_Secondary *SecondaryFilterer) WatchPrimaryTransferred(opts *bind.WatchOpts, sink chan<- *SecondaryPrimaryTransferred) (event.Subscription, error) {

	logs, sub, err := _Secondary.contract.WatchLogs(opts, "PrimaryTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecondaryPrimaryTransferred)
				if err := _Secondary.contract.UnpackLog(event, "PrimaryTransferred", log); err != nil {
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

// ParsePrimaryTransferred is a log parse operation binding the contract event 0x4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9.
//
// Solidity: event PrimaryTransferred(address recipient)
func (_Secondary *SecondaryFilterer) ParsePrimaryTransferred(log types.Log) (*SecondaryPrimaryTransferred, error) {
	event := new(SecondaryPrimaryTransferred)
	if err := _Secondary.contract.UnpackLog(event, "PrimaryTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedCrowdsaleABI is the input ABI used to generate the binding from.
const TimedCrowdsaleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractLBCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// TimedCrowdsaleFuncSigs maps the 4-byte function signature to its string representation.
var TimedCrowdsaleFuncSigs = map[string]string{
	"ec8ac4d8": "buyTokens(address)",
	"4b6753bc": "closingTime()",
	"1515bc2b": "hasClosed()",
	"47535d7b": "isOpen()",
	"b7a8807c": "openingTime()",
	"2c4e722e": "rate()",
	"fc0c546a": "token()",
	"521eb273": "wallet()",
	"4042b66f": "weiRaised()",
}

// TimedCrowdsale is an auto generated Go binding around an Ethereum contract.
type TimedCrowdsale struct {
	TimedCrowdsaleCaller     // Read-only binding to the contract
	TimedCrowdsaleTransactor // Write-only binding to the contract
	TimedCrowdsaleFilterer   // Log filterer for contract events
}

// TimedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedCrowdsaleSession struct {
	Contract     *TimedCrowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedCrowdsaleCallerSession struct {
	Contract *TimedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TimedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedCrowdsaleTransactorSession struct {
	Contract     *TimedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TimedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedCrowdsaleRaw struct {
	Contract *TimedCrowdsale // Generic contract binding to access the raw methods on
}

// TimedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedCrowdsaleCallerRaw struct {
	Contract *TimedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// TimedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedCrowdsaleTransactorRaw struct {
	Contract *TimedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimedCrowdsale creates a new instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsale(address common.Address, backend bind.ContractBackend) (*TimedCrowdsale, error) {
	contract, err := bindTimedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsale{TimedCrowdsaleCaller: TimedCrowdsaleCaller{contract: contract}, TimedCrowdsaleTransactor: TimedCrowdsaleTransactor{contract: contract}, TimedCrowdsaleFilterer: TimedCrowdsaleFilterer{contract: contract}}, nil
}

// NewTimedCrowdsaleCaller creates a new read-only instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*TimedCrowdsaleCaller, error) {
	contract, err := bindTimedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleCaller{contract: contract}, nil
}

// NewTimedCrowdsaleTransactor creates a new write-only instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedCrowdsaleTransactor, error) {
	contract, err := bindTimedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTransactor{contract: contract}, nil
}

// NewTimedCrowdsaleFilterer creates a new log filterer instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*TimedCrowdsaleFilterer, error) {
	contract, err := bindTimedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleFilterer{contract: contract}, nil
}

// bindTimedCrowdsale binds a generic wrapper to an already deployed contract.
func bindTimedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsale.Contract.TimedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TimedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TimedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsale *TimedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsale *TimedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsale *TimedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.ClosingTime(&_TimedCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.ClosingTime(&_TimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleSession) HasClosed() (bool, error) {
	return _TimedCrowdsale.Contract.HasClosed(&_TimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _TimedCrowdsale.Contract.HasClosed(&_TimedCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleSession) IsOpen() (bool, error) {
	return _TimedCrowdsale.Contract.IsOpen(&_TimedCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _TimedCrowdsale.Contract.IsOpen(&_TimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.OpeningTime(&_TimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.OpeningTime(&_TimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) Rate() (*big.Int, error) {
	return _TimedCrowdsale.Contract.Rate(&_TimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _TimedCrowdsale.Contract.Rate(&_TimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Token() (common.Address, error) {
	return _TimedCrowdsale.Contract.Token(&_TimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _TimedCrowdsale.Contract.Token(&_TimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Wallet() (common.Address, error) {
	return _TimedCrowdsale.Contract.Wallet(&_TimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _TimedCrowdsale.Contract.Wallet(&_TimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsale.Contract.WeiRaised(&_TimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsale.Contract.WeiRaised(&_TimedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.BuyTokens(&_TimedCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.BuyTokens(&_TimedCrowdsale.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Fallback(&_TimedCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Fallback(&_TimedCrowdsale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Receive(&_TimedCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Receive(&_TimedCrowdsale.TransactOpts)
}

// TimedCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the TimedCrowdsale contract.
type TimedCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *TimedCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(TimedCrowdsaleTimedCrowdsaleExtended)
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
func (it *TimedCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the TimedCrowdsale contract.
type TimedCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*TimedCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTimedCrowdsaleExtendedIterator{contract: _TimedCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleTimedCrowdsaleExtended)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*TimedCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(TimedCrowdsaleTimedCrowdsaleExtended)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the TimedCrowdsale contract.
type TimedCrowdsaleTokensPurchasedIterator struct {
	Event *TimedCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleTokensPurchased)
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
		it.Event = new(TimedCrowdsaleTokensPurchased)
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
func (it *TimedCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleTokensPurchased represents a TokensPurchased event raised by the TimedCrowdsale contract.
type TimedCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*TimedCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTokensPurchasedIterator{contract: _TimedCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleTokensPurchased)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*TimedCrowdsaleTokensPurchased, error) {
	event := new(TimedCrowdsaleTokensPurchased)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstableTokenVaultABI is the input ABI used to generate the binding from.
const UnstableTokenVaultABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"PrimaryTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"primary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"transferPrimary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UnstableTokenVaultFuncSigs maps the 4-byte function signature to its string representation.
var UnstableTokenVaultFuncSigs = map[string]string{
	"c6dbdf61": "primary()",
	"beabacc8": "transfer(address,address,uint256)",
	"2348238c": "transferPrimary(address)",
}

// UnstableTokenVaultBin is the compiled bytecode used for deploying new contracts.
var UnstableTokenVaultBin = "0x608060405234801561001057600080fd5b50600061001b610076565b600080546001600160a01b0383166001600160a01b0319909116811790915560408051918252519192507f4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9919081900360200190a15061007a565b3390565b61033a806100896000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632348238c14610046578063beabacc81461006e578063c6dbdf61146100a4575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100c8565b005b61006c6004803603606081101561008457600080fd5b506001600160a01b038135811691602081013590911690604001356101ba565b6100ac61029b565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b03166100dc6102aa565b6001600160a01b0316146101215760405162461bcd60e51b815260040180806020018281038252602c8152602001806102d9602c913960400191505060405180910390fd5b6001600160a01b0381166101665760405162461bcd60e51b815260040180806020018281038252602a8152602001806102af602a913960400191505060405180910390fd5b600080546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d99181900360200190a150565b6000546001600160a01b03166101ce6102aa565b6001600160a01b0316146102135760405162461bcd60e51b815260040180806020018281038252602c8152602001806102d9602c913960400191505060405180910390fd5b826001600160a01b031663a9059cbb83836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561026a57600080fd5b505af115801561027e573d6000803e3d6000fd5b505050506040513d602081101561029457600080fd5b5050505050565b6000546001600160a01b031690565b339056fe5365636f6e646172793a206e6577207072696d61727920697320746865207a65726f20616464726573735365636f6e646172793a2063616c6c6572206973206e6f7420746865207072696d617279206163636f756e74a2646970667358221220cdd24c05d08289961d37f26f404cdfc4688c42bbea0ae52bbe09cd50a126f02464736f6c63430007060033"

// DeployUnstableTokenVault deploys a new Ethereum contract, binding an instance of UnstableTokenVault to it.
func DeployUnstableTokenVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnstableTokenVault, error) {
	parsed, err := abi.JSON(strings.NewReader(UnstableTokenVaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnstableTokenVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnstableTokenVault{UnstableTokenVaultCaller: UnstableTokenVaultCaller{contract: contract}, UnstableTokenVaultTransactor: UnstableTokenVaultTransactor{contract: contract}, UnstableTokenVaultFilterer: UnstableTokenVaultFilterer{contract: contract}}, nil
}

// UnstableTokenVault is an auto generated Go binding around an Ethereum contract.
type UnstableTokenVault struct {
	UnstableTokenVaultCaller     // Read-only binding to the contract
	UnstableTokenVaultTransactor // Write-only binding to the contract
	UnstableTokenVaultFilterer   // Log filterer for contract events
}

// UnstableTokenVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnstableTokenVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstableTokenVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnstableTokenVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstableTokenVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnstableTokenVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstableTokenVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnstableTokenVaultSession struct {
	Contract     *UnstableTokenVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnstableTokenVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnstableTokenVaultCallerSession struct {
	Contract *UnstableTokenVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UnstableTokenVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnstableTokenVaultTransactorSession struct {
	Contract     *UnstableTokenVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UnstableTokenVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnstableTokenVaultRaw struct {
	Contract *UnstableTokenVault // Generic contract binding to access the raw methods on
}

// UnstableTokenVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnstableTokenVaultCallerRaw struct {
	Contract *UnstableTokenVaultCaller // Generic read-only contract binding to access the raw methods on
}

// UnstableTokenVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnstableTokenVaultTransactorRaw struct {
	Contract *UnstableTokenVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnstableTokenVault creates a new instance of UnstableTokenVault, bound to a specific deployed contract.
func NewUnstableTokenVault(address common.Address, backend bind.ContractBackend) (*UnstableTokenVault, error) {
	contract, err := bindUnstableTokenVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnstableTokenVault{UnstableTokenVaultCaller: UnstableTokenVaultCaller{contract: contract}, UnstableTokenVaultTransactor: UnstableTokenVaultTransactor{contract: contract}, UnstableTokenVaultFilterer: UnstableTokenVaultFilterer{contract: contract}}, nil
}

// NewUnstableTokenVaultCaller creates a new read-only instance of UnstableTokenVault, bound to a specific deployed contract.
func NewUnstableTokenVaultCaller(address common.Address, caller bind.ContractCaller) (*UnstableTokenVaultCaller, error) {
	contract, err := bindUnstableTokenVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnstableTokenVaultCaller{contract: contract}, nil
}

// NewUnstableTokenVaultTransactor creates a new write-only instance of UnstableTokenVault, bound to a specific deployed contract.
func NewUnstableTokenVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*UnstableTokenVaultTransactor, error) {
	contract, err := bindUnstableTokenVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnstableTokenVaultTransactor{contract: contract}, nil
}

// NewUnstableTokenVaultFilterer creates a new log filterer instance of UnstableTokenVault, bound to a specific deployed contract.
func NewUnstableTokenVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*UnstableTokenVaultFilterer, error) {
	contract, err := bindUnstableTokenVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnstableTokenVaultFilterer{contract: contract}, nil
}

// bindUnstableTokenVault binds a generic wrapper to an already deployed contract.
func bindUnstableTokenVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnstableTokenVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnstableTokenVault *UnstableTokenVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnstableTokenVault.Contract.UnstableTokenVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnstableTokenVault *UnstableTokenVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.UnstableTokenVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnstableTokenVault *UnstableTokenVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.UnstableTokenVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnstableTokenVault *UnstableTokenVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnstableTokenVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnstableTokenVault *UnstableTokenVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnstableTokenVault *UnstableTokenVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.contract.Transact(opts, method, params...)
}

// Primary is a free data retrieval call binding the contract method 0xc6dbdf61.
//
// Solidity: function primary() view returns(address)
func (_UnstableTokenVault *UnstableTokenVaultCaller) Primary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnstableTokenVault.contract.Call(opts, &out, "primary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Primary is a free data retrieval call binding the contract method 0xc6dbdf61.
//
// Solidity: function primary() view returns(address)
func (_UnstableTokenVault *UnstableTokenVaultSession) Primary() (common.Address, error) {
	return _UnstableTokenVault.Contract.Primary(&_UnstableTokenVault.CallOpts)
}

// Primary is a free data retrieval call binding the contract method 0xc6dbdf61.
//
// Solidity: function primary() view returns(address)
func (_UnstableTokenVault *UnstableTokenVaultCallerSession) Primary() (common.Address, error) {
	return _UnstableTokenVault.Contract.Primary(&_UnstableTokenVault.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_UnstableTokenVault *UnstableTokenVaultTransactor) Transfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnstableTokenVault.contract.Transact(opts, "transfer", token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_UnstableTokenVault *UnstableTokenVaultSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.Transfer(&_UnstableTokenVault.TransactOpts, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_UnstableTokenVault *UnstableTokenVaultTransactorSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.Transfer(&_UnstableTokenVault.TransactOpts, token, to, amount)
}

// TransferPrimary is a paid mutator transaction binding the contract method 0x2348238c.
//
// Solidity: function transferPrimary(address recipient) returns()
func (_UnstableTokenVault *UnstableTokenVaultTransactor) TransferPrimary(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _UnstableTokenVault.contract.Transact(opts, "transferPrimary", recipient)
}

// TransferPrimary is a paid mutator transaction binding the contract method 0x2348238c.
//
// Solidity: function transferPrimary(address recipient) returns()
func (_UnstableTokenVault *UnstableTokenVaultSession) TransferPrimary(recipient common.Address) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.TransferPrimary(&_UnstableTokenVault.TransactOpts, recipient)
}

// TransferPrimary is a paid mutator transaction binding the contract method 0x2348238c.
//
// Solidity: function transferPrimary(address recipient) returns()
func (_UnstableTokenVault *UnstableTokenVaultTransactorSession) TransferPrimary(recipient common.Address) (*types.Transaction, error) {
	return _UnstableTokenVault.Contract.TransferPrimary(&_UnstableTokenVault.TransactOpts, recipient)
}

// UnstableTokenVaultPrimaryTransferredIterator is returned from FilterPrimaryTransferred and is used to iterate over the raw logs and unpacked data for PrimaryTransferred events raised by the UnstableTokenVault contract.
type UnstableTokenVaultPrimaryTransferredIterator struct {
	Event *UnstableTokenVaultPrimaryTransferred // Event containing the contract specifics and raw log

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
func (it *UnstableTokenVaultPrimaryTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstableTokenVaultPrimaryTransferred)
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
		it.Event = new(UnstableTokenVaultPrimaryTransferred)
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
func (it *UnstableTokenVaultPrimaryTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstableTokenVaultPrimaryTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstableTokenVaultPrimaryTransferred represents a PrimaryTransferred event raised by the UnstableTokenVault contract.
type UnstableTokenVaultPrimaryTransferred struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimaryTransferred is a free log retrieval operation binding the contract event 0x4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9.
//
// Solidity: event PrimaryTransferred(address recipient)
func (_UnstableTokenVault *UnstableTokenVaultFilterer) FilterPrimaryTransferred(opts *bind.FilterOpts) (*UnstableTokenVaultPrimaryTransferredIterator, error) {

	logs, sub, err := _UnstableTokenVault.contract.FilterLogs(opts, "PrimaryTransferred")
	if err != nil {
		return nil, err
	}
	return &UnstableTokenVaultPrimaryTransferredIterator{contract: _UnstableTokenVault.contract, event: "PrimaryTransferred", logs: logs, sub: sub}, nil
}

// WatchPrimaryTransferred is a free log subscription operation binding the contract event 0x4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9.
//
// Solidity: event PrimaryTransferred(address recipient)
func (_UnstableTokenVault *UnstableTokenVaultFilterer) WatchPrimaryTransferred(opts *bind.WatchOpts, sink chan<- *UnstableTokenVaultPrimaryTransferred) (event.Subscription, error) {

	logs, sub, err := _UnstableTokenVault.contract.WatchLogs(opts, "PrimaryTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstableTokenVaultPrimaryTransferred)
				if err := _UnstableTokenVault.contract.UnpackLog(event, "PrimaryTransferred", log); err != nil {
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

// ParsePrimaryTransferred is a log parse operation binding the contract event 0x4101e71e974f68df5e9730cc223280b41654676bbb052cdcc735c3337e64d2d9.
//
// Solidity: event PrimaryTransferred(address recipient)
func (_UnstableTokenVault *UnstableTokenVaultFilterer) ParsePrimaryTransferred(log types.Log) (*UnstableTokenVaultPrimaryTransferred, error) {
	event := new(UnstableTokenVaultPrimaryTransferred)
	if err := _UnstableTokenVault.contract.UnpackLog(event, "PrimaryTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
