/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package console

import (
	"fmt"
	"context"
	"strconv"

	"github.com/spf13/cobra"
)


// commands
// var bashCompletionCmd = &cobra.Command{
// 	Use:   "bashCompletion",
// 	Short: "Generates bash completion scripts",
// 	Long: `A script "gobcos.sh" will get you completions of the console commands.
// Copy it to 

//     /etc/bash_completion.d/ 

// as described here:

//     https://debian-administration.org/article/316/An_introduction_to_bash_completion_part_1

// and reset your terminal to use autocompletion.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		rootCmd.GenBashCompletionFile("gobcos.sh");
// 		fmt.Println("gobcos.sh created on your current diretory successfully.")
// 	},
// }

// var zshCompletionCmd = &cobra.Command{
// 	Use:   "zshCompletion",
// 	Short: "Generates zsh completion scripts",
// 	Long: `A script "gobcos.zsh" will get you completions of the console commands.
// The recommended way to install this script is to copy to '~/.zsh/_console', and
// then add the following to your ~/.zshrc file:

//     fpath=(~/.zsh $fpath)

// as described here:

//     https://debian-administration.org/article/316/An_introduction_to_bash_completion_part_1

// and reset your terminal to use autocompletion.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		rootCmd.GenZshCompletionFile("_console");
// 		fmt.Println("zsh file _console had created on your current diretory successfully.")
// 	},
// }

// =========== account ==========
var newAccountCmd = &cobra.Command{
	Use:   "newAccount",
	Short: "Create a new account",
	Long: `Create a new account and save the results to ./bin/account/yourAccountName.keystore in encrypted form.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		clientVer, err := RPC.GetClientVersion(context.Background())
		if err != nil {
			fmt.Printf("client version not found: %v\n", err)
			return
		}
		fmt.Printf("Client Version: \n%s\n" , clientVer)
	},
}


// ======= node =======

var getClientVersionCmd = &cobra.Command{
	Use:   "getClientVersion",
	Short: "Get the blockchain version",
	Long: `Returns the specific FISCO BCOS version that runs on the node you connected.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		clientVer, err := RPC.GetClientVersion(context.Background())
		if err != nil {
			fmt.Printf("client version not found: %v\n", err)
			return
		}
		fmt.Printf("Client Version: \n%s\n" , clientVer)
	},
}

var getGroupIDCmd = &cobra.Command{
	Use:   "getGroupID",
	Short: "Get the group ID of the client",
	Long: `Returns the group ID that the console had connected to.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		groupID := RPC.GetGroupID()
		fmt.Printf("Group ID: \n%s\n" , groupID)
	},
}

var getBlockNumberCmd = &cobra.Command{
	Use:   "getBlockNumber",
	Short: "Get the latest block height of the blockchain",
	Long: `Returns the latest block height in the specified group.
The block height is encoded in hex`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		blockNumber,err := RPC.GetBlockNumber(context.Background())
		if err != nil {
			fmt.Printf("block number not found: %v\n", err)
			return
		}
		fmt.Printf("blocknumber: \n%s\n", blockNumber)
	},
}

var getPbftViewCmd = &cobra.Command{
	Use:   "getPbftView",
	Short: "Get the latest PBFT view (only support under PBFT consensus)",
	Long: `Returns the latest PBFT view in the specified group where the node is located.
The PBFT view is encoded in hex`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		pbft,err := RPC.GetPBFTView(context.Background())
		if err != nil {
			fmt.Printf("PBFT view not found: %v\n", err)
			return
		}
		fmt.Printf("PBFT view: \n%s\n" , pbft)
	},
}

var getSealerListCmd = &cobra.Command{
	Use:   "getSealerList",
	Short: "Get the sealers' ID list",
	Long: `Returns an ID list of the sealer nodes within the specified group.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		sealerList,err := RPC.GetSealerList(context.Background())
		if err != nil {
			fmt.Printf("sealer list not found: %v\n", err)
			return
		}
		fmt.Printf("Sealer List: \n%s\n" , sealerList)
	},
}

var getObserverListCmd = &cobra.Command{
	Use:   "getObserverList",
	Short: "Get the observers' ID list",
	Long: `Returns an ID list of observer nodes within the specified group.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		observerList,err := RPC.GetObserverList(context.Background())
		if err != nil {
			fmt.Printf("observer list not found: %v\n", err)
			return
		}
		fmt.Printf("Observer List: \n%s\n" , observerList)
	},
}

var getConsensusStatusCmd = &cobra.Command{
	Use:   "getConsensusStatus",
	Short: "Get consensus status of nodes",
	Long: `Returns consensus status information within the specified group.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		consensusStatus,err := RPC.GetConsensusStatus(context.Background())
		if err != nil {
			fmt.Printf("consensus status not found: %v\n", err)
			return
		}
		fmt.Printf("Consensus Status: \n%s\n" , consensusStatus)
	},
}

var getSyncStatusCmd = &cobra.Command{
	Use:   "getSyncStatus",
	Short: "Get synchronization status of nodes",
	Long: `Returns synchronization status information within the specified group.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		syncStatus,err := RPC.GetSyncStatus(context.Background())
		if err != nil {
			fmt.Printf("synchronization status not found: %v\n", err)
			return
		}
		fmt.Printf("Synchronization Status: \n%s\n" , syncStatus)
	},
}

var getPeersCmd = &cobra.Command{
	Use:   "getPeers",
	Short: "Get the connected peers' information",
	Long: `Returns the information of connected p2p nodes.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers,err := RPC.GetPeers(context.Background())
		if err != nil {
			fmt.Printf("peers not found: %v\n", err)
			return
		}
		fmt.Printf("Peers: \n%s\n" , peers)
	},
}

var getGroupPeersCmd = &cobra.Command{
	Use:   "getGroupPeers",
	Short: "Get all peers' ID within the group",
	Long: `Returns an ID list of consensus nodes and observer nodes within the specified group.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers,err := RPC.GetGroupPeers(context.Background())
		if err != nil {
			fmt.Printf("peers not found: %v\n", err)
			return
		}
		fmt.Printf("Peers: \n%s\n" , peers)
	},
}

var getNodeIDListCmd = &cobra.Command{
	Use:   "getNodeIDList",
	Short: "Get ID list of nodes",
	Long: `Returns an ID list of the node itself and the connected p2p nodes.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers,err := RPC.GetNodeIDList(context.Background())
		if err != nil {
			fmt.Printf("node ID list not found: %v\n", err)
			return
		}
		fmt.Printf("Node ID list: \n%s\n" , peers)
	},
}

var getGroupListCmd = &cobra.Command{
	Use:   "getGroupList",
	Short: "Get ID list of groups that the node belongs",
	Long: `Returns an ID list of groups that the node belongs.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers,err := RPC.GetGroupList(context.Background())
		if err != nil {
			fmt.Printf("group IDs list not found: %v\n", err)
			return
		}
		fmt.Printf("Group ID List: \n%s\n" , peers)
	},
}

// ========= block access ==========

var getBlockByHashCmd = &cobra.Command{
	Use:   "getBlockByHash [block hash] [includeTransactions]",
	Short: "Query the block by its hash",
	Long: `Returns the block information according to the block hash.
Arguments:
         [block hash]: hash string
[includeTransactions]: must be "true" or "false".

For example:

    getBlockByHash 0x910ea44e2a83618c7cc98456678c9984d94977625e224939b24b3c904794b5ec true

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.RangeArgs(1,2),
	Run: func(cmd *cobra.Command, args []string) {
		var bhash string
		var includeTx bool
        if(len(args) == 1) {
			bhash = args[0]
			includeTx = true
		} else {
			bhash = args[0]
			_includeTx, err := strconv.ParseBool(args[1])
			if err != nil {
				fmt.Printf("arguments error: %v\n\n", err)
				return
			}
			includeTx = _includeTx
		}
		peers,err := RPC.GetBlockByHash(context.Background(), bhash, includeTx)
		if err != nil {
			fmt.Printf("block not found: %v\n", err)
			return
		}
		fmt.Printf("Block: \n%s\n" , peers)
	},
}

var getBlockByNumberCmd = &cobra.Command{
	Use:   "getBlockByNumber [block number] [includeTransactions]",
	Short: "Query the block by its number",
	Long: `Returns the block information according to the block number.
Arguments:
       [block number]: can be input in a decimal or in hex(prefix with "0x").
[includeTransactions]: must be "true" or "false".

For example:

    getBlockByNumber 0x9 true

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.RangeArgs(1,2),
	Run: func(cmd *cobra.Command, args []string) {
		var bnumber string
		var includeTx bool
        if(len(args) == 1) {
			bnumber = args[0]
			includeTx = true
		} else {
			bnumber = args[0]
			_includeTx, err := strconv.ParseBool(args[1])
			if err != nil {
				fmt.Printf("error: %v\n\n", err)
				return
			}
			includeTx = _includeTx
		}
		block,err := RPC.GetBlockByNumber(context.Background(), bnumber, includeTx)
		if err != nil {
			fmt.Printf("block not found: %v\n", err)
			return
		}
		fmt.Printf("Block: \n%s\n" , block)
	},
}

var getBlockHashByNumberCmd = &cobra.Command{
	Use:   "getBlockHashByNumber [block number]",
	Short: "Query the block hash by its number",
	Long: `Returns the block hash according to the block number.
Arguments:
[block number]: can be input in a decimal format or in hex(prefix with "0x").

For example:

    getBlockHashByNumber 0x9

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bnumber := args[0]
		bhash,err := RPC.GetBlockHashByNumber(context.Background(), bnumber)
		if err != nil {
			fmt.Printf("block not found: %v\n", err)
			return
		}
		fmt.Printf("Block Hash: \n%s\n" , bhash)
	},
}

// ======== transaction access ========

var getTransactionByHashCmd = &cobra.Command{
	Use:   "getTransactionByHash [transaction hash]",
	Short: "Query the transaction by its hash",
	Long: `Returns the transaction according to the transaction hash.
Arguments:
[transaction hash]: hash string.

For example:
	
    getTransactionByHash 0x7536cf1286b5ce6c110cd4fea5c891467884240c9af366d678eb4191e1c31c6f
	
For more information please refer:
	
    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := args[0]
		tx,err := RPC.GetTransactionByHash(context.Background(), txHash)
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction: \n%s\n" , tx)
	},
}

var getTransactionByBlockHashAndIndexCmd = &cobra.Command{
	Use:   "getTransactionByBlockHashAndIndex [block hash] [transaction index]",
	Short: "Query the transaction by block hash and transaction index",
	Long: `Returns transaction information based on block hash and transaction index inside the block.
Arguments:
       [block hash]: block hash string.
[transaction index]: index for the transaction that must be encoded in hex format(prefix with "0x").

For example:

    getTransactionByBlockHashAndIndex 0x10bfdc1e97901ed22cc18a126d3ebb8125717c2438f61d84602f997959c631fa 0x0

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bhash := args[0]
		txIndex := args[1]
		tx,err := RPC.GetTransactionByBlockHashAndIndex(context.Background(), bhash, txIndex)
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction: \n%s\n" , tx)
	},
}

var getTransactionByBlockNumberAndIndexCmd = &cobra.Command{
	Use:   "getTransactionByBlockNumberAndIndex [block number] [transaction index]",
	Short: "Query the transaction by block number and transaction index",
	Long: `Returns transaction information based on block number and transaction index inside the block.
Arguments:
     [block number]: block number encoded in decimal format or in hex(prefix with "0x").
[transaction index]: index for the transaction that must be encoded in hex format(prefix with "0x").

For example:

    getTransactionByBlockNumberAndIndex 0x9 0x0

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bnumber := args[0]
		txIndex := args[1]
		tx,err := RPC.GetTransactionByBlockNumberAndIndex(context.Background(), bnumber, txIndex)
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction: \n%s\n" , tx)
	},
}

var getTransactionReceiptCmd = &cobra.Command{
	Use:   "getTransactionReceipt [transaction hash]",
	Short: "Query the transaction receipt by transaction hash",
	Long: `Returns transaction receipt information based on transaction hash.
Arguments:
[transaction hash]: transaction hash string.

For example:

    getTransactionReceipt 0x708b5781b62166bd86e543217be6cd954fd815fd192b9a124ee9327580df8f3f

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := args[0]
		tx,err := RPC.GetTransactionReceipt(context.Background(), txHash)
		if err != nil {
			fmt.Printf("transaction receipt not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction Receipt: \n%s\n" , tx)
	},
}

var getPendingTransactionsCmd = &cobra.Command{
	Use:   "getPendingTransactions",
	Short: "Get the pending transactions",
	Long: `Return the transactions that are pending for packaging.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tx, err := RPC.GetPendingTransactions(context.Background())
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Pending Transactions: \n%s\n" , tx)
	},
}

var getPendingTxSizeCmd = &cobra.Command{
	Use:   "getPendingTxSize",
	Short: "Get the count of pending transactions",
	Long: `Return the total count of pending transactions.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tx, err := RPC.GetPendingTxSize(context.Background())
		if err != nil {
			fmt.Printf("transactions not found: %v\n", err)
			return
		}
		fmt.Printf("Peding Transactions Count: \n%s\n" , tx)
	},
}

// ======== contracts =======

var getCodeCmd = &cobra.Command{
	Use:   "getCode [contract address]",
	Short: "Get the contract data from contract address",
	Long: `Return contract data queried according to contract address.
Arguments:
[contract address]: contract hash string.

For example:

    getCode 0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        contractAdd := args[0]
		code, err := RPC.GetCode(context.Background(), contractAdd)
		if err != nil {
			fmt.Printf("contract code not found: %v\n", err)
			return
		}
		fmt.Printf("Contract Code: \n%s\n" , code)
	},
}

var getTotalTransactionCountCmd = &cobra.Command{
	Use:   "getTotalTransactionCount",
	Short: "Get the totoal transactions and the latest block height",
	Long: `Returns the current total number of transactions and block height.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		counts, err := RPC.GetTotalTransactionCount(context.Background())
		if err != nil {
			fmt.Printf("information not found: %v\n", err)
			return
		}
		fmt.Printf("Latest Statistics on Transaction and Block Height: \n%s\n" , counts)
	},
}

var getSystemConfigByKeyCmd = &cobra.Command{
	Use:   "getSystemConfigByKey [key to query]",
	Short: "Get the system configuration through key-value, currently support [tx_count_limit],[tx_gas_limit] ",
	Long: `Returns the system configuration through key-value.
Arguments:
[key to query]: currently only support two key: "tx_count_limit" and "tx_gas_limit".

For example:

    getSystemConfigByKey tx_count_limit

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value, err := RPC.GetSystemConfigByKey(context.Background(), key)
		if err != nil {
			fmt.Printf("information not found: %v\n", err)
			return
		}
		fmt.Printf("Result: \n%s\n" , value)
	},
}

// ======= contract operation =====


func init() {
	// add common command
	// TODO: test the bash scripts
	// rootCmd.AddCommand(bashCompletionCmd, zshCompletionCmd)
	// add node command
	rootCmd.AddCommand(getClientVersionCmd, getGroupIDCmd, getBlockNumberCmd, getPbftViewCmd, getSealerListCmd)
	rootCmd.AddCommand(getObserverListCmd, getConsensusStatusCmd, getSyncStatusCmd, getPeersCmd, getGroupPeersCmd)
	rootCmd.AddCommand(getNodeIDListCmd, getGroupListCmd)
	// add block access command
	rootCmd.AddCommand(getBlockByHashCmd, getBlockByNumberCmd, getBlockHashByNumberCmd)
	// add transaction command
	rootCmd.AddCommand(getTransactionByHashCmd, getTransactionByBlockHashAndIndexCmd, getTransactionByBlockNumberAndIndexCmd)
	rootCmd.AddCommand(getTransactionReceiptCmd, getPendingTransactionsCmd, getPendingTxSizeCmd)
	// add contract command
	rootCmd.AddCommand(getCodeCmd, getTotalTransactionCountCmd, getSystemConfigByKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
