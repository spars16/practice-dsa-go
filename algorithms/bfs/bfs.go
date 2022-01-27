package bfs

const (
	Unvisited State = iota
	Visited
	Visiting
)

type (
	Graph struct {
		nodes []*node
	}

	State int
	node  struct {
		key      int
		adjacent []*node
		state    State
	}

	Queue struct {
		items []*node
	}
)

func (g *Graph) AddVertex(k int) {
	// TODO: check if node with key k already exists in graph
	g.nodes = append(g.nodes, &node{key: k})
}

func (g *Graph) AddEdge(from, to int) {
	fromNode := getNode(g.nodes, from)
	toNode := getNode(g.nodes, to)

	// TODO: check if either nodes returned nil
	// TODO: check if edge already exists

	fromNode.adjacent = append(fromNode.adjacent, toNode)
}

func getNode(nodes []*node, k int) *node {
	for _, n := range nodes {
		if n.key == k {
			return n
		}
	}
	return nil
}

func (q *Queue) Enqueue(n *node) {
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() *node {
	if q.IsEmpty() {
		return nil
	}
	removed := q.items[0]
	q.items = q.items[1:]
	return removed
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (g *Graph) BreadthFirstSearch(start, end *node) bool {
	if start == end {
		return true
	}

	q := Queue{}

	start.state = Visiting
	q.Enqueue(start)
	for !q.IsEmpty() {
		n := q.Dequeue()
		if n != nil {
			for _, v := range n.adjacent {
				if v.state == Unvisited {
					if v == end {
						return true
					} else {
						v.state = Visiting
						q.Enqueue(v)
					}
				}
			}
			n.state = Visited
		}
	}
	return false
}

// type bstNode struct {
// 	data  int
// 	left  *bstNode
// 	right *bstNode
// }

// func (n *bstNode) BreadthFirstSearch() []int {
// 	queue := []*bstNode{n}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue := queue[1:]
// 		if n.left != nil {
// 			queue = append(queue, n.left)
// 		}
// 		if n.right != nil {
// 			queue = append(queue, n.right)
// 		}
// 	}
// 	return array
// }
