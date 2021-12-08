package leet

import (
	"leetcode/sort"

	"github.com/golang/go/src/pkg/fmt"
)

/*
@Time : 2021/12/8 11:09
*/

func backspaceCompare(s string, t string) bool {

	s = getVal(s)
	t = getVal(t)

	return s == t
}

func TestBackspaceCompare() {
	s1 := "ab###c"
	s2 := "ad###c"
	//s1 = getVal(s1)
	//s2 = getVal(s2)
	//if s1 == s2 {
	//	fmt.Println("=====")
	//} else {
	//	fmt.Println("not == ")
	//}

	ret := backspaceCompare(s1, s2)
	fmt.Println("compare resutl ", ret)
}

func getVal(s1 string) string {
	stack := sort.NewStack()
	for _, v := range s1 {
		if string(v) != "#" {
			stack.Push(string(v))
		} else {
			stack.Pop()
		}
	}
	//fmt.Println("stack is:", stack.Lenth)
	s := ""
	for i := stack.Lenth; i > 0; i-- {
		pv := stack.Pop()
		//fmt.Println(" pop out val : ", pv)
		//fmt.Println(" pop out val len : ", stack.Lenth)
		s += fmt.Sprint(pv)
	}

	return s
}
