package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) NotValidateInsert(k *int) {
	if n.Left == nil {
		if k == nil {
			n.Left = nil
		} else {
			n.Left = &Node{Val: *k}
		}
	} else if n.Right == nil {
		if k == nil {
			n.Right = nil
		} else {
			n.Right = &Node{Val: *k}
		}
	} else {
		n.Left.NotValidateInsert(k)
	}
}

func isValidBST(root *Node) bool {
	return validateDFS(root, []*Node{})
}

func validateDFS(currentNode *Node, previousNode []*Node) bool {
	// validate currentNode
	var isCurrentNodeValid bool = true
	l := len(previousNode)

	// validate currentNode with latest node
	if l > 0 {
		if currentNode.Val > previousNode[l-1].Val {
			isCurrentNodeValid = currentNode.Val == previousNode[l-1].Right.Val
		} else {
			isCurrentNodeValid = currentNode.Val == previousNode[l-1].Left.Val
		}
	}

	if !isCurrentNodeValid {
		return false
	}

	if l > 1 {
		temp := previousNode[l-1]
		// validate currentNode with previous nodes
		for i := l - 2; i >= 0; i-- {
			if temp.Val < previousNode[i].Val {
				isCurrentNodeValid = currentNode.Val < previousNode[i].Val
			} else {
				isCurrentNodeValid = currentNode.Val > previousNode[i].Val
			}

			if !isCurrentNodeValid {
				return false
			}
		}
	}

	if currentNode.Left != nil {
		leftPrevious := append(previousNode, currentNode)
		return validateDFS(currentNode.Left, leftPrevious)
	}

	if currentNode.Right != nil {
		rightPrevious := append(previousNode, currentNode)
		return validateDFS(currentNode.Left, rightPrevious)
	} else {
		return true
	}
}

func main() {
	root := &Node{Val: 9}
	values := []interface{}{4, 20, 1, nil, 6, 15, 170}
	for _, v := range values {
		var addr *int
		if 
		fmt.Println(addr == nil)
		root.NotValidateInsert(addr)
	}
	// fmt.Println(isValidBST(root))
}
