package heap

import (
	"fmt"
	"math"
)

// Interface through which heap can be accessed
type Interface interface {
	// Insert an element into heap
	Insert(element int64)
	// Max/Min element at the top of heap. This is just a lookup.
	// It will not have any side effect
	Top() (int64, error)
	// Removes the element from the top of the heap. Depending on the type of heap,
	// either the max or min most element will come to the top.
	Remove() (int64, error)
}

func min(a, b int64) bool { return a < b }
func max(a, b int64) bool { return a > b }

type implementation struct {
	comparator func(int64, int64) bool
	elements   []int64
	sentinel   int64
}

// NewMinHeap can be used to create a min heap
// A MinHeap always has the minimum element at top
func NewMinHeap() Interface {
	return &implementation{
		comparator: min,
		elements:   make([]int64, 0),
		sentinel:   math.MaxInt64,
	}
}

// NewMaxHeap can be used to create a max heap
// A MaxHeap will always have the maximum element at top
func NewMaxHeap() Interface {
	return &implementation{
		comparator: max,
		elements:   make([]int64, 0),
		sentinel:   math.MinInt64,
	}
}

func (h *implementation) Insert(element int64) {
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

func (h *implementation) Top() (int64, error) {
	if len(h.elements) < 1 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.elements[0], nil
}

func (h *implementation) Remove() (minimum int64, err error) {
	if len(h.elements) < 1 {
		return 0, fmt.Errorf("heap is empty")
	}

	// save for later returning
	minimum = h.elements[0]
	if len(h.elements) == 1 {
		h.elements = make([]int64, 0)
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
