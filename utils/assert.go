package utils

import "testing"

func AssertCorrectStringsMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertCorrectNumbersMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertCorrectNumbersAndArrayMessage(t testing.TB, got, want int, arr []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d for %v", got, want, arr)
	}
}
