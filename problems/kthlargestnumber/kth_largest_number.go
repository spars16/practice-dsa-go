package kthlargestnumber

import (
	"fmt"
	"sort"
)

type MaxHeap struct {
	arr []int
}

func (m *MaxHeap) Insert(k int) {
	m.arr = append(m.arr, k)
	m.maxHeapifyUp(len(m.arr) - 1)
}

func (m *MaxHeap) Extract() int {
	lastIndex := len(m.arr)
	if lastIndex < 0 {
		return -1
	}
	m.swap(0, lastIndex-1)
	removed := m.arr[lastIndex-1]
	m.arr = m.arr[:lastIndex-1]
	m.maxHeapifyDown(0)
	return removed
}

func (m *MaxHeap) GetMax() int {
	return m.arr[0]
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	if m.arr[index] > m.arr[parent(index)] {
		m.swap(index, parent(index))
		m.maxHeapifyUp(parent(index))
	}
}

func (m *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(m.arr) - 1
	if index > lastIndex {
		return
	}

	left, right := left(index), right(index)
	childToCompare := 0
	if right > lastIndex || m.arr[left] > m.arr[right] {
		childToCompare = left
	} else {
		childToCompare = right
	}

	if m.arr[index] < m.arr[childToCompare] {
		m.swap(childToCompare, index)
		m.maxHeapifyDown(childToCompare)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index * 2) + 1
}

func right(index int) int {
	return (index * 2) + 2
}

func (m *MaxHeap) swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

// max heap approach: O(log(N))
func KthLargest(arr []int, k int) int {
	m := MaxHeap{}
	for i, k := range arr {
		m.Insert(k)
		fmt.Println(m.arr, i)
	}

	for i := 0; i < k-1; i++ {
		m.Extract()
	}
	return m.GetMax()
}

// sorting approach (merge or quick sort): O(Nlog(N))
func KthLargestSimple(arr []int, k int) int {
	sort.Ints(arr)
	return arr[len(arr)-k]
}
