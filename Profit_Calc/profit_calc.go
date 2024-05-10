package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var revenue float64
	taxRate := 5.5
	var expenses float64

	err := requestInput("Revenue", &revenue)

	if err != nil {
		fmt.Println("there was an error in main")
		fmt.Println(err)
		panic("can't continue. Sorry")
	}

	err = requestInput("Expenes", &expenses)

	if err != nil {
		fmt.Println("there was an error in main")
		fmt.Println(err)
		panic("Can't continue. Sorry")
	}

	err = requestInput("Tax Rate", &taxRate)

	if err != nil {
		fmt.Println("there was an error in main")
		fmt.Println(err)
		panic("Can't continue. Sorry")
	}

	ebt, profit, ratio := calcFinanceInfo(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

	writeInfoToFile(ebt, profit, ratio)

}

func requestInput(question string, value *float64) error {
	fmt.Printf("%s: ", question)
	fmt.Scan(value)

	if *value == 0 {
		return errors.New("value was 0")
	} else if *value < 0 {
		return errors.New("value was negitive")
	}

	return nil
}

func calcFinanceInfo(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio

}

func writeInfoToFile(ebt float64, profit float64, ratio float64) {
	finText := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("finance.txt", []byte(finText), 0644)

}
