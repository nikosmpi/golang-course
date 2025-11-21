package main

import "fmt"

func main() {
	revenue := getUserInput("Give the revenue: ")
	expenses := getUserInput("Give the expenses: ")
	taxRate := getUserInput("Give the tax rate: ")
	earningBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("Earning Before Taxs: %.1f\n", earningBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit
	return earningBeforeTax, profit, ratio
}
