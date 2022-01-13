package trie

import (
	"testing"
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

	if trie.Search("wizard") != false {
		t.Error("expected false")
	}

	if trie.Search("dash") != true {
		t.Error("expected true")
	}
}
