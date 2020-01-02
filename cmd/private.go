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
	"log"
	"os"
	"io/ioutil"
	"bytes"
	"strings"

	"github.com/spf13/cobra"
	toml "github.com/BurntSushi/toml"
	"github.com/lacchain/lacchain-cli/model"
)

const ORION_PATH = "/home/adrianpareja/lacchain/orion/"
const ORION_CONF = "orion.conf"
const CA_PATH = "/home/adrianpareja/lacchain/orion/CAs"

// privateCmd represents the private command
var privateCmd = &cobra.Command{
	Use:   "private",
	Short: "create private channel in Orion",
	Long: `create a new private channel with other organizations using Orion.
	       This add ca certificates from other organizations`,
	Run: func(cmd *cobra.Command, args []string) {
		create(args)
	},
}

func init() {
	rootCmd.AddCommand(privateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// privateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// privateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func create(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments, is necessary 1 CA certificate\nexample: lacchain private /pathCA1/ca1.pem /pathCA2/ca2.pem")
		return
	}
	err := copyFile(ORION_PATH+ORION_CONF,ORION_PATH+"orion_back.conf")
	check(err)
	data, err := ioutil.ReadFile(ORION_PATH+"orion_back.conf")
	check(err)
	
	config := model.OrionConfig{}
	if _, err := toml.Decode(string(data), &config); err != nil {
		log.Fatal(err)
	}

	for _, ival := range args { 
		caDestination := CA_PATH+getNameFile(ival)
		copyFile(ival,caDestination)
		config.Tlsserverchain = append(config.Tlsserverchain,caDestination)
		config.Tlsclientchain = append(config.Tlsclientchain,caDestination)
	}   

	buf := new(bytes.Buffer)
	if err = toml.NewEncoder(buf).Encode(config); err != nil {
    	log.Fatal(err)
	}	

	newData := []byte(buf.String())
	
    err = ioutil.WriteFile(ORION_PATH+"/orion.conf", newData, 0644)
    check(err)

	deleteFile(ORION_PATH+"/orion_back.conf")

	fmt.Println("Addition of CAs was successful")
	fmt.Println("please, execute #>lacchain restart orion")
}

func copyFile(sourceFile string, destinationFile string)(error){
	input, err := ioutil.ReadFile(sourceFile)
    if err != nil{
		fmt.Println("Error reading", sourceFile)
		return err
	}

    err = ioutil.WriteFile(destinationFile, input, 0644)
    if err != nil {
    	fmt.Println("Error creating", destinationFile)
        return err
	}
	
	return nil
}

func deleteFile(pathFile string) {
    // delete file
    err := os.Remove(pathFile)
    if err != nil {
        return
    }
}

func getNameFile(pathFile string)(string){
	slash := strings.LastIndexByte(pathFile, '/')
	nameFile := pathFile[slash:len(pathFile)]
    return nameFile
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
