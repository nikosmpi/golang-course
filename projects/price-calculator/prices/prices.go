package prices

import (
	"fmt"
	"strconv"

	"github.com/nikosmpi/price-calculator/conversion"
	"github.com/nikosmpi/price-calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", taxIncludedPrice), 64)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	return nil
}

func NewTaxIncludedPricesJob(iom iomanager.IOManager, TaxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: iom,
		TaxRate:   TaxRate,
	}
}
