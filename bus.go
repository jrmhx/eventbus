package eventbus

import (
	"reflect"
	"sync"
)

// a global component that:
// manage event routes (map)
// using string to represent evt for the current stadge

type Bus struct {
	topicsMu sync.Mutex
	topics   map[reflect.Type][]subscriber
	write    chan any
}

func (b *Bus) Close() {
	close(b.write)
}

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

func NewBus() *Bus {
	b := &Bus{
		topics: make(map[reflect.Type][]subscriber, 0),
		write:  make(chan any),
	}
	go b.pump()
	return b
}

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
