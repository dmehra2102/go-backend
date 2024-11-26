package util

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUR"
	GBP = "GBP"
	AUD = "AUD"
	CAD = "CAD"
	JPY = "JPY"
	CNY = "CNY"
	NZD = "NZD"
	SGD = "SGD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, GBP, AUD, CAD, JPY, CNY, NZD, SGD:
		return true
	default:
		return false
	}
}
