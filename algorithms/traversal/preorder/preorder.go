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

func PreOrderTraversal(node *node) {
	// iterative
}

func PreOrderTraversalRecursive(node *node) {
	if node != nil {
		fmt.Println(node.data)
		PreOrderTraversalRecursive(node.left)
		PreOrderTraversalRecursive(node.right)
	}
}
