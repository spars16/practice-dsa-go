package calculator

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(str string) {
	s.items = append(s.items, str)
}

func (s *Stack) Pop() string {
	removed := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return removed
}

func Calculator(input string) int {
	strArr := strings.Split(input, "")

	for _, str := range strArr {
		if str != "(" && str != ")" {
			// if str == "+" || str == ""
			fmt.Println()
		}
	}

	// convert strArr into postfix notation: A + B * C --> A B C * +
	// by storing operands in one stack and operators in another stack
	// and popping values from both in sequence

	return -1
}
