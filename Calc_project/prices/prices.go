package prices

import (
	"fmt"

	"example.com/calc/conversion"
	"example.com/calc/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string] string `json:"tax_included"`

}


func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	err := job.LoadData()

	if err != nil {
		fmt.Println(err)
		errorChan <- err
		return
	}


	result :=  make(map[string]string)


	for _, price := range job.InputPrices {

		taxIncludedPrice := fmt.Sprintf("%.2f", price * (1 + job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = taxIncludedPrice
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
	doneChan<-true

}


func (job *TaxIncludedPriceJob) LoadData() error{
	
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


func NewTaxIncludedPriceJob(iom iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		IOManager: iom,
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}
