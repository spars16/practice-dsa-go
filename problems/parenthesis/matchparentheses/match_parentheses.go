package matchparentheses

/*
Snap - 1st Round
Given a string that contains parentheses and letters, remove the minimum number of invalid parentheses to make the input string valid.

input: "()"
output: "()"

1 --> 0 --> -1
input: "(a))"
output: "(a)"

1 --> 2 --> 1
3
input: "((a)"
output: "(a)"

-1 --> 0
input: ")("
output: ""

input: "()))"
output: "()"

input: "()(()"
output: "()()"

input: "(a))()"
output: "(a)()"

3

// stack, when we see an ( , we push )
// when we see ) , pop off the stack to compare

// keep two counters for open and closed parentheses
// check the difference,
//. if it is greater than 0, remove open parenthesis
//. if it is less than 0, remove closed parenthesis

*/
import (
	"strings"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	removed := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return removed
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Time complexity: O(N)
// Space complexity: O(N)
func matchParentheses(str string) string {
	s := Stack{} // O(N) space
	strArray := strings.Split(str, "")

	result := make([]string, len(strArray)) // O(N) space
	copy(result, strArray)                  // O(N)

	for i, currStr := range strArray { // O(N)
		if currStr == "(" {
			s.Push(i)
		} else if currStr == ")" {
			if !s.IsEmpty() {
				s.Pop()
			} else {
				result = append(result[:i], result[i+1:]...)
			}
		}
	}

	for !s.IsEmpty() { // O(N)
		i := s.Pop()
		result = append(result[:i], result[i+1:]...)
	}

	return strings.Join(result, "")
}
