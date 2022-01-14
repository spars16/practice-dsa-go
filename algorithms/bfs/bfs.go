package bfs

import (
	"fmt"
	"math"
)

type (
	Node struct {
		Value int
		Left  *Node
		Right *Node
	}
)

func (n *Node) Insert(val int) {
	if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
		}
		n.Left.Insert(val)
	} else if val > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: val}
		}
		n.Right.Insert(val)
	}
}

// INCOMPLETE
// func (n *Node) BreadthFirstSearch(array []int) []int {
// 	queue := []*Node{n}
// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue := queue[1:]
// 		array = append(array, current.Value)
// 		if n.Left != nil {
// 			queue = append(queue, n.Left)
// 		}
// 		if n.Right != nil {
// 			queue = append(queue, n.Right)
// 		}
// 	}
// 	return array
// }

func (n *Node) LevelOrderTraversal() {
	h := height(n)
	for i := 1; i < int(h); i++ {
		n.printCurrentLevel(i)
	}
}

func height(n *Node) float64 {
	if n == nil {
		return 0
	}
	lHeight := height(n.Left)
	rHeight := height(n.Right)

	return math.Max(lHeight, rHeight) + 1
}

func (n *Node) printCurrentLevel(level int) {
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
