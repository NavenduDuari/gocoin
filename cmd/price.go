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
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/NavenduDuari/gocoin/cmd/utils"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

type Response []CoinData
type CoinData struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Rank   string `json:"rank"`
	OneDay oneDay `json:"1D"`
}
type oneDay struct {
	PriceChange string `json:"price_change"`
}

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "To check price of crypto-currencies",
	Long: `To check price of crypto-currencies. For example:

gocoin check price					//This gives prices with default coin & currency
gocoin check price --coin BTC --conv INR		//This gives price of BTC in INR`,
	Run: func(cmd *cobra.Command, args []string) {
		conv, _ := cmd.Flags().GetBool("conv")
		coin, _ := cmd.Flags().GetBool("coin")

		getPrice(coin, conv, args)
	},
}

func init() {
	checkCmd.AddCommand(priceCmd)

	priceCmd.Flags().BoolP("coin", "", false, "let us choose coin")
	priceCmd.Flags().BoolP("conv", "", false, "let us choose currency")
}

func getPrice(coin bool, conv bool, args []string) {
	baseUrl := "https://api.nomics.com/v1/currencies/ticker?key="
	key := viper.GetString("key.nomics")
	//default id
	ids := "&ids=BTC,ETH,XRP"
	//default currency
	convert := "&convert=EUR"
	var currencySymbol = utils.CurrencySymbol["EUR"]
	if key == "" {
		fmt.Println("No key is set. Please set one.")
		return
	}

	if coin {
		ids = "&ids=" + strings.Join(args, ",")
	}

	if conv {
		convert = "&convert=" + args[0]
		currencySymbol = utils.CurrencySymbol[args[0]]
	}
	finalUrl := baseUrl + key + ids + convert

	res, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println("Unable to get price")
		return
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	//printing data
	for _, coin := range responseObject {
		priceChange, _ := strconv.Atoi(coin.OneDay.PriceChange)
		fmt.Print("Coin: ", utils.Yellow+coin.Id+utils.Reset, " ")
		if priceChange < 0 {
			down := html.UnescapeString("&#" + "11015" + ";")
			fmt.Print("Price: ", utils.Red, currencySymbol, coin.Price, down, utils.Reset, " ")
		} else {
			up := html.UnescapeString("&#" + "11014" + ";")
			fmt.Print("Price: ", utils.Green, currencySymbol, coin.Price, up, utils.Reset, " ")
		}
		fmt.Print("Rank: ", utils.Blue, coin.Rank, utils.Reset, "\n")
	}
}
