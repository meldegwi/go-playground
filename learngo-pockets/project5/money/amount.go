package money

import "fmt"

const (
	ErrIncompatibleDeciCurr = Error("Decimal precision and Currency precision are not compatible")
)

type Amount struct {
	quantity Decimal
	currency Currency
}

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	switch {
	case quantity.precision > currency.precision:
		return Amount{}, fmt.Errorf("%w: too precise",
			ErrIncompatibleDeciCurr)

	case quantity.precision < currency.precision:
		quantity.subunits = quantity.subunits *
			pow10(currency.precision-quantity.precision)
	}

	quantity.precision = currency.precision

	return Amount{
		quantity: quantity,
		currency: currency,
	}, nil
}

func (a *Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return fmt.Errorf("%w: too precise", ErrIncompatibleDeciCurr)
	}

	return nil
}

func (a *Amount) String() string {
	return a.quantity.String() + " " + a.currency.String()
}
