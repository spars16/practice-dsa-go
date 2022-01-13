package parenthesis

type Stack struct {
	items []rune
}

func (s *Stack) push(b rune) {
	s.items = append(s.items, b)
}

func (s *Stack) pop() rune {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func IsValid(s string) bool {
	stack := Stack{}
	for _, c := range s {
		if c == '(' {
			stack.push(')')
		} else if c == '{' {
			stack.push('}')
		} else if c == '[' {
			stack.push(']')
		} else if stack.isEmpty() || stack.pop() != c {
			return false
		}
	}
	return stack.isEmpty()
}
