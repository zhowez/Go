package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountBalFileName = "bal.txt"

func main() {
	var option int

	fmt.Println("Welcome to Go Bank!")

	_, err := readBalFile()

	if err != nil {
		fmt.Println("There was an error in Main")
		fmt.Println(err)
		panic("Can't continue. Sorry")
	}

	for option != 4 {
		printMenu()
		fmt.Scan(&option)
		fmt.Println("")

		switch option {
		case 1:
			checkBalance()

		case 2:
			depositMoney()

		case 3:
			withdrawlAmount()

		case 4:
			fmt.Println("Thank you\nGoodbye")

		default:
			fmt.Println("\nNot a valid option")
		}

	}

}

func printMenu() {

	fmt.Println("\nWhat do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Printf("Your choice: ")
}

func checkBalance() {

	bal, err := readBalFile()

	if err != nil {
		fmt.Println("There was an error in check balance")
		fmt.Println(err)
	}
	fmt.Printf("Your balance is: $%.2f\n", bal)

}

func depositMoney() {
	var deposit float64
	fmt.Printf("How much would you like to deposit: ")
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("The deposit must be greater than 0")
		fmt.Println("Returning to the Main Menu")
		return
	}

	bal, err := readBalFile()

	if err != nil {
		fmt.Println("There was an error in Deposit Money")
		fmt.Println(err)
	}

	bal += deposit

	fmt.Printf("You have deposited: $%.2f\n", deposit)

	writeBalToFile(bal)

	checkBalance()
}

func withdrawlAmount() {
	var withdrawl float64
	bal, err := readBalFile()

	if err != nil {
		fmt.Println("There was an error in withdrawl Money")
		fmt.Println(err)
	}
	fmt.Printf("How much would you like to deposit: ")
	fmt.Scan(&withdrawl)

	if withdrawl <= 0 || withdrawl > bal {
		fmt.Println("The withdrawl amount must be greater than 0 or less than the balance")
		fmt.Println("Returning to the Main Menu")
		return
	}

	bal -= withdrawl

	fmt.Printf("You have withdrawn: $%.2f\n", withdrawl)
	writeBalToFile(bal)

	checkBalance()

}

func writeBalToFile(bal float64) {
	balText := fmt.Sprint(bal)
	os.WriteFile(accountBalFileName, []byte(balText), 0644)
}

func readBalFile() (float64, error) {
	data, err := os.ReadFile(accountBalFileName)

	if err != nil {
		writeBalToFile(1000.0)
		return 1000, errors.New("Failed to find file. Created new one and deposited $1000")
	}
	balText := string(data)
	bal, err := strconv.ParseFloat(balText, 64)

	if err != nil {
		writeBalToFile(1000.0)
		return 1000, errors.New("Failed to read file. Created new one and deposited $1000")
	}

	return bal, nil
}
