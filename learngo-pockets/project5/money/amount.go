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
	switch {
	case quantity.percision > currency.percision:
		return Amount{}, fmt.Errorf("%w: too percise",
			ErrIncompatibleDeciCurr)

	case quantity.percision < currency.percision:
		quantity.subunits = quantity.subunits *
			pow10(currency.percision-quantity.percision)
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

func (a *Amount) String() string {
	return a.quantity.String() + " " + a.currency.String()
}
