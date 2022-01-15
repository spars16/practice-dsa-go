package binarysearchtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &BinarySearchTreeNode{Data: 100}
	input := []int{50, 150}
	for _, inp := range input {
		tree.Insert(inp)
	}
	fmt.Println(tree)

	// search
	assert.False(t, tree.Search(25))
	assert.True(t, tree.Search(50))

	// remove
	assert.True(t, tree.Remove(50))
	assert.False(t, tree.Remove(25))
}
