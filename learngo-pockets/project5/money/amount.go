package money

import "fmt"

const (
	ErrIncompatibleDeciCurr = Error("Decimal percision and Currency percision are not compatible")
)

type Amount struct {
	quantity Decimal
	currency Currency
}

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.percision > currency.percision {
		return Amount{}, fmt.Errorf("%w: too percise", ErrIncompatibleDeciCurr)
	}

	quantity.percision = currency.percision

	return Amount{
		quantity: quantity,
		currency: currency,
	}, nil
}

func (a *Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.percision > a.currency.percision:
		return fmt.Errorf("%w: too percise", ErrIncompatibleDeciCurr)
	}

	return nil
}
