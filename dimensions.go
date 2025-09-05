package dimension


import (
	"fmt"
)


func GetDimension(prompt string) int {
	
	var output int

	for {
		output = getIntData(prompt)

		if invalidDimension(output) {
			fmt.Println("Dimensions must be between 5 and 100")
			fmt.Errorf("error: invalid dimension")
			continue
		}

		break
	}

	return output
}


func invalidDimension(toBeValidated int) bool {
	return toBeValidated < 5 || toBeValidated > 100
}

func getIntData(prompt string) (output int) {
	fmt.Print(prompt)
	fmt.Scanln(&output)
	return output
}



func getVolume(height int, width int, depth int) int {
	return height * width * depth
}
