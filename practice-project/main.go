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

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, rate := range taxRates {
		doneChans[index] = make(chan bool)
		priceJob := prices.New(priceList, rate, fm)
		go priceJob.Process(doneChans[index], errorChans[index])
		taxRateJobList = append(taxRateJobList, *priceJob)
	}

	for i, _ := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}

	}
	fm.WriteJson(&taxRateJobList)

}
