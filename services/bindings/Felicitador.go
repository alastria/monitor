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

// FelicitadorABI is the input ABI used to generate the binding from.
const FelicitadorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"cual\",\"type\":\"uint256\"}],\"name\":\"leerFelicitacion\",\"outputs\":[{\"name\":\"nombreFelicitador\",\"type\":\"string\"},{\"name\":\"mensaje\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"felicitador\",\"type\":\"string\"},{\"name\":\"mensaje\",\"type\":\"string\"}],\"name\":\"felicita\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leerCuantasFelicitaciones\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FelicitadorBin is the compiled bytecode used for deploying new contracts.
const FelicitadorBin = `60606040526000600155341561001457600080fd5b6105f1806100236000396000f300606060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633b7ed0261461005c5780639bb8f1fe14610164578063d88229e014610204575b600080fd5b341561006757600080fd5b61007d600480803590602001909190505061022d565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156100c15780820151818401526020810190506100a6565b50505050905090810190601f1680156100ee5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561012757808201518184015260208101905061010c565b50505050905090810190601f1680156101545780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b341561016f57600080fd5b610202600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506103be565b005b341561020f57600080fd5b610217610449565b6040518082815260200191505060405180910390f35b610235610453565b61023d610453565b60008381548110151561024c57fe5b906000526020600020906002020160000160008481548110151561026c57fe5b9060005260206000209060020201600101818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103125780601f106102e757610100808354040283529160200191610312565b820191906000526020600020905b8154815290600101906020018083116102f557829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103ae5780601f10610383576101008083540402835291602001916103ae565b820191906000526020600020905b81548152906001019060200180831161039157829003601f168201915b5050505050905091509150915091565b600080548060010182816103d29190610467565b9160005260206000209060020201600060408051908101604052808681526020018581525090919091506000820151816000019080519060200190610418929190610499565b506020820151816001019080519060200190610435929190610499565b505050506000805490506001819055505050565b6000600154905090565b602060405190810160405280600081525090565b815481835581811511610494576002028160020283600052602060002091820191016104939190610519565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104da57805160ff1916838001178555610508565b82800160010185558215610508579182015b828111156105075782518255916020019190600101906104ec565b5b5090506105159190610558565b5090565b61055591905b808211156105515760008082016000610538919061057d565b600182016000610548919061057d565b5060020161051f565b5090565b90565b61057a91905b8082111561057657600081600090555060010161055e565b5090565b90565b50805460018160011615610100020316600290046000825580601f106105a357506105c2565b601f0160209004906000526020600020908101906105c19190610558565b5b505600a165627a7a72305820431ee1d48a1fd4d07671c6b03058057b6db78be48cb85244aff374698fe482870029`

// DeployFelicitador deploys a new Ethereum contract, binding an instance of Felicitador to it.
func DeployFelicitador(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Felicitador, error) {
	parsed, err := abi.JSON(strings.NewReader(FelicitadorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FelicitadorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Felicitador{FelicitadorCaller: FelicitadorCaller{contract: contract}, FelicitadorTransactor: FelicitadorTransactor{contract: contract}}, nil
}

// Felicitador is an auto generated Go binding around an Ethereum contract.
type Felicitador struct {
	FelicitadorCaller     // Read-only binding to the contract
	FelicitadorTransactor // Write-only binding to the contract
}

// FelicitadorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FelicitadorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FelicitadorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FelicitadorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FelicitadorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FelicitadorSession struct {
	Contract     *Felicitador      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FelicitadorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FelicitadorCallerSession struct {
	Contract *FelicitadorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FelicitadorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FelicitadorTransactorSession struct {
	Contract     *FelicitadorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FelicitadorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FelicitadorRaw struct {
	Contract *Felicitador // Generic contract binding to access the raw methods on
}

// FelicitadorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FelicitadorCallerRaw struct {
	Contract *FelicitadorCaller // Generic read-only contract binding to access the raw methods on
}

// FelicitadorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FelicitadorTransactorRaw struct {
	Contract *FelicitadorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFelicitador creates a new instance of Felicitador, bound to a specific deployed contract.
func NewFelicitador(address common.Address, backend bind.ContractBackend) (*Felicitador, error) {
	contract, err := bindFelicitador(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Felicitador{FelicitadorCaller: FelicitadorCaller{contract: contract}, FelicitadorTransactor: FelicitadorTransactor{contract: contract}}, nil
}

// NewFelicitadorCaller creates a new read-only instance of Felicitador, bound to a specific deployed contract.
func NewFelicitadorCaller(address common.Address, caller bind.ContractCaller) (*FelicitadorCaller, error) {
	contract, err := bindFelicitador(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FelicitadorCaller{contract: contract}, nil
}

// NewFelicitadorTransactor creates a new write-only instance of Felicitador, bound to a specific deployed contract.
func NewFelicitadorTransactor(address common.Address, transactor bind.ContractTransactor) (*FelicitadorTransactor, error) {
	contract, err := bindFelicitador(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FelicitadorTransactor{contract: contract}, nil
}

// bindFelicitador binds a generic wrapper to an already deployed contract.
func bindFelicitador(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FelicitadorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Felicitador *FelicitadorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Felicitador.Contract.FelicitadorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Felicitador *FelicitadorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Felicitador.Contract.FelicitadorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Felicitador *FelicitadorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Felicitador.Contract.FelicitadorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Felicitador *FelicitadorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Felicitador.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Felicitador *FelicitadorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Felicitador.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Felicitador *FelicitadorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Felicitador.Contract.contract.Transact(opts, method, params...)
}

// LeerCuantasFelicitaciones is a free data retrieval call binding the contract method 0xd88229e0.
//
// Solidity: function leerCuantasFelicitaciones() constant returns(result uint256)
func (_Felicitador *FelicitadorCaller) LeerCuantasFelicitaciones(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Felicitador.contract.Call(opts, out, "leerCuantasFelicitaciones")
	return *ret0, err
}

// LeerCuantasFelicitaciones is a free data retrieval call binding the contract method 0xd88229e0.
//
// Solidity: function leerCuantasFelicitaciones() constant returns(result uint256)
func (_Felicitador *FelicitadorSession) LeerCuantasFelicitaciones() (*big.Int, error) {
	return _Felicitador.Contract.LeerCuantasFelicitaciones(&_Felicitador.CallOpts)
}

// LeerCuantasFelicitaciones is a free data retrieval call binding the contract method 0xd88229e0.
//
// Solidity: function leerCuantasFelicitaciones() constant returns(result uint256)
func (_Felicitador *FelicitadorCallerSession) LeerCuantasFelicitaciones() (*big.Int, error) {
	return _Felicitador.Contract.LeerCuantasFelicitaciones(&_Felicitador.CallOpts)
}

// LeerFelicitacion is a free data retrieval call binding the contract method 0x3b7ed026.
//
// Solidity: function leerFelicitacion(cual uint256) constant returns(nombreFelicitador string, mensaje string)
func (_Felicitador *FelicitadorCaller) LeerFelicitacion(opts *bind.CallOpts, cual *big.Int) (struct {
	NombreFelicitador string
	Mensaje           string
}, error) {
	ret := new(struct {
		NombreFelicitador string
		Mensaje           string
	})
	out := ret
	err := _Felicitador.contract.Call(opts, out, "leerFelicitacion", cual)
	return *ret, err
}

// LeerFelicitacion is a free data retrieval call binding the contract method 0x3b7ed026.
//
// Solidity: function leerFelicitacion(cual uint256) constant returns(nombreFelicitador string, mensaje string)
func (_Felicitador *FelicitadorSession) LeerFelicitacion(cual *big.Int) (struct {
	NombreFelicitador string
	Mensaje           string
}, error) {
	return _Felicitador.Contract.LeerFelicitacion(&_Felicitador.CallOpts, cual)
}

// LeerFelicitacion is a free data retrieval call binding the contract method 0x3b7ed026.
//
// Solidity: function leerFelicitacion(cual uint256) constant returns(nombreFelicitador string, mensaje string)
func (_Felicitador *FelicitadorCallerSession) LeerFelicitacion(cual *big.Int) (struct {
	NombreFelicitador string
	Mensaje           string
}, error) {
	return _Felicitador.Contract.LeerFelicitacion(&_Felicitador.CallOpts, cual)
}

// Felicita is a paid mutator transaction binding the contract method 0x9bb8f1fe.
//
// Solidity: function felicita(felicitador string, mensaje string) returns()
func (_Felicitador *FelicitadorTransactor) Felicita(opts *bind.TransactOpts, felicitador string, mensaje string) (*types.Transaction, error) {
	return _Felicitador.contract.Transact(opts, "felicita", felicitador, mensaje)
}

// Felicita is a paid mutator transaction binding the contract method 0x9bb8f1fe.
//
// Solidity: function felicita(felicitador string, mensaje string) returns()
func (_Felicitador *FelicitadorSession) Felicita(felicitador string, mensaje string) (*types.Transaction, error) {
	return _Felicitador.Contract.Felicita(&_Felicitador.TransactOpts, felicitador, mensaje)
}

// Felicita is a paid mutator transaction binding the contract method 0x9bb8f1fe.
//
// Solidity: function felicita(felicitador string, mensaje string) returns()
func (_Felicitador *FelicitadorTransactorSession) Felicita(felicitador string, mensaje string) (*types.Transaction, error) {
	return _Felicitador.Contract.Felicita(&_Felicitador.TransactOpts, felicitador, mensaje)
}
