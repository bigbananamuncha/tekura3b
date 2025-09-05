package main

import (
	"fmt"
	
)

type customer struct {
	firstName string
	lastName string
	age int
	address string
	island int 
	shippingCost int
}

func main() {
	var cust customer

	// Get the customer details
	cust.age = dimension.GetIntData("Enter your age: ")
	cust.firstName = getStringData("Enter your first name: ")

	fmt.Printf("Hi %s, you are %d years old.\n", cust.firstName, cust.age)

	// Get the package volume
	height := getDimension("Enter the height, must be between 5 and 100: ")
	width := getDimension("Enter the width, must be between 5 and 100: ")
	depth := getDimension("Enter the depth, must be between 5 and 100: ")


	cust.shippingCost

	volume := getVolume(height, width, depth)

	fmt.Printf("Hi %s, your box is %d cm cubed.\n", cust.firstName, volume)

}
	// Get the customer lastname and address 

func costCalculator() {
	var cost int
	if (volume <= 6000) {
		cost = 8
	} else if (volume > 6000 || volume < 100000) {
		cost = 12
	}
	cost = 15
}


func getStringData(prompt string) (output string) {
	fmt.Print(prompt)
	fmt.Scanln(&output)
	return output
}



func getBaseRate(volume int) int {
	const lowRate = 800
	const mediumRate = 1200
	const highRate = 1500

	if volume <= 6000 {
		return lowRate
	} else if volume > 6000 && volume < 100000 {
		return mediumRate
	} 
	return highRate
}



func getShippingRate(baseRate int, island int) int {
	const northIsland = 1
	const northIslandRate = 10
	const southIsland = 2
	const southIslandRate = 15
	const stewartIsland = 3
	const stewartIslandRate = 20
	
	if island == northIsland {
		return baseRate * northIslandRate
	} else if island == southIsland {
		return baseRate * southIslandRate
	}
	return baseRate * stewartIslandRate
}