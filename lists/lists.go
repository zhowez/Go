package main

import (
	"fmt"
)


func main() {
	prices := []float64{10.99,8.99}
	fmt.Println(prices[0])

	prices = append(prices, 5.99)
	prices = prices[1:]

	fmt.Println(prices)


}