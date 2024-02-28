package profitCalculator

import (
	"fmt"
	"os"
)

// to write the values to a file
func writeValuesToFile(ebt float64, profit float64, ratio float64, profitFile string) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(profitFile, []byte(results), 0644)
}
