package main

import (
	"fmt"
	"max/proj1/cmdmanager"
	"max/proj1/filemanager"
	"max/proj1/prices"
)

func main() {
	taxRates := []float64{0, 7, 10, 15, 20}
	var inpOrFile string
	inpFlag := false
	var priceJob *prices.TaxIncludedPricesJob
	fmt.Println("You can either enter pre-tax prices or process prices.txt file.")
	fmt.Print("Do wish to enter pre-tax prices\nY or n >")
	fmt.Scanln(&inpOrFile)
	if inpOrFile == "Y" || inpOrFile == "y" {
		inpFlag = true
	}
	for _, taxRate := range taxRates {
		if inpFlag {
			cmdm := cmdmanager.New()
			priceJob = prices.NewTaxIncludedPricesJob(cmdm, taxRate)
		} else {
			fm := filemanager.New("prices.txt", fmt.Sprintf("Result_%.0f_Tax_Rate.json", taxRate))
			priceJob = prices.NewTaxIncludedPricesJob(fm, taxRate)
		}
		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Could not process job\nError = %v\n", err)
		}
	}
}
