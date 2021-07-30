package main

import (
	"fmt"
	"leetcode/leet"
	"leetcode/sort"
	"math/rand"
	"strconv"
)

func main() {

	//ReverseArr(arr)
	//rates := RateCut()
	//for i:=0;i <10  ;i++    {
	//	GetWord( rates)
	//}
	//sort.Bubble( arr )
	//sort.Inert(arr )
	//leet.ThreeSumMain()
	//arr := sort.GenArr(10, 100)
	//sort.InertSort(arr)
	//fmt.Println("ARR IS :", arr)

	//sort.Split(arr, 0, len(arr))

	//sort.QuickSort(arr)
	//fmt.Printf("last arr:%v\n ", arr)
	//FibStart()
	//FibSync()
	//root := leet.InitTree()
	//
	//arr := leet.InorderTraversal(root)
	//fmt.Println("arr :", arr)
	//
	//root104 := leet.Init104Tree()
	//fmt.Println(" max depth : ", leet.MaxDepth(root104))
	//fmt.Println(" max depth : ", leet.MaxDepth(root))
	node := leet.InitReverselist()
	p := leet.ReverseList(node)
	//for node.Next != nil {
	//	node = node.Next
	//	fmt.Println("node.val :", node.Val)
	//}
	for p.Next != nil {
		fmt.Println("p.val :", p.Val)
		p = p.Next
		//fmt.Println("p.val :", p.Val)
	}
	fmt.Println("p is:", p)

}
func m() {
	a := leet.LengthOfLastWord("aa bbxxx  x")
	fmt.Println(a)
	arr := []int{}
	sort.ShellSortStep(arr)
}
func tmp() {
	s := "()[]x"
	for idx, _ := range s {
		fmt.Println(s[idx])
	}

}
func testx() {
	arr := []int{1, 3, 5, 9, 0, 2}
	mp := make(map[int]int)
	for i := 0; i < 100; i++ {
		randN := arr[rand.Intn(len(arr))]
		mp[randN] += 1
	}
	fmt.Println("map ", mp)
}

func DultiAssign() {
	a, b := 10, 20
	a, b = b, a
	fmt.Println("a , b : ", a, b)
}

func RateCut() [][]float64 {
	arr1 := []int{1, 1, 1, 3, 2, 2}

	ratesum := 0
	rates := [][]float64{}
	midSum := []int{} //中间值
	for _, v := range arr1 {
		ratesum += v
		fmt.Printf(" v %d , sum : %d  \n", v, ratesum)
		midSum = append(midSum, ratesum)
	}

	for k, v := range midSum {
		left := 0.0
		res := float64(v*1.0) / float64(ratesum)
		if k != 0 {
			left = Decimal(float64(midSum[k-1]) / float64(ratesum))
		}
		res = Decimal(res)
		fmt.Println(" res ", left, res)
		rates = append(rates, []float64{left, res})
	}
	return rates
}

func GetWord(rates [][]float64) string {
	//fmt.Println("GetWord ... ")
	arr2 := []string{"w1", "w2", "w3", "w4", "w5", "w6"}
	r := rand.Float64()
	selected := -1
	for k, rate := range rates {
		if len(rate) != 2 {
			continue
		}

		if r >= rate[0] && r <= rate[1] {
			selected = k
			fmt.Println(" current rate is", rate)
			break
		}
	}
	word := ""
	if selected != -1 {
		word = arr2[selected]
		fmt.Printf("rate %f , use word :%s  \n ", r, arr2[selected])
	}

	return word
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func ReverseArr(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	fmt.Println("arr :", arr)
}
