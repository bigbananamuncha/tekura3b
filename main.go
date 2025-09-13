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
	cust.firstName = userinput.GetStringData("Hello, please enter your first name: ")
	cust.lastName = userinput.GetStringData("Please enter your last name: ")
	fmt.Printf("Hi %s ", cust.firstName)
	cust.age = userinput.GetIntData("Please enter your age: ")

	fmt.Printf("Hi %s, you are %d years old.\n", cust.firstName, cust.age)

	// Get the package volume
	height := dimensions.GetDimension("Please enter the height of your package, must be between 5 and 100: ")
	width := dimensions.GetDimension("Please enter the width of your package, must be between 5 and 100: ")
	depth := dimensions.GetDimension("Enter the depth of your package, must be between 5 and 100: ")

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
	fmt.Printf("Here are your details: %+v\n", cust)
}