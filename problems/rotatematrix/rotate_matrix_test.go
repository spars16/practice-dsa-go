package rotatematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "valid 3x3",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{4, 1, 2},
				{7, 5, 3},
				{8, 9, 6},
			},
		},
		{
			name: "valid 4x3",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: [][]int{
				{5, 1, 2, 3},
				{9, 6, 7, 4},
				{10, 11, 12, 8},
			},
		},
		{
			name: "valid 4x4",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: [][]int{
				{5, 1, 2, 3},
				{9, 10, 6, 4},
				{13, 11, 7, 8},
				{14, 15, 16, 12},
			},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, RotateMatrix(tt.input))
	}
}
