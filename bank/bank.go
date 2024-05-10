package main

import (
	"fmt"
)

func main() {
	var option int
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")

	for option != 4 {
		printMenu()
		fmt.Scan(&option)
		fmt.Println("")

		switch option {
		case 1:
			checkBalance(accountBalance)

		case 2:
			depositMoney(&accountBalance)

		case 3:
			withdrawlAmount(&accountBalance)

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

func checkBalance(bal float64) {
	fmt.Printf("Your balance is: $%.2f\n", bal)

}

func depositMoney(bal *float64) {
	var deposit float64
	fmt.Printf("How much would you like to deposit: ")
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("The deposit must be greater than 0")
		fmt.Println("Returning to the Main Menu")
		return
	}

	*bal += deposit

	fmt.Printf("You have deposited: $%.2f\n", deposit)

	checkBalance(*bal)
}

func withdrawlAmount(bal *float64) {
	var withdrawl float64
	fmt.Printf("How much would you like to deposit: ")
	fmt.Scan(&withdrawl)

	if withdrawl <= 0 || withdrawl > *bal {
		fmt.Println("The withdrawl amount must be greater than 0 or less than the balance")
		fmt.Println("Returning to the Main Menu")
		return
	}

	*bal -= withdrawl

	fmt.Printf("You have withdrawn: $%.2f\n", withdrawl)

	checkBalance(*bal)

}
