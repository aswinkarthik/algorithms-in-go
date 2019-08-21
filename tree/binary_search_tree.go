package tree

import "fmt"

// BinarySearchTree is an implementation of BST
type BinarySearchTree struct {
	Root *Node
}

// Node represents each node in a Binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert an element into BST
func (t *BinarySearchTree) Insert(element int) (err error) {
	t.Root, err = insert(t.Root, element)
	return err
}

func insert(node *Node, element int) (*Node, error) {
	var err error
	if node == nil {
		node = &Node{Value: element}
		return node, err
	}

	if element == node.Value {
		return node, fmt.Errorf("node already exists")
	}

	if element < node.Value {
		node.Left, err = insert(node.Left, element)
	}

	if element > node.Value {
		node.Right, err = insert(node.Right, element)
	}

	return node, err
}
