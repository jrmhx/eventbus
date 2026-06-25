package queue

// circular queue

type FIFOQueue[T any] struct {
	buffer   []T
	capacity int
	head     int
	tail     int
	size     int
}

func New[T any](n int) *FIFOQueue[T] {
	return &FIFOQueue[T]{
		buffer:   make([]T, n),
		capacity: n,
		head:     0,
		tail:     0,
		size:     0,
	}
}

func (q *FIFOQueue[T]) Add(val T) bool {
	if q.Full() {
		return false
	}
	q.buffer[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.size += 1
	return true
}
func (q *FIFOQueue[T]) Peek() (val T, ok bool) {
	ok = false
	if !q.Empty() {
		val = q.buffer[q.head]
		ok = true
	}
	return
}
func (q *FIFOQueue[T]) Drop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % q.capacity
	q.size -= 1
	return true
}
func (q *FIFOQueue[T]) Len() int {
	return q.size
}

func (q *FIFOQueue[T]) Full() bool {
	return q.size == q.capacity
}

func (q *FIFOQueue[T]) Empty() bool {
	return q.size == 0
}
