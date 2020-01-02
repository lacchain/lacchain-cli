/*
Copyright © 2020 Adrian Pareja <adriancc5.5@gmail.com>

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

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart both orion's node and besu by default",
	Long: `restart orion's node or besu node. For example:

lacchain restart  ->  restart orion and besu
lacchain restart orion  -> restart only orion
lacchain restart besu  -> restart only besu`,  
	Run: func(cmd *cobra.Command, args []string) {
		restart(args)
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func restart(args []string){
	if len(args) == 1 {
		if args[0] == "orion" {
			restartOrion()
		}else if args[0] == "besu"{
			restartBesu()
		}else{
			fmt.Println("unknown argument, only use [orion/besu]")
		}
	}else if len(args) == 0{
		restartAll()
	}else{
		fmt.Println("Only support 1 argument [orion/besu]")
	}
}

func restartAll(){
	executeService("pantheon","stop")
	executeService("orion","stop")
	executeService("orion","start")
	executeService("pantheon","start")
}

func restartOrion(){
	executeService("orion","restart")
}

func restartBesu(){
	executeService("pantheon","restart")
}

func executeService(service string, option string){
	cmd := exec.Command("service", service, option)
	var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    cmd.Run()
    outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
    fmt.Printf("%s\n%s", outStr, errStr)
}

