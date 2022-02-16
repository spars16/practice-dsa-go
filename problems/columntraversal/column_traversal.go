package columntraversal

/*
Meta - 1st round
*Input:* [3,9,8,4,0,1,7,null,null,null,2,5] (0's right child is 2 and 1's left child is 5)

     3
    /\
   /  \
   9   8
  /\  /\
 /  \/  \
 4  01   7
    /\
   /  \
   5   2

*Output:*

[
  [4],
  [9,5],
  [3,0,1],
  [8,2],
  [7]
]

[
  [3]
]
*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

func ColumnOrderTraversal(root *Node) [][]int {
	return helper(root, 0, [][]int{})
}

func helper(node *Node, currIndex int, result [][]int) [][]int {
	if node != nil {
		helper(node.left, currIndex-1, result)
		helper(node.right, currIndex+1, result)

		if currIndex < 0 {
			result = append([][]int{{node.data}}, result...)
		} else if currIndex > 0 {
			result = append(result, []int{node.data})
		}
	}
	return result
}
