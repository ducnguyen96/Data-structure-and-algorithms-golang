package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int)  {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(400)
	tree.Insert(300)
	fmt.Println(tree)
	exist := tree.Search(120)
	fmt.Println("exist :", exist)
}