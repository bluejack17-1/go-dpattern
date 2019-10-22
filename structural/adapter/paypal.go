package adapter

import "fmt"

type Money struct {
	Amount   int64
	Currency string
}

type PaypalPayment struct {}

func (*PaypalPayment) SendMoney(from string, to string, amount *Money) {
	fmt.Printf("[PAYPAL] SEND MONEY FROM %s to %s (%d %s)", from, to, amount.Amount, amount.Currency)
}
