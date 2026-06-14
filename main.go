package main

type Subscriber struct {
	name    string
	eventCh chan string
}

func main() {
	dns := Subscriber{
		"DNS",
		make(chan string),
	}

	logger := Subscriber{
		"logger",
		make(chan string),
	}

	// node publish a changeIP event

	ev := "changeIP"

	dns.eventCh <- ev
	logger.eventCh <- ev
}

// publisher --event(string)--> subscriber (by chan string)
// need to improve:
// 1. publisher highly coupled with subscriber, pub needs to manage sub life cycle
//  sol: introduce Bus as a global eventbus hub/manager/... that pub -> bus -> sub
// 2. event are just string, not using type checking
//  sol: Generics Subcriber[T]
