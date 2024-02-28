package bank

import (
	"fmt"

	"whatIsThis.com/fileops"
)

const accountBalanceFile = "balance.txt"

// a small bank with balance-check, deposit, withdraw options
func DummyBank() {
	// get balance
	accountBalance, err := fileops.GetFloatFromFile("bank/" + accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("The program can't continue further.")
	}

	fmt.Println("Welcome to GO Bank!")

	// main loop
	for {
		presentOptions()

		// users choice
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println("\n")

		switch choice {

		// balance-check
		case 1:
			fmt.Println("Your balance is", accountBalance)
			fmt.Print("\n")
		// deposit money
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
			fileops.WriteFloatToFile("bank/"+accountBalanceFile, accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
			fmt.Print("\n")
		// withdraw money
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
			fileops.WriteFloatToFile("bank/"+accountBalanceFile, accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
			fmt.Print("\n")
		// exit
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
