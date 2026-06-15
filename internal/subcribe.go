package eventbus

type Subscriber struct {
	evt  string
	read chan string
}

func (s *Subscriber) Event() <-chan string {
	return s.read
}
