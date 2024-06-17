package utils

import (
	"errors"
	"slices"
	"testing"
)

func AssertCorrectErrorMessage(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

func AssertCorrectStringsMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertCorrectIntMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertCorrectFloatMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func AssertCorrectFloatMoreDecimalsMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func AssertCorrectNumbersArraysMessage(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertCorrectNumbersAndArrayMessage(t testing.TB, got, want int, arr []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d for %v", got, want, arr)
	}
}
