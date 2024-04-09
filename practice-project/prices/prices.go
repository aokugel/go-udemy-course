package prices

import (
	"fmt"
)

type TaxIncludedPriceJob struct {
	Prices            []float64          `json:"prices"`
	TaxRate           float64            `json:"taxRate"`
	TaxIncludedPrices map[string]float64 `json:"taxIncludedPrices"`
}

func New(prices []float64, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		Prices:            prices,
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}
}

func (t *TaxIncludedPriceJob) Process() {

	for _, val := range t.Prices {
		s := fmt.Sprintf("%.2f", val)
		t.TaxIncludedPrices[s] = (t.TaxRate + 1) * val
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
