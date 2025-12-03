package main

import (
	"fmt"

	"github.com/nikosmpi/price-calculator/filemanager"
	"github.com/nikosmpi/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	donesChans := make([]chan bool, len(taxRates))
	errorsChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		donesChans[index] = make(chan bool)
		errorsChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		go priceJob.Process(donesChans[index], errorsChans[index])
		//if err != nil {
		//	fmt.Println(err)
		//}
	}

	for i := range taxRates {
		select {
		case err := <-errorsChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-donesChans[i]:
			fmt.Println("done")
		}
	}
}
