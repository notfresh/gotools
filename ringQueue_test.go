package gotools

import "testing"

func TestRingQueue(t *testing.T) {
	q := NewRingQueue(3 + 1)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Print()

}
