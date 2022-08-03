package gotools

import "fmt"

type Queue struct {
	list []interface{}
}

func NewQueue() *Queue {
	list := make([]interface{}, 0)
	return &Queue{list: list}
}

func (q *Queue) Push(data interface{}) {
	q.list = append(q.list, data)
}

func (q *Queue) Pop() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	res := q.list[0]
	q.list = q.list[1:]
	return res
}

func (q *Queue) Size() int {
	return len(q.list)
}

func (q *Queue) Print() {
	for _, val := range q.list {
		fmt.Println(val)
	}
}

//
//————————————————
//原文作者：littlexiaoshuishui
//转自链接：https://learnku.com/articles/45975
//版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。
