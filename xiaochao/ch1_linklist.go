package main

import (
	"fmt"
	"time"
)

func main() {
	//l := build()
	//traVerse(l)
	//traVerse2(l)
	//l.Show()
	//
	//buildListAndShow()
	//buildListAndShow2()
	//ReverseList()
	ReverseN1(1)
	ReverseN1(2)
}

// 单链表的构建
// 一个一个手动构建

func build() *ListNode {
	list := &ListNode{Val: 1}
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	return list
}

// 构建链表并且展示
func buildListAndShow() {
	list := []int{1, 3, 5, 7, 9}
	// 先构建头head
	head := &ListNode{Val: list[0]}
	// 让tail 指向head
	tail := head
	for i := 1; i < len(list); i++ {
		// 构建next 节点
		tail.next = &ListNode{
			Val: list[i],
		}
		//tail.next构建完成后tail后移
		tail = tail.next
	}

	head.Show()

}
func buildListAndShow2() {
	l := &ListNode{}
	l.Insert2(2)
	l.Insert2(4)
	l.Insert2(6)
	l.Insert2(8)
	l.Show()

}

// 单链表的插入 只考虑插入 不考虑头节点空的情况（如果是整数，前面空节点int成了0）

func (h *ListNode) Insert(val int) {
	for h.next != nil {
		h = h.next
	}
	h.next = &ListNode{Val: val}
}

// 单链表的插入
func (p *ListNode) Insert2(val int) {

	for p.next != nil {
		p = p.next
	}

	p.next = &ListNode{Val: val}

	//如果首节点为空的时候（val 用interface{}可判断空） 往前移动
	if p.Val == nil {
		p.Val = p.next.Val
		p.next = p.next.next
	}
}

// 定义基本的单链表结构
type ListNode struct {
	Val  interface{}
	next *ListNode
}

// 使用for 循环的形式遍历
func traVerse(head *ListNode) {
	for {
		fmt.Println("node ", head.Val)
		head = head.next
		if head == nil {
			break
		}
	}
}

// 使用for 循环的另外一种形式（带条件的for）遍历
func traVerse2(head *ListNode) {
	fmt.Println("traVerse2 :", head.Val)
	for head.next != nil {
		head = head.next
		fmt.Println("traVerse2 :", head.Val)
	}
}

//对象的方法调用形式 展示单链表 l 就是头head
func (l *ListNode) Show() {
	fmt.Printf(" %d", l.Val)
	for l.next != nil {
		l = l.next
		fmt.Printf("-> %d", l.Val)
	}
	fmt.Println()
}

// 合并两个有序链表
func MergeOrderLinkList(l1 ListNode, l2 ListNode) {

}

func ReverseN1(n int) {
	l := build()
	fmt.Printf("ReverseN%d: List build ok", n)
	l.Show()

	fmt.Printf("ReverseN%d: now reverse List :\n", n)
	time.Sleep(1 * time.Second)
	l = reverseN(l, n)
	fmt.Printf("ReverseN%d : after reverse List :\n", n)
	l.Show()
}

func ReverseList() {
	l := build()
	println("List build ok")
	l.Show()

	println("now reverse List :")
	time.Sleep(1 * time.Second)
	l = Reverse(l)
	fmt.Println("after reverse List :")
	l.Show()

}

//递归反转整个链表
func Reverse(head *ListNode) *ListNode {
	if head.next == nil {
		return head
	}
	// debug info
	fmt.Println("now Reverse head.val:", head.Val)
	fmt.Println("now Reverse head.Nextis :")
	head.next.Show()

	//head 翻转之后 当做last
	last := Reverse(head.next)
	fmt.Println(" after Reverse last.Val:", last.Val)
	last.Show()
	head.next.next = head
	head.next = nil

	return last
}

// 1129
// 将链表的前 n 个节点反转（n <= 链表长度）
//ListNode reverseN(ListNode head, int n)
var successor *ListNode // 后驱节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第n+1 个几点
		successor = head.next
		return head
	}
	//以head.next 为七点，翻转前n-1 个节点
	last := reverseN(head.next, n-1)
	head.next.next = head

	head.next = successor
	return last

}
