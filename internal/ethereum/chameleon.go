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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"random\",\"type\":\"bytes\"}],\"name\":\"dataCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"random\",\"type\":\"bytes\"}],\"name\":\"dataModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"dataRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bootstrap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"random\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"genesisKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"random\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"newBootstrap\",\"type\":\"string\"}],\"name\":\"initialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"random\",\"type\":\"bytes\"}],\"name\":\"modify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"new_bootstrap\",\"type\":\"string\"}],\"name\":\"setBootstrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b50610ad1806100246000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a746dadb1161005b578063a746dadb146100d6578063a7f0b3de146100f2578063fb969b0a14610110578063fd7230d61461012e5761007d565b806320c5429b146100825780634420220b1461009e57806355235348146100ba575b600080fd5b61009c60048036038101906100979190610394565b61014a565b005b6100b860048036038101906100b3919061047c565b610184565b005b6100d460048036038101906100cf9190610511565b610223565b005b6100f060048036038101906100eb9190610571565b610263565b005b6100fa610279565b60405161010791906105cd565b60405180910390f35b610118610282565b6040516101259190610678565b60405180910390f35b61014860048036038101906101439190610511565b610314565b005b7fc5e5964f72ab19de1a46cbfde95a2eeb3422b17d72be21132f39bd164f47be638160405161017991906105cd565b60405180910390a150565b60008054146101c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bf906106e6565b60405180910390fd5b846000819055508181600191826101e092919061094c565b507f8cfa76e35a985c69a84a9a45394e09f857cb6a0263aa7ab43f0d2cf7bf31ade785858560405161021493929190610a69565b60405180910390a15050505050565b7f4fb1ae1f2c62b665c3780633343fe9a508b8b525f83211d2033769cc58a896eb83838360405161025693929190610a69565b60405180910390a1505050565b81816001918261027492919061094c565b505050565b60008054905090565b6060600180546102919061076f565b80601f01602080910402602001604051908101604052809291908181526020018280546102bd9061076f565b801561030a5780601f106102df5761010080835404028352916020019161030a565b820191906000526020600020905b8154815290600101906020018083116102ed57829003601f168201915b5050505050905090565b7f8cfa76e35a985c69a84a9a45394e09f857cb6a0263aa7ab43f0d2cf7bf31ade783838360405161034793929190610a69565b60405180910390a1505050565b600080fd5b600080fd5b6000819050919050565b6103718161035e565b811461037c57600080fd5b50565b60008135905061038e81610368565b92915050565b6000602082840312156103aa576103a9610354565b5b60006103b88482850161037f565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126103e6576103e56103c1565b5b8235905067ffffffffffffffff811115610403576104026103c6565b5b60208301915083600182028301111561041f5761041e6103cb565b5b9250929050565b60008083601f84011261043c5761043b6103c1565b5b8235905067ffffffffffffffff811115610459576104586103c6565b5b602083019150836001820283011115610475576104746103cb565b5b9250929050565b60008060008060006060868803121561049857610497610354565b5b60006104a68882890161037f565b955050602086013567ffffffffffffffff8111156104c7576104c6610359565b5b6104d3888289016103d0565b9450945050604086013567ffffffffffffffff8111156104f6576104f5610359565b5b61050288828901610426565b92509250509295509295909350565b60008060006040848603121561052a57610529610354565b5b60006105388682870161037f565b935050602084013567ffffffffffffffff81111561055957610558610359565b5b610565868287016103d0565b92509250509250925092565b6000806020838503121561058857610587610354565b5b600083013567ffffffffffffffff8111156105a6576105a5610359565b5b6105b285828601610426565b92509250509250929050565b6105c78161035e565b82525050565b60006020820190506105e260008301846105be565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610622578082015181840152602081019050610607565b60008484015250505050565b6000601f19601f8301169050919050565b600061064a826105e8565b61065481856105f3565b9350610664818560208601610604565b61066d8161062e565b840191505092915050565b60006020820190508181036000830152610692818461063f565b905092915050565b7f47656e65736973204461746120496e697469616c656400000000000000000000600082015250565b60006106d06016836105f3565b91506106db8261069a565b602082019050919050565b600060208201905081810360008301526106ff816106c3565b9050919050565b600082905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061078757607f821691505b60208210810361079a57610799610740565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826107c5565b61080c86836107c5565b95508019841693508086168417925050509392505050565b6000819050919050565b600061084961084461083f8461035e565b610824565b61035e565b9050919050565b6000819050919050565b6108638361082e565b61087761086f82610850565b8484546107d2565b825550505050565b600090565b61088c61087f565b61089781848461085a565b505050565b5b818110156108bb576108b0600082610884565b60018101905061089d565b5050565b601f821115610900576108d1816107a0565b6108da846107b5565b810160208510156108e9578190505b6108fd6108f5856107b5565b83018261089c565b50505b505050565b600082821c905092915050565b600061092360001984600802610905565b1980831691505092915050565b600061093c8383610912565b9150826002028217905092915050565b6109568383610706565b67ffffffffffffffff81111561096f5761096e610711565b5b610979825461076f565b6109848282856108bf565b6000601f8311600181146109b357600084156109a1578287013590505b6109ab8582610930565b865550610a13565b601f1984166109c1866107a0565b60005b828110156109e9578489013582556001820191506020850194506020810190506109c4565b86831015610a065784890135610a02601f891682610912565b8355505b6001600288020188555050505b50505050505050565b600082825260208201905092915050565b82818337600083830152505050565b6000610a488385610a1c565b9350610a55838584610a2d565b610a5e8361062e565b840190509392505050565b6000604082019050610a7e60008301866105be565b8181036020830152610a91818486610a3c565b905094935050505056fea264697066735822122046e2c01e6b2c0cd4b0cad7a65b89663bb439807a5e7f5e50e5fa6a194144d0ac64736f6c63430008110033",
}

// EthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumMetaData.ABI instead.
var EthereumABI = EthereumMetaData.ABI

// EthereumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthereumMetaData.Bin instead.
var EthereumBin = EthereumMetaData.Bin

// DeployEthereum deploys a new Ethereum contract, binding an instance of Ethereum to it.
func DeployEthereum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethereum, error) {
	parsed, err := EthereumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthereumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

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

// Create is a paid mutator transaction binding the contract method 0xfd7230d6.
//
// Solidity: function create(uint256 key, bytes random) returns()
func (_Ethereum *EthereumTransactor) Create(opts *bind.TransactOpts, key *big.Int, random []byte) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "create", key, random)
}

// Create is a paid mutator transaction binding the contract method 0xfd7230d6.
//
// Solidity: function create(uint256 key, bytes random) returns()
func (_Ethereum *EthereumSession) Create(key *big.Int, random []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.Create(&_Ethereum.TransactOpts, key, random)
}

// Create is a paid mutator transaction binding the contract method 0xfd7230d6.
//
// Solidity: function create(uint256 key, bytes random) returns()
func (_Ethereum *EthereumTransactorSession) Create(key *big.Int, random []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.Create(&_Ethereum.TransactOpts, key, random)
}

// Initialization is a paid mutator transaction binding the contract method 0x4420220b.
//
// Solidity: function initialization(uint256 genesisKey, bytes random, string newBootstrap) returns()
func (_Ethereum *EthereumTransactor) Initialization(opts *bind.TransactOpts, genesisKey *big.Int, random []byte, newBootstrap string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "initialization", genesisKey, random, newBootstrap)
}

// Initialization is a paid mutator transaction binding the contract method 0x4420220b.
//
// Solidity: function initialization(uint256 genesisKey, bytes random, string newBootstrap) returns()
func (_Ethereum *EthereumSession) Initialization(genesisKey *big.Int, random []byte, newBootstrap string) (*types.Transaction, error) {
	return _Ethereum.Contract.Initialization(&_Ethereum.TransactOpts, genesisKey, random, newBootstrap)
}

// Initialization is a paid mutator transaction binding the contract method 0x4420220b.
//
// Solidity: function initialization(uint256 genesisKey, bytes random, string newBootstrap) returns()
func (_Ethereum *EthereumTransactorSession) Initialization(genesisKey *big.Int, random []byte, newBootstrap string) (*types.Transaction, error) {
	return _Ethereum.Contract.Initialization(&_Ethereum.TransactOpts, genesisKey, random, newBootstrap)
}

// Modify is a paid mutator transaction binding the contract method 0x55235348.
//
// Solidity: function modify(uint256 key, bytes random) returns()
func (_Ethereum *EthereumTransactor) Modify(opts *bind.TransactOpts, key *big.Int, random []byte) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "modify", key, random)
}

// Modify is a paid mutator transaction binding the contract method 0x55235348.
//
// Solidity: function modify(uint256 key, bytes random) returns()
func (_Ethereum *EthereumSession) Modify(key *big.Int, random []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.Modify(&_Ethereum.TransactOpts, key, random)
}

// Modify is a paid mutator transaction binding the contract method 0x55235348.
//
// Solidity: function modify(uint256 key, bytes random) returns()
func (_Ethereum *EthereumTransactorSession) Modify(key *big.Int, random []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.Modify(&_Ethereum.TransactOpts, key, random)
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
	Random []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataCreated is a free log retrieval operation binding the contract event 0x8cfa76e35a985c69a84a9a45394e09f857cb6a0263aa7ab43f0d2cf7bf31ade7.
//
// Solidity: event dataCreated(uint256 key, bytes random)
func (_Ethereum *EthereumFilterer) FilterDataCreated(opts *bind.FilterOpts) (*EthereumDataCreatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataCreated")
	if err != nil {
		return nil, err
	}
	return &EthereumDataCreatedIterator{contract: _Ethereum.contract, event: "dataCreated", logs: logs, sub: sub}, nil
}

// WatchDataCreated is a free log subscription operation binding the contract event 0x8cfa76e35a985c69a84a9a45394e09f857cb6a0263aa7ab43f0d2cf7bf31ade7.
//
// Solidity: event dataCreated(uint256 key, bytes random)
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

// ParseDataCreated is a log parse operation binding the contract event 0x8cfa76e35a985c69a84a9a45394e09f857cb6a0263aa7ab43f0d2cf7bf31ade7.
//
// Solidity: event dataCreated(uint256 key, bytes random)
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
	Random []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataModified is a free log retrieval operation binding the contract event 0x4fb1ae1f2c62b665c3780633343fe9a508b8b525f83211d2033769cc58a896eb.
//
// Solidity: event dataModified(uint256 key, bytes random)
func (_Ethereum *EthereumFilterer) FilterDataModified(opts *bind.FilterOpts) (*EthereumDataModifiedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataModified")
	if err != nil {
		return nil, err
	}
	return &EthereumDataModifiedIterator{contract: _Ethereum.contract, event: "dataModified", logs: logs, sub: sub}, nil
}

// WatchDataModified is a free log subscription operation binding the contract event 0x4fb1ae1f2c62b665c3780633343fe9a508b8b525f83211d2033769cc58a896eb.
//
// Solidity: event dataModified(uint256 key, bytes random)
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

// ParseDataModified is a log parse operation binding the contract event 0x4fb1ae1f2c62b665c3780633343fe9a508b8b525f83211d2033769cc58a896eb.
//
// Solidity: event dataModified(uint256 key, bytes random)
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
