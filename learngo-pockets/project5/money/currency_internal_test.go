package money

import (
	"testing"
)

func TestParseCurrency(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		curr    Currency
		wantErr error
	}{
		{
			name: "valid code",
			code: "EGP",
			curr: Currency{
				code:      "EGP",
				precision: 2,
			},
			wantErr: nil,
		},
		{
			name:    "contains numbers",
			code:    "LO1",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "empty code",
			code:    "",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "too short",
			code:    "EG",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "too long",
			code:    "EGPP",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "lowercase",
			code:    "egp",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "mixed case",
			code:    "EgP",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "contains special character",
			code:    "E$P",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "contains whitespace",
			code:    "E P",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "leading whitespace",
			code:    " EGP",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name:    "trailing whitespace",
			code:    "EGP ",
			curr:    Currency{},
			wantErr: ErrInvalidCurrencyCode,
		},
		{
			name: "another valid code",
			code: "USD",
			curr: Currency{
				code:      "USD",
				precision: 2,
			},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			curr, err := ParseCurrency(tc.code)

			if curr != tc.curr {
				t.Errorf("currencies mismatch: expected %v, got %v", tc.curr, curr)
			}

			if err != tc.wantErr {
				t.Errorf("errors mismatch: expected %v, got %v", tc.wantErr, err)
			}
		})
	}
}
