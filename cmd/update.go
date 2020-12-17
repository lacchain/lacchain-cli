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
package cmd

import (
	"fmt"
    "context"
    "encoding/hex"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"

    bl "github.com/lacchain/lacchain-cli/blockchain"
    rules "github.com/lacchain/lacchain-cli/blockchain/contracts"
    "github.com/lacchain/lacchain-cli/errors"
    log "github.com/lacchain/lacchain-cli/util"
    "github.com/lacchain/lacchain-cli/model"

	"github.com/spf13/cobra"
)

const TOPIC_ADDED = "0x983a527ad2402ad85d7f70bcae14ec1567e0b0d2e06a6f72ffbcabfe3e8863ea"
const TOPIC_REMOVED ="0xf05dee0659735cf956ff02ae9f4bd9f1c41bb30ea20d7a1a3869a42c7254ca45"

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("update called")
        updateOption, _ := cmd.Flags().GetString("config")

        if updateOption=="bootnode"{
            uploadBootnodes();
        }else{
            fmt.Println("use config flag")
        }
	},
}

func init() {
    rootCmd.AddCommand(updateCmd)
    updateCmd.Flags().StringP("config", "c", "bootnode", "Update initial bootnodes in config.toml")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func uploadBootnodes(){
	client := new(bl.Client)
	err := client.Connect("http://34.74.56.215:4545")
	if err != nil {
		handleError(err)
	}
    defer client.Close()
    
    nodeIngressAddress := common.HexToAddress("0x0000000000000000000000000000000000009999")

    name := [32]byte{}

    decodeByte, err := hex.DecodeString("72756c6573000000000000000000000000000000000000000000000000000000")

    copy(name[:], decodeByte)

    nodeRules, err := client.GetNodeIngressContract(nodeIngressAddress,name)

    log.GeneralLogger.Println("nodeRules", nodeRules)

    contractAddress := common.HexToAddress(nodeRules)
    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(18000000),
        ToBlock:   nil,
        Addresses: []common.Address{
            contractAddress,
		},
	}

    logs, err := client.GetEthclient().FilterLogs(context.Background(), query)
    if err != nil {
        log.GeneralLogger.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(rules.RulesABI)))
    if err != nil {
        log.GeneralLogger.Fatal(err)
	}
    
    addBootList := []model.Event{}
    removeBootList := []model.Event{}
    
	for _, vLog := range logs {
        fmt.Println("BlockHash:",vLog.BlockHash.Hex()) 
        fmt.Println("BlockNumber:",vLog.BlockNumber) 
        fmt.Println("TxHash:",vLog.TxHash.Hex())

        event := model.Event{}
        err := contractAbi.Unpack(&event, "NodeAdded", vLog.Data)
        if err != nil {
            log.GeneralLogger.Fatal(err)
        }

        if TOPIC_ADDED == vLog.Topics[0].Hex(){
            fmt.Println("nodeAdded:",event.NodeAdded)
            fmt.Println(fmt.Sprintf("enodeHigh:0x%s",hex.EncodeToString(event.EnodeHigh[:])))
            fmt.Println(fmt.Sprintf("enodeLow:0x%s",hex.EncodeToString(event.EnodeLow[:])))
            fmt.Println(fmt.Sprintf("enodeIp:0x%s",hex.EncodeToString(event.EnodeIp[:])))
            fmt.Println("enodePort:",event.EnodePort)

            nodeType,err := client.GetNodeInformation(contractAddress,event.EnodeHigh,event.EnodeLow,event.EnodeIp,event.EnodePort,nil)
            if err != nil {
                log.GeneralLogger.Println(err)
            }
            if nodeType == 1{
                fmt.Println("Agregando el bootnode")
                addBootList = append(addBootList,event)
            }else{
                fmt.Println("No hacer nada")
            }
        } else if TOPIC_REMOVED == vLog.Topics[0].Hex(){
            fmt.Println("nodeRemoved:",event.NodeAdded)
            fmt.Println(fmt.Sprintf("enodeHigh:0x%s",hex.EncodeToString(event.EnodeHigh[:])))
            fmt.Println(fmt.Sprintf("enodeLow:0x%s",hex.EncodeToString(event.EnodeLow[:])))
            fmt.Println(fmt.Sprintf("enodeIp:0x%s",hex.EncodeToString(event.EnodeIp[:])))
            fmt.Println("enodePort:",event.EnodePort)

            oldBlockNumber := new(big.Int).SetUint64(vLog.BlockNumber-1) 
            nodeType,err := client.GetNodeInformation(contractAddress,event.EnodeHigh,event.EnodeLow,event.EnodeIp,event.EnodePort,oldBlockNumber)
            if err != nil {
                log.GeneralLogger.Println(err)
            }
            if nodeType == 1{
                fmt.Println("Eliminando el bootnode")
                removeBootList = append(removeBootList,event)
            }else{
                fmt.Println("No es bootnode")
            }
        }

        log.ReplaceConfigFile(addBootList,removeBootList)
	}
    
	eventSignature := []byte("NodeAdded(bool,bytes32,bytes32,bytes32,bytes16,uint)")
    hash := crypto.Keccak256Hash(eventSignature)
    fmt.Println("eventSignature:",hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}

func handleError(err error)(int){
	errorType := errors.GetType(err)
   	switch errorType {
    	case errors.FailedConnection: 
			log.GeneralLogger.Fatal(err.Error())
		case errors.FailedKeystore:
			log.GeneralLogger.Fatal(err.Error())	
		case errors.FailedConfigTransaction:
			log.GeneralLogger.Println(err.Error())
			return 100  
		default: 
			log.GeneralLogger.Println("otro error:",err)
	   }
	  return 100
}
