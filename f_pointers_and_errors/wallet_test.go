package pointers_and_errors

import (
	"go-fundamentals/utils"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	utils.AssertCorrectIntMessage(t, got, want)
}
