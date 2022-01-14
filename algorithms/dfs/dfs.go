package dfs

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

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

func (n *Node) DepthFirstSearch(array []int) []int {
	if n == nil {
		return array
	}
	array = append(array, n.Value)
	array = n.Left.DepthFirstSearch(array)
	array = n.Right.DepthFirstSearch(array)
	return array
}
