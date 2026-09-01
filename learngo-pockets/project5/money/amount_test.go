package money_test

import (
	"errors"
	"testing"

	"moneyconverter/money"
)

func TestNewAmount(t *testing.T) {
	tests := []struct {
		name    string
		quanty  string
		curr    string
		wantErr error
	}{
		{
			name:    "valid case",
			quanty:  "12.34",
			curr:    "EGP",
			wantErr: nil,
		},
		{
			name:    "percision mismatch",
			quanty:  "12.34567",
			curr:    "USD",
			wantErr: money.ErrIncompatibleDeciCurr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			quanty, _ := money.ParseDecimal(tc.quanty)
			curr, _ := money.ParseCurrency(tc.curr)

			_, err := money.NewAmount(quanty, curr)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("errors mismatch: expected %v, got %v", tc.wantErr, err)
			}
		})
	}
}
