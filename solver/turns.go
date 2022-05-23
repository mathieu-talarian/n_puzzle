package solver

import "sync"

type turns struct {
	val   uint
	mutex *sync.Mutex
}

var Turns *turns

func init() {

	Turns = &turns{
		val:   0,
		mutex: &sync.Mutex{},
	}
}

func (t *turns) Inc() {
	t.mutex.Lock()
	t.val++
	t.mutex.Unlock()
}

func (t *turns) Dec() {
	t.mutex.Lock()
	t.val--
	t.mutex.Unlock()
}

func (t *turns) Val() (val uint) {
	t.mutex.Lock()
	val = t.val
	t.mutex.Unlock()
	return
}
