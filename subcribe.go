package eventbus

import "reflect"

type subscriber interface {
	subscribeType() reflect.Type
	dispatch(evt any)
}

type Subscriber[E any] struct {
	typ  reflect.Type
	read chan E
}

func (s *Subscriber[E]) subscribeType() reflect.Type {
	return s.typ
}

func (s *Subscriber[E]) dispatch(evt any) {
	v, ok := evt.(E)
	if !ok {
		return
	}

	s.read <- v
}

func (s *Subscriber[E]) Event() <-chan E {
	return s.read
}
