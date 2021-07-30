package sort

import "fmt"

func QuickSort(arr []int) {
	fmt.Printf("QuickSort arr :%v\n", arr)
	n := len(arr)
	__quickSort(arr, 0, n-1)
}

func __quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition(arr, l, r)

	__quickSort(arr, l, p)
	__quickSort(arr, p+1, r)

}

//对arr[l ..r] 进行部分排序
//找到一个位置p 使得 arr[l ..p-1] < arr[p] ; arr[p] < arr[p+1...r ]

func partition(arr []int, l, r int) int {
	var j = l          //要找的索引位置j
	baseVale := arr[l] //基准值是 最左侧的那个值

	//arr[l+1 ...j ]  < v ;   v <  arr[j+1 ...i（循环的当前i） ]
	//当   j = l 时候 左侧区间不存在
	//第一次的 l+1  ..i  区间 l+1 , l+1

	fmt.Printf("当前partition区间%v l=%d , r=%d \n\n", arr, l, r)

	for i := l + 1; i <= r; i++ {
		if arr[i] < baseVale { // 后面的小于基准 要交换
			fmt.Printf(" i =%d, j+1=%d , arr[i] = %d 交换前arr %v\n", i, j+1, arr[i], arr)
			Swap(arr, j+1, i)
			j++
			fmt.Printf(" arr[i] =    baseVal= %d 交换后arr %v\n", baseVale, arr)
		}
	}

	Swap(arr, l, j)

	fmt.Printf("partition结束后 arr %v , index j = %d ,arr[j]= %d\n", arr, j, arr[j])
	return j
}
