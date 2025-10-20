package cost

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetBaseRate(t *testing.T) {
    tests := []struct {
        name             string
        volume           float64 
        expectedBaseRate float64
    }{
        // Low Rate Tests
        {name: "Low Rate - Inside Range", volume: 1000.0, expectedBaseRate: 800.0},
        {name: "Low Rate - Upper Boundary (6000.0)", volume: 6000.0, expectedBaseRate: 800.0}, 
        
        // Medium Rate Tests
        {name: "Medium Rate - Lower Boundary (6000.001)", volume: 6000.001, expectedBaseRate: 1200.0}, 
        {name: "Medium Rate - Upper Boundary (100000.0)", volume: 100000.0, expectedBaseRate: 1200.0}, 
        
        // High Rate Tests
        {name: "High Rate - Lower Boundary (100000.001)", volume: 100000.001, expectedBaseRate: 1500.0}, 
        {name: "High Rate - Large Volume", volume: 500000.0, expectedBaseRate: 1500.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := GetBaseRate(tt.volume)
            if result != tt.expectedBaseRate {
                t.Errorf("GetBaseRate(%.3f) returned %.2f; want %.2f", 
                    tt.volume, result, tt.expectedBaseRate)
            }
        })
    }
}

func TestGetShippingRate(t *testing.T) {
    const fixedBaseRate = 1000.0

    tests := []struct {
        name             string
        island           string
        expectedTotalRate float64
    }{
        {name: "North Island Surcharge", island: "North Island", expectedTotalRate: 1100.0},
        {name: "South Island Surcharge", island: "South", expectedTotalRate: 1150.0},
        {name: "Stewart Island Surcharge", island: "Stewart Island", expectedTotalRate: 1200.0},
        
        {name: "Invalid/Other Island", island: "Chatham Islands", expectedTotalRate: 1000.0},
        {name: "Empty String", island: "", expectedTotalRate: 1000.0},
        
        {name: "Case Insensitive Check (North)", island: "nOrTh iSlAnD", expectedTotalRate: 1100.0},
        {name: "Case Insensitive Check (South)", island: "sOutH", expectedTotalRate: 1150.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := GetShippingRate(fixedBaseRate, tt.island)
            
            assert.InDelta(t, tt.expectedTotalRate, result, 0.001, "GetShippingRate calculation failed")
        })
    }
}

func TestIsValidIsland(t *testing.T) {
    tests := []struct {
        name  string
        input string
        expected bool
    }{
        // Valid Cases
        {name: "Valid - North Island full name", input: "North Island", expected: true},
        {name: "Valid - South Island lower case", input: "south island", expected: true},
        {name: "Valid - Stewart uppercase only", input: "STEWART", expected: true},
        
        // Invalid Cases
        {name: "Invalid - Blank", input: "", expected: false},
        {name: "Invalid - Other Island", input: "Auckland", expected: false},
        {name: "Invalid - Partial Match", input: "Nortwind", expected: false},
        {name: "Invalid - Just spaces", input: "   ", expected: false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsValidIsland(tt.input)
            
            assert.Equal(t, tt.expected, result, "IsValidIsland returned wrong status")
        })
    }
}