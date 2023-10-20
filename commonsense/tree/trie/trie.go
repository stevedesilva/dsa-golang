package trie

type node struct {
	children map[rune]*node
}

type Trie struct {
	Root *node
}

func NewTrie() *Trie {
	return &Trie{Root: &node{children: make(map[rune]*node)}}
}

// insert a word into the trie
func (t *Trie) Insert(word string) {
	// start at the root
}
