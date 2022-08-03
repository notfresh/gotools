package gotools

import (
	"container/heap"
)

// IntHeap 是一个由整数组成的最小堆。
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push 和 Pop 使用 pointer receiver 作为参数，
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1] //最后一个？
	*h = old[0 : n-1]
	return x
}

///////////////////////////////////////////
// Item 是优先队列中包含的元素。
type Item struct {
	value    interface{} // 元素的值，可以是任意类型。
	priority int         // 元素在队列中的优先级。
	// 元素的索引可以用于更新操作，它由 heap.Interface 定义的方法维护。
	index int // 元素在堆中的索引。
}

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 返回的是最大值而不是最小值，
	// 因此这里使用大于号进行对比。
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) Update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

/////////////////////////////

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueueR []*Item

func NewPriorityQueueR(data []*Item) PriorityQueueR {
	return data
}

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。

func (pq *PriorityQueueR) Len() int { return len(*pq) }

func (pq *PriorityQueueR) Less(i, j int) bool {
	// 我们希望 Pop 返回的是最大值而不是最小值，
	// 因此这里使用大于号进行对比。
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueueR) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueueR) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueR) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueueR) Update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type ZPriorityQueue struct {
	data  []*Item
	order bool
}

func NewZPriorityQueue(flag bool) *ZPriorityQueue {
	pq := &ZPriorityQueue{
		order: flag,
		data:  make([]*Item, 0),
	}
	//heap.Init(pq.data.(PriorityQueue))
	//if flag {
	//	pq.data = make(PriorityQueue, 0)
	//	z := pq.data
	//	heap.Init(&z)
	//} else {
	//	pq.data = make(PriorityQueueR, 0)
	//}
	return pq
}

func (pq *ZPriorityQueue) Len() int { return len(pq.data) }

func (pq *ZPriorityQueue) Push(x interface{}) {

}

func (pq *ZPriorityQueue) Pop() interface{} {
	return nil
}
