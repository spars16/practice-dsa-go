package levelorder

import (
	"fmt"
	"math"
)

type (
	node struct {
		Value int
		Left  *node
		Right *node
	}
)

func (n *node) LevelOrderTraversal() {
	h := height(n)
	for i := 1; i < int(h); i++ {
		n.printCurrentLevel(i)
		fmt.Println()
	}
}

func (n *node) insert(val int) {
	if val < n.Value {
		if n.Left == nil {
			n.Left = &node{Value: val}
		}
		n.Left.insert(val)
	} else if val > n.Value {
		if n.Right == nil {
			n.Right = &node{Value: val}
		}
		n.Right.insert(val)
	}
}

func (n *node) printCurrentLevel(level int) {
	if level == 1 {
		fmt.Printf("%d ", n.Value)
	} else if level > 1 {
		if n.Left != nil {
			n.Left.printCurrentLevel(level - 1)
		}
		if n.Right != nil {
			n.Right.printCurrentLevel(level - 1)
		}
	}
}

func height(n *node) float64 {
	if n == nil {
		return 0
	}
	lHeight := height(n.Left)
	rHeight := height(n.Right)

	return math.Max(lHeight, rHeight) + 1
}
