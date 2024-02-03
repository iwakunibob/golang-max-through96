package prices

import (
	"fmt"
	"max/proj1/conversion"
	"max/proj1/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludePrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.LoadPrices()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate/100.0)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
	job.TaxIncludePrices = result
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPricesJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30, 40, 50},
		TaxRate:     taxRate,
	}
}
