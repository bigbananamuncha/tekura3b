package dimensions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetVolume(t *testing.T) {
	tests := []struct {
		name    string
		height  float64
		width   float64
		depth   float64
		expected float64
	}{
		{
			name: "Standard Box",
			height: 10.0,
			width: 5.0,
			depth: 2.0,
			expected: 100.0,
		},
		{
			name: "Zero Dimension Edge Case",
			height: 10.0,
			width: 0.0, 
			depth: 5.0,
			expected: 0.0,
		},
		{
			name: "Maximum Valid Volume",
			height: 100.0, 
			width: 100.0,
			depth: 100.0,
			expected: 1000000.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetVolume(tt.height, tt.width, tt.depth) 
			assert.InDelta(t, tt.expected, result, 0.001, "GetVolume result mismatch")
		})
	}
}