package linkedlist

type (
	// Doubly Linked List struct implementation
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

func Init() *LinkedList {
	head := &node{}
	tail := &node{
		prev: head,
	}
	head.next = tail
	return &LinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (list *LinkedList) Prepend(data int) {
	second := list.head.next
	first := &node{
		data: data,
		next: second,
		prev: list.head,
	}
	list.head.next = first
	second.prev = first
	list.size++
}

func (list *LinkedList) Append(data int) {
	secondToLast := list.tail.prev
	last := &node{
		data: data,
		next: list.tail,
		prev: secondToLast,
	}
	list.tail.prev = last
	secondToLast.next = last
	list.size++
}

func (list *LinkedList) RemoveLast() *node {
	nodeToRemove := list.tail.prev
	list.tail.prev = nodeToRemove.prev
	nodeToRemove.prev.next = list.tail
	list.size--
	return nodeToRemove
}

func (list *LinkedList) RemoveFirst() *node {
	nodeToRemove := list.head.next
	list.head.next = nodeToRemove.next
	nodeToRemove.next.prev = list.head
	list.size--
	return nodeToRemove
}

// Remove the first occurrence of the passed in value
// RETURN the node that was removed
func (list *LinkedList) Remove(val int) *node {
	curr := list.head.next
	for curr != list.tail {
		if curr.data == val {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			list.size--
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (list *LinkedList) ToArray() []int {
	var arr []int
	curr := list.head.next
	for curr != list.tail {
		arr = append(arr, curr.data)
		curr = curr.next
	}
	return arr
}
