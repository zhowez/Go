package main

import (
	"fmt"
)

func main() {

	var revenue float64
	taxRate := 5.5
	var expenses float64

	requestInput("Revenue", &revenue)

	requestInput("Expenes", &expenses)

	requestInput("Tax Rate", &taxRate)

	fmt.Println(revenue)
	fmt.Println(expenses)
	fmt.Println(taxRate)

	ebt, profit, ratio := calcFinanceInfo(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

}

func requestInput(question string, value *float64) {
	fmt.Printf("%s: ", question)
	fmt.Scan(value)
}

func calcFinanceInfo(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

	return ebt, profit, ratio

}
