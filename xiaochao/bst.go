package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
@Time : 2021/11/18 9:47
@Desc 二叉搜索树（Binary Search Tree 左边节点比根小 右边节点比根大
refer: https://segmentfault.com/a/1190000022331472
*/
var arr []int

func main() {
	//生成树的叶子节点数量
	//nodeNum := 5
	//t := buildTreeWithRandom(nodeNum, 100)
	//fmt.Printf(" t is : %+v \n", t)
	//traverse(t)
	//for i := 0; i < nodeNum; i++ {
	//	targetVal := arr[i]
	//	fmt.Println("search val is : ", targetVal)
	//	v := t.Search(targetVal)
	//	fmt.Println("v is:  ", v.Val, "v.Left is:  ", v.Left, "v.Right is:  ", v.Right)
	//}
	//
	//leftCount := Count(t.Left)
	//fmt.Printf(" 左子树有 【%d】个几点\n", leftCount)
	//rightCount := Count(t.Right)
	//fmt.Printf(" 右子树有 【%d】个几点\n", rightCount)
	arr = []int{8, 10, 11, 7, 6}
	t := buildTreeUserInsert(arr)
	fmt.Println(" 前序遍历...... ")
	traverse(t)
	traverse2(t)
	fmt.Println()
	mirrorTree(t)
	fmt.Println(" 镜像后遍历.... ")
	traverse2(t)
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeUserInsert(arr []int) *TreeNode {
	t := new(TreeNode)
	fmt.Printf("arr is %+v \n", arr)
	for _, v := range arr {
		node := &TreeNode{Val: v}
		t = t.Insert(node)
	}
	return t
}

func buildTreeWithRandom(n int, max int) *TreeNode {
	t := new(TreeNode)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		val := rand.Intn(max)
		arr = append(arr, val)
		fmt.Printf("generate val : %d ", val)
		node := &TreeNode{Val: val}
		t.Insert(node)
		fmt.Printf(" now tree is : %+v \n", t)
	}
	fmt.Println("arr is:", arr)
	return t
}

// root val 问题 第一个
func (t *TreeNode) Insert(node *TreeNode) *TreeNode {
	if t.Val == node.Val {
		fmt.Println("插入的val 和已经存在的节点val相等 return")
		return t
	}

	if node == nil {
		fmt.Println("insert node is nil ")
		return t
	}

	if t == nil || t.Val == 0 {
		t.Val = node.Val
		return t
	}
	//插入到左侧
	if t.Val > node.Val {
		fmt.Printf(" node.val %d < t.val %d  node insert left \n ", node.Val, t.Val)
		if t.Left == nil {
			t.Left = node
		} else {
			t.Left.Insert(node)
		}
	} else {
		fmt.Printf(" node.val %d > t.val %d  ,node insert right \n ", node.Val, t.Val)
		if t.Right == nil {
			t.Right = node
		} else {
			t.Right.Insert(node)
		}
	}

	return t
}

// desc： https://juejin.cn/post/6950288278662676493

func (t *TreeNode) InserttoBST(val int) *TreeNode {
	if t.Val == val {
		return t
	}

	if t == nil {
		t.Val = val
		return t
	}

	newNode := &TreeNode{Val: val}
	p := t

	for p != nil {
		if val < t.Val {
			if p.Left == nil {
				p.Left = newNode
				break
			}
			//遍历
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = newNode
				break
			}
			p = p.Right
		}
	}

	return t
}

func (t *TreeNode) Remove(node *TreeNode) {

}

// 对于有序的二叉树，左子节点小于当前节点，右子节点大于当前节点
func (t *TreeNode) Search(val int) *TreeNode {
	if t == nil {
		return nil
	}

	if val == t.Val {
		//fmt.Println("找到了")
		return t
	}
	// 搜索的val 大于根节点的val 从右侧搜索
	if val > t.Val {
		return t.Right.Search(val)
	}

	if val < t.Val {
		return t.Left.Search(val)
	}
	return nil
}

func traverse(root *TreeNode) {
	// 前序遍历
	if root != nil {
		fmt.Println(" ", root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}
}

func traverse2(root *TreeNode) {
	// 前序遍历
	if root == nil {
		return
	}
	fmt.Printf(" %d **", root.Val)
	traverse2(root.Left)
	traverse2(root.Right)

}

func Count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Count(root.Left) + Count(root.Right)
}

// 节点翻转
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	fmt.Printf("root is %d ,left : %+v , right: %+v \n", root.Val, root.Left, root.Right)
	tmpNode := root.Left
	root.Left = root.Right
	root.Right = tmpNode
	//fmt.Println("now  inverse root.left", root.Left)
	root.Left = mirrorTree(root.Left)
	//fmt.Println("now  inverse root.Right", root.Right)
	root.Right = mirrorTree(root.Right)
	return root
}
