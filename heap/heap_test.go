package heap_test

import (
	"sort"
	"testing"

	"github.com/aswinkarthik/algorithms-in-go/heap"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	t.Run("should perform heap operations on a min heap", func(t *testing.T) {
		h := heap.NewMinHeap()

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
			assert.Equal(t, el, actual)
		}
	})

	t.Run("should perform heap operations on a max heap", func(t *testing.T) {
		h := heap.NewMaxHeap()

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
		assert.Equal(t, sortedElements[0], actual)

		for _, el := range sortedElements {
			actual, err := h.Remove()

			assert.NoError(t, err)
			assert.Equal(t, el, actual)
		}
	})

}
