package traversal

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) insert(val int) {
	if val < n.data {
		if n.left == nil {
			n.left = &node{data: val}
		}
		n.left.insert(val)
	} else if val > n.data {
		if n.right == nil {
			n.right = &node{data: val}
		}
		n.right.insert(val)
	}
}

func InOrderTraversal(node *node) {
	// iterative
}

func InOrderTraversalRecursive(node *node) {
	if node != nil {
		InOrderTraversalRecursive(node.left)
		fmt.Println(node.data)
		InOrderTraversalRecursive(node.right)
	}
}
