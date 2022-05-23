package solver

import (
	"N_Puzzle/bst"
	"sync"
)

type closedList struct {
	list  *bst.Node
	mutex *sync.Mutex
}

var Closed_List *closedList

func init() {
	Closed_List = &closedList{
		list:  nil,
		mutex: &sync.Mutex{},
	}
}

func (c *closedList) Insert(uuid string) {
	c.mutex.Lock()

	if c.list == nil {
		c.list = bst.NewNode(bst.String(uuid))
	} else {
		c.list.Insert(bst.String(uuid))
	}
	c.mutex.Unlock()
}

func (c *closedList) AlreadyClosed(node *Node) bool {
	c.mutex.Lock()

	_, ok := c.list.Find(bst.String(node.State.CreateUuid()))

	c.mutex.Unlock()
	return ok
}
