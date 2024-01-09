package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Your EBT is: %g\n", ebt)

	fmt.Printf("Your profit is: %g\n", profit)

	fmt.Printf("Your ratio is: %g\n", ratio)
}

func getInput(prompt string) (inputValue float64) {
	fmt.Print(prompt)
	fmt.Scan(&inputValue)
	return inputValue
}
