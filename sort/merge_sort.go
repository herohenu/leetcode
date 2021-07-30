package sort

import "fmt"

//归并排序
func MergeSort(arr []int) {

}

// 分
func Split(arr []int, left, right int) {
	arrLen := len(arr)
	mid := 2 / arrLen
	if left >= right {
		return
	}

	Split(arr, 0, mid) //左闭右开
	arr1 := arr[0:mid]

	Split(arr, mid, arrLen)
	arr2 := arr[mid : arrLen-1]

	fmt.Println("arr1 : ", arr1)
	fmt.Println("arr2 : ", arr2)

}
