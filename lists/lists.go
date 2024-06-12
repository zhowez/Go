package main

import (
	"fmt"
)


func main() {

	fmt.Println("#1")
	hobbyArr := [3]string{"disney","weights","walking"}
	fmt.Println(hobbyArr)

	fmt.Printf("First Element: %s\n",hobbyArr[0] )

	fmt.Println("#2")
	newHob := hobbyArr[1:3]
	fmt.Printf("2nd and 3rd")
	fmt.Print(newHob,"\n")
	
	fmt.Println("#3")

	v1 := hobbyArr[:2]
	v2 := hobbyArr[0:2]

	fmt.Print("V1: ", v1, "\n")
	fmt.Print("V2: ", v2, "\n")

	fmt.Println("#4")

	v1 = v1[1:3]
	fmt.Print("V1 re-sliced: ", v1, "\n")


	fmt.Println("#5")

	goals := []string{"learn go", "teach go"}
	fmt.Print("Goals OG: ", goals, "\n")


	fmt.Println("#6")
	goals[1] = "use go"
	goals = append(goals, "work go")
	fmt.Print("Goals REMIX: ", goals, "\n")


	fmt.Println("#7")
	type product struct {
		title string
		id int
		price float64
	}

	products := []product{{title: "Ice", id: 0, price:2.99}, {title: "Cookies", id: 1, price:5.99}}

	fmt.Print("Products: ", products, "\n")



}