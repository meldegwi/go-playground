package ecbank

import (
	"fmt"
	"net/http"

	"moneyconverter/money"
)

// Client can call the bank to retrive exchange rate.
type Client struct{}

const (
	ErrCallingServer        = ecbankError("error calling server")
	ErrClientSide           = ecbankError("error client side")
	ErrServerSide           = ecbankError("error server side")
	ErrUnknownStatusCode    = ecbankError("unknown status code")
	ErrUnexpectedFormat     = ecbankError("unexpected format")
	ErrExchangeRateNotFound = ecbankError("exchange rate not found")
)

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

// FetchExchangeRate fetches the Exchange rate of the day and returns it.
func (c *Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const path = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	resp, err := http.Get(path)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrCallingServer, err.Error())
	}

	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil
}

// checkStatusCode returns a different error
// depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

// httpStatusClass returns the class of a http status code.
func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}
