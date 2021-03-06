package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2021/11/25 19:46
@desc labuladong  ch4
@refer:https://www.cnblogs.com/labuladong/p/13975044.html
*/
//分析 get key 返回val 需要O（1）
// get /put 时候 key 对应的 freq +1  Incr(key)    keyFreqMap
// cap 满了 需要删除 freq 最小的 如果freq 相同 删除最旧的key

//1 使用map 存 key 到val 的映射  Key2Val
//2 需要存储key 对应的 freq 映射  keyFreqMap
//3 需要记录频次到key 的映射  key 有多个
//3.1 记录freq到key的映射  FreqKeyMap
//3.2 需要记录最小的freq  采用变量minFreq
//3.3 多个key 对应一个freq  那么freq-> key 是1：n 即 1个freq 对应列表 还能表示插入顺序
//3.4 希望 freq 对应的 key 的列表是存在时序的，便于快速查找并删除最旧的 key
//3.5 快速删除freqToKeyList 中的任意key ，  key 被访问后 变为了freq+1
// linkList 能解决3.3 4.4  无法满足快速删除
// hash(map) 能解决3.5 能快速方位删除 但是没有顺序
// 因此使用linkedHashMap map[string key ] linkedMap

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

func NewDoubleList(cap int) *DoubleList {
	list := new(DoubleList)
	//head := new(Node)
	//tail := new(Node)
	//list.Head = head
	//list.Tail = tail
	list.Head = list.Tail //init
	list.Capacity = cap
	list.lock = sync.RWMutex{}
	return list
}

type LFUCache struct {
	Key2Val   map[int]*Node       // key -> val 的映射
	Key2Freq  map[int]int         // key -> req 的映射
	Freq2Keys map[int]*DoubleList // freq 到keys 的映射
	MinFreq   int                 // 最小访问频次
	//List      *DoubleList
	Cap  int // 容量
	Size int // 实际存入的元素数量  可以用key2val 长度判断
	lock sync.RWMutex
}

func NewLFUCache(limit int) *LFUCache {
	cache := new(LFUCache)
	cache.Key2Val = make(map[int]*Node)
	cache.Key2Freq = make(map[int]int)
	cache.Freq2Keys = make(map[int]*DoubleList)

	cache.Cap = limit
	cache.MinFreq = 0 //  可以不用初始化
	//cache.List = new(DoubleList)
	//cache.List.Capacity = limit
	//cache.List.lock = new(sync.RWMutex)
	return cache
}
func (list *DoubleList) Traverse() {
	// 链表的遍历
	keys := []int{}
	cur := list.Head
	for cur.Next != nil {
		keys = append(keys, cur.Next.Key)
		cur = cur.Next
	}
}

// 在链表头部添加节点 x，时间 O(1)
// refer: https://segmentfault.com/a/1190000020042196
// node  必须使用指针
func (list *DoubleList) AddFirst(node *Node) bool {
	fmt.Println(" call Addfirst ... ", *node)
	if list == nil {
		fmt.Println("list is nil can't add elem ,please init list")
		return false
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
	fmt.Printf(" list remove node : %+v \n", *node)
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
	if list == nil || list.Tail == nil {
		fmt.Printf(" list is nil %+v \n", list)
		return node
	}

	if list.Tail == nil {
		fmt.Printf(" list.tail is nil,list %+v \n", list)
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

// 移除最小频次的list 最后一个
func (cache *LFUCache) RemoveMinFreqKey() bool {
	l := cache.Freq2Keys[cache.MinFreq]
	if l == nil {
		fmt.Printf("RemoveMinFreqKey error , minFreq is %d \n ", cache.MinFreq)
		return false
	}

	if l.Size == 0 {
		fmt.Printf("RemoveOldByFreqFromlist 出错了\n")
		return false
	}

	// todo ？？  need ?
	if l.Size == 1 {
		delete(cache.Freq2Keys, cache.MinFreq)
		return true
	}

	//update key2FreqList
	oldest := l.removeLast()
	fmt.Printf("removeOld node is : %+v \n", oldest)

	// update kv map ,delete old key

	delete(cache.Key2Val, oldest.Key)
	return true
}

func (cache *LFUCache) increaseFreq(key int) int {
	//todo lock
	freq := cache.Key2Freq[key]

	//cache.Key2Freq   update   √
	cache.Key2Freq[key] = freq + 1

	//kv not change √

	// old freq 2 key list remove key
	list := cache.Freq2Keys[freq]
	node := cache.Key2Val[key]
	list.removeLast()
	if list == nil || list.Size == 0 {
		delete(cache.Freq2Keys, key)
		// important!!!
		if cache.MinFreq == freq {
			cache.MinFreq++
		}
	}

	// new freq 2key list add key
	//这里一定是新的 还是需要判断??
	newlist := cache.Freq2Keys[freq+1]
	if newlist == nil {
		newlist = NewDoubleList(cache.Cap)
	}
	newlist.AddFirst(node)
	cache.Freq2Keys[freq+1] = newlist

	return freq
}

// diff: https://blog.csdn.net/u014779917/article/details/107857521
func (cache *LFUCache) Put(key, val int) {
	fmt.Printf("now put key %d ,val %d \n ", key, val)
	_, ok := cache.Key2Val[key]
	if ok {
		//覆盖老的kv
		cache.Key2Val[key] = &Node{Key: key, Val: val}

		cache.increaseFreq(key) // freq +1

		fmt.Printf("put sucess key: %d , val: %d  \n", key, val)
		return
	}

	//判断容量 是否满了 满了就淘汰最老的
	if cache.Size >= cache.Cap {
		b := cache.RemoveMinFreqKey()
		if b {
			cache.Size--
		} else {
			fmt.Printf(" remove old fail ,return !!!!")
			return
		}
	}

	// 然后重新放到 三个map
	newNode := &Node{Key: key, Val: val}
	cache.Key2Val[key] = newNode

	//update key Freq Map
	cache.MinFreq = 1
	cache.Key2Freq[key] = cache.MinFreq

	l := cache.Freq2Keys[cache.MinFreq]
	if l == nil {
		l = NewDoubleList(cache.Cap)

		l.AddFirst(newNode)
	}
	cache.Freq2Keys[cache.MinFreq] = l

	cache.Size++
}

// 获取的时候直接返回node 的val
// -1 标示不存在
func (cache *LFUCache) Get(key int) int {
	node, ok := cache.Key2Val[key]
	if !ok {
		return -1
	}

	//访问后 频次+1
	cache.increaseFreq(key)
	//更新频次到key list映射表
	//cache.RemoveMinFreqKey() // todo old freq  key list remove this key

	// 更新 key2freq 是否有必要？？
	// cache.Key2Freq[key] = newFreq IncrKeyFreq

	return node.Val
}

// test lfu
func main() {
	cache := NewLFUCache(3)
	cache.Put(1, 111)
	cache.Put(2, 2)

	v := cache.Get(1)
	fmt.Println("v : ", v)

	cache.Put(3, 3)
	cache.Put(1, 1)
	fmt.Println("key 2,val:  ", cache.Get(2))

	fmt.Println("cache.size : ", cache.Size)
}
