package adapter

import (
	"time"
)

type Payment interface {
	Pay(from string, to string, amount int64)
}

type ShopingCart struct {
	PaymentMethod Payment
}

type BankAdapter struct {
	Bank BankPayment
}

func (ba BankAdapter) Pay(from string, to string, amount int64) {
	ba.Bank.ProcessTransaction(&Transaction{
		FromAccount: from,
		ToAccount:   to,
		Amount:      amount,
		Date:        time.Now(),
		Reason:      "random aja",
	})
}

type PaypalAdapter struct {
	Paypal PaypalPayment
}

func (pa PaypalAdapter) Pay(from string, to string, amount int64) {
	pa.Paypal.SendMoney(from, to, &Money{
		Amount:   amount,
		Currency: "USD",
	})
}

