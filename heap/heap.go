package heap

import (
	"fmt"
)

// Interface through which heap can be accessed
type Interface interface {
	// Insert an element into heap
	Insert(element interface{})
	// Max/Min element at the top of heap. This is just a lookup.
	// It will not have any side effect
	Top() (interface{}, error)
	// Removes the element from the top of the heap. Depending on the type of heap,
	// either the max or min most element will come to the top.
	Remove() (interface{}, error)
}

// AssertHeapProperty is a function that has to qualify if heap property is met.
// If true, heap property is considered to be with-held
// If false, its not
// This is used to handle inserts and removals appropriately.
// Make sure to convert interface{} to appropriate type.
type AssertHeapProperty func(a, b interface{}) bool

type implementation struct {
	comparator AssertHeapProperty
	elements   []interface{}
}

// New can be used to instantiate a generic Heap.
// The comparator should assert that the heap property is met.
func New(comparator AssertHeapProperty) Interface {
	return &implementation{
		comparator: comparator,
		elements:   make([]interface{}, 0),
	}
}

// NewMinHeapInt64 can be used to create a min heap of int64 elements
// The interface does not support type. Hence Top and remove should cast appropriately.
// It uses MinHeapInt64 as the comparator
func NewMinHeapInt64() Interface {
	return New(MinHeapInt64)
}

// NewMaxHeapInt64 can be used to create a max heap of int64 elements
// The interface does not support type. Hence Top and remove should cast appropriately.
// It uses MaxHeapInt64 as the comparator
func NewMaxHeapInt64() Interface {
	return New(MaxHeapInt64)
}

func (h *implementation) Insert(element interface{}) {
	h.elements = append(h.elements, element)

	for index := len(h.elements) - 1; index > 0; index = parentOf(index) {

		parentIndex := parentOf(index)

		// If heap property is met, break
		if !h.comparator(h.elements[index], h.elements[parentIndex]) {
			break
		}

		// swap elements
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
	}
}

func (h *implementation) Top() (interface{}, error) {
	if len(h.elements) < 1 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.elements[0], nil
}

func (h *implementation) Remove() (minimum interface{}, err error) {
	if len(h.elements) < 1 {
		return 0, fmt.Errorf("heap is empty")
	}

	// save for later returning
	minimum = h.elements[0]
	if len(h.elements) == 1 {
		h.elements = make([]interface{}, 0)
		return
	}

	// bring up the last element top
	h.elements[0] = h.elements[len(h.elements)-1]

	// reduce length or slice by 1
	h.elements = h.elements[0 : len(h.elements)-1]

	for index := 0; index < len(h.elements); {
		if h.hasRightChild(index) {
			// both children exist

			right := h.elements[rightOf(index)]
			left := h.elements[leftOf(index)]

			// select either left or right node that will maintain heap property
			selectedIndex := leftOf(index)
			if !h.comparator(left, right) {
				selectedIndex = rightOf(index)
			}

			// Check if heap property is met
			if !h.comparator(h.elements[index], h.elements[selectedIndex]) {
				h.elements[index], h.elements[selectedIndex] = h.elements[selectedIndex], h.elements[index]
				index = selectedIndex
				continue
			}
		} else if h.hasLeftChild(index) {
			// only left child exists
			selectedIndex := leftOf(index)

			// Check if heap property is met
			if !h.comparator(h.elements[index], h.elements[selectedIndex]) {
				h.elements[index], h.elements[selectedIndex] = h.elements[selectedIndex], h.elements[index]
				index = selectedIndex
				continue
			}
		}
		break
	}

	return minimum, err
}

func parentOf(pos int) int {
	if pos == 0 {
		return pos
	}
	return (pos - 1) / 2
}

func leftOf(pos int) int {
	return 2*pos + 1
}

func rightOf(pos int) int {
	return 2*pos + 2
}

func (h *implementation) hasLeftChild(pos int) bool {
	return leftOf(pos) < len(h.elements)
}

func (h *implementation) hasRightChild(pos int) bool {
	return rightOf(pos) < len(h.elements)
}
