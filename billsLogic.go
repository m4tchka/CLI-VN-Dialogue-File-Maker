package main

import (
	"fmt"
	"os"
)

type bill struct {
	//Define a bill struct (basically a js object) with 3 "fields"(keys)
	name  string
	items map[string]float64
	tip   float64
}

/* Make new bills */
func newBill(name string) bill {
	//Define a new function
	b := bill{
		//create a new bill variable 'b'
		name: name,
		//Give it the argument passed in as the name parameter, as the value of its 'name' field
		items: map[string]float64{},
		tip:   0,
		//Initialise its 2 other required fields with "empty" values for now
	}
	return b
}

func (b *bill) format() string {
	//Define a new function that is specifically associated with variables of the 'bill' custom type.

	fs := "Bill breakdown: \n"
	// Define a variable 'fs' (formatted string) that will serve as the "title" of the text file
	var total float64 = 0

	/* list items */
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
		/* NOTE: Loop through all the items added to the bill, and append them to the title string.
		The -25 means force the %v variable that is inserted, to take up 25 character spaces,
		The variable will occupy as many spaces as needed, at the start of the line
		+25 would put the variable at the end of those 25 characters instead.
		The colon is appended to the first inserted variable (the key of the item field of the struct) so that it doesn't appear at the end of the 25 character space.
		*/
		// Also add each value to the variable 'total'
	}
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)
	//Same as above, but with tip instead
	/* total */
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	//Same as above, but add the tip to the total as well
	return fs
}

/* update tip */
func (b *bill) updateTip(tip float64) {
	//WARNING: Very specific note <<
	//NOTE: The receiver's type is a POINTER to a bill struct
	/* REVIEW: If the receiver is NOT a pointer to a bill variable (e.g. the receiver's type is just a normal ' bill ' instead of ' *bill '),
	then the bill that is actually passed into the function when it's called is merely a copy of that variable
	and this copy is what gets changed within the function, whilst the actual bill variable is left unchanged.
	Instead, pass a pointer (which itself gets copied - but still points to the original bill variable)

	*/
	b.tip = tip
	/* NOTE:
	Dereferencing structs specifically DOES NOT require explicit derefencing. The following is also valid however:
	(*b).tip = tip
	If the function is associated with a pointer, dereferencing is automatic
	so updating the tip field as above updates the actual value that the pointer points to
		NOTE: Additionally, passing a pointer only makes a copy of said pointer,
		as opposed to making a copy of a potentially huge struct with many fields.
	*/
}

/* add item to bill */
func (b *bill) addItem(newItem string, price float64) {
	// Define a new function associated with bill structs (well, technically pointers to them, but as above effectively the same.)
	b.items[newItem] = price
	// On the b struct passed in, access the items field, which contains a map of items against prices.
	// Within that, access the entry with the key of the same value as the newItem argument (which won't exist)
	// (Since the newITem key doesn't exist on that map, create a new key for that item, with a value of the price.)
}

/* save bill */
func (b *bill) save() {
	// Define the last function associated with bill structs
	data := []byte(b.format())
	// Define a new variable as a slice of bytes, and pass the string created from the format() function.
	// NOTE: In the slice of bytes, EACH CHARACTER IS AN INDIVIDUAL ENTRY
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	/* NOTE: The WriteFile function from the os package will take in 3 parameters:
	1. The filepath to where it should be saved, including extension. If a file of that name & path doesn't exist at that location, make that new file.
	2. The contents to be written to that new file, as an slice/array of bytes
	3. Permissions with/through which the writeFile can change the specified file - 0644 = read & write permissions - Should be the minimum required.
	*/
	if err != nil {
		panic(err)
		//If there is an error, emergency exit the program (since this can't handle any errors that may occur)
	}
	fmt.Println("Bill was saved to file")
}
