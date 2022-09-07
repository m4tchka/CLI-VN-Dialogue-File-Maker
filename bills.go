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

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	//Returns everything typed into the terminal, as a string, right before enter is pressed. Must be in SINGLE quotes, and assign it to the variable name (or error if there is an error).
	return strings.TrimSpace(input), err
	//Update the name variable with whitespace removed
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//Reader that reads from terminal (Stdin = terminal)
	name, _ := getInput("Create a new bill name: ", reader)
	//Prints the string "Create a new bill name", and stores the user's input into the variable 'name' (also returns an error as second return, but not assigned due to the "_" second variable)
	b := newBill(name)
	// Creates a new bill using the newBill func accessed from billsLogic.go (specifically associated with bill structs only), passing in the name inputted by the user.
	fmt.Println("Created the bill - ", b.name)
	//Print  the created bill's name & return the bill struct
	return b
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	//As above
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	// Prints the string "Choose option ..." and stores the user's input into the variable 'opt'
	switch opt {
	//S/C statement to carry out different functions/ store the subsequent input into various different variables. If an invalid option, reprompt the user.
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		//Self-explanatory by now ...
		// >>> NOTE: ALL USER INPUTS TO THE READER ARE STRINGS BY DEFAULT <<<
		p, err := strconv.ParseFloat(price, 64)
		// Convert the variable 'price' into type `float64`
		if err != nil {
			//I.e. If there is an error
			fmt.Println("The price must be a number")
			promptOptions(b)
			// Reprompt the user
		}
		fmt.Println(name, price)
		b.addItem(name, p)
		// Call the addItem method from billsLogic.go (specifically associated with bill structs only), passing in name & price. Creates a new entry on the bill for that item
		fmt.Println("item added - ", name, price)
		promptOptions(b)
		//Reprompt the user to either add another item, add a tip or save the bill
	case "s":
		b.save()
		// Call the save method from billsLogic.go (specifically associated with bill structs only), thus saving the bill as a text file to the bills folder.
		fmt.Println("You saved the file - ", b.name)
	case "t":
		//Basically the same as adding an item to the bill above, except with a different method.
		tip, _ := getInput("Enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		//Passes in the tip to the update tip function.
		/*REVIEW: The function does not need to have the address passed in, despite taking a pointer
		Since the function updateTip is associated with the pointer to a bill i.e. func (b* bill) updateTip,
		a COPY OF THE POINTER will be copied and passed in (instead of a copy of the variable)
		*/
		fmt.Println("tip added - ", tip)
		promptOptions(b)
		// Add the tip entry to the bill with the user-inputted tip, then reprompt the user

	default:
		fmt.Println("Invalid option")
		promptOptions(b)
		// Inform the user the option they submitted was invalid, then reprompt the user
	}
}
func mainx() {
	myBill := createBill()
	// Call the createBill function, then prompt the user to perform 1 of 3 options regarding that bill
	promptOptions(myBill)
}

/*	Running order:
	createBill - returns bill
		calls getInput to get name
	promptOptions(bill)
		calls getInput to get user choice of action - Assuming user chooses a -> t -> s:
	user enters "a"
	calls getInput 2x to get name and price
		casts price to p float64
	bill.addItem(name,p)
	promptOptions(bill)
	user enters "t"
	calls getInput to get tip
		casts tip to t float64
	bill.updateTip(t)
	promptOptions(bill)
	user enters "s"
	bill.saveBill()
	bill is created and saved
*/
