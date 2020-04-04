package utils

type details struct {
	Symbol string
	Name   string
}

var CurrencyDetails = map[string]details{
	"USD": {
		Symbol: "$",
		Name:   "US Dollar",
	}, "CAD": {
		Symbol: "CA$",
		Name:   "Canadian Dollar",
	}, "EUR": {
		Symbol: "€",
		Name:   "Euro",
	}, "AED": {
		Symbol: "د.إ.‏",
		Name:   "United Arab Emirates Dirham",
	}, "AUD": {
		Symbol: "AU$",
		Name:   "Australian Dollar",
	}, "BDT": {
		Symbol: "৳",
		Name:   "Bangladeshi Taka",
	}, "BRL": {
		Symbol: "R$",
		Name:   "Brazilian Real",
	}, "CNY": {
		Symbol: "CN¥",
		Name:   "Chinese Yuan",
	}, "CZK": {
		Symbol: "Kč",
		Name:   "Czech Republic Koruna",
	}, "GBP": {
		Symbol: "£",
		Name:   "British Pound Sterling",
	}, "ILS": {
		Symbol: "₪",
		Name:   "Israeli New Sheqel",
	}, "INR": {
		Symbol: "₹",
		Name:   "Indian Rupee",
	}, "JPY": {
		Symbol: "¥",
		Name:   "Japanese Yen",
	}, "KRW": {
		Symbol: "₩",
		Name:   "South Korean Won",
	}, "RUB": {
		Symbol: "₽.",
		Name:   "Russian Ruble",
	},
}
