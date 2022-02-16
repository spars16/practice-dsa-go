package slidematrixgame

import (
	"fmt"
	"testing"
)

func TestSlideMatrix(t *testing.T) {
	tests := []struct {
		name        string
		inputMatrix [][]*Coordinate
		inputStart  *Coordinate
		inputEnd    *Coordinate
		expected    bool
	}{
		{
			name: "valid",
			inputMatrix: [][]*Coordinate{
				{&Coordinate{value: 0, x: 0, y: 0}, &Coordinate{value: 0, x: 1, y: 0}, &Coordinate{value: 0, x: 2, y: 0}, &Coordinate{value: 1, x: 3, y: 0}},
				{&Coordinate{value: 0, x: 0, y: 1}, &Coordinate{value: 0, x: 1, y: 1}, &Coordinate{value: 0, x: 2, y: 1}, &Coordinate{value: 1, x: 3, y: 1}},
				{&Coordinate{value: 0, x: 0, y: 2}, &Coordinate{value: 0, x: 1, y: 2}, &Coordinate{value: 1, x: 2, y: 2}, &Coordinate{value: 0, x: 3, y: 2}},
				{&Coordinate{value: 1, x: 0, y: 3}, &Coordinate{value: 0, x: 1, y: 3}, &Coordinate{value: 0, x: 2, y: 3}, &Coordinate{value: 0, x: 3, y: 3}},
			},
			inputStart: &Coordinate{x: 0, y: 0},
			inputEnd:   &Coordinate{x: 3, y: 3},
			expected:   true,
		},
	}

	for _, tt := range tests {
		// output := FindPath(tt.inputMatrix, tt.inputStart, tt.inputEnd)
		// if output == tt.expected {
		//     fmt.Println("Passed")
		// } else {
		//     fmt.Printf("Failed %v", tt.name)
		// }
		fmt.Println(checkDirection(tt.inputMatrix, tt.inputStart, DOWN).y)
	}
}
