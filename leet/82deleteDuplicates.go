package leet

/*
@Time : 2021/12/1 13:38
@desc:https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	//先遍历链表

	//如果cur 当前和下一个相同 那么 直接把cur 指向next.next 如果存在的话

	//
	cur := head

	if head.Next != nil && cur.Val == head.Val {
		cur = head.Next.Next
	}

	return cur
}

func main() {

}

func buildList() {

}
