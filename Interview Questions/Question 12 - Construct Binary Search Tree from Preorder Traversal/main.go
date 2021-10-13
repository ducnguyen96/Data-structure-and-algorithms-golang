package main

import "fmt"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	l := len(preorder)
	root := &TreeNode{Val: preorder[0]}
	for i := 1; i < l; i++ {
		helper(root, preorder[i])
	}
	return root
}

func helper(node *TreeNode, val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			helper(node.Left, val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			helper(node.Right, val)
		}
	}
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}
	root := bstFromPreorder(preorder)
	fmt.Println(root)
}
