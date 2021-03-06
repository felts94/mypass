// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Run: func(cmd *cobra.Command, args []string) {
		label, err := cmd.Flags().GetString("label")
		pass, err := cmd.Flags().GetString("password")
		if err != nil {
			panic(err)
		}

		if label == "" {
			println("no label")
			return
		}
		pwd := password()
		if pass == "" {
			println("\nPlease enter the password for " + label)
			pass = password()
		}
		var kr KeyRing
		if _, err := os.Stat("data"); os.IsNotExist(err) {
			kr.Data[label] = pass
			encryptFile("data", kr.Bytes(), pwd)
		} else {
			d := decryptFile("data", pwd)
			kr = KeyRingFromBytes(d)
			kr.Data[label] = pass
			encryptFile("data", kr.Bytes(), pwd)
		}
		print("\n")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().String("label", "", "label for the password")
	addCmd.Flags().String("password", "", "password for the label")
}
