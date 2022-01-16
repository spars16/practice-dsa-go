package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
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
	}

	for _, tt := range tests {
		h := MaxHeap{}
		for _, inp := range tt.input {
			h.Insert(inp)
		}
		assert.Equal(t, tt.expected, h.arr)
	}
}
