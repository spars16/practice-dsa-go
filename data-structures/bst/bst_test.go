package bst

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &BinarySearchTreeNode{Data: 100}
	input := []int{50, 150}
	for _, inp := range input {
		tree.Insert(inp)
	}
	fmt.Println(tree)

	if tree.Search(25) != false {
		t.Errorf("expected false")
	}

	if tree.Search(50) != true {
		t.Errorf("expected true")
	}
}
