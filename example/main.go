package main

import (
	"fmt"

	eventbus "github.com/jrmhx/eventbus"
)

type EventA struct {
	a int
}

type EventB struct {
	b int
}

func main() {
	bus := eventbus.NewBus()
	sub := eventbus.Subscribe(bus, EventA{})

	go eventbus.Publish(bus, EventA{1})
	go eventbus.Publish(bus, EventB{2})
	go eventbus.Publish(bus, EventA{3})

	evt1 := <-sub.Event()
	fmt.Println(evt1)

	ev2 := <-sub.Event()
	fmt.Println(ev2)
}
