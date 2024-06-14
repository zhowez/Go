package cmdmanager

import (
	"fmt"
)



type CMDManger struct {

}

func (cmd CMDManger) ReadLines()  ([]string, error) {

	var prices []string

	fmt.Println("please enter your prices. Confirm every price with ENTER")

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if (price == "0") {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil


}

func (cmd CMDManger) WriteResult(data any) error{

	fmt.Println(data)

	return nil

}

func New()  CMDManger {
	return CMDManger{}
}