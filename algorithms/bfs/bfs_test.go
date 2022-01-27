package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearchGraph(t *testing.T) {
	tests := []struct {
		name     string
		input    [2]int // start, end
		expected bool
	}{
		// {
		// 	name:     "valid - connection",
		// 	input:    [2]int{125, 253},
		// 	expected: true,
		// },
		{
			name:     "valid - no connection",
			input:    [2]int{125, 253},
			expected: false,
		},
	}

	for _, tt := range tests {
		g := Graph{}

		// add nodes
		start := tt.input[0]
		end := tt.input[1]

		g.AddVertex(start)
		g.AddVertex(end)

		startNode := getNode(g.nodes, start)
		endNode := getNode(g.nodes, end)

		assert.Equal(t, tt.expected, g.BreadthFirstSearch(startNode, endNode))
	}
}

// func TestBreadthFirstSearch(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    []int
// 		expected []int
// 	}{
// 		{
// 			name:     "happy path",
// 			input:    []int{},
// 			expected: []int{},
// 		},
// 	}

// 	for _, tt := range tests {
// 		n := &node{data: 10}
// 		var result []int
// 		n.BreadthFirstSearch(result)
// 		assert.Equal(t, tt.expected, result)
// 	}
// }
