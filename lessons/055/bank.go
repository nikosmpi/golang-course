package main

import (
	"fmt"

	"github.com/nikosmpi/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("---------")
		fmt.Println("ERROR:", err)
		fmt.Println("---------")
		// panic("Cant continue without balance file")
	}
	fmt.Println("Welcome to Go Bank!")
	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
		}
	}
}
