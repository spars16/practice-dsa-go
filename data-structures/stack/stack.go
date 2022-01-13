package stack

type Stack struct {
	items []int
	size  int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	s.size++
}

func (s *Stack) Pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.size--
	return item
}
