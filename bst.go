package main

import "fmt"

type BinarySearchTree struct {
	Uuid  *TreeString
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

type TreeString string

// Compare func compare BstString to Item of new BST
func (treeStringA TreeString) Compare(treeStringB TreeString) int {
	if treeStringA < treeStringB {
		return -1
	}
	if treeStringA > treeStringB {
		return 1
	}
	return 0
}

// Insert new `Item` on BST
func (node *BinarySearchTree) Insert(newData TreeString) error {
	if node == nil {
		return fmt.Errorf("cannot insert Value into a Nil tree")
	}

	switch {
	case node.Uuid.Compare(newData) == 0:
		return nil
	case node.Uuid.Compare(newData) > 0:
		if node.Left == nil {
			node.Left = &BinarySearchTree{Uuid: &newData}
			return nil
		}
		return node.Left.Insert(newData)
	case node.Uuid.Compare(newData) < 0:
		if node.Right == nil {
			node.Right = &BinarySearchTree{Uuid: &newData}
			return nil
		}
		return node.Right.Insert(newData)
	}
	return nil
}

// Find `Item` on BST returns nil, false if you can't find item
func (node *BinarySearchTree) Find(newData TreeString) bool {
	if node == nil {
		return false
	}
	switch {
	case node.Uuid.Compare(newData) == 0:
		return true
	case node.Uuid.Compare(newData) > 0:
		return node.Left.Find(newData)
	default:
		return node.Right.Find(newData)
	}
}
