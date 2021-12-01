package main

import "fmt"

/*
@Time : 2021/12/1 9:42
*/

func main() {
	a1 := []int{1, 2, 3}
	a2 := []int{2, 3, 4, 5, 5, 6, 7}
	res := MergeArr(a1, a2)
	fmt.Printf("res is %v \n ", res)
}

func MergeArr(a1, a2 []int) []int {
	l1 := len(a1)
	l2 := len(a2)
	fmt.Println("a1 : ", a1)
	fmt.Println("a2 : ", a2)
	fmt.Printf("l1 : %d \n ", l1)
	fmt.Printf("l2 : %d \n ", l2)
	res := make([]int, l1+l2)
	i, j, k, r := 0, 0, 0, 0

	for {
		if i >= l1 || i >= l2 || j >= l1 || j >= l2 {
			break
		}
		if a1[i] < a2[j] {
			res[k] = a1[i]
			k++
			i++
		} else if a1[i] < a2[j] {
			res[k] = a2[j]
			k++
			j++
		} else {
			j++
			r++
		}
	}
	// 剩余部分
	if i < l1 {
		for t := i; t < l1; t++ {
			res[k] = a1[t]
			k++
		}
	}

	if j < l2 {
		for t := j; t < l2; t++ {
			res[k] = a2[t]
			k++
		}
	}

	fmt.Println("repeat : ", r)
	fmt.Println("len(res) : ", len(res))
	p := len(res) - r
	res = res[:p]
	fmt.Println("res p : ", res)
	return res
}
