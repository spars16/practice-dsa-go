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

func PostOrderTraversal(node *node) {
	// iterative
}

func PostOrderTraversalRecursive(node *node) {
	if node != nil {
		PostOrderTraversalRecursive(node.left)
		PostOrderTraversalRecursive(node.right)
		fmt.Println(node.data)
	}
}
