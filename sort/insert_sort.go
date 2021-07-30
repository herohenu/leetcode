package sort

import "fmt"

//0310
func InertSort(arr []int) {
	//arr = []int{9, 3, 5, 4, 7, 1, 2}
	fmt.Println("插入排序 :", arr)
	//外层 右侧 未排序区域
	for i := 1; i < len(arr); i++ {
		//内层 左侧为排序完成
		//从外层为界限，内部从右往左进行排序
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
		fmt.Printf("第%d次结果： %v \n", i, arr)
	}

}
