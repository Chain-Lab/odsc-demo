// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// EthereumMetaData contains all meta data concerning the Ethereum contract.
var EthereumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"random\",\"type\":\"string\"}],\"name\":\"dataCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"random\",\"type\":\"string\"}],\"name\":\"dataModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"dataRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bootstrap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"random\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"genesisKey\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"random\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newBootstrap\",\"type\":\"string\"}],\"name\":\"initialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"random\",\"type\":\"string\"}],\"name\":\"modify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"new_bootstrap\",\"type\":\"string\"}],\"name\":\"setBootstrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumMetaData.ABI instead.
var EthereumABI = EthereumMetaData.ABI

// Ethereum is an auto generated Go binding around an Ethereum contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// Bootstrap is a free data retrieval call binding the contract method 0xfb969b0a.
//
// Solidity: function bootstrap() view returns(string)
func (_Ethereum *EthereumCaller) Bootstrap(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "bootstrap")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bootstrap is a free data retrieval call binding the contract method 0xfb969b0a.
//
// Solidity: function bootstrap() view returns(string)
func (_Ethereum *EthereumSession) Bootstrap() (string, error) {
	return _Ethereum.Contract.Bootstrap(&_Ethereum.CallOpts)
}

// Bootstrap is a free data retrieval call binding the contract method 0xfb969b0a.
//
// Solidity: function bootstrap() view returns(string)
func (_Ethereum *EthereumCallerSession) Bootstrap() (string, error) {
	return _Ethereum.Contract.Bootstrap(&_Ethereum.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() view returns(uint256)
func (_Ethereum *EthereumCaller) Genesis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "genesis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() view returns(uint256)
func (_Ethereum *EthereumSession) Genesis() (*big.Int, error) {
	return _Ethereum.Contract.Genesis(&_Ethereum.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() view returns(uint256)
func (_Ethereum *EthereumCallerSession) Genesis() (*big.Int, error) {
	return _Ethereum.Contract.Genesis(&_Ethereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumSession) Owner() (common.Address, error) {
	return _Ethereum.Contract.Owner(&_Ethereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumCallerSession) Owner() (common.Address, error) {
	return _Ethereum.Contract.Owner(&_Ethereum.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 key, string random) returns()
func (_Ethereum *EthereumTransactor) Create(opts *bind.TransactOpts, key *big.Int, random string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "create", key, random)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 key, string random) returns()
func (_Ethereum *EthereumSession) Create(key *big.Int, random string) (*types.Transaction, error) {
	return _Ethereum.Contract.Create(&_Ethereum.TransactOpts, key, random)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 key, string random) returns()
func (_Ethereum *EthereumTransactorSession) Create(key *big.Int, random string) (*types.Transaction, error) {
	return _Ethereum.Contract.Create(&_Ethereum.TransactOpts, key, random)
}

// Initialization is a paid mutator transaction binding the contract method 0xe0451b64.
//
// Solidity: function initialization(uint256 genesisKey, string random, string newBootstrap) returns()
func (_Ethereum *EthereumTransactor) Initialization(opts *bind.TransactOpts, genesisKey *big.Int, random string, newBootstrap string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "initialization", genesisKey, random, newBootstrap)
}

// Initialization is a paid mutator transaction binding the contract method 0xe0451b64.
//
// Solidity: function initialization(uint256 genesisKey, string random, string newBootstrap) returns()
func (_Ethereum *EthereumSession) Initialization(genesisKey *big.Int, random string, newBootstrap string) (*types.Transaction, error) {
	return _Ethereum.Contract.Initialization(&_Ethereum.TransactOpts, genesisKey, random, newBootstrap)
}

// Initialization is a paid mutator transaction binding the contract method 0xe0451b64.
//
// Solidity: function initialization(uint256 genesisKey, string random, string newBootstrap) returns()
func (_Ethereum *EthereumTransactorSession) Initialization(genesisKey *big.Int, random string, newBootstrap string) (*types.Transaction, error) {
	return _Ethereum.Contract.Initialization(&_Ethereum.TransactOpts, genesisKey, random, newBootstrap)
}

// Modify is a paid mutator transaction binding the contract method 0x441d10e2.
//
// Solidity: function modify(uint256 key, string random) returns()
func (_Ethereum *EthereumTransactor) Modify(opts *bind.TransactOpts, key *big.Int, random string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "modify", key, random)
}

// Modify is a paid mutator transaction binding the contract method 0x441d10e2.
//
// Solidity: function modify(uint256 key, string random) returns()
func (_Ethereum *EthereumSession) Modify(key *big.Int, random string) (*types.Transaction, error) {
	return _Ethereum.Contract.Modify(&_Ethereum.TransactOpts, key, random)
}

// Modify is a paid mutator transaction binding the contract method 0x441d10e2.
//
// Solidity: function modify(uint256 key, string random) returns()
func (_Ethereum *EthereumTransactorSession) Modify(key *big.Int, random string) (*types.Transaction, error) {
	return _Ethereum.Contract.Modify(&_Ethereum.TransactOpts, key, random)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethereum *EthereumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethereum *EthereumSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethereum.Contract.RenounceOwnership(&_Ethereum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethereum *EthereumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethereum.Contract.RenounceOwnership(&_Ethereum.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 key) returns()
func (_Ethereum *EthereumTransactor) Revoke(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "revoke", key)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 key) returns()
func (_Ethereum *EthereumSession) Revoke(key *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.Revoke(&_Ethereum.TransactOpts, key)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 key) returns()
func (_Ethereum *EthereumTransactorSession) Revoke(key *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.Revoke(&_Ethereum.TransactOpts, key)
}

// SetBootstrap is a paid mutator transaction binding the contract method 0xa746dadb.
//
// Solidity: function setBootstrap(string new_bootstrap) returns()
func (_Ethereum *EthereumTransactor) SetBootstrap(opts *bind.TransactOpts, new_bootstrap string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setBootstrap", new_bootstrap)
}

// SetBootstrap is a paid mutator transaction binding the contract method 0xa746dadb.
//
// Solidity: function setBootstrap(string new_bootstrap) returns()
func (_Ethereum *EthereumSession) SetBootstrap(new_bootstrap string) (*types.Transaction, error) {
	return _Ethereum.Contract.SetBootstrap(&_Ethereum.TransactOpts, new_bootstrap)
}

// SetBootstrap is a paid mutator transaction binding the contract method 0xa746dadb.
//
// Solidity: function setBootstrap(string new_bootstrap) returns()
func (_Ethereum *EthereumTransactorSession) SetBootstrap(new_bootstrap string) (*types.Transaction, error) {
	return _Ethereum.Contract.SetBootstrap(&_Ethereum.TransactOpts, new_bootstrap)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethereum *EthereumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethereum *EthereumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.TransferOwnership(&_Ethereum.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethereum *EthereumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.TransferOwnership(&_Ethereum.TransactOpts, newOwner)
}

// EthereumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ethereum contract.
type EthereumOwnershipTransferredIterator struct {
	Event *EthereumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthereumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumOwnershipTransferred)
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
		it.Event = new(EthereumOwnershipTransferred)
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
func (it *EthereumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumOwnershipTransferred represents a OwnershipTransferred event raised by the Ethereum contract.
type EthereumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethereum *EthereumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthereumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthereumOwnershipTransferredIterator{contract: _Ethereum.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethereum *EthereumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthereumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumOwnershipTransferred)
				if err := _Ethereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseOwnershipTransferred(log types.Log) (*EthereumOwnershipTransferred, error) {
	event := new(EthereumOwnershipTransferred)
	if err := _Ethereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDataCreatedIterator is returned from FilterDataCreated and is used to iterate over the raw logs and unpacked data for DataCreated events raised by the Ethereum contract.
type EthereumDataCreatedIterator struct {
	Event *EthereumDataCreated // Event containing the contract specifics and raw log

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
func (it *EthereumDataCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDataCreated)
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
		it.Event = new(EthereumDataCreated)
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
func (it *EthereumDataCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDataCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDataCreated represents a DataCreated event raised by the Ethereum contract.
type EthereumDataCreated struct {
	Key    *big.Int
	Random string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataCreated is a free log retrieval operation binding the contract event 0xecfc5d2a1f2595baeb9bdcf4cea149dcebb4b04ab893e5fddcd7d7e039f5768b.
//
// Solidity: event dataCreated(uint256 key, string random)
func (_Ethereum *EthereumFilterer) FilterDataCreated(opts *bind.FilterOpts) (*EthereumDataCreatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataCreated")
	if err != nil {
		return nil, err
	}
	return &EthereumDataCreatedIterator{contract: _Ethereum.contract, event: "dataCreated", logs: logs, sub: sub}, nil
}

// WatchDataCreated is a free log subscription operation binding the contract event 0xecfc5d2a1f2595baeb9bdcf4cea149dcebb4b04ab893e5fddcd7d7e039f5768b.
//
// Solidity: event dataCreated(uint256 key, string random)
func (_Ethereum *EthereumFilterer) WatchDataCreated(opts *bind.WatchOpts, sink chan<- *EthereumDataCreated) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "dataCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDataCreated)
				if err := _Ethereum.contract.UnpackLog(event, "dataCreated", log); err != nil {
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

// ParseDataCreated is a log parse operation binding the contract event 0xecfc5d2a1f2595baeb9bdcf4cea149dcebb4b04ab893e5fddcd7d7e039f5768b.
//
// Solidity: event dataCreated(uint256 key, string random)
func (_Ethereum *EthereumFilterer) ParseDataCreated(log types.Log) (*EthereumDataCreated, error) {
	event := new(EthereumDataCreated)
	if err := _Ethereum.contract.UnpackLog(event, "dataCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDataModifiedIterator is returned from FilterDataModified and is used to iterate over the raw logs and unpacked data for DataModified events raised by the Ethereum contract.
type EthereumDataModifiedIterator struct {
	Event *EthereumDataModified // Event containing the contract specifics and raw log

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
func (it *EthereumDataModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDataModified)
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
		it.Event = new(EthereumDataModified)
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
func (it *EthereumDataModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDataModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDataModified represents a DataModified event raised by the Ethereum contract.
type EthereumDataModified struct {
	Key    *big.Int
	Random string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataModified is a free log retrieval operation binding the contract event 0xa6be96c178c67432984374a8be68549a855f46237a530610843abadbf26065b9.
//
// Solidity: event dataModified(uint256 key, string random)
func (_Ethereum *EthereumFilterer) FilterDataModified(opts *bind.FilterOpts) (*EthereumDataModifiedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataModified")
	if err != nil {
		return nil, err
	}
	return &EthereumDataModifiedIterator{contract: _Ethereum.contract, event: "dataModified", logs: logs, sub: sub}, nil
}

// WatchDataModified is a free log subscription operation binding the contract event 0xa6be96c178c67432984374a8be68549a855f46237a530610843abadbf26065b9.
//
// Solidity: event dataModified(uint256 key, string random)
func (_Ethereum *EthereumFilterer) WatchDataModified(opts *bind.WatchOpts, sink chan<- *EthereumDataModified) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "dataModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDataModified)
				if err := _Ethereum.contract.UnpackLog(event, "dataModified", log); err != nil {
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

// ParseDataModified is a log parse operation binding the contract event 0xa6be96c178c67432984374a8be68549a855f46237a530610843abadbf26065b9.
//
// Solidity: event dataModified(uint256 key, string random)
func (_Ethereum *EthereumFilterer) ParseDataModified(log types.Log) (*EthereumDataModified, error) {
	event := new(EthereumDataModified)
	if err := _Ethereum.contract.UnpackLog(event, "dataModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDataRevokedIterator is returned from FilterDataRevoked and is used to iterate over the raw logs and unpacked data for DataRevoked events raised by the Ethereum contract.
type EthereumDataRevokedIterator struct {
	Event *EthereumDataRevoked // Event containing the contract specifics and raw log

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
func (it *EthereumDataRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDataRevoked)
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
		it.Event = new(EthereumDataRevoked)
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
func (it *EthereumDataRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDataRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDataRevoked represents a DataRevoked event raised by the Ethereum contract.
type EthereumDataRevoked struct {
	Key *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDataRevoked is a free log retrieval operation binding the contract event 0xc5e5964f72ab19de1a46cbfde95a2eeb3422b17d72be21132f39bd164f47be63.
//
// Solidity: event dataRevoked(uint256 key)
func (_Ethereum *EthereumFilterer) FilterDataRevoked(opts *bind.FilterOpts) (*EthereumDataRevokedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataRevoked")
	if err != nil {
		return nil, err
	}
	return &EthereumDataRevokedIterator{contract: _Ethereum.contract, event: "dataRevoked", logs: logs, sub: sub}, nil
}

// WatchDataRevoked is a free log subscription operation binding the contract event 0xc5e5964f72ab19de1a46cbfde95a2eeb3422b17d72be21132f39bd164f47be63.
//
// Solidity: event dataRevoked(uint256 key)
func (_Ethereum *EthereumFilterer) WatchDataRevoked(opts *bind.WatchOpts, sink chan<- *EthereumDataRevoked) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "dataRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDataRevoked)
				if err := _Ethereum.contract.UnpackLog(event, "dataRevoked", log); err != nil {
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

// ParseDataRevoked is a log parse operation binding the contract event 0xc5e5964f72ab19de1a46cbfde95a2eeb3422b17d72be21132f39bd164f47be63.
//
// Solidity: event dataRevoked(uint256 key)
func (_Ethereum *EthereumFilterer) ParseDataRevoked(log types.Log) (*EthereumDataRevoked, error) {
	event := new(EthereumDataRevoked)
	if err := _Ethereum.contract.UnpackLog(event, "dataRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
