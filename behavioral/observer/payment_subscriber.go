package observer

import "fmt"

type checkoutPaymentSubscriber struct {
	paymentService PaymentServiceInterface
}

func NewPaymentSubscriber() ISubscriber {
	return &checkoutPaymentSubscriber{
		paymentService: NewPaymentService(),
	}
}

func (ps checkoutPaymentSubscriber) OnSubscribe() {
	fmt.Println("FROM PAYMENT SUBSCRIBE")

	ps.paymentService.Pay()
}
