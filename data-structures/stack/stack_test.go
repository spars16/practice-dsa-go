package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	input := []int{1, 2, 3}

	for _, i := range input {
		s.Push(i)
	}
	fmt.Println(s.items)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
