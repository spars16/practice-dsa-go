package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name       string
		inputArray []int
		target     int
		expected   []int
	}{
		{
			name:       "valid",
			inputArray: []int{9, 2, 5, 7, 11, 3},
			target:     7,
			expected:   []int{1, 2},
		},
		{
			name:       "invalid",
			inputArray: []int{9, 2, 5, 7, 11, 3},
			target:     23,
			expected:   []int{},
		},
	}
	for _, tt := range tests {
		expected := TwoSum(tt.inputArray, tt.target)
		assert.Equal(t, expected, tt.expected)
	}
}
