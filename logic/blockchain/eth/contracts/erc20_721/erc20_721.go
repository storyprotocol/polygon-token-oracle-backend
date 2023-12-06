// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_721

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc20721MetaData contains all meta data concerning the Erc20721 contract.
var Erc20721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc20721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20721MetaData.ABI instead.
var Erc20721ABI = Erc20721MetaData.ABI

// Erc20721 is an auto generated Go binding around an Ethereum contract.
type Erc20721 struct {
	Erc20721Caller     // Read-only binding to the contract
	Erc20721Transactor // Write-only binding to the contract
	Erc20721Filterer   // Log filterer for contract events
}

// Erc20721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20721Session struct {
	Contract     *Erc20721         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20721CallerSession struct {
	Contract *Erc20721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Erc20721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20721TransactorSession struct {
	Contract     *Erc20721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc20721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20721Raw struct {
	Contract *Erc20721 // Generic contract binding to access the raw methods on
}

// Erc20721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20721CallerRaw struct {
	Contract *Erc20721Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20721TransactorRaw struct {
	Contract *Erc20721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20721 creates a new instance of Erc20721, bound to a specific deployed contract.
func NewErc20721(address common.Address, backend bind.ContractBackend) (*Erc20721, error) {
	contract, err := bindErc20721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20721{Erc20721Caller: Erc20721Caller{contract: contract}, Erc20721Transactor: Erc20721Transactor{contract: contract}, Erc20721Filterer: Erc20721Filterer{contract: contract}}, nil
}

// NewErc20721Caller creates a new read-only instance of Erc20721, bound to a specific deployed contract.
func NewErc20721Caller(address common.Address, caller bind.ContractCaller) (*Erc20721Caller, error) {
	contract, err := bindErc20721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20721Caller{contract: contract}, nil
}

// NewErc20721Transactor creates a new write-only instance of Erc20721, bound to a specific deployed contract.
func NewErc20721Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20721Transactor, error) {
	contract, err := bindErc20721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20721Transactor{contract: contract}, nil
}

// NewErc20721Filterer creates a new log filterer instance of Erc20721, bound to a specific deployed contract.
func NewErc20721Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20721Filterer, error) {
	contract, err := bindErc20721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20721Filterer{contract: contract}, nil
}

// bindErc20721 binds a generic wrapper to an already deployed contract.
func bindErc20721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20721 *Erc20721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20721.Contract.Erc20721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20721 *Erc20721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20721.Contract.Erc20721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20721 *Erc20721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20721.Contract.Erc20721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20721 *Erc20721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20721 *Erc20721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20721 *Erc20721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc20721 *Erc20721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc20721 *Erc20721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc20721.Contract.BalanceOf(&_Erc20721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc20721 *Erc20721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc20721.Contract.BalanceOf(&_Erc20721.CallOpts, owner)
}
