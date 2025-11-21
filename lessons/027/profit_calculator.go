package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Give the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Give the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Give the tax rate: ")
	fmt.Scan(&taxRate)
	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit
	fmt.Print("Earning Before Taxs: ")
	fmt.Println(earningBeforeTax)
	fmt.Print("Profit: ")
	fmt.Println(profit)
	fmt.Print("Ratio: ")
	fmt.Println(ratio)

}
