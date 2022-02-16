package arraylist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayList(t *testing.T) {
	tests := []struct {
		name          string
		add           []int
		removeByValue int
		expected      bool
	}{
		{
			name:          "valid - start",
			add:           []int{2, 4, 6, 5, 12},
			removeByValue: 2,
			expected:      true,
		},
		{
			name:          "valid - middle",
			add:           []int{2, 4, 6, 5, 12},
			removeByValue: 5,
			expected:      true,
		},
		{
			name:          "valid - end",
			add:           []int{2, 4, 6, 5, 12},
			removeByValue: 12,
			expected:      true,
		},
		{
			name:          "invalid",
			add:           []int{2, 4, 6, 5, 12},
			removeByValue: 3,
			expected:      false,
		},
	}

	for _, tt := range tests {
		list := ArrayList{}
		for _, item := range tt.add {
			list.Add(item)
		}
		fmt.Println(list.items)
		assert.Equal(t, list.RemoveByValue(tt.removeByValue), tt.expected)
		fmt.Println(list.items)

		fmt.Println(list.RemoveByIndex(0))
		fmt.Println(list.items)
	}
}
