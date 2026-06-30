package queue

// bounded queue with circular buffer

type BoundedQueue[T any] struct {
	buffer   []T
	capacity int
	head     int
	tail     int
	size     int
}

func NewBoundedQueue[T any](cap int) *BoundedQueue[T] {
	if cap <= 0 {
		panic("Bounded queue capacity has to be large than 0")
	}
	return &BoundedQueue[T]{
		buffer:   make([]T, cap),
		capacity: cap,
		head:     0,
		tail:     0,
		size:     0,
	}
}

func (q *BoundedQueue[T]) Add(val T) bool {
	if q.Full() {
		return false
	}
	q.buffer[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.size += 1
	return true
}
func (q *BoundedQueue[T]) Peek() (val T, ok bool) {
	var zero T
	val = zero
	ok = false
	if !q.Empty() {
		val = q.buffer[q.head]
		ok = true
	}
	return
}
func (q *BoundedQueue[T]) Drop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % q.capacity
	q.size -= 1
	return true
}
func (q *BoundedQueue[T]) Len() int {
	return q.size
}

func (q *BoundedQueue[T]) Full() bool {
	return q.size == q.capacity
}

func (q *BoundedQueue[T]) Empty() bool {
	return q.size == 0
}

func (q *BoundedQueue[T]) Cap() int {
	return q.capacity
}
