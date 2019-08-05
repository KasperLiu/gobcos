// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package client provides a client for the FISCO BCOS RPC API.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"gobcos/rpc"
	"math/big"
	"gobcos/core/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// Client defines typed wrappers for the Ethereum RPC API. 
type Client struct {
	c       *rpc.Client
	groupID uint
}

// Dial connects a client to the given URL and groupID.
func Dial(rawurl string, groupID uint) (*Client, error) {
	return DialContext(context.Background(), rawurl, groupID)
}

// DialContext pass the context to the rpc client
func DialContext(ctx context.Context, rawurl string, groupID uint) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c, groupID), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, groupID uint) *Client {
	return &Client{c: c, groupID: groupID}
}

// Close disconnects the rpc
func (gc *Client) Close() {
	gc.c.Close()
}

// Blockchain Access

// ChainID retrieves the current chain ID for transaction replay protection.
func (gc *Client) ChainID(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := gc.c.CallContext(ctx, &result, "eth_chainId")
	if err != nil {
		return nil, err
	}
	return (*big.Int)(&result), err
}

type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}


// HeaderByHash returns the block header with the given hash.
func (gc *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	var head *types.Header
	err := gc.c.CallContext(ctx, &head, "eth_getBlockByHash", hash, false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}
	return head, err
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (gc *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var head *types.Header
	err := gc.c.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}
	return head, err
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

// TransactionCount returns the total number of transactions in the given block.
func (gc *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	var num hexutil.Uint
	err := gc.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByHash", blockHash)
	return uint(num), err
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (gc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := gc.c.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

type rpcProgress struct {
	StartingBlock hexutil.Uint64
	CurrentBlock  hexutil.Uint64
	HighestBlock  hexutil.Uint64
	PulledStates  hexutil.Uint64
	KnownStates   hexutil.Uint64
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (gc *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	var raw json.RawMessage
	if err := gc.c.CallContext(ctx, &raw, "eth_syncing"); err != nil {
		return nil, err
	}
	// Handle the possible response types
	var syncing bool
	if err := json.Unmarshal(raw, &syncing); err == nil {
		return nil, nil // Not syncing (always false)
	}
	var progress *rpcProgress
	if err := json.Unmarshal(raw, &progress); err != nil {
		return nil, err
	}
	return &ethereum.SyncProgress{
		StartingBlock: uint64(progress.StartingBlock),
		CurrentBlock:  uint64(progress.CurrentBlock),
		HighestBlock:  uint64(progress.HighestBlock),
		PulledStates:  uint64(progress.PulledStates),
		KnownStates:   uint64(progress.KnownStates),
	}, nil
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (gc *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return gc.c.EthSubscribe(ctx, ch, "newHeads")
}

// State Access

// NetworkID returns the network ID (also known as the chain ID) for this chain.
func (gc *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	version := new(big.Int)
	var ver string
	if err := gc.c.CallContext(ctx, &ver, "net_version"); err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}

// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (gc *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := gc.c.CallContext(ctx, &result, "eth_getBalance", account, toBlockNumArg(blockNumber))
	return (*big.Int)(&result), err
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (gc *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := gc.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, toBlockNumArg(blockNumber))
	return result, err
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (gc *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var result hexutil.Uint64
	err := gc.c.CallContext(ctx, &result, "eth_getTransactionCount", account, toBlockNumArg(blockNumber))
	return uint64(result), err
}

// Pending State

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (gc *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	var result hexutil.Big
	err := gc.c.CallContext(ctx, &result, "eth_getBalance", account, "pending")
	return (*big.Int)(&result), err
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (gc *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	var result hexutil.Bytes
	err := gc.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, "pending")
	return result, err
}



// PendingTransactionCount returns the total number of transactions in the pending state.
func (gc *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	var num hexutil.Uint
	err := gc.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "pending")
	return uint(num), err
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.


// ========================== edit by KasperLiu ==========================


// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (gc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	// ======================================== KasperLiu =========================================
	err := gc.c.CallContext(ctx, &result, "getCode", gc.groupID, account.String())
	return result, err
}


// Filters

// FilterLogs executes a filter query.
func (gc *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = gc.c.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (gc *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	return gc.c.EthSubscribe(ctx, ch, "logs", arg)
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{
		"address": q.Addresses,
		"topics":  q.Topics,
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

// Pending State

// PendingCodeAt returns the contract code of the given account in the pending state.
func (gc *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var result hexutil.Bytes
	// =================================== KasperLiu ===============================
	// err := gc.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
	err := gc.c.CallContext(ctx, &result, "getCode", gc.groupID, account.String())
	return result, err
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (gc *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var result hexutil.Uint64
	err := gc.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")
	return uint64(result), err
}


// Contract Calling

type callResult struct {
	CurrentBlockNumber string `json:"currentBlockNumber"`
	Output             string `json:"output"`
	Status             string `json:"status"`
}

// CallContract invoke the call method of rpc api
func (gc *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	var cr *callResult
	// err := gc.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	err := gc.c.CallContext(ctx, &cr, "call", gc.groupID, toCallArg(msg))
	if err != nil {
		return nil, err
	}
	hex = common.FromHex(cr.Output)
	return hex, nil
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (gc *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	// ==================================== KasperLiu ===================================
	// err := gc.c.CallContext(ctx, &hex, "eht_call", toCallArg(msg), "pending")
	err := gc.c.CallContext(ctx, &hex, "call", gc.groupID, toCallArg(msg))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (gc *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var hex hexutil.Big
	if err := gc.c.CallContext(ctx, &hex, "eth_gasPrice"); err != nil {
		return nil, err
	}
	return (*big.Int)(&hex), nil
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (gc *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	var hex hexutil.Uint64
	err := gc.c.CallContext(ctx, &hex, "eth_estimateGas", toCallArg(msg))
	if err != nil {
		return 0, err
	}
	return uint64(hex), nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (gc *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	// ============================= KasperLiu ====================================
	// return gc.c.CallContext(ctx, nil, "eth_sendRawTransaction", common.ToHex(data))
	return gc.c.CallContext(ctx, nil, "sendRawTransaction", gc.groupID, common.ToHex(data))
}

// ==================================== edit by KasperLiu =================================
func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From.String(),
		"to":   msg.To.String(),
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data).String()
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value).String()
	}

	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}

	return arg
}

// GetGroupID returns the groupID of the client
func (gc *Client) GetGroupID() uint {
	return gc.groupID
}

// SetGroupID sets the groupID of the client
func (gc *Client) SetGroupID(newID uint) {
	gc.groupID = newID
}

// ============================================== FISCO BCOS Blockchain Access ================================================

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (gc *Client) GetClientVersion(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getClientVersion")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (gc *Client) GetBlockNumber(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockNumber", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (gc *Client) GetPBFTView(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPbftView", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err

	// TODO
	// Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (gc *Client) GetSealerList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getSealerList", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (gc *Client) GetObserverList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getObserverList", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (gc *Client) GetConsensusStatus(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getConsensusStatus", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetSyncStatus returns the synchronization status of the group
func (gc *Client) GetSyncStatus(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getSyncStatus", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPeers returns the information of the connected peers
func (gc *Client) GetPeers(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPeers", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (gc *Client) GetGroupPeers(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getGroupPeers", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (gc *Client) GetNodeIDList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getNodeIDList", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetGroupList returns the groupID list that the node belongs to
func (gc *Client) GetGroupList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getGroupList")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockByHash returns the block information according to the given block hash
func (gc *Client) GetBlockByHash(ctx context.Context, bhash string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockByHash", gc.groupID, bhash, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (gc *Client) GetBlockByNumber(ctx context.Context, bnum string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockByNumber", gc.groupID, bnum, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (gc *Client) GetBlockHashByNumber(ctx context.Context, bnum string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockHashByNumber", gc.groupID, bnum)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (gc *Client) GetTransactionByHash(ctx context.Context, txhash string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionByHash", gc.groupID, txhash)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByBlockHashAndIndex returns the transaction information according to
// the given block hash and transaction index
func (gc *Client) GetTransactionByBlockHashAndIndex(ctx context.Context, bhash string, txindex string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionByBlockHashAndIndex", gc.groupID, bhash, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to
// the given block number and transaction index
func (gc *Client) GetTransactionByBlockNumberAndIndex(ctx context.Context, bnum string, txindex string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionByBlockNumberAndIndex", gc.groupID, bnum, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (gc *Client) GetTransactionReceipt(ctx context.Context, txhash string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionReceipt", gc.groupID, txhash)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPendingTransactions returns information of the pending transactions
func (gc *Client) GetPendingTransactions(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPendingTransactions", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPendingTxSize returns amount of the pending transactions
func (gc *Client) GetPendingTxSize(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPendingTxSize", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetCode returns the contract code according to the contract address
func (gc *Client) GetCode(ctx context.Context, addr string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getCode", gc.groupID, addr)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (gc *Client) GetTotalTransactionCount(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTotalTransactionCount", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (gc *Client) GetSystemConfigByKey(ctx context.Context, findkey string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getSystemConfigByKey", gc.groupID, findkey)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}
