package money

import (
	"errors"
	"testing"
)

func TestNewAmount(t *testing.T) {
	tests := []struct {
		name     string
		quantity Decimal
		currency Currency
		want     Amount
		wantErr  error
	}{
		{
			name:     "valid case",
			quantity: Decimal{1234, 2},
			currency: Currency{"EGP", 2},
			want:     Amount{quantity: Decimal{1234, 2}, currency: Currency{"EGP", 2}},
		},
		{
			name:     "quantity more precise than currency",
			quantity: Decimal{1234, 4},
			currency: Currency{"USD", 2},
			want:     Amount{},
			wantErr:  ErrIncompatibleDeciCurr,
		},
		{
			name:     "quantity precision matches currency",
			quantity: Decimal{123456, 4},
			currency: Currency{"USD", 4},
			want:     Amount{quantity: Decimal{123456, 4}, currency: Currency{"USD", 4}},
		},
		{
			name:     "quantity less precise than currency",
			quantity: Decimal{123456, 2},
			currency: Currency{"USD", 3},
			want:     Amount{quantity: Decimal{123456, 3}, currency: Currency{"USD", 3}},
		},
		{
			name:     "zero precision currency",
			quantity: Decimal{1234, 0},
			currency: Currency{"JPY", 0},
			want:     Amount{quantity: Decimal{1234, 0}, currency: Currency{"JPY", 0}},
		},
		{
			name:     "zero precision quantity with precise currency",
			quantity: Decimal{1234, 0},
			currency: Currency{"KWD", 3},
			want:     Amount{quantity: Decimal{1234, 3}, currency: Currency{"KWD", 3}},
		},
		{
			name:     "one precision with two precision currency",
			quantity: Decimal{12345, 1},
			currency: Currency{"EGP", 2},
			want:     Amount{quantity: Decimal{12345, 2}, currency: Currency{"EGP", 2}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)

			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("errors mismatch: expected %v, got %v", tc.wantErr, err)
			}

			if got != tc.want {
				t.Errorf("amount mismatch: expected %+v, got %+v", tc.want, got)
			}
		})
	}
}
