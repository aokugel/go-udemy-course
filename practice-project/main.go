package main

import (
	"fmt"

	"example.com/practice-project/fileops"
	"example.com/practice-project/prices"
)

func main() {

	fm := fileops.New("prices.txt", "taxrates.txt", "results.json")

	taxRates, err1 := fm.GetInputFromFile("operand")
	priceList, err2 := fm.GetInputFromFile("input")

	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return
	}

	var taxRateJobList []prices.TaxIncludedPriceJob

	for _, rate := range taxRates {
		priceJob := prices.New(priceList, rate)
		priceJob.Process()
		taxRateJobList = append(taxRateJobList, *priceJob)
	}

	fmt.Println(fileops.WriteJson(&taxRateJobList, "testFile.json"))

}
