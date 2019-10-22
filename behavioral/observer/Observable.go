package observer

type ISubscriber interface {
	OnSubscribe()
}

type IObservable interface {
	Register()
	Notify()
}