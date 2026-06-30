package queue

// Queue is an container interface for events
// eventbus uses a concrete queue[Event] as channel buffer
//
// Elements are inspected using Peek and removed using Drop,
// allowing callers to acknowledge processing before removing
// an item from the queue.
type Queue[T any] interface {
	// Init a queue with capacity of n
	New(int) *Queue[T]

	// Add inserts v into the queue.
	// It returns false if the queue is full.
	Add(T) bool

	// Peek returns the next element without removing it.
	Peek() (T, bool)

	// Drop removes the element previously returned by Peek.
	Drop() bool

	// Empty reports whether the queue contains no elements.
	Empty() bool

	// Full reports whether the queue has reached capacity,
	// always return false for unbounded queue
	Full() bool

	// Len returns the number of elements currently in the queue.
	Len() int

	// Cap returns the capacity of the queue, 0 if its unbounded queue
	Cap() int
}
