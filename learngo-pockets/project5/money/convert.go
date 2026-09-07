package money

import "fmt"

func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get rate: %w", err)
	}

	convertedValue := applyExchangeRate(amount, to, r)
	if err := convertedValue.validate(); err != nil {
		return Amount{}, nil
	}

	return convertedValue, nil
}

type ExchangeRate Decimal

func applyExchangeRate(a Amount, target Currency, xr ExchangeRate) Amount {
	converted, err := multiply(a.quantity, xr)
	if err != nil {
		return Amount{}
	}

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		quantity: converted,
		currency: target,
	}
}

func multiply(d Decimal, xr ExchangeRate) (Decimal, error) {
	dec := Decimal{
		subunits:  d.subunits * xr.subunits,
		precision: d.precision + xr.precision,
	}

	// dec.Simplify()

	return dec, nil
}

func (xr ExchangeRate) Divide(xRate ExchangeRate) ExchangeRate {
	dec1 := Decimal(xr)
	dec2 := Decimal(xRate)
	res := dec1.Divide(dec2)
	return ExchangeRate(res)
}
