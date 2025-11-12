// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package task1

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
  _ = abi.ConvertType
)

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
  ABI: "[{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTokenType\",\"name\":\"symbol\",\"type\":\"uint8\"}],\"name\":\"getChainlinkDataFeedLatestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceDropInterval\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"}],\"name\":\"putOnShelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeFromShelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
  AuctionCaller     // Read-only binding to the contract
  AuctionTransactor // Write-only binding to the contract
  AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
  contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
  contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
  contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
  Contract     *Auction          // Generic contract binding to set the session for
  CallOpts     bind.CallOpts     // Call options to use throughout this session
  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
  Contract *AuctionCaller // Generic contract caller binding to set the session for
  CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
  Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
  TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
  Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
  Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
  Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
  contract, err := bindAuction(address, backend, backend, backend)
  if err != nil {
    return nil, err
  }
  return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
  contract, err := bindAuction(address, caller, nil, nil)
  if err != nil {
    return nil, err
  }
  return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
  contract, err := bindAuction(address, nil, transactor, nil)
  if err != nil {
    return nil, err
  }
  return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
  contract, err := bindAuction(address, nil, nil, filterer)
  if err != nil {
    return nil, err
  }
  return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
  parsed, err := AuctionMetaData.GetAbi()
  if err != nil {
    return nil, err
  }
  return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
  return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
  return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
  return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
  return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
  return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
  return _Auction.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Auction *AuctionCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
  var out []interface{}
  err := _Auction.contract.Call(opts, &out, "admin")

  if err != nil {
    return *new(common.Address), err
  }

  out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

  return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Auction *AuctionSession) Admin() (common.Address, error) {
  return _Auction.Contract.Admin(&_Auction.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Auction *AuctionCallerSession) Admin() (common.Address, error) {
  return _Auction.Contract.Admin(&_Auction.CallOpts)
}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xc581ccc9.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(uint8 symbol) view returns(int256 answer, uint8 decimal, uint256 createdAt, uint256 updatedAt)
func (_Auction *AuctionCaller) GetChainlinkDataFeedLatestAnswer(opts *bind.CallOpts, symbol uint8) (struct {
  Answer    *big.Int
  Decimal   uint8
  CreatedAt *big.Int
  UpdatedAt *big.Int
}, error) {
  var out []interface{}
  err := _Auction.contract.Call(opts, &out, "getChainlinkDataFeedLatestAnswer", symbol)

  outstruct := new(struct {
    Answer    *big.Int
    Decimal   uint8
    CreatedAt *big.Int
    UpdatedAt *big.Int
  })
  if err != nil {
    return *outstruct, err
  }

  outstruct.Answer = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
  outstruct.Decimal = *abi.ConvertType(out[1], new(uint8)).(*uint8)
  outstruct.CreatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
  outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

  return *outstruct, err

}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xc581ccc9.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(uint8 symbol) view returns(int256 answer, uint8 decimal, uint256 createdAt, uint256 updatedAt)
func (_Auction *AuctionSession) GetChainlinkDataFeedLatestAnswer(symbol uint8) (struct {
  Answer    *big.Int
  Decimal   uint8
  CreatedAt *big.Int
  UpdatedAt *big.Int
}, error) {
  return _Auction.Contract.GetChainlinkDataFeedLatestAnswer(&_Auction.CallOpts, symbol)
}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xc581ccc9.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(uint8 symbol) view returns(int256 answer, uint8 decimal, uint256 createdAt, uint256 updatedAt)
func (_Auction *AuctionCallerSession) GetChainlinkDataFeedLatestAnswer(symbol uint8) (struct {
  Answer    *big.Int
  Decimal   uint8
  CreatedAt *big.Int
  UpdatedAt *big.Int
}, error) {
  return _Auction.Contract.GetChainlinkDataFeedLatestAnswer(&_Auction.CallOpts, symbol)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Auction *AuctionCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
  var out []interface{}
  err := _Auction.contract.Call(opts, &out, "implementation")

  if err != nil {
    return *new(common.Address), err
  }

  out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

  return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Auction *AuctionSession) Implementation() (common.Address, error) {
  return _Auction.Contract.Implementation(&_Auction.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Auction *AuctionCallerSession) Implementation() (common.Address, error) {
  return _Auction.Contract.Implementation(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
  var out []interface{}
  err := _Auction.contract.Call(opts, &out, "nft")

  if err != nil {
    return *new(common.Address), err
  }

  out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

  return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionSession) Nft() (common.Address, error) {
  return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCallerSession) Nft() (common.Address, error) {
  return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 tokenId) payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
  return _Auction.contract.Transact(opts, "bid", tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 tokenId) payable returns()
func (_Auction *AuctionSession) Bid(tokenId *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.Bid(&_Auction.TransactOpts, tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 tokenId) payable returns()
func (_Auction *AuctionTransactorSession) Bid(tokenId *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.Bid(&_Auction.TransactOpts, tokenId)
}

// End is a paid mutator transaction binding the contract method 0x0ad24528.
//
// Solidity: function end(uint256 limit) payable returns(uint256)
func (_Auction *AuctionTransactor) End(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
  return _Auction.contract.Transact(opts, "end", limit)
}

// End is a paid mutator transaction binding the contract method 0x0ad24528.
//
// Solidity: function end(uint256 limit) payable returns(uint256)
func (_Auction *AuctionSession) End(limit *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.End(&_Auction.TransactOpts, limit)
}

// End is a paid mutator transaction binding the contract method 0x0ad24528.
//
// Solidity: function end(uint256 limit) payable returns(uint256)
func (_Auction *AuctionTransactorSession) End(limit *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.End(&_Auction.TransactOpts, limit)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _admin, address nftAddr, uint256 _startTime, uint256 _endTime, uint256 _priceDropInterval) returns()
func (_Auction *AuctionTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, nftAddr common.Address, _startTime *big.Int, _endTime *big.Int, _priceDropInterval *big.Int) (*types.Transaction, error) {
  return _Auction.contract.Transact(opts, "initialize", _admin, nftAddr, _startTime, _endTime, _priceDropInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _admin, address nftAddr, uint256 _startTime, uint256 _endTime, uint256 _priceDropInterval) returns()
func (_Auction *AuctionSession) Initialize(_admin common.Address, nftAddr common.Address, _startTime *big.Int, _endTime *big.Int, _priceDropInterval *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.Initialize(&_Auction.TransactOpts, _admin, nftAddr, _startTime, _endTime, _priceDropInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _admin, address nftAddr, uint256 _startTime, uint256 _endTime, uint256 _priceDropInterval) returns()
func (_Auction *AuctionTransactorSession) Initialize(_admin common.Address, nftAddr common.Address, _startTime *big.Int, _endTime *big.Int, _priceDropInterval *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.Initialize(&_Auction.TransactOpts, _admin, nftAddr, _startTime, _endTime, _priceDropInterval)
}

// PutOnShelf is a paid mutator transaction binding the contract method 0x780ee92e.
//
// Solidity: function putOnShelf(uint256 tokenId, uint256 maxPrice, uint256 minPrice) returns()
func (_Auction *AuctionTransactor) PutOnShelf(opts *bind.TransactOpts, tokenId *big.Int, maxPrice *big.Int, minPrice *big.Int) (*types.Transaction, error) {
  return _Auction.contract.Transact(opts, "putOnShelf", tokenId, maxPrice, minPrice)
}

// PutOnShelf is a paid mutator transaction binding the contract method 0x780ee92e.
//
// Solidity: function putOnShelf(uint256 tokenId, uint256 maxPrice, uint256 minPrice) returns()
func (_Auction *AuctionSession) PutOnShelf(tokenId *big.Int, maxPrice *big.Int, minPrice *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.PutOnShelf(&_Auction.TransactOpts, tokenId, maxPrice, minPrice)
}

// PutOnShelf is a paid mutator transaction binding the contract method 0x780ee92e.
//
// Solidity: function putOnShelf(uint256 tokenId, uint256 maxPrice, uint256 minPrice) returns()
func (_Auction *AuctionTransactorSession) PutOnShelf(tokenId *big.Int, maxPrice *big.Int, minPrice *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.PutOnShelf(&_Auction.TransactOpts, tokenId, maxPrice, minPrice)
}

// RemoveFromShelf is a paid mutator transaction binding the contract method 0x00769627.
//
// Solidity: function removeFromShelf(uint256 tokenId) returns()
func (_Auction *AuctionTransactor) RemoveFromShelf(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
  return _Auction.contract.Transact(opts, "removeFromShelf", tokenId)
}

// RemoveFromShelf is a paid mutator transaction binding the contract method 0x00769627.
//
// Solidity: function removeFromShelf(uint256 tokenId) returns()
func (_Auction *AuctionSession) RemoveFromShelf(tokenId *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.RemoveFromShelf(&_Auction.TransactOpts, tokenId)
}

// RemoveFromShelf is a paid mutator transaction binding the contract method 0x00769627.
//
// Solidity: function removeFromShelf(uint256 tokenId) returns()
func (_Auction *AuctionTransactorSession) RemoveFromShelf(tokenId *big.Int) (*types.Transaction, error) {
  return _Auction.Contract.RemoveFromShelf(&_Auction.TransactOpts, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
  return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
  return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
  return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}
