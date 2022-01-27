package indexdifference

import (
	"fmt"
	"math"
)

// ByteDance - 1st Round

// Start typing here
/*
For each element in a given array, calculate the absolute value of index difference between it and all other elements of the same value.

Return the resulting values in an array. For example, if ther array elements at indices 2 and 3 are equal, the distance metric for element 2 is |2-3| = 1. For element 3 it is |3-2| = 1.

Example: n = 6, arr = [1,2,1,1,2,3]
        index =       [0,1,2,3,4,5]

1: 0 -> 2 -> 3
2: 1 -> 4
3: 5

answer = [5,3,3,4,3,0]

*/

// map[int]*Node
// Node: data, next

type (
	LinkedList struct {
		head *Node
	}

	Node struct {
		data int
		next *Node
	}
)

func (l *LinkedList) insert(val int) {
	// insert value at the beginning of the linked list
	head := l.head

	toInsert := &Node{
		data: val,
		next: head.next,
	}

	head.next = toInsert
}

func (l *LinkedList) calculate(index int) int {
	curr := l.head.next

	sum := 0
	for curr != nil {
		sum += int(math.Abs(float64(curr.data) - float64(index)))
		curr = curr.next
	}
	return sum
}

func (l *LinkedList) Println() {
	curr := l.head.next
	for curr != nil {
		fmt.Printf("%v ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func FindIndexDifference(input []int) []int {
	result := []int{}

	m := createInputMap(input)
	for i, elem := range input {
		sum := m[elem].calculate(i)
		result = append(result, sum)
	}

	return result
}

func createInputMap(input []int) map[int]*LinkedList {
	m := make(map[int]*LinkedList)
	for i := len(input) - 1; i >= 0; i-- {
		if _, contains := m[input[i]]; !contains {
			m[input[i]] = &LinkedList{
				head: &Node{
					next: &Node{
						data: i,
					},
				},
			}
		} else {
			m[input[i]].insert(i)
		}
	}
	return m
}
