package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("failed to read the file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("failed to parse the float")
	}
	return balance, nil
}

func dummyBank() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("The program can't continue further.")
	}

	fmt.Println("Welcome to GO Bank!")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
			fmt.Print("\n")
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
