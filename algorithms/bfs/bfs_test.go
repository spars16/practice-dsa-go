package bfs

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	node := Node{Value: 100}
	node.Insert(150)
	node.Insert(50)
	node.Insert(25)
	node.Insert(75)
	node.Insert(30)
	node.Insert(80)
	fmt.Println(node)

	node.LevelOrderTraversal()
}
