package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "happy path - prepend",
			input:    []int{1, 2, 3, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		list := LinkedList{}
		for _, input := range tt.input {
			list.Prepend(input)
		}
		if tt.expected != list.head.data {
			t.Errorf("got %d, want %d", tt.expected, list.head.data)
		}
	}
}
