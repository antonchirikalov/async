package async

import "sync"

type Future[T any] interface {
	Done() bool
	Value() T
	Wait() T
}

type feature[T any] struct {
	sync.Mutex
	cond  *sync.Cond
	done  bool
	value T
}

func (f *feature[T]) set(value T) {
	f.Lock()
	f.value = value
	f.done = true
	f.Unlock()
	f.cond.Signal()
}

func (f *feature[T]) Value() T {
	return f.value
}

func (f *feature[T]) Done() bool {
	f.Lock()
	defer f.Unlock()
	return f.done
}

func (f *feature[T]) Wait() T {
	f.Lock()
	defer f.Unlock()
	if !f.done {
		f.cond.Wait()
	}
	return f.value
}
