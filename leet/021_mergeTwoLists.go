package leet

/*
@Time: 2021-12-02 20:07
@Desc:将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
@source:https://leetcode-cn.com/problems/merge-two-sorted-lists/

*/

//参考： https://github.com/wx-satellite/learning-algorithm/tree/master/10_%E5%90%88%E5%B9%B6%E4%B8%A4%E4%B8%AA%E6%9C%89%E5%BA%8F%E9%93%BE%E8%A1%A8
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := new(ListNode)
	var cur = head

	for {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}

		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	return head.Next

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

// https://blog.csdn.net/sinat_38068807/article/details/103811475
func MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var list, cur *ListNode

	if l1.Val < l2.Val {
		list = l1
		//cur = l1
		l1 = l1.Next
	} else {
		list = l2
		//cur = l2
		l2 = l2.Next
	}
	cur = list

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}

	if l2 == nil {
		cur.Next = l1
	}
	return list
}

func MergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {

	list := new(ListNode)
	var cur = list

	p1 := l1
	p2 := l2

	for {
		if p1 == nil {
			cur.Next = p2
			break
		}

		if p2 == nil {
			cur.Next = p1
			break
		}

		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}

	return list.Next

}

func TestMergeTwoLists() {
	arr1 := []int{1, 2, 5, 7}
	arr2 := []int{1, 2, 3, 4}
	l1 := buildList(arr1)
	l1.Show()

	l2 := buildList(arr2)
	l2.Show()
	//l := MergeTwoListsWithRecursive(l1, l2)
	//l.Show()

	l := MergeTwoLists2(l1, l2)
	l.Show()
}
