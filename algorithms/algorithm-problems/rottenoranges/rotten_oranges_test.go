package rottenoranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRottenOranges(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "valid",
			input: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "invalid",
			input: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, RottenOranges(tt.input, 0, -1))
	}
}
