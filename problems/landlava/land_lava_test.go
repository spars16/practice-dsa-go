package landlava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// land = [0,1,2,1,2,0,1,3]
// lava_vol =  4
//               _
//     _   _    |
//   _| |_| |  _|
// _|       |_|
//
//
// land = [3,0,1,2,1,2,0,1,3]
// lava_vol =  14
// _               _
//  |    _   _    |
//  |  _| |_| |  _|
//  |_|       |_|

func TestLandLava(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "valid - original",
			input:    []int{0, 1, 2, 1, 2, 0, 1, 3},
			expected: 4,
		},
		{
			name:     "valid - far walls",
			input:    []int{3, 0, 1, 2, 1, 2, 0, 1, 3},
			expected: 14,
		},
		{
			name:     "valid - reverse original",
			input:    []int{3, 1, 0, 2, 1, 2, 1, 0},
			expected: 4,
		},
		{
			name:     "valid - empty",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, FillLava(tt.input))
	}
}
