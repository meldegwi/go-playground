package money

import (
	"fmt"
	"math"
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

// Decimal is responsible for representing a floating point number with fixed precision.
// example:
// 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Decimal struct {
	subunits  int64
	precision byte
}

// ParseDecimal converts a string to its decimal representation.
// It assumes there is only one decimal separator,
// and that separator is '.'(full stop character).
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
	return Decimal{subunits: su, precision: perc}, nil
}

// Simplify is removing any tailing zeros from your money.
// However, you should use it carefully as it remove entire amounts where they are dividable by 10
// e.g. 10.00 USD will get simplified to 00.00
func (d *Decimal) Simplify() {
	for d.subunits%10 == 0 {
		d.precision--
		d.subunits /= 10
	}
}

func (d *Decimal) String() string {
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.subunits)
	}

	centsPerUnits := pow10(d.precision)
	frac := d.subunits % centsPerUnits
	integer := d.subunits / centsPerUnits

	decimalFormat := "%d.%0" +
		strconv.Itoa(int(d.precision)) + "d"

	return fmt.Sprintf(decimalFormat, integer, frac)
}

func (d *Decimal) Divide(dec Decimal) Decimal {
	res := float64(d.subunits) / float64(dec.subunits)
	var pow byte
	if d.precision >= dec.precision {
		pow = byte(d.precision)
	} else {
		pow = byte(dec.precision)
	}

	subuns := int64(res * float64(pow10(pow)))
	return Decimal{subunits: subuns, precision: d.precision}
}

func pow10(pow byte) int64 {
	switch pow {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(pow)))
	}
}
