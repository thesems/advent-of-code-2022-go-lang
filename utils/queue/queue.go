package queue

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	data := []T{}
	queue := &Queue[T]{data}
	return queue
}

func (q *Queue[T]) Add(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Get() T {
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}
