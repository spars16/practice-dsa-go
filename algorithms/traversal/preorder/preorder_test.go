package traversal

import "testing"

func TestPreorderTraversal(t *testing.T) {

}

func TestPreorderTraversalRecursive(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "happy path",
			input: []int{150, 50, 200, 70, 55, 25},
		},
	}

	for _, tt := range tests {
		n := &node{data: 100}
		for _, val := range tt.input {
			n.insert(val)
		}
		PreOrderTraversalRecursive(n)
	}
}
