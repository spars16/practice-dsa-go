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

func (list *ArrayList) Remove(value int) bool {
	for i, item := range list.items {
		if item == value {
			list.items = append(list.items[0:i], list.items[i+1:]...)
			list.size--
			return true
		}
	}
	return false
}
