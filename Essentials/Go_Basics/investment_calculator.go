package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	futureValue, furtueRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	fmt.Println(futureValue)
	fmt.Println(furtueRealValue)

}

func calculateFutureValues(iA float64, eRR float64, y float64, iR float64) (float64, float64) {
	fV := iA * math.Pow(1+eRR/100, y)
	rFV := fV / math.Pow(1+iR/100, y)
	return fV, rFV
}
