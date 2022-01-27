package queue

type Queue struct {
	items []int
	size  int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	q.size++
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
