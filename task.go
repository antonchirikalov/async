package async

type TaskFunc[T any] func() T

type Task[T any] struct {
	task   TaskFunc[T]
	future *feature[T]
}
