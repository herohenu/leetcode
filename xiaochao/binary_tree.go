package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//t := BuildTree(10)
	//t.PreOrderTravel()
	t := new(TreeNode)
	arr := []int{1, 3, 5, 6, 9, 10, 11, 2}
	t.Array2Tree(arr)
	traverse(t)

}

// 假设树中没有重复的节点
type TreeNode2 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//向树中填充随机数值
func BuildTree(n int) *TreeNode {
	tree := new(TreeNode)
	//设置随机种子为当前时间
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		tmp := rand.Intn(n * 2)
		fmt.Println("tmp is :", tmp)
		tree = tree.Insert(tmp)
	}
	//tree.Val = n

	return tree
}

//todo
func (t *TreeNode) Array2Tree(arr []int) *TreeNode {
	for _, v := range arr {
		t.Insert(v)
	}

	return t
}

//向树中插入节点
func (t *TreeNode) Insert(val int) *TreeNode {
	fmt.Println("call Insert ,val :", val)
	if t == nil {
		t = &TreeNode{val, nil, nil}
		return t
	}

	//假设没有重复的val
	if t.Val == val {
		fmt.Println("val has exists, will not insert ")
		return t
	}

	// 如果值小于根节点 插入到左侧节点
	if val < t.Val {
		t.Left = t.Left.Insert(val)
		return t
	}
	//如果值大于根节点 插入到右侧节点
	t.Right = t.Right.Insert(val)

	return t
}

func (t *TreeNode) Print() {
	fmt.Printf("  %d \t ", t.Val)
}

func (t *TreeNode) PreOrderTravel() {
	if t == nil {
		return
	}
	t.Print()
	t.Left.PreOrderTravel()
	t.Right.PreOrderTravel()
}

func traverse(t *TreeNode) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Val, " ")
	traverse(t.Right)
}
