package levelorder

import (
	"fmt"
	"math"
)

type (
	node struct {
		data  int
		left  *node
		right *node
	}
)

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

func (n *node) LevelOrderTraversal() {
	h := height(n)
	for i := 1; i <= int(h); i++ {
		n.printCurrentLevel(i)
		fmt.Println()
	}
}

func (n *node) printCurrentLevel(level int) {
	if level == 1 {
		fmt.Printf("%d ", n.data)
	} else if level > 1 {
		if n.left != nil {
			n.left.printCurrentLevel(level - 1)
		}
		if n.right != nil {
			n.right.printCurrentLevel(level - 1)
		}
	}
}

func height(n *node) float64 {
	if n == nil {
		return 0
	}
	lHeight := height(n.left)
	rHeight := height(n.right)

	return math.Max(lHeight, rHeight) + 1
}
