package main

import (
	"errors"
	"fmt"
	"os"
)

var accountBalanceFile = "financials.txt"

func main() {

	revenue, err1 := getInput("Enter your revenue: ")
	expenses, err2 := getInput("Enter your expenses: ")
	taxRate, err3 := getInput("Enter your tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1, err2, err3)
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)
	financials := storeResults(ebt, profit, ratio)
	fmt.Println(financials)
}

func getInput(prompt string) (inputValue float64, err error) {
	fmt.Print(prompt)
	fmt.Scan(&inputValue)
	if inputValue <= 0 {
		err = errors.New("Cannot input 0 or negative number")
		return 0, err
	}
	return inputValue, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) string {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(accountBalanceFile, []byte(results), 0644)
	return results
}
