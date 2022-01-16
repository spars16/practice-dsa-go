package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	input := []string{
		"cat",
		"cog",
		"dog",
		"dash",
		"do",
		"cater",
	}

	trie := &Trie{root: &Node{}}
	for _, inp := range input {
		trie.Insert(inp)
	}
	// fmt.Println(trie.root)

	assert.False(t, trie.Search("wizard"))
	assert.True(t, trie.Search("dash"))
}
