package pointers_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(10)

	// For this instance, type is Bitcoin from int
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
