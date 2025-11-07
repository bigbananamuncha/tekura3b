package userinput

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func mockInput(t *testing.T, input string) func() {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	_, err = w.WriteString(input + "\n")
	if err != nil {
		t.Fatalf("Failed to write to pipe: %v", err)
	}
	w.Close()

	oldStdin := os.Stdin
	os.Stdin = r

	return func() {
		os.Stdin = oldStdin
	}
}

func TestGetStringData(t *testing.T) {
	input := "  hello world  "
	defer mockInput(t, input)()
    
	expected := "Hello World" 

	result := GetStringData("Prompt: ") 

	assert.Equal(t, expected, result, "GetStringData should trim and capitalize the string")
}

func TestGetIntData_ValidInput(t *testing.T) {
	input := "42"
	defer mockInput(t, input)()
	
	expected := 42
	result := GetIntData("Enter age: ")
	
	assert.Equal(t, expected, result, "GetIntData should correctly parse a valid integer")
}

func TestGetIntData_Spaces(t *testing.T) {
	input := "  101 "
	defer mockInput(t, input)()
	
	expected := 101
	result := GetIntData("Enter number: ")
	
	assert.Equal(t, expected, result, "GetIntData should handle leading/trailing spaces")
}

// TODO: add tests that produce errors and check that they return the correct error message