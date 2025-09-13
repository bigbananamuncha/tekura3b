package userinput

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetStringData(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	output, _ := reader.ReadString('\n')
	caser := cases.Title(language.Und)
	capitalizedString := caser.String(output)
	return strings.TrimSpace(capitalizedString)
}


func GetIntData(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	output, _ := strconv.Atoi(strings.TrimSpace(input))
	return output
}