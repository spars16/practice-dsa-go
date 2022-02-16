package arraylist

// Go's slice implementation provides arraylist functionality by default
type ArrayList struct {
	items []int
	size  int
}

func (list *ArrayList) Add(value int) {
	list.items = append(list.items, value)
	list.size++
}

func (list *ArrayList) RemoveByIndex(index int) int {
	if index >= len(list.items) {
		return -1
	}
	removed := list.items[index]
	list.items = append(list.items[:index], list.items[index+1:]...)
	return removed
}

func (list *ArrayList) RemoveByValue(value int) bool {
	for i, item := range list.items {
		if item == value {
			list.items = append(list.items[0:i], list.items[i+1:]...)
			list.size--
			return true
		}
	}
	return false
}
