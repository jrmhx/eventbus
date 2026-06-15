package main

import (
	"fmt"

	eventbus "github.com/jrmhx/eventbus/internal"
)

func main() {
	bus := eventbus.NewBus()
	sub := bus.Subscribe("123")

	go bus.Publish("123")
	go bus.Publish("456")
	go bus.Publish("123")

	evt1 := <-sub.Event()
	fmt.Println(evt1)

	ev2 := <-sub.Event()
	fmt.Println(ev2)
}
