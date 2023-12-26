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
			return nil, errors.New(fmt.Sprintf("rune %c not found", c))
		}
	}
	return currNode, nil
}

func (t *Trie) PrintAllWords() ([]string, error) {
	// if trie is empty return error
	if t.Root == nil {
		return nil, errors.New("trie is empty")
	}
	words := collectAllWords(t.Root, "", make([]string, 0))
	return words, nil
}

func collectAllWords(node *Node, word string, words []string) []string {
	// loop over node children,
	// if '*' then add current work to array and return
	// else add key to word and call key node value
	for k, child := range node.children {
		if k == '*' {
			words = append(words, word)
		} else {
			currentWordBeingBuilt := word + string(k)
			words = collectAllWords(child, currentWordBeingBuilt, words)
		}
	}
	return words
}

func (t *Trie) AutoCompleteWord(word string) ([]string, error) {
	// if trie is empty return error
	if t.Root == nil {
		return nil, errors.New("trie is empty")
	}
	words := collectAllWords(t.Root, "", make([]string, 0))
	return words, nil
}

func (t *Trie) CollectAllKeys() ([]rune, error) {
	// if trie is empty return error
	if t.Root == nil {
		return nil, errors.New("trie is empty")
	}
	res := collectAllKeys(make([]rune, 0), t.Root)
	return res, nil
}

func collectAllKeys(res []rune, current *Node) []rune {
	for k, node := range current.children {
		res = append(res, k)
		if k != '*' {
			current = node
			res = collectAllKeys(res, current)
		}
	}
	return res
}
