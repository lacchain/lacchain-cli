/*
Copyright © 2019 Adrian Pareja <adriancc5.5@gmail.com>

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
	"os/exec"
	"bytes"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lacchain/lacchain-cli/model"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test node is working",
	Long: `test orion's node or besu node. For example:

	lacchain test  ->  test orion and besu
	lacchain test orion  -> test only orion
	lacchain test besu  -> test only besu`,
	Run: func(cmd *cobra.Command, args []string) {
		testNode(args)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func testNode(args []string){
	if len(args) == 1 {
		if args[0] == "orion" {
			testOrion()
		}else if args[0] == "besu"{
			testBesu()
		}else{
			fmt.Println("unknown argument, only use [orion/besu]")
		}
	}else if len(args) == 0{
		testAll()
	}else{
		fmt.Println("Only support 1 argument [orion/besu]")
	}
}

func testAll(){
	testBesu()
	testOrion()
}

func testBesu(){
	if isJoinedToNetwork() {
		if isSynchroningBlocks(){
			fmt.Println("well done!. Besu is just working.")
		}else {
			fmt.Println("The node is not synchroning blocks yet")	
		}
	}else {
		fmt.Println("The node is not join to network yet")
	}
	
}

func testOrion(){

}

func isJoinedToNetwork()(bool){
	cmd := exec.Command("curl", "-X", "POST", "--data", "{\"jsonrpc\":\"2.0\",\"method\":\"net_peerCount\",\"params\":[],\"id\":1}", "34.74.56.215:4545")
	var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("%s\n", outStr)
	if err != nil{
		fmt.Printf("%s\n", errStr)
		return false
	}
	peerResult := model.PeerResponse{}
	json.Unmarshal(stdout.Bytes(), &peerResult)
	fmt.Printf("Your node is connected to %d nodes of Lacchain Network\n",hex2int(peerResult.Result))
	return true
}

func isSynchroningBlocks()(bool){
	cmd := exec.Command("tail", "-10", "/home/adrianpareja/lacchain/logs/pantheon_info_cp.log")
	var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    cmd.Run()
    outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("\n%s\n%s", outStr, errStr)
	return strings.Contains(outStr, "INFO Imported #")
}

func hex2int(hexStr string) uint32 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 32)
	return uint32(result)
}
