// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygon_token

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

// PolygonTokenMetaData contains all meta data concerning the PolygonToken contract.
var PolygonTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOwnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"callbackFunctionSignature\",\"type\":\"bytes4\"}],\"name\":\"PolygonTokenBalanceRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"handleCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceThreshold\",\"type\":\"uint256\"}],\"name\":\"requestAsyncCall\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdToRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRequestCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolygonTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonTokenMetaData.ABI instead.
var PolygonTokenABI = PolygonTokenMetaData.ABI

// PolygonToken is an auto generated Go binding around an Ethereum contract.
type PolygonToken struct {
	PolygonTokenCaller     // Read-only binding to the contract
	PolygonTokenTransactor // Write-only binding to the contract
	PolygonTokenFilterer   // Log filterer for contract events
}

// PolygonTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonTokenSession struct {
	Contract     *PolygonToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolygonTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonTokenCallerSession struct {
	Contract *PolygonTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PolygonTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonTokenTransactorSession struct {
	Contract     *PolygonTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PolygonTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonTokenRaw struct {
	Contract *PolygonToken // Generic contract binding to access the raw methods on
}

// PolygonTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonTokenCallerRaw struct {
	Contract *PolygonTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonTokenTransactorRaw struct {
	Contract *PolygonTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonToken creates a new instance of PolygonToken, bound to a specific deployed contract.
func NewPolygonToken(address common.Address, backend bind.ContractBackend) (*PolygonToken, error) {
	contract, err := bindPolygonToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonToken{PolygonTokenCaller: PolygonTokenCaller{contract: contract}, PolygonTokenTransactor: PolygonTokenTransactor{contract: contract}, PolygonTokenFilterer: PolygonTokenFilterer{contract: contract}}, nil
}

// NewPolygonTokenCaller creates a new read-only instance of PolygonToken, bound to a specific deployed contract.
func NewPolygonTokenCaller(address common.Address, caller bind.ContractCaller) (*PolygonTokenCaller, error) {
	contract, err := bindPolygonToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonTokenCaller{contract: contract}, nil
}

// NewPolygonTokenTransactor creates a new write-only instance of PolygonToken, bound to a specific deployed contract.
func NewPolygonTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonTokenTransactor, error) {
	contract, err := bindPolygonToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonTokenTransactor{contract: contract}, nil
}

// NewPolygonTokenFilterer creates a new log filterer instance of PolygonToken, bound to a specific deployed contract.
func NewPolygonTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonTokenFilterer, error) {
	contract, err := bindPolygonToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonTokenFilterer{contract: contract}, nil
}

// bindPolygonToken binds a generic wrapper to an already deployed contract.
func bindPolygonToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolygonTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonToken *PolygonTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonToken.Contract.PolygonTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonToken *PolygonTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonToken.Contract.PolygonTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonToken *PolygonTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonToken.Contract.PolygonTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonToken *PolygonTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonToken *PolygonTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonToken *PolygonTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonToken.Contract.contract.Transact(opts, method, params...)
}

// RequestIdToRequest is a free data retrieval call binding the contract method 0x01e5bbb9.
//
// Solidity: function requestIdToRequest(bytes32 ) view returns(address requester, address tokenAddress, address tokenOwnerAddress, uint256 balanceThreshold, bool isRequestCompleted, bool exists)
func (_PolygonToken *PolygonTokenCaller) RequestIdToRequest(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Requester          common.Address
	TokenAddress       common.Address
	TokenOwnerAddress  common.Address
	BalanceThreshold   *big.Int
	IsRequestCompleted bool
	Exists             bool
}, error) {
	var out []interface{}
	err := _PolygonToken.contract.Call(opts, &out, "requestIdToRequest", arg0)

	outstruct := new(struct {
		Requester          common.Address
		TokenAddress       common.Address
		TokenOwnerAddress  common.Address
		BalanceThreshold   *big.Int
		IsRequestCompleted bool
		Exists             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenOwnerAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.BalanceThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsRequestCompleted = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Exists = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// RequestIdToRequest is a free data retrieval call binding the contract method 0x01e5bbb9.
//
// Solidity: function requestIdToRequest(bytes32 ) view returns(address requester, address tokenAddress, address tokenOwnerAddress, uint256 balanceThreshold, bool isRequestCompleted, bool exists)
func (_PolygonToken *PolygonTokenSession) RequestIdToRequest(arg0 [32]byte) (struct {
	Requester          common.Address
	TokenAddress       common.Address
	TokenOwnerAddress  common.Address
	BalanceThreshold   *big.Int
	IsRequestCompleted bool
	Exists             bool
}, error) {
	return _PolygonToken.Contract.RequestIdToRequest(&_PolygonToken.CallOpts, arg0)
}

// RequestIdToRequest is a free data retrieval call binding the contract method 0x01e5bbb9.
//
// Solidity: function requestIdToRequest(bytes32 ) view returns(address requester, address tokenAddress, address tokenOwnerAddress, uint256 balanceThreshold, bool isRequestCompleted, bool exists)
func (_PolygonToken *PolygonTokenCallerSession) RequestIdToRequest(arg0 [32]byte) (struct {
	Requester          common.Address
	TokenAddress       common.Address
	TokenOwnerAddress  common.Address
	BalanceThreshold   *big.Int
	IsRequestCompleted bool
	Exists             bool
}, error) {
	return _PolygonToken.Contract.RequestIdToRequest(&_PolygonToken.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(string)
func (_PolygonToken *PolygonTokenCaller) Results(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _PolygonToken.contract.Call(opts, &out, "results", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(string)
func (_PolygonToken *PolygonTokenSession) Results(arg0 [32]byte) (string, error) {
	return _PolygonToken.Contract.Results(&_PolygonToken.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(string)
func (_PolygonToken *PolygonTokenCallerSession) Results(arg0 [32]byte) (string, error) {
	return _PolygonToken.Contract.Results(&_PolygonToken.CallOpts, arg0)
}

// HandleCallback is a paid mutator transaction binding the contract method 0x3d0aa6e5.
//
// Solidity: function handleCallback(bytes32 requestId, uint256 balance) returns()
func (_PolygonToken *PolygonTokenTransactor) HandleCallback(opts *bind.TransactOpts, requestId [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _PolygonToken.contract.Transact(opts, "handleCallback", requestId, balance)
}

// HandleCallback is a paid mutator transaction binding the contract method 0x3d0aa6e5.
//
// Solidity: function handleCallback(bytes32 requestId, uint256 balance) returns()
func (_PolygonToken *PolygonTokenSession) HandleCallback(requestId [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _PolygonToken.Contract.HandleCallback(&_PolygonToken.TransactOpts, requestId, balance)
}

// HandleCallback is a paid mutator transaction binding the contract method 0x3d0aa6e5.
//
// Solidity: function handleCallback(bytes32 requestId, uint256 balance) returns()
func (_PolygonToken *PolygonTokenTransactorSession) HandleCallback(requestId [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _PolygonToken.Contract.HandleCallback(&_PolygonToken.TransactOpts, requestId, balance)
}

// RequestAsyncCall is a paid mutator transaction binding the contract method 0x738fd8c5.
//
// Solidity: function requestAsyncCall(address tokenAddress, address tokenOwnerAddress, uint256 balanceThreshold) returns(bytes32 requestId)
func (_PolygonToken *PolygonTokenTransactor) RequestAsyncCall(opts *bind.TransactOpts, tokenAddress common.Address, tokenOwnerAddress common.Address, balanceThreshold *big.Int) (*types.Transaction, error) {
	return _PolygonToken.contract.Transact(opts, "requestAsyncCall", tokenAddress, tokenOwnerAddress, balanceThreshold)
}

// RequestAsyncCall is a paid mutator transaction binding the contract method 0x738fd8c5.
//
// Solidity: function requestAsyncCall(address tokenAddress, address tokenOwnerAddress, uint256 balanceThreshold) returns(bytes32 requestId)
func (_PolygonToken *PolygonTokenSession) RequestAsyncCall(tokenAddress common.Address, tokenOwnerAddress common.Address, balanceThreshold *big.Int) (*types.Transaction, error) {
	return _PolygonToken.Contract.RequestAsyncCall(&_PolygonToken.TransactOpts, tokenAddress, tokenOwnerAddress, balanceThreshold)
}

// RequestAsyncCall is a paid mutator transaction binding the contract method 0x738fd8c5.
//
// Solidity: function requestAsyncCall(address tokenAddress, address tokenOwnerAddress, uint256 balanceThreshold) returns(bytes32 requestId)
func (_PolygonToken *PolygonTokenTransactorSession) RequestAsyncCall(tokenAddress common.Address, tokenOwnerAddress common.Address, balanceThreshold *big.Int) (*types.Transaction, error) {
	return _PolygonToken.Contract.RequestAsyncCall(&_PolygonToken.TransactOpts, tokenAddress, tokenOwnerAddress, balanceThreshold)
}

// PolygonTokenPolygonTokenBalanceRequestIterator is returned from FilterPolygonTokenBalanceRequest and is used to iterate over the raw logs and unpacked data for PolygonTokenBalanceRequest events raised by the PolygonToken contract.
type PolygonTokenPolygonTokenBalanceRequestIterator struct {
	Event *PolygonTokenPolygonTokenBalanceRequest // Event containing the contract specifics and raw log

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
func (it *PolygonTokenPolygonTokenBalanceRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonTokenPolygonTokenBalanceRequest)
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
		it.Event = new(PolygonTokenPolygonTokenBalanceRequest)
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
func (it *PolygonTokenPolygonTokenBalanceRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonTokenPolygonTokenBalanceRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonTokenPolygonTokenBalanceRequest represents a PolygonTokenBalanceRequest event raised by the PolygonToken contract.
type PolygonTokenPolygonTokenBalanceRequest struct {
	RequestId                 [32]byte
	Requester                 common.Address
	TokenAddress              common.Address
	TokenOwnerAddress         common.Address
	CallbackAddr              common.Address
	CallbackFunctionSignature [4]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPolygonTokenBalanceRequest is a free log retrieval operation binding the contract event 0x97462a831b3a28e48c1bc63ec183d4a9c2ab47e01d7a1cdbefdfad544723646a.
//
// Solidity: event PolygonTokenBalanceRequest(bytes32 indexed requestId, address indexed requester, address tokenAddress, address tokenOwnerAddress, address callbackAddr, bytes4 callbackFunctionSignature)
func (_PolygonToken *PolygonTokenFilterer) FilterPolygonTokenBalanceRequest(opts *bind.FilterOpts, requestId [][32]byte, requester []common.Address) (*PolygonTokenPolygonTokenBalanceRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PolygonToken.contract.FilterLogs(opts, "PolygonTokenBalanceRequest", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PolygonTokenPolygonTokenBalanceRequestIterator{contract: _PolygonToken.contract, event: "PolygonTokenBalanceRequest", logs: logs, sub: sub}, nil
}

// WatchPolygonTokenBalanceRequest is a free log subscription operation binding the contract event 0x97462a831b3a28e48c1bc63ec183d4a9c2ab47e01d7a1cdbefdfad544723646a.
//
// Solidity: event PolygonTokenBalanceRequest(bytes32 indexed requestId, address indexed requester, address tokenAddress, address tokenOwnerAddress, address callbackAddr, bytes4 callbackFunctionSignature)
func (_PolygonToken *PolygonTokenFilterer) WatchPolygonTokenBalanceRequest(opts *bind.WatchOpts, sink chan<- *PolygonTokenPolygonTokenBalanceRequest, requestId [][32]byte, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PolygonToken.contract.WatchLogs(opts, "PolygonTokenBalanceRequest", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonTokenPolygonTokenBalanceRequest)
				if err := _PolygonToken.contract.UnpackLog(event, "PolygonTokenBalanceRequest", log); err != nil {
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

// ParsePolygonTokenBalanceRequest is a log parse operation binding the contract event 0x97462a831b3a28e48c1bc63ec183d4a9c2ab47e01d7a1cdbefdfad544723646a.
//
// Solidity: event PolygonTokenBalanceRequest(bytes32 indexed requestId, address indexed requester, address tokenAddress, address tokenOwnerAddress, address callbackAddr, bytes4 callbackFunctionSignature)
func (_PolygonToken *PolygonTokenFilterer) ParsePolygonTokenBalanceRequest(log types.Log) (*PolygonTokenPolygonTokenBalanceRequest, error) {
	event := new(PolygonTokenPolygonTokenBalanceRequest)
	if err := _PolygonToken.contract.UnpackLog(event, "PolygonTokenBalanceRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
