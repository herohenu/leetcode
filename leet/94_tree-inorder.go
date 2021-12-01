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

func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}

	arr = append(arr, inorderTraversal(root.Left)...)
	arr = append(arr, root.Val)

	arr = append(arr, inorderTraversal(root.Right)...)

	return arr
}

func InorderTraversal(root *TreeNode) []int {

	fmt.Println("root , ", root)
	return inorderTraversal(root)
}

func InitTree() *TreeNode {
	t5 := &TreeNode{Val: 5}
	t4 := &TreeNode{Val: 4}
	t3 := &TreeNode{Val: 3, Left: t4, Right: t5}
	t2 := &TreeNode{Val: 2}
	t1 := &TreeNode{Val: 1, Left: t2, Right: t3}
	return t1
}

func NotdiguiTravel(root *TreeNode) []int {
	return []int{}
}
