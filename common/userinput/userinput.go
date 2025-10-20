package userinput

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
//transfering the data from the prompt and from the command line into a string
func GetStringData(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    output, _ := reader.ReadString('\n')
	caser := cases.Title(language.Und)
	//capitalising the first letter of each word
    capitalizedString := caser.String(output)
	//trimming any uneeded spaces 
    return strings.TrimSpace(capitalizedString)
}

//transfering the data from the prompt and from the command line into an int
func GetIntData(prompt string) int {
    reader := bufio.NewReader(os.Stdin)
	//handling invalid inputs 
    for {
        fmt.Print(prompt)
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input. Please try again.")
            continue
        }
		//trimming uneeded space
        trimmedInput := strings.TrimSpace(input)
        number, err := strconv.Atoi(trimmedInput)
        
        if err != nil {
            fmt.Printf("Invalid input: '%s' is not a whole number. Please enter an integer.\n", trimmedInput)
            continue 
        } 
        if number < 0 {
            fmt.Println("Invalid input: The number cannot be negative. Please enter a non-negative integer.")
            continue
        }else {
			return number
		}
    }
}

//transfering the data from the prompt and from the command line into a float
func GetFloatData(prompt string) float64 {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(prompt)
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input. Please try again.")
            continue
        }
        
		//trimming uneeded space
        trimmedInput := strings.TrimSpace(input)
        output, err := strconv.ParseFloat(trimmedInput, 64)
        
		//invalid input handling
        if err != nil {
            fmt.Printf("Invalid input: '%s' is not a valid number. Please enter a number.\n", trimmedInput)
            continue 
        } else {

            return output
        }
    }
}