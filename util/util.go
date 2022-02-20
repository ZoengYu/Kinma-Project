package util

const (
	USD = "USD"
	TWD = "TWD"
	CNY = "CNY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency{
	case USD, TWD, CNY:
		return true
	}

	return false
}