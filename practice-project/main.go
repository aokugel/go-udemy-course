package main

import (
	"fmt"

	"example.com/practice-project/fileops"
	"example.com/practice-project/prices"
)

func main() {

	//pick either the fileops or cmdmanager struct
	fm := fileops.New("prices.txt", "taxrates.txt", "results.json")
	//cmdm := cmdmanager.New()

	priceList, taxRates, err1 := fm.GetInputFromFile()

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	//the slice of Taxincluded price jobs to run the calculations.
	var taxRateJobList []prices.TaxIncludedPriceJob

	for _, rate := range taxRates {
		priceJob := prices.New(priceList, rate, fm)
		priceJob.Process()
		taxRateJobList = append(taxRateJobList, *priceJob)
	}

	fm.WriteJson(&taxRateJobList)

}
