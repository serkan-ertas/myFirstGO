package profitCalculator

import (
	"fmt"
)

// to create file in correct directory
const profitFile = "profitCalculator/profitCalculator.txt"

// used this function to call in main
func CalculateProfit() {
	fmt.Println("Welcome to the \"Profit Calculator\"")

	// get the values
	revenue, err1 := inputValue("Revenue: ")
	expenses, err2 := inputValue("Expenses: ")
	taxRate, err3 := inputValue("Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("ERROR")
		fmt.Println(err1)
		fmt.Println("----------")
		panic("the program can't continue further")
	}

	// calculate, print and save the values
	ebt, profit, ratio := calcValue(revenue, expenses, taxRate)
	writeValuesToFile(ebt, profit, ratio, profitFile)
	printValue(ebt, profit, ratio)
}
