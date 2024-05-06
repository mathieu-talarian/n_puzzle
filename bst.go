package main

import "fmt"

// Bst binary search tree struct
type Bst struct {
	Uuid  *BstString
	Left  *Bst
	Right *Bst
}

// BstString implements the Item interface
type BstString string

// Compare func compare BstString to Item of new BST
func (a BstString) Compare(b BstString) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// NewBst returns *BST
func NewBst(uuid BstString) *Bst {
	return &Bst{Uuid: &uuid}
}

// Insert new `Item` on BST
func (n *Bst) Insert(data BstString) error {
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
func (n *Bst) Find(data BstString) bool {
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
