package queue

type UnboundedQueue[T any] struct {
	buffer []T
	head   int
}

func NewUnboundedQueue[T any]() *UnboundedQueue[T] {
	return &UnboundedQueue[T]{
		buffer: make([]T, 0),
		head:   0,
	}
}

func (q *UnboundedQueue[T]) compact() {
	n := copy(q.buffer, q.buffer[q.head:])

	clear(q.buffer[n:])

	// update slice header
	q.buffer = q.buffer[:n]
	q.head = 0
}

func (q *UnboundedQueue[T]) needShift() bool {
	return len(q.buffer) == cap(q.buffer) && q.head > 0
}

func (q *UnboundedQueue[T]) Add(val T) bool {
	if q.needShift() {
		q.compact()
	}

	q.buffer = append(q.buffer, val)
	return true
}

func (q *UnboundedQueue[T]) Peek() (val T, ok bool) {
	if q.Empty() {
		var zero T
		val = zero
		ok = false
		return
	}

	val = q.buffer[q.head]
	ok = true
	return
}

func (q *UnboundedQueue[T]) Drop() bool {
	if q.Empty() {
		return false
	}
	var zero T
	q.buffer[q.head] = zero
	q.head++
	return true
}

func (q *UnboundedQueue[T]) Len() int {
	return len(q.buffer) - q.head
}

func (q *UnboundedQueue[T]) Full() bool {
	return false
}

func (q *UnboundedQueue[T]) Empty() bool {
	return len(q.buffer) == q.head
}

func (q *UnboundedQueue[T]) Cap() int {
	return 0
}
