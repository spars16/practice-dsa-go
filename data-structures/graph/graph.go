package graph

import "fmt"

type (
	Graph struct {
		vertices []*node
	}

	node struct {
		key      int
		adjacent []*node
	}
)

func (g *Graph) AddVertex(k int) {
	if getNode(g.vertices, k) != nil {
		fmt.Printf("Vertex %v already exists.\n", k)
		return
	}
	n := &node{key: k}
	g.vertices = append(g.vertices, n)
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := getNode(g.vertices, from)
	toVertex := getNode(g.vertices, to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("Vertex does not exist: invalid edge %v --> %v\n", from, to)
	} else if getNode(fromVertex.adjacent, to) != nil {
		fmt.Printf("Edge %v --> %v already exists.\n", from, to)
	} else {
		// add edge
		// directed graph implementation, so we only add toVertex to fromVertex's adjacency list
		// rather than creating a bidirectional connection
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func getNode(arr []*node, k int) *node {
	for _, vertex := range arr {
		if vertex.key == k {
			return vertex
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("Vertex %d:", vertex.key)
		for _, edge := range vertex.adjacent {
			fmt.Printf(" %v", edge.key)
		}
		fmt.Println()
	}
}
