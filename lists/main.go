package main

import "fmt"

type floatMap map[string]float64

func main() {

	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)


	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angualar"] = 4.7

	fmt.Println(courseRatings)


	for index, value := range userNames {

		fmt.Println(index)
		fmt.Println(value)
	}


	
}