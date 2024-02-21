package observer

type IObserver[E any] interface {
	OnNotify(event E)
}
