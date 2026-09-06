package main

import (
	"flag"
	"fmt"
	"os"

	"moneyconverter/ecbank"
	"moneyconverter/money"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")
	flag.Parse()

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse value %q: %s\n", *from, err.Error())
		os.Exit(1)
	}

	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse value %q: %s\n", *to, err.Error())
		os.Exit(1)
	}

	value := flag.Arg(0)
	if value == "" {
		fmt.Fprintf(os.Stderr, "mising amount to convert\n")
		flag.Usage()
		os.Exit(1)
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse value %q: %s\n", value, err.Error())
		os.Exit(1)
	}

	amount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	rates := &ecbank.Client{}

	convertedAmount, err := money.Convert(amount, toCurrency, rates)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s = %s\n", amount.String(), convertedAmount.String())
}
