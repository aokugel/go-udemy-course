package prices

import (
	"fmt"
	"strconv"

	"example.com/practice-project/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManger          iomanager.IOmanager `json:"-"`
	Prices            []string            `json:"prices"`
	TaxRate           string              `json:"taxRate"`
	TaxIncludedPrices map[string]float64  `json:"taxIncludedPrices"`
}

func New(prices []string, taxRate string, io iomanager.IOmanager) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		Prices:            prices,
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
		IOManger:          io,
	}
}

func (t *TaxIncludedPriceJob) Process() {

	for _, val := range t.Prices {
		float64tax, err := strconv.ParseFloat(t.TaxRate, 64)
		float64price, err2 := strconv.ParseFloat(val, 64)
		t.TaxIncludedPrices[val] = (float64tax + 1.0) * float64price
		if err != nil || err2 != nil {
			fmt.Println(err, err2)
		}
	}

	fmt.Println(t.TaxIncludedPrices)

}

func PriceCalculation(prices []float64, taxRate float64) map[string]float64 {
	tip := make(map[string]float64)

	for _, val := range prices {
		s := fmt.Sprintf("%.2f", val)
		tip[s] = (taxRate + 1) * val

	}
	return tip
}
