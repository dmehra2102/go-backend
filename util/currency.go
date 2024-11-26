package util

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR:
		return true
	default:
		return false
	}
}
