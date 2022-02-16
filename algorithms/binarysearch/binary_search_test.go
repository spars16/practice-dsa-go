package binarysearch

import (
	"sort"
	"strings"
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
		assert.Equal(t, tt.expected, BinarySearch(tt.input, tt.target))
		assert.Equal(t, tt.expected, BinarySearchRecursive(tt.input, tt.target, 0, len(tt.input)-1))
	}
}

func TestBinarySearchSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{
			name:     "valid - simple sorted",
			input:    []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55},
			target:   6,
			expected: 2,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, BinarySearchSort(tt.input, tt.target))
	}
}

func TestBinarySearchString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   byte
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
			expected: 2,
		},
		{
			name:     "invalid - simple sorted",
			input:    "abgilxyz",
			target:   'm',
			expected: -1,
		},
	}

	for _, tt := range tests {
		inp := strings.Split(tt.input, "")
		sort.Strings(inp)
		assert.Equal(t, tt.expected, BinarySearchString(strings.Join(inp, ""), tt.target))
	}
}
