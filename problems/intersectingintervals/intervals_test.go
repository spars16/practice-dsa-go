package intersectingintervals

import (
	"fmt"
	"testing"
)

func TestOverlappingIntervals(t *testing.T) {
	tests := []struct {
		name     string
		input    []*Interval
		expected int
	}{
		{
			name: "valid - long",
			input: []*Interval{
				{startTime: 0, endTime: 30},
				{startTime: 5, endTime: 10},
				{startTime: 15, endTime: 20},
			},
			expected: 2,
		},
		{
			name: "valid - no overlap",
			input: []*Interval{
				{startTime: 7, endTime: 10},
				{startTime: 2, endTime: 4},
			},
			expected: 1,
		},
		{
			name: "valid - short",
			input: []*Interval{
				{startTime: 0, endTime: 5},
				{startTime: 5, endTime: 30},
				{startTime: 4, endTime: 8},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		output := FindMinimumRooms(tt.input)
		if output == tt.expected {
			fmt.Printf("Passed %s\n", tt.name)
		} else {
			fmt.Printf("Failed %s: got %v, wanted %v\n", tt.name, output, tt.expected)
		}
	}
}
