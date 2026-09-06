package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"moneyconverter/money"
)

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string  `xml:"currency, attr"`
	Rate     float64 `xml:"rate, attr"`
}

const baseCurrencyCode = "EUR"

func (e envelope) exchangeRates() map[string]money.ExchangeRate {
	rates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	for _, c := range e.Rates {
		dec, _ := money.ParseDecimal(strconv.FormatFloat(c.Rate, 'f', -1, 64))
		rates[c.Currency] = money.ExchangeRate(dec)
	}

	euroDec, _ := money.ParseDecimal("1.0000")
	rates[baseCurrencyCode] = money.ExchangeRate(euroDec)

	return rates
}

func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		euroDec, _ := money.ParseDecimal("1.0000")
		return money.ExchangeRate(euroDec), nil
	}

	rates := e.exchangeRates()

	sourceFactor, ok := rates[source]
	if !ok {
		return money.ExchangeRate{}, fmt.Errorf("failed to find the source currency", source)
	}

	targetFactor, ok := rates[target]
	if !ok {
		return money.ExchangeRate{}, fmt.Errorf("failed to find the target currency", source)
	}

	res := targetFactor.Divide(sourceFactor)

	return res, nil
}

func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {
	decoder := xml.NewDecoder(respBody)

	var ecbMessage envelope
	err := decoder.Decode(&ecbMessage)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}

	rate, err := ecbMessage.exchangeRate(source, target)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrExchangeRateNotFound, err)
	}

	return rate, nil
}
