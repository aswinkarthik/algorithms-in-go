package queue

import "fmt"

// Interface methods exposed by queue
// All return types and insert types are definted as interface{}
// Please cast to the type you want
type Interface interface {
	Enqueue(interface{})
	Dequeue() (interface{}, error)
	First() (interface{}, error)
}

// New to create a new queue
func New() Interface {
	return &implementation{}
}

type node struct {
	value interface{}
	next  *node
}

type implementation struct {
	begin *node
	end   *node
}

// Enqueue elements into queue
func (q *implementation) Enqueue(element interface{}) {
	node := &node{value: element}

	if q.begin != nil && q.end != nil {
		q.end.next = node
		q.end = node
	} else if q.end == nil && q.begin == nil {
		q.begin = node
		q.end = node
	} else {
		panic("queue is in bad state")
	}
}

// Dequeue elements from queue
func (q *implementation) Dequeue() (interface{}, error) {
	if q.begin == nil || q.end == nil {
		return 0, fmt.Errorf("no elements")
	}
	nodeToDequeue := q.begin
	q.begin = q.begin.next
	return nodeToDequeue.value, nil
}

// First will return the first element in queue
func (q *implementation) First() (interface{}, error) {
	if q.begin != nil {
		return q.begin.value, nil
	}
	return 0, fmt.Errorf("no elements")
}
