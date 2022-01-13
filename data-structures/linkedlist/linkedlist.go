package linkedlist

type (
	LinkedList struct {
		head *node
		tail *node
		size int
	}

	node struct {
		data int
		next *node
		prev *node
	}
)

func (list *LinkedList) Prepend(data int) {
	second := list.head
	list.head = &node{
		data: data,
		next: second,
	}
	list.size++
}

func (list *LinkedList) RemoveLast() *node {
	nodeToRemove := list.tail.prev
	list.tail.prev = nodeToRemove.prev
	nodeToRemove.prev.next = list.tail
	list.size--
	return nodeToRemove
}
