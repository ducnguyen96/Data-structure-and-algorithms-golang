package main

import (
	"fmt"
)

const AlphabetLength = 26

// Node represents a node in trie that hold children
type Node struct {
	children [AlphabetLength]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

// Init a trie
func initTrie() *Trie {
	return &Trie{&Node{}}
}

// Insert a word to trie
func (t *Trie) insert(w string) {
	currentNode := t.root
	for _, chr := range w {
		// if chr node is not existed yet, create new node and assign it to current node
		// because we want children start with 'a' has 0 as index,
		// so we have to subtract chr with 'a'
		index := chr - 'a'
		if currentNode.children[index] == nil {
			currentNode.children[index] = &Node{}
		}

		// update currentNode
		currentNode = currentNode.children[index]
	}
	// now currentNode is last char we can assign isEnd to true
	currentNode.isEnd = true
}

// Search
func (t *Trie) search(w string) bool {
	currentNode := t.root
	for _, chr := range w {
		// if chr node is not existed yet, create new node and assign it to current node
		// because we want children start with 'a' has 0 as index,
		// so we have to subtract chr with 'a'
		index := chr - 'a'
		if currentNode.children[index] == nil {
			return false
		}

		// update currentNode
		currentNode = currentNode.children[index]
	}

	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := initTrie()
	fmt.Println(myTrie.root)

	words := [...]string{"hello", "world", "welcome", "to", "my", "new", "world"}

	for _, word := range words {
		myTrie.insert(word)
	}

	fmt.Println(myTrie.search("to"))

	fmt.Println(myTrie.search("welcome"))
	fmt.Println('w' - 'a') // index of letter 'w'
	// loop through children of 'w' we should have
	// node of 'e' in 'welcome' not nil and node of 'o' in 'world'
	for index, child := range myTrie.root.children['w' - 'a'].children {
		if child != nil {
			fmt.Printf("%c ", rune(index + 'a')) // 'e' and 'o'
		}
	}
}