package userinput

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)
func GetStringData(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	output, _ := reader.ReadString('\n')
	return strings.TrimSpace(output)
}


func GetIntData(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	output, _ := strconv.Atoi(strings.TrimSpace(input))
	return output
}