package async

import "sync"

type Pool[T any] interface {
	// Async schedules a task for asynchronous execution.
	Async(groupName string, task TaskFunc[T]) Future[T]
	// Close the port for any future tasks executions.
	Close()
}

type pool[T any] struct {
	ch   map[string]*channel[T]
	size int
}

func NewPool[T any](size int) Pool[T] {
	return &pool[T]{
		size: size,
		ch:   make(map[string]*channel[T]),
	}
}

func (p *pool[T]) Async(groupName string, taskFunc TaskFunc[T]) Future[T] {
	_, ok := p.ch[groupName]
	if !ok {
		p.ch[groupName] = newChannel[T](p.size)
	}

	future := new(feature[T])
	future.cond = sync.NewCond(future)

	task := &Task[T]{
		task:   taskFunc,
		future: future,
	}

	p.ch[groupName].ch <- task
	return task.future
}

func (p *pool[T]) Close() {
	for _, ch := range p.ch {
		ch.close()
	}
}
