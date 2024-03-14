package async

type channel[T any] struct {
	ch chan *Task[T]
}

func newChannel[T any](size int) *channel[T] {
	ch := &channel[T]{
		ch: make(chan *Task[T], size),
	}
	go ch.run()
	return ch
}

func (c *channel[T]) run() {
	for task := range c.ch {
		task.future.set(task.task())
	}
}

func (c *channel[T]) close() {
	close(c.ch)
}
