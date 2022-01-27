package hashtable

import "fmt"

// Go's map[string]string implementation provides hashtable functionality by default
const ArraySize = 5

type (
	HashTable struct {
		arr  [ArraySize]*bucket
		size int
	}

	bucket struct {
		head *bucketNode
		size int
	}

	bucketNode struct {
		key  string
		next *bucketNode
	}
)

// Init
func Init() *HashTable {
	table := &HashTable{}
	for i := range table.arr {
		table.arr[i] = &bucket{
			head: &bucketNode{},
		}
	}
	return table
}

// Insert
func (h *HashTable) Insert(key string) {
	h.arr[hash(key)].insert(key)
	h.size++
}

// Search
func (h *HashTable) Search(key string) int {
	index := hash(key)
	node := h.arr[index].search(key)
	if node != nil {
		return index
	}
	return -1
}

// Delete
func (h *HashTable) Delete(key string) bool {
	index := hash(key)
	if h.arr[index].delete(key) {
		h.size--
		return true
	}
	return false
}

func (h *HashTable) Print() {
	for i, elem := range h.arr {
		fmt.Printf("Index %v:", i)
		elem.Print()
		fmt.Println()
	}
}

// insert
func (b *bucket) insert(key string) {
	// TODO: if bucket already contains key, do not insert
	next := b.head.next
	node := &bucketNode{key: key, next: next}
	b.head.next = node
	b.size++
}

// search
func (b *bucket) search(key string) *bucketNode {
	curr := b.head.next
	for curr != nil {
		if curr.key == key {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// delete
func (b *bucket) delete(key string) bool {
	curr := b.head
	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			b.size--
			return true
		}
	}
	return false
}

func (b *bucket) Print() {
	curr := b.head.next
	for curr != nil {
		fmt.Printf(" %v", curr.key)
		curr = curr.next
	}
}

// basic hash function
func hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}
