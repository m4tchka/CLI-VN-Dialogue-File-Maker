package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	//Returns everything typed into the terminal, as a string, right before enter is pressed. Must be in SINGLE quotes, and assign it to the variable name (or error if there is an error).
	return strings.TrimSpace(input), err
	//Update the name variable with whitespace removed
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//Reader that reads from terminal
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		fmt.Println(name, price)
		b.addItem(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)

	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}

}
func main() {
	/* 	myBill := newBill("Kalndra's bill")
	   	myBill.addItem("onion soup", 4.50)
	   	myBill.addItem("veg pie", 8.95)
	   	myBill.addItem("toffee pudding", 4.95)
	   	myBill.addItem("coffee", 3.25)
	   	myBill.updateTip(10)
	   	fmt.Println(myBill.format()) */
	myBill := createBill()
	promptOptions(myBill)
}
