package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	fmt.Print("Give the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Give the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Give the years: ")
	fmt.Scan(&years)
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//formatedFutureValue := fmt.Sprintf("The future value is: %.1f \n", futureValue)
	//formatedFutureRealValue := fmt.Sprintf("The future value (inflation) is: %.1f \n", futureRealValue)

	// fmt.Println("The future value is: ", futureValue)
	fmt.Printf("The future value is: %.1f \n The future value (inflation) is: %.1f \n", futureValue, futureRealValue)
	//	fmt.Printf(`The future value is: %.1f
	//The future value (inflation) is: %.1f`, futureValue, futureRealValue)
	//fmt.Println("The future value (inflation) is: ", futureRealValue)
	//fmt.Print(formatedFutureValue, formatedFutureRealValue)
}
