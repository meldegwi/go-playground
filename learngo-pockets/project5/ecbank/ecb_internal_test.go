package ecbank

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"moneyconverter/money"
)

func TestEuroCenteralBank_FetchExchangeRate_Sucess(t *testing.T) {
	tServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mockXML := `<?xml version="1.0" encoding="UTF-8"?>
<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-05-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
	<Cube>
		<Cube time="2026-09-06">
			<Cube currency="USD" rate="1.0850"/>
			<Cube currency="RON" rate="4.9750"/>
		</Cube>
	</Cube>
</gesmes:Envelope>`
			fmt.Fprintln(w, mockXML)
		}))
	defer tServer.Close()

	ecb := Client{
		url: tServer.URL,
	}

	_, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))
	if err != nil {
		t.Errorf("expected no error got %s", err.Error())
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
