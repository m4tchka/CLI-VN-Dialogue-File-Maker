package main

func main() {
	myBill := CreateBill()
	// Call the createBill function, then prompt the user to perform 1 of 3 options regarding that bill
	PromptOptions(myBill)
}
