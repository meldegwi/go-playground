package money

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ErrMissingDecimalPoint  = Error("unable to find the decimal point")
	ErrTooManyDecimalPoints = Error("too many decimal points")
	ErrInvalidDecimal       = Error("unable to convert the decimal")
	ErrTooLarge             = Error("quantity over 10^12 is too large")
)

const maxDecimal = 1e12

// Decimal is responsible for representing a floating point number with fixed percision.
// example:
// 1.52 = 152 * 10^(-2) whill be stored as {152, 2}
type Decimal struct {
	subunits  int64
	percision byte
}

// ParseDecimal converts a string to its decimal representation.
// It assumes there is only one decimal seperator,
// and that seperator is '.'(full stop character).
func ParseDecimal(decimal string) (Decimal, error) {
	intPart, fracPart, ok := strings.Cut(decimal, ".")
	if !ok {
		return Decimal{}, ErrMissingDecimalPoint
	}

	if _, _, ok := strings.Cut(fracPart, "."); ok {
		return Decimal{}, ErrTooManyDecimalPoints
	}

	if _, _, ok := strings.Cut(intPart, "."); ok {
		return Decimal{}, ErrTooManyDecimalPoints
	}

	var sb strings.Builder
	sb.WriteString(intPart)
	sb.WriteString(fracPart)

	su, err := strconv.ParseInt(sb.String(), 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err)
	}
	if su > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	perc := byte(len(fracPart))
	return Decimal{subunits: su, percision: perc}, nil
}

func (d *Decimal) Simplify() {
	for d.subunits%10 == 0 {
		d.percision--
		d.subunits /= 10
	}
}
