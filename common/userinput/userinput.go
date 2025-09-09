package userinput

import "fmt"

func GetStringData(prompt string) (output string) {
	fmt.Print(prompt)
	fmt.Scanln(&output)
	return output
}


func GetIntData(prompt string) (output int) {
	fmt.Print(prompt)
	fmt.Scanln(&output)
	return output
}