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

func (t *Trie) CollectAllWords() ([]string, error) {
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

//func (t *Trie) AutoCompleteWord(word string) ([]string, error) {
//	// if trie is empty return error
//	if t.Root == nil {
//		return nil, errors.New("trie is empty")
//	}
//	words := collectAllWords(t.Root, "", make([]string, 0))
//	return words, nil
//}

func (t *Trie) CollectAllKeys() ([]rune, error) {
	// if trie is empty return error
	if t.Root == nil {
		return nil, errors.New("trie is empty")
	}
	res := collectAllKeys(make([]rune, 0), t.Root)
	return res, nil
}

func (t *Trie) AutoComplete(prefix string) ([]string, error) {
	if len(prefix) < 1 {
		return nil, errors.New("word length less than 1")
	}
	// for each char find word
	if word, err := t.Search(prefix); err != nil {
		// prefix not found
		return nil, err
	} else {
		return collectAllWords(word, prefix, make([]string, 0)), nil
	}
}

/*
public String autoCorrect(String prefix) throws IllegalArgumentException {
        if (prefix == null || prefix.length() < 1) {
            throw new IllegalArgumentException();
        }
        final ArrayList<String> words = new ArrayList<>();

        final char[] prefixAsArrayOfChar = prefix.toCharArray();
        if (!root.getChildren().containsKey(prefixAsArrayOfChar[0])) {
            throw new IllegalArgumentException("Prefix not in trie");
        }

        String wordToFind = "";
        Node current = root;
        for (Character prefixChar : prefixAsArrayOfChar) {
            final HashMap<Character, Node> children = current.getChildren();
            if (children.containsKey(prefixChar)) {
                wordToFind += prefixChar;
                current = children.get(prefixChar);
            } else {
                printAll(current, wordToFind, words);
                if (words.isEmpty()) {
                    throw new IllegalArgumentException();
                } else {
                    return words.get(0);
                }
            }
        }
        return wordToFind;
    }
*/

func (t *Trie) AutoCorrect(word string) (string, error) {
	if len(word) < 1 {
		return "", errors.New("word length less than 1")
	}
	// match as much of word as possible
	curr := t.Root
	found := ""
	for _, v := range word {
		children := curr.children
		// if v in map
		if _, ok := children[v]; ok {
			found += string(v)
			curr = children[v]
		} else {
			// collect words for prefix
			node := curr
			if node == nil {
				return "", errors.New("word not found")
			}
			words := collectAllWords(node, found, make([]string, 0))
			if len(words) > 0 {
				return words[0], nil
			} else {
				return "", errors.New("word not found")
			}
		}
	}

	return found, nil
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
