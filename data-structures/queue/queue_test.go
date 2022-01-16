package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		dequeue  int
		expected []int
	}{
		{
			name:     "valid",
			input:    []int{2, 6, 1, 5, 3, 1, 7},
			dequeue:  0,
			expected: []int{2, 6, 1, 5, 3, 1, 7},
		},
		{
			name:     "valid - multiple dequeue",
			input:    []int{2, 6, 1, 5, 3, 1, 7},
			dequeue:  2,
			expected: []int{1, 5, 3, 1, 7},
		},
	}

	for _, tt := range tests {
		q := Queue{}
		for _, item := range tt.input {
			q.Enqueue(item)
		}
		for i := 0; i < tt.dequeue; i++ {
			q.Dequeue()
		}
		assert.Equal(t, tt.expected, q.items)
	}
}
