package main

import "fmt"

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return nil
}

func StartAdd() {
	l12 := ListNode{3, nil}
	l11 := ListNode{4, &l12}
	l10 := ListNode{2, &l11}

	fmt.Println(l10)
	//l22 := ListNode{ }
	//list1 := []ListNode{  l10 ,l11 ,l12}

}
