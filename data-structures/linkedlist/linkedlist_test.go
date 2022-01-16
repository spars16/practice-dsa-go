package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		name         string
		prepend      []int
		append       []int
		remove       []int
		expectedArr  []int
		expectedSize int
	}{
		{
			name:         "happy path - prepend",
			prepend:      []int{1, 2, 3, 4},
			expectedArr:  []int{4, 3, 2, 1},
			expectedSize: 4,
		},
		{
			name:         "happy path - append",
			append:       []int{1, 2, 2, 4, 3},
			expectedArr:  []int{1, 2, 2, 4, 3},
			expectedSize: 5,
		},
		{
			name:         "happy path - mixed",
			prepend:      []int{5, 6, 2},
			append:       []int{1, 2, 3, 4},
			remove:       []int{2, 5},
			expectedArr:  []int{6, 1, 2, 3, 4},
			expectedSize: 5,
		},
	}

	for _, tt := range tests {
		list := Init()
		for _, p := range tt.prepend {
			list.Prepend(p)
		}
		for _, a := range tt.append {
			list.Append(a)
		}
		for _, r := range tt.remove {
			list.Remove(r)
		}
		assert.Equal(t, tt.expectedSize, list.size)
		assert.Equal(t, tt.expectedArr, list.ToArray())
	}
}
