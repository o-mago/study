package queue

import "container/list"

type ContainerListQueue[T any] struct {
	containerList list.List
}

func NewContainerListQueue[T any]() Queue[T] {
	return &ContainerListQueue[T]{}
}

func (q *ContainerListQueue[T]) Len() int {
	return q.containerList.Len()
}

func (q *ContainerListQueue[T]) Enqueue(value T) error {
	q.containerList.PushBack(value)
	return nil
}

func (q *ContainerListQueue[T]) Dequeue() (T, error) {
	elem := q.containerList.Front()
	if elem == nil {
		var empty T
		return empty, ErrQueueEmpty
	}

	q.containerList.Remove(elem)
	return elem.Value.(T), nil
}
