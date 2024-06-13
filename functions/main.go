package main

import "fmt"

type transformFn func(int) int
func main() {

	numbers := []int{1,2,3,4}

	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, tripple)

	fmt.Println(numbers)

	fmt.Println(doubled)

	fmt.Println(trippled)



}


func transformNumbers (numbers *[]int, transform transformFn) []int{
	dNumbers := []int{}
	for _, val:= range *numbers {
		dNumbers = append(dNumbers,transform(val))
	}

	return dNumbers
}

func getTransformerFunction() transformFn {
	return double
}

func double(number int) int {
	return number * 2
}

func tripple(number int) int {
	return number * 3
}