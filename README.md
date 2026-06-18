# EventBus
[![Go Reference](https://pkg.go.dev/badge/github.com/jrmhx/eventbus.svg)](https://pkg.go.dev/github.com/jrmhx/eventbus)
[![Go Report Card](https://goreportcard.com/badge/github.com/jrmhx/eventbus)](https://goreportcard.com/report/github.com/jrmhx/eventbus)

## TODOs

- [x] global topic routing
- [x] generic subscriber and publisher, add event type
- [x] generic facade + ungeneric core implement
- [ ] add client for serilization pub/sub and cascading close
- [ ] implement queue(bounded) as buffer to work with unbuffered channel for bus 
- [ ] implement queue(unbounded) as buffer to work with unbuffered channel for subscriber
- [ ] emit publisher
- [ ] add ctx for bus and client
- [ ] logger for bus
- [ ] debugger for bus
- [ ] NewBusWithConfig()
