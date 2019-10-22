package observer

import "fmt"

type PaymentServiceInterface interface {
	Pay()
}

type paymentService struct {}

func NewPaymentService() PaymentServiceInterface {
	return &paymentService{}
}

func (paymentService) Pay() {
	fmt.Println("DO PAYMENT LOGIC")
}
