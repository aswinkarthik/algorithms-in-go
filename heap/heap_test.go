package heap_test

import (
	"sort"
	"testing"

	"github.com/aswinkarthik/algorithms-in-go/heap"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	t.Run("should perform heap operations on a min heap", func(t *testing.T) {
		h := heap.NewMinHeapInt64()

		elements := []int64{3, 1, 6, 5, 2, 4}
		sortedElements := make([]int64, len(elements))
		copy(sortedElements, elements)
		sort.SliceStable(sortedElements, func(i, j int) bool {
			return sortedElements[i] < sortedElements[j]
		})

		actual, err := h.Top()
		assert.Error(t, err)

		for _, el := range elements {
			h.Insert(el)
		}

		actual, err = h.Top()
		assert.NoError(t, err)
		assert.Equal(t, sortedElements[0], actual)

		for _, el := range sortedElements {
			actual, err := h.Remove()

			assert.NoError(t, err)
			assert.Equal(t, el, actual.(int64))
		}
	})

	t.Run("should perform heap operations on a max heap", func(t *testing.T) {
		h := heap.NewMaxHeapInt64()

		elements := []int64{3, 1, 6, 5, 2, 4}
		sortedElements := make([]int64, len(elements))
		copy(sortedElements, elements)
		sort.SliceStable(sortedElements, func(i, j int) bool {
			return sortedElements[i] > sortedElements[j]
		})

		actual, err := h.Top()
		assert.Error(t, err)

		for _, el := range elements {
			h.Insert(el)
		}

		actual, err = h.Top()
		assert.NoError(t, err)
		assert.Equal(t, sortedElements[0], actual.(int64))

		for _, el := range sortedElements {
			actual, err := h.Remove()

			assert.NoError(t, err)
			assert.Equal(t, el, actual.(int64))
		}
	})

	t.Run("should support a custom type to simulate a priority queue", func(t *testing.T) {
		type myStruct struct {
			value int
			label string
		}

		comparator := func(a, b interface{}) bool {
			return a.(myStruct).value < b.(myStruct).value
		}

		first := myStruct{5, "a"}
		second := myStruct{8, "b"}
		third := myStruct{3, "c"}

		h := heap.New(comparator)

		h.Insert(first)
		h.Insert(second)
		h.Insert(third)

		val, err := h.Top()

		assert.NoError(t, err)
		assert.Equal(t, third, val.(myStruct))

		val, err = h.Remove()

		assert.NoError(t, err)
		assert.Equal(t, third, val.(myStruct))

		val, err = h.Remove()

		assert.NoError(t, err)
		assert.Equal(t, first, val.(myStruct))

		val, err = h.Remove()

		assert.NoError(t, err)
		assert.Equal(t, second, val.(myStruct))

	})
}
