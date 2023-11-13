package trie

import (
	"errors"
	"fmt"
)

type Node struct {
	children map[rune]*Node
}

func newNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

type Trie struct {
	Root *Node
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
		if currNode.children[c] != nil {
			currNode = currNode.children[c]
		} else {
			node := newNode()
			currNode.children[c] = node
			currNode = node
		}
	}
	currNode.children['*'] = nil
	return nil

	// insert into a map code
}

func (t *Trie) Search(word string) (*Node, error) {
	currNode := t.Root
	if len(word) == 0 {
		return currNode, errors.New("no word to input, word length is zero")
	}
	for _, c := range word {
		currentRune := fmt.Sprintf("current rune is: %c", c)
		fmt.Println(currentRune)
		if v, found := currNode.children[c]; found {
			currentRune := fmt.Sprintf("%v, current rune is: %c", v, c)
			fmt.Println(currentRune)
			currNode = currNode.children[c]
		} else {
			r := c
			// rune not found Sprintf message
			msg := fmt.Sprintf("rune %c not found", r)
			return nil, errors.New(msg)
		}
	}
	return currNode, nil
}

func (t *Trie) PrintAllWords() ([]string, error) {
	// if trie is empty return error
	if t.Root == nil {
		return nil, errors.New("trie is empty")
	}
	word := ""
	words := make([]string, 0)
	printAllWords(t.Root, word, words)
	return words, nil
}

func printAllWords(root *Node, word string, words []string) {
	//
}
