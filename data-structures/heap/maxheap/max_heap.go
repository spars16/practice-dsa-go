package maxheap

type MaxHeap struct {
	arr  []int
	size int
}

func (h *MaxHeap) Insert(k int) {
	h.arr = append(h.arr, k)
	h.size++
	h.maxHeapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) Extract() int {
	lastIndex := len(h.arr)
	if lastIndex < 0 {
		return -1
	}
	h.swap(0, lastIndex-1)
	removed := h.arr[lastIndex-1]
	h.arr = h.arr[:lastIndex-1]
	h.size--
	h.maxHeapifyDown(0)
	return removed
}

func (h *MaxHeap) GetMax() int {
	return h.arr[0]
}

// O(logn)
func (h *MaxHeap) maxHeapifyUp(index int) {
	if h.arr[index] > h.arr[parent(index)] {
		h.swap(index, parent(index))
		h.maxHeapifyUp(parent(index))
	}
}

// O(logn)
func (h *MaxHeap) maxHeapifyDown(index int) {
	if len(h.arr) <= index {
		return
	}

	if h.arr[left(index)] > h.arr[index] {
		h.swap(left(index), index)
		h.maxHeapifyDown(left(index))
	} else if h.arr[right(index)] > h.arr[index] {
		h.swap(right(index), h.arr[index])
		h.maxHeapifyDown(right(index))
	}
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.arr[index1], h.arr[index2] = h.arr[index2], h.arr[index1]
}

// the index of the parent node
func parent(index int) int {
	return (index - 1) / 2 // truncated integer division
}

// the index of the left node
func left(index int) int {
	return (index * 2) + 1
}

// the index of the right node
func right(index int) int {
	return (index * 2) + 2
}
