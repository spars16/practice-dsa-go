package indexdifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexDifference(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "valid",
			input:    []int{1, 2, 1, 1, 2, 3},
			expected: []int{5, 3, 3, 4, 3, 0},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, FindIndexDifference(tt.input))
	}
}
