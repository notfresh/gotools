package gotools

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	res1 := stack.Pop()
	res2 := stack.Pop()
	res3 := stack.Pop()
	fmt.Println(res1, res2, res3)

}
