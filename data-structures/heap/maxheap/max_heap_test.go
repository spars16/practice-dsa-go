package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		extract  int
		expected []int
	}{
		{
			name:     "happy path - push",
			input:    []int{10, 20, 30, 14, 52, 23, 63},
			extract:  0,
			expected: []int{63, 30, 52, 10, 14, 20, 23},
		},
		{
			name:     "happy path - extract",
			input:    []int{10, 20, 30, 14, 52, 23, 63},
			extract:  1,
			expected: []int{52, 30, 23, 10, 14, 20},
		},
	}

	for _, tt := range tests {
		h := MaxHeap{}
		for _, inp := range tt.input {
			h.Insert(inp)
		}

		for i := 0; i < tt.extract; i++ {
			h.Extract()
		}
		assert.Equal(t, tt.expected, h.arr)
	}
}
