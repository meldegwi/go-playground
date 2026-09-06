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
		sRate    stubRate
		validate func(t *testing.T, got money.Amount, err error)
	}{
		{
			name:   "34.98 USD to EUR",
			amount: mustParseAmount(t, "34.98", "USD"),
			to:     mustParseCurrency(t, "EUR"),
			sRate: func() stubRate {
				xr, _ := money.ParseDecimal("2.0000")
				return stubRate{
					rate: money.ExchangeRate(xr),
					err:  nil,
				}
			}(),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}
				expected := mustParseAmount(t, "69.96", "EUR")

				if !reflect.DeepEqual(expected, got) {
					t.Errorf("amounts doesn't match. expected %v, got %v", expected, got)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to, tc.sRate)
			tc.validate(t, got, err)
		})
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	curr, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency code %s", code)
	}

	return curr
}

func mustParseAmount(t *testing.T, value, code string) money.Amount {
	t.Helper()

	n, err := money.ParseDecimal(value)
	if err != nil {
		t.Fatalf("invalid number: %s", value)
	}

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("invalid currency code: %s", code)
	}

	amount, err := money.NewAmount(n, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %v and currency code %s",
			n, code)
	}

	return amount
}

// stubRate is a very simple stub for the exchangeRates.
type stubRate struct {
	rate money.ExchangeRate
	err  error
}

// FetchExchangeRate implements the interface exchangeRates with the same
// signature but fields are unused for tests purposes.
func (m stubRate) FetchExchangeRate(
	_, _ money.Currency,
) (money.ExchangeRate, error) {
	return m.rate, m.err
}
