package sort

import (
	"math/rand"
	"sync"
)

func GenArr(n, max int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		r := rand.Intn(max)
		arr = append(arr, r)
	}
	return arr
}

func Swap(arr []int, i, j int) {
	//fmt.Printf("before arr arr=%v , i= %d , j =%d \n", arr, i, j)
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
	//fmt.Println("after swap arr ", arr)
}

type Stack struct {
	Top   *node
	Lenth int
	lock  *sync.RWMutex
}

type node struct {
	Value interface{}
	Front *node
}

func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

func (s Stack) Len() int {
	return s.Lenth
}

func (s Stack) Peek() interface{} {
	if s.Lenth == 0 {
		return nil
	}

	return s.Top.Value
}

func (s Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Lenth == 0 {
		return nil
	}
	n := s.Top
	s.Top = s.Top.Front
	s.Lenth--
	return n.Value
}

func (s Stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	ele := &node{Value: value, Front: s.Top}
	s.Top = ele
	s.Lenth++
}
