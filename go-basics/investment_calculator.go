package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 6.5

func main() {
	const inflationRate float64 = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years: ")
	fmt.Scan(&years)

	fmt.Print("Rate of return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Final amount of money is: is ", futureValue)
	// fmt.Println("Final real value of money of money is: is ", futureRealValue)

	// fmt.Printf("Final amount of money is: is %.2f \n", futureValue)
	// fmt.Printf("Final real value of money of money is: is %.2f \n", futureRealValue)

	// fmt.Printf(`Final amount of money is: is %.1f \n

	// Final real value of money of money is: is %.1f \n`, futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Final real value of money of money is: is %.2f", futureValue)
	formattedRFV := fmt.Sprintf("Final real value of money of money is: is %.2f", futureRealValue)

	outputText(formattedFV, formattedRFV)
	outputText()

	// fmt.Print(formattedFV, formattedRFV)
}

func outputText(text ...string) {
	for _, element := range text {
		fmt.Println(element)
	}
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(expectedReturnRate/100+1, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// func calculateRealFutureValue(futureValue, inflationRate, years float64) float64 {
// 	return futureValue / math.Pow(1+inflationRate/100, years)
// }
