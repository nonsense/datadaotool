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

// DealClientMetaData contains all meta data concerning the DealClient contract.
var DealClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"DealProposalCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"received\",\"type\":\"string\"}],\"name\":\"ReceivedDataCap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHENTICATE_MESSAGE_METHOD_NUM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATACAP_RECEIVER_HOOK_METHOD_NUM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET_NOTIFY_DEAL_METHOD_NUM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"client\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dealProposals\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"}],\"name\":\"getDealProposal\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"method\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"handle_filecoin_method\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"deal\",\"type\":\"bytes\"}],\"name\":\"makeDealProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pieceDeals\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pieceProviders\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pieceToProposal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pieceCid\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"pieceSize\",\"type\":\"uint64\"}],\"name\":\"simpleDealProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"client\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AUTHENTICATEMESSAGEMETHODNUM is a free data retrieval call binding the contract method 0x00706790.
//
// Solidity: function AUTHENTICATE_MESSAGE_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCaller) AUTHENTICATEMESSAGEMETHODNUM(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "AUTHENTICATE_MESSAGE_METHOD_NUM")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AUTHENTICATEMESSAGEMETHODNUM is a free data retrieval call binding the contract method 0x00706790.
//
// Solidity: function AUTHENTICATE_MESSAGE_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientSession) AUTHENTICATEMESSAGEMETHODNUM() (uint64, error) {
	return _DealClient.Contract.AUTHENTICATEMESSAGEMETHODNUM(&_DealClient.CallOpts)
}

// AUTHENTICATEMESSAGEMETHODNUM is a free data retrieval call binding the contract method 0x00706790.
//
// Solidity: function AUTHENTICATE_MESSAGE_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCallerSession) AUTHENTICATEMESSAGEMETHODNUM() (uint64, error) {
	return _DealClient.Contract.AUTHENTICATEMESSAGEMETHODNUM(&_DealClient.CallOpts)
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

// MARKETNOTIFYDEALMETHODNUM is a free data retrieval call binding the contract method 0x6067f454.
//
// Solidity: function MARKET_NOTIFY_DEAL_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCaller) MARKETNOTIFYDEALMETHODNUM(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "MARKET_NOTIFY_DEAL_METHOD_NUM")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MARKETNOTIFYDEALMETHODNUM is a free data retrieval call binding the contract method 0x6067f454.
//
// Solidity: function MARKET_NOTIFY_DEAL_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientSession) MARKETNOTIFYDEALMETHODNUM() (uint64, error) {
	return _DealClient.Contract.MARKETNOTIFYDEALMETHODNUM(&_DealClient.CallOpts)
}

// MARKETNOTIFYDEALMETHODNUM is a free data retrieval call binding the contract method 0x6067f454.
//
// Solidity: function MARKET_NOTIFY_DEAL_METHOD_NUM() view returns(uint64)
func (_DealClient *DealClientCallerSession) MARKETNOTIFYDEALMETHODNUM() (uint64, error) {
	return _DealClient.Contract.MARKETNOTIFYDEALMETHODNUM(&_DealClient.CallOpts)
}

// DealProposals is a free data retrieval call binding the contract method 0x74112d66.
//
// Solidity: function dealProposals(bytes32 ) view returns(bytes)
func (_DealClient *DealClientCaller) DealProposals(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "dealProposals", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DealProposals is a free data retrieval call binding the contract method 0x74112d66.
//
// Solidity: function dealProposals(bytes32 ) view returns(bytes)
func (_DealClient *DealClientSession) DealProposals(arg0 [32]byte) ([]byte, error) {
	return _DealClient.Contract.DealProposals(&_DealClient.CallOpts, arg0)
}

// DealProposals is a free data retrieval call binding the contract method 0x74112d66.
//
// Solidity: function dealProposals(bytes32 ) view returns(bytes)
func (_DealClient *DealClientCallerSession) DealProposals(arg0 [32]byte) ([]byte, error) {
	return _DealClient.Contract.DealProposals(&_DealClient.CallOpts, arg0)
}

// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
//
// Solidity: function getDealProposal(bytes32 proposalId) view returns(bytes)
func (_DealClient *DealClientCaller) GetDealProposal(opts *bind.CallOpts, proposalId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "getDealProposal", proposalId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
//
// Solidity: function getDealProposal(bytes32 proposalId) view returns(bytes)
func (_DealClient *DealClientSession) GetDealProposal(proposalId [32]byte) ([]byte, error) {
	return _DealClient.Contract.GetDealProposal(&_DealClient.CallOpts, proposalId)
}

// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
//
// Solidity: function getDealProposal(bytes32 proposalId) view returns(bytes)
func (_DealClient *DealClientCallerSession) GetDealProposal(proposalId [32]byte) ([]byte, error) {
	return _DealClient.Contract.GetDealProposal(&_DealClient.CallOpts, proposalId)
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

// PieceDeals is a free data retrieval call binding the contract method 0x0a0e0c91.
//
// Solidity: function pieceDeals(bytes ) view returns(uint64)
func (_DealClient *DealClientCaller) PieceDeals(opts *bind.CallOpts, arg0 []byte) (uint64, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "pieceDeals", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PieceDeals is a free data retrieval call binding the contract method 0x0a0e0c91.
//
// Solidity: function pieceDeals(bytes ) view returns(uint64)
func (_DealClient *DealClientSession) PieceDeals(arg0 []byte) (uint64, error) {
	return _DealClient.Contract.PieceDeals(&_DealClient.CallOpts, arg0)
}

// PieceDeals is a free data retrieval call binding the contract method 0x0a0e0c91.
//
// Solidity: function pieceDeals(bytes ) view returns(uint64)
func (_DealClient *DealClientCallerSession) PieceDeals(arg0 []byte) (uint64, error) {
	return _DealClient.Contract.PieceDeals(&_DealClient.CallOpts, arg0)
}

// PieceProviders is a free data retrieval call binding the contract method 0xf82704f0.
//
// Solidity: function pieceProviders(bytes ) view returns(bytes provider, bool valid)
func (_DealClient *DealClientCaller) PieceProviders(opts *bind.CallOpts, arg0 []byte) (struct {
	Provider []byte
	Valid    bool
}, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "pieceProviders", arg0)

	outstruct := new(struct {
		Provider []byte
		Valid    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Provider = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// PieceProviders is a free data retrieval call binding the contract method 0xf82704f0.
//
// Solidity: function pieceProviders(bytes ) view returns(bytes provider, bool valid)
func (_DealClient *DealClientSession) PieceProviders(arg0 []byte) (struct {
	Provider []byte
	Valid    bool
}, error) {
	return _DealClient.Contract.PieceProviders(&_DealClient.CallOpts, arg0)
}

// PieceProviders is a free data retrieval call binding the contract method 0xf82704f0.
//
// Solidity: function pieceProviders(bytes ) view returns(bytes provider, bool valid)
func (_DealClient *DealClientCallerSession) PieceProviders(arg0 []byte) (struct {
	Provider []byte
	Valid    bool
}, error) {
	return _DealClient.Contract.PieceProviders(&_DealClient.CallOpts, arg0)
}

// PieceToProposal is a free data retrieval call binding the contract method 0xdd08881c.
//
// Solidity: function pieceToProposal(bytes ) view returns(bytes32 proposalId, bool valid)
func (_DealClient *DealClientCaller) PieceToProposal(opts *bind.CallOpts, arg0 []byte) (struct {
	ProposalId [32]byte
	Valid      bool
}, error) {
	var out []interface{}
	err := _DealClient.contract.Call(opts, &out, "pieceToProposal", arg0)

	outstruct := new(struct {
		ProposalId [32]byte
		Valid      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// PieceToProposal is a free data retrieval call binding the contract method 0xdd08881c.
//
// Solidity: function pieceToProposal(bytes ) view returns(bytes32 proposalId, bool valid)
func (_DealClient *DealClientSession) PieceToProposal(arg0 []byte) (struct {
	ProposalId [32]byte
	Valid      bool
}, error) {
	return _DealClient.Contract.PieceToProposal(&_DealClient.CallOpts, arg0)
}

// PieceToProposal is a free data retrieval call binding the contract method 0xdd08881c.
//
// Solidity: function pieceToProposal(bytes ) view returns(bytes32 proposalId, bool valid)
func (_DealClient *DealClientCallerSession) PieceToProposal(arg0 []byte) (struct {
	ProposalId [32]byte
	Valid      bool
}, error) {
	return _DealClient.Contract.PieceToProposal(&_DealClient.CallOpts, arg0)
}

// AddBalance is a paid mutator transaction binding the contract method 0xc71c6490.
//
// Solidity: function addBalance(bytes client, uint256 value) returns()
func (_DealClient *DealClientTransactor) AddBalance(opts *bind.TransactOpts, client []byte, value *big.Int) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "addBalance", client, value)
}

// AddBalance is a paid mutator transaction binding the contract method 0xc71c6490.
//
// Solidity: function addBalance(bytes client, uint256 value) returns()
func (_DealClient *DealClientSession) AddBalance(client []byte, value *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.AddBalance(&_DealClient.TransactOpts, client, value)
}

// AddBalance is a paid mutator transaction binding the contract method 0xc71c6490.
//
// Solidity: function addBalance(bytes client, uint256 value) returns()
func (_DealClient *DealClientTransactorSession) AddBalance(client []byte, value *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.AddBalance(&_DealClient.TransactOpts, client, value)
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

// MakeDealProposal is a paid mutator transaction binding the contract method 0x5abba064.
//
// Solidity: function makeDealProposal(bytes deal) returns()
func (_DealClient *DealClientTransactor) MakeDealProposal(opts *bind.TransactOpts, deal []byte) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "makeDealProposal", deal)
}

// MakeDealProposal is a paid mutator transaction binding the contract method 0x5abba064.
//
// Solidity: function makeDealProposal(bytes deal) returns()
func (_DealClient *DealClientSession) MakeDealProposal(deal []byte) (*types.Transaction, error) {
	return _DealClient.Contract.MakeDealProposal(&_DealClient.TransactOpts, deal)
}

// MakeDealProposal is a paid mutator transaction binding the contract method 0x5abba064.
//
// Solidity: function makeDealProposal(bytes deal) returns()
func (_DealClient *DealClientTransactorSession) MakeDealProposal(deal []byte) (*types.Transaction, error) {
	return _DealClient.Contract.MakeDealProposal(&_DealClient.TransactOpts, deal)
}

// SimpleDealProposal is a paid mutator transaction binding the contract method 0xa70d3a9e.
//
// Solidity: function simpleDealProposal(bytes pieceCid, uint64 pieceSize) returns()
func (_DealClient *DealClientTransactor) SimpleDealProposal(opts *bind.TransactOpts, pieceCid []byte, pieceSize uint64) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "simpleDealProposal", pieceCid, pieceSize)
}

// SimpleDealProposal is a paid mutator transaction binding the contract method 0xa70d3a9e.
//
// Solidity: function simpleDealProposal(bytes pieceCid, uint64 pieceSize) returns()
func (_DealClient *DealClientSession) SimpleDealProposal(pieceCid []byte, pieceSize uint64) (*types.Transaction, error) {
	return _DealClient.Contract.SimpleDealProposal(&_DealClient.TransactOpts, pieceCid, pieceSize)
}

// SimpleDealProposal is a paid mutator transaction binding the contract method 0xa70d3a9e.
//
// Solidity: function simpleDealProposal(bytes pieceCid, uint64 pieceSize) returns()
func (_DealClient *DealClientTransactorSession) SimpleDealProposal(pieceCid []byte, pieceSize uint64) (*types.Transaction, error) {
	return _DealClient.Contract.SimpleDealProposal(&_DealClient.TransactOpts, pieceCid, pieceSize)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xfbe22ca3.
//
// Solidity: function withdrawBalance(bytes client, uint256 value) returns(uint256)
func (_DealClient *DealClientTransactor) WithdrawBalance(opts *bind.TransactOpts, client []byte, value *big.Int) (*types.Transaction, error) {
	return _DealClient.contract.Transact(opts, "withdrawBalance", client, value)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xfbe22ca3.
//
// Solidity: function withdrawBalance(bytes client, uint256 value) returns(uint256)
func (_DealClient *DealClientSession) WithdrawBalance(client []byte, value *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.WithdrawBalance(&_DealClient.TransactOpts, client, value)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xfbe22ca3.
//
// Solidity: function withdrawBalance(bytes client, uint256 value) returns(uint256)
func (_DealClient *DealClientTransactorSession) WithdrawBalance(client []byte, value *big.Int) (*types.Transaction, error) {
	return _DealClient.Contract.WithdrawBalance(&_DealClient.TransactOpts, client, value)
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
	Id       [32]byte
	Size     uint64
	Verified bool
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDealProposalCreate is a free log retrieval operation binding the contract event 0xfd6419d07e4c269e58d0c63969756c2124155b4a8d6dd08b8cd46e3a9acbf625.
//
// Solidity: event DealProposalCreate(bytes32 indexed id, uint64 size, bool indexed verified, uint256 price)
func (_DealClient *DealClientFilterer) FilterDealProposalCreate(opts *bind.FilterOpts, id [][32]byte, verified []bool) (*DealClientDealProposalCreateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var verifiedRule []interface{}
	for _, verifiedItem := range verified {
		verifiedRule = append(verifiedRule, verifiedItem)
	}

	logs, sub, err := _DealClient.contract.FilterLogs(opts, "DealProposalCreate", idRule, verifiedRule)
	if err != nil {
		return nil, err
	}
	return &DealClientDealProposalCreateIterator{contract: _DealClient.contract, event: "DealProposalCreate", logs: logs, sub: sub}, nil
}

// WatchDealProposalCreate is a free log subscription operation binding the contract event 0xfd6419d07e4c269e58d0c63969756c2124155b4a8d6dd08b8cd46e3a9acbf625.
//
// Solidity: event DealProposalCreate(bytes32 indexed id, uint64 size, bool indexed verified, uint256 price)
func (_DealClient *DealClientFilterer) WatchDealProposalCreate(opts *bind.WatchOpts, sink chan<- *DealClientDealProposalCreate, id [][32]byte, verified []bool) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var verifiedRule []interface{}
	for _, verifiedItem := range verified {
		verifiedRule = append(verifiedRule, verifiedItem)
	}

	logs, sub, err := _DealClient.contract.WatchLogs(opts, "DealProposalCreate", idRule, verifiedRule)
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

// ParseDealProposalCreate is a log parse operation binding the contract event 0xfd6419d07e4c269e58d0c63969756c2124155b4a8d6dd08b8cd46e3a9acbf625.
//
// Solidity: event DealProposalCreate(bytes32 indexed id, uint64 size, bool indexed verified, uint256 price)
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
