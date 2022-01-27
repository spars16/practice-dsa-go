package kthlargestnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name       string
		inputArray []int
		inputK     int
		expected   int
	}{
		{
			name:       "valid",
			inputArray: []int{1, 6, 2, 5, 3, 7, 4},
			inputK:     3,
			expected:   5,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, KthLargestSimple(tt.inputArray, tt.inputK))
		// assert.Equal(t, tt.expected, KthLargest(tt.inputArray, tt.inputK))
	}
}
