package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	
	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func (number int) int { return number * 2 })

	doubled := transformNumbers(&numbers, double)

	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)

	fmt.Println(tripled)


	fmt.Println(transformed)

	fmt.Println(factorial(3))
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}


func createTransformer(factor int)  func(int) int{ 
	return func(number int) int {
		return number * factor
	}
}

func factorial (number int) int {
	if (number == 0) {
		return 1
	}

	return factorial(number - 1) * number  
}