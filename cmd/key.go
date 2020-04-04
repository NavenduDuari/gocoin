/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/viper"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manages api key for nomics.com",
	Long: `Manages api key for nomics.com. For example:

gocoin key				//This shows your api key 
gocoin key --set yourApikey		//This sets your api key
gocoin key --del			//This deletes your api key`,
	Run: func(cmd *cobra.Command, args []string) {
		set, _ := cmd.Flags().GetBool("set")
		del, _ := cmd.Flags().GetBool("del")
		if set {
			setKey(args)
			return
		} else if del {
			delKey(args)
			return
		} else {
			getKey(args)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)

	keyCmd.Flags().BoolP("set", "s", false, "Set api key")
	keyCmd.Flags().BoolP("get", "g", false, "Get api key")
	keyCmd.Flags().BoolP("del", "d", false, "Delete api key")
}

func setKey(args []string) {
	if len(args) < 1 {
		fmt.Println("Too few arguments")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	viper.Set("key.nomics", args[0])
	viper.WriteConfig()
}

func getKey(args []string) {
	if len(args) > 0 {
		fmt.Println("Too many arguments")
		return
	}

	key := viper.GetString("key.nomics")
	if key == "" {
		fmt.Println("No key is set. Please set one.")
		return
	}
	fmt.Println(key)
}

func delKey(args []string) { //TODO delete key
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	viper.Set("key.nomics", "")
	viper.WriteConfig()
}
