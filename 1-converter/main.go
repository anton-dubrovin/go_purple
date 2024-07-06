package main

import (
	"errors"
	"fmt"
)

const UsdToEur = 0.93
const UsdToRub = 88.2

const scanInputCurrencyPrefix = "Введите тип исходной валюты (USD, EUR, RUB): "
const scanOutputCurrencyPrefix = "Введите тип целевой валюты (USD, EUR, RUB): "
const scanAmountPrefix = "Введите значение исходной валюты: "

const invalidCurrencyErrorText = "введен невалидный тип валюты"
const invalidAmountErrorText = "введено невалидное значение валюты"
const negativeAmountErrorText = "значение валюты не должно быть отрицательным"

func scanCurrency(prefixPrint string) (string, error) {
	var currency string

	fmt.Print(prefixPrint)
	_, err := fmt.Scan(&currency)

	if err != nil || (currency != "USD" && currency != "EUR" && currency != "RUB") {
		return "", errors.New(invalidCurrencyErrorText)
	}

	return currency, nil
}

func scanAmount(prefixPrint string) (float64, error) {
	var amount float64

	fmt.Print(prefixPrint)
	_, err := fmt.Scan(&amount)

	if err != nil {
		var discard string
		_, _ = fmt.Scan(&discard)
		return 0, errors.New(invalidAmountErrorText)
	}

	if amount < 0 {
		return 0, errors.New(negativeAmountErrorText)
	}

	return amount, nil
}

func convert(inputAmount float64, inputCurrency string, outputCurrency string) float64 {
	if inputCurrency == outputCurrency {
		return inputAmount
	}

	outputAmount := inputAmount

	if inputCurrency == "EUR" {
		outputAmount /= UsdToEur
	} else if inputCurrency == "RUB" {
		outputAmount /= UsdToRub
	}

	if outputCurrency == "EUR" {
		outputAmount *= UsdToEur
	} else if outputCurrency == "RUB" {
		outputAmount *= UsdToRub
	}

	return outputAmount
}

func main() {
	var inputCurrency, outputCurrency string
	var inputAmount, outputAmount float64
	var err error

	for {
		inputCurrency, err = scanCurrency(scanInputCurrencyPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	for {
		inputAmount, err = scanAmount(scanAmountPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	for {
		outputCurrency, err = scanCurrency(scanOutputCurrencyPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	outputAmount = convert(inputAmount, inputCurrency, outputCurrency)
	fmt.Printf("%0.2f", outputAmount)

}
