package main

import "fmt"

// TreeNode represents the components of a binary search tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// {1, 4, nil, nil, 3, 6}
func NotValidateInsert(queue []*TreeNode, k []*int, index *int) {
	l := len(k)
	if *index > l {
		return
	}

	currentNode := queue[0]
	queue = queue[1:]

	if *index < l {
		if k[*index] != nil {
			currentNode.Left = &TreeNode{Val: *k[*index]} // {1}
			queue = append(queue, currentNode.Left)
		}
		*index++ // 1 4
	} else {
		return
	}

	if *index < l {
		if k[*index] != nil {
			currentNode.Right = &TreeNode{Val: *k[*index]} // {4}
			queue = append(queue, currentNode.Right)
		}
		*index++ // 2 5
	} else {
		return
	}

	NotValidateInsert(queue, k, index)
}

func IsValidBST(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	left := validateDFS(root.Left, []*TreeNode{root})
	if !left {
		return false
	}
	right := validateDFS(root.Right, []*TreeNode{root})
	return left && right
}

func validateDFS(currentNode *TreeNode, previousNode []*TreeNode) bool {
	if currentNode == nil {
		return true
	}

	// validate currentNode
	var isCurrentNodeValid bool = true
	l := len(previousNode)

	// validate currentNode with latest node
	if l > 0 {
		if currentNode.Val > previousNode[l-1].Val {
			if previousNode[l-1].Right == nil {
				return false
			}
			isCurrentNodeValid = currentNode.Val == previousNode[l-1].Right.Val
		} else if currentNode.Val < previousNode[l-1].Val {
			if previousNode[l-1].Left == nil {
				return false
			}
			isCurrentNodeValid = currentNode.Val == previousNode[l-1].Left.Val
		} else {
			return false
		}
	}

	if !isCurrentNodeValid {
		return false
	}

	if l > 1 {
		// temp := previousNode[l-1]
		// validate currentNode with previous nodes
		for i := l - 1; i >= 1; i-- {
			if previousNode[i].Val < previousNode[i-1].Val {
				isCurrentNodeValid = currentNode.Val < previousNode[i-1].Val
			} else if previousNode[i].Val > previousNode[i-1].Val {
				isCurrentNodeValid = currentNode.Val > previousNode[i-1].Val
			} else {
				return false
			}

			if !isCurrentNodeValid {
				return false
			}
		}
	}

	if currentNode.Left != nil && isCurrentNodeValid {
		leftPrevious := append(previousNode, currentNode)
		isCurrentNodeValid = validateDFS(currentNode.Left, leftPrevious)
	}

	if currentNode.Right != nil && isCurrentNodeValid {
		rightPrevious := append(previousNode, currentNode)
		isCurrentNodeValid = validateDFS(currentNode.Right, rightPrevious)
	}

	return isCurrentNodeValid
}
func AlternativeSolution(root *TreeNode) bool {
	return RecValidate(root, nil, nil)
}

func RecValidate(n, min, max *TreeNode) bool {
	if n == nil {
		return true
	}
	if min != nil && n.Val <= min.Val {
		return false
	}
	if max != nil && n.Val >= max.Val {
		return false
	}
	return RecValidate(n.Left, min, n) && RecValidate(n.Right, n, max)
}

func main() {
	// [2,1,3]
	// [5,1,4,null,null,3,6]
	// [32,26,47,19,null,null,56,null,27, 55] v
	// [9, 4, 20, 1, 6, 15, 170]
	// [1, 1]
	// [0,nil,-1]
	// [3,1,5,0,2,4,6,null,null,null,3]
	// [5,4,6,null,null,3,7]

	root := &TreeNode{Val: 5}
	values := []interface{}{4, 6, nil, nil, 3, 7}
	nodes := make([]*int, len(values))
	var index *int = new(int)
	*index = 0
	for i, v := range values {
		var addr *int
		if v != nil {
			addr = new(int)
			*addr = v.(int)
		}
		nodes[i] = addr
	}

	NotValidateInsert([]*TreeNode{root}, nodes, index)

	isValid := AlternativeSolution(root)
	fmt.Println("isValid :", isValid)
}
