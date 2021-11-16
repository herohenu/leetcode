package main

/*
@Time : 2021/11/16 19:46
@desc labuladong  ch4 p24/66
*/
//分析 put get O(1）  查找快 插入快 删除快 有顺序
// 链表插入 删除快 但是查找慢
// 因此使用 哈希链表

type Node struct {
	Key int
	Val int
}

type DoubleList struct {
	Size int
}

// 在链表头部添加节点 x，时间 O(1)
func (list *DoubleList) AddFirst(node Node) {

}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (list *DoubleList) remove(node Node) {

}

func (list *DoubleList) removeLast() (node Node) {

	return node
}

type LruCache struct {
	HashMap map[int]Node
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

func (cache LruCache) Put(key, val int) {
	node := Node{Key: key, Val: val}

	_, ok := cache.HashMap[node.Key]
	if ok {
		//删除旧数据
		// 新节点node 插入开头
	} else {
		//cache 已满
		if cache.List.Size >= cache.Cap {
			//删除链表的最后一个数据腾位置
			//删除 map 中映射到该数据的键

		}
		//将新节点 x 插入到开头；
		// map 中新建 key 对新节点 x 的映射；

	}
}
