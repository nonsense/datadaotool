// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dc

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

// DealClientDealProposal is an auto generated low-level Go binding around an user-defined struct.
type DealClientDealProposal struct {
	PieceCid             []byte
	PaddedPieceSize      uint64
	VerifiedDeal         bool
	Client               []byte
	Label                []byte
	StartEpoch           uint64
	EndEpoch             uint64
	StoragePricePerEpoch uint64
	ProviderCollateral   uint64
	ClientCollateral     uint64
	Version              string
	Params               []byte
}

// DealClientMetaData contains all meta data concerning the DealClient contract.
var DealClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DealProposalCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"received\",\"type\":\"string\"}],\"name\":\"ReceivedDataCap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORIZE_MESSAGE_METHOD_NUM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATACAP_RECEIVER_HOOK_METHOD_NUM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cidraw\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"addCID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cidraw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"authorizeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cidProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cidSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cidSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getDealProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pieceCid\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"paddedPieceSize\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"verifiedDeal\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"client\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"label\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storagePricePerEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"providerCollateral\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"clientCollateral\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structDealClient.DealProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"method\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"handle_filecoin_method\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pieceCid\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"paddedPieceSize\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"verifiedDeal\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"client\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"label\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storagePricePerEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"providerCollateral\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"clientCollateral\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structDealClient.DealProposal\",\"name\":\"_deal\",\"type\":\"tuple\"}],\"name\":\"makeDealProposal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pieceCid\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"paddedPieceSize\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"verifiedDeal\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"client\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"label\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storagePricePerEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"providerCollateral\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"clientCollateral\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"raw_auth_params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"}],\"name\":\"publish_deal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DealClientABI is the input ABI used to generate the binding from.
// Deprecated: Use DealClientMetaData.ABI instead.
var DealClientABI = DealClientMetaData.ABI

// DealClient is an auto generated Go binding around an Ethereum contract.
type DealClient struct {
	DealClientCaller     // Read-only binding to the contract
	DealClientTransactor // Write-only binding to the contract
	DealClientFilterer   // Log filterer for contract events
}

// DealClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type DealClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DealClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DealClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DealClientSession struct {
	Contract     *DealClient       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DealClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DealClientCallerSession struct {
	Contract *DealClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DealClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DealClientTransactorSession struct {
	Contract     *DealClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DealClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type DealClientRaw struct {
	Contract *DealClient // Generic contract binding to access the raw methods on
}

// DealClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DealClientCallerRaw struct {
	Contract *DealClientCaller // Generic read-only contract binding to access the raw methods on
}

// DealClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DealClientTransactorRaw struct {
	Contract *DealClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDealClient creates a new instance of DealClient, bound to a specific deployed contract.
func NewDealClient(address common.Address, backend bind.ContractBackend) (*DealClient, error) {
	contract, err := bindDealClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DealClient{DealClientCaller: DealClientCaller{contract: contract}, DealClientTransactor: DealClientTransactor{contract: contract}, DealClientFilterer: DealClientFilterer{contract: contract}}, nil
}

// NewDealClientCaller creates a new read-only instance of DealClient, bound to a specific deployed contract.
func NewDealClientCaller(address common.Address, caller bind.ContractCaller) (*DealClientCaller, error) {
	contract, err := bindDealClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DealClientCaller{contract: contract}, nil
}

// NewDealClientTransactor creates a new write-only instance of DealClient, bound to a specific deployed contract.
func NewDealClientTransactor(address common.Address, transactor bind.ContractTransactor) (*DealClientTransactor, error) {
	contract, err := bindDealClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DealClientTransactor{contract: contract}, nil
}

// NewDealClientFilterer creates a new log filterer instance of DealClient, bound to a specific deployed contract.
func NewDealClientFilterer(address common.Address, filterer bind.ContractFilterer) (*DealClientFilterer, error) {
	contract, err := bindDealClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DealClientFilterer{contract: contract}, nil
}

// bindDealClient binds a generic wrapper to an already deployed contract.
func bindDealClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DealClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DealClient *DealClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DealClient.Contract.DealClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DealClient *DealClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DealClient.Contract.DealClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DealClient *DealClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DealClient.Contract.DealClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DealClient *DealClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DealClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DealClient *DealClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DealClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DealClient *DealClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DealClient.Contract.contract.Transact(opts, method, params...)
}

// AUTHORIZEMESSAGEMETHODNUM is a free data retrieval call binding the contract method 0xc3bd579b.
//
// Solidity: function AUTHORIZE_MESSAGE_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCaller) AUTHORIZEMESSAGEMETHODNUM(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "AUTHORIZE_MESSAGE_METHOD_NUM")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AUTHORIZEMESSAGEMETHODNUM is a free data retrieval call binding the contract method 0xc3bd579b.
//
// Solidity: function AUTHORIZE_MESSAGE_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientSession) AUTHORIZEMESSAGEMETHODNUM() (uint64, error) {
	return _DealClient.Contract.AUTHORIZEMESSAGEMETHODNUM(&_DealClient.CallOpts)
}

// AUTHORIZEMESSAGEMETHODNUM is a free data retrieval call binding the contract method 0xc3bd579b.
//
// Solidity: function AUTHORIZE_MESSAGE_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCallerSession) AUTHORIZEMESSAGEMETHODNUM() (uint64, error) {
	return _DealClient.Contract.AUTHORIZEMESSAGEMETHODNUM(&_DealClient.CallOpts)
}

// DATACAPRECEIVERHOOKMETHODNUM is a free data retrieval call binding the contract method 0xb34ba252.
//
// Solidity: function DATACAP_RECEIVER_HOOK_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCaller) DATACAPRECEIVERHOOKMETHODNUM(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "DATACAP_RECEIVER_HOOK_METHOD_NUM")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DATACAPRECEIVERHOOKMETHODNUM is a free data retrieval call binding the contract method 0xb34ba252.
//
// Solidity: function DATACAP_RECEIVER_HOOK_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientSession) DATACAPRECEIVERHOOKMETHODNUM() (uint64, error) {
	return _DealClient.Contract.DATACAPRECEIVERHOOKMETHODNUM(&_DealClient.CallOpts)
}

// DATACAPRECEIVERHOOKMETHODNUM is a free data retrieval call binding the contract method 0xb34ba252.
//
// Solidity: function DATACAP_RECEIVER_HOOK_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCallerSession) DATACAPRECEIVERHOOKMETHODNUM() (uint64, error) {
	return _DealClient.Contract.DATACAPRECEIVERHOOKMETHODNUM(&_DealClient.CallOpts)
}

// CidProviders is a free data retrieval call binding the contract method 0x6fc6a748.
//
// Solidity: function cidProviders(bytes , bytes ) view returns(bool)
func (_DealClient *DealClientCaller) CidProviders(opts *bind.CallOpts, arg0 []byte, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "cidProviders", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CidProviders is a free data retrieval call binding the contract method 0x6fc6a748.
//
// Solidity: function cidProviders(bytes , bytes ) view returns(bool)
func (_DealClient *DealClientSession) CidProviders(arg0 []byte, arg1 []byte) (bool, error) {
	return _DealClient.Contract.CidProviders(&_DealClient.CallOpts, arg0, arg1)
}

// CidProviders is a free data retrieval call binding the contract method 0x6fc6a748.
//
// Solidity: function cidProviders(bytes , bytes ) view returns(bool)
func (_DealClient *DealClientCallerSession) CidProviders(arg0 []byte, arg1 []byte) (bool, error) {
	return _DealClient.Contract.CidProviders(&_DealClient.CallOpts, arg0, arg1)
}

// CidSet is a free data retrieval call binding the contract method 0x3d7650b2.
//
// Solidity: function cidSet(bytes ) view returns(bool)
func (_DealClient *DealClientCaller) CidSet(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "cidSet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CidSet is a free data retrieval call binding the contract method 0x3d7650b2.
//
// Solidity: function cidSet(bytes ) view returns(bool)
func (_DealClient *DealClientSession) CidSet(arg0 []byte) (bool, error) {
	return _DealClient.Contract.CidSet(&_DealClient.CallOpts, arg0)
}

// CidSet is a free data retrieval call binding the contract method 0x3d7650b2.
//
// Solidity: function cidSet(bytes ) view returns(bool)
func (_DealClient *DealClientCallerSession) CidSet(arg0 []byte) (bool, error) {
	return _DealClient.Contract.CidSet(&_DealClient.CallOpts, arg0)
}

// CidSizes is a free data retrieval call binding the contract method 0x8f8a6d8f.
//
// Solidity: function cidSizes(bytes ) view returns(uint256)
func (_DealClient *DealClientCaller) CidSizes(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "cidSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CidSizes is a free data retrieval call binding the contract method 0x8f8a6d8f.
//
// Solidity: function cidSizes(bytes ) view returns(uint256)
func (_DealClient *DealClientSession) CidSizes(arg0 []byte) (*big.Int, error) {
	return _DealClient.Contract.CidSizes(&_DealClient.CallOpts, arg0)
}

// CidSizes is a free data retrieval call binding the contract method 0x8f8a6d8f.
//
// Solidity: function cidSizes(bytes ) view returns(uint256)
func (_DealClient *DealClientCallerSession) CidSizes(arg0 []byte) (*big.Int, error) {
	return _DealClient.Contract.CidSizes(&_DealClient.CallOpts, arg0)
}

// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
//
// Solidity: function getDealProposal(bytes32 id) view returns((bytes,uint64,bool,bytes,bytes,uint64,uint64,uint64,uint64,uint64,string,bytes))
func (_DealClient *DealClientCaller) GetDealProposal(opts *bind.CallOpts, id [32]byte) (DealClientDealProposal, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "getDealProposal", id)

	if err != nil {
		return *new(DealClientDealProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(DealClientDealProposal)).(*DealClientDealProposal)

	return out0, err

}

// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
//
// Solidity: function getDealProposal(bytes32 id) view returns((bytes,uint64,bool,bytes,bytes,uint64,uint64,uint64,uint64,uint64,string,bytes))
func (_DealClient *DealClientSession) GetDealProposal(id [32]byte) (DealClientDealProposal, error) {
	return _DealClient.Contract.GetDealProposal(&_DealClient.CallOpts, id)
}

// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
//
// Solidity: function getDealProposal(bytes32 id) view returns((bytes,uint64,bool,bytes,bytes,uint64,uint64,uint64,uint64,uint64,string,bytes))
func (_DealClient *DealClientCallerSession) GetDealProposal(id [32]byte) (DealClientDealProposal, error) {
	return _DealClient.Contract.GetDealProposal(&_DealClient.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DealClient *DealClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DealClient *DealClientSession) Owner() (common.Address, error) {
	return _DealClient.Contract.Owner(&_DealClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DealClient *DealClientCallerSession) Owner() (common.Address, error) {
	return _DealClient.Contract.Owner(&_DealClient.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(bytes pieceCid, uint64 paddedPieceSize, bool verifiedDeal, bytes client, bytes label, uint64 startEpoch, uint64 endEpoch, uint64 storagePricePerEpoch, uint64 providerCollateral, uint64 clientCollateral, string version, bytes params)
func (_DealClient *DealClientCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	PieceCid             []byte
	PaddedPieceSize      uint64
	VerifiedDeal         bool
	Client               []byte
	Label                []byte
	StartEpoch           uint64
	EndEpoch             uint64
	StoragePricePerEpoch uint64
	ProviderCollateral   uint64
	ClientCollateral     uint64
	Version              string
	Params               []byte
}, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		PieceCid             []byte
		PaddedPieceSize      uint64
		VerifiedDeal         bool
		Client               []byte
		Label                []byte
		StartEpoch           uint64
		EndEpoch             uint64
		StoragePricePerEpoch uint64
		ProviderCollateral   uint64
		ClientCollateral     uint64
		Version              string
		Params               []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PieceCid = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.PaddedPieceSize = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.VerifiedDeal = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Client = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Label = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.StartEpoch = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.EndEpoch = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.StoragePricePerEpoch = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.ProviderCollateral = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.ClientCollateral = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.Version = *abi.ConvertType(out[10], new(string)).(*string)
	outstruct.Params = *abi.ConvertType(out[11], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(bytes pieceCid, uint64 paddedPieceSize, bool verifiedDeal, bytes client, bytes label, uint64 startEpoch, uint64 endEpoch, uint64 storagePricePerEpoch, uint64 providerCollateral, uint64 clientCollateral, string version, bytes params)
func (_DealClient *DealClientSession) Proposals(arg0 [32]byte) (struct {
	PieceCid             []byte
	PaddedPieceSize      uint64
	VerifiedDeal         bool
	Client               []byte
	Label                []byte
	StartEpoch           uint64
	EndEpoch             uint64
	StoragePricePerEpoch uint64
	ProviderCollateral   uint64
	ClientCollateral     uint64
	Version              string
	Params               []byte
}, error) {
	return _DealClient.Contract.Proposals(&_DealClient.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(bytes pieceCid, uint64 paddedPieceSize, bool verifiedDeal, bytes client, bytes label, uint64 startEpoch, uint64 endEpoch, uint64 storagePricePerEpoch, uint64 providerCollateral, uint64 clientCollateral, string version, bytes params)
func (_DealClient *DealClientCallerSession) Proposals(arg0 [32]byte) (struct {
	PieceCid             []byte
	PaddedPieceSize      uint64
	VerifiedDeal         bool
	Client               []byte
	Label                []byte
	StartEpoch           uint64
	EndEpoch             uint64
	StoragePricePerEpoch uint64
	ProviderCollateral   uint64
	ClientCollateral     uint64
	Version              string
	Params               []byte
}, error) {
	return _DealClient.Contract.Proposals(&_DealClient.CallOpts, arg0)
}

// AddCID is a paid mutator transaction binding the contract method 0xd4a0cd0a.
//
// Solidity: function addCID(bytes cidraw, uint256 size) returns()
func (_DealClient *DealClientTransactor) AddCID(opts *bind.TransactOpts, cidraw []byte, size *big.Int) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "addCID", cidraw, size)
}

// AddCID is a paid mutator transaction binding the contract method 0xd4a0cd0a.
//
// Solidity: function addCID(bytes cidraw, uint256 size) returns()
func (_DealClient *DealClientSession) AddCID(cidraw []byte, size *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.AddCID(&_DealClient.TransactOpts, cidraw, size)
}

// AddCID is a paid mutator transaction binding the contract method 0xd4a0cd0a.
//
// Solidity: function addCID(bytes cidraw, uint256 size) returns()
func (_DealClient *DealClientTransactorSession) AddCID(cidraw []byte, size *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.AddCID(&_DealClient.TransactOpts, cidraw, size)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xa6b9cea6.
//
// Solidity: function authorizeData(bytes cidraw, bytes provider, uint256 size) returns()
func (_DealClient *DealClientTransactor) AuthorizeData(opts *bind.TransactOpts, cidraw []byte, provider []byte, size *big.Int) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "authorizeData", cidraw, provider, size)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xa6b9cea6.
//
// Solidity: function authorizeData(bytes cidraw, bytes provider, uint256 size) returns()
func (_DealClient *DealClientSession) AuthorizeData(cidraw []byte, provider []byte, size *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.AuthorizeData(&_DealClient.TransactOpts, cidraw, provider, size)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xa6b9cea6.
//
// Solidity: function authorizeData(bytes cidraw, bytes provider, uint256 size) returns()
func (_DealClient *DealClientTransactorSession) AuthorizeData(cidraw []byte, provider []byte, size *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.AuthorizeData(&_DealClient.TransactOpts, cidraw, provider, size)
}

// HandleFilecoinMethod is a paid mutator transaction binding the contract method 0x868e10c4.
//
// Solidity: function handle_filecoin_method(uint64 method, uint64 , bytes params) returns()
func (_DealClient *DealClientTransactor) HandleFilecoinMethod(opts *bind.TransactOpts, method uint64, arg1 uint64, params []byte) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "handle_filecoin_method", method, arg1, params)
}

// HandleFilecoinMethod is a paid mutator transaction binding the contract method 0x868e10c4.
//
// Solidity: function handle_filecoin_method(uint64 method, uint64 , bytes params) returns()
func (_DealClient *DealClientSession) HandleFilecoinMethod(method uint64, arg1 uint64, params []byte) (*types.Transaction, error) {
	return _DealClient.Contract.HandleFilecoinMethod(&_DealClient.TransactOpts, method, arg1, params)
}

// HandleFilecoinMethod is a paid mutator transaction binding the contract method 0x868e10c4.
//
// Solidity: function handle_filecoin_method(uint64 method, uint64 , bytes params) returns()
func (_DealClient *DealClientTransactorSession) HandleFilecoinMethod(method uint64, arg1 uint64, params []byte) (*types.Transaction, error) {
	return _DealClient.Contract.HandleFilecoinMethod(&_DealClient.TransactOpts, method, arg1, params)
}

// MakeDealProposal is a paid mutator transaction binding the contract method 0xc1117182.
//
// Solidity: function makeDealProposal((bytes,uint64,bool,bytes,bytes,uint64,uint64,uint64,uint64,uint64,string,bytes) _deal) returns(bytes32)
func (_DealClient *DealClientTransactor) MakeDealProposal(opts *bind.TransactOpts, _deal DealClientDealProposal) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "makeDealProposal", _deal)
}

// MakeDealProposal is a paid mutator transaction binding the contract method 0xc1117182.
//
// Solidity: function makeDealProposal((bytes,uint64,bool,bytes,bytes,uint64,uint64,uint64,uint64,uint64,string,bytes) _deal) returns(bytes32)
func (_DealClient *DealClientSession) MakeDealProposal(_deal DealClientDealProposal) (*types.Transaction, error) {
	return _DealClient.Contract.MakeDealProposal(&_DealClient.TransactOpts, _deal)
}

// MakeDealProposal is a paid mutator transaction binding the contract method 0xc1117182.
//
// Solidity: function makeDealProposal((bytes,uint64,bool,bytes,bytes,uint64,uint64,uint64,uint64,uint64,string,bytes) _deal) returns(bytes32)
func (_DealClient *DealClientTransactorSession) MakeDealProposal(_deal DealClientDealProposal) (*types.Transaction, error) {
	return _DealClient.Contract.MakeDealProposal(&_DealClient.TransactOpts, _deal)
}

// PublishDeal is a paid mutator transaction binding the contract method 0xd75fb3c8.
//
// Solidity: function publish_deal(bytes raw_auth_params, address callee) returns()
func (_DealClient *DealClientTransactor) PublishDeal(opts *bind.TransactOpts, raw_auth_params []byte, callee common.Address) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "publish_deal", raw_auth_params, callee)
}

// PublishDeal is a paid mutator transaction binding the contract method 0xd75fb3c8.
//
// Solidity: function publish_deal(bytes raw_auth_params, address callee) returns()
func (_DealClient *DealClientSession) PublishDeal(raw_auth_params []byte, callee common.Address) (*types.Transaction, error) {
	return _DealClient.Contract.PublishDeal(&_DealClient.TransactOpts, raw_auth_params, callee)
}

// PublishDeal is a paid mutator transaction binding the contract method 0xd75fb3c8.
//
// Solidity: function publish_deal(bytes raw_auth_params, address callee) returns()
func (_DealClient *DealClientTransactorSession) PublishDeal(raw_auth_params []byte, callee common.Address) (*types.Transaction, error) {
	return _DealClient.Contract.PublishDeal(&_DealClient.TransactOpts, raw_auth_params, callee)
}

// DealClientDealProposalCreateIterator is returned from FilterDealProposalCreate and is used to iterate over the raw logs and unpacked data for DealProposalCreate events raised by the DealClient contract.
type DealClientDealProposalCreateIterator struct {
	Event *DealClientDealProposalCreate // Event containing the contract specifics and raw log

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
func (it *DealClientDealProposalCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DealClientDealProposalCreate)
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
		it.Event = new(DealClientDealProposalCreate)
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
func (it *DealClientDealProposalCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DealClientDealProposalCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DealClientDealProposalCreate represents a DealProposalCreate event raised by the DealClient contract.
type DealClientDealProposalCreate struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDealProposalCreate is a free log retrieval operation binding the contract event 0xc15e53e899e372ce2933436c5b2c22ed67a40809a9d557a7846a6104fea03972.
//
// Solidity: event DealProposalCreate(bytes32 indexed id)
func (_DealClient *DealClientFilterer) FilterDealProposalCreate(opts *bind.FilterOpts, id [][32]byte) (*DealClientDealProposalCreateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DealClient.contract.FilterLogs(opts, "DealProposalCreate", idRule)
	if err != nil {
		return nil, err
	}
	return &DealClientDealProposalCreateIterator{contract: _DealClient.contract, event: "DealProposalCreate", logs: logs, sub: sub}, nil
}

// WatchDealProposalCreate is a free log subscription operation binding the contract event 0xc15e53e899e372ce2933436c5b2c22ed67a40809a9d557a7846a6104fea03972.
//
// Solidity: event DealProposalCreate(bytes32 indexed id)
func (_DealClient *DealClientFilterer) WatchDealProposalCreate(opts *bind.WatchOpts, sink chan<- *DealClientDealProposalCreate, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DealClient.contract.WatchLogs(opts, "DealProposalCreate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DealClientDealProposalCreate)
				if err := _DealClient.contract.UnpackLog(event, "DealProposalCreate", log); err != nil {
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

// ParseDealProposalCreate is a log parse operation binding the contract event 0xc15e53e899e372ce2933436c5b2c22ed67a40809a9d557a7846a6104fea03972.
//
// Solidity: event DealProposalCreate(bytes32 indexed id)
func (_DealClient *DealClientFilterer) ParseDealProposalCreate(log types.Log) (*DealClientDealProposalCreate, error) {
	event := new(DealClientDealProposalCreate)
	if err := _DealClient.contract.UnpackLog(event, "DealProposalCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DealClientReceivedDataCapIterator is returned from FilterReceivedDataCap and is used to iterate over the raw logs and unpacked data for ReceivedDataCap events raised by the DealClient contract.
type DealClientReceivedDataCapIterator struct {
	Event *DealClientReceivedDataCap // Event containing the contract specifics and raw log

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
func (it *DealClientReceivedDataCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DealClientReceivedDataCap)
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
		it.Event = new(DealClientReceivedDataCap)
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
func (it *DealClientReceivedDataCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DealClientReceivedDataCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DealClientReceivedDataCap represents a ReceivedDataCap event raised by the DealClient contract.
type DealClientReceivedDataCap struct {
	Received string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceivedDataCap is a free log retrieval operation binding the contract event 0x10aa319ed8cad9bceb033c0c2788c4ae17469ac844e4c6e2c2e20e74ca8a7be8.
//
// Solidity: event ReceivedDataCap(string received)
func (_DealClient *DealClientFilterer) FilterReceivedDataCap(opts *bind.FilterOpts) (*DealClientReceivedDataCapIterator, error) {

	logs, sub, err := _DealClient.contract.FilterLogs(opts, "ReceivedDataCap")
	if err != nil {
		return nil, err
	}
	return &DealClientReceivedDataCapIterator{contract: _DealClient.contract, event: "ReceivedDataCap", logs: logs, sub: sub}, nil
}

// WatchReceivedDataCap is a free log subscription operation binding the contract event 0x10aa319ed8cad9bceb033c0c2788c4ae17469ac844e4c6e2c2e20e74ca8a7be8.
//
// Solidity: event ReceivedDataCap(string received)
func (_DealClient *DealClientFilterer) WatchReceivedDataCap(opts *bind.WatchOpts, sink chan<- *DealClientReceivedDataCap) (event.Subscription, error) {

	logs, sub, err := _DealClient.contract.WatchLogs(opts, "ReceivedDataCap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DealClientReceivedDataCap)
				if err := _DealClient.contract.UnpackLog(event, "ReceivedDataCap", log); err != nil {
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

// ParseReceivedDataCap is a log parse operation binding the contract event 0x10aa319ed8cad9bceb033c0c2788c4ae17469ac844e4c6e2c2e20e74ca8a7be8.
//
// Solidity: event ReceivedDataCap(string received)
func (_DealClient *DealClientFilterer) ParseReceivedDataCap(log types.Log) (*DealClientReceivedDataCap, error) {
	event := new(DealClientReceivedDataCap)
	if err := _DealClient.contract.UnpackLog(event, "ReceivedDataCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
