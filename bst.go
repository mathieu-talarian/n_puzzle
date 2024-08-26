package main

import "fmt"

type BinarySearchTree struct {
	Uuid  *TreeString
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

type TreeString string

// Compare func compare BstString to Item of new BST
func (a TreeString) Compare(b TreeString) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// NewBst returns *BST
func NewBst(uuid TreeString) *BinarySearchTree {
	return &BinarySearchTree{Uuid: &uuid}
}

// Insert new `Item` on BST
func (n *BinarySearchTree) Insert(data TreeString) error {
	if n == nil {
		return fmt.Errorf("cannot insert Value into a Nil tree")
	}

	switch {
	case n.Uuid.Compare(data) == 0:
		return nil
	case n.Uuid.Compare(data) > 0:
		if n.Left == nil {
			n.Left = &Bst{Uuid: &data}
			return nil
		}
		return n.Left.Insert(data)
	case n.Uuid.Compare(data) < 0:
		if n.Right == nil {
			n.Right = &Bst{Uuid: &data}
			return nil
		}
		return n.Right.Insert(data)
	}
	return nil
}

// Find `Item` on BST returns nil, false if can't find item
func (n *BinarySearchTree) Find(data TreeString) bool {
	if n == nil {
		return false
	}
	switch {
	case n.Uuid.Compare(data) == 0:
		return true
	case n.Uuid.Compare(data) > 0:
		return n.Left.Find(data)
	default:
		return n.Right.Find(data)
	}
}
