package main

/*
@Time : 2021/11/12 9:39
@from: https://labuladong.gitee.io/algo/2/20/48/
@Desc: 二叉堆  用到的操作 上浮 + 下沉
应用场景： 堆排序 和优先队列
说明 是个完全二叉树 根节点索引从0开始
*/

//code： https://www.jianshu.com/p/37bca5f2a6e9
// code2 : https://blog.nowcoder.net/n/d6800fbabade4d04997194ac620bd744?from=nowcoder_improve
type MaxHeap struct {
	size int
	nums []int
}

// return parent Index
func parent(root int) int {
	if root == 0 {
		return 0
	}
	return (root - 1) / 2
}

// return left child index
func left(root int) int {

	return root*2 + 1
}

// return right child index
func right(root int) int {
	return root * 2
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (heap *MaxHeap) GetSize() int {
	return heap.size
}

func (heap *MaxHeap) IsEmpty() bool {
	return len(heap.nums) == 0
}

// 插入元素
func (heap *MaxHeap) SiftUp(i int) {

}
