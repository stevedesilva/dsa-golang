package trie

import "errors"

type node struct {
	children map[rune]*node
}

func newNode() *node {
	return &node{children: make(map[rune]*node)}
}

type Trie struct {
	Root *node
}

func NewTrie() *Trie {
	return &Trie{
		Root: newNode(),
	}
}

func (t *Trie) Insert(word string) error {
	if len(word) == 0 {
		return errors.New("no word to input, word length is zero")
	}
	currNode := t.Root
	for _, c := range word {
		if _, found := currNode.children[c]; found {
			currNode = currNode.children[c]
		} else {
			currNode.children[c] = newNode()
		}
	}
	currNode.children['*'] = nil
	return nil

	// insert into a map code
}

func (t *Trie) Search(word string) []string {
	return nil
}
