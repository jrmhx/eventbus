package eventbus

import "sync"

// a global component that:
// manage event routes (map)
// using string to represent evt for the current stadge

type Bus struct {
	topicsMu sync.Mutex
	topics   map[string][]*Subscriber
	write    chan string
}

func (b *Bus) Close() {
	close(b.write)
}

func (b *Bus) Publish(evt string) {
	b.write <- evt
}

func (b *Bus) pump() {
	for evt := range b.write {
		b.topicsMu.Lock()
		subs := b.topics[evt]
		b.topicsMu.Unlock()

		for _, sub := range subs {
			sub.read <- evt
		}
	}
}

func NewBus() *Bus {
	b := &Bus{
		topics: make(map[string][]*Subscriber, 0),
		write:  make(chan string),
	}
	go b.pump()
	return b
}

func (b *Bus) Subscribe(evt string) *Subscriber {
	s := &Subscriber{
		evt:  evt,
		read: make(chan string),
	}
	b.topicsMu.Lock()
	defer b.topicsMu.Unlock()
	b.topics[evt] = append(b.topics[evt], s)
	return s
}
