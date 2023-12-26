package main

import (
	"errors"
	"fmt"
	"os"
)

const profitFile = "profitCalculator.txt"

func writeValuesToFile(ebt float64, profit float64, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(profitFile, []byte(results), 0644)
}

func calculateProfit() {
	fmt.Println("Welcome to the \"Profit Calculator\"")

	revenue, err1 := inputValue("Revenue: ")
	expenses, err2 := inputValue("Expenses: ")
	taxRate, err3 := inputValue("Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("ERROR")
		fmt.Println(err1)
		fmt.Println("----------")
		panic("the program can't continue further")
	}

	ebt, profit, ratio := calcValue(revenue, expenses, taxRate)
	writeValuesToFile(ebt, profit, ratio)
	printValue(ebt, profit, ratio)
}

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
