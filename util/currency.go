package util

const (
	USD = "USD"
	EUR = "EUR"
	HRK = "HRK"
	KM  = "KM"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, HRK, KM:
		return true
	}
	return false
}
