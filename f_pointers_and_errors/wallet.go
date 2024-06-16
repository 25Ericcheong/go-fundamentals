package pointers_and_errors

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Stringer Interface already defined in fmt package and essentially is called when printing using the %s format
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
