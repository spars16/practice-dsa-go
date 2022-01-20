package minheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
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
			expected: []int{10, 14, 23, 20, 52, 30, 63},
		},
	}

	for _, tt := range tests {
		h := MinHeap{}
		for _, inp := range tt.input {
			h.Insert(inp)
		}
		assert.Equal(t, tt.expected, h.arr)
	}
}
