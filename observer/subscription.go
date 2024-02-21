package observer

type Subscription[E any] struct {
	observerList []IObserver[E]
}

func (s Subscription[E]) Register(obs IObserver[E]) {
	s.observerList = append(s.observerList, obs)
}

func (s Subscription[E]) NotifyAll(event E) {
	for _, observer := range s.observerList {
		observer.OnNotify(event)
	}
}
