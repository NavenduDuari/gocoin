package utils

import (
	"fmt"
)

func PrintPriceDown(currencySymbol string, price string, priceChangePercent string) {
	fmt.Print("Price: ", Red, currencySymbol, price, "(", priceChangePercent, "%)", Bold, down, Reset, " ")
}

func PrintPriceUp(currencySymbol string, price string, priceChangePercent string) {
	fmt.Print("Price: ", Green, currencySymbol, price, "(+", priceChangePercent, "%)", Bold, up, Reset, " ")
}

func PrintCoinInfo(name string, id string) {
	fmt.Print("Coin: ", Cyan, name, Yellow, "(", id, ")", Reset, " ")

}

func PrintRank(rank string) {
	fmt.Print("Rank: ", Blue, Bold, rank, Reset, "\n")

}

func PrintCoinSuggestion() {
	fmt.Println(Green, Bold, "Use coin Id with --coin flag", Reset)
	for id, name := range coinDetails {
		fmt.Print("Id: ", Yellow, id, Reset)
		fmt.Println("  Name: ", Cyan, name, Reset)
	}
}

func PrintConvSuggestion() {
	fmt.Println(Green, Bold, "Use coin Id with --conv flag", Reset)
	for id, details := range CurrencyDetails {
		fmt.Print("Id: ", Yellow, id, Reset)
		fmt.Print(" Symbol: ", Red, Bold, details.Symbol, Reset)
		fmt.Println("  Name:", Cyan, details.Name, Reset)
	}
}
