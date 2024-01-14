package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		//os.WriteFile(accountBalanceFile, []byte("0"), 0644)
		return 0, errors.New("accountBalanceFile not found")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var choice int
	var amount float64

	for {
		var balance, err = getBalanceFromFile()
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("---------")
			panic("Can't continue, sorry.")
		}
		fmt.Println("What do you want to do?\n1. Check Balance \n2. Deposit Money\n3. Withdraw money\n4. Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Balance is $%.2f\n", balance)
		case 2:
			fmt.Println("How much do you want to deposit?")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount, please try again")
				continue
			} else {
				balance += amount
				fmt.Println("balance updated!  New amount is: ", balance)
				writeBalanceToFile(balance)
			}
		case 3:
			fmt.Println("How much do you want to withdraw")
			fmt.Scan(&amount)
			if amount < balance {
				balance -= amount
				fmt.Println("balance updated!  New amount is: ", balance)
				writeBalanceToFile(balance)
			} else {
				fmt.Println("Insufficient funds")
			}
		default:
			fmt.Println("Peace")
			fmt.Println("Thanks for the bank")
			return
		}
	}

}
