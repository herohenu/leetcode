package main

import "fmt"

/*
@Time : 2021/12/1 9:42
@desc: 合并两个有序数组 golang实现
@author: iokde.com
*/

func main() {
	a1 := []int{1, 2, 3, 4}    // 5
	a2 := []int{1, 3, 3, 4, 5} // 3
	res := MergeArr(a1, a2)
	fmt.Printf("res is %v \n ", res)
}

func MergeArr(a1, a2 []int) []int {

	l1 := len(a1)
	l2 := len(a2)

	if l1 == 0 {
		return a2
	}

	if l2 == 0 {
		return a1
	}

	fmt.Println("a1 : ", a1)
	fmt.Println("a2 : ", a2)
	fmt.Printf("l1 : %d \n ", l1)
	fmt.Printf("l2 : %d \n ", l2)

	res := make([]int, l1+l2)

	i, j, k := 0, 0, 0

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
			fmt.Println("repeat ele is :", a2[j])
			j++
			//r++
		}
	}
	fmt.Println(" now res :", res)
	fmt.Println(" i is :", i)
	fmt.Println(" j is :", j)
	fmt.Println(" k is :", k)
	// 剩余部分

	if i < l1 {
		leftArr := a1[i:]
		fmt.Println("a1 left to be merge is : ", leftArr)
	}

	if j < l2 {
		fmt.Println("l2 left is :", a2[j:])
	}
	res = res[0:k]
	res = append(res, MergeArr(a1[i:], a2[j:])...)

	fmt.Println("res p : ", res)
	return res
}
