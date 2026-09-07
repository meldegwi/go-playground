package money

import (
	"reflect"
	"testing"
)

func TestApplyExchangeRate(t *testing.T) {
	tt := []struct {
		name     string
		amount   Amount
		target   Currency
		xRate    ExchangeRate
		expected Amount
	}{
		{
			name: "xr is the same",
			amount: Amount{
				quantity: Decimal{subunits: 123, precision: 2},
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 4},
			xRate:  ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 12300, precision: 4},
				currency: Currency{code: "TRG", precision: 4},
			},
		},
		{
			name: "same currency and same precision",
			amount: Amount{
				quantity: Decimal{subunits: 152, precision: 2},
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 2},
			xRate:  ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 152, precision: 2},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		{
			name: "target precision removes decimal part",
			amount: Amount{
				quantity: Decimal{subunits: 250, precision: 2},
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 0},
			xRate:  ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 2, precision: 0},
				currency: Currency{code: "TRG", precision: 0},
			},
		},
		{
			name: "increase target precision",
			amount: Amount{
				quantity: Decimal{subunits: 250, precision: 2},
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 4},
			xRate:  ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 25000, precision: 4},
				currency: Currency{code: "TRG", precision: 4},
			},
		},
		{
			name: "real life exchange rate",
			amount: Amount{
				quantity: Decimal{subunits: 4250, precision: 2}, // 42.50
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 2},
			xRate:  ExchangeRate{subunits: 11101, precision: 4}, // 1.1101
			expected: Amount{
				quantity: Decimal{subunits: 4717, precision: 2}, // 47.17
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		{
			name: "keep precision of rate",
			amount: Amount{
				quantity: Decimal{subunits: 1000, precision: 0}, // 1000
				currency: Currency{code: "TST", precision: 0},
			},
			target: Currency{code: "TRG", precision: 2},
			xRate:  ExchangeRate{subunits: 11101, precision: 4}, // 1.1101
			expected: Amount{
				quantity: Decimal{subunits: 111010, precision: 2},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		{
			name: "large numbers",
			amount: Amount{
				quantity: Decimal{subunits: 10000000000122, precision: 4}, // 1,000,000,000.0122
				currency: Currency{code: "TST", precision: 4},
			},
			target: Currency{code: "TRG", precision: 2},
			xRate:  ExchangeRate{subunits: 12387, precision: 2}, // 123.87
			expected: Amount{
				quantity: Decimal{subunits: 12387000000151, precision: 2},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		{
			name: "very small exchange rate",
			amount: Amount{
				quantity: Decimal{subunits: 1000000, precision: 2}, // 10,000
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 2},
			xRate:  ExchangeRate{subunits: 505935, precision: 10}, // 0.000505935
			expected: Amount{
				quantity: Decimal{subunits: 50, precision: 2},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		{
			name: "adding precision when input has none",
			amount: Amount{
				quantity: Decimal{subunits: 133, precision: 0}, // 133
				currency: Currency{code: "TST", precision: 0},
			},
			target: Currency{code: "TRG", precision: 5},
			xRate:  ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 13300000, precision: 5},
				currency: Currency{code: "TRG", precision: 5},
			},
		},
		{
			name: "increasing output precision",
			amount: Amount{
				quantity: Decimal{subunits: 133, precision: 2}, // 1.33
				currency: Currency{code: "TST", precision: 2},
			},
			target: Currency{code: "TRG", precision: 5},
			xRate:  ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 133000, precision: 5},
				currency: Currency{code: "TRG", precision: 5},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := applyExchangeRate(tc.amount, tc.target, tc.xRate)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("conversion failed: amounts mismatch: got %v, expected %v", got, tc.expected)
			}
		})
	}
}
