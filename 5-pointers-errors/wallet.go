package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// this is a global value within package
var ErrorInsufficientFunds = errors.New("amount wanting to withdraw is more than total balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	// error is a type of interface and can be nil in Go
	if amount > w.balance {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}
