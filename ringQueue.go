// 环形队列
// 设计思路见 https://zhuanlan.zhihu.com/p/23948886
package gotools

import (
	"fmt"
)

type RingQueue struct {
	list  []interface{}
	front int
	end   int
	size  int
}

func NewRingQueue(size int) *RingQueue {
	rq := &RingQueue{}
	rq.list = make([]interface{}, size) // 多分配一格，用来检测头尾
	rq.size = size
	rq.front = 0
	rq.end = 0
	return rq
}

func (q *RingQueue) Front(data interface{}) {
	if q.front == q.end {
		data = nil
	}
	data = q.list[q.front]
}

func (q *RingQueue) Tail(data interface{}) {
	if q.front == q.end {
		data = nil
	}
	data = q.list[q.end]
}

func (q *RingQueue) adjustIdx(idx *int) {
	if *idx == q.size {
		*idx = 0
	}
}

func (q *RingQueue) Push(data interface{}) {
	if (q.end+1)%q.size == q.front { // 满了
		fmt.Println("队列已满，清除头部元素")
		q.front++ // 这一步不能漏，否则会出现问题
		q.adjustIdx(&q.front)
	}
	q.list[q.end] = data
	q.end++
	q.adjustIdx(&q.end)
}

func (q *RingQueue) Pop() (data interface{}) {
	if q.front == q.end {
		return
	}
	data = q.list[q.front]
	q.front++
	q.adjustIdx(&q.front)
	return
}

func (q *RingQueue) Print() {
	idx := q.front
	fmt.Println("开始遍历环形队列...")
	for {
		if idx == q.end {
			break
		}
		fmt.Println("idx", q.list[idx])
		idx++
		q.adjustIdx(&idx)
	}
	fmt.Println("结束遍历环形队列...")
}
