package queue

type ChannelQueue[T any] struct {
	channel chan T
}

func NewChannelQueue[T any](capacity int) Queue[T] {
	return &ChannelQueue[T]{
		channel: make(chan T, capacity),
	}
}

func (q *ChannelQueue[T]) Len() int {
	return len(q.channel)
}

func (q *ChannelQueue[T]) Enqueue(value T) error {
	select {
	case q.channel <- value:
		return nil
	default:
		return ErrQueueFull
	}
}

func (q *ChannelQueue[T]) Dequeue() (T, error) {
	select {
	case val := <-q.channel:
		return val, nil
	default:
		var empty T
		return empty, ErrQueueEmpty
	}
}
