package leet

import (
	"fmt"
	"time"
)

/*
@Time: 2021-12-02 20:07
@Desc:将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
@source:https://leetcode-cn.com/problems/merge-two-sorted-lists/

*/

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	list := new(ListNode)
	dumy := list
	p1 := l1
	p2 := l2
	for {
		if p1 == nil || p2 == nil {
			fmt.Println("l1.next ", p1)
			fmt.Println("l2.next ", p2)
			break
		}
		fmt.Println("l1.val is: ", l1.Val)
		fmt.Println("l2.val is: ", l2.Val)
		time.Sleep(1)
		if p1.Val < p2.Val {
			time.Sleep(1)
			fmt.Println(" l1.val < l2.val   ,list will point l1 ", l1.Val)
			list.Next = p1

			p1 = p1.Next
			list = list.Next
		} else if p1.Val > p2.Val {
			time.Sleep(1)
			fmt.Println(" l1.val > l2.val  ,l1.val :  ", l1.Val)
			list.Next = p2
			p2 = p2.Next
			list = list.Next
		} else {
			time.Sleep(1)

			list.Next = p1
			p1 = p1.Next

			list = list.Next //list->p1
			list.Next = p2
			p2 = p2.Next
			//list = list.Next
		}

	}

	if l1.Next != nil {
		list.Next = l1
	}
	if l2.Next != nil {
		list.Next = l2
	}

	return dumy.Next

}

//refer: https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0021.Merge-Two-Sorted-Lists/
func MergeTwoListsWithRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoListsWithRecursive(l1.Next, l2)
		return l1
	}

	l2.Next = MergeTwoListsWithRecursive(l1, l2.Next)

	return l2
}

func TestMergeTwoLists() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := buildList(arr1)
	l1.Show()

	l2 := buildList(arr2)
	l2.Show()
	//l := MergeTwoListsWithRecursive(l1, l2)
	//l.Show()

	l := MergeTwoLists(l1, l2)
	l.Show()
}
