//lint:file-ignore U1000 dont wanna see unused val/func
package main

import (
	"fmt"

	"whatIsThis.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func dummyBank() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("The program can't continue further.")
	}

	fmt.Println("Welcome to GO Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
			fmt.Print("\n")
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				fmt.Print("\n")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
			fmt.Print("\n")
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				fmt.Print("\n")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cannot withdraw more than you have.")
				fmt.Print("\n")
				continue
			}

			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
			fmt.Print("\n")
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
