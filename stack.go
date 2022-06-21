package gotools

import "fmt"

type Stack struct {
	list []interface{}
}

func NewStack() *Stack {
	return new(Stack)
}
func (s *Stack) Push(value interface{}) {
	s.list = append(s.list, value)
}

func (s *Stack) Pop() interface{} {
	len := len(s.list)
	var value interface{}
	if len == 0 {
		return nil
	} else if len == 1 {
		value = s.list[0]
		s.list = []interface{}{}
	} else {
		value = s.list[len-1]
		s.list = s.list[len-2 : len-1]
	}
	return value
}
func (s *Stack) Print() {
	for _, v := range s.list {
		fmt.Printf("%v", v)
	}
	fmt.Printf("\n")
}
