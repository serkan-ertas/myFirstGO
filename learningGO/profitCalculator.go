package main

import (
	"fmt"
)

func calculateProfit() {
	fmt.Println("Welcome to the \"Profit Calculator\"")

	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses //earnings before tax
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Ebt is : ", ebt)
	fmt.Println("Profit is: ", profit)
	fmt.Println("Ratio = ", ratio)
}
