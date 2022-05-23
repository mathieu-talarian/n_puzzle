package solver

import (
	"sync"

	go_heaps "github.com/theodesp/go-heaps"
	rank_paring "github.com/theodesp/go-heaps/rank_pairing"
)

type open_list struct {
	list  *rank_paring.RPHeap
	mutex *sync.Mutex
}

var Open_List *open_list

func init() {

	Open_List = &open_list{
		list:  rank_paring.New(),
		mutex: &sync.Mutex{},
	}
}

func (o *open_list) Insert(el go_heaps.Item) {
	o.mutex.Lock()

	o.list.Insert(el)

	o.mutex.Unlock()
}

func (o *open_list) Size() int {
	o.mutex.Lock()

	size := o.list.Size()

	o.mutex.Unlock()
	return size
}

func (o *open_list) DeleteMin() go_heaps.Item {
	o.mutex.Lock()

	min := o.list.DeleteMin()

	o.mutex.Unlock()

	return min
}
