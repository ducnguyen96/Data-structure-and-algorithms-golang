package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
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

func (n *Node) BinarySearch(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.BinarySearch(k)
	} else if n.Key > k {
		return n.Left.BinarySearch(k)
	}
	return true
}

func (n *Node) BreadthFirstSearch() []int {
	currentNode := n
	var list []int
	var queue []*Node
	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		list = append(list, currentNode.Key)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	return list
}

func (n *Node) BreadthFirstSearchRecursive(queue []*Node, list []int) []int {
	if len(queue) == 0 {
		return list
	}

	currentNode := queue[0]
	queue = queue[1:]
	list = append(list, currentNode.Key)

	if currentNode.Left != nil {
		queue = append(queue, currentNode.Left)
	}
	if currentNode.Right != nil {
		queue = append(queue, currentNode.Right)
	}

	return currentNode.BreadthFirstSearchRecursive(queue, list)
}

func (n *Node) DepthFirstSearchInOrder() []int {
	return traverseInOrder(n, []int{})
}

func (n *Node) DepthFirstSearchPreOrder() []int {
	return traversePreOrder(n, []int{})
}

func (n *Node) DepthFirstSearchPostOrder() []int {
	return traversePostOrder(n, []int{})
}

func traverseInOrder(node *Node, list []int) []int {
	if node.Left != nil {
		list = traverseInOrder(node.Left, list)
	}
	list = append(list, node.Key)
	if node.Right != nil {
		list = traverseInOrder(node.Right, list)
	}
	return list
}

func traversePreOrder(node *Node, list []int) []int {
	list = append(list, node.Key)
	if node.Left != nil {
		list = traversePreOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePreOrder(node.Right, list)
	}
	return list
}

func traversePostOrder(node *Node, list []int) []int {
	if node.Left != nil {
		list = traversePostOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePostOrder(node.Right, list)
	}
	list = append(list, node.Key)
	return list
}

func main() {
	tree := &Node{Key: 9}
	buildTree := []int{4, 6, 20, 170, 15, 1}
	for _, node := range buildTree {
		tree.Insert(node)
	}
	// fmt.Println(tree.BreadthFirstSearchRecursive([]*Node{tree}, []int{}))
	fmt.Println(tree.DepthFirstSearchPreOrder())
}
