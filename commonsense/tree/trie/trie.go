package trie

type node struct {
	children map[rune]*node
}

type Trie struct {
	Root *node
}

func NewTrie() *Trie {
	return &Trie{
		Root: &node{children: make(map[rune]*node)},
	}
}

func (t *Trie) Insert(word string) {

}

func (t *Trie) Search(word string) []string {
	return nil
}
