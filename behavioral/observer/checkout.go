package observer

import "fmt"

type CheckoutServiceInterface interface {
	Checkout()
}

type checkoutService struct {
	subs []ISubscriber
}

func NewCheckoutService() CheckoutServiceInterface {
	cs := checkoutService{
		subs: []ISubscriber{},
	}

	cs.Register(NewPaymentSubscriber())
	cs.Register(NewShipmentSubscriber())

	return &cs
}

func (s *checkoutService) Register(sub ISubscriber) {
	s.subs = append(s.subs, sub)
}

func (s *checkoutService) Notify() {
	for _, sub := range s.subs {
		sub.OnSubscribe()
	}
}

func (s *checkoutService) Checkout() {
	fmt.Println("DO CHECKOUT LOGIC")

	s.Notify()
}