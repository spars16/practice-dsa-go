package calculator

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "valid",
			input:    "(1+(4+5+2)-3)+(6+8)",
			expected: 23,
		},
	}

	for _, tt := range tests {
		// assert.Equal(t, tt.expected, Calculator(tt.input))
		fmt.Println(Calculator(tt.input))
	}
}
