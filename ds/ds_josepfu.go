package main

import "fmt"

/***
0513 约瑟夫问题
*/
type Boy struct {
	No   int
	next *Boy
}

type CircleSigleLinkedList struct {
	First *Boy
}

func (list *CircleSigleLinkedList) AddBoy(size int) {
	if size < 1 {
		return
	}

	//创建链表 辅助指针
	var cur *Boy
	//var list CircleSigleLinkedList

	for i := 1; i <= size; i++ {
		boy := &Boy{No: i}
		if i == 1 { //单个元素构建环形链表 后续的加入环
			list.First = boy
			list.First.next = list.First
			cur = boy
			fmt.Println("添加了第一个成员 ：", cur.No)
		} else {
			cur.next = boy
			boy.next = list.First
			cur = boy
			fmt.Println("添加了成员 ：", cur.No)
		}

	}
}

func (list *CircleSigleLinkedList) Show() {
	if list == nil {
		fmt.Println("没有任何成员")
		return
	}
	var cur *Boy
	cur = list.First
	for {
		fmt.Println("boy 编号 no", cur.No)
		if cur.next == list.First {
			break
		}

		if cur.next != nil {
			cur = cur.next
		}
	}

}

func main() {
	list := new(CircleSigleLinkedList)
	list.AddBoy(5)
	list.Show()
}
