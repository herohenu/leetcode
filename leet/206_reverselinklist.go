package leet

/*
@Time : 2021/7/14 9:31
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *ListNode) *ListNode {
	pre := new(ListNode)
	cur := head
	for cur != nil {
		next := cur.Next //1
		cur.Next = pre
		pre = cur
		cur = next

	}

	return pre
}

func InitReverselist() *ListNode {
	head := new(ListNode)
	v1 := &ListNode{Val: 1}
	v2 := &ListNode{Val: 2}
	v3 := &ListNode{Val: 3}
	v4 := &ListNode{Val: 4}
	v5 := &ListNode{Val: 5}

	v1.Next = v2
	v2.Next = v3
	v3.Next = v4
	v4.Next = v5
	head.Next = v1
	return head
}

//fmt.Println("cur is ", cur.Val)
//next := cur.Next //? head.next  1
//fmt.Println("next is ", next)
//
//cur.Next = pre
//fmt.Println("cur.next is pre:", pre)
//pre = cur
//fmt.Println("pre is :", pre)
//cur = next
//fmt.Println("cur is :", cur)
