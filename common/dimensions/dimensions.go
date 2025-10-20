package dimensions

import (
	"fmt"
	"github.com/bigbananamuncha/tekura3b/common/userinput"
)

//retrieving the dimensions of the package using the command line 
func GetDimension(prompt string) float64 {
	
	var output float64 

	//checking for valid inputs
	for {
		output = userinput.GetFloatData(prompt)
		if invalidDimension(output) {
			fmt.Println("Dimensions must be between 5 and 100")
			fmt.Errorf("error: invalid dimension")
			continue
		}

		break
	}

	return output
}

//checking the input is between the required rage 
func invalidDimension(toBeValidated float64) bool {
	return toBeValidated < 5 || toBeValidated > 100
}

//calculating the volume of the package
func GetVolume(height float64, width float64, depth float64) float64 {
	return height * width * depth
}
