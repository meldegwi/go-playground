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
				quantity: Decimal{subunits: 123, percision: 2},
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 4},
			xRate:  ExchangeRate{subunits: 1, percision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 12300, percision: 4},
				currency: Currency{code: "TRG", percision: 4},
			},
		},
		{
			name: "same currency and same precision",
			amount: Amount{
				quantity: Decimal{subunits: 152, percision: 2},
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 2},
			xRate:  ExchangeRate{subunits: 1, percision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 152, percision: 2},
				currency: Currency{code: "TRG", percision: 2},
			},
		},
		{
			name: "target precision removes decimal part",
			amount: Amount{
				quantity: Decimal{subunits: 250, percision: 2},
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 0},
			xRate:  ExchangeRate{subunits: 1, percision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 2, percision: 0},
				currency: Currency{code: "TRG", percision: 0},
			},
		},
		{
			name: "increase target precision",
			amount: Amount{
				quantity: Decimal{subunits: 250, percision: 2},
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 4},
			xRate:  ExchangeRate{subunits: 1, percision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 25000, percision: 4},
				currency: Currency{code: "TRG", percision: 4},
			},
		},
		{
			name: "real life exchange rate",
			amount: Amount{
				quantity: Decimal{subunits: 4250, percision: 2}, // 42.50
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 2},
			xRate:  ExchangeRate{subunits: 11101, percision: 4}, // 1.1101
			expected: Amount{
				quantity: Decimal{subunits: 4717, percision: 2}, // 47.17
				currency: Currency{code: "TRG", percision: 2},
			},
		},
		{
			name: "keep precision of rate",
			amount: Amount{
				quantity: Decimal{subunits: 1000, percision: 0}, // 1000
				currency: Currency{code: "TST", percision: 0},
			},
			target: Currency{code: "TRG", percision: 2},
			xRate:  ExchangeRate{subunits: 11101, percision: 4}, // 1.1101
			expected: Amount{
				quantity: Decimal{subunits: 111010, percision: 2},
				currency: Currency{code: "TRG", percision: 2},
			},
		},
		{
			name: "large numbers",
			amount: Amount{
				quantity: Decimal{subunits: 10000000000122, percision: 4}, // 1,000,000,000.0122
				currency: Currency{code: "TST", percision: 4},
			},
			target: Currency{code: "TRG", percision: 2},
			xRate:  ExchangeRate{subunits: 12387, percision: 2}, // 123.87
			expected: Amount{
				quantity: Decimal{subunits: 12387000000151, percision: 2},
				currency: Currency{code: "TRG", percision: 2},
			},
		},
		{
			name: "very small exchange rate",
			amount: Amount{
				quantity: Decimal{subunits: 1000000, percision: 2}, // 10,000
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 2},
			xRate:  ExchangeRate{subunits: 505935, percision: 10}, // 0.000505935
			expected: Amount{
				quantity: Decimal{subunits: 50, percision: 2},
				currency: Currency{code: "TRG", percision: 2},
			},
		},
		{
			name: "adding precision when input has none",
			amount: Amount{
				quantity: Decimal{subunits: 133, percision: 0}, // 133
				currency: Currency{code: "TST", percision: 0},
			},
			target: Currency{code: "TRG", percision: 5},
			xRate:  ExchangeRate{subunits: 1, percision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 13300000, percision: 5},
				currency: Currency{code: "TRG", percision: 5},
			},
		},
		{
			name: "increasing output precision",
			amount: Amount{
				quantity: Decimal{subunits: 133, percision: 2}, // 1.33
				currency: Currency{code: "TST", percision: 2},
			},
			target: Currency{code: "TRG", percision: 5},
			xRate:  ExchangeRate{subunits: 1, percision: 0},
			expected: Amount{
				quantity: Decimal{subunits: 133000, percision: 5},
				currency: Currency{code: "TRG", percision: 5},
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
