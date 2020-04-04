package cmd

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/NavenduDuari/gocoin/cmd/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CoinDataArr []CoinData
type CoinData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rank   string `json:"rank"`
	OneDay oneDay `json:"1D"`
}
type oneDay struct {
	PriceChange string `json:"price_change"`
}

type CoinMetaDataArr []CoinMetaData
type CoinMetaData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var (
	baseUrl        = "https://api.nomics.com/v1/currencies/ticker?key="
	key            string
	ids            = "&ids=BTC,ETH,XRP"
	convert        = "&convert=INR"
	currencySymbol = utils.CurrencyDetails["INR"].Symbol
)

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

	priceCmd.Flags().BoolP("coin", "", false, "Let us choose coin")
	priceCmd.Flags().BoolP("conv", "", false, "Let us choose currency")
	priceCmd.Flags().BoolP("suggest", "s", false, "Gives suggestion")
}

func getKeyValue() string {
	key = viper.GetString("key.nomics")
	if key == "" {
		fmt.Println("No key is set. Please set one.")
		os.Exit(1)
	}
	return key
}
func getPrice(coin bool, conv bool, args []string) {
	key = getKeyValue()
	if coin {
		ids = "&ids=" + strings.Join(args, ",")
	}

	if conv {
		convert = "&convert=" + args[0]
		currencySymbol = utils.CurrencyDetails[args[0]].Symbol
	}
	finalUrl := baseUrl + key + ids + "&interval=1d" + convert

	res, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println("Unable to get price")
		return
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	var coinDataArrObj CoinDataArr
	json.Unmarshal(responseData, &coinDataArrObj)

	//printing data
	showPrice(coinDataArrObj)
}

func showPrice(coinDataArrObj CoinDataArr) {
	for _, coin := range coinDataArrObj {
		priceChange, _ := strconv.ParseFloat(coin.OneDay.PriceChange, 64)
		price, _ := strconv.ParseFloat(coin.Price, 64)
		priceChangePercent := fmt.Sprintf("%.2f", priceChange/(priceChange+price)*100)
		fmt.Print("Coin: ", utils.Cyan, coin.Name, utils.Yellow, "(", coin.Id, ")", utils.Reset, " ")
		if priceChange < 0 {
			down := html.UnescapeString("&#" + "11015" + ";")
			fmt.Print("Price: ", utils.Red, currencySymbol, coin.Price, "(", priceChangePercent, "%)", down, utils.Reset, " ")
		} else {
			up := html.UnescapeString("&#" + "11014" + ";")
			fmt.Print("Price: ", utils.Green, currencySymbol, coin.Price, "(+", priceChangePercent, "%)", up, utils.Reset, " ")
		}
		fmt.Print("Rank: ", utils.Blue, coin.Rank, utils.Reset, "\n")
	}
}

func getSuggestion() {
	//Coin suggestion
	fmt.Println(utils.Green, "Use coin Id with --coin flag", utils.Reset)
	for id, name := range utils.CoinDetails {
		fmt.Print("Id: ", utils.Yellow, id, utils.Reset)
		fmt.Println("  Name: ", utils.Cyan, name, utils.Reset)
	}

	fmt.Println("------------------------------")
	//conversion currency suggestion
	fmt.Println(utils.Green, "Use coin Id with --conv flag", utils.Reset)
	for id, details := range utils.CurrencyDetails {
		fmt.Print("Id: ", utils.Yellow, id, utils.Reset)
		fmt.Print(" Symbol: ", utils.Red, details.Symbol, utils.Reset)
		fmt.Println("  Name:", utils.Cyan, details.Name, utils.Reset)
	}
}
