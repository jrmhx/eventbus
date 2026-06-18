package eventbus

import (
	"reflect"
	"sync"
)

// Bus routes published events to subscribers of the same concrete type.
type Bus struct {
	topicsMu sync.Mutex
	topics   map[reflect.Type][]subscriber
	write    chan any
}

// Close stops the bus from accepting new events.
func (b *Bus) Close() {
	close(b.write)
}

// Publish sends evt to subscribers registered for its concrete type.
func Publish[E any](b *Bus, evt E) {
	b.write <- evt
}

func (b *Bus) pump() {
	for evt := range b.write {
		b.topicsMu.Lock()
		subs := b.topics[reflect.TypeOf(evt)]
		b.topicsMu.Unlock()

		for _, sub := range subs {
			sub.dispatch(evt)
		}
	}
}

// NewBus creates a bus and starts its routing goroutine.
func NewBus() *Bus {
	b := &Bus{
		topics: make(map[reflect.Type][]subscriber, 0),
		write:  make(chan any),
	}
	go b.pump()
	return b
}

// Subscribe registers a subscriber for events of type E.
func Subscribe[E any](b *Bus, evt E) *Subscriber[E] {
	s := &Subscriber[E]{
		typ:  reflect.TypeFor[E](),
		read: make(chan E),
	}
	b.topicsMu.Lock()
	defer b.topicsMu.Unlock()
	b.topics[s.typ] = append(b.topics[s.typ], s)
	return s
}
