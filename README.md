# eventbus

[![Go Reference](https://pkg.go.dev/badge/github.com/jrmhx/eventbus.svg)](https://pkg.go.dev/github.com/jrmhx/eventbus)
[![Go Report Card](https://goreportcard.com/badge/github.com/jrmhx/eventbus)](https://goreportcard.com/report/github.com/jrmhx/eventbus)

## Overview

Package eventbus provides a small typed in-process event bus.

This is an early v0.1.0 implementation. It is intentionally simple: events are
routed by their concrete Go type, subscribers receive values through typed
channels, and the bus uses one goroutine to deliver events in order.

## Install

```sh
go get github.com/jrmhx/eventbus
```

## Example

```go
package main

import (
	"fmt"

	"github.com/jrmhx/eventbus"
)

type UserCreated struct {
	ID string
}

func main() {
	bus := eventbus.NewBus()
	defer bus.Close()

	sub := eventbus.Subscribe(bus, UserCreated{})

	go eventbus.Publish(bus, UserCreated{ID: "u1"})

	event := <-sub.Event()
	fmt.Println(event.ID)
}
```

## Current Behavior

Subscribers are registered by event type:

```go
sub := eventbus.Subscribe(bus, UserCreated{})
```

Publishing a value sends it to subscribers registered for the same concrete
type:

```go
eventbus.Publish(bus, UserCreated{ID: "u1"})
```

Delivery is synchronous inside the bus pump. For a single subscriber, events are
received in the order the bus reads them from its publish channel.

## Limitations

This version does not isolate slow subscribers. Subscriber channels are
unbuffered, so if a subscriber does not read from its channel, the bus blocks
while delivering to it. That also blocks later publishes.

There is also no unsubscribe API, context support, per-subscriber queue, or
client abstraction yet.

## Roadmap

- [x] global topic routing
- [x] generic subscriber and publisher, add event type
- [x] generic facade + ungeneric core implement
- [ ] add client for serialization pub/sub and cascading close
- [ ] implement queue(bounded) as buffer to work with unbuffered channel for bus
- [ ] implement queue(unbounded) as buffer to work with unbuffered channel for subscriber
- [ ] emit publisher
- [ ] add ctx for bus and client
- [ ] logger for bus
- [ ] debugger for bus
- [ ] NewBusWithConfig()
