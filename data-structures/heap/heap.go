package heap

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Push(input int) int {
	return -1
}

func (h *MaxHeap) Pop() int {
	return -1
}

func (h *MaxHeap) maxHeapifyUp() {

}

func (h *MaxHeap) maxHeapifyDown() {

}

func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

// returns the index of the parent node
func parent(index int) int {
	return index - 1
}

// returns the index of the left node
func left(index int) int {
	return index*2 + 1
}

// returns the index of the right node
func right(index int) int {
	return index*2 + 2
}
