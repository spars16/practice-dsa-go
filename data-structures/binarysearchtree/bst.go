package binarysearchtree

type BinarySearchTreeNode struct {
	Data  int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

func (t *BinarySearchTreeNode) Insert(k int) {
	if t.Data < k {
		if t.Right == nil {
			t.Right = &BinarySearchTreeNode{Data: k}
			return
		}
		t.Right.Insert(k)
	} else if t.Data > k {
		if t.Left == nil {
			t.Left = &BinarySearchTreeNode{Data: k}
			return
		}
		t.Left.Insert(k)
	}
}

func (t *BinarySearchTreeNode) Search(k int) bool {
	if t == nil {
		return false
	}

	if t.Data < k {
		return t.Right.Search(k)
	} else if t.Data > k {
		return t.Left.Search(k)
	}
	return true
}

func (t *BinarySearchTreeNode) Remove(k int) bool {
	if t == nil {
		return false
	}

	if t.Data < k {
		if t.Right != nil && t.Right.Data == k {
			t.Right = nil
			return true
		}
		return t.Right.Remove(k)
	} else if t.Data > k {
		if t.Left != nil && t.Left.Data == k {
			t.Left = nil
			return true
		}
		return t.Left.Remove(k)
	}
	return false
}
