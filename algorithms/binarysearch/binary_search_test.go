package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{
			name:     "valid - simple sorted",
			input:    []int{2, 4, 6, 8, 25, 63, 73, 163, 174, 540},
			target:   8,
			expected: 3,
		},
		{
			name:     "invalid - simple sorted",
			input:    []int{2, 4, 6, 8, 25, 63, 73, 163, 174, 540},
			target:   5,
			expected: -1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, BinarySearch(tt.input, tt.target), tt.expected)
	}
}

func TestBinarySearchString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   rune
		expected int
	}{
		{
			name:     "valid - simple sorted",
			input:    "abgilxyz",
			target:   'y',
			expected: 6,
		},
		{
			name:     "valid - complex",
			input:    "jfksdaLJAf",
			target:   'L',
			expected: 6,
		},
		{
			name:     "invalid - simple sorted",
			input:    "abgilxyz",
			target:   'm',
			expected: -1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, BinarySearchString(tt.input, tt.target), tt.expected)
	}
}
