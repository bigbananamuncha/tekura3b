package cost

import (
	"testing"
)

func TestGetBaseRate(t *testing.T) {
	volume := 1000
	expectedBaseRate := 800

	result := GetBaseRate(volume)
	if result != expectedBaseRate {
		t.Errorf("GetBaseRate(%d) = %d; want %d", volume, result, expectedBaseRate)
	}

	volume = 6000
	expectedBaseRate = 800

	result = GetBaseRate(volume)
	if result != expectedBaseRate {
		t.Errorf("GetBaseRate(%d) = %d; want %d", volume, result, expectedBaseRate)
	}
}


// from the root of the project, run:
// go test ./...
// this will run all the tests in the project
// All the test files should be named *_test.go
// The test functions should be named Test*