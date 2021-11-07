package main

import "fmt"

func main() {
	l := build()
	traVerse(l)
	traVerse2(l)
	l.Show()

	buildListAndShow()
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

// 单链表的插入
func (h *ListNode) Insert(val int) {
	for h.next != nil {
		h = h.next
	}
	h.next = &ListNode{Val: val}
}

// 定义基本的单链表结构
type ListNode struct {
	Val  int
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
	fmt.Println("show :", l.Val)
	for l.next != nil {
		l = l.next
		fmt.Println("show :", l.Val)
	}
}
