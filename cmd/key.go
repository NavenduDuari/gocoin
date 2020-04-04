package cmd

import (
	"fmt"
	"os"

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
			fmt.Println(getKey(args))
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

func getKey(args []string) string {
	if len(args) > 0 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	key := viper.GetString("key.nomics")
	if key == "" {
		fmt.Println("No key is set. Please set one.")
		os.Exit(1)
	}
	return key
}

func delKey(args []string) { //TODO delete key
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	viper.Set("key.nomics", "")
	viper.WriteConfig()
}
