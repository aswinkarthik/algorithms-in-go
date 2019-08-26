package stack

import "fmt"

// Interface provides methods for interaction with stack
// All return types and insert types are definted as interface{}
// Please cast to the type you want
type Interface interface {
	// Push elements into stack
	Push(interface{})
	// Pop elements out of stack
	Pop() (interface{}, error)
	// Peek what is the top most element
	Peek() (interface{}, error)
}

// New creates a new stack
func New() Interface {
	return &implementation{}
}

type implementation struct {
	top *node
}

type node struct {
	value interface{}
	next  *node
}

// Push element into stack
func (s *implementation) Push(element interface{}) {
	node := &node{value: element}
	if s.top == nil {
		s.top = node
		return
	}

	node.next = s.top
	s.top = node
}

// Pop element out of stack
func (s *implementation) Pop() (interface{}, error) {
	if s.top == nil {
		return 0, fmt.Errorf("no elements in stack")
	}
	nodeToPop := s.top
	s.top = s.top.next
	return nodeToPop.value, nil
}

// Peek the top most element in stack
func (s *implementation) Peek() (interface{}, error) {
	if s.top == nil {
		return 0, fmt.Errorf("no elements in stack")
	}
	return s.top.value, nil
}
