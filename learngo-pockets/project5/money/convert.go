package money

func Convert(amount Amount, to Currency) (Amount, error) {
	// implement me
	return Amount{}, nil
}

type ExchangeRate Decimal

func applyExchangeRate(a Amount, target Currency, xr ExchangeRate) Amount {
	converted, err := multiply(a.quantity, xr)
	if err != nil {
		return Amount{}
	}

	switch {
	case converted.percision > target.percision:
		converted.subunits = converted.subunits / pow10(converted.percision-target.percision)
	case converted.percision < target.percision:
		converted.subunits = converted.subunits * pow10(target.percision-converted.percision)
	}

	converted.percision = target.percision

	return Amount{
		quantity: converted,
		currency: target,
	}
}

func multiply(d Decimal, xr ExchangeRate) (Decimal, error) {
	dec := Decimal{
		subunits:  d.subunits * xr.subunits,
		percision: d.percision + xr.percision,
	}

	dec.Simplify()

	return dec, nil
}
