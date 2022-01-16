package graph

import "testing"

func TestDirectedGraph(t *testing.T) {
	tests := []struct {
		name     string
		vertices []int
		edges    [][2]int
	}{
		{
			name:     "valid",
			vertices: []int{2, 6, 3, 7, 1, 6},
			edges: [][2]int{
				{2, 6},
				{7, 3},
				{4, 2}, // invalid, vertex 4 does not exist
				{2, 5}, // invalid, vertex 5 does not exist
				{2, 6}, // invalid, edge already exists
				{3, 1},
			},
		},
	}

	for _, tt := range tests {
		graph := Graph{}
		for _, vertex := range tt.vertices {
			graph.AddVertex(vertex)
		}
		for _, edge := range tt.edges {
			graph.AddEdge(edge[0], edge[1])
		}
		graph.Print()
	}
}
