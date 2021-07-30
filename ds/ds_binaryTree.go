package main

import "fmt"

/**
 *https://leetcode-cn.com/problems/merge-two-binary-trees/
 * 20201207
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitTreeNode1() *TreeNode {
	var TreeNode3 = &TreeNode{
		Val: 3,
	}

	var TreeNode2 = &TreeNode{
		Val: 2,
	}

	var TreeNode5 = &TreeNode{
		Val: 5,
	}

	TreeNode3.Left = TreeNode5

	var TreeNode1 = &TreeNode{Val: 1, Left: TreeNode3, Right: TreeNode2}

	return TreeNode1
}

func InitTreeNode2() *TreeNode {
	var TreeNode4 = &TreeNode{
		Val: 4,
	}

	var TreeNode7 = &TreeNode{
		Val: 7,
	}

	var TreeNode1 = &TreeNode{
		Val:   1,
		Right: TreeNode4,
	}

	var TreeNode3 = &TreeNode{
		Val:   3,
		Right: TreeNode7,
	}

	var TreeNode2 = &TreeNode{
		Val:   2,
		Left:  TreeNode1,
		Right: TreeNode3,
	}

	return TreeNode2
}

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	if t1.Left == nil && t2.Left != nil {
		t1.Left = &TreeNode{}
	}
	if t1.Right == nil && t2.Right != nil {
		t1.Right = &TreeNode{}
	}

	// 都有值 直接相加
	t1.Val += t2.Val
	MergeTrees(t1.Left, t2.Left)
	MergeTrees(t1.Right, t2.Right)

	return t1
}

func PreTravelTree(t *TreeNode) {
	if t == nil {
		//fmt.Println(" ")
		return
	}

	fmt.Printf("%d \t ", t.Val)
	PreTravelTree(t.Left)
	PreTravelTree(t.Right)
}

func InOrderTree(t *TreeNode) {
	if t == nil {
		return
	}
	InOrderTree(t.Left)
	fmt.Printf("%d \t", t.Val)
	InOrderTree(t.Right)
}

func InitTree() *TreeNode {
	var t = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(" tree ", t)
	return t
}
