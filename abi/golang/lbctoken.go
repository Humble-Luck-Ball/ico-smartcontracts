// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lbctoken

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
