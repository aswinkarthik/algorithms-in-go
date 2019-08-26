package stack_test

import (
	"testing"

	"github.com/aswinkarthik/algorithms-in-go/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("should do stack operations", func(t *testing.T) {
		s := stack.New()
		elements := []int{5, 2, 6, 7, 12, 6}

		for _, el := range elements {
			s.Push(el)
		}

		actual, err := s.Peek()

		assert.NoError(t, err)
		assert.Equal(t, 6, actual.(int))

		for i := len(elements) - 1; i >= 0; i-- {
			el, err := s.Pop()

			assert.NoError(t, err)
			assert.Equal(t, elements[i], el.(int))
		}

		_, err = s.Pop()
		assert.Error(t, err)

		_, err = s.Peek()
		assert.Error(t, err)
	})
}
