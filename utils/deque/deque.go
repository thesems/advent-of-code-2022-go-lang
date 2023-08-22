package deque

import "errors"

type Deque[T any] struct {
    items []T
}

func New[T any]() *Deque[T] {
    deque := &Deque[T]{make([]T, 0)}
    return deque
}

func (d *Deque[T]) Push(item T) {
    d.items = append(d.items, item)
}

func (d *Deque[T]) Front() (*T, error) {
    if len(d.items) == 0 {
        return nil, errors.New("queue is empty")
    }

    item := d.items[0]
    d.items = d.items[1:]
    return &item, nil
}

func (d *Deque[T]) IsEmpty() bool {
    return len(d.items) == 0
}
