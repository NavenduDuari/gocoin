package cmd

import (
	"fmt"

	"github.com/NavenduDuari/gocoin/cmd/utils"
)

func getSuggestion() {
	utils.PrintCoinSuggestion()
	fmt.Println("------------------------------")
	utils.PrintConvSuggestion()

}
