package sort

import "fmt"

/*
@Time : 2021/6/23 19:50
@File : shellsort
*/

func ShellSort(arr []int) {

}

//func InertSort(arr []int) {
//	//arr = []int{9, 3, 5, 4, 7, 1, 2}
//	fmt.Println("插入排序 :", arr)
//	//外层 右侧 未排序区域
//	for i := 1; i < len(arr); i++ {
//		//内层 左侧为排序完成
//		//从外层为界限，内部从右往左进行排序
//		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
//			arr[j-1], arr[j] = arr[j], arr[j-1]
//		}
//		fmt.Printf("第%d次结果： %v \n", i, arr)
//	}
//
//}

//	arr = []int{}
func ShellSortStep(arr []int) {
	arr = []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	fmt.Println("初始化 ：  arr ", arr)
	// 分组 [ 8,3],[9,5],[1,4],[7,6],[3,0]
	// 预期 [3,8] ,[5,9],[1,4],[6,7],[0,3]
	//  3,5,1,6,0 ,8,9,4,7,3
	//step 1 分成10 /2 = 5组 , 两两比较
	step := 10 / 2 // len(arr) / 2
	// 第一个for 循环是插入排序 从 第2个元素开始到数组末尾 ）
	for i := 5; i < len(arr); i++ {
		// 比较 i 和 i-step 的对比  进行粗略的对比移动 交换
		for j := i - 5; j >= 0; j -= 5 {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println("第1次后：arr ", arr)

	// 第二次 步长折半
	step = step / 2
	for i := step; i < len(arr); i++ {
		// 比较 i 和 i-step 的对比  进行粗略的对比移动 交换
		for j := i - step; j >= 0; j -= step {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println("第2次后：arr ", arr)
	//
	step = step / 2
	for i := step; i < len(arr); i++ {
		// 比较 i 和 i-step 的对比  进行粗略的对比移动 交换
		for j := i - step; j >= 0; j -= step {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println("第3次后：arr ", arr)

}
