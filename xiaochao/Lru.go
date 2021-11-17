package main

import (
	"fmt"
	"sync"
)

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
	Size     int // 当前容量
	Capacity int
	lock     sync.RWMutex
	Head     *Node // 初始化
	Tail     *Node
}

// 在链表头部添加节点 x，时间 O(1)
// refer: https://segmentfault.com/a/1190000020042196
// node  必须使用指针
func (list *DoubleList) AddFirst(node *Node) bool {
	fmt.Println(" call Addfirst ... ", *node)
	if list == nil {
		fmt.Println("list is nil ,will new list ")
		list = new(DoubleList)
	}

	if list.Capacity == 0 {
		//false
		fmt.Println("list size ===0 return false ")
		return false
	}
	list.lock.Lock()
	defer list.lock.Unlock()

	//判断头结点是否存在
	if list.Head == nil {
		fmt.Println("head  不存在 。。 ")
		list.Head = node
		list.Tail = node
	} else {
		fmt.Println("head  已经存在 ！！ ")
		//头结点存在
		list.Head.Pre = node
		node.Next = list.Head

		//设置新的头节点
		list.Head = node
		list.Head.Pre = nil
	}
	list.Size++
	fmt.Printf("now list is :%+v \n", list)

	fmt.Printf("now list.size is :%+v \n", list.Size)
	return true
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
// node  必须使用指针 不然传入的是新的对象 找不到pre 会报错
func (list *DoubleList) remove(node *Node) *Node {
	fmt.Println(" 节点 被淘汰了: ", *node)
	fmt.Println(" list will remove addr : ", &node)
	fmt.Println(" list head is : ", list.Head)
	if node == list.Head {
		return list.removeHead()
	}

	fmt.Println(" list Tail is : ", list.Tail)
	if node == list.Tail {
		return list.Tail
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	fmt.Println(" list is : ", list)
	fmt.Println(" list size : ", list.Size)
	//如果是中间节点
	node.Pre.Next = node.Next // todo fix  panic
	node.Next.Pre = node.Pre
	list.Size--
	return node

}

func (list *DoubleList) removeHead() (node *Node) {
	if list.Size == 0 || list.Head == nil {
		return nil
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	node = list.Tail
	//判断尾的前一个是否存在
	if node.Next != nil {
		list.Head = node.Next
		node.Next.Pre = nil
	}
	list.Size--

	return node
}

func (list *DoubleList) removeLast() (node *Node) {
	if list.Tail == nil {
		return node
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	node = list.Tail
	//判断尾的前一个是否存在
	if node.Pre != nil {
		list.Tail = node.Pre
		list.Tail.Next = nil
	}
	list.Size--

	return node
}

type LruCache struct {
	HashMap map[int]*Node // 删除list 中node 的时候 根据key 删除map中的val
	List    *DoubleList
	Cap     int
}

func NewLruCache(limit int) *LruCache {
	cache := new(LruCache)
	cache.HashMap = make(map[int]*Node)
	cache.Cap = limit
	cache.List = new(DoubleList)
	cache.List.Capacity = limit
	//cache.List.lock = new(sync.RWMutex)
	return cache
}

func (cache LruCache) Put(key, val int) {
	newNode := &Node{Key: key, Val: val}
	fmt.Println(" new node :", newNode)

	oldNode, ok := cache.HashMap[key]
	if ok {
		fmt.Printf(" key  exists %d , node : %+v", key, cache.HashMap)
		//删除旧数据
		cache.List.remove(oldNode)

		// 新节点node 插入开头
		cache.HashMap[key] = newNode
		cache.List.AddFirst(newNode)

	} else {
		fmt.Println(" newNode 不在 cache.list ", newNode)
		//cache 已满
		if cache.List.Size >= cache.Cap {
			//删除链表的最后一个数据腾位置
			//删除 map 中映射到该数据的键
			last := cache.List.removeLast()
			delete(cache.HashMap, last.Key)
			fmt.Println(" 容量已满 ,删除节点", *last, " last.key :", last.Key)
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

// test lru
func main() {
	cache := NewLruCache(3)
	cache.Put(1, 111)
	cache.Put(2, 2)

	v := cache.Get(1)
	fmt.Println("v : ", v)

	cache.Put(3, 3)
	cache.Put(1, 1)
	fmt.Println("key 2,val:  ", cache.Get(2))

}
