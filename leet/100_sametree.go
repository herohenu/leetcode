package main

/***
date: 2021-5-8
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {

	brr := DiffTree(p, q)
	b := true
	for _, v := range brr {
		b = b && v
		if !b {
			return false
		}
	}
	return b
}

func DiffTree(p, q *TreeNode) []bool {
	brr := []bool{}
	if p == nil && q == nil {
		brr = append(brr, true)
		return brr
	}
	if p == nil || q == nil {
		brr = append(brr, false)
		return brr
	}

	if p.Val != q.Val {
		brr = append(brr, false)
		return brr
	}

	brr = append(brr, DiffTree(p.Left, q.Left)...)
	brr = append(brr, DiffTree(p.Right, q.Right)...)

	return brr
}
