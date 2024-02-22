package profitCalculator

import (
	"errors"
	"fmt"
)

func inputValue(printText string) (float64, error) {
	fmt.Print(printText)
	var scanVal float64
	fmt.Scan(&scanVal)

	if scanVal < 0 {
		return scanVal, errors.New("it can't be less than 0")
	} else if scanVal == 0 {
		return scanVal, errors.New("it can't be equal to zero")
	} else {
		return scanVal, nil
	}
}

func calcValue(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses //earnings before tax
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func printValue(ebt float64, profit float64, ratio float64) {
	fmt.Printf("Earning Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
