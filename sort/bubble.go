package sort

import "fmt"

func Bubble(arr []int )  {
	for i:=0;  i< len(arr)-1 ; i++ {
		for j :=i+1 ; j<len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j],arr[i]
			}
		}
	}
	PrintArr(arr )
}

func PrintArr(arr []int)  {
	for _,v := range arr {
		fmt.Printf("%v ", v )
	}

}

