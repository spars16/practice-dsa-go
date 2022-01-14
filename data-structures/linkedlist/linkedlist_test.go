package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedSize  int
		expectedValue int
	}{
		{
			name:          "happy path - prepend",
			input:         []int{1, 2, 3, 4},
			expectedSize:  4,
			expectedValue: 4,
		},
	}

	for _, tt := range tests {
		list := Init()
		for _, input := range tt.input {
			list.Prepend(input)
		}
		if tt.expectedSize != list.size {
			t.Errorf("got %d, want %d", list.size, tt.expectedSize)
		}
		if tt.expectedValue != list.head.next.data {
			t.Errorf("got %d, want %d", list.head.next.data, tt.expectedValue)
		}
	}
}
