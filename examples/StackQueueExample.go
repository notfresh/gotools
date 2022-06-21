package main

import (
	"fmt"
	zx "github.com/notfresh/gotools"
)

func TestStack() {
	stack := zx.NewStack()
	stack.Push(1)
	stack.Push(2)
	res1 := stack.Pop()
	res2 := stack.Pop()
	res3 := stack.Pop()
	fmt.Println(res1, res2, res3)

}

func TestQueue() {
	q := zx.NewQueue()
	q.Push(12)
	q.Push(15)
	q.Pop()
	q.Print()
}

func main() {
	TestStack()
	TestQueue()
}
