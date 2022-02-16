package intersectingintervals

import (
	"container/heap"
	"math"
	"sort"
)

/*
Meta - 1st round
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
find the minimum number of meeting rooms required.

*Example 1:Input:* [[0, 30],[5, 10],[15, 20]]*Output:* 2
*Example 2:Input:* [[7,10],[2,4]] *Output:* 1

*Example 3:Input:* [[0, 5],[5, 30],[4, 8]]
[[0, 5],[4, 8],[5, 30]]

-------------------
    ---
	    ---
*/

type MinHeap struct {
	items []int
}

// Sort interface: Len, Swap, Less
func (h MinHeap) Len() int {
	return len(h.items)
}

func (h MinHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h.items[i] < h.items[j]
}

// Push
func (h *MinHeap) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}

// Pop
func (h *MinHeap) Pop() interface{} {
	if len(h.items) <= 0 {
		return -1
	}
	removed := h.items[0]
	h.items = h.items[1:]
	return removed
}

// Peek
func (h *MinHeap) Peek() interface{} {
	if len(h.items) <= 0 {
		return -1
	}
	return h.items[0]
}

type Interval struct {
	startTime int
	endTime   int
}

func FindMinimumRooms(intervals []*Interval) int {
	byStartTime := sortIntervalsByStartTime(intervals)

	numRooms := 0
	h := &MinHeap{}
	heap.Init(h)
	h.Push(math.MaxInt)
	for _, interval := range byStartTime {
		min := h.Peek().(int)
		for min < interval.startTime {
			h.Pop()
		}
		numRooms += len(h.items) - 1
		h.Push(interval.endTime)
	}
	return numRooms
}

func sortIntervalsByStartTime(intervals []*Interval) []*Interval {
	result := make([]*Interval, len(intervals))
	copy(result, intervals)

	sort.Slice(result, func(i, j int) bool {
		return result[i].startTime < result[j].startTime
	})
	return result
}
