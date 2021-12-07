package leet

import "fmt"

/*
@Time : 2021/12/2 14:15
*/

// usage
//002  082
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList2(arr []int) *ListNode {
	node := new(ListNode)
	for _, v := range arr {
		node.Next = &ListNode{Val: v}
	}
	return node
}

func buildList(arr []int) *ListNode {
	//arr := []int{1, 3, 5, 7, 9}
	// 先构建头head
	head := &ListNode{Val: arr[0]}
	// 让tail 指向head
	tail := head
	for i := 1; i < len(arr); i++ {
		// 构建next 节点
		tail.Next = &ListNode{
			Val: arr[i],
		}
		//tail.next构建完成后tail后移
		tail = tail.Next
	}

	return head
}

func buildListWithNilNode(arr []int) *ListNode {

	return nil
}

//对象的方法调用形式 展示单链表 l 就是头head
func (l *ListNode) Show() {
	fmt.Printf(" %d", l.Val)
	//fmt.Printf(" %d ,%p  ", l.Val, l)

	for l.Next != nil {
		l = l.Next
		fmt.Printf("-> %d", l.Val)
		//fmt.Printf("-> %d ,%p", l.Val, l)
	}
	fmt.Println()
}
