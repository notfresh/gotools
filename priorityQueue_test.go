package gotools

import (
	"container/heap"
	"fmt"
	"testing"
	"time"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

// 这个示例首先会创建一个优先队列，并在队列中包含一些元素
// 接着将一个新元素添加到队列里面，并对其进行操作
// 最后按优先级有序地移除队列中的各个元素。
func TestPriorityQueue(t *testing.T) {
	// 一些元素以及它们的优先级。
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 创建一个优先队列，并将上述元素放入到队列里面，
	// 然后对队列进行初始化以满足优先队列（堆）的不变性。
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新元素，然后修改它的优先级。
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.value, 5)

	// 以降序形式取出并打印队列中的所有元素。
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}
}

func TestPriorityQueueR(t *testing.T) {
	// 一些元素以及它们的优先级。
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 创建一个优先队列，并将上述元素放入到队列里面，
	// 然后对队列进行初始化以满足优先队列（堆）的不变性。
	pq := make(PriorityQueueR, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新元素，然后修改它的优先级。
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.value, 5)

	// 以降序形式取出并打印队列中的所有元素。
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}
}

func TestPriorityQueueRWithTime(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(time.Second * -1)
	t3 := t2.Add(time.Second * -1)
	item1 := &Item{
		"aaaa",
		int(t1.Unix()),
		1,
	}
	item2 := &Item{
		"bbb",
		int(t2.Unix()),
		2,
	}
	item3 := &Item{
		"bbb",
		int(t3.Unix()),
		1,
	}
	q := make(PriorityQueueR, 0)
	heap.Init(&q)
	heap.Push(&q, item1)
	heap.Push(&q, item2)
	heap.Push(&q, item3)
	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}

}
