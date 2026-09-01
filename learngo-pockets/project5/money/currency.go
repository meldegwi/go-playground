package money

import "unicode"

const (
	ErrInvalidCurrencyCode = Error("invalid currency code")
)

// Currency defines code of the currency.
type Currency struct {
	code      string
	percision byte
}

func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	for _, r := range code {
		if !unicode.IsLetter(r) || !unicode.IsUpper(r) {
			return Currency{}, ErrInvalidCurrencyCode
		}
	}

	var perc byte
	switch code {
	case "IRR":
		perc = 0
	case "CNY", "VND":
		perc = 1
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		perc = 3
	default:
		perc = 2
	}

	return Currency{
		code:      code,
		percision: perc,
	}, nil
}
