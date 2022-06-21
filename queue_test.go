package gotools

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(12)
	q.Push(15)
	q.Pop()
	q.Print()
}
