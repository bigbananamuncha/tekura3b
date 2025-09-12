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
	shippingCost int
}

func main() {
	var cust Customer

	// Get the inital customer input
	cust.age = userinput.GetIntData("Enter your age: ")
	cust.firstName = userinput.GetStringData("Enter your first name: ")

	fmt.Printf("Hi %s, you are %d years old.\n", cust.firstName, cust.age)

	// Get the package volume
	height := dimensions.GetDimension("Enter the height, must be between 5 and 100: ")
	width := dimensions.GetDimension("Enter the width, must be between 5 and 100: ")
	depth := dimensions.GetDimension("Enter the depth, must be between 5 and 100: ")

	volume := dimensions.GetVolume(height, width, depth)

	fmt.Printf("Hi %s, your box is %d cm cubed.\n", cust.firstName, volume)

	cust.island = userinput.GetStringData("Enter your island (North Island, South Island, or Stewart Island): ")
	
	baseRate := cost.GetBaseRate(volume)
	cust.shippingCost = cost.GetShippingRate(baseRate, cust.island)
	fmt.Printf("Your shipping cost is %d.\n", cust.shippingCost)

	cust.address = userinput.GetStringData("Enter your address: ")
	fmt.Printf("Your address is %s.\n", cust.address)
	cust.telephone = userinput.GetStringData("Enter your telephone number: ")
	fmt.Printf("Your telephone number is %s.\n", cust.telephone)
}

// You need to be able to handle invalid input. And handle "North" "north" "North Island"
// look up the "strings" package for help and maybe "regexp"

// TODO: Add a function to print the customer details
// TODO: Add a function that gets the volume of the package separate to the main function
// TODO: look for any other improvements you can make to the code where you move logical chunks out to functions