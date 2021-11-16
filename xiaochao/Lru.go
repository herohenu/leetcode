package main

import "sync"

/*
@Time : 2021/11/16 19:46
@desc labuladong  ch4 p24/66
*/
//分析 put get O(1）  查找快 插入快 删除快 有顺序
// 链表插入 删除快 但是查找慢
// 因此使用 哈希链表

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
	Freq int
}

type DoubleList struct {
	Size int // 当前容量
	lock *sync.RWMutex
	Head *Node
	Tail *Node
}

// 在链表头部添加节点 x，时间 O(1)
// refer: https://segmentfault.com/a/1190000020042196
func (list *DoubleList) AddFirst(node Node) bool {
	if list == nil {
		list = new(DoubleList)
	}
	if list.Size == 0 {
		//false
		return false
	}
	list.lock.Lock()
	defer list.lock.Unlock()

	//判断头结点是否存在
	if list.Head == nil {
		list.Head = &node
		list.Tail = &node
	} else {
		//头结点存在
		list.Head.Pre = &node
		node.Next = list.Head

		//设置新的头节点
		list.Head = &node
		list.Head.Pre = nil
	}
	list.Size++
	return true
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (list *DoubleList) remove(node Node) {

}

func (list *DoubleList) removeLast() (node Node) {

	return node
}

type LruCache struct {
	HashMap map[int]Node // 删除list 中node 的时候 根据key 删除map中的val
	List    *DoubleList
	Cap     int
}

func NewLruCache(limit int) *LruCache {
	cache := new(LruCache)
	cache.HashMap = make(map[int]Node)
	cache.Cap = limit
	cache.List = new(DoubleList)
	return cache
}

//todo
func (cache LruCache) Put(key, val int) {
	newNode := Node{Key: key, Val: val}

	oldNode, ok := cache.HashMap[key]
	if ok {
		//删除旧数据
		cache.List.remove(oldNode)
		// 新节点node 插入开头
		cache.HashMap[key] = newNode
		cache.List.AddFirst(newNode)

	} else {
		//cache 已满
		if cache.List.Size >= cache.Cap {
			//删除链表的最后一个数据腾位置
			//删除 map 中映射到该数据的键
			last := cache.List.removeLast()
			delete(cache.HashMap, last.Key)

		}
		// 将新节点 x 插入到开头；
		// map 中新建 key 对新节点 x 的映射；
		cache.List.AddFirst(newNode)
		cache.HashMap[key] = newNode
	}
}

func (cache LruCache) Get(key int) int {
	nodeVal, ok := cache.HashMap[key]
	if !ok {
		return -1
	}

	//访问后 数据提前
	cache.Put(key, nodeVal.Val)

	return nodeVal.Val
}
