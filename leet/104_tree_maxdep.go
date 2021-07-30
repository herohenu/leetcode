package leet

import "github.com/tal-tech/go-zero/core/mathx"

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return mathx.MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}

func Init104Tree() *TreeNode {
	t6 := &TreeNode{Val: 6}
	t5 := &TreeNode{Val: 5, Left: t6}
	t4 := &TreeNode{Val: 4}
	t3 := &TreeNode{Val: 3, Left: t4, Right: t5}
	t2 := &TreeNode{Val: 2}
	t1 := &TreeNode{Val: 1, Left: t2, Right: t3}
	return t1
}
