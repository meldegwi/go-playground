package money

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tests := []struct {
		name     string
		decimal  string
		expected Decimal
		wantErr  error
	}{
		{
			name:    "valid decimal",
			decimal: "1.23",
			expected: Decimal{
				subunits:  123,
				percision: 2,
			},
			wantErr: nil,
		},
		{
			name:    "2 decimal digits",
			decimal: "1.52",
			expected: Decimal{
				subunits:  152,
				percision: 2,
			},
			wantErr: nil,
		},
		{
			name:     "no decimal digits",
			decimal:  "123",
			expected: Decimal{},
			wantErr:  ErrMissingDecimalPoint,
		},
		{
			name:    "suffix 0 as decimal digits",
			decimal: "1.50",
			expected: Decimal{
				subunits:  150,
				percision: 2,
			},
			wantErr: nil,
		},
		{
			name:    "prefix 0 as decimal digits",
			decimal: "0.52",
			expected: Decimal{
				subunits:  52,
				percision: 2,
			},
			wantErr: nil,
		},
		{
			name:    "multiple of 10",
			decimal: "10.00",
			expected: Decimal{
				subunits:  1000,
				percision: 2,
			},
			wantErr: nil,
		},
		{
			name:     "no decimal point",
			decimal:  "123",
			expected: Decimal{},
			wantErr:  ErrMissingDecimalPoint,
		},
		{
			name:     "too many decimal points",
			decimal:  "1.2.3",
			expected: Decimal{},
			wantErr:  ErrTooManyDecimalPoints,
		},
		{
			name:     "invalid decimal",
			decimal:  "1,.23",
			expected: Decimal{},
			wantErr:  ErrInvalidDecimal,
		},
		{
			name:     "invalid decimal part",
			decimal:  "1.ab",
			expected: Decimal{},
			wantErr:  ErrInvalidDecimal,
		},
		{
			name:     "not a number",
			decimal:  "NaN",
			expected: Decimal{},
			wantErr:  ErrMissingDecimalPoint,
		},
		{
			name:     "empty string",
			decimal:  "",
			expected: Decimal{},
			wantErr:  ErrMissingDecimalPoint,
		},
		{
			name:     "too large decimal",
			decimal:  "1.1234567891112",
			expected: Decimal{},
			wantErr:  ErrTooLarge,
		},
		{
			name:     "too large integer",
			decimal:  "1234567890123.0",
			expected: Decimal{},
			wantErr:  ErrTooLarge,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)

			if got != tc.expected {
				t.Errorf("decimals doesn't match. expected %v, got %v", tc.expected, got)
			}
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("error doesn't match. expected %v, got %v", tc.wantErr, err)
			}
		})
	}
}
