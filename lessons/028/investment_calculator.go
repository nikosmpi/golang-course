package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	//fmt.Print("Give the investment amount: ")
	outputText("Give the investment amount: ")
	fmt.Scan(&investmentAmount)
	//fmt.Print("Give the expected return rate: ")
	outputText("Give the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	//fmt.Print("Give the years: ")
	outputText("Give the years: ")
	fmt.Scan(&years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	//formatedFutureValue := fmt.Sprintf("The future value is: %.1f \n", futureValue)
	//formatedFutureRealValue := fmt.Sprintf("The future value (inflation) is: %.1f \n", futureRealValue)

	// fmt.Println("The future value is: ", futureValue)
	fmt.Printf("The future value is: %.1f \n The future value (inflation) is: %.1f \n", futureValue, futureRealValue)
	//	fmt.Printf(`The future value is: %.1f
	//The future value (inflation) is: %.1f`, futureValue, futureRealValue)
	//fmt.Println("The future value (inflation) is: ", futureRealValue)
	//fmt.Print(formatedFutureValue, formatedFutureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	//return
}
