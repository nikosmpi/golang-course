package main

import (
	"errors"
	"fmt"
	"os"
)

func writeResults(earningBeforeTax, profit, ratio float64) {
	results := fmt.Sprintf("Earning Before Taxs: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningBeforeTax, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {
	revenue, err := getUserInput("Give the revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Give the expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Give the tax rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeResults(earningBeforeTax, profit, ratio)
	fmt.Printf("Earning Before Taxs: %.1f\n", earningBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid number")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit
	return earningBeforeTax, profit, ratio
}
