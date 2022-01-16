package levelorder

import (
	"fmt"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	n := node{Value: 100}
	n.insert(150)
	n.insert(50)
	n.insert(25)
	n.insert(75)
	n.insert(30)
	n.insert(80)
	fmt.Println(n)

	n.LevelOrderTraversal()
}
