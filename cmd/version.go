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

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Lacchain's cli version",
	Long: `Version of Lacchain's cli node Orion and Besu`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LACChain version 0.0.9")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
