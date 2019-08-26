package queue_test

import (
	"testing"

	"github.com/aswinkarthik/algorithms-in-go/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("should allow queue operations", func(t *testing.T) {
		elements := []int{5, 2, 6, 7, 12, 6}
		q := queue.New()
		for _, el := range elements {
			q.Enqueue(el)
		}

		actual, err := q.First()

		assert.NoError(t, err)
		assert.Equal(t, 5, actual.(int))

		for _, expected := range elements {
			val, err := q.Dequeue()
			assert.NoError(t, err)
			assert.Equal(t, expected, val.(int))
		}

		_, err = q.Dequeue()
		assert.Error(t, err)
	})
}
