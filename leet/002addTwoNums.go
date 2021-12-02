package leet

import "fmt"

/*
@Time : 2021/12/2 14:10
*/

/**

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

题目大意 #
2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。

*/

func main() {
	fmt.Println("-----------")
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	l1 := buildList(arr1)
	l2 := buildList(arr2)
	l1.Show()
	l2.Show()

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
