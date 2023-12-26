package main

import (
	"fmt"
	"math"
)

func calculateInvestment() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.2f\n", realFutureValue)

	fmt.Print(formattedFV, formattedRFV)
}
