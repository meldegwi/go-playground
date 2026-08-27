package money_test

import (
	"reflect"
	"testing"

	"moneyconverter/money"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		{
			name:   "34.98 USD to EUR",
			amount: money.Amount{},
			to:     money.Currency{},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
					expected := money.Amount{}
					if !reflect.DeepEqual(expected, got) {
						t.Errorf("amounts doesn't match. expected %v, got %v", expected, got)
					}
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}
