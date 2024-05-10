package main

import (
	"fmt"
)

func main() {

	var revenue float64
	taxRate := 5.5
	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expeneses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := (revenue*(1-taxRate) - expenses)
	ratio := ebt / profit

	fmt.Printf("EBT: %f\n", ebt)
	fmt.Printf("Profit: %f\n", profit)
	fmt.Printf("Ratio: %f\n", ratio)

}
