package heap

import "testing"

func TestHeap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "happy path - push",
			input:    []int{10, 20, 30, 14, 52, 23, 63},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		h := MaxHeap{}
		for i := range tt.input {
			h.Push(tt.input[i])
		}
		for j := range h.array {
			if h.array[j] != tt.expected[j] {
				t.Errorf("at index %d: got %v, want %v", j, h.array[j], tt.expected[j])
			}
		}
	}
}
