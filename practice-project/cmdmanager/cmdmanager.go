package cmdmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type CMDManager struct{}

func (cmd *CMDManager) GetInputFromFile() ([]string, []string, error) {

	fmt.Printf("Please enter Prices, separate each price with an enter press\nEnter -1 to escape\n")

	var prices []string
	for {
		var price string
		fmt.Printf("Price: ")
		fmt.Scan(&price)
		if price == "-1" {
			break
		}
		prices = append(prices, price)
	}

	fmt.Printf("Please enter Tax Rates, separate each Tax Rates with an enter press\nEnter -1 to escape\n")

	var taxRates []string
	for {
		var taxRate string
		fmt.Printf("Tax Rate: ")
		fmt.Scan(&taxRate)
		if taxRate == "-1" {
			break
		}
		taxRates = append(taxRates, taxRate)
	}

	return prices, taxRates, nil

}

func (cmd *CMDManager) WriteJson(data interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
