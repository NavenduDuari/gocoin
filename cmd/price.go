package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/NavenduDuari/gocoin/cmd/utils"

	"github.com/spf13/cobra"
)

type coinData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rank   string `json:"rank"`
	OneDay oneDay `json:"1D"`
}

type oneDay struct {
	PriceChange string `json:"price_change"`
}

var (
	baseUrl        = "https://api.nomics.com/v1/currencies/ticker?key="
	key            string
	ids            = "&ids=BTC,ETH,XRP"
	convert        = "&convert=INR"
	currencySymbol = utils.CurrencyDetails["INR"].Symbol
)

var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "To check price of crypto-currencies",
	Long: `To check price of crypto-currencies. For example:

gocoin check price					//This gives prices with default coin & currency
gocoin check price --coin=LTC,BNB --conv=INR		//This gives price of BTC in INR`,
	Run: func(cmd *cobra.Command, args []string) {
		conv, _ := cmd.Flags().GetString("conv")
		coin, _ := cmd.Flags().GetString("coin")
		suggest, _ := cmd.Flags().GetBool("suggest")
		if suggest {
			getSuggestion()
			return
		}
		getPrice(coin, conv, args)
	},
}

func init() {
	checkCmd.AddCommand(priceCmd)

	priceCmd.Flags().StringP("coin", "", "", "Let us choose coin")
	priceCmd.Flags().StringP("conv", "", "", "Let us choose currency")
	priceCmd.Flags().BoolP("suggest", "s", false, "Gives suggestion")
}

func getPrice(coin string, conv string, args []string) {
	key = getKeyValue()
	if len(coin) > 0 {
		ids = "&ids=" + coin
	}

	if len(conv) > 0 {
		convert = "&convert=" + conv
		currencySymbol = utils.CurrencyDetails[conv].Symbol
	}
	finalUrl := baseUrl + key + ids + "&interval=1d" + convert

	res, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println("Unable to get price")
		return
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	var coinDataArrObj []coinData
	json.Unmarshal(responseData, &coinDataArrObj)

	showPrice(coinDataArrObj)
}

func showPrice(coinDataArrObj []coinData) {
	for _, coin := range coinDataArrObj {
		priceChange, _ := strconv.ParseFloat(coin.OneDay.PriceChange, 64)
		price, _ := strconv.ParseFloat(coin.Price, 64)
		priceChangePercent := fmt.Sprintf("%.2f", priceChange/(priceChange+price)*100)
		utils.PrintCoinInfo(coin.Id, coin.Name)
		if priceChange < 0 {
			utils.PrintPriceDown(currencySymbol, coin.Price, priceChangePercent)
		} else {
			utils.PrintPriceUp(currencySymbol, coin.Price, priceChangePercent)
		}
		utils.PrintRank(coin.Rank)
	}
}
