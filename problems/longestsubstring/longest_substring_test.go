package longestsubstring

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "valid",
			input:    "ABCDEFGABEF",
			expected: 7,
		},
		{
			name:     "valid - end",
			input:    "ABCABCD",
			expected: 4,
		},
		{
			name:     "valid - start",
			input:    "ABCDBCD", // ABCD
			expected: 4,
		},
	}

	for _, tt := range tests {
		output := FindLongestSubstring(tt.input)
		if output != tt.expected {
			fmt.Printf("Failed %v: want %v got %v\n", tt.name, tt.expected, output)
		} else {
			fmt.Printf("Passed %v\n", tt.name)
		}
	}
}
