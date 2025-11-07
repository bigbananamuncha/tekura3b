package main

import (
	"fmt"
	"github.com/bigbananamuncha/tekura3b/common/userinput"
	"github.com/bigbananamuncha/tekura3b/common/dimensions"
	"github.com/bigbananamuncha/tekura3b/common/cost"
)

type Customer struct {
	firstName string
	lastName string
	age int
	address string
	telephone string
	island string
	shippingCost float64
}

func main() {
	var cust Customer

	// --- create a function that does this and put it here - use a pointer for cust
	// Get the inital customer input
	// Name, last name, and age 
	cust.firstName = userinput.GetStringData("Hello, please enter your first name: ")
	cust.lastName = userinput.GetStringData("Please enter your last name: ")
	fmt.Printf("Hi %s ", cust.firstName)
	cust.age = userinput.GetIntData("Please enter your age: ")

	// Print the customers details
	fmt.Printf("Hi %s, you are %d years old.\n", cust.firstName, cust.age)
	// <----

	// ---> same here
	// Get the package dimensions
	height := dimensions.GetDimension("Please enter the height of your package, must be between 5 and 100: ")
	width := dimensions.GetDimension("Please enter the width of your package, must be between 5 and 100: ")
	depth := dimensions.GetDimension("Enter the depth of your package, must be between 5 and 100: ")

	// Find the volume of the package
	volume := dimensions.GetVolume(height, width, depth)

	// Print the package volume
	fmt.Printf("Hi %s, your box is %f cm cubed.\n", cust.firstName, volume)
	// <---

	// ---> same here
	// Get the island of residence for the customer
	
	for {
        // Use the userinput function to get the data
		cust.island = userinput.GetStringData("Please enter your island:\nNorth Island\nSouth Island\nStewart Island\n: ")        
        // Check if the input is valid using a function from the cost package
        if cost.IsValidIsland(cust.island) {
            break // Exit the loop if valid
        }
        
        // Error message for the user
        fmt.Println("Invalid island entered")
    }
	// <---
	
	// ---> same here
	// Find the cost of shipping
	baseRate := cost.GetBaseRate(volume)
	cust.shippingCost = cost.GetShippingRate(baseRate, cust.island)
	fmt.Printf("Your shipping cost is NZD $%.2f \n", cust.shippingCost)
	// <---

	// --- same here
	// Get the last details: adress and phone number, then print all the collected details. 
	cust.address = userinput.GetStringData("Please enter your address: ")
	fmt.Printf("Your address is %s.\n", cust.address)
	cust.telephone = userinput.GetStringData("Enter your telephone number: ")
	fmt.Printf("Your telephone number is %s.\n", cust.telephone)
	// Customer Details
	fmt.Printf("Here are your details\n")
	fmt.Printf("Customer Name:\t\t%s %s\n", cust.firstName, cust.lastName)
	fmt.Printf("Age:\t\t\t%d\n", cust.age)
	fmt.Printf("Phone:\t\t\t%s\n", cust.telephone)
	fmt.Printf("Shipping Address:\t%s\n", cust.address)
	// <---
}