package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var choice int
	var amount float64

	fmt.Println("welcome to go bank!")
	fmt.Println("Reach us 24/7 at: ", randomdata.PhoneNumber())

	for {
		var balance, err = fileops.GetFloatFromFile(accountBalanceFile)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("---------")
			// panic("Can't continue, sorry.")
		}
		presentOptions()
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
				fileops.WriteFloatToFile(balance, accountBalanceFile)
			}
		case 3:
			fmt.Println("How much do you want to withdraw")
			fmt.Scan(&amount)
			if amount < balance {
				balance -= amount
				fmt.Println("balance updated!  New amount is: ", balance)
				fileops.WriteFloatToFile(balance, accountBalanceFile)
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
