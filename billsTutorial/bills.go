package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Default highlights >>> TODO: FIXME:
// Custom highlights >>> REVIEW: NOTE: WARNING:
type Bill struct {
	//Define a bill struct (basically a js object) with 3 "fields"(keys)
	name  string
	items map[string]float64
	tip   float64
}

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	//Returns everything typed into the terminal, as a string, right before enter is pressed. Must be in SINGLE quotes, and assign it to the variable name (or error if there is an error).
	return strings.TrimSpace(input), err
	//Update the name variable with whitespace removed
}
func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	//Reader that reads from terminal (Stdin = terminal)
	name, _ := GetInput("Create a new bill name: ", reader)
	//Prints the string "Create a new bill name", and stores the user's input into the variable 'name' (also returns an error as second return, but not assigned due to the "_" second variable)
	b := NewBill(name)
	// Creates a new bill using the newBill func accessed from billsLogic.go (specifically associated with bill structs only), passing in the name inputted by the user.
	fmt.Println("Created the bill - ", b.name)
	//Print  the created bill's name & return the bill struct
	return b
}
func PromptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	//As above
	opt, _ := GetInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	// Prints the string "Choose option ..." and stores the user's input into the variable 'opt'
	switch opt {
	//S/C statement to carry out different functions/ store the subsequent input into various different variables. If an invalid option, reprompt the user.
	case "a":
		name, _ := GetInput("Item name: ", reader)
		price, _ := GetInput("Item price: ", reader)
		//Self-explanatory by now ...
		// >>> NOTE: ALL USER INPUTS TO THE READER ARE STRINGS BY DEFAULT <<<
		p, err := strconv.ParseFloat(price, 64)
		// Convert the variable 'price' into type `float64`
		if err != nil {
			//I.e. If there is an error
			fmt.Println("The price must be a number")
			PromptOptions(b)
			// Reprompt the user
		}
		fmt.Println(name, price)
		b.AddItem(name, p)
		// Call the addItem method from billsLogic.go (specifically associated with bill structs only), passing in name & price. Creates a new entry on the bill for that item
		fmt.Println("item added - ", name, price)
		PromptOptions(b)
		//Reprompt the user to either add another item, add a tip or save the bill
	case "s":
		b.Save()
		// Call the save method from billsLogic.go (specifically associated with bill structs only), thus saving the bill as a text file to the bills folder.
		fmt.Println("You saved the file - ", b.name)
	case "t":
		//Basically the same as adding an item to the bill above, except with a different method.
		tip, _ := GetInput("Enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			PromptOptions(b)
		}
		b.UpdateTip(t)
		//Passes in the tip to the update tip function.
		/*REVIEW: The function does not need to have the address passed in, despite taking a pointer
		Since the function updateTip is associated with the pointer to a bill i.e. func (b* bill) updateTip,
		a COPY OF THE POINTER will be copied and passed in (instead of a copy of the variable)
		*/
		fmt.Println("tip added - ", tip)
		PromptOptions(b)
		// Add the tip entry to the bill with the user-inputted tip, then reprompt the user

	default:
		fmt.Println("Invalid option")
		PromptOptions(b)
		// Inform the user the option they submitted was invalid, then reprompt the user
	}
}

/*	Running order:
	createBill - returns bill
		calls GetInput to get name
	PromptOptions(bill)
		calls GetInput to get user choice of action - Assuming user chooses a -> t -> s:
	user enters "a"
	calls GetInput 2x to get name and price
		casts price to p float64
	bill.addItem(name,p)
	PromptOptions(bill)
	user enters "t"
	calls GetInput to get tip
		casts tip to t float64
	bill.updateTip(t)
	PromptOptions(bill)
	user enters "s"
	bill.saveBill()
	bill is created and saved
*/
