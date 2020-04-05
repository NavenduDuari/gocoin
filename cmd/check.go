package cmd

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "To check data of crypto-currencies",
	Long: `To check data of crypto-currencies. For example:

gocoin check price		//This shows price`,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
