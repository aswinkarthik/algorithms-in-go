package tree_test

import (
	"testing"

	"github.com/aswinkarthik/algorithms-in-go/tree"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	bst := tree.BinarySearchTree{}
	bst.Insert(10)

	assert.Equal(t, 10, bst.Root.Value)

	bst.Insert(8)
	if assert.NotNil(t, bst.Root.Left) {
		assert.Equal(t, 8, bst.Root.Left.Value)
	}

	bst.Insert(12)
	if assert.NotNil(t, bst.Root.Right) {
		assert.Equal(t, 12, bst.Root.Right.Value)
	}

	bst.Insert(14)
	if assert.NotNil(t, bst.Root.Right.Right) {
		assert.Equal(t, 14, bst.Root.Right.Right.Value)
	}

	err := bst.Insert(14)
	assert.Error(t, err)
}
