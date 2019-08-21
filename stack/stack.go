package stack

import "fmt"

// Interface provides methods for interaction with stack
type Interface interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
}

// New creates a new stack
func New() Interface {
	return &intStack{}
}

type intStack struct {
	top *node
}

type node struct {
	value int
	next  *node
}

// Push element into stack
func (s *intStack) Push(element int) {
	node := &node{value: element}
	if s.top == nil {
		s.top = node
		return
	}

	node.next = s.top
	s.top = node
}

// Pop element out of stack
func (s *intStack) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("no elements in stack")
	}
	nodeToPop := s.top
	s.top = s.top.next
	return nodeToPop.value, nil
}

// Peek the top most element in stack
func (s *intStack) Peek() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("no elements in stack")
	}
	return s.top.value, nil
}
