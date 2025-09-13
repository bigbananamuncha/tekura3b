package dimensions

import (
	"fmt"
	"github.com/bigbananamuncha/tekura3b/common/userinput"
)

func GetDimension(prompt string) int {
	
	var output int

	for {
		output = userinput.GetIntData(prompt)
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

func GetVolume(height int, width int, depth int) int {
	return height * width * depth
}
