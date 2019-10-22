package adapter

import (
	"fmt"
	"time"
)

type BankPayment struct {}

type Transaction struct {
	FromAccount string
	ToAccount   string
	Amount      int64
	Date        time.Time
	Reason      string
}

func (*BankPayment) ProcessTransaction(t *Transaction) {
	fmt.Printf("[BANK - %v] PROCESS TRANSACTION %s to %s %d (%s)", t.Date, t.FromAccount, t.ToAccount, t.Amount, t.Reason)
}

